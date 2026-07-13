package gmarkdown_test

import (
	"fmt"

	"github.com/dengmengmian/ghelper/gmarkdown"
)

func ExampleMdToHTML() {
	html := gmarkdown.MdToHTML("**bold** and *italic*")
	fmt.Println(html)
	// Output:
	// <p><strong>bold</strong> and <em>italic</em></p>
}
