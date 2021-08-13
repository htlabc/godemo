package ioutil

import "github.com/gogo/protobuf/io"

//1.6. Discard 变量
//Discard 对应的类型（type devNull int）实现了 io.Writer 接口，同时，为了优化 io.Copy 到 Discard，避免不必要的工作，实现了 io.ReaderFrom 接口。
//
//devNull 在实现 io.Writer 接口时，只是简单的返回（标准库文件：src/pkg/io/ioutil.go)。

var Discard io.Writer

//func (dis Discard) Write(p []byte) (int, error) {
//	return len(p), nil
//}
