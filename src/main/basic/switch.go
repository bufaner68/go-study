package main

import "fmt"

//switch里不需要break语句，直接选择匹配的选项执行
func main(){
	var score int = 90
	switch score{
	case 90: fmt.Println("perfect")
	case 80: fmt.Println("fine")
	case 70: fmt.Println("ok")
	case 60: fmt.Println("pass")
	default: fmt.Println("fail")
	}


}