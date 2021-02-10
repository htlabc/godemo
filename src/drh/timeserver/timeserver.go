package timeserver

import (
	"fmt"
	"time"
)

func TimeDemo() {

	//var period int=10
	//t:=time.Duration(10)*time.Second
	//
	//fmt.Println(time.Now().Add(time.Duration(10)*time.Second))
	//fmt.Println(time.Now())
	//if time.Now().Add(time.Second*10).After(time.Now()){
	//	fmt.Println("after 10 second......")
	//}

	if time.Now().Before(time.Now().Add(time.Duration(3) * time.Second)) {
		fmt.Println("before 3 seconds")
	}

}
