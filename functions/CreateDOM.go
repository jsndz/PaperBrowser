package functions

type Node struct {
	Element   string
	Type      string
	ChildNode *Node
}

type DOM struct {
	root *Node
}

func CreateDOM(tokens []Token) {
	stack := make([]Node, 0)
	for _, token := range tokens {
		if token.Type == "startTag" {
			stack = append(stack, Node{Element: token.Content, Type: token.Type})
		} else if token.Type == "text" {
			stack = append(stack, Node{Element: token.Content, Type: token.Type})
		}
	}
}
