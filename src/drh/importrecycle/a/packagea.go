package a

var Namea string = "namea"

type A struct {
	Id int
}

func (a *A) GetA() *A {
	return a
}

func NewA() *A {
	return &A{}
}

func (a *A) test() {

}
