package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string] int)		//string, int 键值对
	length := len(counts)
	fmt.Println("counts length: ", length)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {					//input.Scan() 是否有值
		if input.Text() == "EOF" {
			break
		}
		counts[input.Text()]++
	}

	length2 := len(counts)
	fmt.Println("counts length: ", length2)

	for line, n := range counts {	//不是按顺序取值的？
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}