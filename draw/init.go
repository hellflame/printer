package draw

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

const VERSION = "v0.3.4"
const FontUrl = "https://raw.githubusercontent.com/hellflame/terminal_printer/" +
	"808004a7cd41b4383bfe6aa310c491c69d9b2556/fonts/"
const FillLength = 224

var DefaultWidth = 50
var DefaultHeight = 80
var IsShabby = false
var FillBytes []int
var FontBase string
var DefaultFonts = []string{"shuyan.ttf", "letter.ttf",
	"Haibaoyuanyuan.ttf", "fengyun.ttf"}

// decide default terminal size & initiate fill bytes & find out windows
func init() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin // this is important
	result, err := cmd.Output()
	if err == nil {
		parse := strings.Split(strings.TrimRight(string(result), "\n"),
			" ")
		DefaultHeight, _ = strconv.Atoi(parse[0])
		DefaultWidth, _ = strconv.Atoi(parse[1])
	} else {
		IsShabby = true
	}

	for i := 32; i < 32+FillLength; i++ {
		FillBytes = append(FillBytes, i)
	}

	usr, err := user.Current()
	var homeDIR string
	if err != nil {
		homeDIR = "."
	} else {
		homeDIR = usr.HomeDir
	}
	FontBase = homeDIR + "/.terminal_fonts/"
}
