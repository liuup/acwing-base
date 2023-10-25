package main

import (
	"bufio"
	"container/heap"
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
	_ = []interface{}{rc, ri}

	n := ri()

	pq := InitHeap(Minheap)
	for i := 0; i < n; i++ {
		pq.push_(pair{ri()})
	}

	ans := 0
	for pq.Len() > 1 {
		a := pq.pop_()
		b := pq.pop_()

		pq.push_(pair{a.val + b.val})
		ans += a.val + b.val
	}

	fmt.Fprintln(out, ans)
}

// 优先队列/堆
// h := InitHeap(Maxheap)
// h.push_(pair{1})
// p := h.pop_()

const Maxheap = true
const Minheap = false

type pair struct {
	// id       int
	// distance int
	val int
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
		return h.vals[i].val > h.vals[j].val
	}

	return h.vals[i].val < h.vals[j].val
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

// func (h *hp) Contain(v int) bool {	// 用的不多
// 	for i := 0; i < h.Len(); i++ {
// 		if (*h).vals[i].x == v {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	_debug()
}
