package cyclebarrier

import (
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/net/context"
	"time"
)

func CyclebarrierDemo(stopch <-chan struct{}) {

	cnt := 0
	b := cyclicbarrier.NewWithAction(10, func() error {
		cnt++
		return nil
	})

	//wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ { // create 10 goroutines (the same count as barrier parties)
		//wg.Add(1)
		go func(i int) {
			for j := 0; j < 5; j++ {
				// do some hard work 5 times
				time.Sleep(1 * time.Second)
				err := b.Await(context.TODO()) // ..and wait for other parties on the barrier.
				// Last arrived goroutine will do the barrier action
				// and then pass all other goroutines to the next round
				fmt.Printf(" round: %d  times: %d\n ", i, j)

				if err != nil {
					panic(err)
				}
			}
			//wg.Done()
		}(i)
	}
	select {
	case <-stopch:
		fmt.Println(cnt) // cnt = 5, it means that the barrier was passed 5 times
		return
	}

	//wg.Wait()
}
