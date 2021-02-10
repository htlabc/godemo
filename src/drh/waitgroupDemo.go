package drh

import (
	"fmt"
	"sync"
	"time"
)

func RunSync() {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val)
			time.Sleep(5 * time.Second)
			wg.Done()
		}(i)

	}
	wg.Wait()
}
