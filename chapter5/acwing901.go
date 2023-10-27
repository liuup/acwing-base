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

	// acwing 901
	n, m := ri(), ri()

	memo = make2dimen(n, m)

	grid := make2dimen(n, m)
	for i, v := range grid {
		for j := range v {
			grid[i][j] = ri()
		}
	}

	ans := -1
	for i, v := range grid {
		for j := range v {
			ans = max_i(ans, dfs(grid, i, j))
		}
	}

	fmt.Fprintln(out, ans)
}

func dfs(grid [][]int, i, j int) int {
	if memo[i][j] != 0 { // 如果已经访问过了 直接返回
		return memo[i][j]
	}

	memo[i][j]++

	for _, d := range directions {
		di := i + d[0]
		dj := j + d[1]
		if isok(grid, di, dj) && grid[i][j] > grid[di][dj] {
			memo[i][j] = max_i(memo[i][j], dfs(grid, di, dj)+1)
		}
	}

	return memo[i][j]
}

var memo [][]int
var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向

func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
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
