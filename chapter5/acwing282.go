package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	// in := bufio.NewReader(os.Stdin)	// 搭配fmt.Scan()使用
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
	// 读一个长度为 n 的仅包含小写字母的字符串
	rsn := func(n int) []byte {
		b := rc()
		// for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		// }
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}
	_ = []interface{}{rc, ri, rs, rsn}

	n := ri()

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = ri()
		s[i] += s[i-1]
	}

	dp := make2dimen(n+1, n+1)

	// dp
	// 先枚举区间长度
	for length := 2; length <= n; length++ {
		// 左端点
		for i := 1; i+length-1 <= n; i++ {
			j := i + length - 1 // 右端点
			dp[i][j] = 1e8
			for k := i; k < j; k++ {
				dp[i][j] = min_i(dp[i][j], dp[i][k]+dp[k+1][j]+s[j]-s[i-1])
			}
		}
	}
	fmt.Fprintln(out, dp[1][n])
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 创建指定维度的二维数组
// n: rows; m: cols
func make2dimen(n, m int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	return
}

func main() {
	_debug()
}
