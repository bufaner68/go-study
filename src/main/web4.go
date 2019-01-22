package main

import (
	"net/http"
	"fmt"
	"log"
)
func main(){
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func(d dollars) String() string{
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

/*
func(db database) ServeHTTP(w http.ResponseWriter,req *http.Request){
	for item,price := range db{
		fmt.Fprintf(w,"%s:%s\n",item,price)
	}
}
*/
func (db database) ServeHTTP(w http.ResponseWriter,req *http.Request){
	//switch匹配URL
	switch req.URL.Path{
		//匹配到/list时，输出全部
	case "/list":
		for item,price := range db{
			fmt.Fprintf(w,"%s: %s\n",item,price)
		}
		//匹配到/price时，输出对应item的price
	case "/price":
		item := req.URL.Query().Get("item")
		price,ok := db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"no such item:%q\n",item)
			return
		}
		fmt.Fprintf(w,"%s\n",price)
		//匹配到其他时，输出找不到
	default:
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w,"no such page:%s\n",req.URL)
		msg := fmt.Sprintf("no such page:%s\n",req.URL)
		http.Error(w,msg,http.StatusNotFound)
	}
}