package main

import (
	"fmt"
)

//定义一个最简单的结构体
//空结构体
type test struct{}

//含有数据的结构体
type person struct {
	Name string
	Age  int
}

//复合型结构
type human struct {
	Sex  int
	Name string
	Age  int
}

//包含了上一个结构，直接使用结构体名即可
type student struct {
	human
	Class string
}

type teacher struct {
	human
	Course string
}

type student1 struct {
	h1   human
	year int
}

//同名字段
type student2 struct {
	human
	Name string
}

func main() {
	a1 := test{}
	fmt.Println(a1)

	a2 := person{}
	fmt.Println(a2)

	//给结构体初始化
	a2.Name = "jake"
	a2.Age = 19
	fmt.Println(a2)

	//通过字面值初始化
	a3 := person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a3)

	//值拷贝
	A(a3)
	fmt.Println(a3)

	//地址拷贝
	B(&a3)
	fmt.Println(a3)

	//定义是直接定义为指针，建议这样操作
	a4 := &person{
		Name: "lilei",
		Age:  20,
	}
	//
	//*a4.Name = "sam"
	a4.Name = "sam"
	fmt.Println(a4)

	B(a4)
	fmt.Println(a4)

	//匿名结构，定义及初始化
	b1 := &struct {
		Name string
		City string
	}{
		Name: "lilei",
		City: "ShenZhen",
	}

	fmt.Println(b1)

	//匿名结构嵌套
	type msg struct {
		Name    string
		Age     int
		Contact struct {
			Phone, City string
		}
	}

	//初始化
	b2 := msg{Name: "Sam", Age: 21}
	b2.Contact.Phone = "18111111111"
	b2.Contact.City = "ShenZhen"
	fmt.Println(b2)

	//匿名字段
	type msg1 struct {
		string
		int
	}
	//匿名字段，赋值顺序必须正确
	b3 := msg1{"Joe", 22}
	fmt.Println(b3)

	//可以赋值和比较,必须是类型相同
	fmt.Println(a2 == a3)

	//复合类型，初始化时，复合的结构体初始化，直接使用结构体名
	c1 := teacher{Course: "English", human: human{Name: "Jim", Age: 20, Sex: 1}}
	//但是赋值可以直接使用
	c1.Name = "Joe"
	fmt.Println(c1)

	//可以直接使用
	c2 := student{}
	c2.Name = "Cathy"
	c2.Age = 18
	c2.Class = "302"
	c2.Sex = 0
	fmt.Println(c2)

	c3 := student1{year: 2016, h1: human{Name: "Jim", Age: 20, Sex: 1}}
	fmt.Println(c3)

	c4 := student2{Name: "Jim1", human: human{Name: "Jim2", Age: 20, Sex: 1}}
	fmt.Println(c4)
	c4.Name = "Jim3"
	c4.human.Name = "Jim4" //同名可以使用结构体名再选择操作
	fmt.Println(c4)

}

func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 13
	fmt.Println("B", per)
}
