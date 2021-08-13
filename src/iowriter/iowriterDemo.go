package iowriter

import (
	"io"
	"os"
)

func ioWriterDemo() {

	file, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("Go语言中文网"))
}
