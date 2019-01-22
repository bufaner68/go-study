package main
import "fmt"
var (
    firstName, lastName, s  string
    i, year ,a,b,c int
    f float32
    input = "56.12 / 5212 / Go"
    format = "%f / %d / %s"
    birthday = "2018/07/09"
    geshi = "%d/%d/%d"
)
func main() {
    fmt.Println("Please input your full name: ")
    //输入两个字符串firstName和lastName
    fmt.Scanln(&firstName, &lastName)
 // fmt.Scanf(“%s %s”, &firstName, &lastName) 
    //输出字符串
    fmt.Printf("Hi %s %s!\n", firstName, lastName)
    //从字符串变量input, format里读取输入
    fmt.Sscanf(input, format, &f, &i, &s)
    fmt.Println("From the string we read: ", f, i, s)
    
    //练习使用Scanln输入和Sscanf从变量读取输入
    fmt.Println("请输入你的年龄：")
    fmt.Scanln(&year)
    fmt.Printf("%s的年龄是%d",lastName,year)
    fmt.Sscanf(birthday,geshi,&a,&b,&c)
    fmt.Printf("  %s的出生年月是:",firstName)
    fmt.Println(a,b,c)
}