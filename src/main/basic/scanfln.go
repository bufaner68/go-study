package main

import "fmt"

func main(){
	var (
		name,address string
		birthday,year int
		num = "18"
		format = "%d"
	)
	//Scanln 将从标准输入的带有空格的字符串值保存到相应的变量里去，并以一个新行结束输入
	//Scanf做相同的工作，但它使用第一个参数指时输入格式
	//Sscan系列函数也是读取输入，但它是用来从字符串变量里读取，而不是从标准（os.Stdin）里读取
	fmt.Println("请输入您的姓名：")
	//从命令行获取name输入
	fmt.Scanln(&name)
	//从变量直接获取生日
	fmt.Sscanf(num,format,&year)
	fmt.Println("请输入您的生日：")
	//从命令行获取birthday输入
	fmt.Scanln(&birthday)
	fmt.Println("请输入您的地址：")
	//从命令行获取address输入
	fmt.Scanln(&address)
	fmt.Printf("%s的年龄是%d,生日是：%d,家住在%s",name,year,birthday,address)
}