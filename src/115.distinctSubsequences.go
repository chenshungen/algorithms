package main

import "fmt"

func numDistinct2(s string, t string) int {
	m, n := len(s), len(t)

	dp := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = 1
	}

	// dp[i] 在更新前的值，即 s[:i] 与 t[:j] 的匹配值
	// 有可能会在更新 dp[i+1] ，即 s[:i] 与 t[:j+1] 的匹配值，的时候用到
	// 利用 prev 保存起来
	var prev int

	for j := 0; j < n; j++ {
		dp[j], prev = 0, dp[j]

		for i := j + 1; i < m+1; i++ {
			// 对于 dp[i] 来说
			// s[:i] 中符合条件的子字符串，按照是否包含 s[i-1]，可以分成两个部分：
			// 第一部分：
			//     **不**包含 s[i-1]，
			//     这部分的数量，等于 dp[i-1]
			// 第二部分：
			//     包含 s[i-1]
			//     这部分，只有当 s[i-1] == t[j-1] 的时候，才存在
			//     存在的话，这部分的数量，等于 prev
			if t[j] == s[i-1] {
				dp[i], prev = dp[i-1]+prev, dp[i]
			} else {
				dp[i], prev = dp[i-1], dp[i]
			}
		}
	}

	return dp[m]
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

func testNumDistinct115() {
	fmt.Println(numDistinct2("rabbbit", "rabbit"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}
