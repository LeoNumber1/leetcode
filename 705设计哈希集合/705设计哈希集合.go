package main

import (
	"container/list"
	"fmt"
)

func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	fmt.Println(hs.Contains(1))

}

const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) hash(key int) int {
	return key % base
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
