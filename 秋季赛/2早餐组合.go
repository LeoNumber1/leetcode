package main

import (
	"fmt"
	"sort"
)

func main() {
	staple := []int{10, 20, 5}
	drinks := []int{5, 5, 2}
	x := 15

	staple = []int{2, 1, 1}
	drinks = []int{8, 9, 5, 1}
	x = 9

	//fmt.Println(breakfastNumber(staple, drinks, x))
	//fmt.Println(breakfastNumber1(staple, drinks, x))
	fmt.Println(breakfastNumber2(staple, drinks, x))
	fmt.Println(breakfastNumber3(staple, drinks, x))
	fmt.Println(breakfastNumber4(staple, drinks, x))

	//fmt.Println(NUMBER)
}

const NUMBER = 1000000007

//超时了
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)

	var count int
for1:
	for i := 0; i < len(staple); i++ {
		for j := 0; j < len(drinks); j++ {
			if staple[i]+drinks[j] <= x {
				count++
			} else {
				if staple[i] > x {
					break for1
				}
				if drinks[j] > x {
					break
				}
			}
		}
	}

	return count % NUMBER
}

//超时了
func breakfastNumber1(staple []int, drinks []int, x int) int {
	stapleNew := [100001]int{}
	drinksNew := [100001]int{}
	n := 0
	ns := len(staple)
	nd := len(drinks)
	if ns > nd {
		n = ns
	} else {
		n = nd
	}
	for i := 0; i < n; i++ {
		if i < ns && staple[i] < x {
			stapleNew[staple[i]]++
		}
		if i < nd && drinks[i] < x {
			drinksNew[drinks[i]]++
		}
	}
	var count int

	for i := 0; i < 100001; i++ {
		if i >= x {
			break
		}
		if stapleNew[i] != 0 {
			for j := 0; j < 100001; j++ {
				if j >= x {
					break
				}
				if drinksNew[j] != 0 && i+j <= x {
					count += stapleNew[i] * drinksNew[j]
				}
			}
		}
	}

	return count % NUMBER
}

//超时了
func breakfastNumber2(staple []int, drinks []int, x int) int {
	n := 0
	ns := len(staple)
	nd := len(drinks)
	if ns > nd {
		n = ns
	} else {
		n = nd
	}

	stapleMap := map[int]int{}
	drinksMap := map[int]int{}

	for i := 0; i < n; i++ {
		if i < ns && staple[i] < x {
			stapleMap[staple[i]]++
		}
		if i < nd && drinks[i] < x {
			drinksMap[drinks[i]]++
		}
	}

	var count int

	for ks, vs := range stapleMap {
		for kd, vd := range drinksMap {
			if ks+kd <= x {
				count += vs * vd
			}
		}
	}
	return count % NUMBER
}

func breakfastNumber3(staple []int, drinks []int, x int) int {
	sort.Ints(drinks)
	var count int
	for i := 0; i < len(staple); i++ {
		if staple[i] >= x {
			continue
		}
		start, end := 0, len(drinks)-1
		for start <= end {
			middle := (start + end) / 2
			if staple[i]+drinks[middle] <= x {
				count += middle + 1 - start
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}

	return count % NUMBER
}

func breakfastNumber4(staple []int, drinks []int, x int) int {
	const NUMBER = 1000000007
	var count = 0
	sort.Ints(staple)
	sort.Ints(drinks)
	for i := 0; i < len(staple); i++ {
		if staple[i] > x {
			break
		}
		var low = 0
		var high = len(drinks)
		var target = x - staple[i]
		//采用二分查找
		for j := len(drinks) - 1; j >= 0 && low < high; j-- {
			var mid = (low + high) / 2
			//比目标值大，说明目标值在当前mid值左边，接着往左分
			if drinks[mid] > target {
				high = mid
			} else {
				//比目标值小，往右+1查找
				low = mid + 1
			}
		}
		//找到小于等于目标值(target)的元素下标，即[0，low]中的值都满足。
		count += low
	}
	return count % NUMBER
}
