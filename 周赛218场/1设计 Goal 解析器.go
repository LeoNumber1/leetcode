package main

import "fmt"

func main() {
	command := "G()(al)"

	fmt.Println(interpret(command))
}

func interpret(command string) string {
	var ans string
	for i := 0; i < len(command); {
		if command[i] == 'G' {
			ans += "G"
			i++
			continue
		}
		if command[i] == '(' && command[i+1] == ')' {
			ans += "o"
			i += 2
			continue
		}
		ans += "al"
		i += 4
	}
	return ans
}
