package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	trie := InitTrie()

	var n int
	fmt.Fscan(in, &n)

	var op string
	var s string
	for ; n > 0; n-- {
		fmt.Fscan(in, &op)
		if op == "I" {
			fmt.Fscan(in, &s)
			trie.Insert(s)
		} else if op == "Q" {
			fmt.Fscan(in, &s)
			_, cnt := trie.Search(s)
			fmt.Fprintln(out, cnt)
		}
	}
}

// 前缀树
type Trie struct {
	isEnd bool
	next  [26]*Trie
	cnt   int // 以这个字母为结尾的单词存在多少次
}

func InitTrie() Trie {
	return Trie{
		isEnd: false,
		next:  [26]*Trie{},
		cnt:   0,
	}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		if node.next[word[i]-'a'] == nil {
			node.next[word[i]-'a'] = &Trie{}
		}
		node = node.next[word[i]-'a']
	}
	node.isEnd = true
	node.cnt++
}

func (t *Trie) Search(word string) (bool, int) {
	node := t
	for i := 0; i < len(word); i++ {
		node = node.next[word[i]-'a']
		if node == nil {
			return false, 0
		}
	}

	if node.isEnd {
		return true, node.cnt
	}
	return false, 0
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for i := 0; i < len(prefix); i++ {
		node = node.next[prefix[i]-'a']
		if node == nil {
			return false
		}
	}
	return true
}

func (t *Trie) Empty() bool {
	for _, son := range t.next {
		if son != nil {
			return false
		}
	}
	return true
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
