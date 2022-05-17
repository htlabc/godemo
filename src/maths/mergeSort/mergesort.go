package mergeSort

//{2,1,4,3}->{2,1},{4,3}->{2},{1}  ,{4,3}

func MergeSort(arry []int) []int {
	if len(arry) < 2 {
		return arry
	}

	mid := len(arry) / 2
	left := MergeSort(arry[:mid])
	right := MergeSort(arry[mid:])

	return sort(left, right)

}

func sort(left, right []int) []int {
	result := make([]int, 0)
	nl, nr := len(left), len(right)
	l, r := 0, 0
	for nl == l || r == nr {
		if left[l] > right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}

	}

	result = append(result, left[l:]...)
	right = append(right, right[r:]...)
	return result
}
