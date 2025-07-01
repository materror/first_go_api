package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)
// Goal: make a website API that allows the user to parse HTMLs
func main() {
	htmlData := ""

	// Parse HTML string
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	fmt.Printf("%v", doc)
}