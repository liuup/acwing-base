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

	n, m := ri(), ri()

	colors := make([]int, n+1) // 0未染色 1红色 2黑色

	es := make([][]int, n+1)

	for ; m > 0; m-- {
		a, b := ri(), ri()
		es[a] = append(es[a], b)
		es[b] = append(es[b], a)
	}

	for i := 1; i <= n; i++ { // 开始染色
		if colors[i] == 0 {
			if !dfs(colors, 1, i, es) {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")

}

// 染色该点 并递归处理与之临近的点
func dfs(colors []int, color, u int, es [][]int) bool {
	colors[u] = color

	// 遍历与u邻近的点
	for _, x := range es[u] {
		if colors[x] == 0 { // 邻近点没有染色
			if !dfs(colors, 3-color, x, es) {
				return false
			}
		} else if colors[x] != 3-color { // 已经染色 判断是否冲突
			return false
		}
	}
	return true
}

func main() {
	_debug()
}
