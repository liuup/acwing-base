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
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			weight[i][j] = ri()
		}
	}

	for i, v := range f {
		for j := range v {
			f[i][j] = 1<<31 - 1
		}
	}

	// 从0开始 那么状态为 ...001 路径长度为0
	f[1][0] = 0

	for i := 0; i < 1<<n; i++ { // 枚举所有状态
		for j := 0; j < n; j++ { // n个点
			if i>>j&1 == 1 { // 如果这个点经过了
				for k := 0; k < n; k++ { // 前面的点
					if i>>k&1 == 1 { // k这个点经过了

						/*
							假设
							i = 100100 六个点
							j = 2
							1<<j = 100

							i-(1<<j) 把这j的状态减掉 拿到之前的状态
						*/

						f[i][j] = min_i(f[i][j], f[i-(1<<j)][k]+weight[k][j])
					}
				}
			}
		}
	}
	fmt.Fprintln(out, f[(1<<n)-1][n-1])
}

const N int = 20
const M int = 1 << 20

var (
	f      [M][N]int
	weight [N][N]int
)

/*
哪些点被用过
目前停在哪个点


状态表示
state表示当前点没用过 j表示最后停在了哪个点上
f[state][j] = f[state_k][k] + weight[k][j], state_k = state除掉之后的集合

集合


属性
路径总长度的最小值

状态压缩 state用位运算表示一个路径

state=10011

f[state][j] = min(f[state])


*/

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_debug()
}
