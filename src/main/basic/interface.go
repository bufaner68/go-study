package main

import "fmt"

//定义一个接口Phone
type Phone interface{
	call()
}

//定义一个结构体NokiaPhone
type NokiaPhone struct{
	
}

//结构体NokiaPhone实现接口里的call方法
func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am nokiaPhone, I can call")
}

//定义IPhone结构体
type IPhone struct{
	
}

//结构体iPhone实现接口里的call方法
func (iPhone IPhone) call(){
	fmt.Println("I am iPhone, I can call")
}


func main(){
	//定义接口类型的变量
	var phone Phone

	//为接口变量赋值
	phone = new (NokiaPhone)
	//调用call方法
	phone.call()

	phone = new (IPhone)
	phone.call()
}