package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("apple", 3)
	obj.Insert("app", 2)
	obj.Insert("app", 1)
	fmt.Println(obj.Sum("ap"))
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

type MapSum struct {
	root *TrieNode
	cnt  map[string]int
}

type TrieNode struct {
	children [26]*TrieNode
	val      int
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if this.cnt[key] > 0 {
		delta -= this.cnt[key]
	}
	this.cnt[key] = val
	node := this.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
		node.val += delta
	}
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}
