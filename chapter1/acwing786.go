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
	n, k := ri(), ri()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = ri()
	}
	quicksort(nums, 0, n-1)
	fmt.Fprintln(out, nums[k-1])
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := nums[(l+r)>>1]
	for i < j {
		i++
		for ; nums[i] < x; i++ {
		}
		j--
		for ; nums[j] > x; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quicksort(nums, l, j)
	quicksort(nums, j+1, r)
}
