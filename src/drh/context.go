package drh

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type worker struct {
	id int
	f  func()
}

var arry []int = make([]int, 0)

func ContextDemo() {

	var stopchan chan struct{} = make(chan struct{}, 1)
	ctx1 := context.Background()
	ctx2, cancel := context.WithCancel(ctx1)
	for v := 0; v <= 3; v++ {
		f := func(i int) {
			defer cancel()
			arry = append(arry, 1)
			fmt.Println("begin to start" + strconv.Itoa(i))
			time.Sleep(time.Duration(rand.Intn(10)))
			fmt.Println("after 10 second" + strconv.Itoa(i))
			cancel()
			<-stopchan
			fmt.Println("aucatal exit")
		}
		go f(v)
	}

	for {
		select {
		case <-ctx2.Done():

			break
		}
	}

	fmt.Println(arry)
}

func watch(ctx context.Context, id int) {
	fmt.Println("execute func ...")
	value := ctx.Value(id)
	fmt.Println("ctx value", value)
	time.Sleep(11 * time.Second)
	ctx.Done()
	fmt.Println("time out  to run otherthing")
}

func CtxTimeoutDemo() {
	ctx := context.Background()
	ctx1, cancel := context.WithTimeout(ctx, time.Duration(5*time.Second))
	ctx2 := context.WithValue(ctx1, 1, 1)
	ctx3 := context.WithValue(ctx1, 2, 2)
	ctx4 := context.WithValue(ctx1, 3, 3)

	//ctx5,cancel:=context.WithTimeout(ctx2,time.Duration(5*time.Second))
	go watch(ctx2, 1)
	go watch(ctx3, 2)
	go watch(ctx4, 3)
	cancel()

	time.Sleep(10 * time.Second)
	fmt.Println("recive cancel singal")
}
