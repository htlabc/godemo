package time

import (
	"fmt"
	"time"
)

func TimeDemo() {
	now := time.Now()

	local1, _ := time.LoadLocation("")

	local2, _ := time.LoadLocation("Local")
	local3, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))

}
