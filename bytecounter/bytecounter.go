package bytecounter

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	n = len(p)
	*b += ByteCounter(n)
	return
}
