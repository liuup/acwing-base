package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	n, m := ri(), ri()
	us := InitUnion(n + 1)

	for ; m > 0; m-- {
		op, a, b := string(rs()), ri(), ri()
		if op == "M" {
			us.Union(a, b)
		} else if op == "Q" {
			if us.IsConnected(a, b) {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}

// union set 并查集

type unionset struct {
	father []int
	// father map[int]int
	count int // 连通分量的个数
}

// type unionset struct {
// 	father map[int]int
// 	count  int // 连通分量的个数
// }

func InitUnion(n int) unionset {
	fa := make([]int, n)
	for i := 1; i < len(fa); i++ {
		fa[i] = i
	}
	return unionset{fa, n}
}

// func InitUnion(n int) unionset {	// 有时候数据比较离散，可以考虑用哈希
// 	us := unionset{make(map[int]int)}
// 	for _, x := range nums {
// 		us.father[x] = x
// 	}
// 	return us
// }

func (u *unionset) Find(i int) int {
	if u.father[i] == i {
		return i
	}
	u.father[i] = u.Find(u.father[i]) // 路径压缩
	return u.father[i]
}

func (u *unionset) Union(i, j int) {
	i_fa, j_fa := u.Find(i), u.Find(j)
	if i_fa == j_fa {
		return
	}
	u.father[i_fa] = j_fa
	u.count--
}

func (u *unionset) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
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
