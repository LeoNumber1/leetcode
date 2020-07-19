package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"

	//s1 = "a"
	//s2 = ""
	//s3 = "c"

	//s1 = "aa"
	//s2 = "ab"
	//s3 = "aaba"

	//s1 = "aabd"
	//s2 = "abdc"
	//s3 = "aabdabcd"

	//s1 = "aabcc"
	//s2 = "dbbca"
	//s3 = "aadbcbbcac"
	fmt.Println(isInterleave6(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveTrue(s1, s2, s3) || isInterleaveTrue(s2, s1, s3)
}

func isInterleaveTrue(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var index1, index2 int
	whichString := 1
	for _, v := range s3 {
		if whichString == 1 {
			if index1 < len(s1) {
				if v == int32(s1[index1]) {
					whichString = 1
					index1++
					continue
				}
			}
			if len(s2) != 0 && index2 < len(s2) {
				if v == int32(s2[index2]) {
					whichString = 2
					index2++
					continue
				}
			}
		} else {
			if index2 < len(s2) {
				if v == int32(s2[index2]) {
					whichString = 2
					index2++
					continue
				}
			}
			if len(s1) != 0 && index1 < len(s1) {
				if v == int32(s1[index1]) {
					whichString = 1
					index1++
					continue
				}
			}
		}
		return false
	}
	return true
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return isInterleaveTrue2(s1, s2, s3) || isInterleaveTrue2(s2, s1, s3)
}

func isInterleaveTrue2(s1, s2, s3 string) bool {
	arr1 := strings.Split(s1, "")
	arr3 := strings.Split(s3, "")

	var index1 int
	for k, v := range arr3 {
		if index1 < len(arr1) {
			if v == arr1[index1] {
				arr3[k] = ""
				index1++
			}
		} else {
			break
		}
	}

	s3 = strings.Join(arr3, "")

	if s2 == s3 {
		return true
	}

	return false
}

func isInterleave3(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	m := isInterleaveTrue3(s1, s2, s3)
	if _, isExist := m[s2]; isExist {
		return true
	}
	return false
}

func isInterleaveTrue3(s1 string, s2 string, s3 string) map[string]int {
	m := map[string]int{}
	arr1 := strings.Split(s1, "")
	arr3 := strings.Split(s3, "")
	var s4 string

	var index1 int
	var lastK int
	var lastV string
	for k, v := range arr3 {
		if index1 < len(arr1) {
			if v == arr1[index1] {
				arr3[k] = ""
				//if index1+1 < len(arr1) {
				index1++
				//}
				lastK = k
				lastV = v
			}
		} else {
			s4 = strings.Join(arr3, "")
			if _, isExist := m[s4]; !isExist {
				m[s4] = 1
			}
			index1--
			arr3[lastK] = lastV
			//}
		}
	}

	s4 = strings.Join(arr3, "")
	if _, isExist := m[s4]; !isExist {
		m[s4] = 1
	}

	return m
}

func isInterleave4(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	mAll := map[string]int{}

	for i := 0; i < len(s1); i++ {
		m := map[string]int{}
		m = isInterleaveTrue4(s1, s2, s3, i+1)
		for k, v := range m {
			mAll[k] = v
		}
	}

	//m := isInterleaveTrue4(s1, s2, s3, 1)
	if _, isExist := mAll[s2]; isExist {
		return true
	}
	return false
}

func isInterleaveTrue4(s1, s2, s3 string, count int) map[string]int {
	m := map[string]int{}
	arr1 := strings.Split(s1, "")
	arr3 := strings.Split(s3, "")
	//arr4 := arr3
	var s4 string

	var index1 int
	//var lastK int
	//var lastV string
	var arr []int
	//mapCount := map[int]string{}
	for k, v := range arr3 {
		if index1 < len(arr1) {
			if v == arr1[index1] {
				if index1+1 == len(arr1) {

					arr3[k] = ""
					s4 = strings.Join(arr3, "")
					if _, isExist := m[s4]; !isExist {
						m[s4] = 1
					}
					if count == 1 {
						arr3[arr[0]] = v
					}
					for z := 0; z < count; z++ {

					}
				} else {
					//mapCount[k] = v
					arr = append(arr, k)
					arr3[k] = ""
					index1++
				}
				//lastK = k
				//lastV = v
			}
		} else {
			s4 = strings.Join(arr3, "")
			if _, isExist := m[s4]; !isExist {
				m[s4] = 1
			}
			index1--
			//for j := lastK; j < len(arr3[lastK:]); j++ {
			//	arr3[j] = arr4[j]
			//}
			//arr3[lastK:] = arr4[lastK:]
			//}
		}
	}

	//s4 = strings.Join(arr3, "")
	//if _, isExist := m[s4]; !isExist {
	//	m[s4] = 1
	//}

	return m
}

var result map[string]int

func isInterleave5(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	result = map[string]int{}

	//for i := 0; i < len(s1); i++ {
	arr1 := strings.Split(s1, "")
	arr3 := strings.Split(s3, "")
	arr4 := make([]string, len(arr3), len(arr3))
	fmt.Println(copy(arr4, arr3))
	isInterleaveTrue5(arr1, s2, arr3, arr4, 0)
	//}

	//m := isInterleaveTrue4(s1, s2, s3, 1)
	if _, isExist := result[s2]; isExist {
		return true
	}
	return false
}

func isInterleaveTrue5(arr1 []string, s2 string, arr3, arr4 []string, start int) bool {
	//s1 = "aabd"
	//s2 = "abdc"
	//s3 = "aabdabcd"

	//var position int
	//arr4 := arr3
	//var arr4 []string

	//arr5 := make([]string, len(arr3), len(arr3))
	//fmt.Println(arr3  arr5)
	var s4 string

	var index1 int
	//var arr []int
	for i := start; i < len(arr3); i++ {
		if index1 < len(arr1) {
			if arr3[i] == arr1[index1] {
				if index1+1 == len(arr1) {
					arr4[i] = ""
					s4 = strings.Join(arr4, "")
					if s2 == s4 {
						return true
					}
					arr4[i] = arr3[i]
					index1++
					i--
					//	start = i + 1
					//	return isInterleaveTrue5(arr1, s2, arr3, arr4, start)
					//	//i = i - 1
					//
					//	//if count == 1 {
					//	//	arr3[arr[0]] = v
					//	//}
				} else {
					//if count == index1 {
					//
					//}
					//arr = append(arr, k)
					arr4[i] = ""
					index1++
				}
				//lastK = k
				//lastV = v
			}
		} else {
			arr4[i] = ""
			s4 = strings.Join(arr4, "")
			if s2 == s4 {
				return true
			}
			arr4[i] = arr3[i]
			start = i + 1
			return isInterleaveTrue5(arr1, s2, arr3, arr4, start)

			s4 = strings.Join(arr4, "")
			if s4 == s2 {
				return true
			}
			index1--
			//for j := lastK; j < len(arr3[lastK:]); j++ {
			//	arr3[j] = arr4[j]
			//}
			//arr3[lastK:] = arr4[lastK:]
			//}
		}
	}
	return false
}

func isInterleave6(s1, s2, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}

	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	fmt.Println(f)
	return f[n][m]
}
