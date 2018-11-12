package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
	文件操作
	1）文件是否存在
	2）文件创建
	3）文件读取
	4）文件写入
	5）文件删除
 */

//判断文件或文件夹是否存在
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取执行文件所在的目录
func GetCurrentPath() (currentPath string, err error) {
	currentPath, err = filepath.Abs("")
	return
}

//标准库 io/ioutil os bufio 均有提供对文件读写的方法
//func WriteFile(path string) {
//
//}

//读取文件内容返回string类结果
// path可为可执行文件的相对路径或者绝对路径
func ReadContentToString(path string) string {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	content, err := ioutil.ReadAll(fp)
	return string(content)
}
