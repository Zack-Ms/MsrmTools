package common

/*
 * @desc
 * @author 	zk
 * @date 	2021/9/15 18:57
 ****************************************
 */

// Rotate 从结尾往前k个元素，整体左移
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}

// Reverse 反转字符串数组
func Reverse(array []string) []string {
	l, r := 0, len(array)-1
	for l < r {
		array[r], array[l] = array[l], array[r]
		r--
		l++
	}
	return array
}
