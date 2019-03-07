package main

import (
	"bufio"
	"fmt"
	"os"
)

//流的方式读取文件
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string] int)		//
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)	//未传参数的话则从标准输入取。os.Stdin 直接赋值给 *os.File ? Stdin 属于 File类型的常量
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Println("line: ", line, "times:", n)
	}

	//TODO: 打印重复的行都出现在哪些文件

	//使用 ioutil.ReadFile() 读取文件

}