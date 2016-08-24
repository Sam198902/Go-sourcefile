package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	Name string
}

type TC int

func main() {
	//调用方法
	a := A{}
	a.Print()

	b := B{}
	b.Print()
	fmt.Println(b.Name) //由于是值传递，所以并不会改变

	c := C{}
	c.Print()
	fmt.Println(c.Name) //采用指针传递，所以会改变值

	//牛逼的地方
	//Go可以为任何一种类型定义方法
	var d TC
	fmt.Printf("d = %d\n", d)
	d.Print() //Method Value
	fmt.Printf("d = %d\n", d)
	(*TC).Print(&d) //Method Expression

	d.Increase(100)
	fmt.Printf("d = %d\n", d)

}

//与A的struct绑定,函数名前有一个receiver 接收者
//这样A就结构体和此方法进行了绑定
func (a A) Print() {
	fmt.Println("A")
}

//Go 语言不能进行方法的重载，所以下述不能使用
/*func (a A) Print(b int) {
	fmt.Println("A")
}*/

//由于接收者不一样，所以此处虽然方法名是一样，但是并不是重载，是不同的方法
//接收者不一样
func (b B) Print() {
	b.Name = "BB" //此处是值传递
	fmt.Println("B")
}

//此处是指针传递
func (c *C) Print() {
	c.Name = "CC"
	fmt.Println("C")
}

func (a *TC) Print() {
	*a = 10
	fmt.Println("TC")
}

func (a *TC) Increase(num int) {
	//虽然底层都是int，但是type之后，类型是不同的
	*a += TC(num)
	fmt.Println("TC num")
}

//方法可以访问私有字段
