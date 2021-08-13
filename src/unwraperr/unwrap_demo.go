package unwraperr

import (
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
)

func Demo() {
	//返回原始错误
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(errors2.Cause(e))
}
