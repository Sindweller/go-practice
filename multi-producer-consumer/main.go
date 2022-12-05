package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wgProc := &sync.WaitGroup{}
	wgCons := &sync.WaitGroup{}

	// 生产者5个
	for i := 0; i < 5; i++ {
		wgProc.Add(1)
		go producer(ch, wgProc, i)
	}

	// 消费者三个
	for i := 0; i < 3; i++ {
		wgCons.Add(1)
		go consumer(ch, wgCons, i)
	}

	// 等待生产者都执行完成
	wgProc.Wait()
	fmt.Println("all producer done")
	close(ch)
	wgCons.Wait()
	fmt.Println("all consumer done")
	return
}

// 生产者
func producer(ch chan int, wg *sync.WaitGroup, ii int) {
	defer fmt.Printf("producer %d quit\n", ii)
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("producer %d: %d \n", ii, i)
		ch <- i
		time.Sleep(time.Second)
	}
	return
}

// 消费者
func consumer(ch chan int, wg *sync.WaitGroup, ii int) {
	defer fmt.Printf("consumer %d quit\n", ii)
	defer wg.Done()
	for data := range ch {
		time.Sleep(2 * time.Second)
		fmt.Printf("consumer %d: %d \n", ii, data)
	}

	return
}
