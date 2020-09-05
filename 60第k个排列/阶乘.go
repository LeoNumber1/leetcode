package main

import "fmt"

var facVal uint64 = 1

var i int = 1
var n int

/* function declaration */

func factorial1(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			facVal *= uint64(i)
		}
	}
	return facVal

}

func main() {
	fmt.Println("Enter a positive integer between 0 - 50 : ")
	fmt.Scan(&n)
	fmt.Print("Factorial is: ", factorial1(n))
}
