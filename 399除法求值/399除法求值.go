package main

import "fmt"

func main() {

}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mCalc := make(map[string]float64)
	mExsit := make(map[string][]string)
	for key, val := range equations {
		mExsit[val[0]] = append(mExsit[val[0]], val[0]+"/"+val[1])
		if _, has := mExsit[val[1]]; !has {
			mExsit[val[1]] = []string{}
		}
		mCalc[val[0]+"/"+val[1]] = values[key]
	}
	fmt.Println(mExsit)
	fmt.Println(mCalc)
	var ans = make([]float64, len(queries))
	for k, val := range queries {
		_, val1has := mExsit[val[1]]
		_, val0has := mExsit[val[0]]
		if val0has && val1has {
			if val[1] == val[0] {
				ans[k] = 1
				continue
			}
			if v, has := mCalc[val[0]+"/"+val[1]]; has {
				ans[k] = v
				continue
			}
			if v, has := mCalc[val[1]+"/"+val[0]]; has {
				ans[k] = 1 / v
				continue
			}
			// ans[k] = v0 / v1
		} else {
			//两个有一个不存在，则是-1
			ans[k] = -1.0
		}
	}
	return ans
}
