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

	// 分组背包问题
	n, m := ri(), ri()

	v := make2dimen(N, N) // 体积
	w := make2dimen(N, N) // 价值
	s := make([]int, N)   // 代表第i组物品的个数

	f := make2dimen(N, N)

	for i := 1; i <= n; i++ {
		s[i] = ri()
		for j := 0; j < s[i]; j++ {
			v[i][j] = ri()
			w[i][j] = ri()
		}
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j] // 不选
			for k := 0; k < s[i]; k++ {
				if j >= v[i][k] {
					f[i][j] = max_i(f[i][j], f[i-1][j-v[i][k]]+w[i][k])
				}
			}
		}
	}

	fmt.Fprintln(out, f[n][m])
}

const N = 110

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
