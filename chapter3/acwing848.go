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

	mp := map[int][]int{}
	for ; m > 0; m-- {
		a, b := ri(), ri()
		mp[a] = append(mp[a], b)
	}

	in := make([]int, n+1)
	for _, row := range mp {
		for _, x := range row {
			in[x]++
		}
	}

	q := []int{}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	count := 0
	path := []int{}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		count++
		path = append(path, cur)

		for _, next := range mp[cur] {
			in[next]--
			if in[next] == 0 {
				q = append(q, next)
			}
		}
	}

	if count == n {
		for _, x := range path {
			fmt.Fprint(out, x, " ")
		}
	} else {
		fmt.Fprintln(out, -1)
	}
}

func main() {
	_debug()
}
