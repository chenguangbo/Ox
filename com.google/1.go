package main

import (
	"fmt"
	"unicode/utf8"
)
import _ "unicode/utf8"

func main() {

	//var days1 int = 97
	//fmt.Println(days1)

	//aa()
	//bb()
	//cc()
	//var arr [10]int
	//dd(arr, aa)

	//ff()
	//gg()

	//a1()

	//a4()

}

func a4() {
	sum := 0.0
	fmt.Println()
	fmt.Println(sum)
}

func a3() {
	s := "foobar"
	a := []rune(s)
	fmt.Println(a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("%s\n", string(a))
}

func a2() {
	str := "hello"
	b := []byte(str)
	fmt.Println(str)
	fmt.Println(b)

	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))

	fmt.Println("123--: ", string(b))
}

func a1() {
	i := 0
Loop:
	fmt.Printf("%d\n", i)
	if i < 1000000 {
		i++
		goto Loop
	}
}

func aa() {
	print("asdanda")
}

func bb() {
	var arr [10]int
	arr[0] = 12
	arr[2] = 30
	fmt.Println(arr)
}

func cc() {

	var aaa int = 10
	aaa = 20
	fmt.Println("cc : ", aaa)

}

func dd(a [10]int, b func()) {
	fmt.Println(a)
	b()
	return
}

func ff() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i)
	}
}

func gg() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Println(monthdays)

	fmt.Println(monthdays["Jan"])
	fmt.Println(monthdays["Nov"])
	var value int
	var present bool
	value, present = monthdays["Nov"]
	fmt.Println("value : ", value)
	fmt.Println("present : ", present)
}
