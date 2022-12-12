package dare

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"io"
)

type headerV20 []byte

// func (h headerV20) Version() byte         { return h[0] }
// func (h headerV20) SetVersion()           { h[0] = Version20 }
func (h headerV20) Cipher() byte          { return h[1] }
func (h headerV20) SetCipher(cipher byte) { h[1] = cipher }
func (h headerV20) Length() int           { return int(binary.LittleEndian.Uint16(h[2:4])) + 1 }
func (h headerV20) SetLength(length int)  { binary.LittleEndian.PutUint16(h[2:4], uint16(length-1)) }
func (h headerV20) IsFinal() bool         { return h[4]&0x80 == 0x80 }
func (h headerV20) Nonce() []byte         { return h[4:HeaderSize] }
func (h headerV20) AddData() []byte       { return h[:4] }
func (h headerV20) SetRand(randVal []byte, final bool) {
	copy(h[4:], randVal)
	if final {
		h[4] |= 0x80
	} else {
		h[4] &= 0x7F
	}
}

type packageV20 []byte

func (p packageV20) Header() headerV20  { return headerV20(p[:HeaderSize]) }
func (p packageV20) Payload() []byte    { return p[HeaderSize : HeaderSize+p.Header().Length()] }
func (p packageV20) Ciphertext() []byte { return p[HeaderSize:p.Length()] }
func (p packageV20) Length() int        { return HeaderSize + TagSize + p.Header().Length() }

type authEnc struct {
	// CipherID byte
	SeqNum  uint32
	Cipher  cipher.AEAD
	RandVal []byte
}

type authDec struct {
	SeqNum uint32
	Cipher cipher.AEAD
}

type authEncV20 struct {
	authEnc
	finalized bool
}

var newAesGcm = func(key []byte) (cipher.AEAD, error) {
	aes256, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(aes256)
}

func newAuthEncV20(cfg *Config) (authEncV20, error) {
	cipher, err := newAesGcm(cfg.Key)
	if err != nil {
		return authEncV20{}, err
	}
	var randVal [12]byte
	if _, err = io.ReadFull(cfg.Rand, randVal[:]); err != nil {
		return authEncV20{}, err
	}
	return authEncV20{
		authEnc: authEnc{
			// CipherID: cipherID,
			RandVal: randVal[:],
			Cipher:  cipher,
			SeqNum:  cfg.SequenceNumber,
		},
	}, nil
}

func (ae *authEncV20) Seal(dst, src []byte)      { ae.seal(dst, src, false) }
func (ae *authEncV20) SealFinal(dst, src []byte) { ae.seal(dst, src, true) }

func (ae *authEncV20) seal(dst, src []byte, finalize bool) {
	if ae.finalized { // callers are not supposed to call Seal(Final) after a SealFinal call happened
		panic("sio: cannot seal any package after final one")
	}
	ae.finalized = finalize

	header := headerV20(dst[:HeaderSize])
	// header.SetVersion()
	// header.SetCipher(ae.CipherID)
	header.SetLength(len(src))
	header.SetRand(ae.RandVal, finalize)

	var nonce [12]byte
	copy(nonce[:], header.Nonce())
	binary.LittleEndian.PutUint32(nonce[8:], binary.LittleEndian.Uint32(nonce[8:])^ae.SeqNum)

	ae.Cipher.Seal(dst[HeaderSize:HeaderSize], nonce[:], src, header.AddData())
	ae.SeqNum++
}

type authDecV20 struct {
	authDec
	refHeader headerV20
	finalized bool
}

func newAuthDecV20(cfg *Config) (authDecV20, error) {
	cipher, err := newAesGcm(cfg.Key)
	if err != nil {
		return authDecV20{}, err
	}
	return authDecV20{
		authDec: authDec{
			SeqNum: cfg.SequenceNumber,
			Cipher: cipher,
		},
	}, nil
}

func (ad *authDecV20) Open(dst, src []byte) error {
	if ad.finalized {
		return fmt.Errorf("unexpected data after final package")
	}
	if len(src) <= HeaderSize+TagSize {
		return fmt.Errorf("invalid payload size")
	}

	header := packageV20(src).Header()
	if ad.refHeader == nil {
		ad.refHeader = make([]byte, HeaderSize)
		copy(ad.refHeader, header)
	}

	if HeaderSize+header.Length()+TagSize != len(src) {
		return fmt.Errorf("invalid payload size")
	}
	if !header.IsFinal() && header.Length() != MaxPayloadSize {
		return fmt.Errorf("invalid payload size")
	}
	refNonce := ad.refHeader.Nonce()
	if header.IsFinal() {
		ad.finalized = true
		refNonce[0] |= 0x80 // set final flag
	}
	if subtle.ConstantTimeCompare(header.Nonce(), refNonce[:]) != 1 {
		return fmt.Errorf("header nonce mismatch")
	}

	var nonce [12]byte
	copy(nonce[:], header.Nonce())
	binary.LittleEndian.PutUint32(nonce[8:], binary.LittleEndian.Uint32(nonce[8:])^ad.SeqNum)
	cipher := ad.Cipher
	if _, err := cipher.Open(dst[:0], nonce[:], src[HeaderSize:HeaderSize+header.Length()+TagSize], header.AddData()); err != nil {
		return fmt.Errorf("authentication failed")
	}
	ad.SeqNum++
	return nil
}
