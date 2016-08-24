//当前程序的包名
package main

//导入其它的包
//import "fmt"
//import std "fmt"  别名
//import . "fmt"  不建议。省略调用
import (
	"fmt"
	_ "math"
	//	"io"
	//	"os"
	//	"strings"
	//	"time"
)

//常量定义
const PI = 3.14

const (
	AI     = 2.14
	const1 = "1"
	const2 = 2
	const3 = 3
)

const a, b, c = 'a', 1, 2

//全局变量的声明与赋值
var name = "gopher"

//变量组的使用只能在函数体外声明与赋值
//函数体内还有另一种方式
var (
	name1 = "go"
	name2 = "1"
	name3 = 2
	name4 = 3
)

//一般类型声明
type newType int

//结构体的声明
type gopher struct{}

//接口声明
type golang interface{}

//由main函数作为函数的入口点启动
func main() {
	fmt.Println("Hello World! 您好，深圳！")
}
