package main

import (
	"fmt"
)

//运算符

/*
6: 0110
11:1011
*/

func main() {
	fmt.Println(1 << 10 << 10 >> 10)

	//& 与
	fmt.Println(6 & 11)

	//|  或
	fmt.Println(6 | 11)

	//^  异或
	fmt.Println(6 ^ 11)

	//&^  按位清除
	fmt.Println(6 &^ 11)

	/*指针*/
	a := 12
	var p *int = &a
	fmt.Println(p, *p)

}
