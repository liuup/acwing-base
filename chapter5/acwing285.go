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

	/// acwing 900
	n := ri()

	dp := make2dimen(n+1, 2)

	happy := make([]int, n+1)
	for i := 1; i <= n; i++ {
		happy[i] = ri()
	}

	vis := make([]bool, n+1) // 是否有上司

	boss := make([][]int, n+1) // 使用邻接表存储
	for i := 0; i < n-1; i++ {
		a, b := ri(), ri()
		boss[b] = append(boss[b], a) // b是a的上司
		vis[a] = true                // a有上司了
	}

	root := 1
	// 找到没有上司的根节点 董事长
	for i := 1; i <= n; i++ {
		if !vis[i] {
			root = i
			break
		}
	}

	dfs(root, happy, boss, dp)

	fmt.Fprintln(out, max_i(dp[root][0], dp[root][1])) // 来或不来的最大值
}

func dfs(b int, happy []int, boss [][]int, dp [][]int) {
	dp[b][0] = 0
	dp[b][1] = happy[b] // 1代表这个boss要来 先加上他的利益

	for i := 0; i < len(boss[b]); i++ {
		y := boss[b][i]

		dfs(y, happy, boss, dp)

		dp[b][0] += max_i(dp[y][0], dp[y][1])
		dp[b][1] += dp[y][0]
	}
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
