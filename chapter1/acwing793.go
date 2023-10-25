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

	buf := make([]byte, 1e6) // read string buffer
	in.Buffer(buf, 1e6)
	// ===== fast io =====

	rev := func(b []byte) {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
	}

	n, m := rs(), ri()

	rev(n)

	ans := highmul(n, m)
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(out, string(ans[i]))
	}
}

// 高精度乘低精度
func highmul(a []byte, b int) (ans []byte) {
	t := 0
	for i := 0; i < len(a) || t > 0; i++ {
		if i < len(a) {
			t += int(a[i]-'0') * b
		}
		ans = append(ans, byte(t%10+'0'))
		t /= 10
	}
	for len(ans) > 1 && ans[len(ans)-1] == '0' {
		ans = ans[:len(ans)-1]
	}
	return ans
}
