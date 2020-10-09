package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Member struct {
	Ids []int
}

func main() {
	//testSlice()
	go dealSignal()
	exited := make(chan struct{}, 1)
	go channel1(exited)
	count := 0
	t := time.Tick(time.Second)
Loop:
	for {
		select {
		case <-t:
			count++
			fmt.Printf("main run %d\n", count)
		case <-exited:
			fmt.Println("main exit begin")
			break Loop
		}
	}
	fmt.Println("main exit end")
}

func testSlice() {
	member := &Member{}
	member.Ids = append(member.Ids, 1)
	member.Ids = append(member.Ids, 2)
	member.Ids = append(member.Ids, 3)
	//member.Ids = append(member.Ids, 0)
	a := append(member.Ids, 4)
	b := append(member.Ids, 5)
	c := append(member.Ids, 6)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

var exit = make(chan string, 1)

func dealSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit <- "shutdown"
	}()
}

func channel1(exited chan<- struct{}) {
	t := time.Tick(time.Second)
	count := 0
	for {
		select {
		case <-t:
			count++
			fmt.Printf("channel1 run %d\n", count)
		case <-exit:
			fmt.Println("channel1 exit")
			close(exited)
			return
		}
	}
}
