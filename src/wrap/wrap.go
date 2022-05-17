package wrap

import (
	"fmt"
	"runtime"
)

func WrapDemo() {
	for i := 0; i <= 4; i++ {
		test(i)
	}

}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, link, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name()
	fmt.Println(fmt.Sprintf("%v %s %d %t %s ", pc, file, link, ok, pcName))
}
