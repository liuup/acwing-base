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
	rc := func() byte {                      // 读一个字符
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
	ri := func() (x int) { // 读一个整数，支持负数
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
	rs := func() (s []byte) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() { // 'A' 'Z'
		}
		// for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		// 	s = append(s, b)
		// }
		for ; '0' <= b && b <= '9'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	_ = []interface{}{rc, ri, rs}

	n := ri()

	rgs := make([][]int, 0, n) // 预分配内存
	for i := 0; i < n; i++ {
		rgs = append(rgs, []int{ri(), ri()})
	}

	sort.Slice(rgs, func(i, j int) bool {
		return rgs[i][1] < rgs[j][1]
	})

	res := 0
	end := -1 << 31

	for i := 0; i < n; i++ {
		if end >= rgs[i][0] && end <= rgs[i][1] {
			continue
		} else {
			// 选中区间数+1 更新右端点
			res++
			end = rgs[i][1]
		}
	}
	fmt.Fprintln(out, res)
}

func main() {
	_debug()
}
