package main

import "fmt"

func main() {
	sum(10, 20)
}

//   当go 执行到一个defer时,不会立即执行defer后的语句,而是将defer后的语句亚茹到一个栈中
//   当函数执行完毕后,在从defer栈中,依次从栈顶取出语句执行   (遵守栈,先入后出的机制)
func sum(n1 int, n2 int) int {

	defer fmt.Println("n1 : ", n1)
	defer fmt.Println("n2 : ", n2)
	res := n1 + n2
	fmt.Println("res : ", res)
	return res
}
