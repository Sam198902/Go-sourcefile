package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ",My name is ", u.Name)
}

func (u User) A(age int, name string) {
	fmt.Println("Hello", name, ",My name is ", u.Name, ",I'm ", age)
}

func main() {
	u := User{1, "ok", 12}
	u.Hello("Cathy")
	Info(u)

	type TC int
	var a1 TC
	Info(a1)

	a2 := Manager{User: User{2, "No", 21}, title: "YYY"}
	//获取类型
	t1 := reflect.TypeOf(a2)
	//获取索引为0的字段信息
	fmt.Printf("%#v\n", t1.Field(0))
	//获取索引为0的字段信息
	fmt.Printf("%#v\n", t1.Field(1))

	//这是获取匿名字段中某个具体字段
	//FieldByIndex中传入slice
	//第一个是相对于上层的索引，User是0，第二个是相对于User内的索引
	//故查看的是Id字段的信息
	fmt.Printf("%#v\n", t1.FieldByIndex([]int{0, 0}))

	/*通过放射修改类型的值*/

	//一般类型如何修改值
	x := 123
	//获取类型，这里传指针，这样就是指针的操作
	v := reflect.ValueOf(&x)
	//取得value Elem方法
	//如果不是指针传递，则这里SetInt改变不了x本身值的
	v.Elem().SetInt(999)
	//将x的值修改为999
	fmt.Println(x)

	//修改结构体
	a3 := User{100, "Jim", 30}
	Set(&a3)
	fmt.Println(a3)

	//如何利用放射调用方法
	//1、获取类型
	v3 := reflect.ValueOf(a3)
	//2、获取方法名
	m3 := v3.MethodByName("Hello")
	//此时传送参数不再是单一的类型，而是一个类型集合
	//所以需要用一个slice
	//注意写法
	args1 := []reflect.Value{reflect.ValueOf("Joe")}
	//3、调用
	m3.Call(args1)

	//利用放射动态调用
	m4 := v3.MethodByName("A")
	args2 := []reflect.Value{reflect.ValueOf(15), reflect.ValueOf("Joe")}
	m4.Call(args2)
}

//如何获取某一个对象字信息和类型信息
//这个最好是用于struct
func Info(o interface{}) {
	//判断struct类型
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//如果不是struct，返回
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("XX")
		return
	}

	//获得struct类型值
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	//反射出struct结构内容
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	//获取struct结构的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {

	//获得类型
	v := reflect.ValueOf(o)
	//kind 是获取类型
	//Elem获取值
	//CanSet是否可以被修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("ZZZ")
		return
	} else {
		v = v.Elem()
	}

	//获取对象
	//FieldByName 寻找对象的某个字段

	f := v.FieldByName("Name")
	//这里判断当前是否找到了这个字段
	//也可以理解为这个变量是否有效
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	//再判断类型是否是String
	if f.Kind() == reflect.String {
		f.SetString("Sam")
	}
}
