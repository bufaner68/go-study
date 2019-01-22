package main

import (
	"fmt"
)

func main() {
	sum := 0
	//for循环，与其他语言不同的是没有小括号
	//go语言里只有for循环
	//for循环可以省略初始值，循环条件和每次循环后执行的语句
	for i := 0; i < 6; i++ {
		sum += i
	}
	fmt.Println(sum)

	//省略初始值和每次循环后执行的语句
	s := 1
	for s < 6 {
		s += s
	}
	fmt.Println(s)

	//range区域内循环
	ss,sep := "",""
	for _,arg :=range os.Args[1:]{
		s += sep +arg
		sep = " "
	}
	fmt.Println(ss)
}
