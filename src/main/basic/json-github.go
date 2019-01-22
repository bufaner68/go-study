package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

//const IssueURL = "https://api.github.com/search/issues"
const IssueURL = "https://api.github.com/search/issues"
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"create_at"`
	Body     string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {

	q := url.QueryEscape("application/vnd.github.symmetra-preview+json")
	urls := IssueURL + "?q=" + q
	resp, err := http.Get(urls)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("start encode ...")

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("error...")
		return nil, err
	}
	return &result, nil
}

func main() {
	//调用SearchIssues对GitHub进行数据搜索
	result, err := SearchIssues(os.Args[1:])
	//如果搜索出错，返回错误信息
	if err != nil {
		log.Fatal(err)
	}
	//对搜索的结果条数进行计数，返回条数
	fmt.Printf("%d issue: \n", result.TotalCount)
	//把搜索的结果按条输出
	for _, item := range result.Items {
		fmt.Printf("#<%-5d> <%9.9s> <%.55s>\n\n", item.Number, item.User.Login, item.Title)
	}
}
