package main

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"html/template"
)

//1.增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。例如，一
//个形如  /update?item=socks&price=6  的请求会更新库存清单里一个货品的价格并且当这个货
//品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）

//2.修改/list的handler让它把输出打印成一个HTML的表格而不是文本。
//html/template包(§4.6)可能会对你有帮助。

func main(){
	db := database{"shoes":50,"socks":5}
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	http.HandleFunc("/update",db.update)
	http.HandleFunc("/delete",db.delete)
	//http.HandleFunc("create",db.create)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

type dollars float32

func(d dollars) String() string{
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

func(db database) list(w http.ResponseWriter,req *http.Request){
	//for item,price := range db{
	//	fmt.Fprintf(w,"%s: %s\n",item,price)
	//}
	//输出格式是HTML
	var shopList = template.Must(template.New("shopList").Parse(`
    <h1>shopList</h1>
    <table>
    <tr style='text-align: left'>
    <th>item</th>
    <th>    </th>
    <th>price</th>
    </tr>
    </table>
    `))
	shopList.Execute(w, nil)

	const templ = `<p>{{.A}}    {{.B}}</p>`
	type data struct {
		A string
		B dollars
	}
	t := template.Must(template.New("escape").Parse(templ))

	for item, price := range db {

		var dat = data{item, price}

		err := t.Execute(w, dat)
		if err != nil {
			log.Fatal(err)
		}
	}
}


func(db database) price(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	price,ok := db[item]
	if !ok{
		fmt.Fprintf(w,"no such item:%s",item)
	}
	fmt.Fprintf(w,"%s",price)
}

//实现对map集合的删除功能
func(db database) delete(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	delete(db,item)
	var price dollars
	for item,price = range db{
		fmt.Fprintf(w,"%s: %s",item,price)
	}
}

//实现对map集合的数据修改功能
func(db database) update(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "price value error: %q\n", price)
		return
	}
	db[item] = dollars(newPrice)
	for it,pr := range db{
		fmt.Fprintf(w,"%s: %s\n",it,pr)
	}
}