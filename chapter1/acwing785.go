package main

import (
	"bufio"   // io
	"fmt"     // Fprintln(out)
	"os"      //io
	"strconv" //io
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

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = ri()
	}
	quicksort(nums, 0, len(nums)-1)

	for _, x := range nums {
		fmt.Fprint(out, x, " ")
	}
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := nums[(l+r)>>1]
	for i < j {
		i++
		for ; nums[i] < x; i++ { // 可以理解让左边都小于x
		}
		j--
		for ; nums[j] > x; j-- {
		}

		if i < j {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	quicksort(nums, l, j)
	quicksort(nums, j+1, r)
}
