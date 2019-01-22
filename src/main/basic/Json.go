package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main(){
	//json.Marshal()用于返回一个编码后的字节
	//json.MarshalIndent()用于将数据按格式转换后返回
	//当结构体成员为空或为零值时，不生成JSON对象
	data,err := json.MarshalIndent(movies,"","    ")
	if err!=nil{
		log.Fatal("JSON marshaling failed:%s",err)
	}
	fmt.Printf("%s\n",data)

	//解码:将JSON数据解码城go语言的数据结构，使用unmarshal()函数完成，
	// 可以只选择自己感兴趣的部分解码，那么其他JSON成员将被忽略
	var titles[]struct{Title string}
	if err := json.Unmarshal(data,&titles);err!= nil{
		log.Fatal("JSON unmarshaling failed:%s",err)
	}
	fmt.Println(titles)
	}