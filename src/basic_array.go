package main

import (
	"fmt"
)

func main() {
	var a1 [2]int
	var a2 [1]int
	var a3 [2]string
	fmt.Println(a1, a2, a3)
	//上述是两个不同的类型，数组中长度是类型的一部分

	//声明并赋值
	b1 := [2]int{1, 1}
	fmt.Println(b1)

	//可以利用索引，索引是从0开始的
	b2 := [10]int{9: 1} //前面默认使用初始化零值
	fmt.Println(b2)

	//可以不指定长度，Go字段推断
	b3 := [...]int{1, 2, 3, 4}
	fmt.Println(b3)
	b4 := [...]int{1: 2, 3: 2, 8: 3, 2: 4} //可以无序指定
	fmt.Println(b4)
	b5 := [...]int{19: 1} //长度无20的数组，索引为19的位值是1
	fmt.Println(b5)

	//指向数组的指针
	c1 := [...]int{5: 8}
	var p *[6]int = &c1
	fmt.Println(p)
	fmt.Println(*p)
	//fmt.Println(*(p + 4))

	//指针数组
	x, y := 1, 2
	c2 := [...]*int{&x, &y}
	fmt.Println(c2)

	//数组比较，只能使用 == 和!= ，不能使用 > 和 <
	//比较必须是同类型
	fmt.Println(a1 == b1)

	//使用new关键字创建，返回的也是数组的指针
	p1 := new([10]int)
	fmt.Println(p1)

	//对于赋值，于数组本身或者指向数组的指针，均可以使用下标来操作
	a1[1] = 2
	p1[2] = 10
	fmt.Println(a1, p1)

	//多维数组,自右往左看
	c3 := [2][3]int{
		{1, 1, 1},
		{2, 3, 4},
	}
	fmt.Println(c3)

	c4 := [2][3]int{
		{1: 10},
		{2: 9},
	}
	fmt.Println(c4)

	//非顶级的数组不能使用系统自动推断的方式
	//为了避免混淆，最好还是确定
	c5 := [...][3]int{
		{1: 10},
		{2: 9},
	}
	fmt.Println(c5)

	//冒泡排序
	d := [...]int{5, 22, 14, 1, 3123, 0, -1, 342}
	fmt.Printf("排序前: %v\n", d)

	num := len(d)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if d[i] > d[j] {
				temp := d[i]
				d[i] = d[j]
				d[j] = temp
			}
		}
	}

	fmt.Printf("排序后: %v\n", d)
}
