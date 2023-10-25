package main

import (
	"bufio" // io
	"fmt"   // io
	"math"  // io
	"os"    // io
	"sort"
	"strconv" // io
)

func iomode() {
	n, m := ri(), ri()

	inputs := [][]int{} // 0index 1val
	for i := 0; i < n; i++ {
		x, c := ri(), ri()
		inputs = append(inputs, []int{x, c})
	}
	inputs_mp := map[int]int{}
	for _, k := range inputs {
		inputs_mp[k[0]] += k[1]
	}

	lr := [][]int{} // 0left 1right
	for i := 0; i < m; i++ {
		l, r := ri(), ri()
		// lr[l] = r
		lr = append(lr, []int{l, r})
	}

	// 去重
	tmp := map[int]struct{}{}
	for _, k := range inputs {
		tmp[k[0]] = struct{}{}
	}
	for _, k := range lr {
		tmp[k[0]] = struct{}{}
		tmp[k[1]] = struct{}{}
	}

	// 放数组里排序
	alls := make([]int, len(tmp))
	i := 0
	for k := range tmp {
		alls[i] = k
		i++
	}
	sort.Ints(alls)

	// 计算前缀和
	sums := make([]int, len(alls)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + inputs_mp[alls[i-1]]
	}

	// 用哈希降低时间复杂度
	alls_mp := map[int]int{}
	for i, x := range alls { // alls -> index
		alls_mp[x] = i
	}

	for _, x := range lr {
		left := bsearch1_(alls, x[0]) + 1 // 左边界+1
		right := bsearch1_(alls, x[1]) + 1

		fmt.Fprintln(out, sums[right]-sums[left-1])
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

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

// "runtime/debug"
// func init() { debug.SetGCPercent(-1) }

const eof = 0

var (
	in  *bufio.Scanner
	out *bufio.Writer

	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB

	outs = make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmps = [64]byte{}              // 可根据绝对值的十进制长度的上限调整
)

func r() int { // fast read int
	in.Scan()
	x, _ := strconv.Atoi(string(in.Bytes()))
	return x
}

func rf() float64 { // fast read float64
	in.Scan()
	s := in.Bytes()
	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	dotPos := len(s) - 1
	f := 0
	for i, b := range s {
		if b == '.' {
			dotPos = i
		} else {
			f = f*10 + int(b&15)
		}
	}
	if neg {
		f = -f
	}
	return float64(f) / math.Pow10(len(s)-1-dotPos) // 放心，math.Pow10 会直接查表，非常快
}

// buf := make([]byte, 1e6+10)
// in.Buffer(buf, 1e6+10)
func rs() []byte { in.Scan(); return in.Bytes() } // read string

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

// suggest
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

// how to use
// wInt(1)
// os.Stdout.Write(outS)
func wint(x int) {
	if x == 0 { // 如果保证不为零则去掉
		outs = append(outs, '0')
		return
	}
	if x < 0 { // 如果保证是非负数则去掉
		x = -x
		outs = append(outs, '-')
	}
	p := len(tmps)
	for ; x > 0; x /= 10 {
		p--
		tmps[p] = '0' | byte(x%10)
	}
	outs = append(outs, tmps[p:]...)
	// outs = append(outs, '\n') // 看情况使用s
}

// ===== ===== fast io ===== =====

func main() {
	// ===== ===== fast io ===== =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// sbuf := make([]byte, 1e6) // read string buffer init
	// in.Buffer(sbuf, 1e6)
	// ===== ===== fast io ===== =====

	iomode()
}
