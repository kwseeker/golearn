package main

import "fmt"

/**
	空接口，类似Java Object 般的存在
	空接口变量可以接收任何类型
 */

type Any interface {}

type StructType struct {
	val1 int
	val2 string
}

func main() {
	varInt := 1
	varStr := "Hello"
	varStructType := StructType{2, "World"}
	arrInt := []int{11,12,13}

	arrAny := []Any{varInt, varStr, varStructType, arrInt}
	for _, v := range arrAny {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case StructType:
			fmt.Println("StructType")
		case []int:
			fmt.Println("[]int")
		default:
		}
	}
}