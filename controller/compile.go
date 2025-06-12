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

	//Imprimir el arbol
	fmt.Println("\n ARBOL NO FORMATEADO \n" + tree.ToStringTree(nil, parser))

	// --- Print the formatted tree ---
	fmt.Println("Árbol FORMATEADO:")
	// Get the raw string from ANTLR
	rawTreeString := tree.ToStringTree(nil, parser)
	// Format it using our new function
	formattedOutput := FormatAntlrTree(rawTreeString)
	fmt.Println(formattedOutput) // Print the nicely formatted tree

	/*
		// --- Generación del archivo .dot ---
		// Creamos una nueva instancia del parser DOT
		dotParser := NewTreeStringParser()
		dotFilePath := "salida.dot" // Define el nombre del archivo .dot de salida

		// Llamamos al método que se encarga de parsear la cadena y generar el archivo DOT
		err := dotParser.GenerateDotFromParseTreeString(rawTreeString, dotFilePath)
		if err != nil {
			fmt.Println("Error al generar o escribir el archivo DOT:", err)
		} else {
			fmt.Printf("Archivo DOT '%s' generado exitosamente.\n", dotFilePath)
		}
	*/
	// Con esto:
	dotFilePath := "salida.dot"
	err := GenerateDotFromFormattedTreeString(formattedOutput, dotFilePath)
	println(err)
	var searchTree SearchTree = *NewSearchTree()
	searchTree.Visit(tree)
	// Visitor
	visitor := compile.NewCompilerVisitor() // luego lo cambias por tu visitor real

	// Visitar cada declaración y sentencia en orden
	for _, dclSimple := range searchTree.DeclaracionesSimples {
		visitor.Visit(dclSimple)
	}

	for _, dclSlice := range searchTree.DeclaracionesArreglos {
		visitor.Visit(dclSlice)
	}

	for _, dclStruct := range searchTree.DeclaracionesStructs {
		visitor.Visit(dclStruct)
	}

	for _, dclStruct2 := range searchTree.DeclaracionesStructs2 {
		visitor.Visit(dclStruct2)
	}

	for _, asign := range searchTree.Asignaciones {
		visitor.Visit(asign)
	}

	for _, stmt := range searchTree.Funciones {
		visitor.Visit(stmt)
	}

	for _, stmtStruct := range searchTree.FunctStruct {
		visitor.Visit(stmtStruct)
	}

	for _, stmtMain := range searchTree.FunctMain {
		fmt.Println("Visitando función main")
		visitor.Visit(stmtMain)
	}

	//visitor.Visit(tree) descomentar si quieres visitar todo el árbol

	fmt.Println("\n SCOPE GLOBAL PARA VER DECLARACIONES \n" + visitor.ReportScope())
	fmt.Println("\n SCOPE PARA VER FUNCIONES GLOBALES \n" + visitor.ReportFunctions())

	return visitor.Salida
}
