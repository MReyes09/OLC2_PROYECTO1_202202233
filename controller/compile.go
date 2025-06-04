package controller

import (
	"OLC2CLIENTE/compile"
	"OLC2CLIENTE/gramatica/gramAntlr"

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

	return visitor.Salida
}
