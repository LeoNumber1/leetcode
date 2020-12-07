package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'Z', 'Z', 'Z', 'Z'}
	tasks = []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	tasks = []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2

	//fmt.Println(leastInterval(tasks, n))
	fmt.Println(leastIntervalOfficial1(tasks, n))
}

//3480 ms-5.45%	7.1 MB-5.45%
func leastInterval0(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	total := len(tasks)
	//map的key是cpu可以运行的状态，0为可以运行，1为下一次运行；二维数组里放了需要运行的程序集，其中每一维的0代表程序序号，1代表该程序剩余运行次数
	m := make(map[int][][2]int)
	byteArr := [26]int{}
	for _, task := range tasks {
		byteArr[task-'A']++
	}
	for k, v := range byteArr {
		if v != 0 {
			m[0] = append(m[0], [2]int{k, v})
		}
	}
	mySort := func() {
		sort.Slice(m[0], func(i, j int) bool {
			return m[0][i][1] > m[0][j][1]
		})
	}
	mySort()
	var ans int

	for total > 0 {
		if len(m[0]) > 0 {
			task := m[0][0]
			m[0] = m[0][1:]
			task[1]--
			total--
			if task[1] != 0 {
				m[n+1] = append(m[n+1], task)
			}
		}
		for i := 1; i <= n+1; i++ {
			if v, has := m[i]; has {
				m[i-1] = append(m[i-1], v...)
				delete(m, i)
			}
		}
		mySort()
		ans++
	}

	return ans
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	total := len(tasks)
	//map的key是cpu可以运行的状态，0为可以运行，1为下一次运行；二维数组里放了需要运行的程序集，其中每一维的0代表程序序号，1代表该程序剩余运行次数
	m := make(map[int][]int)
	byteArr := [26]int{}
	for _, task := range tasks {
		byteArr[task-'A']++
	}
	for _, v := range byteArr {
		if v != 0 {
			m[0] = append(m[0], v)
		}
	}
	mySort := func() {
		sort.Slice(m[0], func(i, j int) bool {
			return m[0][i] > m[0][j]
		})
	}
	mySort()
	var ans int

	for total > 0 {
		if len(m[0]) > 0 {
			task := m[0][0]
			m[0] = m[0][1:]
			task--
			total--
			if task != 0 {
				m[n+1] = append(m[n+1], task)
			}
		}
		for i := 1; i <= n+1; i++ {
			if v, has := m[i]; has {
				m[i-1] = append(m[i-1], v...)
				delete(m, i)
			}
		}
		mySort()
		ans++
	}

	return ans
}

func leastIntervalOfficial(tasks []byte, n int) (minTime int) {
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}

	nextValid := make([]int, 0, len(cnt))
	rest := make([]int, 0, len(cnt))
	for _, c := range cnt {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}

	for range tasks {
		minTime++
		minNextValid := math.MaxInt64
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return
}

func leastIntervalOfficial1(tasks []byte, n int) int {
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}
