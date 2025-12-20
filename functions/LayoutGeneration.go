package functions

import "fmt"

type State struct {
	CursorX int
	CursorY int
}

type PaintInstruction struct {
	CursorX  int
	CursorY  int
	Function string
	Content  string
}

func LayoutGeneration(node *Node, instructions *[]*PaintInstruction, state *State) {
	if node == nil {
		return
	}
	switch node.Tag {
	case "p":
		state.CursorX += 1
		state.CursorY += 1
	}

	if node.Text != "" {
		*instructions = append(*instructions, &PaintInstruction{CursorX: state.CursorX, CursorY: state.CursorY, Function: "WriteText", Content: node.Text})
		state.CursorX += len(node.Text)
	}

	for _, v := range node.Children {
		LayoutGeneration(v, instructions, state)
	}
}

func PrintLayout(pi *[]*PaintInstruction) {
	for _, v := range *pi {
		fmt.Println(v)
	}
}
