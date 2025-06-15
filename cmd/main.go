package main

import (
	"fmt"
	"log"

	"github.com/claudiocaldeirao/experimental_compiler/internal/lexical"
	"github.com/claudiocaldeirao/experimental_compiler/internal/reader"
	"github.com/claudiocaldeirao/experimental_compiler/internal/semantic"
	"github.com/claudiocaldeirao/experimental_compiler/internal/syntatic"
)

func main() {
	textBuffer, err := reader.ReadFile("example/source.cdl")
	if err != nil {
		log.Fatal(err)
	}

	tokens := lexical.Tokenize(textBuffer)
	for _, t := range tokens {
		fmt.Printf("Token: %-10s Value: %s\n", t.Type, t.Lexeme)
	}

	// TODO
	syntatic.Parse()
	semantic.Parse()
}
