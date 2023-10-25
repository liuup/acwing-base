package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	us := InitUnion(n + 1)

	var op string
	var a, b int

	for ; m > 0; m-- {
		fmt.Fscan(in, &op)
		if op == "C" {
			fmt.Fscan(in, &a, &b)
			us.Union(a, b)
		} else if op == "Q1" {
			fmt.Fscan(in, &a, &b)

			if us.IsConnected(a, b) {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		} else if op == "Q2" {
			fmt.Fscan(in, &a)

			fmt.Fprintln(out, us.Count(a))
		}
	}
}

// union set 并查集

type unionset struct {
	father []int
	count  int // 连通分量的个数

	mp map[int]int
}

func InitUnion(n int) unionset {
	fa := make([]int, n+1)
	mp := map[int]int{}
	for i := 1; i <= n; i++ {
		fa[i] = i
		mp[i] = 1
	}
	return unionset{fa, n, mp}
}

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

	u.mp[j_fa] += u.mp[i_fa]

	u.count--
}

func (u *unionset) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func (u *unionset) Count(x int) int {
	return u.mp[u.Find(x)]
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func main() {
	// ===== ===== fast io ===== =====
	in = bufio.NewReader(os.Stdin) // 搭配Fscan使用
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
