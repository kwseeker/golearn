# Go 基础

## 1 Go 与 Java
Go是直接编译成机器语言，可以直接运行，Java是解释执行语言需要依赖JVM运行；

Go同样是面向对象的，但与Java是不同的实现方式；

Go有类似于Java的垃圾回收机制；

Go是一种函数式语言；

Go支持多返回值，Java只能通过将多个返回包装成新的类返回；

Go具备动态语言特性，在对象初始化、构造和序列化等方面提供了无比简洁的表述方式；

Go相对于Java内存消耗更低。

## 2 入门

#### 2.1 Hello,world

Go语言允许为struct(类似Java的class，不同于C/C++的struct)提供动态拓展；如
```golang
package main

import "fmt"

type Person struct {
    name string
}

func (person *Person) setName(name string) {
    person.name = name
}

func (person *Person) GetInfo() {
    fmt.Println(person.name)
}

func main() {
    p := Person{"zhangsan"}
    p.setName("lisi")
    p.GetInfo()
}
```
通过(person *Person)为Person拓展了两个新的方法，(person *Person)并不是返回值，Go方法返回值是定义在参数之后的，
里面的person指针其实类似Java的this。

Go 通过方法名的首字母大小写区分public和private，而且Go只有public和private这两种
权限，public表示所有包可以使用，private表示只有包内可以使用。

#### 2.2 命令行参数

os.Args 保存所有的参数 os.Args[0] 是命令本身， os.Args[1:] 是传给命令的参数。

局部变量定义,详细参考3.3节。
```
var s, sep string
```

Go 代码行结束不需要";"标示。

Go 的 for循环没有"()"。

Go 没有while循环，而是使用下面的方式。
```
for condition {
    // ...
}

//无限循环
for {
    // ...
}
```

Go 的foreach实现。
```
for _, arg := range os.Args[1:] {   // _ 表示忽略第一个返回值
    // ...
}
```

Go 字符串高效拼接
```
strings.Join(os.Args[1:], " ")
```

#### 2.3 查找重复的行

make() 是builtin包的函数（注意不是方法）；用于分配和初始化某类型对象，如slice，map，chan；
```
func make(t Type, size ...IntegerType) Type
```

Go 语言的函数和方法的区分

方法（是面向对象的概念）是包含接收者的，接收者可以是命名类型或者结构体类型的一个值或者是一个指针。如2.1章节。
注意：Go语言不允许为简单的内置类型添加方法。
```
func (person *Person) setName(name string) {
    person.name = name
}
```
而函数不包含接受者，如上面的make函数。
同样函数也是通过函数名大小写区分public和private的。

然后可能会有疑问为什么make函数，看起来是一个private的函数，却可以在自己的代码随意使用，
还不用带包名？
因为make函数是Go语言的内建函数。

内建函数

Go预定义了少数内建函数（属于语言本身而不是包），这意味着无需引用任何包就可以使用它们。
而且内建函数都是小写开头的，但是可以public访问的。
```sh
# 查看内建函数和类型
godoc builtin
```

Go 的 map写法
```
//Go
counts = make(map[string] int)
//Java
Map counts = new HashMap<String, int>();
```
Go语言的map，可以使用 key 做索引，像数组一样访问;






#### 2.4 GIF动画

#### 2.5 获取URL

#### 2.6 并发获取多个URL

#### 2.7 Web服务

## 3 程序结构

#### 3.1 命令

#### 3.2 声明

#### 3.3 变量

#### 3.4 赋值

#### 3.5 类型

#### 3.6 包和文件

#### 3.7 作用域

## 4 基础数据类型

#### 4.1 整型

#### 4.2 浮点型

#### 4.3 复数

#### 4.4 布尔型

#### 4.5 字符串

#### 4.6 常量

## 5 复合数据类型

#### 5.1 数组

#### 5.2 Slice

#### 5.3 Map

#### 5.4 结构体

#### 5.5 JSON

#### 5.6 文本和HTML模版

## 6 函数

#### 6.1 函数声明

#### 6.2 递归

#### 6.3 多返回值

#### 6.4 错误

#### 6.5 函数值

#### 6.6 匿名函数

#### 6.7 可变参数

#### 6.8 Deferred函数

#### 6.9 Panic异常

#### 6.10 Recover捕获异常

## 7 方法

#### 7.1 方法声明

#### 7.2 基于指针对象的方法

#### 7.3 通过嵌入结构体来拓展类型

#### 7.4 方法值和方法表达式

#### 7.5 示例：Bit数组

#### 7.6 封装

## 8 接口

#### 8.1 接口是合约

#### 8.2 接口类型

#### 8.3 实现接口的条件

#### 8.4 flag.Value接口

#### 8.5 接口值

#### 8.6 sort.Interface接口

#### 8.7 http.Handler接口

#### 8.8 error接口

#### 8.9 示例：表达式求值

#### 8.10 类型断言

#### 8.11 基于类型断言识别错误类型

#### 8.12 通过类型断言查询接口

#### 8.13 类型分支

#### 8.14 示例：基于标记的XML解码

#### 8.15 补充几点

## 9 Goroutines 和 Channels

#### 9.1 Goroutines

#### 9.2 示例：并发的Clock服务

#### 9.3 示例：并发的Echo服务

#### 9.4 Channels

#### 9.5 并发的循环

#### 9.6 示例：并发的web爬虫

#### 9.7 基于select的多路复用

#### 9.8 示例：并发的字典遍历

#### 9.9 并发的退出

#### 9.10 示例：聊天程序

## 10 基于共享变量的并发

#### 10.1 竞争条件

#### 10.2 sync.Mutex互斥锁

#### 10.3 sync.RWMutex读写锁

#### 10.4 内存同步

#### 10.5 sync.Once初始化

#### 10.6 竞争条件检测

#### 10.7 示例：并发的非阻塞缓存

#### 10.8 Goroutines和线程

## 11 包和工具

#### 11.1

#### 11.2

#### 11.3

#### 11.4

#### 11.5

#### 11.6

#### 11.7

## 12 测试

#### 12.1

#### 12.2

#### 12.3

#### 12.4

#### 12.5

#### 12.6

## 13 反射

#### 13.1

#### 13.2

#### 13.3

#### 13.4

#### 13.5

#### 13.6

#### 13.7

#### 13.8

#### 13.9

## 14 底层编程

#### 14.1

#### 14.2

#### 14.3

#### 14.4

#### 14.5