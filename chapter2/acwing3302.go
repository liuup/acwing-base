package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func _debug() {
	var s string
	fmt.Fscan(in, &s)

	fmt.Fprintln(out, backCalc(mid2back(s)))
}

// 中缀转后缀
func mid2back(s string) []string {
	priority := func(x byte) int {
		if x == '*' || x == '/' {
			return 2
		} else if x == '+' || x == '-' {
			return 1
		}
		return 0
	}

	stk := []byte{}
	ans := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' { // 数字直接输出 处理多位数字
			var j int
			for j = i; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			ans = append(ans, s[i:j])
			i = j - 1
		} else if s[i] == '(' {
			stk = append(stk, s[i])
		} else if s[i] == ')' {
			for len(stk) > 0 && stk[len(stk)-1] != '(' {
				ans = append(ans, string(stk[len(stk)-1]))
				stk = stk[:len(stk)-1]
			}
			stk = stk[:len(stk)-1]
		} else { // 运算符 要判断优先级
			// 只要栈顶 符号的优先级不低于新符号，就不断取出栈顶并输出
			for len(stk) > 0 && priority(stk[len(stk)-1]) >= priority(s[i]) {
				ans = append(ans, string(stk[len(stk)-1]))
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, s[i])
		}
	}
	for len(stk) > 0 {
		ans = append(ans, string(stk[len(stk)-1]))
		stk = stk[:len(stk)-1]
	}
	return ans
}

// 后缀计算
func backCalc(s []string) (ans int) {
	op := func(b, a int, x string) int {
		if x == "+" {
			return b + a
		} else if x == "-" {
			return b - a
		} else if x == "*" {
			return b * a
		} else if x == "/" {
			return b / a
		}
		return -1 << 31
	}

	stk := []int{}
	for _, x := range s {
		if x == "+" || x == "-" || x == "*" || x == "/" {
			a := stk[len(stk)-1]
			b := stk[len(stk)-2]
			stk = stk[:len(stk)-2]
			stk = append(stk, op(b, a, x))
		} else {
			ii, _ := strconv.Atoi(x)
			stk = append(stk, ii)
		}
	}
	return stk[0] // 最后肯定只剩一个
}

// ===== ===== fast io ===== =====
// golang fast io from 0x3F https://github.com/EndlessCheng/codeforces-go/

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func main() {
	in = bufio.NewReader(os.Stdin) // 搭配Fscan使用
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== ===== fast io ===== =====
	_debug()
}
