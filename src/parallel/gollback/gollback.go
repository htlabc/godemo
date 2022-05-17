package gollback

import (
	"errors"
	"fmt"
	"github.com/vardius/gollback"
	"golang.org/x/net/context"
	"time"
)

func gollbackDemo() {

	rs, errs := gollback.All(
		context.Background(),
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(3 * time.Second)
			return 1, nil
		},
		func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failed")
		},
	)

	fmt.Println(rs)
	fmt.Println(errs)

	//Race 方法Race 方法跟 All 方法类似，只不过，在使用 Race 方法的时候，只要一个异步函数执行没有错误，就立马返回，而不会返回所有的子任务信息。如果所有的子任务都没有成功，就会返回最后一个 error 信息。Race 方法签名如下：

	//Retry 不是执行一组子任务，而是执行一个子任务。如果子任务执行失败，它会尝试一定的次数，如果一直不成功 ，就会返回失败错误 ，如果执行成功，它会立即返回。如果 retires 等于 0，它会永远尝试，直到成功。

}
