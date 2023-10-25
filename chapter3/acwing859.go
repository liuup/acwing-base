package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n, m := ri(), ri()

	edges := []edge{}

	for ; m > 0; m-- {
		a, b, c := ri(), ri(), ri()
		edges = append(edges, edge{a, b, c})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].val < edges[j].val
	})

	us := InitUnion(n)
	ans := 0
	for i := range edges {
		if !us.IsConnected(edges[i].from, edges[i].to) {
			ans += edges[i].val
			us.Union(edges[i].from, edges[i].to)
		}
	}

	if us.count != 1 {
		fmt.Fprintln(out, "impossible")
	} else {
		fmt.Fprintln(out, ans)
	}
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type edge struct {
	from, to int
	val      int
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
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
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

func main() {
	_debug()
}
