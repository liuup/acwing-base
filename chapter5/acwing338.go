package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	// in := bufio.NewReader(os.Stdin)	// 搭配fmt.Scan()使用
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
	// 读一个仅包含小写字母的字符串
	rs := func() (s []byte) {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		}
		for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	// 读一个长度为 n 的仅包含小写字母的字符串
	rsn := func(n int) []byte {
		b := rc()
		// for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		// }
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}
	_ = []interface{}{rc, ri, rs, rsn}

	for {
		a, b := ri(), ri()
		if a == 0 && b == 0 {
			break
		}

		if a > b { // 交换一下
			a, b = b, a
		}

		//
		for i := 0; i < 10; i++ {
			fmt.Fprint(out, count(b, i)-count(a-1, i), " ")
		}
		fmt.Fprintln(out)
	}
}

// 1-n中 从左到右 x出现的次数 x0-9
func count(n, i int) int {
	res := 0
	d := dgt(n)

	for j := 1; j <= d; j++ { // 从右往左
		// l r 第j位左边的数 右边的数
		// dj 第j位的数字
		p := pow10(j - 1)
		l := n / p / 10
		r := n % p
		dj := (n / p) % 10

		// 计算dij位左边的整数小于1 视频中xxx = 000 ~abc-1的情况
		if i > 0 {
			res += l * p
		}
		if i == 0 && l > 0 { // 如果i=0 左边高位不能全为0 xxx=001~abc-1
			res += (l - 1) * p
		}
		// 计算第j位左边的整数等于1的情况
		if dj > i && (i > 0 || l > 0) {
			res += p
		}
		if dj == i && (i > 0 || l > 0) {
			res += (r + 1)
		}
	}
	return res
}

// 统计n有多少位
func dgt(n int) (ans int) {
	for n > 0 {
		ans++
		n /= 10
	}
	return
}

// 右边的数
func pow10(x int) (ans int) {
	ans = 1
	for ; x > 0; x-- {
		ans *= 10
	}
	return ans
}

func main() {
	_debug()
}
