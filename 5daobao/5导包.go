package main

import (
	"fmt"
	"hello/5daobao/mulu"
)

func main() {
	mulu.Bao()          //目录软件包单文件都行
	fmt.Println(mulu.A) //目录软件包单文件都行
	Bao()               //用目录/软件包形式构建运行可以,不能单文件
	fmt.Println(A)      //用目录/软件包形式构建运行可以,不能单文件
}
func init() { //此函数优先运行
	fmt.Println("init")
}
