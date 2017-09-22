package main

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	ss, err := GetSubmissions("UnethicalLifeProTips")
	if err != nil {
		log.Fatal(err)
	}
	i := rand.Intn(len(ss))
	im, err := MakeImage(ss[i].Title)
	if err != nil {
		log.Fatal(err)
	}
	gg.SavePNG("out.png", im)
}

func GenerateImages() {
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
