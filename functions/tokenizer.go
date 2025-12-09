package functions

import (
	"fmt"
)

type Token struct {
	Type string

	Content string
}

func Tokenizer(html string) []Token {
	tokens := make([]Token, 0)
	var isTag bool
	var isContent bool
	var tag string
	var content string
	var isClosingTag bool
	for index, ch := range html {
		fmt.Println(string(ch))
		if ch == '<' {
			if isContent && content != "" {
				tokens = append(tokens, Token{Type: "Text", Content: content})
				content = ""
				fmt.Println(tokens)
				isContent = false
			}
			if index+1 < len(html) && html[index+1] == '/' {
				isClosingTag = true
			} else {
				isTag = true
			}
		}
		if ch == '>' {
			if isTag {
				tokens = append(tokens, Token{Type: "startTag", Content: tag})
				isTag = false
				tag = ""
				fmt.Println(tokens)
			}
			if isClosingTag {
				tokens = append(tokens, Token{Type: "endTag", Content: tag})
				tag = ""
				isClosingTag = false
				fmt.Println(tokens)
			}
			if index+1 < len(html) && html[index+1] != '<' {
				isContent = true
			}

		}
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == ' ' {
			if isTag || isClosingTag {
				if ch != ' ' {
					tag = tag + string(ch)
				}
			}
			if isContent {
				content = content + string(ch)
			}

		}

	}
	return tokens
}
