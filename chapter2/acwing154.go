package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	n, k := ri(), ri()

	nums := make([]int, n)
	for i := range nums {
		nums[i] = ri()
	}

	q := []int{}
	mins := []int{}
	maxs := []int{}

	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[i] <= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		left := i - k + 1
		for q[0] < left {
			q = q[1:]
		}
		if i >= k-1 {
			mins = append(mins, nums[q[0]])
		}
	}

	q = []int{}
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		left := i - k + 1
		for q[0] < left {
			q = q[1:]
		}
		if i >= k-1 {
			maxs = append(maxs, nums[q[0]])
		}
	}
	for _, x := range mins {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
	for _, x := range maxs {
		fmt.Fprint(out, x, " ")
	}
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

const eof = 0

var (
	out         *bufio.Writer
	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB
)

func rc() byte { // faster read one byte
	if _i == _n {
		_n, _ = os.Stdin.Read(buf)
		if _n == 0 { // EOF
			return eof
		}
		_i = 0
	}
	b := buf[_i]
	_i++
	return b
}

// 读一个仅包含小写/大写字母的字符串，必要时进行修改
func rs() (s []byte) {
	b := rc()
	for ; 'A' > b || b > 'Z'; b = rc() { // 'A' 'Z'	// 把非必要字符过滤掉
	}
	for ; 'A' <= b && b <= 'Z'; b = rc() { // 'A' 'Z'
		s = append(s, b)
	}
	return
}

// 读一个长度为 n 的仅包含小写/大写字母的字符串，必要时进行修改
func rsn(n int) []byte {
	b := rc()
	// 只读取大小写字母和数字
	for ; !('a' <= b && b <= 'z') && !('A' <= b && b <= 'Z') && !('0' <= b && b <= '9'); b = rc() { // 'A' 'Z'
	}
	s := make([]byte, 0, n)
	s = append(s, b)
	for i := 1; i < n; i++ {
		s = append(s, rc())
	}
	return s
}

func ri() (x int) { // faster read int, support negative
	neg := false
	b := rc()
	for ; '0' > b || b > '9'; b = rc() {
		// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
		if b == eof {
			return
		}
		if b == '-' {
			neg = true
		}
	}
	for ; '0' <= b && b <= '9'; b = rc() {
		x = x*10 + int(b&15)
	}
	if neg {
		return -x
	}
	return
}

func main() {
	// ===== ===== fast io ===== =====
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
