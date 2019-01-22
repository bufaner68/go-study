package main

import "fmt"

//var用于声明一个变量，可以与函数、包在同一级别声明，变量名后跟变量类型
var str1 string

//:=不能再函数外使用，函数外必须以func或var开头
//bb := 9

func main(){
	//变量声明和调用
	var str2 string
	str1 = "hello"
	str2 = "Go"
	fmt.Println(str1,str2)

	//变量声明时可以直接初始化，如果直接初始化，则可以不需要加上变量类型
	var inta,intb = 3,3

	//var numm 变量如果只声明不使用，会报错
	fmt.Println(inta,intb)

	//:= 是一种简洁的赋值语句，在函数内使用时可以代替var
	aa :=6
	fmt.Println(aa)
}
