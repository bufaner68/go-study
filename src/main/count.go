package main
 
import (
        "bufio"
        "fmt"
        "strings"
)
 
func main() {
 
        var c WordsCounter
        fmt.Fprintf(&c, "hello world 123")
        fmt.Println(c) //输出 3
}

//下面*ByteCounter类型里的Write方法，仅仅在丢失写向它的字节前统计它们的长度。
//(在这个+=赋值语句中，让len(p)的类型和*c的类型匹配的转换是必须的。)
type ByteCounter int
 
func (c *ByteCounter) Write(p []byte) (int, error) {
        *c += ByteCounter(len(p)) // convert int to ByteCounter
        return len(p), nil
}
 
//定义类型
type WordsCounter int
 
//满足相同接口的类型
func (w *WordsCounter) Write(p []byte) (int, error) {
        //分隔字符串
        s := strings.NewReader(string(p))
        bs := bufio.NewScanner(s)
        bs.Split(bufio.ScanWords)
        sum := 0
        for bs.Scan() {
                sum++
        }  
        *w = WordsCounter(sum)
        return sum, nil
}