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

	n1, n2, m := ri(), ri(), ri()

	g := make([][]int, n1+n2+1)

	for ; m > 0; m-- {
		a, b := ri(), ri()
		g[a] = append(g[a], b) // 存一边就够了
	}

	st := make([]int, n1+n2+1)    // 标记是否找到了c
	match := make([]int, n1+n2+1) // match[x] 和x匹配的编号

	tmp := make([]int, n1+n2+1) // 辅助数组

	ans := 0
	for i := 1; i <= n1; i++ { // 从n1去找n2
		copy(st, tmp) // 归零

		if find(i, st, match, g) {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}

func find(i int, st, match []int, g [][]int) bool {
	for _, x := range g[i] { // 遍历所有孩子
		if st[x] == 0 {
			st[x] = 1
			if match[x] == 0 || find(match[x], st, match, g) {
				match[x] = i
				return true
			}
		}
	}
	return false
}

func main() {
	_debug()
}
