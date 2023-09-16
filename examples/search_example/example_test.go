package search

import (
	"fmt"
	"go-utils/search"
	"testing"
)

func Test_Example1(t *testing.T) {
	data := []int{70, 90, 20, 0, -10}
	sortedData := []int{10, 20, 30, 40, 50}
	strData := []string{"123", "1222", "12222"}

	// 第一个小于1的元素下标，输出 3
	fmt.Println(search.SimpleSearch(data, func(v int) bool {
		return v < 1
	}))

	// 查找所有大于1的元素下标，输出 [0 1 2]
	fmt.Println(search.SimpleSearchMultiple(data, func(v int) bool {
		return v > 1
	}))

	// 查找第一个大于等于3的元素下标，输出3
	fmt.Println(search.BinarySearch(sortedData, func(v int) bool {
		return v > 30
	}))

	// 查找第一个大于等于100的元素下标，输出-1
	fmt.Println(search.BinarySearch(sortedData, func(v int) bool {
		return v > 100
	}))

	// 查找最大的元素下标，输出1
	fmt.Println(search.ExtremeSearch(data, func(v1, v2 int) bool {
		return v1 > v2
	}))

	// 查找最长的字符串下标，输出2
	fmt.Println(search.ExtremeSearch(strData, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	}))
}
