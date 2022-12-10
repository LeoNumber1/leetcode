package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "5F3Z-2e-9-w"
	k := 4

	s = "2-5g-3-J"
	k = 2

	fmt.Println(licenseKeyFormatting(s, k))
	fmt.Println(math.MinInt32)
}

func licenseKeyFormatting(s string, k int) string {
	n := len(s)
	sb := &strings.Builder{}
	for i := n - 1; i >= 0; {
		for j := 0; j < k; {
			if i >= 0 && s[i] != '-' {
				sb.WriteByte(s[i])
				j++
			}
			i--
			if i < 0 {
				break
			}
		}
		sb.WriteByte('-')
	}
	temp := sb.String()
	ans := &strings.Builder{}
	for i := len(temp) - 1; i >= 0; i-- {
		ans.WriteByte(temp[i])
	}
	return strings.TrimLeft(strings.ToUpper(ans.String()), "-")
}
