package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	path []int
	vis  = 0
)

func _debug() {
	path = []int{}
	vis = 0
	n := ri()
	back(n)
}

func back(n int) {
	if len(path) == n {
		for _, x := range path {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprint(out, "\n")
		return
	}

	for j := 1; j <= n; j++ {
		if (vis>>j)&1 == 1 {
			continue
		}
		path = append(path, j)
		vis |= (1 << j)

		back(n)

		path = path[:len(path)-1]
		vis &= ^(1 << j)
	}
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

const eof = 0

var (
	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB rc rs rsn
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
	in *bufio.Reader // 搭配Fscan使用
	// in  *bufio.Scanner
	out *bufio.Writer
)

func main() {
	in = bufio.NewReader(os.Stdin) // 搭配Fscan使用
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
