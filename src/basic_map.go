package main

import (
	"fmt"
	"sort"
)

func main() {
	/*map 定义 */
	//1
	var m1 map[int]string
	m1 = map[int]string{}
	fmt.Println(m1)

	//2,右边必须是map类型
	var m2 map[int]string = map[int]string{}
	fmt.Println(m2)

	//3，
	m3 := map[int]string{}
	fmt.Println(m3)

	//4、make
	m4 := make(map[int]string)
	fmt.Println(m4)

	//赋值
	a1 := make(map[int]string)
	a1[1] = "ok"
	a1[10] = "No"
	//输出所有map
	fmt.Println(a1)
	//根据key得到value
	fmt.Println(a1[1], a1[10])
	//如果去取一个不存在的key，也不会报错,会是对应类型的零值
	fmt.Println(a1[2])

	//如下使用请注意，会去到零
	a2 := make(map[int]int)
	a2[1] = 10
	fmt.Println(a2)
	fmt.Println(a2[10])
	fmt.Println(len(a2))

	//复杂的map
	var b1 map[int]map[int]string
	b1 = make(map[int]map[int]string)
	//由于只是初始化了第一个map，所以直接赋值是会报错的
	//b1[1][1] = "ok"
	b1[1] = make(map[int]string)
	b1[1][1] = "ok"
	//如下也是不行的，因为只是初始化了 b1[1]的map，而没有初始化b1[2]的map
	//b1[2][1] = "No"
	fmt.Println(b1)

	//如果确定是否被初始化了，可以利用多返回值
	a, ok := b1[2][1]
	fmt.Println(a, ok)

	//可以利用if
	if !ok {
		b1[2] = make(map[int]string)
		b1[2][1] = "No"
	}

	a, ok = b1[2][1]
	fmt.Println(a, ok)

	/*迭代操作*/

	//slice
	//v是一个值拷贝，所以对于v的修改不会修改值本身
	s1 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	for i, v := range s1 {
		fmt.Printf("[%d %s] ", i, string(v))
	}
	fmt.Println()

	//如需要改变某个值，可以使用下标i
	for i, _ := range s1 {
		s1[i] += 1
	}
	fmt.Println(string(s1))

	//map
	sm := make([]map[int]string, 5)
	//下面就可以看出，当前v只是map的拷贝
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)

	//直接对象进行操作，利用k
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][i] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	//map
	b2 := map[string]string{"日": "Sunday", "一": "Monday", "二": "Tuesday", "三": "Wednesday", "四": "Thursday", "五": "Friday", "六": "Saturday"}
	for i, v := range b2 {
		fmt.Printf("星期[%s] 英文[%s]\n", i, v)
	}
	//上述的例子非常好的体现了map 是无序的

	b3 := map[int]string{0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday"}
	s3 := make([]int, len(b3))
	i := 0
	for k, _ := range b3 {
		s3[i] = k
		i++
	}

	sort.Ints(s3)
	fmt.Println(s3)
	nlen := len(s3)
	for i := 0; i < nlen; i++ {
		fmt.Printf("%s ", b3[s3[i]])
	}
	fmt.Println()

	//将数序颠倒
	b4 := make(map[string]int, len(b3))
	for i, v := range b3 {
		b4[v] = i
	}

	fmt.Println(b4)

}
