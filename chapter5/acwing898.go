package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	// in := bufio.NewReader(os.Stdin)
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

	// line dp
	n := ri()

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = ri()
		}
	}

	// left border
	for i := 1; i < n; i++ {
		dp[i][0] += dp[i-1][0]
	}
	// right border
	for i := 1; i < n; i++ {
		dp[i][len(dp[i])-1] += dp[i-1][len(dp[i-1])-1]
	}
	// other
	for i := 1; i < n; i++ {
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] += max_i(dp[i-1][j-1], dp[i-1][j])
		}
	}
	// find max and print
	mm := -1 << 31
	for i := 0; i < n; i++ {
		mm = max_i(mm, dp[n-1][i])
	}
	fmt.Fprintln(out, mm)
}

// 用的比较多 这俩max min先放这吧

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	_debug()
}
