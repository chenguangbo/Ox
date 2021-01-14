package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "我爱你中国!"
	fmt.Println("str len = ", len(str))

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println("字符=  : ", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("1231a2")
	if err != nil {
		fmt.Println("转换错误: ", err, n)
	} else {
		fmt.Println("转换正常: ", err, n)
	}
	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T", str, str)

	// 字符串 转 []byte:
	var bytes = []byte("hello go")
	fmt.Println("bytes=%v\n", bytes)
	// []byte 转 字符串
	strc := string(bytes)
	fmt.Println("byte数组转字符串 : 2", strc)

}
