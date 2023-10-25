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

	n := ri()

	g := make([][]int, n+1)

	for i := 0; i < n-1; i++ {
		a, b := ri(), ri()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	st := make([]bool, n+1)

	dfs(1, st, g)

	fmt.Fprintln(out, ans)
}

var (
	ans int = 1<<31 - 1
)

func dfs(u int, st []bool, g [][]int) int {
	res := 0
	st[u] = true
	sum := 1

	for _, x := range g[u] {
		if !st[x] {
			s := dfs(x, st, g)
			res = max_i(res, s)
			sum += s
		}
	}

	res = max_i(res, len(g)-1-sum) // n-sum
	ans = min_i(res, ans)
	return sum
}

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
