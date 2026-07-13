package gmarkdown

import (
	"strings"
	"testing"
)

func TestMdToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:     "bold",
			input:    "**bold**",
			contains: []string{"<strong>", "</strong>", "bold"},
		},
		{
			name:     "italic",
			input:    "*italic*",
			contains: []string{"<em>", "</em>", "italic"},
		},
		{
			name:     "header",
			input:    "# Title",
			contains: []string{"<h1", "Title", "</h1>"},
		},
		{
			name:     "link",
			input:    "[link](https://example.com)",
			contains: []string{"<a", "href=\"https://example.com\"", "link", "</a>"},
		},
		{
			name:     "code_inline",
			input:    "`code`",
			contains: []string{"<code>", "code", "</code>"},
		},
		{
			name:     "empty",
			input:    "",
			contains: nil,
		},
		{
			name:     "paragraph",
			input:    "Hello world",
			contains: []string{"<p>", "Hello world", "</p>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := MdToHTML(tt.input)
			for _, want := range tt.contains {
				if want != "" && !strings.Contains(html, want) {
					t.Errorf("MdToHTML(%q) did not contain %q, got: %s", tt.input, want, html)
				}
			}
		})
	}
}
