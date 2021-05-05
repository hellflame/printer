// resize image to target size

package draw

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

// LoadImage for next use
func LoadImage(imgPath string) image.Image {
	f, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
