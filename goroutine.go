package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
	test7()
}

func test1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var skip int
		for {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			fmt.Printf("%s:%d\n", file, line)
			skip++
		}
		wg.Done()
	}()
	wg.Wait()
}

func test2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("_______________")
		runtime.Goexit()
		fmt.Println("123123")
	}()
	wg.Wait()
}

func test3() {
	fmt.Println("main 函数 开始...")
	go func() {
		fmt.Println("父 协程 开始...")
		go func() {
			for {
				fmt.Println("子 协程 执行中...")
				timer := time.NewTimer(time.Second * 2)
				<-timer.C
			}
		}()
		time.Sleep(time.Second * 5)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("main 函数 退出")
}

func test4() {
	fmt.Println("main 函数 开始...")
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		fmt.Println("父 协程 开始...")
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("子 协程 接受停止信号...")
					return
				default:
					fmt.Println("子 协程 执行中...")
					timer := time.NewTimer(time.Second * 1)
					<-timer.C
				}
			}
		}(ctx)
		time.Sleep(time.Second * 5)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("main 函数 退出")

}

func test5() {
	ch := make(chan int)
	//定义一个WaitGroup，阻塞主线程执行
	var wg sync.WaitGroup
	//添加一个goroutine等待
	wg.Add(1)
	//goroutine超时
	go func() {
		//执行完成，减少一个goroutine等待
		defer wg.Done()
		for {
			select {
			case i := <-ch:
				fmt.Println(i)
			//goroutine内部3秒超时
			case <-time.After(3 * time.Second):
				fmt.Println("goroutine1 timed out")
				return
			}
		}
	}()
	ch <- 111
	//新增一个1秒执行一次的计时器
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	//新增一个10秒超时的上下文
	background := context.Background()
	ctx, _ := context.WithTimeout(background, 10*time.Second)
	//添加一个goroutine等待
	wg.Add(1)
	go func(ctx context.Context) {
		//执行完成，减少一个goroutine等待
		defer wg.Done()
		j := 1
		for {
			select {
			//每秒一次
			case <-ticker.C:
				fmt.Println("tick " + strconv.Itoa(j))
				j++
				ch <- j
			//内部超时，不会被执行
			case <-time.After(5 * time.Second):
				fmt.Println("goroutine2 timed out")
				return
			//上下文传递超时信息，结束goroutine
			case <-ctx.Done():
				fmt.Println("goroutine2 done")
				return
			}
		}
	}(ctx)
	//等待所有goroutine执行完成
	wg.Wait()
}

func test6() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	charCha := make(chan int, 1)
	numCha := make(chan int, 1)
	//end := make(chan int, 0)
	numCha <- 1
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-charCha
			fmt.Print(i + 1)
			numCha <- 1
		}
		//end <- 1
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-numCha
			fmt.Print(string('A' + i))
			charCha <- 1
		}
	}()

	//<-end
	wg.Wait()
	return
}

func test7() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	charChan := make(chan byte, 1)
	numChan := make(chan int, 1)
	var a byte
	var b int
	n := 20
	i := 0
	go func() {
		defer wg.Done()
		for {
			select {
			case a = <-charChan:
				fmt.Print(string(a))
				numChan <- int(a-'A') + 1
				i++
				//time.Sleep(time.Second)
				if i >= n {
					return
				}
			case b = <-numChan:
				fmt.Print(b)
				charChan <- byte('A' + b)
				i++
				//time.Sleep(time.Second)
				if i >= n {
					return
				}
			default:

			}
		}
	}()
	charChan <- 'A'
	wg.Wait()
}
