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
	// ===== fast io =====

	n, m := ri(), ri()

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = ri()
	}

	ans := partial_sum(nums)

	for i := 0; i < m; i++ {
		l, r := ri(), ri()
		fmt.Fprintln(out, ans[r]-ans[l-1])
	}
}

// 一维前缀和，补0
func partial_sum(vals []int) (ans []int) {
	ans = make([]int, len(vals)+1)
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] + vals[i-1]
	}
	return
}
