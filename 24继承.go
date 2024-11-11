package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (an *Animal) Shout() {
	fmt.Println("喊叫")
}
func (an *Animal) ShowInfo() {
	fmt.Println("年龄体重", an.Age, an.Weight)
}

type Cat struct {
	Animal
	food string
}

func (c *Cat) scratch() {
	fmt.Println("挠人")
}

func main() {
	cat := Cat{}
	cat.Age = 3
	cat.Animal.Weight = 10.6 //似乎两种方法都行
	cat.food = "鱼"
	cat.Shout()
	cat.ShowInfo()
}
