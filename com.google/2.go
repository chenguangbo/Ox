package main

import (
	"fmt"
	"time"
)

var (
	fun1 = func(a1 int, n2 int) int {
		return n2 + a1
	}
)

func main() {
	var name string
	var age int
	fmt.Println("程序开始")
	fmt.Scan(&name)
	fmt.Println("输入年龄:")
	fmt.Scanln(&age)
	fmt.Println(name)
	fmt.Println(&age)
	fmt.Println(time.July)

	if 0 > -1 {
		fmt.Println("表达是成立")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("2021年1月11日19:17:04 : ", i)
	}
	var i int64
	for {
		i++
		fmt.Println(i)
	}

}
