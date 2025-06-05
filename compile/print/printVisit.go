package print

import (
	"fmt"

	"OLC2CLIENTE/gramatica/gramAntlr"

	"github.com/antlr4-go/antlr/v4"
)

// PrintVisitor es una estructura auxiliar para manejar impresión
type PrintVisitor struct {
	Salida *string // puntero a salida para modificar directamente
	Visit  func(antlr.ParseTree) interface{}
}

// Constructor
func NewPrintVisitor(salida *string, visitFunc func(antlr.ParseTree) interface{}) *PrintVisitor {
	return &PrintVisitor{
		Salida: salida,
		Visit:  visitFunc,
	}
}

func (pv *PrintVisitor) VisitPrint(ctx *gramAntlr.PrintContext) interface{} {
	pv.handlePrintExprs(ctx.AllExpr(), false)
	return nil
}

func (pv *PrintVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	pv.handlePrintExprs(ctx.AllExpr(), true)
	return nil
}

func (pv *PrintVisitor) handlePrintExprs(exprs []gramAntlr.IExprContext, newline bool) {
	for _, expr := range exprs {
		value := pv.Visit(expr)
		if value == nil {
			continue
		}
		*pv.Salida += fmt.Sprintf("%v ", value)
	}
	if newline {
		*pv.Salida += "\n"
	}
}
