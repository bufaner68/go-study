// Server2 is a minimal "echo" and counter server.
package main
import (
"fmt"
"log"
"net/http"
"sync"
)
var mu sync.Mutex
var count int
func main() {
	//对/的请求，调用handler函数进行处理
	http.HandleFunc("/", handler)
	//对/count的请求，调用counter函数进行处理
	http.HandleFunc("/count", counter)
	//监听localhost：8000地址
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
mu.Lock()
//对处理次数进行计数
count++
	fmt.Fprintf(w, "Count %d\n", count)
mu.Unlock()
//输出请求的URL地址
fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
mu.Lock()
//输出处理次数count
fmt.Fprintf(w, "Count %d\n", count)
mu.Unlock()
}