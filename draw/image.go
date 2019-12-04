// resize image to target size

package draw

import (
	_ "image/jpeg"
)


func colorToGray(r, g, b, a uint32) uint8{
	return uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
}
