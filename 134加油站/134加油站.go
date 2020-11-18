package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	//gas = []int{2, 3, 4}
	//cost = []int{3, 4, 3}

	//gas = []int{5, 1, 2, 3, 4}
	//cost = []int{4, 4, 1, 5, 1}

	fmt.Println(canCompleteCircuit0(gas, cost))
	fmt.Println(canCompleteCircuit(gas, cost))
}

//执行耗时:176 ms,击败了24.26% 的Go用户
//内存消耗:4 MB,击败了5.65% 的Go用户
func canCompleteCircuit0(gas []int, cost []int) int {
	n := len(gas)
	var gasTotal, costTotal int
	for k, v := range gas {
		gasTotal += v
		costTotal += cost[k]
	}
	if costTotal > gasTotal {
		return -1
	}
	var start = 0
	var dfs func(index, total int, first bool) int
	dfs = func(index, total int, first bool) int {
		if !first && index == start {
			return index
		}
		if cost[index] > total {
			return -1
		}
		if index < n-1 {
			total += gas[index+1] - cost[index]
			return dfs(index+1, total, false)
		} else {
			total += gas[0] - cost[index]
			return dfs(0, total, false)
		}
	}
	for i := 0; i < n; i++ {
		start = i
		ans := dfs(start, gas[start], true)
		if ans != -1 {
			return ans
		}
	}
	return -1
}

//执行耗时:20 ms,击败了32.34% 的Go用户
//内存消耗:2.9 MB,击败了100.00% 的Go用户
func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	var gasTotal, costTotal int
	index := -1
	for k, v := range gas {
		if index == -1 && v >= cost[k] {
			index = k
		}
		gasTotal += v
		costTotal += cost[k]
	}
	if costTotal > gasTotal {
		return -1
	}
for1:
	for i := index; i < n; i++ {
		var total int
		for j := i; j < n; j++ {
			total += gas[j] - cost[j]
			if total < 0 {
				continue for1
			}
		}
		for w := 0; w <= i; w++ {
			total += gas[w] - cost[w]
			if total < 0 {
				continue for1
			}
		}
		return i
	}
	return -1
}

//执行耗时:20 ms,击败了32.34% 的Go用户
//内存消耗:3.4 MB,击败了5.65% 的Go用户
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	var gasTotal, costTotal int
	list := []int{}
	for k, v := range gas {
		if v >= cost[k] {
			list = append(list, k)
		}
		gasTotal += v
		costTotal += cost[k]
	}
	if costTotal > gasTotal {
		return -1
	}
for1:
	for _, i := range list {
		var total int
		for j := i; j < n; j++ {
			total += gas[j] - cost[j]
			if total < 0 {
				continue for1
			}
		}
		for w := 0; w <= i; w++ {
			total += gas[w] - cost[w]
			if total < 0 {
				continue for1
			}
		}
		return i
	}
	return -1
}

//执行耗时:4 ms,击败了92.34% 的Go用户
//内存消耗:2.9 MB,击败了84.35% 的Go用户
func canCompleteCircuitOfficial(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
