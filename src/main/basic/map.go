package main

import "fmt"

func main(){
	//创建map集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	//给map映射的键值对赋值
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["India"] = "新德里"

	//输出map映射的键值对
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是：",countryCapitalMap[country])
	}
	//第二种获取集合中值得方法，同时还可以查看元素在映射里是否存在
	capital,ok := countryCapitalMap["郑州"]
	if ok{
		fmt.Println("郑州的首都：",capital)
	} else {
		fmt.Println("郑州的首都不存在")
	}

	fmt.Println("=======================================")
	//delete()函数用于删除map集合的元素，参数是map和对应的key
	delete(countryCapitalMap,"Italy")
	fmt.Println("Italy已被删除")
	fmt.Println("删除后的map为：")
	for country := range countryCapitalMap{
		fmt.Println(country,"->",countryCapitalMap[country])
	}
}