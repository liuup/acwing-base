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

	// 多源汇 floyd最短路
	n, m, k := ri(), ri(), ri()

	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = 1<<31 - 1
			}
		}
	}

	for ; m > 0; m-- {
		x, y, z := ri(), ri(), ri()
		dist[x][y] = min_i(dist[x][y], z) // 保存最小的边
	}

	// floyd
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				dist[i][j] = min_i(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	for ; k > 0; k-- { // k次询问
		a, b := ri(), ri()

		if dist[a][b] > (1<<31-1)/2 {
			fmt.Fprintln(out, "impossible")
		} else {
			fmt.Fprintln(out, dist[a][b])
		}
	}
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
