package main

import "fmt"

func main() {
	slice := make([]int, 4, 10) //三个参数，一切片类型，二，切片长度，三，切片容量。
	slice2 := []int{1, 2, 3}
	fmt.Println(slice, slice2)
	slice[0] = 99
	slice = append(slice, 1)
	fmt.Println(slice)
	for i, v := range slice {
		fmt.Printf("索引: %d, 值: %d\n", i, v)
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("索引: %d, 值: %d\n", i, slice[i])
	}

}
