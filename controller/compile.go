// controller/compile.go
package controller

import (
	"OLC2CLIENTE/compile"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func CompileCode(code string) string {
	//Inicia y genera lo que necesitamos con antlr

	inputStream := antlr.NewInputStream(code)

	lexer := gramAntlr.NewgramaticaLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := gramAntlr.NewgramaticaParser(tokenStream)
	parser.BuildParseTrees = true

	tree := parser.Inicio() // O el rule root de tu gramática

	// Visitor
	visitor := compile.NewCompilerVisitor() // luego lo cambias por tu visitor real
	visitor.Visit(tree)

	fmt.Println("\n SCOPE GLOBAL PARA VER DECLARACIONES \n" + visitor.ReportScope())
	fmt.Println("\n SCOPE PARA VER FUNCIONES GLOBALES \n" + visitor.ReportFunctions())

	return visitor.Salida
}
