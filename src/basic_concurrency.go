package main

import (
	"fmt"
	"runtime"
	"sync" //同步
	"time"
)

func main() {
	//Goroutine
	//当直接运行，由于 main函数运行快，所以如果不sleep，main退出了，Go还没有运行

	go Go()
	//单位不一样
	time.Sleep(2 * time.Second)

	//Go引用了Channel，通道
	//相当于锁
	//创建一个没有任何缓存的channel，也可以理解为计数器
	//当前创建一个计数器为0的channel，类型为bool
	c1 := make(chan bool)
	//运行一个匿名函数
	go func() {
		fmt.Println("Chan!!!")
		//这里是在往channel中添加一个内容
		c1 <- true
	}()

	//这里是从channel中取出一个值，由于创建之初，初始值是0，则channel中没有任何内容
	//这里就出现了等待，知道内部有内容
	//当匿名Goroutine运行后，会添加一个内容，所以这里会有值
	<-c1

	c2 := make(chan bool)
	go func() {
		fmt.Println("Chan2!!!")
		c2 <- true
		//所以这里必须主动关闭
		close(c2)
	}()

	//如果使用for range也可以实现等待，当前是一直从channel中接收内容，让后输出
	//所以如果for range语句外不主动关闭channel的话，会一直去读取，然后出现死循环
	for v := range c2 {
		fmt.Println(v)
	}

	//make 创建的都是双向通道，即可存，也可取

	//这里还是创建了一个无缓存的channel
	c3 := make(chan bool)
	//运行一个匿名函数
	go func() {
		fmt.Println("Chan3!!!")
		//这里是在往channel中添加一个内容
		<-c3
	}()

	//这里也会阻塞，无缓存的channel，是没有控件预留的，所以当前存入了一个数据。必须被读取才会结束，不然一直阻塞
	c3 <- true

	//创建一个缓存为1的channel
	c4 := make(chan bool, 1)
	//运行一个匿名函数
	go func() {
		fmt.Println("Chan4!!!")
		//这里是在往channel中添加一个内容
		<-c4
	}()

	//因为有一个缓存的预留，所以这里不会阻塞，那么Chan4不会被输出，main会早于Goroutine结束
	c4 <- true

	//创建一个缓存为1的channel
	c5 := make(chan bool, 1)
	//运行一个匿名函数
	go func() {
		fmt.Println("Chan5!!!")
		//这里是在往channel中添加一个内容
		c5 <- true
	}()

	<-c5

	//所以可以理解为有缓存为异步的，无缓存为同步的

	c6 := make(chan bool)

	for i := 0; i < 10; i++ {
		go Go1(c6, i, "c6")
	}

	<-c6

	////

	//用GOMAXPROCS设置使用CPU的核数
	//NumCPU返回当前的核数
	//多核，提升效率
	runtime.GOMAXPROCS(runtime.NumCPU())

	c7 := make(chan bool)

	for i := 0; i < 10; i++ {
		go Go1(c7, i, "c7")
	}

	<-c7

	//解决没有全部运行的问题
	//1,创建一个含有10个缓存的channel
	c8 := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go Go2(c8, i, "c8")
	}

	//这里采用从channel 里读取10次的方式来保证全部运行完成，不然就等待的方式
	for i := 0; i < 10; i++ {
		<-c8
	}

	//2、这种方法是采用同步的方式
	//创建一个同步组
	wg := sync.WaitGroup{}
	//往同步组内加入10个数据，相当于信号量
	wg.Add(10)
	for i := 0; i < 10; i++ {
		//这里采用了指针传递
		go Go3(&wg, i, "wg")
	}
	//一直到10个全部执行结束，这里等待
	wg.Wait()

	//多个channel的时候，使用select
	cs1, cs2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v, ok := <-cs1:
				if !ok {
					break
				}
				fmt.Println("cs1", v)
			case v, ok := <-cs2:
				if !ok {
					break
				}
				fmt.Println("cs2", v)
			}
		}
	}()

	cs1 <- 1
	cs2 <- "hi"
	cs1 <- 3
	cs2 <- "Hello"

	close(cs1)
	close(cs2)

	//利用select作为发送方
	cs3 := make(chan int)
	go func() {
		for v := range cs3 {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case cs3 <- 0:
		case cs3 <- 1:
		}
	}

	//可以无线阻塞
	//当使用空的select就是阻塞

	//select {}

	//可以设置超时
	cs4 := make(chan bool)
	select {
	case v := <-cs4:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("TimeOut")
	}
}

func Go() {
	fmt.Println("Go Go Go !!!")
}

func Go1(c chan bool, index int, name string) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(name, index, a)

	if index == 9 {
		c <- true
	}
}

func Go2(c chan bool, index int, name string) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(name, index, a)

	c <- true
}

//采用指针传递，这样就能修改
func Go3(wg *sync.WaitGroup, index int, name string) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(name, index, a)

	wg.Done()
}
