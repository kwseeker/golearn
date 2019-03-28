package main

import (
	"fmt"
)

/**
	函数作为返回值(函数式编程的写法)
	这样写有什么实际的用途？

	闭包函数积累内部变量的值
 */
func main() {
	methodVal := Plus2(2)
	fmt.Println(methodVal(3))

	var f = Adder()				//f中保持一个局部变量x，对每一次调用x都是共享的
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Println()

	var f2 = Adder()
	fmt.Println(f2(1))

	//f3 := fibonacciPro()
	//fmt.Println(f3(6))

	f4 := fibonacciPro2()
	for i := 0; i < 10; i++ {
		fmt.Println(f4())
	}
}

//使用闭包实现斐波那契数列，因为闭包内部的变量是每次调用效果叠加的
func fibonacciPro2() func() int {
	a := -1
	b := 1
	return func() int {
		a, b = b, a + b
		return b
	}
}

//func fibonacciPro() func(n int) int {
//	var ret1 int	//记录最终结果左边的左边的值
//	var ret2 int 	//记录最终结果左边的值
// 	return func(n int) int {
// 		if n == 1 || n == 2 {
// 			return 1
//		} else {
//			for i:=1; i<=n; i++ {
//				if i == 1 {
//					ret1 = 1
//				} else if i == 2 {
//					ret2 = 1
//				} else {
//					ret1, ret2 = ret2, ret1 + ret2
//				}
//			}
//			return ret2
//		}
//	}
//}


//TODO: 使用goroutine和channel解决递归堆栈溢出问题

//TODO: 懒惰求值解决递归堆栈溢出问题

//斐波那契数列传统实现：数据很大的话存在堆栈溢出的风险
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//返回函数的函数
func Plus2(step int) func(a int) int {
	return func(a int) int {
		return a + step
	}
}

//闭包函数积累内部变量的值
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
