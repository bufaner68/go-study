package main

import "fmt"

//阶乘
func Factorial(n uint64)(result uint64){
	if n>0 {
		result := n*Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int)int{
	if n<2 {
		return 1
	}else{
		return fibonacci(n-1)+fibonacci(n-2)
	}
}

func main(){
	var num = 6
	fmt.Printf("%d的阶层是：%d",num,Factorial(uint64(num))) 
	fmt.Println()
	fmt.Printf("斐波那契数列为：%d",fibonacci(5))
}