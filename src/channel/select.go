package channel

import (
	"fmt"
	"reflect"
	"strconv"
)

func selectForReflectDemo() {

	var chInt = make(chan int)
	var chFloat64 = make(chan float64)
	var chString = make(chan string)
	var closeCh = make(chan struct{})

	defer close(closeCh)

	go func(c chan int, ch4close chan struct{}) {
		for i := 0; i < 3; i++ {
			chInt <- i
		}
		close(chInt)
		ch4close <- struct{}{}

	}(chInt, closeCh)

	go func(c chan float64, ch4close chan struct{}) {
		for i := 0; i < 3; i++ {
			c <- float64(i)
		}
		close(c)
		closeCh <- struct{}{}

	}(chFloat64, closeCh)

	go func(c chan string, ch4close chan struct{}) {
		for i := 0; i < 3; i++ {
			c <- strconv.Itoa(i)
		}
		close(chInt)
		closeCh <- struct{}{}
	}(chString, closeCh)

	var selectCase = make([]reflect.SelectCase, 4)
	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(chInt)

	selectCase[1].Dir = reflect.SelectRecv
	selectCase[1].Chan = reflect.ValueOf(chFloat64)

	selectCase[2].Dir = reflect.SelectRecv
	selectCase[2].Chan = reflect.ValueOf(chString)

	selectCase[3].Dir = reflect.SelectRecv
	selectCase[3].Chan = reflect.ValueOf(closeCh)

	done := 0
	finished := 0

	for finished <= len(selectCase)-1 {
		chosen, recv, revOk := reflect.Select(selectCase)
		if revOk {
			done++
			switch chosen {
			case 0:
				fmt.Println(chosen, recv.Int())
			case 1:
				fmt.Println(chosen, recv.Float())
			case 2:
				fmt.Println(chosen, recv.String())
			case 3:
				finished = finished + 1
				done--
				fmt.Println("finished\t", finished)
			}

		}

	}

}
