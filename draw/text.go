// text drawer

package draw

import (
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"unicode/utf8"
)

func Text(text string, fontPath string) *image.Gray {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Fatal(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	fontSize := 20.
	c := freetype.NewContext()
	c.SetDPI(96)
	c.SetFont(font)
	c.SetFontSize(fontSize)

	height := int(c.PointToFixed(fontSize) >> 6)
	width := utf8.RuneCountInString(text) * height // still need to crop

	fg, bg := image.Black, image.White

	canvas := image.NewGray(image.Rect(0, 0, width, height+5))
	draw.Draw(canvas, canvas.Bounds(), bg, image.Point{}, draw.Src)

	c.SetClip(canvas.Bounds())
	c.SetDst(canvas)
	c.SetSrc(fg)

	_, err = c.DrawString(text, freetype.Pt(0, height))
	if err != nil {
		log.Fatal(err)
	}

	return canvas
}

// clip text image to right size
// for golang doesn't have any direct way to get render string size
func Clip(img *image.Gray) image.Image {

	bound := img.Bounds().Max
	X, Y := bound.X, bound.Y

	var rowPixels [][]bool
	for y := 0; y < Y; y++ {
		var row []bool
		for x := 0; x < X; x++ {
			row = append(row, img.GrayAt(x, y).Y == 255)
		}
		rowPixels = append(rowPixels, row)
	}

	upBorder, upBoarderFound := 0, false
	downBorder := 0

	for y := 0; y < Y; y++ {
		isBlank := true
		for x := 0; x < X; x++ {
			if !rowPixels[y][x] {
				isBlank = false
			}
		}
		if !upBoarderFound && isBlank {
			upBorder++
		}
		if !isBlank {
			upBoarderFound = true
		}
		if downBorder == 0 && isBlank && upBoarderFound {
			downBorder = y
		}
	}

	leftBorder := 0
	rightBorder := X

	for x := 0; x < X; x++ {
		isBlank := true
		for y := 0; y < Y; y++ {
			if !rowPixels[y][x] {
				isBlank = false
			}
		}
		if !isBlank {
			break
		}

		if isBlank {
			leftBorder++
		}
	}
	for x := X - 1; x >= 0; x-- {
		isBlank := true
		for y := 0; y < Y; y++ {
			if !rowPixels[y][x] {
				isBlank = false
			}
		}
		if isBlank {
			rightBorder--
		}
		if !isBlank {
			break
		}

	}

	// possible issue
	if leftBorder > 0 {
		leftBorder -= 1
	}
	if downBorder <= 0 {
		downBorder = Y
	}
	rightBorder += 1

	return img.SubImage(image.Rect(leftBorder, upBorder, rightBorder, downBorder))
}
