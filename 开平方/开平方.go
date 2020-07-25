package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	v := 1000000000
	//v = 10
	//v = 9
	t := 0.00001
	//t = 0.1
	//t = 0.21

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		t1 := time.Now()
		fmt.Println("暴力破解法计算的值 =", sqrt(v, t)) //暴力破解法
		fmt.Println("暴力破解法耗时：", time.Since(t1))
		wg.Done()
	}()

	go func() {
		t2 := time.Now()
		fmt.Println("二分法计算的值 =", sqrtDichotomy(v, t)) //二分法
		fmt.Println("二分法耗时：", time.Since(t2))
		wg.Done()
	}()

	go func() {
		t3 := time.Now()
		fmt.Println("实际的值 =", math.Sqrt(float64(v)))
		fmt.Println("实际计算耗时：", time.Since(t3))
		wg.Done()
	}()

	wg.Wait()
}

func sqrt(v int, t float64) float64 {
	return sqrt2(0, 1, float64(v), t)
}

func sqrt2(start float64, deepth float64, v float64, t float64) float64 {
	var num float64
	for i := start; i < v; i += deepth {
		if i*i <= v {
			num = i
			if deepth < t {
				return num
			}
		} else {
			break
		}
	}
	return sqrt2(num, deepth/10, v, t)
}

//二分法
func sqrtDichotomy(v int, t float64) float64 {
	vFloat := float64(v)
	var left, right float64 = 0, vFloat
	for left < right {
		mid := (left + right) / 2
		if mid*mid <= vFloat {
			left = mid
		} else {
			right = mid
		}
		if math.Abs(vFloat-mid*mid) <= t {
			return mid
		}
	}
	return 0
}
