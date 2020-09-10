package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Money = 100

func main() {
	room := [100]int{}
	for i := 0; i < 100; i++ {
		room[i] = Money
	}
	for i := 0; i < 10000; i++ {
		change(&room)
	}
	fmt.Println(room)
}

func change(room *[100]int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		if (*room)[i] > 0 {
			(*room)[i]--
			(*room)[rand.Intn(100)]++
		}
	}
}
