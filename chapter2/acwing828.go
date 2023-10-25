package main

import (
	"bufio" // io
	"fmt"
	"os" // io
)

func _debug() {
	stk := []int{}

	n := ri()

	for i := 0; i < n; i++ {
		op := string(rs())

		if op == "push" {
			x := ri()
			stk = append(stk, x)
		} else if op == "pop" {
			stk = stk[:len(stk)-1]
		} else if op == "empty" {
			if len(stk) != 0 {
				fmt.Fprintln(out, "NO")
			} else {
				fmt.Fprintln(out, "YES")
			}
		} else if op == "query" {
			fmt.Fprintln(out, stk[len(stk)-1])
		}
	}
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

const eof = 0

var (
	in  *bufio.Scanner
	out *bufio.Writer

	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB
)

// 读一个仅包含小写字母的字符串
func rs() (s []byte) {
	b := rc()
	for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
	}
	for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		s = append(s, b)
	}
	return
}

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

// ===== ===== fast io ===== =====

func main() {
	// ===== ===== fast io ===== =====
	in = bufio.NewScanner(os.Stdin)
	// in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
