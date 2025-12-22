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
	dom := functions.CreateDOM(tokens)
	functions.PrintDOM(dom.Root, 0)
	instructions := []*functions.PaintInstruction{}
	functions.LayoutGeneration(dom.Root, &instructions, &functions.State{CursorX: 0, CursorY: 0})
	functions.PrintLayout(&instructions)
	functions.Paint(&instructions)
}
