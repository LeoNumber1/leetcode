package main

import "fmt"

func main() {
	s := "?j?qg??b"
	s = "?"
	s = "??yw?ipkj?"
	//s = "???"

	fmt.Println(modifyString(s))
}

var word []byte

//m := map[byte]byte{
//	'a': 'a',
//	'b': 'b',
//	'c': 'c',
//	'd': 'd',
//	'e': 'e',
//	'f': 'f',
//	'g': 'g',
//	'h': 'h',
//	'i': 'i',
//	'j': 'j',
//	'k': 'k',
//	'l': 'l',
//	'm': 'm',
//	'n': 'n',
//	'o': 'o',
//	'p': 'p',
//	'q': 'q',
//	'r': 'r',
//	's': 's',
//	't': 't',
//	'u': 'u',
//	'v': 'v',
//	'w': 'w',
//	'x': 'x',
//	'y': 'y',
//	'z': 'z',
//}

func modifyString(s string) string {
	word := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	var index int
	var ans string

	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			if i == len(s)-1 {
				if len(ans) > 0 {
					for _, v := range word {
						if v != ans[index] {
							ans += string(v)
							break
						}
					}
				} else {
					ans += string(word[0])
				}
			} else {
				for _, v := range word {
					if len(ans) > 0 {
						if v != ans[index] && v != s[i+1] {
							ans += string(v)
							break
						}
					} else {
						if v != s[i+1] {
							ans += string(v)
							break
						}
					}
				}
			}
		} else {
			ans += string(s[i])
		}
		index = i
	}

	return ans
}
