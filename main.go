package main

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"os"
)

var font *truetype.Font

func init() {
	data, err := ioutil.ReadFile("/Library/Fonts/Arial.ttf")
	if err != nil {
		panic(err)
	}
	font, err = truetype.Parse(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	titles, err := ReadTitles("titles.txt")
	if err != nil {
		panic(err)
	}
	for i, title := range titles {
		im, err := MakeImage(title)
		if err != nil {
			panic(err)
		}
		fname := fmt.Sprintf("out%d.png", i)
		gg.SavePNG(fname, im)
	}
}

func ReadTitles(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var lines []string
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func SetFontFace(dc *gg.Context, points float64) {
	dc.SetFontFace(
		truetype.NewFace(font, &truetype.Options{
			Size: points,
		}),
	)
}

func SetBestFontFace(dc *gg.Context, s string, lineHeight, h, w float64) {
	points := float64(40)
	for {
		SetFontFace(dc, points)
		_, fontHeight := dc.MeasureString(s)
		nLines := float64(len(dc.WordWrap(s, w)))
		wrappedHeight := fontHeight * lineHeight * nLines
		if wrappedHeight < h {
			fmt.Printf("difference", wrappedHeight-h)
			break
		}
		points--
	}
	SetFontFace(dc, points)
}

func MakeImage(text string) (image.Image, error) {
	const W = 300
	const H = 300
	const P = 16
	const LH = 1.75
	dc := gg.NewContext(W, H)
	dc.SetColor(color.Black)
	dc.Clear()
	dc.SetColor(color.White)
	dc.DrawRectangle(P, P, W-P-P, H-P-P)
	dc.SetLineWidth(8)
	dc.Stroke()
	dc.Height()
	SetBestFontFace(dc, text, LH, H-P-P-P-P, W-P-P-P-P)
	dc.DrawStringWrapped(text, P+P, P+P, 0, 0, W-P-P-P-P, LH, gg.AlignCenter)
	return dc.Image(), nil
}
