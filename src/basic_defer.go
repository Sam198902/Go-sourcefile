package main

import (
	"fmt"
)

func main() {

	//defer 是逆序调用
	/*	fmt.Println("a")
		defer fmt.Println("b")
		defer fmt.Println("c")

		for i := 0; i < 3; i++ {
			defer fmt.Println(i)
		}

		//利用了闭包
		//这里是使用匿名函数，i是一个引用拷贝，则i保留的一直是地址
		for i := 10; i < 13; i++ {
			defer func() {
				fmt.Println(i)
			}()
		}

		for i := 20; i < 23; i++ {
			defer func(a int) {
				fmt.Println(a)
			}(i)
		}

		B()
	*/
	var fs = [4]func(){}

	//匿名函数中变量i本身函数是没有定义的，所以它是从外部得到的
	//得到的是一个引用
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

func A() {
	defer fmt.Println("defer A ")
	i := 10
	j := 0
	a := i / j
	fmt.Println(a)

}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	i := 10
	j := 0
	a := i / j
	fmt.Println(a)
}
