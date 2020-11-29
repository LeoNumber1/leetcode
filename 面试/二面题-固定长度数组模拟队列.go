package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	//testNormal()
	testConcurrence()
}

//测试普通调用
func testNormal() {
	queue := Queue{}
	var err error
	for i := 0; i < 12; i++ {
		if err = queue.In(i); err != nil {
			fmt.Println("入队失败,错误信息:", err)
		} else {
			fmt.Printf("%d入队->", i)
		}
	}
	fmt.Println("----------------------------")
	fmt.Println("当前队列长度:", queue.Len())
	fmt.Println("当前队列是否已满?", queue.IsFull())
	fmt.Println("----------------------------")
	for i := 0; i < 12; i++ {
		if v, err := queue.Out(); err != nil {
			fmt.Println("出队列失败,错误信息:", err)
		} else {
			fmt.Printf("%d出队->", v)
		}
	}
}

//测试并发调用
func testConcurrence() {
	queue := Queue{}
	var err error
	wg := sync.WaitGroup{}
	count := 10000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(w int) {
			defer wg.Done()
			for j := w; j < 12000; j++ {
				if err = queue.In(j); err != nil {
					//fmt.Println("入队失败,错误信息:", err)
				} else {
					fmt.Printf("%d入队->", j)
				}
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------------------")
	fmt.Println("当前队列长度:", queue.Len())
	fmt.Println("当前队列是否已满?", queue.IsFull())
	fmt.Println("----------------------------")
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for queue.len > 0 {
				if v, err := queue.Out(); err != nil {
					fmt.Println("出队列失败,错误信息:", err)
				} else {
					fmt.Printf("%d出队->", v)
				}
			}
		}()
	}
	wg.Wait()
}

/*
 *
 * 具体实现
 *
 */

const MAX = 10

type Queue struct {
	len        int      //队列长度
	head, tail int      //头尾指针位置
	arr        [MAX]int //实际队列元素存在的固定长度数组
	//lock       sync.Mutex
}

func (q *Queue) In(x int) error {
	//q.lock.Lock()
	//defer q.lock.Unlock()
	if q.len == MAX {
		return errors.New("队列已满")
	}
	q.arr[q.tail] = x   //入队将元素插入到队尾指针处
	q.tail++            //队尾指针向后移位
	if q.tail > MAX-1 { //如果队尾指针大于数组边界，则队尾指针指回0
		q.tail = 0
	}
	q.len++ //队列长度+1
	return nil
}

func (q *Queue) Out() (int, error) {
	//q.lock.Lock()
	//defer q.lock.Unlock()
	if q.len == 0 {
		return 0, errors.New("队列已空")
	}
	x := q.arr[q.head]  //出队时将队头指针出位置元素出队
	q.head++            //同时，对头指针向后移一位
	if q.head > MAX-1 { //如果队头指针大于数组边界，则队头指针指回0
		q.head = 0
	}
	q.len-- //队列长度-1
	return x, nil
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) IsFull() bool {
	if q.len < MAX {
		return false
	}
	return true
}
