package rgb24

import (
	"image"
	"image/color"
	"io"
)

type Options struct {
	Width  int
	Height int
}

func Decode(r io.Reader, options *Options) (image.Image, error) {

	img := image.NewNRGBA(image.Rect(0, 0, options.Width, options.Height))
	buf := make([]byte, 3*options.Width)
	x, y := 0, 0

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {

			for i := 0; i < n; i = i + 3 {

				r := buf[i]
				g := buf[i+1]
				b := buf[i+2]

				color := color.RGBA{r, g, b, 0xff}
				img.Set(x, y, color)

				x++
			}
			x = 0
			y++
		}
	}

	return img, nil
}

func Encode(w io.Writer, img image.Image) error {

	for y := 0; y < img.Bounds().Dy(); y++ {
		i := 0
		buf := make([]byte, 3*img.Bounds().Dx())
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// best effort, but if rgb values are > 255 the result will be strange
			red := 0xFF & (r >> 8)
			green := 0xFF & (g >> 8)
			blue := 0xFF & (b >> 8)

			buf[i] = byte(red)
			buf[i+1] = byte(green)
			buf[i+2] = byte(blue)
			i += 3

		}
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
