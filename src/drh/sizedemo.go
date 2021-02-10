package drh

import "fmt"

func SizeDemo() {

	var size int = 102

	for {
		sz := size
		if sz > 100-1 {
			sz = 100 - 1
		}
		size -= sz
		if size == 0 {
			break
		}

		fmt.Println("sz size", sz)
		fmt.Println("size size ", size)
	}

}
