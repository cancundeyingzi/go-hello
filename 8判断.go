package main

import "fmt"

func main() {
	var count = 30
	if count < 30 {
		fmt.Println("库存不足")
	} else if count > 50 {
		fmt.Println("库存还够")
	} else {
		fmt.Println("快不够了")
	}

	switch count / 10 {
	case 3, 4, 5, 6:
		fmt.Println("30")
	case 2:
		fmt.Println("20")
	default:
		fmt.Println("其他")
	}

}
