package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	mBool   bool
	mstring string
)

//下面全是数值类型
var (
	mByte       byte //int8的别名，主要是为了读代码的非常清晰
	mrun        rune //int32的别名
	mint        int
	muint       uint
	mint8       int8
	muint8      uint8
	muint16     uint16
	mint16      int16
	mint32      int32
	muint32     uint32
	mint64      int64
	muint64     uint64
	mfloat32    float32
	mfloat64    float64
	mcomplex    complex64
	mcomplex128 complex128
)

type (
	newType1 int
	type1    float32
	type2    string
	type3    byte
	//byte  int8 最好不要

	文本 string //utf-8的用途，说明可以做到，但是最好不要用
)

func main() {

	//零值的测试
	var a int
	var b string
	var c bool
	var d [2]byte
	var e []int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//如果得到当前计算机某个类型具体的范围
	//使用math包中，Max和Min相关的变量
	fmt.Println(math.MaxInt32)

	/*UTF-8支持*/
	var str 文本 = "Gopher"
	fmt.Println(str)

	/*变量的声明与赋值*/

	//1、声明与赋值分开
	var num1 int
	//num1 = 1.1  这里不会自动隐式转换，会报错
	num1 = 1
	fmt.Println(num1)

	//2、声明与赋值一起
	var num2 int = 2
	fmt.Println(num2)

	//3、自动类型推断，去掉类型
	var num3 = 3
	fmt.Println(num3)

	//4、去掉var 变量说明符,最简洁,全局变量不可以这样声明和赋值
	//Go语言很牛的方法，后续函数返回值使用最多
	num4 := 4
	fmt.Println(num4)

	/*多个变量声明与赋值*/
	//1、并行
	var a1, a2, a3 int = 5, 6, 7
	fmt.Println(a1, a2, a3)

	//2、省略var和类型
	b1, b2, b3 := 8, 9, 10
	fmt.Println(b1, b2, b3)

	//3、忽略赋值
	//直接忽略了第二个赋值，这个也是一个很棒的功能，未来常用于函数返回值
	c1, _, c3, c4 := 11, 12, 13, 14
	fmt.Println(c1, c3, c4)

	/*类型转换*/
	//Go语言不存在隐式转换，所有转换必须显式的进行
	//且转换必须是在两者兼容的类型之间
	var d1 float32 = 100.1
	d2 := int(d1)
	//常见错误
	//d2 := (int)d1
	fmt.Println(d1, d2)

	//在Go语言中bool类型只有false和true，所以不能和int之间类型转换
	//var d3 bool = false
	//d4 := int(d3)

	//在Go语言中，当强制转换为字符串的时候，会先将本地数字变成二进制，然后转换成字符
	var d5 int = 65
	d6 := string(d5)
	fmt.Println(d6)

	//如果确确实实需要65的话，则需要使用strconv库
	d7 := strconv.Itoa(d5)
	fmt.Println(d7)
	d5, _ = strconv.Atoi(d7)
	fmt.Println(d5)

}
