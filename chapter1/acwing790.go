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

func main() {
	// ===== fast io =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== fast io =====

	n := rf()
	fmt.Fprintf(out, "%.6f\n", floatBsearch_(n))
}

// 浮点二分 求一个数的三次方根，将就改改
func floatBsearch_(y float64) float64 {
	l, r := float64(-100), float64(100)
	for r-l > 1e-7 {
		mid := (l + r) / float64(2)
		if mid*mid*mid <= y {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
