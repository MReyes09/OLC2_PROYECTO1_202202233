// controller/compile.go
package controller

import (
	"OLC2CLIENTE/compile"
	"OLC2CLIENTE/compile/generatorARM"
	"os"
	"os/exec"

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

	// Obtener el árbol en texto desde ANTLR
	rawTreeString := tree.ToStringTree(nil, parser)

	// Formatear el árbol (usa tu función personalizada)
	formattedOutput := FormatAntlrTree(rawTreeString)

	// Imprimir el árbol formateado (opcional)
	fmt.Println("Árbol FORMATEADO:")
	fmt.Println(formattedOutput)

	// Generar el archivo DOT
	dotFilePath := "reports/ast.dot"
	err := GenerateDotFromFormattedTreeString(formattedOutput, dotFilePath)
	if err != nil {
		fmt.Println("Error al generar el archivo DOT:", err)
	} else {
		pdfFilePath := "reports/ast.pdf"
		cmd := exec.Command("dot", "-Tpdf", dotFilePath, "-o", pdfFilePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error al convertir DOT a PDF:", err)
		}
		fmt.Printf("Archivo PDF '%s' generado exitosamente.\n", pdfFilePath)
	}

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

	erri := visitor.TablaSimbolos.GenerarReporteHTML("reports/tabla_simbolos.html")
	if err != nil {
		fmt.Println("Error al generar reporte:", erri)
	}
	erro := compile.GenerarReporteHTML(visitor.Errores, "reports/reporte_errores.html")
	if erro != nil {
		fmt.Println("Error al generar reporte:", erro)
	}
	fmt.Println(" ----------------------- Ejecutamos ARM -----------------------")
	vistorARM := generatorARM.NewCompileARMVisitor()
	tree.Accept(vistorARM)
	fmt.Println(" ----------------------- ARM generado -----------------------")

	return visitor.Salida
}
