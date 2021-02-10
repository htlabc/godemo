package drh

import "fmt"

func GotoDemo() {
	num := 1
breakHere:
	fmt.Println("done: ", num)
	num++

	// 手动返回, 避免执行进入标签
	goto breakHere
	return
	// 标签

}
