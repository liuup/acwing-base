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

	n, q := ri(), ri()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = ri()
	}

	for i := 0; i < q; i++ {
		query := ri()
		left := bsearch1_(nums, query)
		if left >= len(nums) || nums[left] != query {
			fmt.Fprint(out, -1)
		} else {
			fmt.Fprint(out, left)
		}

		fmt.Fprint(out, " ")

		right := bsearch2_(nums, query)
		if right < 0 || nums[right] != query {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, right)
		}
	}
}

// find 5 left border; 1 2 3 {5 5 5 7 8}
// if left >= len(nums) || nums[left] != query
func bsearch1_(n []int, t int) int {
	l, r := -1, len(n)
	for l+1 < r {
		mid := (l + r) >> 1
		if n[mid] >= t {
			r = mid
		} else {
			l = mid
		}
	}
	return r // 函数外判断是否越右边界
}

// find 5 right border; {1 2 3 5 5 5} 7 8
// right < 0 || nums[right] != query
func bsearch2_(n []int, t int) int {
	l, r := -1, len(n)
	for l+1 < r {
		mid := (l + r) >> 1
		if n[mid] <= t {
			l = mid
		} else {
			r = mid
		}
	}
	return l // 函数外判断是否越左边界
}
