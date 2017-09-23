package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"math/rand"
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
