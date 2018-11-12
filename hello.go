// golang 的编译安装运行命令
// 		go install	编译并安装生成的静态库或可执行文件
// 		go run
//		go build	测试编译，并不会安装
//		go test 	测试，使用GO的测试框架做单元测试

// 远程包的导入
//		import "github.com/golang/example/stringutil"
//		像Git或Mercurial这样的版本控制系统，可根据导入路径的描述来获取包源代码。go 工具可通过此特性来从远程代码库自动获取包。
//		也可手动获取构建并安装：如 go get github.com/golang/example/hello

// 关于上述内容参考： https://go-zh.org/doc/code.html
// 关于GO项目管理（包括管理工具）参考： https://www.zhihu.com/question/47538279
// 关于GO项目编译原理参考：https://halfrost.com/go_command/ （这篇文章写得很好）

// 现在GO不仅支持静态链接也支持动态链接

//package golearn		//package 不为main，go install 会打包成静态库，在 $GOPATH/pkg/<系统信息>/<包名>/<包静态库>
package main			//package 为main， go install 会打包成可执行文件，在 $GOPATH/bin/下

import (
	"fmt"
	"kwseeker.top/golearn/file"
	"os"
)

func main() {
	fmt.Println("Hello, world.")

	//===============================================================
	////flag.Parse()
	////fmt.Println("级别:", *levelFlag)
	////fmt.Println("份数:", bnFlag)
	//pbi, pwi := cmdline.ParseArgs()
	//if pbi != nil {
	//	fmt.Println("生日：", pbi.GetBirthDay())
	//}
	//if pwi != nil {
	//	fmt.Println("工作列表：", pwi.GetWorkList())
	//}

	//===============================================================
	path, err := file.GetCurrentPath()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current path: ", path);

	if isExist, _ := file.IsExist("test.txt"); !isExist {
		fmt.Println("file is not exist")
		_, err := os.Create("test.txt")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("file is already exist")
	}

	//fp, err := os.Open("test.txt")
	fp, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY,0666)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	wb, err := fp.Write([]byte("This is a sentence in test.txt"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("write %d bytes into test.txt\n", wb)
	fmt.Println(file.ReadContentToString("test.txt"))

}
