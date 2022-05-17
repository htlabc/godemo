package atomic

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//store
//
//load
//
//swap

type Config struct {
	NodeName string
	Addr     string
	Count    int
}

func loadNewoConfig() Config {
	return Config{
		NodeName: "test",
		Addr:     "192.168.0.1",
		Count:    rand.Intn(10),
	}
}

func AtomoicTest() {
	var config atomic.Value
	config.Store(loadNewoConfig())

	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Duration((5 + rand.Int63n(5))) * time.Second)
			config.Store(loadNewoConfig())
			cond.Broadcast()
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			fmt.Println("new config ", c)
			cond.L.Unlock()
		}
	}()

	select {}
}
