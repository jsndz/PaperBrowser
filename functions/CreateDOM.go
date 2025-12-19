package functions

import "fmt"

type Node struct {
	Tag      string
	Text     string
	Children []*Node
}

type DOM struct {
	Root *Node
}

func CreateDOM(tokens []Token) *DOM {
	var root *Node
	stack := make([]*Node, 0)

	for _, token := range tokens {
		switch token.Type {
		case "startTag":

			n := &Node{Tag: token.Content}

			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, n)

			} else {
				root = n
			}
			stack = append(stack, n)
		case "endTag":

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case "Text":
			n := &Node{Text: token.Content}
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, n)
		default:

		}
	}
	return &DOM{Root: root}
}

func PrintDOM(node *Node, depth int) {
	if node == nil {
		return
	}

	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	if node.Tag != "" {
		fmt.Println("<" + node.Tag + ">")
	} else if node.Text != "" {
		fmt.Println("\"" + node.Text + "\"")
	}

	for _, child := range node.Children {
		PrintDOM(child, depth+1)
	}
}
