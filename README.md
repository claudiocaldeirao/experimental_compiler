# Experimental Compiler

This is a simple educational programming language designed as a pseudocode-like syntax that compiles to JavaScript. This project is written entirely in Go and implements lexical, syntactic, and semantic analysis phases, producing JavaScript as output.

---

## Project Structure

```bash
experimental_compiler/
├── cmd/    # Entry point (CLI to compile code)
│   └── main.go
├── internals/
│   ├── lexical/
│   │   └── lexical_analyzer.go     # Lexer (tokenizer)
│   ├── syntactic/
│   │   └── syntactic_analyzer.go   # Parser (syntax validator)
│   ├── semantic/
│   │   └── semantic_analyzer.go    # Semantic Analyzer (e.g: check unused / used before declaration)
│   ├── reader/
│   │   └── reader.go               # Reads the source code file (using made up extension .cdl that extend for CaldeiLang, our example pseudo-code)
│   ├── generator/
│   │   └── generator.go            # Generate target language (in our example, JS)
├── go.mod
├── README.md
```

## Token table

| Category        | Token               | Example Code          | Description               |
| --------------- | ------------------- | --------------------- | ------------------------- |
| **Keywords**    | `IF`                | `IF x > 0 THEN`       | Conditional               |
|                 | `THEN`, `ELSE`      |                       | Conditional branches      |
|                 | `WHILE`, `DO`       | `WHILE x < 10 DO`     | Loops                     |
|                 | `BEGIN`, `END`      | `BEGIN` ... `END`     | Block delimiters          |
|                 | `PRINT`             | `PRINT x`             | Output command            |
| **Operators**   | `+`, `-`, `*`, `/`  | `x = a + b`           | Arithmetic                |
|                 | `=`, `==`, `<`, `>` |                       | Assignment and comparison |
| **Identifiers** | `ID`                | `x`, `total`, `count` | Variable names            |
| **Literals**    | `NUM`               | `10`, `3.14`          | Numeric constants         |
| **Symbols**     | `(`, `)`            | `PRINT(x)`            | Parentheses               |
|                 | `;`                 | `x = 5;`              | End of statement          |

## 🚀 Getting Started

Clone the repository and run the compiler with:

```bash
go run ./cmd/main.go
```

Example source:

```c
BEGIN
x = 10;
y = x + 2;
PRINT y;
END
```

Will generate:

```javascript
let x = 10;
let y = x + 2;
console.log(y);
```

## Analyzers

`lexical`
Contains the tokenizer which processes raw source code and generates a list of tokens.

`syntactic`
Implements the parser which validates the grammar and optionally builds an AST.

`semantic`
Implement semantic checks like type validation and scope rules.

## License

This project it's under the [MIT LICENSE](LICENSE)
