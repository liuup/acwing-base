package main

import (
	"bufio"
	"os"
)

func _debug() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	// 读一个字符
	rc := func() byte {
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

	// 读一个整数，支持负数
	ri := func() (x int) {
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
	_ = []interface{}{rc, ri}

	// 手写输出，适用于有大量（~1e6）输出的场景，CF 上可以再快 60~90ms
	// 使用前 https://codeforces.com/contest/1208/submission/176961129
	// 使用后 https://codeforces.com/contest/1208/submission/176963572
	outS := make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmpS := [20]byte{}              // 可根据绝对值的十进制长度的上限调整
	wInt := func(x int) {
		if x == 0 { // 如果保证不为零则去掉
			outS = append(outS, '0')
			return
		}
		if x < 0 { // 如果保证是非负数则去掉
			x = -x
			outS = append(outS, '-')
		}
		p := len(tmpS)
		for ; x > 0; x /= 10 {
			p--
			tmpS[p] = '0' | byte(x%10)
		}
		outS = append(outS, tmpS[p:]...)
	}

	// acwing 154
	n, k := ri(), ri()

	nums := make([]int, n)
	for i := range nums {
		nums[i] = ri()
	}

	// 找区间最大值
	mmax := make([]int, 0, n-k+1)
	q := []int{}
	for i := 0; i < n; i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		left := i - k + 1
		for q[0] < left {
			q = q[1:]
		}
		if left >= 0 {
			mmax = append(mmax, nums[q[0]])
		}
	}

	// 找区间最小值
	mmin := make([]int, 0, n-k+1)
	q = []int{}
	for i := 0; i < n; i++ {
		for len(q) > 0 && nums[i] <= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		left := i - k + 1
		for q[0] < left {
			q = q[1:]
		}
		if left >= 0 {
			mmin = append(mmin, nums[q[0]])
		}
	}

	for _, x := range mmin {
		wInt(x)
		outS = append(outS, ' ')
	}
	outS = append(outS, '\n')

	for _, x := range mmax {
		wInt(x)
		outS = append(outS, ' ')
	}
	os.Stdout.Write(outS)
}

func main() {
	_debug()
}
