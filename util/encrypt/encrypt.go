package encrypt

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/nextdotid/creator_suite/util/dare"
	"golang.org/x/crypto/scrypt"
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
	output := fmt.Sprintf("%s.enc", content)
	src, err := os.Open(content)
	if err != nil {
		return "", fmt.Errorf("failed to open '%s': %v", content, err)
	}
	dst, err := os.Create(output)
	if err != nil {
		return "", fmt.Errorf("failed to create '%s': %v", output, err)
	}

	key, err := DeriveKey([]byte(publicKey), src, dst)
	if err != nil {
		return "", err
	}
	cfg := dare.Config{Key: key}
	// TODO: use defer to clean file when encrypt failed.
	if _, err := AesEncrypt(src, dst, cfg); err != nil {
		return "", err
	}
	return output, nil
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
