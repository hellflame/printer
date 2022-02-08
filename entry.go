package main

import (
	"fmt"
	"github.com/hellflame/argparse"
	"github.com/hellflame/printer/draw"
	"github.com/nfnt/resize"
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

func showEpilog() string {
	return `try:
	printer -t hello
	printer -t hello --filter 34
	printer -t hello --filter 34 -r
	printer -t hello --filter 34 -r --color 32
	printer -t hello --filter 34 -r --color 32 -w 100
	printer -t hello --filter 34 -r --color 32 -w 100 -e 25	
	...

	printer /path/to/your/favorite/picture
	...
there's more to try

more info please visit https://github.com/hellflame/printer`
}

func main() {
	parser := argparse.NewParser("printer", "terminal printer, print your words & image in terminal", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
		EpiLog:                 showEpilog(),
		AddShellCompletion:     true,
		WithHint:               true})
	showVersion := parser.Flag("v", "version", &argparse.Option{Help: "show version info"})
	text := parser.String("t", "text", &argparse.Option{
		Default: "hellflame", Help: "render text content", Group: "Text Options"})
	width := parser.Int("w", "width", &argparse.Option{
		Default: strconv.Itoa(draw.DefaultWidth), Help: "default console width"})
	height := parser.Int("e", "height", &argparse.Option{
		Default: strconv.Itoa(draw.DefaultHeight), Help: "default console height"})
	filter := parser.Int("f", "filter", &argparse.Option{
		Default: "73", Help: fmt.Sprintf("filter ascii code, 0 ~ %d", draw.FillLength-1)})
	color := parser.Int("c", "color", &argparse.Option{
		Default: "0", Help: "color code, 30 ~ 50"})
	gray := parser.Flag("g", "gray", &argparse.Option{Help: "gray mode"})
	shade := parser.Int("s", "shade", &argparse.Option{
		Default: "128", Help: "shade cliff"})
	font := parser.String("", "font", &argparse.Option{
		Default: "0", Help: "font path or font index", Group: "Text Options"})
	reverse := parser.Flag("r", "reverse", &argparse.Option{Help: "reverse back & foreground"})
	imgPath := parser.String("", "img", &argparse.Option{
		Positional: true, Help: "image path", Validate: func(arg string) error {
			if s, err := os.Stat(arg); err == nil {
				if !s.IsDir() {
					return nil
				}
			}
			return fmt.Errorf("can't access file '%s'", arg)
		}})
	if e := parser.Parse(nil); e != nil {
		switch e.(type) {
		case argparse.BreakAfterHelp, argparse.BreakAfterShellScript:
			return // no print
		default:
			fmt.Println(e.Error())
		}
		return
	}
	if *showVersion {
		fmt.Println(draw.VERSION)
		return
	}

	var img image.Image

	if *imgPath != "" {
		// picture mode
		img = draw.LoadImage(*imgPath)
	} else {
		// text mode
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
			} else {
				fmt.Printf("font index should be 0 ~ %d\n", len(draw.DefaultFonts)-1)
				return
			}
		}
		// text mode special
		*gray = !*gray
	}
	img = resize.Resize(uint(*width), uint(*height), img, resize.Bilinear)
	fmt.Println(draw.GeneratePixel(&img, *filter, *color, *reverse, *gray, uint8(*shade)))
}
