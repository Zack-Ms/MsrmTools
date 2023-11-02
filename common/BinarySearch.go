package common

import (
	"fmt"
	"strings"
)

func test() {
	nums := []int{1, 2, 4, 5, 6}
	target := 2
	fmt.Println(BinarySearchInsert(target, nums))
}

// BinarySearch 二分查找 返回索引
func BinarySearch(target int, array []int) int {
	left, right, middle := 0, len(array)-1, 0
	for left <= right {
		middle = (right + left) / 2
		value := array[middle]
		if value < target {
			left = middle + 1
		} else if value > target {
			right = middle - 1
		} else if value == target {
			return middle
		}
	}
	return -1
}

// BinarySearchFirst 二分查找 返回第一个出现的目标索引
func BinarySearchFirst(value string, array []string) int {
	left, right, middle := 0, len(array), 0
	for left < right {
		middle = left + (right-left)/2
		if strings.EqualFold(value, array[middle]) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return right
}

// BinarySearchInsert 二分查找 查询索引，返回目标索引，如果不存在则返回目标按顺序插入后的索引
func BinarySearchInsert(value int, array []int) int {
	left, right, middle := 0, len(array), 0
	for left < right {
		middle = left + (right-left)/2
		if array[middle] > value {
			right = middle
		} else if array[middle] < value {
			left = middle + 1
		} else {
			return middle
		}
	}
	return right
}
