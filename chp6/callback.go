package main

import "fmt"

func Add(x int, y int) int {
	return x + y;
}

func callback(a int, b int, f func(int, int) int) int {
	return Add(a, b)
}

func main() {
	fmt.Println(callback(1,2, Add))
}