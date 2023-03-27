package encrypt

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/nacl/box"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"golang.org/x/crypto/scrypt"

	"github.com/nextdotid/creator_suite/util/dare"
)

// EncryptContentByPublicKey Encrypt content using the public key
// Returns the encrypted content file path
func EncryptPasswordByPublicKey(password string, publicKey string) (string, error) {
	publicKey = strings.TrimPrefix(publicKey, "04")
	if publicKey == "" || len(publicKey) != 128 {
		return "", fmt.Errorf("invalid public key")
	}
	if password == "" {
		return "", fmt.Errorf("invalid input content")
	}
	keyByte, err := DerivePublicKey(publicKey)
	if err != nil {
		return "", err
	}
	encryptDataByte, err := EciesEncrypt([]byte(password), keyByte)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(encryptDataByte), nil
}

func EncryptContentByPublicKey(filePath string, publicKey string) (string, error) {
	publicKey = strings.TrimPrefix(publicKey, "04")
	if publicKey == "" || len(publicKey) != 128 {
		return "", fmt.Errorf("invalid public key")
	}
	keyByte, err := DerivePublicKey(publicKey)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("fail to get the file")
	}
	//fmt.Printf("content bytes: %s", bytes)
	encryptDataByte, err := EciesEncrypt(bytes, keyByte)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(encryptDataByte), nil
}

func EncryptPasswordWithEncryptionPublicKey(encryptionPublicKey string, password string) (string, error) {
	fmt.Printf("inside fun %s", encryptionPublicKey)
	publicKeyBytes, err := base64.StdEncoding.DecodeString(encryptionPublicKey)
	if err != nil {
		return "", fmt.Errorf("DecodeString EncryptionPublicKey err:%v", err)
	}
	var publicKey [32]byte
	copy(publicKey[:], publicKeyBytes)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return "", fmt.Errorf("generate nonce for encryption err:%v", err)
	}

	pk, sk, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return "", fmt.Errorf("GenerateKey for encryption err:%v", err)
	}
	encryptedBytes := box.Seal(nil, []byte(password), &nonce, &publicKey, sk)

	//return encryptedMessage, nil
	encryptionData := map[string]interface{}{
		"version":        "x25519-xsalsa20-poly1305",
		"nonce":          base64.StdEncoding.EncodeToString(nonce[:]),
		"ephemPublicKey": base64.StdEncoding.EncodeToString(pk[:]),
		"ciphertext":     base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	encryptedData, err := json.Marshal(encryptionData)
	if err != nil {
		return "", fmt.Errorf("Json Marshal EncryptionData struct err:%v", err)
	}
	return hexutil.Encode(encryptedData), nil
}

func EncryptFileWithEncryptionPublicKey(encryptionPublicKey string, filePath string) (string, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(encryptionPublicKey)
	if err != nil {
		return "", fmt.Errorf("DecodeString EncryptionPublicKey err:%v", err)
	}
	var publicKey [32]byte
	copy(publicKey[:], publicKeyBytes)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return "", fmt.Errorf("generate nonce for encryption err:%v", err)
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("fail to get the file")
	}

	pk, sk, err := box.GenerateKey(rand.Reader)
	encryptedBytes := box.Seal(nil, []byte(fileBytes), &nonce, &publicKey, sk)

	//return encryptedMessage, nil
	encryptionData := map[string]interface{}{
		"version":        "x25519-xsalsa20-poly1305",
		"nonce":          base64.StdEncoding.EncodeToString(nonce[:]),
		"ephemPublicKey": base64.StdEncoding.EncodeToString(pk[:]),
		"ciphertext":     base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	encryptedData, err := json.Marshal(encryptionData)
	if err != nil {
		return "", fmt.Errorf("Json Marshal EncryptionData struct err:%v", err)
	}
	return hexutil.Encode(encryptedData), nil
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
