package main

import (
	"fmt"
	"math"
)

func main(){
	//未从包内导出的名字在包外无法访问
	//fmt.Println(math.pi)
	//大写字母开头的名字可以被访问，因为它们是已导出的
	fmt.Println(math.Pi)
}