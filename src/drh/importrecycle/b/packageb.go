package b

import (
	"drh/importrecycle/a"
	"fmt"
)

type B struct {
	B      int
	Binter Binterface
}

type Binterface interface {
	GetA() *a.A
}

func NewB(b Binterface) *B {
	return &B{Binter: b}
}

func (b *B) Test() {
}

func (b *B) GetATest() {
	fmt.Println(a.Namea)
	fmt.Println(b.Binter.GetA().Id)
}
