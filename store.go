package main

import (
	"fmt"
	"os"
)

func IsUsed(s string) (bool, error) {
	_, err := os.Stat(fmt.Sprintf("used/%s.txt", s))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func MarkUsed(s string, text string) error {
	f, err := os.Create(fmt.Sprintf("used/%s.txt", s))
	if err != nil {
		return err
	}
	_, err = f.WriteString(text)
	return err
}
