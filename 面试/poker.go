package main

import "fmt"

func main() {
	//	我手中有一堆扑克牌， 但是观众不知道它的顺序。
	//
	//第一步， 我从牌顶拿出一张牌， 放到桌子上。
	//
	//第二步， 我从牌顶再拿一张牌， 放在手上牌的底部。
	//
	//第三步， 重复第一/二步的操作， 直到我手中所有的牌都放到了桌子上。
	//
	//最后， 观众可以看到桌子上牌的顺序是：13\12\11\10\9\8\7\6\5\4\3\2\1 请问， 我刚开始拿在手里的牌的顺序是什么？用代码实现以下
	nums := []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(poker(nums))
}

func poker(nums []int) []int {
	var ans []int
	for len(nums) > 0 {
		if len(ans) > 0 {
			top := ans[len(ans)-1]
			ans = append([]int{top}, ans[:len(ans)-1]...)
		}
		top := nums[0]
		nums = nums[1:]
		ans = append([]int{top}, ans...)
	}
	return ans
}
