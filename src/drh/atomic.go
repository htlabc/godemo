package drh

import (
	"fmt"
	"sync/atomic"
)

func AtomicDemo() {

	data := int32(1)
	for i := 0; i < 100; i++ {
		atomic.AddInt32(&data, 1)
	}

	fmt.Println(data)
}
