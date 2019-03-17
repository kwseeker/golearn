package main

import (
	"fmt"
	"sort"
)

/**
	Sorter 接口原理与使用 sort.sort.go

	默认已经支持int string float 切片的排序

	Sorter接口用于定义排序规则，除了上面三种类型对其他类型排序要首先实现Sorter接口，具体实现方法可以参考上面三种类型
	其他函数或方法都是在其基础上依靠各种算法实现的
	如：
		type IntSlice []int

		func (p IntSlice) Len() int           { return len(p) }
		func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
		func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

		// Sort is a convenience method.
		func (p IntSlice) Sort() { Sort(p) }
		func Sort(data Interface) {
			n := data.Len()
			quickSort(data, 0, n, maxDepth(n))
		}
*/

type Student struct {
	name string
	age int
}

type StudentSlice []Student

func (stuList StudentSlice) Len() int {
	return len(stuList)
}
func (stuList StudentSlice) Less(i, j int) bool {
	return stuList[i].name < stuList[j].name
}
func (stuList StudentSlice) Swap(i, j int) {
	stuList[i], stuList[j] = stuList[j], stuList[i]
}

//排序算法可以使用sort包 Sort(data Interface) 方法
func (stuList StudentSlice) Sort() {
	sort.Sort(stuList)
}

//也可以自己实现


func main()  {
	//[]int 标准库已经支持
	arr := sort.IntSlice{1, 4, 2, 9, 7, 8}
	arr.Sort()	//内部实现是快排
	fmt.Println("After sort: ", arr)

	//对Student结构体变量排序
	stuList := StudentSlice{Student{"Cindy", 21}, Student{"Arvin", 22},
		Student{"Frank", 22}, Student{"Erik", 20}}
	stuList.Sort()
	fmt.Println(stuList)
}
