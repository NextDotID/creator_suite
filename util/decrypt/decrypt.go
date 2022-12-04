package decrypt

import (
	"fmt"
	"io"
	"os"

	"github.com/nextdotid/creator_suite/util/dare"
	"golang.org/x/crypto/scrypt"
)

// ******************************* Use dare *****************************************

// Decrypt reads from src until it encounters an io.EOF and decrypts all received
// data. The decrypted data is written to dst. It returns the number of bytes
// decrypted and the first error encountered while decrypting, if any.
//
// Decrypt returns the number of bytes written to dst. Decrypt only writes data to
// dst if the data was decrypted successfully. It returns an error of type sio.Error
// if decryption fails.
func AesDecrypt(src io.Reader, dst io.Writer, config dare.Config) (n int64, err error) {
	decReader, err := AesDecryptReader(src, config)
	if err != nil {
		return 0, err
	}
	return io.CopyBuffer(dst, decReader, make([]byte, dare.MaxPayloadSize))
}

// AesDecryptReader wraps the given src and returns an io.Reader which decrypts
// all received data. DecryptReader returns an error if the provided decryption
// configuration is invalid. The returned io.Reader returns an error of
// type sio.Error if the decryption fails.
func AesDecryptReader(src io.Reader, config dare.Config) (io.Reader, error) {
	if err := dare.SetConfigDefaults(&config); err != nil {
		return nil, err
	}
	return dare.NewDecryptReaderV20(src, &config)
}

func DeriveKey(pswd []byte, src *os.File, dst *os.File) ([]byte, error) {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(src, salt); err != nil {
		return nil, fmt.Errorf("failed to read salt from '%s'", src.Name())
	}
	key, err := scrypt.Key(pswd, salt, 32768, 16, 1, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to derive key from password and salt")
	}
	return key, nil
}
