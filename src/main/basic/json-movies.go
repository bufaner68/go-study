package main

import (
	"image"
	"net/url"
	"fmt"
	"net/http"
	"log"
)

const urls = "http://www.omdbapi.com/"

type moviesSearhResult struct{
	TotalCount int
	Image image.Image
}

func searchMovies()(*moviesSearhResult,error){
	var content string
	fmt.Println("请输入您想要找的电影名字：\n")
	fmt.Scanln(&content)
	q := url.QueryEscape(content)
	surl := urls + "?q=" + q
	resp,err := http.Get(surl)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()
	fmt.Println("start coding")

	if resp.StatusCode!= http.StatusOK{
		return nil, fmt.Errorf("search failed：%s",resp.Status)
	}
	var result moviesSearhResult
	return &result,nil
}

func main(){
	result,err := searchMovies()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(result.TotalCount)
	fmt.Println(result)
}