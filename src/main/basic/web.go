// Server1 is a minimal "echo" server.
package main
import (
"fmt"
"log"
"net/http"
)

//main函数将所有发送到/路径下的请求和handler函数关联起来，
// /开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口。
//发送到这个服务的“请求”是一个http.Request类型的对象，这个对象中包含了请求中的一系列相关字段
//其中就包括我们需要的URL。
//当请求到达服务器时，这个请求会被传给handler函数来处理
//这个函数会将/hello这个路径从请求的URL中解析出来，然后把其发送到响应中，这里我们用的是标准输出流的fmt.Fprintf
func main() {
http.HandleFunc("/", handler) // each request calls handler
log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}