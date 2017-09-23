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
	for _, i := range rand.Perm(len(ss)) {
		s := ss[i]
		used, err := IsUsed(s.Id)
		if err != nil {
			log.Fatal(err)
		}
		if !used {
			im, err := MakeImage(s.Title)
			if err != nil {
				log.Fatal(err)
			}
			if err := MarkUsed(s.Id, s.Title); err != nil {
				log.Fatal(err)
			}
			fmt.Println(s.Title)
			gg.SavePNG("out.png", im)
			break
		}
	}
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
