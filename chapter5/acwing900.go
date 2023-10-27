package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	// 读一个字符
	rc := func() byte {
		if _i == _n {
			_n, _ = os.Stdin.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}

	// 读一个整数，支持负数
	ri := func() (x int) {
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}
	_ = []interface{}{rc, ri}

	/// acwing 900
	n := ri()

	// 二维dp
	// dp[i][j]表示前i个整数恰好拼成j的方案数
	// dp[i][j] = dp[i-1][j] + dp[i-1][j-i] + dp[i-1][j-2*i]+...
	// dp[i][j-i] = dp[i-1][j-i]+dp[i-1][j-2*i]+...
	// dp[i][j] = dp[i-1][j] + dp[i][j-i]
	mod := int64(1e9 + 7)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
		dp[i][0] = 1 // 当容量为0时，全不选也是一种方案
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ { // 容量从0开始
			dp[i][j] = dp[i-1][j] % mod
			if j >= i {
				dp[i][j] = (dp[i-1][j] + dp[i][j-i]) % mod
			}
		}
	}
	fmt.Fprintln(out, dp[n][n])
}

func main() {
	_debug()
}
