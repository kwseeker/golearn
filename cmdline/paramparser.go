package cmdline

/**
	paramparser.go 命令行参数解析器
	从命令行提取参数值到指定的struct结构体对象中

	下面实例是将命令行参数读取到用户基本信息和用户工作信息中
 */

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

//===============================================================
// 非命令集合的方式
//var (
//	levelFlag = flag.Int("level", 0, "级别")	//levelFlag 是 int 类型
//	bnFlag int
//)
//func init() {
//	flag.IntVar(&bnFlag, "bn", 3, "份数")
//}

type JobArray []string		//取别名
func (ja *JobArray) String() string {
	return fmt.Sprint([]string(*ja))
}
func (ja *JobArray) Set(value string) error {
	*ja = append(*ja, value)
	return nil			//返回 error 为 nil
}

type PersonBasicInfo struct  {	//定义一个结构体
	name 		string
	age 		int
	sex 		string
	birthDay	*time.Time
}

func (pbi *PersonBasicInfo) GetBirthDay() string {
	return pbi.birthDay.String()
}

type PersonWorkInfo struct {
	name 		string
	job 		*JobArray
}

func (pwi *PersonWorkInfo) GetWorkList() string {
	return pwi.job.String()
}


func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("/t./golearn pbi -name <Arvin> -age <25> -sex <male>")
	fmt.Println("/t./golearn pwi -name <Arvin> -job <...> [-job ...]")
}

func ParsePBIArgs() *PersonBasicInfo {
	// 定义参数标志集合	./golearn pbi -name <Arvin> -age <25> -sex <male>
	pbiFlagSet := flag.NewFlagSet("pbi", flag.ExitOnError)
	// 参数集合中参数
	name := pbiFlagSet.String("name", "", "the name of the person")
	age := pbiFlagSet.Int("age", 0, "the age of the person")
	sex := pbiFlagSet.String("sex", "", "the sex of the person")
	birthDayStr := pbiFlagSet.String("birthday", "", "the birthday of the person")
	//paramDuration := pbiFlagSet.Duration("", 2*time.Minute, "time duration param")

	err := pbiFlagSet.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
		os.Exit(21)
	}

	const shortForm = "2006-01-02"	//这里这个值必须是 2006-01-02 这一天，否则会报错
	birthDay, err := time.Parse(shortForm, *birthDayStr)
	if err != nil {
		log.Panic(err)
		os.Exit(22)
	}
	return &PersonBasicInfo{*name, *age, *sex, &birthDay}
}

func ParsePWIArgs() *PersonWorkInfo {
	// 定义可以解析的另一类参数集合 ./golearn pwi -name <Arvin> -job <...> [-job ...]
	pwiFlagSet := flag.NewFlagSet("pwi", flag.ExitOnError)
	name := pwiFlagSet.String("name", "", "the name of the person")
	jobArray := JobArray{}
	pwiFlagSet.Var(&jobArray, "job", "the jobs of the person")

	err := pwiFlagSet.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
		os.Exit(23)
	}

	return &PersonWorkInfo{*name, &jobArray}
}

// 从命令行获取参数
func ParseArgs() (pbi *PersonBasicInfo, pwi *PersonWorkInfo) {

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(24)
	}

	//argsLen := len(os.Args)
	//fmt.Println("附带了" + strconv.Itoa(argsLen) + "个参数")

	switch os.Args[1] {
	case "pbi":
		pbi = ParsePBIArgs()
	case "pwi":
		pwi = ParsePWIArgs()
	default:
		printUsage()
		os.Exit(25)
	}

	//fmt.Println("命令行：", os.Args[0:])
	return		//相当于 return pbi, pwi
}
