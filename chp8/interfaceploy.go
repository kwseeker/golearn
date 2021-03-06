package main

import "fmt"

/**
	接口多态
	接口变量可以接收任意实现接口方法的结构体变量而产生不同的行为
 */
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width;
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Circumference() float32 {
	return 4*sq.side
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r,q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var sq Shaper
	sq = &Square{6.0}
	fmt.Println("Area of this shape is: ", sq.Area())

	//接口变量的类型断言
	if v, ok := sq.(*Square); ok {
		//v 是 Square指针类型
		fmt.Println("So, this is a square, I can get it's circumference: ", v.Circumference())
	}

	//变量是否继承接口
	if _, ok := sq.(Shaper); ok {
		fmt.Println("So, this value extends Shaper")
	}
}