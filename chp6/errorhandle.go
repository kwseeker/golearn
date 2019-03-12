package main

import "fmt"

//os.PathError的定义, 详见error.go
//type PathError struct {
//	Op string
//	Path string
//	Err error	//继承 error 内建接口
//}
//
//func (e *PathError) String() string {
//	return e.Op + " " + e.Path + " " + e.Err.Error()
//}

type MyError struct {
	Operation string
	Err error	// error接口有一个 Error()方法 要实现
}

//TODO: 接口继承 与 Go结构体指针
//公有方法
func (e *MyError) Error() string {
	return e.Operation + " " + e.Err.Error()
}

//公有函数
func NewMyError(oper string, err error) error {	//类似Java error err = new MyError()
	if err == nil {
		return nil
	}
	return &MyError{oper, err}
}

//测试函数，抛出 MyError 错误
func TestThrowMyError() {

}

//测试函数， fmt.Errorf() 抛出错误
func TestFmtErrorf() (int, error) {
	return 0, fmt.Errorf("some error happend")
}

func main() {



}
