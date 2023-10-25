package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	n, m := ri(), ri()

	grid := make_two_dimen(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid[i][j] = ri()
		}
	}
	bfs(grid)
}

func bfs(grid [][]int) {
	q := [][]int{}

	dist := make_two_dimen_with_val(len(grid), len(grid[0]), -1)

	q = append(q, []int{0, 0})
	dist[0][0] = 0

	for len(q) != 0 {
		x := q[0][0]
		y := q[0][1]

		q = q[1:]

		for _, d := range directions {
			dx := x + d[0]
			dy := y + d[1]
			if isok(grid, dx, dy) && grid[x][y] == 0 && dist[dx][dy] == -1 {
				dist[dx][dy] = dist[x][y] + 1
				q = append(q, []int{dx, dy})
			}
		}
	}
	fmt.Fprintln(out, dist[len(grid)-1][len(grid[0])-1])
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向
func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

// 创建指定维度的二维数组
// n: rows; m: cols
func make_two_dimen(n, m int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	return
}

// 创建指定维度的二维数组，并赋特定值
func make_two_dimen_with_val(n, m, val int) (ans [][]int) {
	tmp := make([]int, m)
	for i := range tmp {
		tmp[i] = val
	}

	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		copy(ans[i], tmp)
	}
	return
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

const eof = 0

var (
	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB rc rs rsn

	outs = make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmps = [64]byte{}              // 可根据绝对值的十进制长度的上限调整
)

func rc() byte { // faster read one byte
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

func ri() (x int) { // faster read int, support negative
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

var (
	in  *bufio.Reader // 搭配Fscan使用
	out *bufio.Writer
)

func main() {
	in = bufio.NewReader(os.Stdin) // 搭配Fscan使用
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
