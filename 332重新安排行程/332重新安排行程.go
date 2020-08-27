package main

import (
	"fmt"
	"sort"
)

func main() {
	//	[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
	tickets := [][]string{
		//{"MUC", "LHR"},
		//{"JFK", "MUC"},
		//{"SFO", "SJC"},
		//{"LHR", "SFO"},

		//{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},

		//{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"},

		{"EZE", "TIA"},
		{"EZE", "HBA"},
		{"AXA", "TIA"},
		{"JFK", "AXA"},
		{"ANU", "JFK"},
		{"ADL", "ANU"},
		{"TIA", "AUA"},
		{"ANU", "AUA"},
		{"ADL", "EZE"},
		{"ADL", "EZE"},
		{"EZE", "ADL"},
		{"AXA", "EZE"},
		{"AUA", "AXA"},
		{"JFK", "AXA"},
		{"AXA", "AUA"},
		{"AUA", "ADL"},
		{"ANU", "EZE"},
		{"TIA", "ADL"},
		{"EZE", "ANU"},
		{"AUA", "ANU"},
	}

	fmt.Println(findItinerary(tickets))
	//fmt.Println(findItineraryOfficial(tickets))
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	} else if len(tickets) == 1 {
		return tickets[0]
	}

	m := map[string][]string{}
	for _, v := range tickets {
		val, ok := m[v[0]]
		if !ok {
			m[v[0]] = []string{v[1]}
		} else {
			m[v[0]] = insert(val, v[1])
		}
	}

	var res = []string{}
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	/*
		方法有问题，可能会出现提前进入终点的情况
	*/
	//var res = []string{"JFK"}
	//for i := 0; i < len(tickets); i++ {
	//	k := res[len(res)-1]
	//	for key, val := range m[k] {
	//		if _, ok := m[val]; ok {
	//			res = append(res, m[k][key])
	//			m[k] = del(m[k], m[k][key])
	//			break
	//		} else if i == len(tickets)-1 {
	//			res = append(res, m[k][key])
	//		}
	//	}
	//}

	return res
}

func insert(arr []string, target string) []string {
	var tmp bool
	for k, v := range arr {
		if target < v {
			arr = append(arr, "")
			copy(arr[k+1:], arr[k:])
			arr[k] = target
			tmp = true
			break
		}
	}
	if !tmp {
		arr = append(arr, target)
	}
	return arr
}

func del(arr []string, target string) []string {
	for k, v := range arr {
		if v == target {
			arr = append(arr[:k], arr[k+1:]...)
			break
		}
	}
	return arr
}

func findItineraryOfficial(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}
	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
