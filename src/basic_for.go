package main

import (
	"fmt"
)

func main() {

	/*for*/
	a := 1
	//无限循环
	for {
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a)

	//有条件
	a = 1
	for a <= 3 {
		a++
	}
	fmt.Println(a)

	//上述中两种均为while的变种
	//
	a = 1
	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)

	//for中条件表达式尽量具体变量，最好不要是表达式，因为for条件表达式，每次都会判断
	b := "sss"
	a = 1
	nlen := len(b)
	for i := 0; i < nlen; i++ {
		a++
	}
	fmt.Println(a)
}
