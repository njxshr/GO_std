package main

import "fmt"

// go 运算符 和其他语言一样
// 假定 A 值为 10，B 值为 20
//+	相加	A + B 输出结果 30
//-	相减	A - B 输出结果 -10
//*	相乘	A * B 输出结果 200
///	相除	B / A 输出结果 2
//%	求余	B % A 输出结果 0
//++	自增	A++ 输出结果 11
//--	自减	A-- 输出结果 9

func main() {
	var a bool = true
	var b bool = true
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	//修改 a 和 b 的值
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
