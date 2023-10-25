package main

import (
	"bufio" // io
	"fmt"   // io
	"math/bits"
	"os"      // io
	"strconv" // io
)

var (
	in  *bufio.Scanner
	out *bufio.Writer
)

func ri() int { // fast read int
	in.Scan()
	x, _ := strconv.Atoi(string(in.Bytes()))
	return x
}

func rf() float64 { // fast read float64
	in.Scan()
	f, _ := strconv.ParseFloat(string(in.Bytes()), 64)
	return f
}

func main() {
	// ===== fast io =====
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // 分割
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// ===== fast io =====

	n := ri()
	for i := 0; i < n; i++ {
		tmp := ri()
		fmt.Fprint(out, bits.OnesCount(uint(tmp)), " ")
	}
	// 或者不用库函数
	// 	n := ri()
	// 	for i := 0; i < n; i++ {
	// 		tmp := ri()
	// 		cnt := 0
	// 		for tmp > 0 {
	// 			if tmp&1 == 1 {
	// 				cnt++
	// 			}
	// 			tmp >>= 1
	// 		}
	// 		fmt.Fprint(out, cnt, " ")
	// 	}
}
