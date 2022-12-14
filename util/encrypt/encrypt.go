package encrypt

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"golang.org/x/crypto/scrypt"

	"github.com/nextdotid/creator_suite/util/dare"
)

// EncryptContentByPublicKey Encrypt content using the public key
// Returns the encrypted content file path
func EncryptContentByPublicKey(content string, publicKey string) (string, error) {
	if publicKey == "" || len(publicKey) != 128 {
		return "", fmt.Errorf("invalid public key")
	}
	if content == "" {
		return "", fmt.Errorf("invalid input content")
	}
	keyByte, err := DerivePublicKey(publicKey)
	if err != nil {
		return "", err
	}
	encryptDataByte, err := EciesEncrypt([]byte(content), keyByte)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(encryptDataByte), nil
}

// ******************************* Use dare *****************************************

// Encrypt reads from src until it encounters an io.EOF and encrypts all received data.
// The encrypted data is written to dst. It returns the number of bytes and first error encountered while encrypting.
// Encrypt returns the number of bytes written to dst.
func AesEncrypt(src io.Reader, dst io.Writer, config dare.Config) (int64, error) {
	encReader, err := AesEncryptReader(src, config)
	if err != nil {
		return 0, err
	}

	return io.CopyBuffer(dst, encReader, make([]byte, dare.HeaderSize+dare.MaxPackageSize+dare.TagSize))
}

// AesEncryptReader wraps the given src and returns an io.Reader which encrypts all received data.
func AesEncryptReader(src io.Reader, config dare.Config) (io.Reader, error) {
	if err := dare.SetConfigDefaults(&config); err != nil {
		return nil, err
	}
	return dare.NewEncryptReaderV20(src, &config)

}

func DeriveKey(pswd []byte, src *os.File, dst *os.File) ([]byte, error) {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("failed to generate random salt '%s'", src.Name())
	}
	if _, err := dst.Write(salt); err != nil {
		return nil, fmt.Errorf("failed to write salt to '%s'", dst.Name())
	}
	key, err := scrypt.Key(pswd, salt, 32768, 16, 1, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to derive key from password and salt")
	}
	return key, nil
}

// ******************************* Use go-ethereum/crypto/ecies ************************************

func EciesEncrypt(content []byte, publicKey []byte) ([]byte, error) {
	if len(publicKey) <= 0 {
		return nil, fmt.Errorf("public key must not be null")
	}
	if len(content) <= 0 {
		return nil, fmt.Errorf("encrypt content must not be null")
	}
	pubkey, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return nil, err
	}
	pubkeyEcies := ecies.ImportECDSAPublic(pubkey)
	encryptContent, err := ecies.Encrypt(rand.Reader, pubkeyEcies, content, nil, nil)
	if err != nil {
		return nil, err
	}
	return encryptContent, nil
}

// DerivePublicKey
// We strip off the 0x and the first 2 characters 04
// which is always the EC prefix and is not required in public key
func DerivePublicKey(pubkey string) ([]byte, error) {
	key := fmt.Sprintf("0x04%s", pubkey)
	keyByte, err := hexutil.Decode(key)
	if err != nil {
		return nil, err
	}
	additionByte, err := hexutil.Decode("0x04")
	if err != nil {
		return nil, err
	}
	additionByteLen := len(additionByte)
	if len(keyByte) != (64 + additionByteLen) {
		return nil, fmt.Errorf("public key must be equal to 64 bytes")
	}
	return keyByte, nil
}
