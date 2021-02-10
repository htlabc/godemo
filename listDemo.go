package main

import (
	"container/list"
	"drh"
	"fmt"
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

	user := drh.User{Username: "snv", Socre: 10}
	drh.GetFieldName("Username", &user)
	fmt.Println(user)
}
