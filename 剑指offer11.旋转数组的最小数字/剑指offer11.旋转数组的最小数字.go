package main

import "fmt"

func main() {
	numbers := []int{
		3, 4, 5, 1, 2,
	}

	fmt.Println(minArrayOfficial(numbers))
}

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min int = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			return numbers[i]
		}
	}
	return min
}

func minArrayOfficial(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	left := 0
	right := len(numbers) - 1

	for right > left {
		middle := left + (right-left)/2
		if numbers[middle] < numbers[right] {
			right = middle
		} else if numbers[middle] > numbers[right] {
			left = middle + 1
		} else {
			right--
		}
	}
	return numbers[left]
}
