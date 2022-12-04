package dare

import (
	"crypto/rand"
	"fmt"
	"io"
)

// Config contains the format configuration.
// The only field which must always be set manually is the secret key.
type Config struct {
	// The secret encryption key. It must be 32 bytes long.
	Key []byte

	// The first expected sequence number.
	// It should only be set manually when decrypting a range within a stream.
	SequenceNumber uint32

	// The RNG used to generate random values.
	// If not set the default value (crypto/rand.Reader) is used.
	Rand io.Reader

	// The size of the encrypted payload in bytes.
	// The default value is 64KB. The payload size must be between 1 and 64 KB.
	// It should be used to restrict the size of encrypted packages.
	PayloadSize int
}

func SetConfigDefaults(config *Config) error {
	if len(config.Key) != KeySize {
		return fmt.Errorf("invalid key size")
	}
	if config.PayloadSize > MaxPayloadSize {
		return fmt.Errorf("sio: payload size is too large")
	}
	if config.Rand == nil {
		config.Rand = rand.Reader
	}
	if config.PayloadSize == 0 {
		config.PayloadSize = MaxPayloadSize
	}
	return nil
}
