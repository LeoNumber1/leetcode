package main

import "fmt"

func main() {
	a := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(a))
}

//执行耗时:28 ms,击败了54.40% 的Go用户
//内存消耗:6.8 MB,击败了5.52% 的Go用户
func sortArrayByParityII0(a []int) []int {
	queue0 := []int{}
	queue1 := []int{}
	for i := 0; i < len(a); i++ {
		if a[i]&1 == 0 { //偶数
			if i&1 != 0 {
				if len(queue1) > 0 {
					j := queue1[0]
					queue1 = queue1[1:]
					a[i], a[j] = a[j], a[i]
				} else {
					queue0 = append(queue0, i)
				}
			}
		} else { //奇数
			if i&1 == 0 {
				if len(queue0) > 0 {
					j := queue0[0]
					queue0 = queue0[1:]
					a[i], a[j] = a[j], a[i]
				} else {
					queue1 = append(queue1, i)
				}
			}
		}
	}
	return a
}

//执行耗时:28 ms,击败了54.40% 的Go用户
//内存消耗:6.3 MB,击败了53.04% 的Go用户
func sortArrayByParityII(a []int) []int {
	for i, j := 0, 1; i < len(a) && j < len(a); {
		x := a[i] & 1
		if x != 0 && a[j]&1 == 0 {
			a[i], a[j] = a[j], a[i]
			i += 2
			j += 2
			continue
		}
		if x == 0 {
			i += 2
			continue
		}
		j += 2
	}
	return a
}

//执行耗时:24 ms,击败了87.36% 的Go用户
//内存消耗:6.3 MB,击败了53.04% 的Go用户
func sortArrayByParityIIOfficial(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}
