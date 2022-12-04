package dare

const (
	// AES_256_GCM specifies the cipher suite AES-GCM with 256 bit keys.
	AES_256_GCM byte = iota

	KeySize = 32

	HeaderSize     = 16
	MaxPayloadSize = 1 << 16
	TagSize        = 16
	MaxPackageSize = HeaderSize + MaxPayloadSize + TagSize

	MaxDecryptedSize = 1 << 48 // 32 TB
	MaxEncryptedSize = MaxDecryptedSize + ((HeaderSize + TagSize) * 1 << 32)
)
