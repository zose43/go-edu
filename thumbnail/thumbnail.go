package thumbnail

// image scale to 128 x 128

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Thumb struct {
	ThumbFile string
	Err       error
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageToFile(outfile, infile)
}

func ImageToFile(outfile, infile string) error {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer func() { _ = in.Close() }()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	err = ImageStream(out, in)
	if err != nil {
		_ = out.Close()
		return fmt.Errorf("scaling %s to %s: %s\n", infile, outfile, err)
	}
	return out.Close()
}

func ImageStream(out io.Writer, in io.Reader) error {
	src, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(out, dst, nil)
}

func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(aspect * 128)
	} else {
		height = int(128 / aspect)
	}

	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

func Handle(filenames []string) int64 {
	size := make(chan int64)
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)

			if err != nil {
				log.Print(err)
				return
			}

			stat, _ := os.Stat(thumb)
			size <- stat.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(size)
	}()

	var total int64
	for v := range size {
		total += v
	}
	return total
}
