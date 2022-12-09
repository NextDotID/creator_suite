package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/xerrors"
	"strings"
)

// StringToPubkey is compatible with comressed / uncompressed pubkey
// hex, and with / without '0x' head.
func StringToPublicKey(pk_str string) (*ecdsa.PublicKey, error) {
	pk_str_parsed := strings.TrimPrefix(pk_str, "0x")
	pk_str_parsed = strings.ToLower(pk_str_parsed)
	pk_bytes := common.Hex2Bytes(pk_str_parsed)
	return BytesToPubKey(pk_bytes)
}

// BytesToPubKey is compatible with comressed / uncompressed pubkey
// bytes.
func BytesToPubKey(pk_bytes []byte) (*ecdsa.PublicKey, error) {
	var result *ecdsa.PublicKey
	var err error
	if len(pk_bytes) == 33 { // compressed
		result, err = crypto.DecompressPubkey(pk_bytes)
	} else {
		result, err = crypto.UnmarshalPubkey(pk_bytes)
	}
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return result, nil
}
