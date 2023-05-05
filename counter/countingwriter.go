package counter

import "io"

type CountingWriter struct {
	num    *int64
	writer io.Writer
}

func (cw *CountingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	if err == nil {
		*cw.num += int64(n)
	}
	return
}

func (cw *CountingWriter) CountingWriter(w io.Writer) (io.Writer, *int64) {
	var n int64
	cw.writer = w
	cw.num = &(n)
	return cw.writer, cw.num
}
