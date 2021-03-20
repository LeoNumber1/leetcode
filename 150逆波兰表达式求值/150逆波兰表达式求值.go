package main

import (
	"fmt"
	"strconv"
)

func main() {

	type args struct {
		tokens []string
	}

	tests := []struct {
		index  int
		args   args
		target int
	}{
		{1, args{[]string{"2", "1", "+", "3", "*"}}, 9},
		{2, args{[]string{"4", "13", "5", "/", "+"}}, 6},
		{3, args{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, 22},
	}

	var errNum bool
	for _, test := range tests {
		result := evalRPN(test.args.tokens)
		if test.target != result {
			errNum = true
			fmt.Println("—————— err in index:", test.index, "except:", test.target, " result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests are OK! -------")
	}
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err != nil {
			calculate(&stack, token)
			continue
		}
		inStack(&stack, v)
	}
	return outStack(&stack)
}

func calculate(stack *[]int, token string) {
	b, a := outStack(stack), outStack(stack)
	var val int
	switch token {
	case "+":
		val = a + b
	case "-":
		val = a - b
	case "*":
		val = a * b
	case "/":
		val = a / b
	}
	inStack(stack, val)
}

func inStack(stack *[]int, val int) {
	*stack = append(*stack, val)
}

func outStack(stack *[]int) int {
	n := len(*stack)
	v := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	return v
}

func evalRPNOfficial(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
