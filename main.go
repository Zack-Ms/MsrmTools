package main

import (
	"fmt"
)

func main() {
	// nums := []int{3,2,4}
	// target := 6

	str := "abcabcbb"
	substring := lengthOfLongestSubstring(str)
	fmt.Println(substring)
}


func lengthOfLongestSubstring(s string) int {
	return 0
}


func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		for j := i+1; j < len(numbers); j ++ {
			if target == numbers[i] + numbers[j] {
				return []int{i+1 , j+1}
			}
		}
	}
	return nil
}


type ListNode struct {
	Val int
	Next *ListNode
}
// 快慢指针
func middleNode(list *ListNode) *ListNode {
	f,r := list, list
	for f.Next != nil {
		f = list.Next.Next
		r = list.Next
	}
	return r
}