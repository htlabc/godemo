package bytes

import (
	"fmt"
	"io"
	"strings"
)

func lean1() {

	r := strings.NewReader("abcdefghijklmn")

	var buf []byte
	buf = make([]byte, 5)
	r.Read(buf)
	fmt.Println(string(buf)) //adcde

	buf = make([]byte, 5)
	r.Seek(-2, io.SeekCurrent) //从当前位置向前偏移两位 （5-2)
	r.Read(buf)
	fmt.Println(string(buf)) //defgh

	buf = make([]byte, 5)
	r.Seek(-3, io.SeekEnd) //设置当前位置是末尾前移三位
	r.Read(buf)
	fmt.Println(string(buf)) //lmn

	buf = make([]byte, 5)
	r.Seek(3, io.SeekStart) //设置当前位置是起始位后移三位
	r.Read(buf)
	fmt.Println(string(buf)) //defgh

}
