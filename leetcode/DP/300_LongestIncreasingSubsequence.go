package DP


/* 计算最长递增子序列
算法一：
dp[i]:前i个元素的，以第i个数字结尾的最长上升子序列的长度
num[i]:第i个元素第数值
状态转移方程：dp[i] = max(dp[j]) + 1, 0 <= j < i, num[j] < num[i]
base case: dp[0] = 1
结果：LISlength = max(dp[i]), 0 <= i < n
时间复杂度：O(n^2)
*/
func LengthOfLIS(nums []int) int {
	var LISlength int
	dp := make([]int, len(nums))	// dp[i]：子串nums[:i]的最长递增子序列的长度
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])	// 取dp[j]的最大值，0 <= j < i, num[j] < num[i]
			}
		}
		dp[i] += 1	// 自增1，因为加上了自己
	}

	// 返回dp[i]中的最大值
	for _, length := range dp {
		LISlength = max(LISlength, length)
	}
	return LISlength
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/* 算法二：贪心+二分查找
贪心思想：
	如果我们要使上升子序列尽可能的长，则我们需要让序列上升得尽可能慢，
	因此我们希望每次在上升子序列最后加上的那个数尽可能的小。
数据结构：
	1.d[i]:长度为i的最长上升子序列的末尾元素的最小值。d[i]单调递增
	2.len:记录目前最长上升子序列的长度。
	起始值: d[1] = nums[0], len = 1
算法：
	设当前已求出的最长上升子序列的长度为len（初始时为 1），从前往后遍历数组nums，
	在遍历到nums[i]时：
		如果nums[i]>d[len]，则直接加入到d数组末尾，更新len=len+1;
		否则(更新数组d历史数据），在d数组中二分查找，找到第一个比nums[i]小的数d[k]，更新d[k+1]=nums[i]
*/

func LengthOfLIS2(nums []int) int {
	if len(nums) == 0 { return 0 }
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]		// dp[i]表示长度为i的最长递增子序列末尾元素的最小值
	length := 1			// length表示当前最长子序列的长度
	var l, r, pos int
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[length] {	// 加入dp
			dp[length+1] = nums[i]
			length++
		} else {					// 更新dp
			l, r = 1, length
			pos = 0  // 如果找不到说明所有的数都比 nums[i] 大，此时要更新 d[1]，所以这里将 pos 设为 0
			for l <= r {
				mid := (l+r)/2
				if nums[i] > dp[mid] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]
		}
	}
	return length
}
