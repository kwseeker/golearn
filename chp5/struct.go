package main

import "fmt"

/**
	Go的struct和Java的class差不多
 */
type Human struct {
	//基本数据类型
	name string
	age int
	height, weight float32

	//复杂数据类型
	hobby map[string]string

	//结构体
	girlFriend *Human	//结构体内部包含结构体自身变量，必须使用指针，不然无法确定结构体大小报异常

	//接口,函数，方法都通过接收者指定
}

type Skill interface {
	program()
	cook(ingredient ... string) string
}

//内部函数
func (person Human) live() {
	fmt.Println(person.name + " love music, movie, game!")
}

//方法
func (person *Human)work()  {
	fmt.Println(person.name + " dreaming to make a masterpiece!")
}

//接口实现
func (person *Human)program() {
	fmt.Println("Hum ...")
}
func (person *Human) cook(ingredient ... string) string {
	var food string
	for _, val := range ingredient {
		food += val
		food += "+"
	}
	return food
}

func main() {

	me := new(Human)		// new() 返回的是指针
	me.name = "Arvin"
	me.age = 25
	me.height = 180.0
	me.weight = 70.0
	me.hobby = map[string]string{"sing":"high", "draw":"low"}
	me.girlFriend = &Human{"Von", 25, 170, 60, map[string]string{"sing":"high", "draw":"high"}, nil}

	me.live()
	me.work()
	me.program()
	fmt.Println(me.cook("egg", "tomato"))
	fmt.Println(me.girlFriend.name)
	me.girlFriend.live()

	//skills := []Skill{me, me.girlFriend}
	var mySkill Skill
	mySkill = me
	mySkill.program()
	fmt.Println(mySkill.cook("beef", "onion"))

	var myGirlSkill Skill
	myGirlSkill= me.girlFriend
	myGirlSkill.program()
	fmt.Println(myGirlSkill.cook("beef", "onion"))
}
