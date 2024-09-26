package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayHello() { //这里是又复制了一个p。所以这边改动了不影响main函数里的
	p.Name = "露露"
	fmt.Println(p.Name)
}

func (p *Person) SayHello2() { //这里用了指针。传递内存，所以会直接修改main函数里的p。
	(*p).Name = "露露"
}
func main() {
	var p Person
	p.Name = "丽丽"
	p.SayHello()
	fmt.Println(p.Name)
	(&p).SayHello2()
	fmt.Println(p.Name)

}
