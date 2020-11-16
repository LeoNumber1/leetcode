package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	//people = [][]int{{3, 0}, {8, 0}, {4, 1}, {7, 1}, {5, 2}, {4, 4}}
	//people = [][]int{{8, 0}, {4, 4}, {7, 1}, {3, 0}, {4, 1}, {5, 2}}

	//fmt.Println(reconstructQueue(people))
	fmt.Println(reconstructQueueOfficial1(people))
}

//执行耗时:44 ms,击败了9.05% 的Go用户
//内存消耗:6.2 MB,击败了34.45% 的Go用户
func reconstructQueue0(people [][]int) [][]int {
	n := len(people)
	if n <= 1 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i][1] < people[j][1] || (people[i][1] == people[j][1] && people[i][0] < people[j][0])
	})

	isOk := make([]bool, n)
	isOk[0] = true

	for i := 1; i < n; i++ {
		if !isOk[i] {
			var num int
			index := i
			for before := index - 1; before >= 0; {
				if num > people[index][1] {
					people[index-1], people[index] = people[index], people[index-1]
					isOk[index-1], isOk[index] = isOk[index], isOk[index-1]
					index--
					num--
					before--
					continue
				}
				if people[before][0] >= people[index][0] {
					num++
					if num > people[index][1] {
						continue
					}
				}
				before--
			}
			isOk[index] = true
		}
	}

	return people
}

//执行耗时:36 ms,击败了14.29% 的Go用户
//内存消耗:6.2 MB,击败了34.45% 的Go用户
func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	if n <= 1 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i][1] < people[j][1] || (people[i][1] == people[j][1] && people[i][0] < people[j][0])
	})

	isOk := make([]bool, n)
	isOk[0] = true

	for i := 1; i < n; i++ {
		if !isOk[i] {
			var num int
			index := i
			for before := index - 1; before >= 0; {
				if num > people[index][1] {
					people[index-1], people[index] = people[index], people[index-1]
					isOk[index-1], isOk[index] = isOk[index], isOk[index-1]
					index--
					num--
					before--
					continue
				}
				if people[before][0] >= people[index][0] {
					num++
					if num > people[index][1] {
						continue
					}
				} else {
					a := num + people[before][1]
					if a == people[index][1] {
						break
					}
				}
				before--
			}
			isOk[index] = true
		}
	}

	return people
}

//执行耗时:24 ms,击败了28.57% 的Go用户
//内存消耗:6.2 MB,击败了40.67% 的Go用户
func reconstructQueueOfficial(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}

//执行耗时:36 ms,击败了14.29% 的Go用户
//内存消耗:7.7 MB,击败了7.18% 的Go用户
func reconstructQueueOfficial1(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}
