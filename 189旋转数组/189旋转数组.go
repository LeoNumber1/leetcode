package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	//rotate(nums, k)
	rotateOfficial(nums, k)

	fmt.Println(nums)
}

//4 ms	3.2 MB
func rotate0(nums []int, k int) {
	n := len(nums)
	if n < 2 || k == 0 {
		return
	}
	k %= n
	temp := make([]int, n-k)
	copy(temp, nums[:n-k])
	for i := n - k; i < n; i++ {
		nums[(i+k)%n] = nums[i]
	}
	for i := k; i < n; i++ {
		nums[i] = temp[i-k]
	}
}

//104 ms	3.2 MB
func rotate1(nums []int, k int) {
	n := len(nums)
	if n < 2 || k == 0 {
		return
	}
	k %= n
	for k > 0 {
		temp := nums[0]
		nums[0] = nums[n-1]
		for i := 1; i < n; i++ {
			pre := nums[i]
			nums[i] = temp
			temp = pre
		}
		k--
	}
}

//8 ms	3.1 MB
func rotate(nums []int, k int) {
	n := len(nums)
	if n < 2 || k == 0 || k == n {
		return
	}
	k %= n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

//4 ms	3.1 MB
func rotateOfficial(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
