/* 0-1背包问题
状态：
	1.两个状态：【可选择的物品】【背包的容量】
	2.表示方法：dp[i][w]，对于前i个物品，容量为w时，能装下物品的最大价值
选择：
	1.背包容量为w时，对于第i件物品，选择放入背包或者不放入背包（容量不足时不放入，放入的价值小于不放入的价值也不放入）
	2.选择不放入背包：	dp[i][w] = dp[i-1][w]
	3.选择放入背包： 	dp[i][w] = dp[i-1][w-wt[i]] + val[i]  （代码理解，若要放入第i件物品，则放入前第状态为dp[i-1][w-wt[i]]
	4.若该物品能放入背包，则选择3和4结果的最大值
base case:
	dp[0][..] = 0; dp[..][0] = 0
易错点：
	由于我们从1而非0开始计数，因此用到wt和val时应该在数组前加一个元素，使得下标也从1开始

伪代码：
	总共有N个物品，最大重量为W；wt[i]表示第i件物品第重量，val[i]表示第i件物品第价值
	int dp[N+1][W+1]
	dp[0][..] = 0
	dp[..][0] = 0
	for i in [1..N]:
		for w in [1..W]:
			// 判断能否放入背包
			if w - wt[i] < 0:	// 注意wt和val取值的时候下标-1
				dp[i][j] = dp[i-1][j]
			else:
				dp[i][j] = max(dp[i-1][w], dp[i-1][w-wt[i]]+val[i])
	return dp[N][W]
*/

package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func backpack(N, W int, wt, val []int) int {
	// 创建二维数组；
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	// base case：dp[0][..] = dp[..][0] = 0
	// 状态转移：
	for i := 1; i < N+1; i++ {
		for w := 1; w < W+1; w++ { // 当可选择的物品为前i个，背包容量为w时，背包可容纳的最大价值
			// 选择状态：
			if w-wt[i] < 0 {	// 若放不下第i件物品，则对于当前容量，前i件物品的最大价值等于前i-1件物品的最大价值
				dp[i][w] = dp[i-1][w]
			} else {			// 若放得下第i件物品，则选择放入或无放入的最大值
				dp[i][w] = max(dp[i-1][w-wt[i]] + val[i], dp[i-1][w])
			}
		}
	}

	return dp[N][W]
}

func main() {
	N := 5
	W := 100
	wt := []int{100, 77, 22, 29, 50, 99}
	val := []int{5, 92, 22, 87, 46, 90}
	wt = append([]int{0}, wt...)
	val = append([]int{0}, val...)
	res := backpack(N, W, wt, val)
	fmt.Println(res)
}
