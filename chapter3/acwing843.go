package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	n := ri()
	tmp := make([]byte, n)
	for i := range tmp {
		tmp[i] = '.'
	}

	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		copy(grid[i], tmp)
	}

	dfs(grid, 0, n)
}

func dfs(grid [][]byte, row, n int) {
	if row == n {
		for _, row := range grid {
			fmt.Fprintln(out, string(row))
		}
		fmt.Fprintln(out)
		return
	}

	for i := 0; i < n; i++ {
		if isvalid(grid, row, i, n) {
			grid[row][i] = 'Q'
			dfs(grid, row+1, n)
			grid[row][i] = '.'
		}
	}
}

func isvalid(grid [][]byte, row, col int, n int) bool {
	// 检查列
	for i := 0; i < row; i++ {
		if grid[i][col] == 'Q' {
			return false
		}
	}

	// 检查斜角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if grid[i][j] == 'Q' {
			return false
		}
	}

	// 检查斜角
	// 检查斜角
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if grid[i][j] == 'Q' {
			return false
		}
	}
	return true
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
