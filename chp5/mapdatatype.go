package main

import (
	"errors"
	"fmt"
)

/**
	C风格函数
	不存在则插入，存在则替换
 */
//func PutAbsent(ages map[string]int, name string, age int) {
//	if someoneAge, ok := ages[name]; !ok {	//不存在， 这种写法Java也可以有，C的话待验证
//		ages[name] = age
//	} else {	//存在
//
//	}
//}

func GetAbsent(ages map[string]int, name string) int {
	if age, ok := ages[name]; ok {	//不存在， 这种写法Java也可以有，C的话待验证
		return age
	} else {
		return -1		//TODO：重新定义为一个不合理的数值，根据Go的设计理念这里是不是不应该像Java那样抛一个自定义的异常？
	}
}

//函数不能重载
func GetAbsent2(ages map[string]int, name string) (int, error) {
	if age, ok := ages[name]; ok {	//不存在， 这种写法Java也可以有，C的话待验证
		return age, nil
	} else {
		return 0, errors.New("key of this map is not exist")
	}
}

func main()  {

	//map类型分配空间
	ages := make(map[string]int)

	//添加新的键值对
	ages["Arvin"] = 21
	ages["Blank"] = 22
	ages["Cindy"] = 23
	//修改某个键值对
	ages["Arvin"] = 24
	for name, age := range ages {
		fmt.Println("name:", name, " age:", age)
	}

	ages["David"] = 25
	fmt.Println("len: ", len(ages))
	delete(ages, "David")
	fmt.Println("len: ", len(ages))

	fmt.Println(ages["David"])					//不存在的话默认返回0，这种会不会有问题，万一写这行的人以为"David"还存在（其实被删除了）
	fmt.Println(GetAbsent(ages, "David")) //重实现一个同样功能的函数但是返回一个不合理的值
	if davidAge, err := GetAbsent2(ages, "David"); err != nil {
		//先确认David年龄
		ages["David"] = 25
	} else {
		fmt.Println(davidAge)
	}

}