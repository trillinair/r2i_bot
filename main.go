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
	"math"
	"os"
)

var font *truetype.Font

func init() {
	data, err := ioutil.ReadFile("trade-gothic-bold-condensed-20.ttf")
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

func AdjustPoints(err, points float64) float64 {
	const limit = 1000
	const speed = 5
	return points - math.Max(((math.Min(limit, err)/limit)*speed), 0.001)
}

func SetBestFontFace(dc *gg.Context, s string, lineHeight, h, w float64) float64 {
	points := float64(40)
	prev := points
	for {
		SetFontFace(dc, points)
		_, fontHeight := dc.MeasureString(s)
		nLines := float64(len(dc.WordWrap(s, w)))
		wrappedHeight := fontHeight * lineHeight * nLines
		err := wrappedHeight - h
		if err <= 0 {
			SetFontFace(dc, prev)
			return wrappedHeight
		}
		prev = points
		points = AdjustPoints(err, points)
	}
}

func MakeImage(text string) (image.Image, error) {
	const W = 612
	const H = 612
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

	textHeight := float64(H - P - P - P - P)
	textWidth := float64(W - P - P - P - P)
	actualTextHeight := SetBestFontFace(dc, text, LH, textHeight, textWidth)
	offset := (textHeight - actualTextHeight) / 2

	dc.DrawStringWrapped(text, P+P, P+P+offset, 0, 0, textWidth, LH, gg.AlignCenter)
	return dc.Image(), nil
}
