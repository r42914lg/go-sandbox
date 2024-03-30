package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffered()
	buffered_closed()
	select_case()
	waitGroup()
}

// buffered channel - sleep to postone exit in main go-routine

func buffered() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	go job1(ch)
	time.Sleep(time.Second)
}

func job1(ch chan int) {
	for val := range ch {
		fmt.Println("Val is --> ", val)
	}
}

// buffered channel - close channel to postone exit to aviod deadlock

func buffered_closed() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	job2(ch)
}

func job2(ch chan int) {
	for val := range ch {
		fmt.Println("Val is --> ", val)
	}
}

// channel, go-routine & select

func select_case() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go job_ch1(ch1)
	go job_ch2(ch2)

	ch_1_complete := false
	ch_2_complete := false

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch_1_complete = true
			} else {
				fmt.Println("Channel #1 Val is --> ", msg)
			}
		case msg, ok := <-ch2:
			if !ok {
				ch_2_complete = true
			} else {
				fmt.Println("Channel #2 Val is --> ", msg)
			}
		}
		if ch_1_complete && ch_2_complete {
			return
		}
	}
}

func job_ch1(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 250)
		ch <- i
	}
	close(ch)
}

func job_ch2(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- i
	}
	close(ch)
}

//  WaitGroup

func waitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done")
}

func worker(num int) {
	fmt.Println("Starting job for i --> ", num)
	time.Sleep(time.Second)
	fmt.Println("Finished ", num)
}
