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

func rs() []byte { in.Scan(); return in.Bytes() } // read string

func main() {
	// ===== fast io =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	buf := make([]byte, 1e6+10) // string input buffer
	in.Buffer(buf, 1e6+10)
	// ===== fast io =====

	n, m := rs(), rs()

	ans := string(highadd(n, m))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(out, string(ans[i]))
	}
}

func highadd(a []byte, b []byte) (ans []byte) {
	rev := func(b []byte) {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
	}
	if len(a) < len(b) {
		return highadd(b, a)
	}

	rev(a)
	rev(b)

	t := 0
	for i := 0; i < len(a); i++ {
		t += int(a[i] - '0')
		if i < len(b) {
			t += int(b[i] - '0')
		}
		ans = append(ans, byte(t%10+'0'))
		t /= 10
	}
	if t > 0 {
		ans = append(ans, byte(t+'0'))
	}
	return
}
