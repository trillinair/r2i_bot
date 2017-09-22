package main

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"io"
	"os"
)

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

func LoadBestFont(dc *gg.Context, s string, lineHeight, h, w float64) error {
	points := float64(40)
	for {
		if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", points); err != nil {
			return err
		}
		_, fontHeight := dc.MeasureString(s)
		nLines := float64(len(dc.WordWrap(s, w)))
		wrappedHeight := fontHeight * lineHeight * nLines

		if wrappedHeight < h {
			break
		}
		points--
	}
	return dc.LoadFontFace("/Library/Fonts/Arial.ttf", points)
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
	if err := LoadBestFont(dc, text, LH, H-P-P-P-P, W-P-P-P-P); err != nil {
		return nil, err
	}
	dc.DrawStringWrapped(text, P+P, P+P, 0, 0, W-P-P-P-P, LH, gg.AlignCenter)
	return dc.Image(), nil
}
