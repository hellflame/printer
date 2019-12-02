// resize image to target size

package draw

import (
	"image"
	"image/color"
	_ "image/jpeg"
)

func ConvertToGray(img image.Image) *image.RGBA {
	bound := img.Bounds().Max
	X, Y := bound.X, bound.Y
	dst := image.NewRGBA(img.Bounds())
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			dst.Set(x, y, color.Gray{Y: uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)})
		}
	}
	return dst
}
