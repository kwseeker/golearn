# Go 基础

参考 [Go官方文档](https://golang.org/doc/)

包含以下部分：

[Go内存模型](https://golang.org/ref/mem)

[Go编程语言规范](https://golang.org/ref/spec#Introduction)

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

make() 与 new() 函数

make() 分配空间并初始化一个slice，map 或 chan 类型的对象;
返回一个参数指定类型的对象。

new() 返回参数指定类型的指针（内容为0）。

Go 的 map写法
```
//Go
counts = make(map[string] int)
//Java
Map counts = new HashMap<String, int>();
```
Go语言的map，可以使用 key 做索引，像数组一样访问;

Go的指针和C的指针

[Go与C指针的异同](https://blog.csdn.net/whatday/article/details/74931416)

在传参方面, 大体上看与C的相同；
C语言的数组本质是指针，而Go的数组和指针没关系，Go数组做参数会检查数组长度；
GO数组名只是变量名不是表示首个地址，C数组名表示数组首个地址；
Go语言取指针以及指针存储的值都是"."，而C分别是"."和"->"；
"&"与C语言一样是得到变量的指针；
取Go中指针的内容的值是不需要使用"*"操作符的，但是指针的指针(或者继续多层时)需要显式使用"*"符号。

Go文件 os.File

iota 常量生成器(iota从0自增)
```
const {
    _ = 1 << (10 * iota)
    KiB     //1024
    MiB     //1048576
    GiB     //1073741824
}
```




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

Go 对整形，浮点型和复数，支持"+"，"-"，一元操作。

整数位运算：&  |  ^  &^  <<  >>

前四种位运算不区分有符号和无符号。

#### 4.2 浮点型

注意精度和取值范围。

#### 4.3 复数

复数处理包 math/cmplx。

#### 4.4 布尔型

注意 Go/C/C++/Java 的 && 比 || 优先级高。

#### 4.5 字符串

Go 字符串处理包 bytes、strings、strconv、unicode。
实现了许多诸如字符串的查询、替换、比较、截断、拆分和合并等功能。

Go 字符串支持切片操作
```
fmt.Println(s[:5])
```

Go string 支持 == < > 比较。

#### 4.6 常量

常量表达式在编译期计算，而不是运行期。常量的潜在类型都是基础类型。

```
//单独声明
const pi = 3.14
//批量声明, 未赋值的值的初始值以前一个变量为准
const {
    e = 2.71828
    pi = 3.1415926
    a = 1.1
    b           //1.1
}
```



## 5 复合数据类型

#### 5.1 数组

固定长度，不够灵活所以使用较少。
更多的是使用Slice，因为Slice可以动态增长和收缩，使用更方便灵活。
```
var v1 [3]int
symbol := [...]string {USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
r := [...]int{99: -1}
```

```
for _, v := range r {
    fmt.Print(v, " ")
}
fmt.Print(r)            // "[...]"
```

#### 5.2 Slice

fmt.Println() 可以直接打印数组和Slice的值。

不超过Slice容量的

Go的骚操作，一个=，批量赋值。
```
func reverse(s []int)  {
	for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
		s[i], s[j] = s[j], s[i]     //省去了中间变量
	}
}
```

数组可以使用"比较"两个数组的值是否完全相同，但是Slice不可以这么做。


#### 5.3 Map

map的分配及初始化
```
ages := make(map[string]int)
ages := map[string]int {    //其实和数组的初始化还是挺像的，不过一个是数字做索引，一个使用字符串做key
    "alice": 21,
    "charlie": 22,
}
ages["alice"] = 21  //上面等同于
ages["charlie"] = 22
map[string]int{}    //创建空的map，但不是nil
```

map的常用操作
+ 添加新的键值对/为已有键赋值
+ 判断键是否存在并取值
+ 删除键值对
+ map遍历
+ 比较两个map是否相等

Go里面没有Set类型，而是以Map的key替代的。




#### 5.4 结构体

#### 5.5 JSON

#### 5.6 文本和HTML模版

#### 5.7 指针

Go 语言为程序员提供了控制数据结构的指针的能力；但是，你不能进行指针运算。
通过给予程序员基本内存布局，Go 语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式，
这些对构建运行良好的系统是非常重要的：指针对于性能的影响是不言而喻的，
而如果你想要做的是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。

!!! 对Go和C指针使用上的异同，参考前面的链接，整体上很像。
但是从功能, Go的指针是一个丢失了灵魂技能的"伪指针"，因为不支持指针运算，
导致"内存偏移"这一便捷快速的操作无法使用。

需要注意的几点：

1）Go指针不允许指针运算，因为使用不当的话，很容易造成内存泄漏以及程序崩溃。

2）Go语言不能获取const常量的地址，因为const常量是不允许修改的。



## 6 函数

#### 6.1 函数声明

#### 6.2 递归

#### 6.3 多返回值

#### 6.4 错误

#### 6.5 函数值

#### 6.6 匿名函数

#### 6.7 可变参数

#### 6.8/9/10 Deferred Panic Recover 错误处理机制

Go不推荐使用异常处理；

Go的设计者认为try/catch机制太过耗费资源，所以提供了更轻量级的异常处理机制，
且只推荐作为处理错误（普通的异常而非错误就不要用）的最终措施。

错误定义(可以参考Go源码的实现，如：os error.go)
```
//使用errors.New定义
var (
	ErrInvalid    = errors.New("invalid argument") // methods on File will return this error when the receiver is nil
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
	ErrClosed     = errors.New("file already closed")
)
//自定义错误数据结构
type MyError struct {
	Operation string
	Err error	// error接口有一个 Error()方法 要实现
}
//用fmt创建错误对象
if f < 0 {
    return 0, fmt.Errorf("some error happened")
}
```




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
