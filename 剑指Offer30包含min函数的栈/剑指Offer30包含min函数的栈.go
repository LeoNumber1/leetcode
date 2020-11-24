package main

import "fmt"

func main() {
	obj := Constructor()
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

type MinStack struct {
	A []int
	B []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.A = append(this.A, x)
	if len(this.B) == 0 || this.B[len(this.B)-1] >= x {
		this.B = append(this.B, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.A) > 0 {
		v := this.A[len(this.A)-1]
		this.A = this.A[:len(this.A)-1]
		if len(this.B) > 0 && this.B[len(this.B)-1] == v {
			this.B = this.B[:len(this.B)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.A) == 0 {
		return 0
	}
	return this.A[len(this.A)-1]
}

func (this *MinStack) Min() int {
	if len(this.B) == 0 {
		return 0
	}
	return this.B[len(this.B)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
