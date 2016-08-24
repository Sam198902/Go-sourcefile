package main

import (
	"fmt"
	"os"
)

func main() {

	a := true
	if a {
		fmt.Println("a is true")
	}

	//支持初始化，但是作用域是在语句块之内
	//如果外部是定义了同名的变量，语句块中仍然使用的是语句块内的变量，语句块之前的变量会被隐藏
	//左大括号必须放在与if同一行
	if a, b, c := 1, 2, "test"; a+b > 2 {
		fmt.Println(c)
	} else {
		fmt.Println("小于2")
		fmt.Println(a)
	}

	fmt.Println(a)

	var s = "Go并发编程"

	fmt.Println(len(s))

	_, err := os.Open("xxx")
	if err != nil {
		fmt.Println(err)
		return
	}

}
