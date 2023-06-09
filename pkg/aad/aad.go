package aad

import (
	"AADDecryptor/pkg/cli"
	"AADDecryptor/pkg/utils"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/wat4r/dpapitk/cngblob"
	"github.com/wat4r/dpapitk/masterkey"
)

var OutputData string

var Result = `[+] AppId: %s
[+] UserName: %s
[+] Refresh token: 
%s
[+] Access token:
%s
`

func readMasterKeyGuid(blobData []byte) string {
	var guidMasterKeyData []byte
	var guidMasterKey [16]byte

	blobDataSplit := bytes.Split(blobData, []byte{0x00, 0x00, 0x00})
	if len(blobDataSplit) >= 2 {
		guidMasterKeyData = blobDataSplit[2]
	}

	err := binary.Read(bytes.NewReader(guidMasterKeyData), binary.LittleEndian, &guidMasterKey)
	if err != nil {
		fmt.Println("[-] " + err.Error())
		os.Exit(0)
	}
	masterKeyGuid := utils.GuidMasterKeyConvert(guidMasterKey)
	return masterKeyGuid
}

func DecryptDefFiles(opt *cli.Options, defFiles []string) {
	var masterKeys map[string][]byte

	for _, defFilePath := range defFiles {
		msg(strings.Repeat("-", 46))
		msg("[+] Def file: " + defFilePath)

		// Decode
		defFileData := utils.ReadFile(defFilePath)
		b64Data := defFileData[6:]
		blobData := utils.Base64Decode(b64Data)

		// Get master key file name
		masterKeyFileName := readMasterKeyGuid(blobData)
		msg("[+] Master key file: " + masterKeyFileName)

		masterKey, exists := masterKeys[masterKeyFileName]
		if !exists {
			masterKeyFilePath := utils.FindFile(opt.MasterKeyFiles, masterKeyFileName)
			if masterKeyFilePath == "" {
				msg(fmt.Sprintf("[-] master key file `%s` not found!", masterKeyFileName))
				continue
			}

			// Check user SID
			if opt.Sid == "" {
				absPath, _ := filepath.Abs(masterKeyFilePath)
				folderPath := filepath.Dir(absPath)
				masterKeyFolder := filepath.Base(folderPath)
				if strings.HasPrefix(masterKeyFolder, "S-1-5-21") {
					opt.Sid = masterKeyFolder
				}
			}
			if opt.Sid == "" {
				fmt.Println("[-] User sid not found!")
				os.Exit(0)
			}

			// Decrypt master key file
			masterKeyFileBytes := utils.ReadFile(masterKeyFilePath)
			masterKeyFile := masterkey.InitMasterKeyFile(masterKeyFileBytes)
			if opt.Password != "" {
				masterKeyFile.DecryptWithPassword(opt.Sid, opt.Password)
			} else if opt.Hash != "" {
				masterKeyFile.DecryptWithHash(opt.Sid, opt.Hash)
			} else if opt.DomainBackupKey != "" {
				domainBackupKey := utils.ReadFile(opt.DomainBackupKey)
				masterKeyFile.DecryptWithPvk(domainBackupKey)
			} else {
				fmt.Println("[-] Password or hash or domain backup key not found!")
				os.Exit(0)
			}
			if !masterKeyFile.Decrypted {
				msg("[-] Master key decrypt failed!")
				continue
			}
			masterKey = masterKeyFile.Key
		}

		result, err := cngblob.DecryptWithMasterKey(blobData, masterKey, nil)
		if err != nil {
			msg("[-] " + err.Error())
			continue
		}

		if len(result) == 0 {
			msg("[-] Decrypt failed!")
			continue
		}

		outputData := utils.ZlibUnCompress(result[4:])
		parseDefFileData(string(outputData))
	}

	if opt.OutputFile != "" {
		output(opt.OutputFile)
	}
}

func parseDefFileData(data string) {
	reAppId := regexp.MustCompile(`\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`)
	reUserName := regexp.MustCompile(`[\w\-\.\_]+@[\w\-\.\_]+\$`)
	reRefreshToken := regexp.MustCompile(`0\.AV[\d\-\.\_a-zA-Z]*`)
	reAccessToken := regexp.MustCompile(`eyJ0eXA[\d\-\.\_a-zA-Z]*`)
	appId := reAppId.FindString(data)
	userName := reUserName.FindString(data)
	refreshToken := reRefreshToken.FindString(data)
	accessToken := reAccessToken.FindString(data)

	output := fmt.Sprintf(Result, appId+getAppName(appId), userName[:len(userName)-1], refreshToken, accessToken)
	msg(output)
}

func msg(message string) {
	OutputData += message + "\n"
	fmt.Println(message)
}

func output(outputFile string) {
	status := utils.WriteFile(outputFile, []byte(OutputData))
	if status {
		fmt.Printf("[+] Output save to: %s\n", outputFile)
	}
}
