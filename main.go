package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/easygithdev/imageio/rgb24"
)

const DIR_IN = "./sample/in"
const DIR_OUT = "./sample/out"

// Generate jpg, png and gif from a rgb file
func readFromRGB(filename string, width int, height int) {
	f, err := os.Open(DIR_IN + "/" + filename)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	// Read RGG file
	img, _ := rgb24.Decode(f, &rgb24.Options{Width: width, Height: height})

	prefixe := strings.SplitAfter(filename, ".")[0]

	// Wite PNG file
	pout, err := os.Create(DIR_OUT + "/" + prefixe + ".png")
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
	defer pout.Close()
	png.Encode(pout, img)

	// Wite JPEG file
	jout, err := os.Create(DIR_OUT + "/" + prefixe + ".jpeg")
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
	defer jout.Close()
	jpeg.Encode(jout, img, &jpeg.Options{Quality: 100})

	// Wite GIF file
	gout, err := os.Create(DIR_OUT + "/" + prefixe + ".gif")
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
	defer gout.Close()
	gif.Encode(gout, img, &gif.Options{})
}

// Get image from an URL and write the rgb file
func writeToRGB(url string) (string, int, int) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("unable to read file: %v", err)
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		log.Fatalf("unable to decode file: %v", err)
	}

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	filename := "img_" + strconv.Itoa(w) + "x" + strconv.Itoa(h) + ".rgb"
	out, _ := os.Create(DIR_IN + "/" + filename)
	defer out.Close()

	rgb24.Encode(out, img)

	return filename, w, h
}

func main() {

	readFromRGB("lena_512x512.rgb", 512, 512)

	filename, w, h := writeToRGB("https://upload.wikimedia.org/wikipedia/commons/d/df/Go_gopher_app_engine_color.jpg")

	readFromRGB(filename, w, h)

}
