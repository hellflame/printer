// parse image to console

package draw

import (
	"fmt"
	"image"
	"math/rand"
	"strings"
	"time"
)

// generate terminal string
func GeneratePixel(img *image.Image, fillIndex int, colorCode int,
	reverse bool, grayMode bool) string {
	var renders []string
	src := *img
	bound := src.Bounds().Max
	X, Y := bound.X, bound.Y
	var fill uint8
	if fillIndex < 0 {
		fill = 4
	} else {
		fill = uint8(fillIndex)
	}
	colorPool := rand.New(rand.NewSource(time.Now().UnixNano()))
	for y := 0; y < Y; y++ {
		var row []string
		for x := 0; x < X; x++ {
			var pix string
			r, g, b, _ := src.At(x, y).RGBA()
			if grayMode {
				shade := uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24) // convert to gray
				if shade < 128 {
					shade = 0 // image binary
				} else {
					shade = 255
				}
				if reverse {
					pix = string(FillBytes[(255-shade)*fill/FillLength])
				} else {
					pix = string(FillBytes[shade*fill/FillLength])
				}

			} else {
				pix = fmt.Sprintf("\033[0;38;2;%d;%d;%dm%s",
					r/257, g/257, b/257, string(FillBytes[fill]))
			}
			if colorCode == 0 {
				// random color mode
				pix = fmt.Sprintf("\033[01;%dm%s", colorPool.Intn(10)+30, pix)
			}
			row = append(row, pix)
		}
		renders = append(renders, strings.Join(row, ""))
	}
	renderResult := strings.Join(renders, "\n")
	if colorCode > 0 {
		// specific color mode
		renderResult = fmt.Sprintf("\033[01;%dm", colorCode) + renderResult
	}
	// colorCode < 0
	return renderResult
}
