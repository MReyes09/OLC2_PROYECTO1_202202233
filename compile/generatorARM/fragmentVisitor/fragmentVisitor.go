package fragmentvisitor

import (
	"OLC2CLIENTE/gramatica/gramAntlr"

	"github.com/antlr4-go/antlr/v4"
)

// Parte del código que maneja los fragmentos de código en ARM
type FragmentElement struct {
	Name   string
	Offset int
}

type FragmentVisitor struct {
	*gramAntlr.BasegramaticaVisitor
	Fragment    []FragmentElement
	LocalOffSet int
	BaseOffSet  int
}

func NewFragmentVisitor(baseOffset int) *FragmentVisitor {
	f := &FragmentVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Fragment:             make([]FragmentElement, 0),
		LocalOffSet:          0,
		BaseOffSet:           baseOffset,
	}
	return f
}

// Implementación del método Visit para el FragmentVisitor
func (f *FragmentVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(f)
}

// -------------------------- DECLARACION DE VARIABLES --------------------------
// Instrucciones: VarDeclStmt
func (f *FragmentVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	f.Visit(ctx.VarDcl())
	return nil
}

// varDcl: 'mut'? ID_VARIABLE ':=' expr           # VarDclWithInference
func (f *FragmentVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	varName := ctx.ID_VARIABLE().GetText()
	f.Fragment = append(f.Fragment, FragmentElement{
		Name:   varName,
		Offset: f.LocalOffSet + f.BaseOffSet,
	})
	f.LocalOffSet += 1
	f.Visit(ctx.Expr())
	return nil
}

func (f *FragmentVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {
	varName := ctx.ID_VARIABLE().GetText()
	f.Fragment = append(f.Fragment, FragmentElement{
		Name:   varName,
		Offset: f.LocalOffSet + f.BaseOffSet,
	})
	f.LocalOffSet += 1
	f.Visit(ctx.Expr())
	return nil
}

func (f *FragmentVisitor) VisitVarDclWithTypeOnly(ctx *gramAntlr.VarDclWithTypeOnlyContext) interface{} {
	varName := ctx.ID_VARIABLE().GetText()
	f.Fragment = append(f.Fragment, FragmentElement{
		Name:   varName,
		Offset: f.LocalOffSet + f.BaseOffSet,
	})
	f.LocalOffSet += 1
	return nil
}

func (f *FragmentVisitor) VisitPrintStmt(ctx *gramAntlr.PrintStmtContext) interface{} {
	//Visitamos imprimir
	f.Visit(ctx.Imprimir())
	return nil
}

func (f *FragmentVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	// recorremos todas las expresiones dentro del println
	for _, expr := range ctx.AllExpr() {
		// recuerda aumentar el offset de la base
		f.Visit(expr)
	}
	return nil
}

func (f *FragmentVisitor) VisitString(ctx *gramAntlr.StringContext) interface{} {
	// Aquí podrías manejar la cadena literal si es necesario
	// LLegamos a un terminal! ya podemos aumentar el offset pues string necesita su espacio
	f.LocalOffSet += 1

	return nil
}
