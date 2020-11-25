package main

func main() {

}

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		//return byte(" ")
		return ' '
	}
	var m = make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range s {
		if val, has := m[v]; has && val == 1 {
			return byte(v)
		}
	}
	return ' '
}
