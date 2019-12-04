// parse image to console

package draw

import (
	"fmt"
	"image"
	"math/rand"
	"strings"
	"time"
)

func GeneratePixel(img *image.Image, fillIndex int, colorCode int,
	reverse bool, grayMode bool) string {
	var renders []string
	src := *img
	bound := src.Bounds().Max
	X, Y := bound.X, bound.Y
	var fill uint8
	if fillIndex < 0 {
		fill = 4
	}else {
		fill = uint8(fillIndex)
	}
	colorPool := rand.New(rand.NewSource(time.Now().UnixNano()))
	for y := 0; y < Y; y ++ {
		var row []string
		for x := 0; x < X; x ++ {
			var pix string
			r, g, b, a := src.At(x, y).RGBA()
			if grayMode {
				shade := colorToGray(r, g, b, a)
				if reverse {
					pix = string(FillBytes[(255 - shade) * fill / 255])
				}else {
					pix = string(FillBytes[shade * fill / 255])
				}
			}else {
				pix = fmt.Sprintf("\033[0;38;2;%d;%d;%dm%s", r, g, b, string(FillBytes[fill]))
			}
			if colorCode == 0 {
				// random color mode
				pix = fmt.Sprintf("\033[01;%dm%s", colorPool.Intn(10) + 30, pix)
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