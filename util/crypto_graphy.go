package util

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/xerrors"
)

// ValidSignatureAndGetTheAddress
func ValidSignatureAndGetTheAddress(signaturePayload string, signature string) (string, error) {
	signByte, err := hexutil.Decode(signature)

	if err != nil {
		return "", err
	}
	publicKeyRecovered, err := RecoverPubkeyFromPersonalSignature(signaturePayload, signByte)
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return crypto.PubkeyToAddress(*publicKeyRecovered).String(), nil
}

func RecoverPubkeyFromPersonalSignature(payload string, signature []byte) (pubkey *ecdsa.PublicKey, err error) {
	// Recover pubkey from signature
	if len(signature) != 65 {
		return nil, xerrors.Errorf("Error: Signature length invalid: %d instead of 65", len(signature))
	}
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	if signature[64] != 0 && signature[64] != 1 {
		return nil, xerrors.Errorf("Error: Signature Recovery ID not supported: %d", signature[64])
	}
	pubkeyRecovered, err := crypto.SigToPub(signPersonalHash([]byte(payload)), signature)
	if err != nil {
		return nil, xerrors.Errorf("Error when recovering pubkey from signature: %s", err.Error())
	}

	return pubkeyRecovered, nil
}

func signPersonalHash(data []byte) []byte {
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(message))
}
