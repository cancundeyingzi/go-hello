package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(int(time.Now().Month()))
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Minute())

	fmt.Println(time.Now().Format("2006/01-02 15/04:05")) //这几个数字必须这么写,,就会显示当前时间.符号无所谓。
	fmt.Println(time.Now().Format("2006 15:04:05"))
}
