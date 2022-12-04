# creator_suite
## cryptool

### creator_suite command tool `cryptool`
A command-line applications to encrypt or decrypt your data.
`AES-GCM` can be used to encrypt/decrypt files, backups or even large object.

### Install
```
$ git clone https://github.com/nextdotid/creator_suite
$ cd creator_suite && make install-tool
```


### Help
Help command:
```
$ ./cryptool help

NAME:
   cryptool - Crypto Tools for CreatorSuite

USAGE:
   cryptool [global options] command [command options] [arguments...]

VERSION:
   1.0.0.rc1

COMMANDS:
   encrypt, encrypt, en  encrypt a file
   decrypt, decrypt, de  decrypt a file
   help, h               Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```


### Test

**Case 1**
Encrypt a file use `AES-GCM`
```
$ ./cryptool encrypt --in README.md --out README.md.enc
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
  ▸ AES
    ECC

✔ AES
Enter Password: ***********************(128 bit)
```

Decrypt a file use `AES-GCM`
```
$ ./cryptool decrypt --in README.md.enc --out README_dec.md
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
  ▸ AES
    ECC

✔ AES
Enter Password: ***********************(128 bit)
```

**Case 2**
Encrypt a file use `AES-GCM`
```
$ ./cryptool encrypt --in test-nextdotid.png --out test-nextdotid.png.enc
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
  ▸ AES
    ECC

✔ AES
Enter Password: ***********************(128 bit)
```

Decrypt a file use `AES-GCM`
```
$ ./cryptool decrypt --in test-nextdotid.png.enc --out test-nextdotid_dec.png
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
  ▸ AES
    ECC

✔ AES
Enter Password: ***********************(128 bit)
```