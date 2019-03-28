package main

import (
	"fmt"
	"time"
)

/**
	带缓冲的通道
 */

//使用不带缓冲的通道实现的Semaphore

//TODO: Linux C 的信号量 和 Java JUC 中的信号量的区别
type Semaphore struct{
	sems map[string]chan int	//键为string，值为 chan int
}
func (semaphore *Semaphore) SemInit(subSem string, ch chan int) {
	fmt.Println("SemInit", subSem)
	if semaphore.sems == nil {
		semaphore.sems = make(map[string]chan int)
	}
	semaphore.sems[subSem] = ch
}
func (semaphore *Semaphore) Signal(subSem string, val int) {
	fmt.Println("Signal", subSem, val)
	semaphore.sems[subSem] <- val
}
func (semaphore *Semaphore) CheckSignal() {
	fmt.Println("CheckSignal")
	for _, v := range semaphore.sems {
		<- v
	}
}


//测试使用带缓冲的通道实现gouroutine的通信

func accept(ch chan string, semaphore *Semaphore) { 	//1秒1取
	//semaphore.SemInit("accept", make(chan int))
	//defer semaphore.Signal("accept", 0)				//DONE：这一句为何没有执行？因为不可能被执行到

	for {
		fmt.Println("Acceptor: ", <- ch, time.Now().String())	//channel数据为空的话，会阻塞，所以这个函数不可能主动退出
		time.Sleep(1e9)
	}
	fmt.Println("Acceptor: You will never see this message be printed out！")
}

func send(ch chan string, semaphore *Semaphore) { 	//刚开始3秒1发，然后逐渐变快
	var i int64
	semaphore.SemInit("send", make(chan int))
	defer semaphore.Signal("send", 0)

	for i = 3000; i > 0; i=i-200 {
		ch <- fmt.Sprintf("%s %d", "msg", i)
		fmt.Println("Sender: ", "msg", i, time.Now().String())
		//time.Sleep(i * time.Microsecond)						//TODO: 为什么这么写不行
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
	//fmt.Println("close channel")
	//close(ch)
}

func monitor(ch chan string) {
	lastCap:= 0
	curCap := 0
	for {
		curCap = len(ch)
		if lastCap != curCap {
			fmt.Println("Channel len: ", curCap, time.Now().String())
			lastCap = curCap
		}
		time.Sleep(100*time.Millisecond)
	}
}

func main() {
	bufSize := 10
	ch := make(chan string, bufSize)
	semaphore := new(Semaphore)

	go monitor(ch)
	go accept(ch, semaphore)
	go send(ch, semaphore)

	//time.Sleep(60*1e9)
	//使用信号量模式优雅退出
	time.Sleep(3e9)
	semaphore.CheckSignal()
}
