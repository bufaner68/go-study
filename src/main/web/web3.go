package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	//定义数据库
	db := database{"shoes": 50, "socks": 5}
	//监听8000端口，调用db的处理函数
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
type dollars float32
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
type database map[string]dollars
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}