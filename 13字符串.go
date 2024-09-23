package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "你好"
	fmt.Println(len(str)) //统计字符串长度
	for i, v := range str {
		fmt.Println(i, v, string(v))
	}

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println("第二种", string(r[i]))
	}

	num1, _ := strconv.Atoi("123")
	fmt.Println(num1) //字符串转整数

	str1 := strconv.Itoa(88)
	fmt.Println(str1, "整数转字符串")

	count := strings.Count("sgwesfgwasedfqwedqwrewsfdrwqedwqefwefwe", "e")
	fmt.Println(count, "统计出现次数")

	fmt.Println(strings.EqualFold("hello", "HELLo"), "不区分大小写的比较")

	fmt.Println("hello" == "HELLo", "区分大小写的比较")

	fmt.Println(strings.Index("afawdaqwedwa", "q"), "第一次出现的位置找不到就返回-1。")

	fmt.Println(strings.Repeat("-", 100))

	fmt.Println(strings.Replace("sdfgsrdgfasgaesghasrgaskgm,vpsadgksp", "a", "aaaaaaaaaaaa", -1), "替换字符,最后一个数字表示替换次数..-1表示全部替换。")

	fmt.Println(strings.Split("go-python-java", "-"), "切割字符串")

	fmt.Println(strings.ToLower("SEGSGSDH通通变成小写。"))

	fmt.Println(strings.ToUpper("dghrsdhsdrhj通通变成大写。"))

	fmt.Println(strings.TrimSpace("      去掉左右   两边的空格。    "))

	fmt.Println(strings.Trim("```````啊啊啊啊```````", "`"), "去掉左右两边指定符号。")

	fmt.Println(strings.TrimLeft("```````啊啊啊啊```````", "`"), "去掉左边指定符号。右边的就不展示了。")

	fmt.Println(strings.HasSuffix("判断是否已指定字符串结束。", "结束。"))
	fmt.Println(strings.HasPrefix("那么这个就是开始", "那么."))
}
