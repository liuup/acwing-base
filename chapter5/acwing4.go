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

	// 读一个仅包含小写字母的字符串
	rs := func() (s []byte) {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		}
		for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	_ = []interface{}{rc, ri, rs}

	// 多重背包

	n, v := ri(), ri()

	vs, ws, ss := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		vs[i] = ri()
		ws[i] = ri()
		ss[i] = ri()
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= v; j++ {
			for k := 0; k*vs[i-1] <= j && k <= ss[i-1]; k++ { // sk是最大数量
				dp[i][j] = max_i(dp[i][j], dp[i-1][j-k*vs[i-1]]+k*ws[i-1])
			}
		}
	}
	fmt.Fprintln(out, dp[n][v])
}

// 用的比较多 这俩max min先放这吧

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_debug()
}
