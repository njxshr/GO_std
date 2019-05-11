// Go 语言常量
//常量是一个简单值的标识符，在程序运行时，不会被修改的量。
//
//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
//
//常量的定义格式：
// const identifier [type] = value
//你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
//
//显式类型定义： const b string = "abc"
//隐式类型定义： const b = "abc"
//多个相同类型的声明可以简写为：
//
//const c_name1, c_name2 = value1, value2

package main

//func main() {
//	const LENGTH  int = 10
//	const WIDTH  int = 5
//	var area int
//	const a, b, c = 1, false, "str" // 多重赋值
//	area = LENGTH * WIDTH
//	fmt.Println("面积为: %d", area)
//	println()
//	println(a, b, c)
//}

// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
//const (
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)
//
//func main() {
//	println(a, b, c)
//}

// iota
// 可以被编译器改变的常量
//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
//iota 可以被用作枚举值：
const (
	a = iota
	b = iota
	c = iota
)

func main() {
	println(a, b, c)
}
