package collections

import "fmt"

// 哈希表
// GO 中一个map就是一个哈希表的引用， 可以写为 map[K]V, K和V分别对应key和value
// map中的元素并不是一个变量，并不能对map的元素进行取址操作
type MyMap struct {
	MapVar 	map[string]int
}

//func (mymap *MyMap) Create() {
//	//ages := make(map[string]int)
//	//agsg := map[string]int{}
//	ages := map[string]int {
//		"Arvin": 25,
//		"Bob": 26,
//		"Cindy": 25,
//		"David": 28,
//		"Eric":23,
//	}
//}

func (mymap *MyMap) Add(key string, value int) {
	mymap.MapVar[key] = value
}

func (mymap *MyMap) Delete(key string) {
	delete(mymap.MapVar, key)
}

func (mymap *MyMap) Update(key string, value int) {
	if _, isPresent := mymap.MapVar[key]; isPresent {
		mymap.MapVar[key] = value
	} else {
		fmt.Println("map does not contain ", key)
	}
}

func (mymap *MyMap) Get(key string) int {
	return mymap.MapVar[key]
}

func (mymap *MyMap) String() string {
	var str string
	for key, value := range mymap.MapVar {
		str += fmt.Sprint(key, ":", value, "\n")
	}
	return str
}
