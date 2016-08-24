package main

import (
	"fmt"
)

const (
	PI     = 3.14
	const1 = "1"
	const2 = 2
	const3 = 3
)

const a = 1

//const b = 'a'
//const a, b, c = 'a', 1, 2

//由于程序开始前，ssss长度是未知的，所有此处用len获取长度，也是未知，故不是常量
//string ssss = "test"
//const c2 = len(ssss)

//如果出现如下方式定义常量，则下属与上一个格式相同的常量延用上一个常量的值
//此处就是初始化规则，如果出现不一致会报错
const (
	a1, a2 = "a", 'a'
	a3, a4
	//c5
)

//因为a本身是常量，所有如下操作是可取的
const (
	b1 = a + 1
	b2 = a * a
	//此处由于传c2是确定的常量，所有len获取也可以是常量，当然常量表达式中的函数必须是内置函数，len就是
	b3 = "test"
	b4 = len(b3)
	b5
)

//如下，iota代表为枚举类型，当出现iota则，当前值为，之前定义的常量个数，从0开始计算
//下面，a输出65,b3也会输出65，因为b3延续上一个常量
//c3为枚举，因为前面已经定义了两个常量，则从0开始，c3为2，d3也为枚举，则为 3
const (
	c1 = 'A'
	c2
	c3 = iota
	c4
)

//上述定义其实均不规范，对于一个真正使用的常量必须遵循编码规范
//必须均为大写，便于阅读的调用

//星期
const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

//
const (
	MAX_COUNT = 512
	//如果不想被外部调用的话，可以
	_MAX_COUNT = 512
	nMAX_COUNT = 512
)
const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
)

func main() {

	fmt.Println(a1, a2, a3, a4)
	fmt.Println(b1, b2, b3, b4, b5)
	fmt.Println(c1, c2, c3, c4)

	fmt.Println(KB)

}
