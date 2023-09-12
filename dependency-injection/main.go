package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type FileDatabase struct {
	f io.Writer
}

func NewDatabase(f io.Writer) *FileDatabase {
	return &FileDatabase{f: f}
}

func main() {

	var buf bytes.Buffer
	WriteToBuffer(&buf, "hello world")

	fmt.Println("written to buffer: ", buf.String())

}

func WriteToBuffer(buf *bytes.Buffer, msg string) error {
	_, err := fmt.Fprintf(buf, msg)
	return err
}

func WriteToFile(file *os.File, msg string) error {
	_, err := fmt.Fprintf(file, msg)
	return err
}
