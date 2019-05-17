package main

// 函数作为实参
import (
	"fmt"
	"math"
)

func main() {
	// 申明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	// 使用函数
	fmt.Println(getSquareRoot(9))
}
