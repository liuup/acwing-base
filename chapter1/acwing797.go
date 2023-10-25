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
	// ===== ===== fast io ===== =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====

	n, m := ri(), ri()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = ri()
	}

	subs := subs_arr(nums)

	for i := 0; i < m; i++ {
		l, r, c := ri(), ri(), ri()
		subs[l-1] += c
		subs[r] -= c
	}

	subs_rec_arr(nums, subs)

	for _, x := range nums {
		fmt.Fprint(out, x, " ")
	}
}

// 一维差分数组 尾补0
// subs[l-1] += c
// subs[r] -= c
func subs_arr(nums []int) (ans []int) {
	ans = make([]int, len(nums)+1)
	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i] - nums[i-1]
	}
	return
}

// 一维差分恢复（原地）
func subs_rec_arr(nums, subs []int) {
	n := len(nums)
	nums[0] = subs[0]
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + subs[i]
	}
}
