package buf

type Buf struct {
	buffer   []byte
	initiate [64]byte
}

func NewBuf(buffer []byte) *Buf {
	return &Buf{buffer: buffer}
}

func (b *Buf) Buffer() []byte {
	return b.buffer
}

func (b *Buf) Grow(n int) {
	if b.buffer == nil {
		b.buffer = b.initiate[:0]
	}
	if b.Len()+n > cap(b.buffer) {
		buf := make([]byte, b.Len(), 2*cap(b.buffer)+n)
		copy(buf, b.buffer)
		b.buffer = buf
	}
}

func (b *Buf) Len() int {
	return len(b.buffer)
}
