package main

import (
	"math/cmplx"
	"fmt"
	"math"
)

/*常见的变量类型：
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
// 表示一个 Unicode 码点

float32 float64

complex64 complex128
//int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
//当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

func main() {
	//规范输出：%T 是输出变量类型，% v 是输出变量值
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*没有明确初始值的变量会被赋予零值
	数值类型 ：0
	布尔类型 ：false
	字符串：“”（空字符串）
	*/
	var i int
	var f float64
	var b bool
	var s string
	//输出字符串应该是%s,但是如果想输出字符串的零值，要用%q ???
	fmt.Printf("%d %f %t %q\n", i,f,b,s)

	//与C语言不同时是，Go在不同类型的项之间赋值时需要显式转换
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)

	//在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出
	v := 45.000
	fmt.Printf("v is of type%T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	/*
	//int 可以存放最大64位的整数，根据平台不同有时会更少。
	const Big = 1 << 100
	const Small = Big >> 99
	fmt.Println(Small)
	fmt.Println(Big)
	*/
}

