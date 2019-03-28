package main

import (
	"fmt"
	"time"
)

/**
	使用无缓冲的channel实现的信号量模式
 */
func main() {

	ch := make(chan int)	//信号量

	go func() {
		//do something need lots of time
		time.Sleep(1e9)
		fmt.Println("I have finished my job")
		ch <- 1
	}()

	<- ch
}