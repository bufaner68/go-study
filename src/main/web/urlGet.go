// Fetch prints the content found at a URL.
package main
import (
"fmt"
"io"
"net/http"
"os"
"strings"
)

func main() {
//读取URL
for _, url := range os.Args[1:] {
	//判断URL是否是以“http：//”开头，使用string.HasPrefix函数判断
	if strings.HasPrefix(url,"http://")== false{
		//如果不是，就加上
		url = "http://" + url
	}
//http.Get函数是创建HTTP请求,如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
//resp的Body字段包括一个可读的服务器响应流
resp, err := http.Get(url)
fmt.Println(resp.Status)
if err != nil {
//如果错误信息不为空，则将日志输出到os.Stderr(go默认将日志输出到os.Stdout里)
//并输出错误信息err
fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//释放资源
os.Exit(1)
}
//将访问的请求结果写入os.Stdout,
io.Copy(os.Stderr,resp.Body)
//关闭资源
resp.Body.Close()
if err != nil {
fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
os.Exit(1)
}

}
}