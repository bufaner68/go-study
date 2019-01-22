package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"bufio"
)

func main() {
	//使用readFile（）读取文件内容
	file_bytes, err := ioutil.ReadFile("F:\\浪潮\\学习资料\\RESTful.txt")
	if err != nil {
		panic(err)
	}
	//使用，分割读取出来的内容并打印出来
	lines := strings.Split(string(file_bytes), ",")
	fmt.Println(lines)
	fmt.Println("=================")


	//f里放的是文件路径
	f,err := os.Open("F:\\浪潮\\学习资料\\RESTful.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//使用bufio.NewScanner()读取文件每行数据
	scanner := bufio.NewScanner(f)
	//每次读入一行
	for scanner.Scan(){
		//打印一行
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}