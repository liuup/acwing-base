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

	n := ri()
	mergetmp = make([]int, n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = ri()
	}

	mergesort(nums, 0, n-1)

	fmt.Fprintln(out, result)
}

var mergetmp []int

var result int = 0

// nums, 0, len(nums)-1
func mergesort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergesort(nums, l, mid)
	mergesort(nums, mid+1, r)

	k := 0
	i, j := l, mid+1
	for ; i <= mid && j <= r; k++ {
		if nums[i] <= nums[j] {
			mergetmp[k] = nums[i]
			i++
		} else {
			mergetmp[k] = nums[j]
			j++
			result += (mid - i + 1) // 求逆序对
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		mergetmp[k] = nums[i]
	}
	for ; j <= r; j, k = j+1, k+1 {
		mergetmp[k] = nums[j]
	}
	for i, j := l, 0; i <= r; i, j = i+1, j+1 { // 全部数组
		nums[i] = mergetmp[j]
	}
}
