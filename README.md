# AADDecryptor
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/release/wat4r/AADDecryptor)](https://github.com/wat4r/AADDecryptor/releases)


## Introduction
AADDecryptor is a tool for windows AAD BrokerPlugin def files decryption.


## Features
 - Support offline decryption
 - Support multiple ways of decrypting the Master key (password, hash, domain backup key)
 - Output: appid, app name, username, refresh token, access token

## Usage
```sh
AADDecryptor -h
```
This will display help for the tool. Here are all the switches it supports.

```yaml
______                           _
|  _  \    AAD BrokerPlugin     | |
| | | |___  ___ _ __ _   _ _ __ | |_ ___  _ __
| | | / _ \/ __| '__| | | | '_ \| __/ _ \| '__|
| |/ /  __/ (__| |  | |_| | |_) | || (_) | |
|___/ \___|\___|_|   \__, | .__/ \__\___/|_|
                      __/ | |
                     |___/|_|      v1.0.0

Windows AAD BrokerPlugin def files decryption utility.

Usage:
  AADDecryptor [flags]

Flags:
INPUT:
   -df, -def-files string  AAD BrokerPlugin def files folder path

DPAPI:
   -mkf, -master-key-files string  master key files folder
   -s, -sid string                 user sid
   -p, -password string            password
   -hash string                    sha1 hash or ntlm hash
   -pvk string                     domain backup key (.pvk)

OUTPUT:
   -o, -output string  output path
```

## Example
### Directory structure
```yaml
+---aad
|       a_k0d3p7ef9etr0bab3ndnhosd.def
|
+---masterKeyFolder
        +---S-1-5-21-1099483827-325504281-218701502-1001
        |       cd8b97f2-c875-4e9e-85d8-bb8d73954e50
        |       Preferred
```


### Decrypt def files
```sh
AADDecryptor -df ./aad -mkf ./masterKeyFolder -p password123 -o output.txt
```


## License
This project is licensed under the [Apache 2.0 license](LICENSE).


## Contact
If you have any issues or feature requests, please contact us. PR is welcomed.
 - https://github.com/wat4r/AADDecryptor/issues

