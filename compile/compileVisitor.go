package compile

import (
	"OLC2CLIENTE/compile/expresiones/operaciones"
	"OLC2CLIENTE/compile/print"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
)

// Esta sería tu "clase" CompilerVisitor
type CompilerVisitor struct {
	*gramAntlr.BasegramaticaVisitor        // composición, como si heredara
	Salida                          string // lo que quieras almacenar
	printVisitor                    *print.PrintVisitor
	operacionesVisitor              *operaciones.OperacionesVisitor // Visitor para operaciones aritméticas
}

// Constructor opcional
func NewCompilerVisitor() *CompilerVisitor {
	v := &CompilerVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Salida:               "",
	}
	// Inicializa el printVisitor pasándole la función Visit y la referencia a la salida
	v.printVisitor = print.NewPrintVisitor(&v.Salida, v.Visit)
	v.operacionesVisitor = operaciones.NewOperacionesVisitor(&v.Salida, v.Visit)
	return v
}

func (v *CompilerVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// Sobrescribiendo un método del visitor
func (v *CompilerVisitor) VisitInicio(ctx *gramAntlr.InicioContext) interface{} {
	for _, stmtCtx := range ctx.AllInstrucciones() {
		v.Visit(stmtCtx)
	}
	return nil
}

// Visita de Imprimir y sus print - println
func (v *CompilerVisitor) VisitPrintStmt(ctx *gramAntlr.PrintStmtContext) interface{} {
	v.Visit(ctx.Imprimir())
	return nil
}

func (v *CompilerVisitor) VisitPrint(ctx *gramAntlr.PrintContext) interface{} {
	return v.printVisitor.VisitPrint(ctx)
}

func (v *CompilerVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	return v.printVisitor.VisitPrintln(ctx)
}

// ----------------------------- Final de print -----------------------------

// ----------------------------- Expresiones -----------------------------
// VisitParens
func (v *CompilerVisitor) VisitParens(ctx *gramAntlr.ParensContext) interface{} {
	return v.Visit(ctx.Expr()) // Visita la expresión dentro de los paréntesis
}

// ----------------------------- TIPOS DE DATOS -----------------------------
// VisitNumber
func (v *CompilerVisitor) VisitInteger(ctx *gramAntlr.IntegerContext) interface{} {
	numero, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		v.Salida += "Error en la conversión de número: " + ctx.GetText()
		return nil
	} else {
		return numero
	}
}

// VisitFloat
func (v *CompilerVisitor) VisitDouble(ctx *gramAntlr.DoubleContext) interface{} {
	texto := ctx.GetText()
	numero, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		v.Salida += "Error en la conversión de double: " + texto
		return nil
	}
	return numero
}

// VisitString
func (v *CompilerVisitor) VisitString(ctx *gramAntlr.StringContext) interface{} {
	texto := ctx.GetText()
	if strings.HasPrefix(texto, "\"") && strings.HasSuffix(texto, "\"") {
		texto = texto[1 : len(texto)-1] // eliminar comillas
	}

	texto = strings.ReplaceAll(texto, `\n`, "\n")
	texto = strings.ReplaceAll(texto, `\t`, "\t")
	texto = strings.ReplaceAll(texto, `\r`, "\r")
	texto = strings.ReplaceAll(texto, `\"`, `"`)
	texto = strings.ReplaceAll(texto, `\\`, `\`)

	return texto
}

// VisitChar - Run
func (v *CompilerVisitor) VisitChar(ctx *gramAntlr.CharContext) interface{} {
	texto := ctx.GetText() // e.g. `'a'`
	texto = strings.Trim(texto, "'")
	r, _ := utf8.DecodeRuneInString(texto)
	return r
}

// VisitNull
func (v *CompilerVisitor) VisitNil(ctx *gramAntlr.NilContext) interface{} {
	return "nil"
}

// VisitBoolean
func (v *CompilerVisitor) VisitBoolean(ctx *gramAntlr.BooleanContext) interface{} {
	texto := ctx.GetText() // e.g. "true" o "false"
	if texto == "true" {
		return true
	} else if texto == "false" {
		return false
	}
	v.Salida += "Error en el valor booleano: " + texto
	return nil
}

// ------------------------------ FIN TIPOS DE DATOS -----------------------------

// ------------------------------ OPERADORES ARITMÉTICOS -----------------------------
// VisitNegate
func (v *CompilerVisitor) VisitNegate(ctx *gramAntlr.NegateContext) interface{} {
	return v.operacionesVisitor.VisitNegate(ctx)
}

// VisitMulDivModulo
func (v *CompilerVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	return v.operacionesVisitor.VisitMulDivModulo(ctx)
}

// VisitAddSub
func (v *CompilerVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {
	return v.operacionesVisitor.VisitAddSub(ctx)
}

// ------------------------------ FIN DE EXPR ARITMÉTICAS -----------------------------

// ------------------------------- OPERADORES LOGICAS -----------------------------
// VisitEqualsNotEquals
func (v *CompilerVisitor) VisitEqualsNotEquals(ctx *gramAntlr.EqualsNotEqualsContext) interface{} {
	return v.operacionesVisitor.VisitEqualsNotEquals(ctx)
}

// VisitMinorMajorEqual
func (v *CompilerVisitor) VisitMinorMajorEqual(ctx *gramAntlr.MinorMajorEqualContext) interface{} {
	return v.operacionesVisitor.VisitMinorMajorEqual(ctx)
}

// VisitLogical
func (v *CompilerVisitor) VisitLogical(ctx *gramAntlr.LogicalContext) interface{} {
	return v.operacionesVisitor.VisitLogical(ctx)
}

// VisitNot
func (v *CompilerVisitor) VisitNot(ctx *gramAntlr.NotContext) interface{} {
	return v.operacionesVisitor.VisitNot(ctx)
}

// ------------------------------- FIN OPERADORES LOGICOS -----------------------------
