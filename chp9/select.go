package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(10e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 10 + 1
		time.Sleep(1e9)
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 10 + 2
		time.Sleep(2e9)
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		//default:		//default不会阻塞
		//	fmt.Printf("waiting ...\n")
		}
	}
}