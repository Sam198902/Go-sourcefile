package main

import (
	"fmt"
)

func Pingpong(s []int) {
	s = append(s, 4)
}

func Pingpong1(s []int) []int {
	s = append(s, 4)
	return s
}

func main() {

	//创建了一个个数为0，深度为0的切片
	s1 := make([]int, 0)
	fmt.Println(s1)
	//由于此处传入的切片，虽然切片是引用传递，但是由于当前切片为0，则
	//在函数里又append了，而个数不满足，则会新分配一个地址
	//那么函数里的s指向不再是传入的s指向了，底层地址改变了
	Pingpong(s1)
	//所以这里也是空
	fmt.Println(s1)

	s1 = Pingpong1(s1)
	fmt.Println(s1)

	s2 := []string{"a", "b", "c"}

	//这里面的匿名函数中v是外面的，所以获得的是引用
	//当goroutine运行是，for已经基本运行完了
	//所以当前v是“C"
	//如果要是正确的，那么建议匿名函数要加参数
	for _, v := range s2 {
		go func() {
			fmt.Println(v)
		}()
	}

	select {}

}
