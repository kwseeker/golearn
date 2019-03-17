package main

import "fmt"

type List []int

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func (l List) Len() int {				//List 继承 Lener 接口
	return len(l)
}

func (l *List) Append(val int) {		//*List 继承 Appender 接口
	*l = append(*l, val)
}

func CountInto(a Appender, start, end int) {	//*List
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {					//List， *List (之所以可以传入*List类型变量是因为Go指针可以被自动解引用)
	return l.Len()*10 > 42
}

func main()  {
	var lst List
	CountInto(&lst, 1, 10)
	if LongEnough(lst) {
		fmt.Println("- lst is long enough")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Println("- plst is long enough")
	}
}