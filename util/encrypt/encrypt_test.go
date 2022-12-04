package encrypt

import (
	"fmt"
	"os"
	"testing"

	"github.com/nextdotid/creator_suite/util/dare"
)

func TestGenerateKeyPair(t *testing.T) {
	pairs := GenerateKeyPair()
	fmt.Printf("pairs = %v\n", pairs)
}

func TestDecrypt(t *testing.T) {
	aesKey := "8de2760ff6c6610d6d79358587d5282eeeeef2f5dd9c93c5fd6afa606015c5a935a5856df97ec7b23731a006c1a1bdb8cd8d840931c34035fa15c10b2ab431e7"
	inFile := "../../cmd/cryptool/myfile.txt"
	outFile := "../../cmd/cryptool/myfile.enc"
	in, err := os.Open(inFile)
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to open '%s': %v\n", inFile, err))
	}
	out, err := os.Create(outFile)
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to create '%s': %v\n", outFile, err))
	}
	key, err := DeriveKey([]byte(aesKey), in, out)
	cfg := dare.Config{Key: key}
	if _, err := AesEncrypt(in, out, cfg); err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
}
