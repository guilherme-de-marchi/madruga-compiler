package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Guilherme-De-Marchi/madruga-compiler/compiler"
)

func InitRepl() {
	cmp := compiler.NewCompiler(compiler.Source{})
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n-----------* Madruga language *-----------")
	for {
		fmt.Print("\n>> ")
		src, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[REPL ERROR] ", err)
		}
		cmp.Lexer = compiler.Source(src)
		tks, err := cmp.Lexer.Scan()
		if err != nil {
			fmt.Println(err)
			continue
		}
		compiler.PrintTokens(tks...)
	}
}
