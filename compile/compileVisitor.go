package compile

import (
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
)

// Esta sería tu "clase" CompilerVisitor
type CompilerVisitor struct {
	*gramAntlr.BasegramaticaVisitor        // composición, como si heredara
	Salida                          string // lo que quieras almacenar
}

// Constructor opcional
func NewCompilerVisitor() *CompilerVisitor {
	return &CompilerVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Salida:               "",
	}
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
	for _, expr := range ctx.AllExpr() {
		value := v.Visit(expr)
		if value == nil {
			continue
		}
		v.Salida += fmt.Sprintf("%v ", value) // Agrega un espacio entre valores
	}
	return nil
}

func (v *CompilerVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	for _, expr := range ctx.AllExpr() {
		value := v.Visit(expr)
		if value == nil {
			continue
		}
		v.Salida += fmt.Sprintf("%v ", value) // Agrega un espacio entre valores
	}
	v.Salida += "\n" // Añade un salto de línea al final
	return nil
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
	value := v.Visit(ctx.Expr())

	if intVal, ok := value.(int); ok {
		return -intVal
	}

	if floatVal, ok := value.(float64); ok {
		return -floatVal
	}

	v.Salida += "Error: solo se pueden negar números"
	return nil
}

func (v *CompilerVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	// Validación de tipos
	_, leftIsInt := left.(int)
	_, leftIsFloat := left.(float64)
	_, rightIsInt := right.(int)
	_, rightIsFloat := right.(float64)

	if !(leftIsInt || leftIsFloat) || !(rightIsInt || rightIsFloat) {
		v.Salida += fmt.Sprintf("Error semántico: no se pueden operar '%s' y '%s' con el operador '%s'.",
			fmt.Sprintf("%v", left), fmt.Sprintf("%v", right), op)
		return nil
	}

	// Si ambos son enteros
	if l, ok := left.(int); ok {
		r := right.(int)
		switch op {
		case "*":
			return l * r
		case "/":
			if r == 0 {
				v.Salida += "Error: división por cero"
				return nil
			}
			return l / r
		case "%":
			return l % r
		}
	}

	// Si ambos son flotantes
	if l, ok := left.(float64); ok {
		r := right.(float64)
		switch op {
		case "*":
			return l * r
		case "/":
			if r == 0 {
				v.Salida += "Error: división por cero"
				return nil
			}
			return l / r
		default:
			v.Salida += "Error: operador no válido para floats (solo * y / permitidos)"
			return nil
		}
	}

	v.Salida += "Error en operador multiplicativo"
	return nil
}

// VisitAddSub
func (v *CompilerVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch l := left.(type) {
	case int:
		r := right.(int)
		switch op {
		case "+":
			return l + r
		case "-":
			return l - r
		}
	case float64:
		r := right.(float64)
		switch op {
		case "+":
			return l + r
		case "-":
			return l - r
		}
	case string:
		if r, ok := right.(string); ok && op == "+" {
			return l + r
		}
	}

	v.Salida += "Error en operador aritmetic"
	return nil
}

// ------------------------------ FIN DE EXPR ARITMÉTICAS -----------------------------

// ------------------------------- OPERADORES LOGICAS -----------------------------
// VisitEqualsNotEquals
func (v *CompilerVisitor) VisitEqualsNotEquals(ctx *gramAntlr.EqualsNotEqualsContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	default:
		v.Salida += "Operador de comparación desconocido: " + op
		return nil
	}
}

// VisitMinorMajorEqual
func (v *CompilerVisitor) VisitMinorMajorEqual(ctx *gramAntlr.MinorMajorEqualContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch l := left.(type) {
	case int:
		r := right.(int)
		switch op {
		case "<":
			return l < r
		case "<=":
			return l <= r
		case ">":
			return l > r
		case ">=":
			return l >= r
		}
	case float64:
		r := right.(float64)
		switch op {
		case "<":
			return l < r
		case "<=":
			return l <= r
		case ">":
			return l > r
		case ">=":
			return l >= r
		}
	default:
		v.Salida += fmt.Sprintf("Error: comparación no soportada entre %T y %T", left, right)
		return nil
	}

	v.Salida += "Operador de comparación no reconocido: " + op
	return nil
}

// VisitLogical
func (v *CompilerVisitor) VisitLogical(ctx *gramAntlr.LogicalContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	lBool, lok := left.(bool)
	rBool, rok := right.(bool)

	if !lok || !rok {
		v.Salida += fmt.Sprintf("Error: operación lógica %s requiere booleanos", op)
		return nil
	}

	switch op {
	case "&&":
		return lBool && rBool
	case "||":
		return lBool || rBool
	default:
		v.Salida += "Operador lógico no reconocido: " + op
		return nil
	}
}

// VisitNot
func (v *CompilerVisitor) VisitNot(ctx *gramAntlr.NotContext) interface{} {
	val := v.Visit(ctx.Expr())
	boolVal, ok := val.(bool)
	if !ok {
		v.Salida += "Error: operador ! requiere un valor booleano"
		return nil
	}
	return !boolVal
}

// ------------------------------- FIN OPERADORES LOGICOS -----------------------------
