//Go 语言循环语句
// 循环控制语句
//循环控制语句可以控制循环体内语句的执行过程。
//
//GO 语言支持以下几种循环控制语句：
//
//控制语句	描述
//break 语句	经常用于中断当前 for 循环或跳出 switch 语句
//continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
//goto 语句	将控制转移到被标记的语句。

package main

import "fmt"

func main() {
	var num int = 1
	for true {
		fmt.Printf("%d test \n", num)
		num += 1
	}
}
