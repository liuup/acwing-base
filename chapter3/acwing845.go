package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func _debug() {
	start := make([]byte, 9)
	for i := range start {
		start[i] = rsn(1)[0]
	}

	q := [][]byte{}
	dist := map[string]int{}

	q = append(q, start)
	dist[string(start)] = 0

	for len(q) != 0 {
		cur := q[0] // []byte
		q = q[1:]
		distance := dist[string(cur)]

		if string(cur) == "12345678x" {
			fmt.Fprintln(out, distance)
			return
		}

		idx := strings.Index(string(cur), "x")
		x := idx / 3
		y := idx % 3

		for _, d := range directions {
			dx := x + d[0]
			dy := y + d[1]

			if !_isok(dx, dy, 3, 3) {
				continue
			}
			idx_new := dx*3 + dy

			cur[idx], cur[idx_new] = cur[idx_new], cur[idx]
			if _, ok := dist[string(cur)]; !ok {
				dist[string(cur)] = distance + 1
				tmp := append([]byte{}, cur...)
				q = append(q, tmp) // copy and append
			}
			cur[idx], cur[idx_new] = cur[idx_new], cur[idx]
		}
	}
	fmt.Fprintln(out, -1)
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向

func _isok(i, j int, n, m int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

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
