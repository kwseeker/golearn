package main

import "fmt"

//常量
const (
	_ = 1 << (10 * iota)
	KiB     //1024
	MiB     //1048576
	GiB     //1073741824
)

func main() {
	//整数
	//var v1 uint8
	//var v2 int8
	//var v3 uint16
	//var v4 int16
	//var v5 uint32
	//var v6 int32
	//var v7 uint64
	//var v8 int64
	//
	//var v9 uint			//和具体的CPU平台有关，32位或者64位
	//var v10 int
	//var v11 uintptr 	//bit不确定的无符号整形，但足够容纳指针，通常用于存储指针地址


	//浮点型
	//var v12 float32		//1个符号位，23个整数位，8个浮点位	1<<24 = 16777216
	//var v13 float64

	//复数
	//var v14 complex64 = 1 + 2i
	var v15 complex128 = complex(1, 2) //1+2i
	var v16 complex128 = complex(3, 4) 	//3+4i
	v17 := v15 * v16
	fmt.Println(v17)
	fmt.Println(real(v17))	//real()内建函数
	fmt.Println(imag(v17))

	//布尔型
	//var v18 bool = true

	//字符串
	//v19 := "this is a string var"
	//var v20 string

	//常量
	fmt.Println("KiB:", KiB)
}
