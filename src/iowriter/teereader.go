package iowriter

import (
	"io"
	"os"
	"strings"
)

func teereaderDemo() {
	reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
	reader.Read(make([]byte, 20))
}
