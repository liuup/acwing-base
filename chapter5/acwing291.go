package main

import (
	"bufio"
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

	// 手写输出，适用于有大量（~1e6）输出的场景，CF 上可以再快 60~90ms
	// 使用前 https://codeforces.com/contest/1208/submission/176961129
	// 使用后 https://codeforces.com/contest/1208/submission/176963572
	outS := make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmpS := [32]byte{}              // 可根据绝对值的十进制长度的上限调整
	wInt := func(x int64) {
		if x == 0 { // 如果保证不为零则去掉
			outS = append(outS, '0')
			return
		}
		if x < 0 { // 如果保证是非负数则去掉
			x = -x
			outS = append(outS, '-')
		}
		p := len(tmpS)
		for ; x > 0; x /= 10 {
			p--
			tmpS[p] = '0' | byte(x%10)
		}
		outS = append(outS, tmpS[p:]...)
	}

	for {
		n, m := ri(), ri()
		if n|m == 0 {
			break
		}

		for i, v := range f {
			for j := range v {
				f[i][j] = 0
			}
		}

		for i := 0; i < (1 << n); i++ {
			st[i] = true
			cnt := 0
			for j := 0; j < n; j++ {
				if i>>j&1 == 1 {
					if cnt&1 == 1 {
						st[i] = false
						// break
					}
					cnt = 0
				} else {
					cnt++
				}
			}
			if cnt&1 == 1 {
				st[i] = false
			}
		}

		f[0][0] = 1

		for i := 1; i <= m; i++ { // 遍历每一列
			for j := 0; j < 1<<n; j++ { // 遍历每一列的每一种状态
				for k := 0; k < 1<<n; k++ { // 枚举i-1列每一种状态
					// f[i-1][k]成功转移到发f[i][j]
					if (j&k == 0) && st[j|k] {
						f[i][j] += f[i-1][k]
					}
				}
			}
		}
		wInt(f[m][0])
		outS = append(outS, '\n')
	}
	os.Stdout.Write(outS)
}

const N = 12
const M = 1 << N

var (
	f  [N][M]int64
	st [M]bool // 判断某个状态是否合法
)

func main() {
	_debug()
}
