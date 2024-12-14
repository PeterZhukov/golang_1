package io

import (
	"io"
	"log"
	"os"
)

func sendReversedString(s string, w io.Writer) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	_, err := w.Write([]byte(string(runes)))
	return err
}

func WriteExampleMain() {
	f, err := os.Create("./strings.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = sendReversedString("hello world", f)
	if err != nil {
		log.Fatal(err)
	}

	err = sendReversedString("hello world", os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
