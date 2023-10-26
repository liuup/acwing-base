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
	ri := func() (x int64) {
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
			x = x*10 + int64(b&15)
		}
		if neg {
			return -x
		}
		return
	}
	_ = []interface{}{rc, ri}

	// acwing 125
	n := int(ri())

	type milk struct {
		ws int64
		s  int64
	}

	ms := make([]milk, n)

	for i := 0; i < n; i++ {
		a, b := ri(), ri()
		ms[i] = milk{a + b, b}
	}

	// 按ws从大到小进行排列
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].ws < ms[j].ws
	})

	ans := int64(-1 << 31)
	sum := int64(0)
	for i := 0; i < n; i++ {
		sum -= ms[i].s
		ans = int64(max_i64(ans, sum))
		sum += ms[i].ws
	}
	fmt.Fprintln(out, ans)
}

func max_i64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func main() {
	_debug()
}
