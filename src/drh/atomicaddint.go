package drh

import (
	"fmt"
	"sync/atomic"
)

func AtomicAddint32() {
	num := int32(10)

	for i := 0; i <= 100; i++ {
		num = atomic.AddInt32(&num, 1)
		fmt.Println(num)
	}
}
