package drh

import "fmt"

func FuncVariable() {

	arryf := make([]func(i int), 0)

	for i := 0; i < 10; i++ {
		f := func(i int) {
			fmt.Println(i)
		}

		arryf = append(arryf, f)
	}

	arryf[0](1)

}
