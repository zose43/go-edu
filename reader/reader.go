package reader

import "io"

type LimitReader struct {
	r io.Reader
	n int64
}

func (lr *LimitReader) LimitReader(p []byte) (n int, err error) {
	if lr.n > int64(len(p)) {
		n, err = lr.r.Read(p)
		if err != nil {
			return 0, err
		}
		lr.n -= int64(n)
	} else {
		p = p[:lr.n]
		n, err = lr.r.Read(p)
		if err != nil {
			return 0, err
		}
		err = io.EOF
	}
	return
}

func NewLimitReader(reader io.Reader, n int64) LimitReader {
	return LimitReader{reader, n}
}
