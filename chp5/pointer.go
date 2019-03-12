package main

import "fmt"

func main() {

	var i1 = 5;
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int	//默认为nil
	intP = &i1
	*intP = 6
	fmt.Printf("An integer: %d, it's location in memory: %p\n", *intP, intP)



}