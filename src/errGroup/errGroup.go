package errGroup

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func ErrGroupDemo() {
	var g errgroup.Group
	var result = make([]error, 3)

	g.Go(
		func() error {
			time.Sleep(5 * time.Second)
			fmt.Println("exec #1")
			result[0] = nil
			return nil
		},
	)

	g.Go(
		func() error {
			result[1] = errors.New("error test")
			return result[1]
		},
	)

	if err := g.Wait(); err != nil {
		fmt.Println("Successfully exec all,result:%v\n", result)
	} else {
		fmt.Println("failed :%v\n", result)
	}

}
