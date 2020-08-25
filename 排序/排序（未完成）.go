package main

import "fmt"

func main() {
	arr := []int{
		3, 5, 4, 2, 1,
		//23, 46, 0, 8, 11, 18,
	}
	//fmt.Println(BubbleSort(arr))
	//fmt.Println(SelectionSort(arr))
	//fmt.Println(InsertSort(arr))
	fmt.Println(QuickSort(arr))
}

//冒泡排序
func BubbleSort(arr []int) []int {
	var lastIndex int = len(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < lastIndex; j++ {
			if j+1 < len(arr) && arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				if j+1 == lastIndex-1 {
					lastIndex--
					break
				}
			}
		}
	}
	//fmt.Println(lastIndex)
	return arr
}

//选择排序
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//插入排序
func InsertSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			j := i
			for j-1 >= 0 && arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				j--
			}
		}
	}
	return arr
}

//快速排序
func QuickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var f func(arr []int, low, high int)

	f = func(arr []int, low, high int) {
		if low < high {
			i, j := low, high
			key := arr[(low+high)/2]
			for i <= j {
				for arr[i] < key {
					i++
				}
				for arr[j] > key {
					j--
				}
				if i <= j {
					arr[i], arr[j] = arr[j], arr[i]
					i++
					j--
				}
			}

			if low < j {
				f(arr, low, j)
			}
			if high > i {
				f(arr, i, high)
			}
		}

		return
	}
	f(arr, 0, len(arr)-1)

	return arr
}
