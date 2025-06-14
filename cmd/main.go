package main

import (
	"github.com/claudiocaldeirao/experimental_compiler/internals/lexical"
	"github.com/claudiocaldeirao/experimental_compiler/internals/semantic"
	"github.com/claudiocaldeirao/experimental_compiler/internals/syntatic"
)

func main() {
	lexical.Parse()
	syntatic.Parse()
	semantic.Parse()
}
