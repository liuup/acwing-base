package main

import (
	"bufio"   // io
	"fmt"     // io
	"os"      // io
	"strconv" // io
)

var (
	in  *bufio.Scanner
	out *bufio.Writer
)

func ri() int { // fast read int
	in.Scan()
	x, _ := strconv.Atoi(string(in.Bytes()))
	return x
}

func rf() float64 { // fast read float64
	in.Scan()
	f, _ := strconv.ParseFloat(string(in.Bytes()), 64)
	return f
}

// buf := make([]byte, 1e9)
// in.Buffer(buf, 1e9)
func rs() []byte { in.Scan(); return in.Bytes() } // read string

func main() {
	// ===== fast io =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	buf := make([]byte, 1e6)
	in.Buffer(buf, 1e6)

	// ===== fast io =====

	rev := func(b []byte) {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
	}

	n, m := rs(), rs()

	rev(n)
	rev(m)

	res := cmp(n, m)

	var ans []byte
	if res == 0 {
		fmt.Fprint(out, 0)
		return
	} else if res == 1 {
		ans = highsub(n, m)
	} else if res == -1 {
		ans = highsub(m, n)
		fmt.Fprint(out, "-")
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(out, string(ans[i]))
	}
}

func cmp(a, b []byte) int {
	if len(a) != len(b) {
		if len(a) > len(b) {
			return 1
		} else {
			return -1
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}

// 高精度减 a b正常顺序，ans逆向
func highsub(a, b []byte) (ans []byte) {

	for i, t := 0, 0; i < len(a); i++ {
		t = int(a[i]-'0') - t
		if i < len(b) {
			t -= int(b[i] - '0')
		}
		ans = append(ans, byte((t+10)%10+'0'))
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	for len(ans) > 1 && ans[len(ans)-1] == '0' {
		ans = ans[:len(ans)-1]
	}
	// rev(ans)	// modify
	return
}
