package main

import "fmt"

func main() {
	perim := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance()) //调用方法
	fmt.Println()                 // 包名.类型名.函数名

}
