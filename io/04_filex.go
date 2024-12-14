package io

import (
	"fmt"
	"log"
	"os"
)

func FilesMain() {
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = os.WriteFile(f.Name(), []byte("Текст"), 0666)
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Данные файла:", data)

	var file *os.File
	file, err = os.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 6)
	n, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	buf = buf[:n]
	fmt.Printf("Данные файла:\n%s\n", buf)
}
