package main

import "fmt"

func main() {
	A := []int{2, 1, 4, 7, 3, 2, 5}
	A = []int{2, 1, 4, 7, 3, 2, 2, 1, 2, 5, 8, 4, 2, 1}
	//A = []int{2, 1, 4, 7, 3, 2, 5, 8, 4, 2, 1}
	//A = []int{2, 3, 3, 2, 0, 2}
	//A = []int{875, 884, 239, 731, 723, 685}

	//fmt.Println(longestMountain(A))
	//fmt.Println(longestMountainOfficial(A))
	fmt.Println(longestMountainOfficial1(A))
}

//20 ms-78.79%	6.2 MB-50.00%
func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	var longer, longest, start int
	var hasMountain bool
	var status int = 0 //0为初始值，1为开始增，2为开始降

	for i := 1; i < len(A); i++ {
		if status == 0 {
			if A[i] > A[i-1] {
				status = 1
				if start == 0 {
					start = i - 1
				}
			} else {
				start = i
			}
		} else if status == 1 {
			if A[i] < A[i-1] { //递增状态下开始下降，证明有山
				hasMountain = true
				longer = i - start + 1
				if longer > longest {
					longest = longer
				}
				status = 2
			} else if A[i] == A[i-1] {
				//递增的状态下持平
				status = 0
				start = 0
			}
		} else {
			if A[i] < A[i-1] { //递减的状态继续下降
				longer = i - start + 1
				if longer > longest {
					longest = longer
				}
			} else if A[i] > A[i-1] {
				//递减的状态下开始上升，状态回1，左山脚start为上一位
				status = 1
				start = i - 1
			} else {
				//递减的状态下持平，状态回0，左山脚start为自己
				status = 0
				start = i
			}
		}
	}

	if hasMountain && longest >= 3 {
		return longest
	} else {
		return 0
	}
}

//24 ms-45.45%	6.5 MB-13.33%
func longestMountainOfficial(a []int) (ans int) {
	n := len(a)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if a[i-1] < a[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if a[i+1] < a[i] {
			right[i] = right[i+1] + 1
		}
	}
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = l + r + 1
		}
	}
	return
}

//20 ms-78.79%	6.2 MB-50.00%
func longestMountainOfficial1(a []int) (ans int) {
	n := len(a)
	left := 0
	for left+2 < n {
		right := left + 1
		if a[left] < a[left+1] {
			for right+1 < n && a[right] < a[right+1] {
				right++
			}
			if right < n-1 && a[right] > a[right+1] {
				for right+1 < n && a[right] > a[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return
}
