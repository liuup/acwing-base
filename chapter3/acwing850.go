package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func _debug() {
	n, m := ri(), ri()

	graph := make([][]edge, n+1)
	for i := range graph {
		graph[i] = make([]edge, 0)
	}

	// 建图
	for ; m > 0; m-- {
		x, y, z := ri(), ri(), ri()
		graph[x] = append(graph[x], edge{y, z})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}

	dist[1] = 0 // 因为是从1开始

	hp := InitHeap(Minheap)
	hp.push_(pair{1, 0})

	for hp.Len() != 0 {
		cur := hp.pop_()

		if dist[cur.id] < cur.distance {
			continue
		}

		// 所有孩子
		for _, e := range graph[cur.id] {
			// 计算距离
			d := dist[cur.id] + e.val
			if d < dist[e.to] {
				dist[e.to] = d
				hp.push_(pair{e.to, d})
			}
		}
	}

	if dist[n] == 1<<31-1 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, dist[n])
	}
}

type edge struct {
	to  int
	val int
}

const Maxheap = true
const Minheap = false

type pair struct {
	id       int
	distance int
}

// type hp []pair
type hp struct {
	vals []pair
	mode bool // true maxheap; false minheap
}

func InitHeap(mode bool) hp { // init heap
	h := hp{mode: mode}
	heap.Init(&h)
	return h
}

func (h hp) Less(i, j int) bool {
	if h.mode { // true maxheap; 修改对应的if分支
		return h.vals[i].distance > h.vals[j].distance
	}

	return h.vals[i].distance < h.vals[j].distance
}

func (h hp) Len() int            { return len(h.vals) }
func (h hp) Swap(i, j int)       { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }
func (h *hp) Push(v interface{}) { h.vals = append(h.vals, v.(pair)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a.vals[len(a.vals)-1]
	*h = hp{a.vals[:len(a.vals)-1], a.mode}
	return v
}
func (h *hp) push_(v pair) { heap.Push(h, v) } // 两个自定义push pop
func (h *hp) pop_() pair   { return heap.Pop(h).(pair) }
func (h *hp) Peek() pair   { return (h.vals)[0] } // 有越界风险

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/
const eof = 0

var (
	_i, _n, buf = 0, 0, make([]byte, 1<<12) // 4KB rc rs rsn

	outs = make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmps = [64]byte{}              // 可根据绝对值的十进制长度的上限调整
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

var (
	in  *bufio.Reader // 搭配Fscan使用
	out *bufio.Writer
)

func main() {
	in = bufio.NewReader(os.Stdin) // 搭配Fscan使用
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
