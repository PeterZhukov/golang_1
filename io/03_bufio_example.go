package io

import (
	"bufio"
	"io"
	"log"
	"os"
)

func BufIoMain() {
	f, err := os.Open("./io/03_bufio_example.go")
	if err != nil {
		log.Fatal(err)
	}
	b, err := get(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	f, err = os.Create("./io/03_bufio_example.go.bak")
	if err != nil {
		log.Fatal(err)
	}
	err = store(f, b)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func store(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

func get(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)

	var b []byte
	for scanner.Scan() {
		b = append(b, []byte(scanner.Text()+"\n")...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return b, nil
}
