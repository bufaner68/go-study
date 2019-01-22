package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("a+b的值是：%d\n",c)
	c = a - b
	fmt.Printf("a-b的值是：%d\n",c)
	c = a * b
	fmt.Printf("a*b的值是：%d\n",c)
	c = a / b
	fmt.Printf("a/b的值是：%d\n",c)
	c = a % b
	fmt.Printf("a取余b的值是：%d\n",c)
	a++
	fmt.Printf("a++的值是：%d\n",c)
	a = 21
	a--
	fmt.Printf("a--的值是：%d\n",c)
}