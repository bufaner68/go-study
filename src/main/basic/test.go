package main

import (
	"io"
	"os"
	"fmt"
)

type Node struct {
	Type NodeType
	Data string
	Attr []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)
type Attribute struct {
	Key, Val string
}
//func Parse(r io.Reader) (*Node, error)

func main() {
	doc, err := Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *Node) []string {
	if n.Type == ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}