package cli

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"os"
)

type Options struct {
	// AAD
	DefFiles string

	// DPAPI
	MasterKeyFiles  string
	Sid             string
	Password        string
	Hash            string // sha1 or ntlm
	DomainBackupKey string

	// Output
	OutputFile string
}

func ParseParameter() *Options {
	opt := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Windows AAD BrokerPlugin def files decryption utility.")

	flagSet.CreateGroup("INPUT", "INPUT",
		flagSet.StringVarP(&opt.DefFiles, "def-files", "df", "", "AAD BrokerPlugin def files folder path"),
	)

	flagSet.CreateGroup("DPAPI", "DPAPI",
		flagSet.StringVarP(&opt.MasterKeyFiles, "master-key-files", "mkf", "", "master key files folder"),
		flagSet.StringVarP(&opt.Sid, "sid", "s", "", "user sid"),
		flagSet.StringVarP(&opt.Password, "password", "p", "", "password"),
		flagSet.StringVar(&opt.Hash, "hash", "", "sha1 hash or ntlm hash"),
		flagSet.StringVar(&opt.DomainBackupKey, "pvk", "", "domain backup key (.pvk)"),
	)

	flagSet.CreateGroup("OUTPUT", "OUTPUT",
		flagSet.StringVarP(&opt.OutputFile, "output", "o", "", "output path"),
	)

	if err := flagSet.Parse(); err != nil {
		fmt.Printf("[-] Could not parse flags: %s\n", err)
		os.Exit(0)
	}

	// Check input
	if opt.DefFiles == "" {
		fmt.Println("[-] Def files path is required!")
		os.Exit(0)
	}
	if opt.MasterKeyFiles == "" {
		fmt.Println("[-] Master key files path is required!")
		os.Exit(0)
	}
	if opt.Password == "" && opt.Hash == "" && opt.DomainBackupKey == "" {
		fmt.Println("[-] Password or hash or domain backup key is required!")
		os.Exit(0)
	}

	return opt
}
