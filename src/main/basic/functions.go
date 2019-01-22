package main

import "fmt"

/*函数声明方式：
func 函数名(参数名1 类型,参数名2 类型)返回值类型{

}
*/
func add(x int,y int)int{
	return x + y
}

//如果一个函数的所有参数类型都是一样的，那么只需在最后一个参数后边加上函数类型，其他的都可以省略
func plus(x,y int)int{
	return x*y
}

//函数可以返回多个返回值,swap函数返回了两个字符串
func swap(a,b string)(string,string) {
	return a,b
}

//没有参数的return语句直接返回已经命名的返回值num1和num2
//直接返回语句适合于短函数，在复杂函数中会使可读性变差
func calculate(left,right int)(num1,num2 int){
	num1 =left + right
	//println(num1)为什么不打印这句？？？
	return
}

func main(){
	fmt.Println(plus(2,6))
	fmt.Println(add(2,6))
	//分别将swap的两个返回值赋给a和b
	a,b :=swap("Hello","World")
	fmt.Println(a,b)
	fmt.Println(swap("Hello"," World"))
	fmt.Println(calculate(8,8))
}

