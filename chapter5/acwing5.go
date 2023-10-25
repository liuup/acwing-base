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
	n, m := ri(), ri()

	v := make([]int, 20010) // 空间不好判断 手动开大点吧
	w := make([]int, 20010)

	cnt := 0
	for i := 1; i <= n; i++ {
		a, b, s := ri(), ri(), ri()

		k := 1
		for k <= s {
			cnt++
			v[cnt] = a * k
			w[cnt] = b * k
			s -= k
			k *= 2
		}
		if s > 0 {
			cnt++
			v[cnt] = a * s
			w[cnt] = b * s
		}
	}
	n = cnt

	dp := make([]int, m+10)
	// 01 bag
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			dp[j] = max_i(dp[j], dp[j-v[i]]+w[i])
		}
	}
	fmt.Fprintln(out, dp[m])
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
