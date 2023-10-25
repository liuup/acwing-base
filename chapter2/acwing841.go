package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	n, m := ri(), ri()
	sb := make([]byte, n+1) // 前面空出来一个位置
	copy(sb[1:], rsn(n))

	shash_init(string(sb), shash_power)

	for ; m > 0; m-- { // m次
		l1, r1, l2, r2 := ri(), ri(), ri(), ri()
		if shash_equal(l1, r1, l2, r2) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

// 1000000 + 10
var shash_p [100000 + 10]int
var shash_h [100000 + 10]int

const shash_power int = 131

func shash_init(s string, power int) { // 预处理前缀和
	shash_p[0] = 1
	shash_h[0] = 0
	for i := 1; i <= len(s)-1; i++ {
		shash_p[i] = shash_p[i-1] * power
		shash_h[i] = shash_h[i-1]*power + int(s[i]-'a'+1)
	}
}

// 计算s[l~r]的hash
func shash(l, r int) int {
	return shash_h[r] - shash_h[l-1]*shash_p[r-l+1]
}

func shash_equal(l1, r1 int, l2, r2 int) bool { // 判断两子串是否相等
	return shash(l1, r1) == shash(l2, r2)
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
