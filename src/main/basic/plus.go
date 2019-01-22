package main

import "fmt"

//九九乘法表
func plus(){
	var row  = 1
	var column  = 1
	//循环次数9次
	for i:= row;i<10;i++{
		//每次循环
		for j:=column;j<=i;j++{
			res:=i*j
			fmt.Print(j,"*",i,"=",res,"  ")
		}
		fmt.Print("\n")
	}
}

func main(){
	plus()
}