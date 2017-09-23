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
	if err := DoPost(); err != nil {
		log.Fatal(err)
	}
}

func DoPost() error {
	ss, err := GetSubmissions("UnethicalLifeProTips")
	if err != nil {
		return err
	}
	for _, i := range rand.Perm(len(ss)) {
		s := ss[i]
		used, err := IsUsed(s.Id)
		if err != nil {
			return err
		}
		if used {
			continue
		}
		im, err := MakeImage(s.Title)
		if err != nil {
			return err
		}
		if err := MarkUsed(s.Id, s.Title); err != nil {
			return nil
		}
		fmt.Println(s.Title)
		gg.SavePNG("out.png", im)
		return nil
	}
	return fmt.Errorf("all %d submissions are used", len(ss))
}
