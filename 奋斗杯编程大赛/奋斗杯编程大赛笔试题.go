package main

import "fmt"

func main() {
	subject13()
	//subject21()
}

func subject13() {
	n := 3
	haoni(n, 'A', 'B', 'C')
}

func haoni(n int, a, b, c byte) {
	if n == 1 {
		fmt.Println(n, ":", string(a), "-->", string(c))
	} else {
		haoni(n-1, a, c, b)
		fmt.Println(n, ":", string(a), "-->", string(c))
		haoni(n-1, b, a, c)
	}
}

func subject21() {
	const N = 110000
	const P = 10007
	var n, length, ans int
	var a [N]int
	getDivisor := func() {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				length++
				a[length] = i
				//if n/i != i {
				//	length++
				//	a[length] = n / i
				//}
			}
		}
	}

	var gcd func(int, int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	n = 4
	getDivisor()
	for i := 1; i <= length; i++ {
		for j := i + 1; j <= length; j++ {
			ans += gcd(a[i], a[j])
		}
	}

	fmt.Println(ans)
}
