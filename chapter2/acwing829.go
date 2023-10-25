package main

import (
	"bufio" // io
	"fmt"
	"os" // io
)

func _debug() {
	queue := []int{}

	n := ri()

	for i := 0; i < n; i++ {
		op := string(rs())

		if op == "push" {
			x := ri()
			queue = append(queue, x)
		} else if op == "pop" {
			queue = queue[1:]
		} else if op == "empty" {
			if len(queue) != 0 {
				fmt.Fprintln(out, "NO")
			} else {
				fmt.Fprintln(out, "YES")
			}
		} else if op == "query" {
			fmt.Fprintln(out, queue[0])
		}
	}
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/
const eof = 0

var (
	out *bufio.Writer

	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB
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

// 读一个仅包含小写字母的字符串，必要时进行修改
func rs() (s []byte) {
	b := rc()
	for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
	}
	for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		s = append(s, b)
	}
	return
}

// suggest
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

func main() {
	// ===== ===== fast io ===== =====
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
