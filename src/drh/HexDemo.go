package drh

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type leafPageElement struct {
	flags uint32 // 类型
	pos   uint32 // 偏移量(数据偏移距离)
	ksize uint32 // key大小
	vsize uint32 // 值大小
}

type page struct {
	id       int    // page id
	flags    uint16 // 区分不同类型的 page
	count    uint16 // page data 中的数据个数
	overflow uint32 // 若单个 page 大小不够，会分配多个 page
}

//type pages struct {
//	ptr []bytes
//}

const maxAllocSize = 0xFF

func PageDemo() {
	p := &page{
		id:    10,
		flags: 0x01,
	}

	//p1:=&page{
	//	id:11,
	//	flags :0x02,
	//}

	fmt.Println(unsafe.Sizeof(p))
	ptr := (*[16]byte)(unsafe.Pointer(p))

	fmt.Println(unsafe.Sizeof(&ptr[0]))
	page := (*page)(unsafe.Pointer(&ptr[0]))
	fmt.Println(page)
}

type Data struct {
	a, b int32
	c    map[string]interface{}
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func SafePointDemo() {
	d := Data{
		a: 1,
		b: 2,
		c: map[string]interface{}{
			"name": "thomas",
			"age":  89,
		},
	}
	/*
	   因[]byte底层数据结构与slicemock一致，需要构造一个完全一样
	   的数据结构进行转换
	*/
	len := unsafe.Sizeof(d)
	sliceMockTest := SliceMock{
		addr: uintptr(unsafe.Pointer(&d)),
		len:  int(len),
		cap:  int(len),
	}

	fmt.Println("address size of:", unsafe.Sizeof(sliceMockTest))
	structToByte := *(*[]byte)(unsafe.Pointer(&sliceMockTest))
	fmt.Printf("structToByte: %T, %v\n", structToByte, structToByte)

	/*
	   []byte转换成数据结构，只需取出addr地址即可，然后转换成对应的
	   数据结构类型即可
	*/
	byteToStruct := *(*Data)(unsafe.Pointer(&structToByte[0]))
	fmt.Println(&structToByte[0])
	fmt.Println("byteToStuct: ", byteToStruct)
	return
}

func putStructToPage() *page {
	p := &page{
		1, 0x01, 10, 20,
	}

	len := unsafe.Sizeof(p)
	fmt.Println("struct size of ", len)
	return p
}

func NewleafPageElement() *leafPageElement {
	return &leafPageElement{
		flags: 10,
		pos:   1,
		ksize: 1,
		vsize: 1,
	}
}

func (n *leafPageElement) TranslateHexDemo() []byte {
	buf := (*[0xfffff7]byte)(unsafe.Pointer(n))
	return (*[0xfffff7]byte)(unsafe.Pointer(&buf[n.pos]))[:n.ksize:n.ksize]
}

func Hex() {
	//byteLen:=len([]byte("test"))
	//fmt.Println([]byte("test"))
	//fmt.Println(byteLen)
	//
	//
	//fmt.Println(string([]byte{116}))
	//
	//var hexArry [0xff]byte
	//
	//hexArry[0]=0
	//hexArry[1]=100
	//fmt.Println(hexArry)

	var uint8int uint8 = 1
	var uint16int uint16 = 1
	var unit32int uint32 = 257
	var unit64int uint64 = 1

	fmt.Println(unsafe.Sizeof(uint8int))  //1字节(byte)
	fmt.Println(unsafe.Sizeof(uint16int)) //2字节(byte)
	fmt.Println(unsafe.Sizeof(unit32int)) //4字节(byte)
	fmt.Println(unsafe.Sizeof(unit64int)) //8字节(byte)

	fmt.Println(unit32int)

	fmt.Println(Uint32ToBytes(unit32int))

	bytestouint64()

	//十进制257
	//1 0000 0001
	// 将 257转成二进制就是
	// | 00000000 | 00000000 | 00000001 | 00000001 |
	// | b2[0]    | b2[1]    | b2[2]    | b2[3]    | // 这里表示b2数组每个下标里面存放的值

}

func IntDemo() {
	int64data := int64(257)
	//int8data:=int(257)
	//int32data:=int32(257)
	//int16data:=int16(257)

	//fmt.Println(int8(int64data))
	//fmt.Println(int16(int64data))
	//fmt.Println(int32(int64data))

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, &int64data)
	fmt.Println(bytesBuffer.Bytes())

}

func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

func bytestouint64() {
	i := uint32(257)
	fmt.Println(i)

	//fmt.Println(i)

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)

	//fmt.Println(b[:])

	i = uint32(binary.BigEndian.Uint32(b))
	fmt.Println(i)
}

//根据对象计算出golang对象跟地址的偏移量计算
