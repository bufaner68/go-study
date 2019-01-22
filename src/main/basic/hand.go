package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )

   //空指针的值为nil,指代零值
   var ptr *int
   //Printf可以输出格式化的字符串，不能输出整型或变量
   //println可以输出字符串，变量，整型
   fmt.Printf("ptr的值为：%x\n",ptr)
   //空指针判断
   if ptr == nil{
   		fmt.Println("是空指针")
   }
}