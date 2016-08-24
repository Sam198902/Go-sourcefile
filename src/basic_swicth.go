package main

import (
	"fmt"
)

func main() {
	/* swtich */
	//与之前的有很大不同

	//可以使用任何类型和表达式作为条件语句
	//case 语句不需要break才能跳出循环，只要执行结束就会跳出
	c := "test"
	switch c {
	case "test":
		fmt.Println("find test")
	case "release":
		fmt.Println("find release")
	default:
		fmt.Println("default")
	}

	//如果需要执行下一个case，需要使用fallthrough
	switch c {
	case "test":
		fmt.Println("find test")
		fallthrough
	case "release":
		fmt.Println("find release")
		fallthrough
	default:
		fmt.Println("default")
	}

	//支持初始化表达式，右侧必须要一个分号
	a := true
	switch a := 1; {
	case a > 1:
		fmt.Println("a>1")
	case a == 1:
		fmt.Println("a=1")
	default:
		fmt.Println("a<1")
	}

	fmt.Println(a)

}
