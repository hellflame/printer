package main

import (
	"fmt"
	"github.com/hellflame/printer/draw"
	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func exist(p string) bool {
	s, err := os.Stat(p)
	if err == nil {
		if !s.IsDir() {
			return true
		}
	}
	return false
}

func downloadFont(fontName string) {
	savePath := draw.FontBase + fontName
	if exist(savePath) {
		return
	}
	fmt.Printf("downloading font %s ...", fontName)
	resp, err := http.Get(draw.FontUrl + fontName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	stat, err := os.Stat(draw.FontBase)
	if err == nil {
		if !stat.IsDir() {
			log.Fatalf("save path is not dir, %s", draw.FontBase)
		}
	} else {
		if err = os.MkdirAll(draw.FontBase, 0755); err != nil {
			log.Fatal(err)
		}
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	saver, err := os.Create(savePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = saver.Close() }()
	_, _ = saver.Write(content)
	fmt.Println("font saved")
}

func runner() {
	var imgPath string
	var run bool
	cmd := &cobra.Command{
		Use:     "printer",
		Short:   "terminal printer",
		Version: draw.VERSION,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				imgPath = args[0]
			}
			run = true
		},
	}
	flag := cmd.Flags()
	text := flag.StringP("text", "t", "hellflame", "render text content")
	width := flag.UintP("width", "w", uint(draw.DefaultWidth), "default console width")
	height := flag.UintP("height", "e", uint(draw.DefaultHeight), "default console height")
	filter := flag.IntP("filter", "f", 73, "filter ascii code")
	color := flag.IntP("color", "c", 0, "color code, 30 ~ 50")
	gray := flag.BoolP("gray", "g", false, "gray mode")
	shade := flag.Uint8P("shade", "s", 128, "shade cliff")
	font := flag.String("font", "0", "font path or font index")
	reverse := flag.BoolP("reverse", "r", false, "reverse back & foreground")
	err := cmd.Execute()
	if err != nil {
		return
	}
	if !run {
		return
	}

	var img image.Image

	if imgPath != "" && exist(imgPath) {
		img = draw.LoadImage(imgPath)
	} else {
		fontIndex, err := strconv.Atoi(*font)
		if err != nil {
			// font is given by user
			if exist(*font) {
				img = draw.Clip(draw.Text(*text, *font))
			} else {
				fmt.Println("font path not exist")
				return
			}
		} else {
			if 0 <= fontIndex && fontIndex < len(draw.DefaultFonts) {
				// font is provide
				fontPath := draw.DefaultFonts[fontIndex]
				downloadFont(fontPath)
				img = draw.Clip(draw.Text(*text, draw.FontBase+fontPath))
			}else {
				fmt.Printf("font index should be 0 ~ %d\n", len(draw.DefaultFonts) - 1)
				return
			}
		}
		*gray = true
	}

	img = resize.Resize(*width, *height, img, resize.Bilinear)
	fmt.Println(draw.GeneratePixel(&img, *filter, *color, *reverse, *gray, *shade))
}

func main() {
	runner()
}
