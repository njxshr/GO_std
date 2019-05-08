//变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。
//变量可以通过变量名访问。
//Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
//声明变量的一般形式是使用 var 关键字：
// var identifier type
// 可以一次命名多个变量
// identifier1, identifier2 type

//package main
//
//import "fmt"
//
//func main()  {
//	fmt.Println("实例1")
//	var a string = "Runoob"
//	fmt.Println(a)
//
//	var b, c int = 1, 2
//	fmt.Println(b, c)
//}

// 变量申明
// 第一种，指定变量类型， 如果没有初始化， 则变量默认为零值
// var v_name n_type
// v_name = value

package GO_std

import "fmt"

func main() {
	fmt.Println("实例2")
	var a = "RUNOOB"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	var d *int
	fmt.Println(d)

	var f string = "RUNNING MAN"
	fmt.Println(f)

	// 简写
	g := "RUNNING MAN S"
	fmt.Println(g)
}

// 注意事项
// 如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。
//
//如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。
//
//如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误，例如下面这个例子当中的变量 a：
