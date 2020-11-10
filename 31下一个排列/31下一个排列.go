package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	nums = []int{1, 2, 3, 3} //[1 3 2 3]
	//nums = []int{3, 2, 1}
	//nums = []int{3, 2, 1, 1}  //[1 1 2 3]
	nums = []int{0, 3, 2, 1} //[1 0 2 3]
	//nums = []int{0, 3, 2, 1, 1} //[1 0 1 2 3]
	//nums = []int{2, 3, 1}       //[3 1 2]
	//nums = []int{2, 3, 1, 1}    //[3 1 1 2]
	//nums = []int{2, 4, 3, 1}    //[3 1 2 4]
	nums = []int{1, 3, 2, 1} //[2 1 1 3]
	//nums = []int{1, 3, 2}
	//nums = []int{2, 1, 3}
	//nums = []int{4, 2, 4, 4, 3} //[4 3 2 4 4]

	//nextPermutation(nums)
	nextPermutationOfficial(nums)
	fmt.Println(nums)
}

//4 ms-69.02%	2.5 MB-46.90%
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	if n == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	const init = -2

	var decreaseStart, increseStart int = init, 0 //逆序开始和结束的下标
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			increseStart = i - 1
			break
		} else if nums[i] <= nums[i-1] {
			//开始逆序
			decreaseStart = i - 1
		}
	}

	if decreaseStart == init { //没有递减，直接把非递减的两个交换就行
		nums[increseStart], nums[increseStart+1] = nums[increseStart+1], nums[increseStart]
		return
	}

	if decreaseStart == 0 { //没有递增，直接反转
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
		return
	}

	var index int
	//找到比左边的数大的最小右边数
	for i := decreaseStart; i < n; i++ {
		if nums[i] > nums[decreaseStart-1] {
			index = i
		} else {
			break
		}
	}

	nums[decreaseStart-1], nums[index] = nums[index], nums[decreaseStart-1]

	// 剩下的反转
	for i := decreaseStart; i < (n+decreaseStart)/2; i++ {
		nums[i], nums[n-1-(i-decreaseStart)] = nums[n-1-(i-decreaseStart)], nums[i]
	}
	return
}

//4 ms-69.02%	2.5 MB-66.47%
func nextPermutationOfficial(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//0 ms-100.00%	2.5 MB-16.67%
func nextPermutation1(nums []int) {
	if len(nums) < 2 {
		return
	}
	p, q := -1, 0
	for tp, tq := 0, 1; tq < len(nums); {
		if nums[tp] < nums[tq] {
			p, q = tp, tq
		} else if nums[tp] > nums[tq] {
			if p >= 0 && nums[tq] > nums[p] && nums[tq] < nums[q] {
				q = tq
			}
		}
		tp++
		tq++
	}
	if p >= 0 {
		nums[p], nums[q] = nums[q], nums[p]
	}
	sort.Ints(nums[p+1:])
}
