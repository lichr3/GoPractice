/* HJ24.合唱队  动态规划 队列
相当于最长递增子序列的改版。分别从两端求最长递增子序列，然后求到哪个元素时候两端的最长递增子序列加起来最大即可
 */

package main

import (
	"fmt"
	"io"
)

func main() {
	for true {
		var N int
		_, err := fmt.Scan(&N)
		if err == io.EOF { break }

		leftNums := make([]int, N)
		rightNums := make([]int, N)
		for i := 0; i < N; i++ {
			var num int
			fmt.Scan(&num)
			leftNums[i] = num
			rightNums[N-i-1] = num
		}
		fromLeft := LengthOfLIS2(leftNums)
		fromRight := LengthOfLIS2(rightNums)
		for i, j := 0, N-1; i < j; i, j = i + 1, j - 1 {
			fromRight[i], fromRight[j] = fromRight[j], fromRight[i]
		}

		//fmt.Println(fromLeft)
		//fmt.Println(fromRight)

		var res int
		for i := 0; i < N; i++ {
			res = max(fromLeft[i]+fromRight[i]-1, res)
		}
		fmt.Println(N-(res))  // res-1为合唱队行的长度，因为中间那个元素多计算了一次
	}
}

/* 最长递增子序列算法一
状态：
	dp[i]：前i个元素的最长递增子序列长度
base case:
	dp[0] = 0
状态转移：
	dp[i] = max(dp[j])+1, 0 <= j < i, nums[j] < nums[i[
*/
func LengthOfLIS(nums []int) []int {
	dp := make([]int, len(nums))	// dp[i]：子串nums[:i]的最长递增子序列的长度
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j], dp[i])	// 取dp[j]的最大值，0 <= j < i, num[j] < num[i]
			}
		}
		dp[i] += 1
	}
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/* 最长递增子序列算法二
贪心思想：同样长度的最长递增子序列，优先选择数值更小的，这也加上下一个元素能形成更长的递增子序列的几率更大
数据结构：
	dp[i]:长度为i的最长递增子序列的末尾元素的最小值。注意：d[i]单调递增
	length：当前最长递增子序列长度
base case：
	dp[1] = nums[0]
	length = 1
算法：
	对于nums中的每一个元素nums[i]
	if nums[i] > dp[length]:	// 往dp添加数据
		length += 1
		dp[length] = nums[i]
	else:  // 更新dp历史数据: 找到dp[1:length]从右往左第一个小于nums[i]的数的下标k，另dp[k+1] = nums[i]
		// 由于dp单调递增，可用二分查找，减少时间复杂度。注意有可能dp中所有的数都比nums[i]大，所以k默认值可设为0
		l, r = 1, length
		k = 0
		for l <= r:
			mid = (l+r)/2
			if nums[i] > dp[mid]:
				k = mid
				l = mid + 1
			else:
				r = mid - 1
		dp[k+1] = nums[i]
*/

func LengthOfLIS2(nums []int) []int {
	res := make([]int, len(nums))
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	length := 1; res[0] = length
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[length] {
			length++
			dp[length] = nums[i]
			res[i] = length
		} else {
			l, r := 1, length
			k := 0
			for l <= r {
				mid := (l+r)/2
				if nums[i] > dp[mid] {
					k = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[k+1] = nums[i]
			res[i] = k+1
		}
	}
	return res
}