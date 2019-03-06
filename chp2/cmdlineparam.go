package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {	// os.Args 保存所有的参数 os.Args[0] 是命令本身， os.Args[1:] 是传给命令的参数
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//strings包提供的字符串高效拼接
	s1 := strings.Join(os.Args[1:], " ")	//取字符串数组切片， 以字符串的格式打印
	fmt.Println(s1)

	//不关系格式也可以直接将字符串数组输出
	fmt.Println(os.Args)						//以数组的格式打印

	//练习1
	fmt.Println(os.Args[0])

	//练习2
	for j := 0; j < len(os.Args); j++ {
		fmt.Println(j, " ", os.Args[j])
	}
}
