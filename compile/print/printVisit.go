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
		fmt.Println("Valor obtenido:", value)
		if value == nil {
			continue
		}
		// Si es slice, lo mostramos bonito
		if slice, ok := value.([]interface{}); ok {
			*pv.Salida += "{" + getStringSlice(slice, "") + "}"
		} else {
			*pv.Salida += fmt.Sprintf("%v", value)
		}
		*pv.Salida += " "
	}
	if newline {
		*pv.Salida += "\n"
	}
}

func getStringSlice(lista []interface{}, cadena string) string {
	for _, item := range lista {
		switch v := item.(type) {
		case []interface{}:
			cadena += "\n\t{"
			cadena = getStringSlice(v, cadena)
			cadena += " },\n"
		default:
			cadena += " " + fmt.Sprintf("%v", v)
		}
	}
	return cadena
}
