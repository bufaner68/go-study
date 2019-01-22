package main

import "fmt"

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main(){
	//切片，实际的是获取数组的某一部分，len切片<=cap切片<=len数组
	//切片由三部分组成：指向底层数组的指针、len、cap。
	
	//创建切片
	var numbers = []int {0,1,2,3,4,5,6,7}
	printSlice(numbers)

	//打印原始切片
	fmt.Println("numbers ==",numbers)

	//打印子切片从索引1到4
	fmt.Println("numbers[1:4] ==",numbers[1:4])

	/* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)

   fmt.Println("=============================")
   //append（）和copy（）函数使用：从切片copy元素和向切片追加元素
   var num []int
   printSlice(num)

   //允许追加空切片
   num = append(num,0)
   printSlice(num)

   //向切片添加一个元素
   num = append(num,2)
   printSlice(num)

   //同时添加多个元素
   num = append(num,3,4,5,6)
   printSlice(num)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   num1 := make([]int,len(num),(cap(num))*2)

   /*copy num的内容到num1*/
   copy(num1,num)
   printSlice(num1)

}