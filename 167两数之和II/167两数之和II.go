package main

import "fmt"

func main() {
	numbers := []int{
		//2, 5, 7, 9, 12, 13, 15,
		//2, 3, 4,
		1, 2, 3, 4, 4, 9, 56, 90,
	}
	//target := 19
	//target := 6
	target := 8
	fmt.Println(twoSum(numbers, target))
	fmt.Println(twoSum2(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n < 2 {
		return []int{0, 0}
	}

	var count int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count++
			if numbers[i]+numbers[j] == target {
				fmt.Println(count)
				return []int{i + 1, j + 1}
			} else if numbers[i]+numbers[j] > target {
				break
			}
		}
	}
	return []int{0, 0}
}

func twoSum2(numbers []int, target int) []int {
	n := len(numbers)
	if n < 2 {
		return []int{0, 0}
	}

	var count int

	mid := n / 2
	if target > 2*numbers[mid] {
		for i := mid; i < n; i++ {
			for j := mid; j >= 0; j-- {
				count++
				if numbers[i]+numbers[j] == target {
					fmt.Println(count)
					return []int{j + 1, i + 1}
				} else if numbers[i]+numbers[j] < target {
					break
				}
			}
		}
	} else {
		for i := 0; i <= mid; i++ {
			for j := i + 1; j < n; j++ {
				count++
				if numbers[i]+numbers[j] == target {
					fmt.Println(count)
					return []int{i + 1, j + 1}
				} else if numbers[i]+numbers[j] > target {
					break
				}
			}
		}
	}

	return []int{0, 0}
}
