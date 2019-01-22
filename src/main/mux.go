package main

import (
	"fmt"
	"net/http"
	"log"
)


//1.想象一个电子商务网站，为了销售它的数据库将它物品的价格映射成美元。下面这个程序可
//能是能想到的最简单的实现了。它将库存清单模型化为一个命名为database的map类型，我
//们给这个类型一个ServeHttp方法，这样它可以满足http.Handler接口。这个handler会遍历整
//个map并输出物品信息。

//2.让我们使用/list来调用已经存在的
//这个行为并且增加另一个/price调用表明单个货品的价格，像这样/price?item=socks来指定一
//个请求参数。
func main(){
	db := database{"shoes":50,"socks":5}
	//mux := http.NewServeMux()
	///HandlerFunc是一个让函数值满足一个接口的适配器，使得一个单一的类型database以多种方式满足http.Handler接口：
	//一种是通过它的list方法，一种是通过它的price方法
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	//mux.Handle("/list",http.HandlerFunc(db.list))
	//mux.Handle("/price",http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

type dollars float32

func(d dollars) String() string{
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

func(db database) list(w http.ResponseWriter,req *http.Request){
	for item,price := range db{
		fmt.Fprintf(w,"%s: %s\n",item,price)
	}
}

func (db database)price(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	price,ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such item:%q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
}