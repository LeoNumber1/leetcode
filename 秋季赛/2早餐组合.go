package main

import (
	"fmt"
	"sort"
)

func main() {
	staple := []int{10, 20, 5}
	drinks := []int{5, 5, 2}
	x := 15

	//staple = []int{2, 1, 1}
	//drinks = []int{8, 9, 5, 1}
	//x = 9

	//fmt.Println(breakfastNumber(staple, drinks, x))
	//fmt.Println(breakfastNumber1(staple, drinks, x))
	fmt.Println(breakfastNumber2(staple, drinks, x))

	//fmt.Println(NUMBER)
}

const NUMBER = 1000000007

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
