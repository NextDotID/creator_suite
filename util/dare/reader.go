package dare

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
)

type encReaderV20 struct {
	authEncV20
	src io.Reader

	buffer   packageV20
	offset   int
	lastByte byte
	stateErr error

	firstRead bool
}

var packageBufferPool = sync.Pool{New: func() interface{} {
	out := make([]byte, MaxPackageSize)
	return out
}}

func recyclePackageBufferPool(b []byte) {
	if cap(b) < MaxPackageSize {
		return
	}
	// Clear so we don't potentially leak info between callers
	for i := range b {
		b[i] = 0
	}
	packageBufferPool.Put(b)
}

// encryptReaderV20 returns an io.Reader wrapping the given io.Reader.
// The returned io.Reader encrypts everything it reads using `dare`
func encryptReaderV20(src io.Reader, config *Config) (*encReaderV20, error) {
	ae, err := newAuthEncV20(config)
	if err != nil {
		return nil, err
	}
	return &encReaderV20{
		authEncV20: ae,
		src:        src,
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
		firstRead:  true,
	}, nil
}

func (r *encReaderV20) recycle() {
	recyclePackageBufferPool(r.buffer)
	r.buffer = nil
}

func (r *encReaderV20) Read(p []byte) (n int, err error) {
	if r.stateErr != nil {
		return 0, r.stateErr
	}
	if r.firstRead {
		r.firstRead = false
		_, err = io.ReadFull(r.src, r.buffer[HeaderSize:HeaderSize+1])
		if err != nil && err != io.EOF { // ErrUnexpectedEOF cannot happen b/c we read just one single byte
			return 0, err
		}
		if err == io.EOF {
			r.finalized = true
			r.stateErr = io.EOF
			r.recycle()
			return 0, io.EOF
		}
		r.lastByte = r.buffer[HeaderSize]
	}

	if r.offset > 0 { // write the buffered package to p
		remaining := r.buffer.Length() - r.offset
		if len(p) < remaining {
			r.offset += copy(p, r.buffer[r.offset:r.offset+len(p)])
			return len(p), nil
		}
		n = copy(p, r.buffer[r.offset:r.offset+remaining])
		p = p[remaining:]
		r.offset = 0
	}
	if r.finalized {
		r.stateErr = io.EOF
		r.recycle()
		return n, io.EOF
	}
	for len(p) >= MaxPackageSize {
		r.buffer[HeaderSize] = r.lastByte
		// try to read the max. payload
		nn, err := io.ReadFull(r.src, r.buffer[HeaderSize+1:HeaderSize+1+MaxPayloadSize])
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return n, err // failed to read from src
		}
		if err == io.EOF || err == io.ErrUnexpectedEOF { // read less than 64KB -> final package
			r.SealFinal(p, r.buffer[HeaderSize:HeaderSize+1+nn])
			return n + HeaderSize + TagSize + 1 + nn, io.EOF
		}
		r.lastByte = r.buffer[HeaderSize+MaxPayloadSize] // save last read byte for the next package
		r.Seal(p, r.buffer[HeaderSize:HeaderSize+MaxPayloadSize])
		p = p[MaxPackageSize:]
		n += MaxPackageSize
	}
	if len(p) > 0 {
		r.buffer[HeaderSize] = r.lastByte
		// try to read the max. payload
		nn, err := io.ReadFull(r.src, r.buffer[HeaderSize+1:HeaderSize+1+MaxPayloadSize])
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			r.stateErr = err
			r.recycle()
			return n, err // failed to read from src
		}
		if err == io.EOF || err == io.ErrUnexpectedEOF { // read less than 64KB -> final package
			r.SealFinal(r.buffer, r.buffer[HeaderSize:HeaderSize+1+nn])
			if len(p) > r.buffer.Length() {
				n += copy(p, r.buffer[:r.buffer.Length()])
				r.stateErr = io.EOF
				r.recycle()
				return n, io.EOF
			}
		} else {
			r.lastByte = r.buffer[HeaderSize+MaxPayloadSize] // save last read byte for the next package
			r.Seal(r.buffer, r.buffer[HeaderSize:HeaderSize+MaxPayloadSize])
		}
		r.offset = copy(p, r.buffer[:len(p)]) // len(p) < len(r.buffer) - otherwise we would be still in the for-loop
		n += r.offset
	}
	return n, nil
}

type decReaderV20 struct {
	authDecV20
	src io.Reader

	stateErr error
	buffer   packageV20
	offset   int
}

