package search

// 1.18泛型不支持结构体方法，所以就直接写成普通函数的形式

// SimpleSearch 返回第一个满足条件的元素下标，传入的data必须是一个切片，否则会panic
func SimpleSearch[T any](data []T, validator func(v T) bool) int {
	for i, _ := range data {
		if validator(data[i]) {
			return i
		}
	}
	return -1
}

// SimpleSearchMultiple 返回满足条件的所有元素的下标
func SimpleSearchMultiple[T any](data []T, validator func(v T) bool) []int {
	r := make([]int, 0)
	for i, _ := range data {
		if validator(data[i]) {
			r = append(r, i)
		}
	}
	return r
}

// BinarySearch 二分查找，数据必须有序，返回从左往右数第一个满足条件的元素的下标
func BinarySearch[T any](sortedData []T, validator func(v T) bool) int {
	if len(sortedData) == 0 {
		return -1
	}
	// 维护左右指针l，r，进行二分搜索
	l, r := 0, len(sortedData)-1
	for l < r {
		mid := (l + r) / 2
		if validator(sortedData[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if validator(sortedData[l]) {
		return l
	} else {
		return -1
	}
}

// ExtremeSearch 根据传入的比较函数，返回最值对应的下标
func ExtremeSearch[T any](data []T, compare func(v1, v2 T) bool) int {
	if len(data) == 0 {
		return -1
	}
	r := 0
	for i, _ := range data {
		if compare(data[i], data[r]) {
			r = i
		}
	}
	return r
}
