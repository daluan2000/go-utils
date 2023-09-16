package search_type

// 1.18泛型不支持结构体方法，所以就直接写成普通函数的形式

// 下面是函数的类型声明

// SimpleSearchT  返回第一个满足条件的元素下标，传入的data必须是一个切片，否则会panic
type SimpleSearchT[T any] func(data []T, validator func(v T) bool) int

// SimpleSearchMultipleT 返回满足条件的所有元素的下标
type SimpleSearchMultipleT[T any] func(data []T, validator func(v T) bool) []int

// BinarySearchT 二分查找，数据必须有序，返回从左往右数第一个满足条件的元素的下标
type BinarySearchT[T any] func(sortedData []T, validator func(v T) bool) int

// ExtremeSearchT 根据传入的比较函数，返回最值对应的下标
type ExtremeSearchT[T any] func(data []T, compare func(v1, v2 T) bool) int
