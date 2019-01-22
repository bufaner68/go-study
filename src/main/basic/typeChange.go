package  main

import "fmt"

func main(){
	var sum int = 17
	var count int = 5
	var mean float64

	//数据类型强制转换
	mean = float64(sum)/float64(count)
	fmt.Printf("mean的值为：%f\n",mean)

}