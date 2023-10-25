package main

import (
	"bufio" // io
	"fmt"   // io
	"os"    // io
)

func _debug() {
	n, m := ri(), ri()
	n1 := make([]int, n)
	n2 := make([]int, m)

	for i := 0; i < n; i++ {
		n1[i] = ri()
	}
	for i := 0; i < m; i++ {
		n2[i] = ri()
	}

	i := 0
	for j := 0; j < m; j++ {
		if i < n && n1[i] == n2[j] {
			i++
		}
	}
	if i == n {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

// "runtime/debug"
// func init() { debug.SetGCPercent(-1) }

const eof = 0

var (
	in  *bufio.Scanner
	out *bufio.Writer

	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB
)

// buf := make([]byte, 1e6+10)
// in.Buffer(buf, 1e6+10)
func rs() []byte { in.Scan(); return in.Bytes() } // read string

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
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// sbuf := make([]byte, 1e6) // read string buffer init
	// in.Buffer(sbuf, 1e6)
	// ===== ===== fast io ===== =====

	_debug()
}
