# drive-encrypt

Tool for encrypting files. Usage is through CLI (later to be expanded to Web)

### Installation
Find the suitable binary for your OS from the [releases.](https://github.com/crossphoton/drive-encrypt/releases)

## Usage
```
Usage:
  drive-encrypt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ls          list files with path
  pull        Decrypt a file
  push        encrypt file
  rm          Remove a file
  serve       Start the HTTP file server

Flags:
  -h, --help              help for drive-encrypt
  -p, --password string   Password

Use "drive-encrypt [command] --help" for more information about a command.
```

### Notes:
```
- You're not supposed to add the executable to your path. All the operations are done in the same directory as the binary.
- If you don't want to pass the password everytime, use alias `alias drive-encrypt=./drive-encrypt -p <password>`
```
