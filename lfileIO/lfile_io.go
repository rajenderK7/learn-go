package lfileio

import (
	"fmt"
	"io"
	"os"

	"github.com/rajenderK7/learn-go/lfunc"
)

const text = "Hello there! this is Batman 🦇"

func CreateCustomFile(filename *string) bool {
	file, err := os.Create(*filename)
	lfunc.CheckNilErr(err)
	defer file.Close()
	length, err := io.WriteString(file, text)
	lfunc.CheckNilErr(err)
	fmt.Printf("%v length contents added to file\n", length)
	return true
}

func ReadFromFile(filename *string) {
	file, err := os.Open(*filename)
	lfunc.CheckNilErr(err)
	defer file.Close()
	bytes, err := io.ReadAll(file)
	lfunc.CheckNilErr(err)
	fmt.Println("Contents of the file are:")
	fmt.Println(string(bytes))
}
