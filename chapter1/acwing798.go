package main

import (
	"bufio" // io
	_ "fmt" // io
	"math"  // io
	"os"    // io
	// io
)

// ===== ===== fast io ===== =====
// golang fast io from 0x3F(https://github.com/EndlessCheng/codeforces-go/)

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

// func ri() int { // fast read int
// 	in.Scan()
// 	x, _ := strconv.Atoi(string(in.Bytes()))
// 	return x
// }

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
	// outs = append(outs, '\n') // mofify
}

// ===== ===== fast io ===== =====

func main() {
	// ===== ===== fast io ===== =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// buf := make([]byte, 1e6) // read string buffer init
	// in.Buffer(buf, 1e6)

	// ===== ===== fast io ===== =====

	n, m, q := ri(), ri(), ri()
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nums[i][j] = ri()
		}
	}

	subs := subs_mat(nums)

	for i := 0; i < q; i++ {
		l1, r1, l2, r2, k := ri(), ri(), ri(), ri(), ri()

		subs[l1-1][r1-1] += k
		subs[l2][r2] += k

		subs[l2][r1-1] -= k
		subs[l1-1][r2] -= k
	}

	subs_rec_mat(nums, subs)

	for _, row := range nums {
		for _, x := range row {
			// fmt.Fprint(out, x, " ")
			wint(x)
			outs = append(outs, ' ')
			// os.Stdout.Write(outs)
		}
		// fmt.Fprint(out, "\n")
		outs = append(outs, '\n')
	}

	os.Stdout.Write(outs)
}

// 二维差分数组 尾补0
// subs[l1-1][r1-1] += k
// subs[l2][r2] += k
// subs[l2][r1-1] -= k
// subs[l1-1][r2] -= k
func subs_mat(nums [][]int) (ans [][]int) {
	n, m := len(nums), len(nums[0])
	ans = make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, m+1)
	}
	ans[0][0] = nums[0][0]
	for j := 1; j < m; j++ { // 第一行
		ans[0][j] = nums[0][j] - nums[0][j-1]
	}
	for i := 1; i < n; i++ { // 第一列
		ans[i][0] = nums[i][0] - nums[i-1][0]
	}
	for i := 1; i < n; i++ { // 其他
		for j := 1; j < m; j++ {
			ans[i][j] = nums[i][j] - nums[i-1][j] - nums[i][j-1] + nums[i-1][j-1]
		}
	}
	return
}

// 二维差分恢复（原地）
func subs_rec_mat(nums, subs [][]int) {
	n, m := len(nums), len(nums[0])
	nums[0][0] = subs[0][0]
	for j := 1; j < m; j++ { // 第一行
		nums[0][j] = nums[0][j-1] + subs[0][j]
	}
	for i := 1; i < n; i++ { // 第一列
		nums[i][0] = nums[i-1][0] + subs[i][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			nums[i][j] = nums[i-1][j] + nums[i][j-1] - nums[i-1][j-1] + subs[i][j]
		}
	}
}
