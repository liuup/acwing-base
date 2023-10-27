package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	a := [][]int{}

	for i := 0; i < n; i++ {
		l, r := ri(), ri()
		a = append(a, []int{l, r})
	}

	// 按右端点排序
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})

	// res答案 end当前选的点
	res := 0
	end := -1 << 31
	for i := 0; i < n; i++ {
		if end >= a[i][0] && end <= a[i][1] {
			continue
		} else {
			// 选的点+1 选的点更新为区间右端点
			res++
			end = a[i][1]
		}
	}
	fmt.Fprintln(out, res)
}

func main() {
	_debug()
}
