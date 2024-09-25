package main //同一个目录所有的package名字要一样

import "fmt"

var A = 555 //变量首字母大写是公有,小写私有

func Bao() { //函数名首字母大写是公有,小写私有
	fmt.Println("外面的")
	return
}
