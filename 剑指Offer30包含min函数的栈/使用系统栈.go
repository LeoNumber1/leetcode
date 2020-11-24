package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor1()
	nums := []int{-2, 0, -3}
	for _, num := range nums {
		obj.Push(num)
	}
	obj.Pop()
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 := obj.Min()
	fmt.Println(param_4)
}

type MinStack1 struct {
	A list.List
	B list.List
}

/** initialize your data structure here. */
func Constructor1() MinStack1 {
	return MinStack1{}
}

func (this *MinStack1) Push(x int) {
	this.A.PushBack(x)
	if this.B.Len() == 0 || this.B.Back().Value.(int) >= x {
		this.B.PushBack(x)
	}
}

func (this *MinStack1) Pop() {
	if this.A.Len() > 0 {
		a := this.A.Back()
		this.A.Remove(a)
		if this.B.Len() > 0 && this.B.Back().Value == a.Value {
			this.B.Remove(this.B.Back())
		}
	}
}

func (this *MinStack1) Top() int {
	if this.A.Len() == 0 {
		return 0
	}
	return this.A.Back().Value.(int)
}

func (this *MinStack1) Min() int {
	if this.B.Len() == 0 {
		return 0
	}
	return this.B.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
