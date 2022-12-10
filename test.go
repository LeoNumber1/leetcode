package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/cheggaaa/pb/v3"
)

type Member struct {
	Ids []int
}

func main() {
	pbTest()
	//pbTest1()
	//regTest()
	return
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

func pbTest() {
	count := 100000
	// create and start new bar
	//bar := pb.StartNew(count)

	// start bar from 'default' template
	//bar := pb.Default.Start(count)

	// start bar from 'simple' template
	//bar := pb.Simple.Start(count)

	// start bar from 'full' template
	bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}

func pbTest1() {
	var limit int64 = 1024 * 1024 * 1024 * 5
	// we will copy 200 Mb from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar := pb.Full.Start64(limit)
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	io.Copy(writer, barReader)
	// finish bar
	bar.Finish()
}

func regTest() {
	str := "123sino123"
	str1 := "123SiNo123"
	matched, err := regexp.MatchString("(?is)sino", str)
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("(?is)sino", str1)
	fmt.Println(matched, err)
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
