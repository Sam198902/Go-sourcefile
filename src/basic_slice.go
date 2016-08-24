package main

import (
	"fmt"
)

func main() {
	//不需要指定长度，就代表slice，如果是有长度就是数组

	/* 定义切片*/
	//切片的零值是nil
	var s1 []int
	fmt.Println(s1)

	//指向数组
	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := a1[9] //取一个索引为9的元素
	fmt.Println(s2)

	//取第6个到第9个元素
	//切片中，左索引是包含，右索引是不包含，即[5,9)
	s3 := a1[5:9]
	fmt.Println(s3)

	//获取后5个元素
	//1,指定下标，后面是不到10
	s4 := a1[5:10]
	fmt.Println(s4)
	//可以用长度代替
	s5 := a1[5:len(a1)]
	fmt.Println(s5)
	//可以省略后面索引方式，代表到结尾
	s6 := a1[5:]
	fmt.Println(s6)

	//获取前5个元素
	s7 := a1[:5]
	fmt.Println(s7)

	/*make*/
	//make创建切片。这种方式比较正式
	//make 3个参数，第一个是类型，第二个是个数，第三个是深度
	b1 := make([]int, 3, 10) //初始化3个int放到切片中
	fmt.Println(b1, len(b1), cap(b1))

	//切片指向底层数组
	a := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Printf("sa[%s]\n", string(sa))
	sb := a[3:5]
	fmt.Printf("sb[%s]\n", string(sb))

	/*Reslice*/
	//从切片再次获取，那么下标是相对于切片的，而不是底层数组
	//sa为cde，则取出de为，从相对于sa索引为[1,3)
	sc := sa[1:3]
	fmt.Printf("sc[%s]\n", string(sc))

	//这里是可以这样操作，这就是切片是指向底层数组的概念
	//只要不超过底层数组位数这样均是可以的
	//故输出的是 fg
	sd := sa[3:5]
	fmt.Printf("sd[%s]\n", string(sd))

	//由于是共享底层数组，当改变其中某一个值，其他均会改变
	sa[1] = 'x'
	fmt.Printf("sa[%s] sb[%s] sc[%s] \n", string(sa), string(sb), string(sc))

	fmt.Printf("sa[%d][%d] sb[%d][%d] sc[%d][%d] sd[%d][%d]\n", len(sa), cap(sa), len(sb), cap(sb), len(sc), cap(sc), len(sd), cap(sd))

	//索引是相对于被slice的数组的，不得超过底层数组的长度
	//se := sa[9:11] //或报错

	/*Append*/
	//可以在slice尾部追加元素
	c1 := make([]int, 3, 6)
	fmt.Printf("%v [%p] cap [%d]\n", c1, c1, cap(c1))
	c1 = append(c1, 1, 2, 3)
	fmt.Printf("%v [%p] cap [%d]\n", c1, c1, cap(c1))

	//当追加超过原来的容量，则会重新分配内存地址
	//容量会成倍增长
	c1 = append(c1, 4, 5, 6)
	fmt.Printf("%v [%p] cap [%d]\n", c1, c1, cap(c1))

	//当append追加超过原数组，则指向的是新数组，老数组的改变不会影像新数组

	/*
		sa = append(sa, 'y', 'y', 'y', 'y')
		fmt.Printf("sa[%s] sb[%s] sc[%s] sd[%s] \n", string(sa), string(sb), string(sc), string(sd))
		fmt.Printf("sa[%d][%d] sb[%d][%d] sc[%d][%d] sd[%d][%d]\n", len(sa), cap(sa), len(sb), cap(sb), len(sc), cap(sc), len(sd), cap(sd))
	*/

	sa = append(sa, 'z', 'z', 'z', 'z', 'z', 'z', 'z')
	sa[1] = 'M'
	fmt.Printf("sa[%s] sb[%s] sc[%s] sd[%s] \n", string(sa), string(sb), string(sc), string(sd))
	fmt.Printf("sa[%d][%d] sb[%d][%d] sc[%d][%d] sd[%d][%d]\n", len(sa), cap(sa), len(sb), cap(sb), len(sc), cap(sc), len(sd), cap(sd))

	d3 := a1[2:5]
	fmt.Printf("d3:%v len[%d],cap[%d]\n", d3, len(d3), cap(d3))
	d3 = append(d3, 11, 11, 11, 11, 11, 11, 11)
	fmt.Printf("d3:%v len[%d],cap[%d]\n", d3, len(d3), cap(d3))

	/* copy */

	d1 := []int{1, 2, 3, 4, 5, 6}
	d2 := []int{7, 8, 9, 10, 0, 0, 0, 0, 0, 0}
	//1
	//copy(d1, d2)
	//2
	//copy(d2, d1)
	//3
	//copy(d2[2:4], d1[1:3])
	//4
	//d2 = d1[:2]  //由于切片都是一个类型[]int，故可以赋值
	//5
	//d2 = d1[2:4]
	//6、整体拷贝
	d2 = d1[:]
	fmt.Println(d2)

}
