package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	//输出请求的URL路径到工作台
	fmt.Println("Path:",r.URL.Path)

	fmt.Fprintf(w,"hello go")
}

func main(){
	//把所有发送到/站点的请求转给sayhelloName函数
	http.HandleFunc("/",sayhelloName)
	//监听9090端口，如果遇到错误，输出到日志里
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServer:",err)
	}
}
