package main

import "fmt"

func main() {
	s := ")()())())()" //4
	//s = "()(()"        //2
	//s = "()(())"       //6
	s = "(()())" //6
	//s = "(()(((()" //2
	//s = "(()(((()((" //2

	//fmt.Println(longestValidParentheses(s))
	//fmt.Println(longestValidParenthesesStackOfficial(s))
	fmt.Println(longestValidParenthesesDPOfficial(s))
}

//我自己写的栈，有问题
func longestValidParentheses(s string) int {
	stack := []byte{}
	var curr, ans int
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			stack = append(stack, '(')
			i++
			continue
		}
		if len(stack) == 0 {
			curr = 0
			i++
			continue
		}
		temp := 0
		var index int
		for j := i; j < len(s); j++ {
			index = j
			if len(stack) > 0 {
				if s[j] == ')' {
					temp += 2
					stack = stack[:len(stack)-1]
				} else {
					if len(stack) >= len(s)-j {
						if temp > ans {
							ans = temp
						}
						break
					}
					stack = append(stack, '(')
				}
			} else {
				break
			}
		}
		//todo
		if len(stack) > 0 {
			stack = nil
			curr = temp
		} else {
			curr += temp
		}
		if curr > ans {
			ans = curr
		}
		i = index
	}
	return ans
}

//官方的栈，0 ms-100.00%	3.5 MB-26.26%
func longestValidParenthesesStackOfficial(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestValidParenthesesDPOfficial(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}
