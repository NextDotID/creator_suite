package dare

import (
	"fmt"
	"io"
)

type encWriterV20 struct {
	authEncV20
	dst io.Writer

	buffer   packageV20
	offset   int
	closeErr error
}

func flush(w io.Writer, p []byte) error {
	n, err := w.Write(p)
	if err != nil {
		return err
	}
	if n != len(p) { // not necessary if the w follows the io.Writer doc *precisely*
		return io.ErrShortWrite
	}
	return nil
}

// encryptWriterV20 returns an io.WriteCloser wrapping the given io.Writer.
// The returned io.WriteCloser encrypts everything written to it using `dare`
// and writes all encrypted ciphertext as well as the package header and tag
// to the wrapped io.Writer.
//
// The io.WriteCloser must be closed to finalize the encryption successfully.
func encryptWriterV20(dst io.Writer, config *Config) (*encWriterV20, error) {
	ae, err := newAuthEncV20(config)
	if err != nil {
		return nil, err
	}
	return &encWriterV20{
		authEncV20: ae,
		dst:        dst,
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
	}, nil
}

func (w *encWriterV20) Write(p []byte) (n int, err error) {
	if w.finalized {
		// The caller closed the encWriterV20 instance (called encWriterV20.Close()).
		// This is a bug in the calling code - Write after Close is not allowed.
		panic("critical error: write to stream after close")
	}
	if w.offset > 0 { // buffer the plaintext data
		remaining := MaxPayloadSize - w.offset
		if len(p) <= remaining { // <= is important here to buffer up to 64 KB (inclusively) - see: Close()
			w.offset += copy(w.buffer[HeaderSize+w.offset:], p)
			return len(p), nil
		}
		n = copy(w.buffer[HeaderSize+w.offset:], p[:remaining])
		w.Seal(w.buffer, w.buffer[HeaderSize:HeaderSize+MaxPayloadSize])
		if err = flush(w.dst, w.buffer); err != nil { // write to underlying io.Writer
			return n, err
		}
		p = p[remaining:]
		w.offset = 0
	}

	for len(p) > MaxPayloadSize {
		// > is important here to call Seal (not SealFinal) only if there is at least on package left - see: Close()
		w.Seal(w.buffer, p[:MaxPayloadSize])
		if err = flush(w.dst, w.buffer); err != nil { // write to underlying io.Writer
			return n, err
		}
		p = p[MaxPayloadSize:]
		n += MaxPayloadSize
	}
	if len(p) > 0 {
		w.offset = copy(w.buffer[HeaderSize:], p)
		n += w.offset
	}
	return n, nil
}

func (w *encWriterV20) Close() (err error) {
	if w.buffer == nil {
		return w.closeErr
	}
	defer func() {
		w.closeErr = err
		recyclePackageBufferPool(w.buffer)
		w.buffer = nil
	}()

	if w.offset > 0 { // true if at least one Write call happened
		w.SealFinal(w.buffer, w.buffer[HeaderSize:HeaderSize+w.offset])
		if err = flush(w.dst, w.buffer[:HeaderSize+w.offset+TagSize]); err != nil { // write to underlying io.Writer
			return err
		}
		w.offset = 0
	}
	if closer, ok := w.dst.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

type decWriterV20 struct {
	authDecV20
	dst io.Writer

	buffer   packageV20
	offset   int
	closeErr error
}

// decryptWriterV20 returns an io.WriteCloser wrapping the given io.Writer.
// The returned io.WriteCloser decrypts everything written to it using DARE 2.0
// and writes all decrypted plaintext to the wrapped io.Writer.
//
// The io.WriteCloser must be closed to finalize the decryption successfully.
func decryptWriterV20(dst io.Writer, config *Config) (*decWriterV20, error) {
	ad, err := newAuthDecV20(config)
	if err != nil {
		return nil, err
	}
	return &decWriterV20{
		authDecV20: ad,
		dst:        dst,
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
	}, nil
}

func (w *decWriterV20) Write(p []byte) (n int, err error) {
	if w.offset > 0 { // buffer package
		remaining := HeaderSize + MaxPayloadSize + TagSize - w.offset
		if len(p) < remaining {
			w.offset += copy(w.buffer[w.offset:], p)
			return len(p), nil
		}
		n = copy(w.buffer[w.offset:], p[:remaining])
		plaintext := w.buffer[HeaderSize : HeaderSize+MaxPayloadSize]
		if err = w.Open(plaintext, w.buffer); err != nil {
			return n, err
		}
		if err = flush(w.dst, plaintext); err != nil { // write to underlying io.Writer
			return n, err
		}
		p = p[remaining:]
		w.offset = 0
	}
	for len(p) >= MaxPackageSize {
		plaintext := w.buffer[HeaderSize : HeaderSize+MaxPayloadSize]
		if err = w.Open(plaintext, p[:MaxPackageSize]); err != nil {
			return n, err
		}
		if err = flush(w.dst, plaintext); err != nil { // write to underlying io.Writer
			return n, err
		}
		p = p[MaxPackageSize:]
		n += MaxPackageSize
	}
	if len(p) > 0 {
		if w.finalized {
			return n, fmt.Errorf("unexpected data after final package")
		}
		w.offset = copy(w.buffer[:], p)
		n += w.offset
	}
	return n, nil
}

func (w *decWriterV20) Close() (err error) {
	if w.buffer == nil {
		return w.closeErr
	}
	defer func() {
		w.closeErr = err
		recyclePackageBufferPool(w.buffer)
		w.buffer = nil
	}()
	if w.offset > 0 {
		if w.offset <= HeaderSize+TagSize { // the payload is always > 0
			return fmt.Errorf("invalid payload size")
		}
		if err = w.Open(w.buffer[HeaderSize:w.offset-TagSize], w.buffer[:w.offset]); err != nil {
			return err
		}
		if err = flush(w.dst, w.buffer[HeaderSize:w.offset-TagSize]); err != nil { // write to underlying io.Writer
			return err
		}
		w.offset = 0
	}

	if closer, ok := w.dst.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// NewDecryptReaderV20 returns an io.Reader wrapping the given io.Reader.
// The returned io.Reader decrypts everything it reads using DARE 2.0.
func NewDecryptReaderV20(src io.Reader, config *Config) (*decReaderV20, error) {
	ad, err := newAuthDecV20(config)
	if err != nil {
		return nil, err
	}
	return &decReaderV20{
		authDecV20: ad,
		src:        src,
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
	}, nil
}
