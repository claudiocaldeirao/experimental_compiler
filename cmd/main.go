package main

import (
	"fmt"
	"log"

	"github.com/claudiocaldeirao/experimental_compiler/internals/lexical"
	"github.com/claudiocaldeirao/experimental_compiler/internals/reader"
	"github.com/claudiocaldeirao/experimental_compiler/internals/semantic"
	"github.com/claudiocaldeirao/experimental_compiler/internals/syntatic"
)

func main() {
	textBuffer, err := reader.ReadFile("example/source.cdl")
	if err != nil {
		log.Fatal(err)
	}

	tokens := lexical.Tokenize(textBuffer)
	for _, t := range tokens {
		fmt.Printf("Token: %-10s Value: %s\n", t.Type, t.Value)
	}

	// TODO
	syntatic.Parse()
	semantic.Parse()
}
