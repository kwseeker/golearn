package collections

import (
	"fmt"
	"testing"
)

func TestMapUtil(t *testing.T) {
	//maptem := map[string]int{}
	maptem := make(map[string]int)
	mymap := &MyMap{maptem}
	mymap.Add("A", 1)
	mymap.Add("B", 2)
	fmt.Println(mymap.String())
	mymap.Add("C", 3)
	mymap.Update("B", 22)
	mymap.Add("D", 4)
	mymap.Delete("C")
	fmt.Println(mymap.String())
}
