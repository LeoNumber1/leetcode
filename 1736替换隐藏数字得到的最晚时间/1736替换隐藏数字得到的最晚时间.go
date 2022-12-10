package main

import "fmt"

func main() {
	fmt.Println(maximumTime("??:??"))
	fmt.Println(maximumTime("?4:03"))
}

func maximumTime(time string) string {
	timeByte := []byte(time)
	for i := 0; i < 5; i++ {
		if time[i] == '?' {
			switch i {
			case 0:
				if timeByte[1] != '?' && timeByte[1] > '3' {
					timeByte[i] = '1'
				} else {
					timeByte[i] = '2'
				}
			case 1:
				if timeByte[0] == '2' {
					timeByte[i] = '3'
				} else {
					timeByte[i] = '9'
				}
			case 3:
				timeByte[i] = '5'
			case 4:
				timeByte[i] = '9'
			}
		}
	}
	return string(timeByte)
}
