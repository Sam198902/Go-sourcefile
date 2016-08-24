package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.Format(time.ANSIC))

	//最好不要这样使用，建议用上面的常量
	//fmt.Println(t.Format("Mon Jan _2 15:04:06 2006"))

	fmt.Println(t.Format(time.Kitchen))
}
