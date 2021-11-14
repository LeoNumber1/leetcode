package main

import "fmt"

func main() {
	tests := struct {
		args  []string
		wants []bool
	}{
		args: []string{
			"aaF", "USA", "asdas", "FlaG", "Flag", "aA", "A", "AAa",
		},
		wants: []bool{
			false, true, true, false, true, false, true, false,
		},
	}
	var all bool = true
	for k, v := range tests.args {
		if detectCapitalUse(v) != tests.wants[k] {
			fmt.Printf("err in %d, string is %s", k, v)
			all = false
		}
	}
	if all {
		fmt.Println("all tests ok")
	}
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	var first, allBig, hasSmall bool
	var bigCount int
	if word[0] <= 'Z' {
		first = true
		allBig = true
		bigCount++
	}
	for _, ch := range word[1:] {
		if ch <= 'Z' {
			if first {
				if hasSmall {
					return false
				}
				bigCount++
				allBig = true
			} else {
				return false
			}
		} else {
			hasSmall = true
			allBig = false
			if bigCount > 1 {
				return false
			}
		}
	}
	if allBig {
		return true
	}
	return hasSmall
}
