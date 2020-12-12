package main

import "fmt"

func main() {
	s := "RD"  //R
	s = "RDD"  //D
	s = "RDDR" //R
	//s = "RDDRDDR"              //D
	//s = "RRDDD"                //R
	//s = "RRDDDD"               //D
	//s = "DRRDRDRDRDDRDRDR"     //R
	//s = "DDRDRRRDRDDRRRDDRRDR" //R

	fmt.Println(predictPartyVictory1(s))
	//fmt.Println(predictPartyVictoryOfficial(s))
}

//8 ms-50.00%	6.8 MB-10.00%
func predictPartyVictory(senate string) string {
	m := make(map[int32]int, 2)
	for _, ch := range senate {
		m[ch]++
	}
	arr := []byte{}
all:
	for len(m) > 1 {
		n := len(senate)
		for i := 0; i < n; {
			arr = append(arr, senate[i])
			prev := senate[i]
			count := 1
			if i+1 >= n {
				i++
				goto next
			}
			for j := i + 1; j < n && count > 0; j++ {
				i = j + 1
				if senate[j] == prev {
					count++
					arr = append(arr, senate[j])
				} else {
					count--
					if m[int32(senate[j])] > 1 {
						m[int32(senate[j])]--
					} else {
						delete(m, int32(senate[j]))
						break all
					}
					if count == 0 {
						break
					}
				}
			}
		next:
			if count > 0 {
				for z := 0; z < len(arr)-1; z++ {
					if arr[z] != prev {
						count--
						if m[int32(arr[z])] > 1 {
							m[int32(arr[z])]--
						} else {
							delete(m, int32(arr[z]))
							break all
						}
						temp := make([]byte, len(arr[z+1:]))
						copy(temp, arr[z+1:])
						arr = append(arr[:z], temp...)
						if count == 0 {
							break
						}
					}
				}
			}
		}

		senate = string(arr)
		arr = []byte{}
	}
	for k := range m {
		if k == 'R' {
			return "Radiant"
		}
		return "Dire"
	}
	return ""
}

func predictPartyVictory1(senate string) string {
	markd := make([]bool, len(senate))
	bans := []int{0, 0} // R, D
	for {
		lefts := []int{0, 0} // R, D
		for i := 0; i < len(senate); i++ {
			if markd[i] {
				continue
			}
			var this_index int = 0
			var anoth_index int = 1
			if senate[i] == 'D' {
				this_index = 1
				anoth_index = 0
			}
			if bans[this_index] > 0 {
				bans[this_index] -= 1
				markd[i] = true
				continue
			}

			bans[anoth_index] += 1
			lefts[this_index] += 1

		}
		if lefts[0] == 0 {
			return "Dire"
		}
		if lefts[1] == 0 {
			return "Radiant"
		}
	}
}

func predictPartyVictoryOfficial(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
