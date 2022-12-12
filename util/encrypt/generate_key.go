package encrypt

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	// uncompressed (65 bytes) format 0x04
	return hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)[2:]
}

// GenerateKeyPair Generates a keypair. By secp256k1 curve
func GenerateKeyPair() []string {
	// TODO: generateKeyPair
	privateKey, _ := crypto.GenerateKey()
	publicKey := &privateKey.PublicKey
	priv, pub := encode(privateKey, publicKey)
	return []string{priv, pub}
}
