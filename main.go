package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Guilherme-De-Marchi/madruga-compiler/compiler"
)

func main() {
	switch len(os.Args) {
	case 1:
		InitRepl()
	case 2:
		src, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Println(err)
		}
		cmp := compiler.NewCompiler(compiler.Source(src))
		tks, err := cmp.Lexer.Scan()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tokens: ", tks)
	default:
		fmt.Println("Usage: go run . [optional: src path]")
	}
}
