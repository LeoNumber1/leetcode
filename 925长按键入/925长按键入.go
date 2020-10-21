package main

import "fmt"

func main() {
	name := "alex"
	typed := "aaleex"

	//name = "saeed"
	//typed = "ssaaedd"

	//name = "leelee"
	//typed = "lleeelee"
	//typed = "leelee"

	name = "vtkgn"
	typed = "vttkgnn"

	//fmt.Println(isLongPressedName(name, typed))
	fmt.Println(isLongPressedNameOfficial(name, typed))
}

//0 ms-100.00%	2 MB-73.33%
func isLongPressedName(name string, typed string) bool {
	lenName := len(name)
	lenTyped := len(typed)
	if lenTyped < lenName {
		return false
	}
	var i, j int
	for i < lenName && j < lenTyped {
		if name[i] == typed[j] {
			i++
			j++
			continue
		}
		if j > 0 && typed[j] == typed[j-1] {
			j++
			continue
		}
		return false
	}
	if i == lenName && j < lenTyped {
		last := name[lenName-1]
		for i := j; i < lenTyped; i++ {
			if typed[i] != last {
				return false
			}
			j++
		}
	}
	return i == lenName && j == lenTyped
}

//0 ms-100.00%	2 MB-73.33%
func isLongPressedNameOfficial(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}