// decryptReaderV20 returns an io.Reader wrapping the given io.Reader.
// The returned io.Reader decrypts everything it reads using `dare`.
func decryptReaderV20(src io.Reader, config *Config) (*decReaderV20, error) {
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

func (r *decReaderV20) recycle() {
	recyclePackageBufferPool(r.buffer)
	r.buffer = nil
}

func (r *decReaderV20) Read(p []byte) (n int, err error) {
	if r.stateErr != nil {
		return 0, r.stateErr
	}
	if r.offset > 0 { // write the buffered plaintext to p
		remaining := len(r.buffer.Payload()) - r.offset
		if len(p) < remaining {
			n = copy(p, r.buffer.Payload()[r.offset:r.offset+len(p)])
			r.offset += n
			return n, nil
		}
		n = copy(p, r.buffer.Payload()[r.offset:])
		p = p[remaining:]
		r.offset = 0
	}
	for len(p) >= MaxPayloadSize {
		nn, err := io.ReadFull(r.src, r.buffer)
		if err == io.EOF && !r.finalized {
			return n, fmt.Errorf("unexpected EOF") // reached EOF but not seen final package yet
		}
		if err != nil && err != io.ErrUnexpectedEOF {
			r.recycle()
			r.stateErr = err
			return n, err // reading from src failed or reached EOF
		}
		if err = r.Open(p, r.buffer[:nn]); err != nil {
			r.recycle()
			r.stateErr = err
			return n, err // decryption failed
		}
		p = p[len(r.buffer.Payload()):]
		n += len(r.buffer.Payload())
	}
	if len(p) > 0 {
		nn, err := io.ReadFull(r.src, r.buffer)
		if err == io.EOF && !r.finalized {
			r.recycle()
			r.stateErr = fmt.Errorf("unexpected EOF")
			return n, fmt.Errorf("unexpected EOF") // reached EOF but not seen final package yet
		}
		if err != nil && err != io.ErrUnexpectedEOF {
			r.recycle()
			r.stateErr = err
			return n, err // reading from src failed or reached EOF
		}
		if err = r.Open(r.buffer[HeaderSize:], r.buffer[:nn]); err != nil {
			return n, err // decryption failed
		}
		if payload := r.buffer.Payload(); len(p) < len(payload) {
			r.offset = copy(p, payload[:len(p)])
			n += r.offset
		} else {
			n += copy(p, payload)
		}
	}
	return n, nil
}

// decryptBufferV20 will append plaintext to dst and return the result.
func decryptBufferV20(dst, src []byte, config *Config) ([]byte, error) {
	ad, err := newAuthDecV20(config)
	if err != nil {
		return nil, err
	}
	for len(src) > 0 {
		buffer := packageV20(src)
		// Truncate to max package size
		if len(buffer) > MaxPackageSize {
			buffer = buffer[:MaxPackageSize]
		}

		// Make space in dst
		payloadLen := buffer.Header().Length()
		if cap(dst) >= len(dst)+payloadLen {
			dst = dst[:len(dst)+payloadLen]
		} else {
			dst = append(dst, make([]byte, payloadLen)...)
		}

		// Write directly to dst.
		if err = ad.Open(dst[len(dst)-payloadLen:], buffer); err != nil {
			return nil, err // decryption failed
		}
		// Forward to next block.
		src = src[buffer.Length():]
	}
	if !ad.finalized {
		return nil, fmt.Errorf("unexpected EOF")
	}
	return dst, nil
}

type decReaderAtV20 struct {
	src io.ReaderAt
	ad  authDecV20
}

// decryptReaderAtV20 returns an io.ReaderAt wrapping the given io.ReaderAt.
// The returned io.ReaderAt decrypts everything it reads using `dare`.
func decryptReaderAtV20(src io.ReaderAt, config *Config) (*decReaderAtV20, error) {
	ad, err := newAuthDecV20(config)
	if err != nil {
		return nil, err
	}
	r := &decReaderAtV20{
		ad:  ad,
		src: src,
	}

	return r, nil
}

func (r *decReaderAtV20) ReadAt(p []byte, offset int64) (n int, err error) {
	if offset < 0 {
		return 0, errors.New("sio: DecReaderAt.ReadAt: offset is negative")
	}

	t := offset / int64(MaxPayloadSize)
	if t+1 > (1<<32)-1 {
		return 0, fmt.Errorf("size is too large for dare")
	}

	decReader := decReaderV20{
		authDecV20: r.ad,
		src:        &sectionReader{r.src, t * MaxPackageSize},
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
		offset:     0,
	}
	defer decReader.recycle()
	decReader.SeqNum = uint32(t)
	if k := offset % int64(MaxPayloadSize); k > 0 {
		if _, err := io.CopyN(ioutil.Discard, &decReader, k); err != nil {
			return 0, err
		}
	}

	for n < len(p) && err == nil {
		var nn int
		nn, err = (&decReader).Read(p[n:])
		n += nn
	}
	if err == io.EOF && n == len(p) {
		err = nil
	}
	return n, err
}

// Use a custom sectionReader since io.SectionReader
// demands a read limit.

type sectionReader struct {
	r   io.ReaderAt
	off int64
}

func (r *sectionReader) Read(p []byte) (int, error) {
	n, err := r.r.ReadAt(p, r.off)
	r.off += int64(n)
	return n, err
}

// NewEncryptReaderV20 returns an io.Reader wrapping the given io.Reader.
// The returned io.Reader encrypts everything it reads using DARE 2.0.
func NewEncryptReaderV20(src io.Reader, config *Config) (*encReaderV20, error) {
	ae, err := newAuthEncV20(config)
	if err != nil {
		return nil, err
	}
	return &encReaderV20{
		authEncV20: ae,
		src:        src,
		buffer:     packageBufferPool.Get().([]byte)[:MaxPackageSize],
		firstRead:  true,
	}, nil
}
