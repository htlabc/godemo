package bag

func bag(n int) int {
	var arry [3]int
	for i := 2; i < n; i++ {
		arry[2] = arry[0] + arry[1]
		arry[0] = arry[1]
		arry[1] = arry[2]
	}
	return arry[1]
}
