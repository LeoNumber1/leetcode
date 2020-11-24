package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	popped = []int{4, 3, 5, 1, 2}
	//fmt.Println(validateStackSequences(pushed, popped))
	fmt.Println(validateStackSequences1(pushed, popped))
}

//8 ms-82.19%	3.8 MB-40.82%
func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n == 0 {
		return true
	}
	stack := []int{}
	index := 0
	i := 0
	for i < n {
		if len(stack) == 0 {
			stack = append(stack, pushed[i])
			i++
		} else {
			if popped[index] != stack[len(stack)-1] {
				stack = append(stack, pushed[i])
				i++
			} else {
				stack = stack[:len(stack)-1]
				index++
			}
		}
	}
	for index != n && popped[index] == stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		index++
	}
	return len(stack) == 0
}

func validateStackSequences1(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0
	// 模拟进栈
	for i := 0; i < len(pushed); i++ {
		// 进栈
		stack = append(stack, pushed[i])
		// 出栈
		for popped[index] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			index++
			// 防止slice越界
			if len(stack) == 0 {
				break
			}
		}
	}
	return len(popped) == index
}
