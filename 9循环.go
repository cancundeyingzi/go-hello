package main

import "fmt"

// break可以有特殊用法"标签"https://www.bilibili.com/video/BV1bN4y1Z7BT?p=53&vd_source=9f46f1783a12a11bff2495237183ab9b
// continue跳出本次循环,同样能用"标签"
func main() {
	//两个方法,一个类似while
	var sum = 0
	for sum < 10 {
		sum++
		fmt.Println(sum) //输出1-10
	}
	//另一个就是普通for
	sum = 0
	for i := 1; i <= 5; i++ { //i只能在这里定义，不能拿出来var
		sum += i
		fmt.Println(sum)
	}
	//这里还有个死循环
	//for {
	//	fmt.Println("死循环")
	//}
	//这里是遍历字符串
	str := "hhhh你好"
	for i := 0; i < len(str); i++ { //用小于号，别用小于等于，不然会下标越界
		fmt.Println(string(str[i]))
	}
	for i, value := range str {
		fmt.Println(i, string(value))
	}

}
