package draw

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const VERSION = "0.1.0"

var DefaultWidth = 50
var DefaultHeight = 80
var IsWindows = false

func init() {
	if runtime.GOOS == "windows" {
		IsWindows = true
	} else {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin // this is important
		result, err := cmd.Output()
		if err == nil {
			parse := strings.Split(strings.TrimRight(string(result), "\n"),
				" ")
			DefaultHeight, _ = strconv.Atoi(parse[0])
			DefaultWidth, _ = strconv.Atoi(parse[1])
		}
	}
}
