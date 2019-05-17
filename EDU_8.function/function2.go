package main

import "fmt"

// 交换两个字符串位置

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Google", "Github")
	fmt.Println(a, b)
}
