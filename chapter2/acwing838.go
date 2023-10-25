package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func _debug() {
	hp := InitHeap(Minheap)

	n, m := ri(), ri()
	for i := 0; i < n; i++ {
		hp.push_(pair{ri()})
	}

	for i := 0; i < m; i++ {
		fmt.Fprint(out, hp.pop_().x, " ")
	}
}

// 优先队列/堆
// h := InitHeap(Maxheap)
// h.push_(pair{1})
// p := h.pop_()

const Maxheap = true
const Minheap = false

type pair struct {
	x int
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
		// return h.vals[i].val > h.vals[j].val
		return true
	}

	return h.vals[i].x < h.vals[j].x
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
	out *bufio.Writer

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

// 读一个仅包含小写字母的字符串，必要时进行修改
func rs() (s []byte) {
	b := rc()
	for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
	}
	for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		s = append(s, b)
	}
	return
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

func main() {
	// ===== ===== fast io ===== =====
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
