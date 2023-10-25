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

	n, m, q := ri(), ri(), ri()
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, m)
	}

	for i := range nums {
		for j := range nums[0] {
			nums[i][j] = ri()
		}
	}

	sums := partial_sum_matrix(nums)

	for i := 0; i < q; i++ {
		l1, r1, l2, r2 := ri(), ri(), ri(), ri()
		fmt.Fprintln(out, sums[l2][r2]-sums[l1-1][r2]-sums[l2][r1-1]+sums[l1-1][r1-1])
	}
}

// 二维前缀和，补0
func partial_sum_matrix(vals [][]int) (ans [][]int) {
	ans = make([][]int, len(vals)+1)
	for i := range ans {
		ans[i] = make([]int, len(vals[0])+1)
	}
	for i := 1; i <= len(vals); i++ {
		for j := 1; j <= len(vals[0]); j++ {
			ans[i][j] = ans[i-1][j] + ans[i][j-1] - ans[i-1][j-1] + vals[i-1][j-1]
		}
	}
	return
}
