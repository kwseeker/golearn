package main

import "fmt"

/**
	使用空接口接收类型不确定且变长的参数
 */
func checkType(values ... interface{}) {
	for i, value := range values {
		switch value := value.(type) {
		case int:
			fmt.Println(i, "int:", value)
		case float64:
			fmt.Println(i, "float64:", value)
		case float32:
			fmt.Println(i, "float32:", value)
		case string:
			fmt.Println(i, "string:", value)
		case bool:
			fmt.Println(i, "bool:", value)
		default:
			fmt.Println(i, "other type:", value)
		}
	}
}

func main() {
	checkType(1, 1.2, true, "Arvin")
}
