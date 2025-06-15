package main

import (
	"fmt"
	"log"

	"github.com/claudiocaldeirao/experimental_compiler/internal/lexical"
	"github.com/claudiocaldeirao/experimental_compiler/internal/reader"
	"github.com/claudiocaldeirao/experimental_compiler/internal/semantic"
	"github.com/claudiocaldeirao/experimental_compiler/internal/syntactic"
)

func main() {
	textBuffer, err := reader.ReadFile("example/source.cdl")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-------------------- Source Code --------------------")
	fmt.Println(textBuffer)

	fmt.Println("Running lexical analysis...")
	lexer := lexical.NewLexer(textBuffer)
	tokens := lexer.Tokenize()

	fmt.Println("-------------------- Token Table --------------------")
	for _, t := range tokens {
		fmt.Printf("Token: %-10s Value: %s\n", t.Type, t.Lexeme)
	}

	fmt.Println("\nRunning syntactic analysis...")
	syntacticParser := syntactic.NewParser(tokens)
	statements := syntacticParser.ParseProgram()

	if statements == nil {
		log.Fatal("Failed to parse source code")
	}

	semanticAnalyzer := semantic.NewSemanticAnalyzer(statements)
	semanticAnalyzer.Analyze()
}
