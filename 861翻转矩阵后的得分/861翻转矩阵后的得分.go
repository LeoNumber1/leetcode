package main

import (
	"fmt"
)

func main() {
	A := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	fmt.Println(matrixScore(A))
	//fmt.Println(matrixScoreOfficial(A))
}

func matrixScore(A [][]int) int {
	m := len(A)
	n := len(A[0])
	for _, nums := range A {
		if nums[0] == 0 {
			for k, num := range nums {
				if num == 0 {
					nums[k] = 1
				} else {
					nums[k] = 0
				}
			}
		}
	}
	for j := 1; j < n; j++ {
		countZero := 0
		for i := 0; i < m; i++ {
			if A[i][j] == 0 {
				countZero++
			}
		}
		if countZero > m/2 {
			for i := 0; i < m; i++ {
				if A[i][j] == 1 {
					A[i][j] = 0
				} else {
					A[i][j] = 1
				}
			}
		}
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				//ans += int(math.Exp2(float64(n - 1 - j)))
				ans += 1 << (n - 1 - j)
			}
		}
	}
	return ans
}

func matrixScoreOfficial(a [][]int) int {
	m, n := len(a), len(a[0])
	ans := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		ones := 0
		for _, row := range a {
			if row[j] == row[0] {
				ones++
			}
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}
