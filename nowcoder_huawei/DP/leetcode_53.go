/* 53. Maximum Subarray
算法：
	1.从左到右遍历数组，维持两个变量：
		current_subarray:选择已有字串和当前元素的最大值
		max_subarray:记录历史最大值
*/

package main

import "fmt"

func main() {
	var size int
	var nums []int
	var num int
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	res := maxSubArray(nums)
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	var currentSubarray = nums[0]
	var maxSubarray = nums[0]
	for _, num := range nums[1:] {
		currentSubarray = myMax(currentSubarray+ num, num)
		maxSubarray = myMax(maxSubarray, currentSubarray)
	}
	return maxSubarray
}

func myMax(a, b int) int {
	if a > b { return a }
	return b
}

