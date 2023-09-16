package search_type

// 1.18泛型不支持结构体方法，所以就直接写成普通函数的形式

// SimpleSearch 返回第一个满足条件的元素下标，传入的data必须是一个切片，否则会panic
type SimpleSearch[T any] func(data []T, validator func(v T) bool) int

// SimpleSearchMultiple 返回满足条件的所有元素的下标
type SimpleSearchMultiple[T any] func(data []T, validator func(v T) bool) []int

// BinarySearch 二分查找，数据必须有序，返回从左往右数第一个满足条件的元素的下标
type BinarySearch[T any] func(sortedData []T, validator func(v T) bool) int

// ExtremeSearch 根据传入的比较函数，返回最值对应的下标
type ExtremeSearch[T any] func(data []T, compare func(v1, v2 T) bool) int
