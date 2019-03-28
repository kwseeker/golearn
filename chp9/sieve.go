package main

import "fmt"

/**
	使用有向通道实现的打印素数的功能
	虽然代码很短，但是代码写的很高效，也不太好理解
	https://go.tanglei.name/content/images/14.2_fig14.2.png?raw=true
 */
// Send the sequence 2, 3, 4, ... to returned channel
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)				//out用于接收筛选出的素数
	go func() {
		ch := generate()				//在一个协程中往通道ch不断插入2，3，4 ...
		for {
			prime := <-ch
			ch = filter(ch, prime)		//每筛选出一个素数就会多创建一个goroutine（内部会创建一个新的通道）并加入到对后续的整数的整除判断中
										//TODO: 这里感觉有点问题内存应该会占用越来越大
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
