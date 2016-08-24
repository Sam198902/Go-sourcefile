package main

import (
	"fmt"
)

func main() {

	/* break */
	a := 1
	for i := 0; i < 10; i++ {
		if i > 3 {
			a = i
			break
		}
	}
	fmt.Println(a)

	//此处LABEL1标签和第一个for是同一个级别
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//当只使用break时，只会跳出当前第二层for，还是死循环
				//break
				//如果加上标签，则会跳到标签，调出标签同一级别的循环
				break LABEL1

			}
		}
	}

	fmt.Println("Break OK")

	/*goto*/
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//当只使用break时，只会跳出当前第二层for，还是死循环
				//break
				//如果加上标签，则会跳到标签，调出标签同一级别的循环
				goto LABEL2

			}
		}
	}
LABEL2:
	fmt.Println("goto OK")

	/*continue*/
	a = 1
	for i := 0; i < 10; i++ {
		if i > 3 {
			a = i
			continue
		}
	}
	fmt.Println(a)

LABEL3:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL3
		}
	}

	fmt.Println("continue OK")

	//goto是调整执行位置
LABEL4:
	for i := 0; i < 10; i++ {
		for {
			goto LABEL4
			fmt.Println(i)
		}
	}

	fmt.Println("continue OK")
}
