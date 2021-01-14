package main

import (
	"fmt"
)

//a := func (n1 int, n2 int) int {
//	return a1 - a2
//}

func main() {
	//aa = a(1, 2)

	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 = ", res1)

	a := func(a1 int, a2 int) int {
		return a1 + a2
	}

	b := a(1000, 2000)
	fmt.Println("b :", b)
	c := a(9, 1)
	fmt.Println("C : ", c)

}
