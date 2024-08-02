package main

func merge_sort(array []int) []int{
	if len(array) <= 1 {return array}

	index_middle := len(array)/2
	left := merge_sort(array[:index_middle])
	right := merge_sort(array[index_middle:])

	return merge(left,right)
}

func merge(left []int, right []int) []int{
	result := []int{}
	// indexes for left and right
	i,j := 0,0
	for i < len(left) && j < len(right){
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
