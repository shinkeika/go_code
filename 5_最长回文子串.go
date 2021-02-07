package main

func longestPalindrome(s string) string {
	// 动态规划
	// 创建二维数组
	length := len(s)
	var ans = ""
	dp := make([][]int, length)
	for index := 0; index < length; index++ {
		dp[index] = make([]int, length)
	}
	for l := 0; l < length; l++ {
		for i := 0; i+l < length; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if dp[i+1][j-1] == 1 && s[i] == s[j] {
					dp[i][j] = 1
				}
			}
			if dp[i][j] == 1 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

func testLongestPalindrome() {
	longestPalindrome("babad")
}
