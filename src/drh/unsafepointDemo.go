package drh

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))            //unsafe.Pointer
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f)))) //*uint64
	return *(*uint64)(unsafe.Pointer(&f))
}

type parent struct {
	name string
	num  int
}

type children struct {
	name string
}

func ModifyByUnsafepointUnion() {

	type Person struct {
		name string
		age  int
		c    *children
	}

	who := Person{"John", 30, &children{}}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (**children)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.c)))
	*pname = "Alice"
	*page = 28
	(*pgender).name = "test"
	fmt.Println(who.c.name) // {Alice 28 false}
}

func ModifyByUnsafepoint() {
	var x struct {
		a bool
		b int16
		c []int
	}

	//var y struct{
	//	data interface{}
	//	name string
	//	age int
	//}

	/**
	  unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	*/

	/**
	  uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	  指针的运算
	*/
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"

	var b interface{}
	fmt.Println("interface size :", unsafe.Sizeof(b))

}
