package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func reverse(s []string)  {
	for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i:= range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	//数组
	var v1 [3]int		//默认元素会被初始化为0
	for i, v := range v1 {
		fmt.Println("index:", i, "value:", v)
	}

	symbol := [...]string {USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{99: -1}	//索引从0-99的数组，最后一个元素初始化为-1
	for _, v := range r {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Print(r)
	fmt.Println()

	//Slice
	months := [...]string{		//这是个数组
		1:"January",
		2:"February",
		3:"March",
		4:"April",
		5:"May",
		6:"June",
		7:"July",
		8:"August",
		9:"September",
		10:"October",
		11:"November",
		12:"December"}
	Q2 := months[4:7]		//前闭后开,索引并不会继承，这是个Slice
	fmt.Println(Q2)
	fmt.Println(len(Q2))	//3
	fmt.Println(cap(Q2))	//9, 4-12, 对一个数组做切片，则切片容量表示从左切点到数组最后元素

	fmt.Println(Q2[:8])		//从index 4 开始拓展8个
	//fmt.Println(Q2[:12])	//panic
	fmt.Println(len(Q2))

	reverse(Q2)
	for i,v := range Q2 {
		fmt.Println(i, v)
	}

	Q3 := months[4:7]
	fmt.Println("Q2==Q3?:", equal(Q2, Q3))

	//Slice添加新元素
	


	//Map



	//结构体

	//JSON

	//文本和HTML模版

}
