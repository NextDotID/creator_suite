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

### Ecies: encrypt & decrypt

`./cryptool ecc`

```
? Choose action:
   ▸ encrypt
      Enter Raw Password: 1234567890123456
      Enter Public Key(64 bytes): ********************************************************************************************************************************
      Encrypted Password: 0x042ce177172b8a54e11b57e4914b2ecc5070a0c6676999f4905884f1b805e97abfd2ca414f3fe94d5d78b9908f0f25230470a04056f6ab8735e3caacb606bdefce1a5f6db49bb0de18b3df17421679acaf2c16c158311972bca1ccc3131efe22620fbb2162f3ed36aa4f9be718e78eeff16538632abb73c80c759b6d47ed9e7dfb
   ▸ decrypt
      Enter Private Key(32 bytes): ****************************************************************
      Raw Password: 1234567890123456
```

### Encrypt & Upload

`$ ./cryptool encrypt`

```
Encrypting and publishing files to ipfs

 Origin File (Input): test-nextdotid.png
 Encrypt File (Output): nextdotid.png.encrypt
 Enter Raw Password: ********************
 Encrypt content finished! nextdotid.png.encrypt
 Password saved. [key id is 6 ]
 ? Do you want to publish your file to IPFS?:
   ▸ Yes
      IPFS Host: http://localhost
      IPFS API Port: 5001
      IPFS Gateway Port: 8080
      IPFS Peer ID (Authorization):
      IPFS Public Key (Authorization):
      Upload successfully!
      Path: http://localhost:8080/ipfs/{Cid}
   ▸ No
      Finished!
```

### Decrypt & Download

`$ ./cryptool decrypt`

```
Download and Decrypting files from ipfs
Use the arrow keys to navigate: ↓ ↑ → ←
   ? Do you want to download your file from IPFS?:
   ▸ Yes
      IPFS Host: http://localhost
      IPFS API Port: 5001
      IPFS Gateway Port: 8080
      IPFS Peer ID (Authorization):
      IPFS Public Key (Authorization):
      IPFS Location Url (/ipfs/${cid}): /ipfs/Cid
      Download File: nextdotid.png.download
      Download successfully!
      Path: nextdotid.png.download
      ✔  Decrypt File (Output): nextdotid_decrypt.png
      ? Choose password mode:
         ▸ Raw Password
            Enter Raw Password: ****************
            Encrypt content finished! nextdotid_decrypt.png
   ▸ No
      Origin File (Input): nextdotid.png.encrypt
      Decrypt File (Output): nextdotid_decrypt2.png
      ? Choose password mode:
         ▸ Encrypted Password
```
