package main

import (
	"fmt"
)

//接口是一个或多个方法签名的集合

type USB interface {
	Name() string
	Connect
}

//可以和结构体类似，进行嵌套
type Connect interface {
	Connect()
}

/*只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示
声明实现了哪个接口，这称为 Structural Typing
*/

//以下PhoneConnect实现了接口USB的所有方法签名，即实现了USB接口
//接口类似父类
type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string {
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("Connect:", pc.name)
}

type PhoneConnect1 struct {
	name string
}

func (pc PhoneConnect1) Name() string {
	return pc.name
}

func (pc PhoneConnect1) Connect() {
	fmt.Println("Connect:", pc.name)
}

//以下是空接口，根据原理，只要是实现了接口中所有方法签名，就实现了该接口，那么对于空接口来说，所有类型均实现了该接口
//所以任何类型均有空接口
type USB_MSG interface {
}

func main() {

	//于是此处可以，定义一个USB类型变量
	var a USB
	//因为PhoneConnect类型实现了接口，故可以直接赋值
	a = PhoneConnect{" A PhoneConnecter"}
	//此处可以调用该对象的方法
	a.Connect()
	Disconnect(a)

	b := PhoneConnect{"B PhoneConnecter"}
	b.Connect()
	//上面只能说明实现了该类型的方法
	//下述是否是可以算 USB类型呢
	b.name = "OK"
	b.Connect()
	Disconnect1(b)

	c := PhoneConnect1{"C PhoneConnecter"}
	c.Connect()
	Disconnect1(c)

	d := PhoneConnect1{"d PhoneConnecter"}
	Disconnect2(d)

	type TC int
	var e TC
	Disconnect2(e)

	//接口也可以进行强制类型转换
	//但是原则是高层可转换为底层，底层不可转换为高层
	//也可以这样理解，可以向下转换
	a1 := PhoneConnect{" X PhoneConnecter"}
	var a2 Connect
	a2 = Connect(a1)
	a2.Connect()

	//a2拿到的只是a1的拷贝
	a1.name = "OK"
	a2.Connect()

	var b1 interface{}
	//空接口是nil
	fmt.Println(b1 == nil)
	//只有当接口存储的类型和对象都为nil时，接口才等于nil
	var p1 *int = nil
	b1 = p1
	//这里存储的是一个nil指针，所有存储的类型不为nil
	fmt.Println(b1 == nil)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnent")
}

func Disconnect1(usb USB) {
	//此处是判断是否具有该类型
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("Disconnent1 PhoneConnect", pc.name)
		return
	} else if pc, ok := usb.(PhoneConnect1); ok {
		fmt.Println("Disconnent1 PhoneConnect1", pc.name)
		return
	}

	fmt.Println("Unknown decive")

}

//此处是空接口，根据上面的理解，所以这个方法是可以传入任何类型的
func Disconnect2(usb interface{}) {
	//此处是判断是否具有该类型

	switch v := usb.(type) {
	case PhoneConnect:
		fmt.Println("Disconnent2 PhoneConnect", v.name)
	case PhoneConnect1:
		fmt.Println("Disconnent2 PhoneConnect1", v.name)
	default:
		fmt.Println("Unknown2 decive")
	}

}
