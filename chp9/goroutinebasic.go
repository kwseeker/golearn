package main

import (
	"fmt"
	"time"
)

func longWait()  {
	fmt.Println("Begin longWait() ...")
	time.Sleep(5*1e9)
	fmt.Println("End of longWait() ...")
}

func shortWait()  {
	fmt.Println("Begin shortWait() ...")
	time.Sleep(2*1e9)
	fmt.Println("End of shortWait() ...")
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(1e9)
	ch <- "Tripoli"
	time.Sleep(1e9)
	ch <- "London"
}

func sendData2(ch chan string) {
	time.Sleep(5e9)
	ch <- "Beijing"
	ch <- "Tokio"
	ch <- "Hongkong"
}

func getData(ch chan string) {				//TODO: sendData()执行完之后，好像getData()知道还有函数要往channel中写东西似的而没有退出,Go怎么实现的
											//Go会检查所有协程是否在等待读写channel
	var input string
	for {									//TODO：为何通道读完之后就正常退出了,而不是继续循环
		fmt.Printf("Circle ... ")
		input = <- ch
		fmt.Printf("%s ", input)
	}
	fmt.Printf("Out of circle ... ")	//这个并没有被执行，好像channel空之后直接return了
}

func dealLock() {
	ch1 := make(chan string)
	sendData(ch1)
	getData(ch1)
}

func main()  {

	//fmt.Println("Begin main()...")
	//go longWait()
	//go shortWait()
	//fmt.Println("About to sleep in main() ...")
	////time.Sleep(10*1e9)
	//time.Sleep(4*1e9)
	//fmt.Println("At end of main() ...")

	/*=============================================*/

	//ch := make(chan string)
	//go sendData(ch)
	//go getData(ch)
	//time.Sleep(2e9)
	//go sendData2(ch)
	//time.Sleep(8e9)

	/*=============================================*/

	//ch1 := make(chan string)
	//ch1 <- "Hello"
	//temp := <- ch1
	//fmt.Println(temp)

	go dealLock()			//TODO：上面操作放到哪个协程哪个协程阻塞，为什么？放到main协程直接Panic，放到子协程无输出
	time.Sleep(5e9)
}