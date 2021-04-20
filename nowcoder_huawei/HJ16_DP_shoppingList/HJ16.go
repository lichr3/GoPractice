/* HJ16_DP_shoppingList.购物单 动态规划问题
状态：
	1.两个状态：【预算】【可选择的物品（主件和附件的组合）】
	2.表示方法：dp[i][w]，对于前i个主件（主件和附件的4种组合情况），预算为w时，能购买物品的最大价值（物品价格乘以重要程度）
选择：
	1.预算w时，对于第i个物品，选择是否购买
	2.选择不购买：dp[i][w] = dp[i-1][w]
	3.选择购买： 	dp[i][w] = max(选择组合j的价值，不选择组合j的价值) （有4种组合可能）
	4.若该物品能放入背包，则选择3和4结果的最大值
base case:
	dp[0][..] = 0; dp[..][0] = 0
限制：
	相较于0-1背包问题，多出的限制是，只有购买了主件才能购买附件。
	即遇到主件时候，要考虑主件和附件的4种组合情况，遇到附件的时候，直接dp[i][w]=dp[i-1][w]即可。
注意事项：
	1.注意题目给出的有些限制，使得我们更好处理。例如最多60个物品，2个附件。则可以使用两个长为61*3的数组分别表示每个物品的价格和价值（价格*重要等级）
伪代码：
	总共有N个主件，预算为W；
	wt[i][k]：k=0表示主件价格，k=1表示附件1价格，k=2表示附件2价格；
	val[i][k]：k=0表示主件价值，k=1表示附件1价值，k=2表示附件2价值；（价格*重要程度）
	int dp[N+1][W+1]
	dp[0][..] = 0
	dp[..][0] = 0
	for i in [1..N]:
		for w in [1..W]:
			if val[i][0] == 0: dp[i][w] = dp[i-1][w]; continue // val[i][0]为空，i为附件,赋值后continue
			// 把主件i与其附件的4种组合遍历：[0] [0,1] [0,2] [0,1,2]
			val0 = val[i][0]; wt0 = wt[i][0]
			val1 = val[i][1]; wt1 = wt[i][1]
			val2 = val[i][3]; wt2 = wt[i][3]
			if w - wt0 < 0: dp[i][w] = dp[i-1][w] continue	// 组合1
			else: dp[i][w] = max(dp[i-1][w-wt0] + val0, dp[i-1][w])
			if w - (wt0 + wt1) >= 0: dp[i][w] = max(dp[i-1][w-wt0-wt1] + val0 + val1, dp[i][w])	// 组合2
			if w - (wt0 + wt2) >= 0: dp[i][w] = max(dp[i-1][w-wt0-wt2] + val0 + val2, dp[i][w]) / /组合3
			if w - (wt0 + wt1 + wt2): dp[i][w] = max(dp[i-1][w-wt0-wt1-wt2] + val0 + val1 + val2, dp[i][w]) // 组合3
	return dp[N][W]
*/
package main

import "fmt"

func main() {
	var N, m int // 总预算为N，可选择的物品数有m个
	fmt.Scan(&N, &m)
	wt := make([][3]int, 61)	// 存储价格
	val := make([][3]int, 61)	// 存储价值

	// 存储工作
	for i := 1; i < m+1; i++ {	// 下标从1开始
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		// 判断主件还是附件，然后把价格和价值存到相应位置
		if c == 0 {	// 若为主件i
			wt[i][0] = a; val[i][0] = a * b
		} else if wt[c][1] == 0 {	// 若为主件i的附件
			wt[c][1] = a; val[c][1] = a * b
		} else {
			wt[c][2] = a; val[c][2] = a * b
		}
	}

	// 建立dp二维数组存储结果
	dp := make([][]int, m+1)  		// m 个物品
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, N+1)	// 总预算为N
	}

	// 状态转移和选择过程
	for i := 1; i < m+1; i++ {
		for w := 1; w < N+1; w++ {
			if wt[i][0] == 0 { dp[i][w] = dp[i-1][w]; continue } 	// 若为附件则跳过
			// 把主件i与其附件的4种组合遍历：[0] [0,1] [0,2] [0,1,2]
			wt0 := wt[i][0]; val0 := val[i][0]
			wt1 := wt[i][1]; val1 := val[i][1]
			wt2 := wt[i][2]; val2 := val[i][2]
			// 组合1
			if w - wt0 < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w-wt0] + val0, dp[i-1][w])
			}
			// 组合2。 注意，此时dp[i][w]已有取值，所以取max的时候应该将之前的dp[i-1][w]换成dp[i][w]
			if w - wt0 - wt1 >= 0 { dp[i][w] = max(dp[i-1][w-wt0-wt1] + val0 + val1, dp[i][w]) }
			if w - wt0 - wt2 >= 0 { dp[i][w] = max(dp[i-1][w-wt0-wt2] + val0 + val2, dp[i][w]) }	// 组合3
			if w - wt0 - wt1 - wt2 >= 0 { dp[i][w] = max(dp[i-1][w-wt0-wt1-wt2] + val0 + val1 + val2, dp[i][w]) } // 组合4
		}
	}

	fmt.Println(dp[m][N])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
