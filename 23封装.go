package main

import "fmt"

type person struct {
	Name string
	age  int
}

// NewPerson 定义一个工厂函数，用于创建 person 实例
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// GetAge 获取年龄
func (p *person) GetAge() int {
	return p.age
}

// SetAge 设置年龄
func (p *person) SetAge(age int) {
	p.age = age
}

func main() {
	p := NewPerson("丽丽")
	p.SetAge(20)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
}
