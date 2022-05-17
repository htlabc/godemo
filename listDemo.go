package main

import (
	"container/list"
	"fmt"
	"htl.com/src/drh"
	"htl.com/src/iowriter"
	time2 "htl.com/src/time"
	"htl.com/src/wrap"
	"sync"
	"time"
)

type goods struct {
	id int
}

func listDemo() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushBack(67)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func main() {

	wrap.WrapDemo()

	drh.UnsafePointDemo()

	time2.TimeDemo()
	iowriter.TeeReaderDemo()

	intChan := make(chan int, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			<-intChan
			fmt.Println("xxxxx")
		}()
		//time.Sleep(10*time.Second)
		for {
			goto end
			//goto end
			select {
			case val := <-intChan:
				fmt.Println(val)
			default:

				fmt.Println("default")
			}
		}

	end:
		//fmt.Println("test")
		//fmt.Println(<-intChan)

	}()
	fmt.Println("get data")

	intChan <- 1
	//close(intChan	`	0)

	wg.Wait()
	fmt.Println("get data")

	for {
		time.Sleep(3 * time.Second)
		break
	}

	//etcd.EtcdDemo()
	//channel.ExampleSendDataWithChan()
	//currenttime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(currenttime)
	//fmt.Println(time.Now().UnixNano()/1e6)
	//fmt.Println(time.Now().UnixNano() / int64(time.Millisecond))
	//bb:=&b.B{}
	//
	//bb.GetATest()

	//drh.ContextDemo()
	//buf:=drh.NewleafPageElement().TranslateHexDemo()
	//fmt.Println(buf)

	//drh.PageDemo()

	//drh.CtxTimeoutDemo()

	//drh.RunSync()

	//drh.HashDemo()

	//user := drh.User{Username: "snv", Socre: 10}
	//drh.GetFieldName("Username", &user)
	//fmt.Println(user)

	//unwraperr.Demo()

	//coba.Execute()

	//ast.AstDemo()

	//go:generate go run main.go
	//go:generate go version
	//fmt.Println("http://c.biancheng.net/golang/")
	//errGroup.ErrGroupDemo()

	//stopch := make(chan struct{})
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	stopch <- struct{}{}
	//}()
	//cyclebarrier.CyclebarrierDemo(stopch)

	//
	//for{
	//
	//	fmt.Println("begin")
	//	fmt.Println(time.Now().Unix())
	//	time.Sleep(2*time.Second)
	//	fmt.Println(time.Now().Unix())
	//
	//}

}
