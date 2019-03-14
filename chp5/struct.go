package main

/**
	Go的struct和Java的class差不多
 */
type Person struct {
	//基本数据类型
	name string
	age int
	height, weight float32

	//复杂数据类型
	hobby map[string]string

	//结构体
	girlFriend Person

	//接口
	mySkill Skill

	//函数
	
}

type Skill interface {
	program()
	cook(ingredient ... string) string
}

func main() {



}
