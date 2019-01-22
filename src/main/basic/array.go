package main

import "fmt"

//数组的声明，初始化和调用
func main(){
	var arr1[6] int 
	for i:=0;i<6;i++{
		arr1[i] = i+1
	}
	
	fmt.Print(arr1)
}