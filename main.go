package main

import (
	"fmt"

	"github.com/jsndz/PaperBrowser/functions"
)

var html = `<html><body><p >Hello This should render a paragraph</p></body></html>`

func main() {
	fmt.Println("Hello, World!")
	tokens := functions.Tokenizer(html)
	fmt.Println(tokens)
}
