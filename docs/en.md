# printer

[EN Doc]()

golang version of terminal printer

python version of terminal printer : [https://github.com/hellflame/terminal_printer](https://github.com/hellflame/terminal_printer)

### show case

- default output (font file is needed, which will be mentioned below)

![](../image/example1.png)

- with given text output

![](../image/example2.png)

- colorful picture output

![](../image/example3.png)

- shaded picture output (under high resolution terminal)

![](../image/example4.png)

- with given size output

![](../image/example5.png)

- hollow out (which is just switch fill characters)

![](../image/example6.png)

- adjust shade border (with bigger `shade` , hollow area is larger, which has the same effect on picture mode)

![](../image/example7.png)

![](../image/example8.png)

- change filter character

![](../image/example9.png)

### Install

```bash
go get -u github.com/hellflame/printer
```

##### download binary

checkout https://github.com/hellflame/printer/releases 

download your os version of binary, extract file and place it in you `PATH` or using `./printer` to execute

### Usage

```
usage: printer [-h] [-v] [-t TEXT] [-w WIDTH] [-e HEIGHT] [-f FILTER] [-c COLOR] [-g] [-s SHADE] [--font FONT] [-r] [IMG]

terminal printer, print your words & image in terminal

positional arguments:
  IMG                         image path

optional arguments:
  -h, --help                  show this help message
  -v, --version               show version info
  -w WIDTH, --width WIDTH     default console width
  -e HEIGHT, --height HEIGHT  default console height
  -f FILTER, --filter FILTER  filter ascii code, 0 ~ 223
  -c COLOR, --color COLOR     color code, 30 ~ 50
  -g, --gray                  gray mode
  -s SHADE, --shade SHADE     shade cliff
  -r, --reverse               reverse back & foreground

Text Options:
  -t TEXT, --text TEXT        render text content
  --font FONT                 font path or font index

more info please visit https://github.com/hellflame/printer
```

- `-t`  Set render text，Default 'hellflame' (Font support is needed)
- `-c` Set render color，Default 0 to be random color
- `-f` Set index of filter charater，Default to choose `i`; Note: some charater is not printable
- `--font` Set font to use, if given a number, it will use default font path; if given a path, it will use font in the given path (Note: render picture don't need font support)
- `-g` Set to use gray mode
- `-s` Set border of gray shade
- `-e` Set display height, Default to be terminal height or a fixed value
- `-w` Set display width, Default to be terminal width of a fixed value

some show case is above, the other usage can be explored with fun

#### CodeUsage

> just for show

```go
import "github.com/hellflame/printer/draw"

func main(){
  img = draw.Clip(draw.Text(text, fontPath))
  result := draw.GeneratePixel(&img, filter, color, reverse?, gray?, shade)
  do_something_with(result)
}
```

### AboutFont

For it's a python migration, there is some fonts @ https://github.com/hellflame/terminal_printer/tree/master/fonts

The project default use the python version's font path, but only support `ttf` type of fonts

the fonts can be downlowd manually, place them at  `${HOME}/.terminal_fonts/` will do

### License

The MIT License (MIT)

Copyright (c) 2019 hellflame

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
