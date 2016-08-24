package main

import (
	"fmt"
)

//1、无参数无返回
//func A() {}

//2、有参数，有返回
//func A(a int, b string) (int, string) {}

//3、相同类型参数
//func A(a, b, c int) string {}
//func A() (a, b, c int)     {}
//func A() (int, int, int)   {}

//不定长变参
func A(a ...int) {
	fmt.Println(a)
}

func A1(a ...int) {
	a[0] = 5
	a[1] = 6
	a[2] = 7
	a[3] = 8
}

func A2(a []int) {
	a[0] = 5
	a[1] = 6
	a[2] = 7
	a[3] = 8
}
func A3(a *int) {
	*a = 10
}

func B() {
	fmt.Println("Func B")
}

func main() {
	A(1, 2, 3)
	A(1, 2, 4, 4, 5)

	//值拷贝
	a1 := [4]int{1, 2, 3, 4}
	A1(a1[0], a1[1], a1[2], a1[3])
	fmt.Println(a1)

	//引用拷贝，也是地址拷贝
	s1 := a1[:]
	A2(s1)
	fmt.Println(a1)

	//指针操作
	num := 1
	A3(&num)
	fmt.Println(num)

	//函数也是类型
	a2 := B
	a2()

	//匿名函数
	a3 := func() {
		fmt.Println("fucn anonymous")
	}
	a3()

	//闭包

	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}

//闭包
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
