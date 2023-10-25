package main

import (
	"bufio"
	"fmt"
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

	n, m := ri(), ri() // n个点 m个边 s起点

	graph := make([][]edge, n+1)
	for ; m > 0; m-- {
		a, b, c := ri(), ri(), ri()
		graph[a] = append(graph[a], edge{b, c})
	}

	s := 1
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}

	vis := make([]bool, n+1)
	cnt := make([]int, n+1)
	q := []int{}

	// init
	dist[s] = 0
	vis[s] = true
	q = append(q, s)

	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		vis[u] = false

		for _, ed := range graph[u] {
			v := ed.to
			w := ed.val
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				cnt[v] = cnt[u] + 1 // 记录边数
				if cnt[v] >= n {
					fmt.Fprintln(out, "Yes")
					return
				}
				if !vis[v] {
					q = append(q, v)
					vis[v] = true
				}
			}
		}
	}

	if dist[n] == 1<<31-1 {
		fmt.Fprintln(out, "impossible")
	} else {
		fmt.Fprintln(out, dist[n])
	}
}

type edge struct {
	to  int
	val int
}

func main() {
	_debug()
}
