package main

import (
	"fmt"
	"math"
)

func main() {
	type args struct {
		s string
	}

	tests := []struct {
		index  int
		args   args
		target int
	}{
		{1, args{"0032"}, 32},
		{2, args{"42"}, 42},
		{3, args{"   -42"}, -42},
		{4, args{"4193 with words"}, 4193},
		{5, args{"-0"}, 0},
		{6, args{"-+12"}, 0},
		{7, args{"+-12"}, 0},
		{8, args{"words and 987"}, 0},
		{9, args{"w10734368512"}, 0},
		{10, args{"+1"}, 1},
		{11, args{"10734368512"}, 2147483647},
		{12, args{"00000-42a1234"}, 0},
		{13, args{"-13+8"}, -13},
	}

	var errNum bool
	for _, tt := range tests {
		result := myAtoi(tt.args.s)
		if tt.target != result {
			errNum = true
			fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " get result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests are OK! -------")
	}
}

func myAtoi(s string) int {
	ans := 0
	negative := false
	start := true
	max := false
	for _, b := range []byte(s) {
		if b == ' ' {
			if !start {
				break
			}
			continue
		}
		if !negative && b == '-' {
			if start {
				negative = true
				start = false
				continue
			}
			break
		}
		if b == '+' {
			if !start {
				break
			}
			start = false
			negative = false
			continue
		}
		if ok, i := convertNum(b); ok {
			ans = ans*10 + i
			start = false
		} else {
			break
		}
		if ans > math.MaxInt32 {
			max = true
			ans = math.MaxInt32
			break
		}
	}
	if max {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if negative {
		return -ans
	}
	return ans
}

var m = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func convertNum(b byte) (bool, int) {
	if v, ok := m[b]; ok {
		return true, v
	}
	return false, 0
}
