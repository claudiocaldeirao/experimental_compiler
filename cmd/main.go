package main

import (
	"fmt"
	"log"
	"os"

	"github.com/claudiocaldeirao/experimental_compiler/internal/generator"
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
		fmt.Printf("Token: %-10s Lexeme: %s\n", t.Type, t.Lexeme)
	}

	fmt.Println("\nRunning syntactic analysis...")
	syntacticParser := syntactic.NewParser(tokens)
	ast := syntacticParser.ParseProgram()

	if ast == nil {
		log.Fatal("Failed to parse source code")
	}

	fmt.Println("-------------------- Abstract Syntax Tree (AST) --------------------")
	for _, interfaceStatement := range ast {
		assignStatement, assignOk := interfaceStatement.(syntactic.AssignStatement)
		printStatement, printOk := interfaceStatement.(syntactic.PrintStatement)

		if assignOk {
			fmt.Printf("Assign Statement: Variable: %s, Expression: %s\n", assignStatement.VarName, assignStatement.ExpressionTokens)
		} else if printOk {
			fmt.Printf("Print Statement: Variable: %s, Expression: %s\n", printStatement.VarName, printStatement.ExpressionTokens)
		} else {
			fmt.Println("Unknown statement type in AST")
		}

	}

	fmt.Println("\nRunning semantic analysis...")
	semanticAnalyzer := semantic.NewSemanticAnalyzer(ast)
	semanticAnalyzer.Analyze()

	fmt.Println("\nGenerating JavaScript code...")
	compiledCode := generator.GenerateJS(ast)
	fmt.Println("-------------------- Compiled Code --------------------")
	fmt.Println(compiledCode)

	os.WriteFile("index.js", []byte(compiledCode), 0644)
}
