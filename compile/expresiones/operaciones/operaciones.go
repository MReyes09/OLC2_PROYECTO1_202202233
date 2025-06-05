package operaciones

import (
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type OperacionesVisitor struct {
	Salida *string // puntero a salida para modificar directamente
	Visit  func(antlr.ParseTree) interface{}
}

func NewOperacionesVisitor(salida *string, visitFunc func(antlr.ParseTree) interface{}) *OperacionesVisitor {
	return &OperacionesVisitor{
		Salida: salida,
		Visit:  visitFunc,
	}
}

// ------------------------------ OPERADORES ARITMÉTICOS -----------------------------
// VisitNegate
func (ov *OperacionesVisitor) VisitNegate(ctx *gramAntlr.NegateContext) interface{} {
	value := ov.Visit(ctx.Expr())

	if intVal, ok := value.(int); ok {
		return -intVal
	}

	if floatVal, ok := value.(float64); ok {
		return -floatVal
	}

	*ov.Salida += "Error: solo se pueden negar números"
	return nil
}

// VisitMulDivModulo
func (ov *OperacionesVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	left := ov.Visit(ctx.Expr(0))
	right := ov.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	// Validación de tipos
	_, leftIsInt := left.(int)
	_, leftIsFloat := left.(float64)
	_, rightIsInt := right.(int)
	_, rightIsFloat := right.(float64)

	if !(leftIsInt || leftIsFloat) || !(rightIsInt || rightIsFloat) {
		*ov.Salida += fmt.Sprintf("Error semántico: no se pueden operar '%s' y '%s' con el operador '%s'.",
			fmt.Sprintf("%ov", left), fmt.Sprintf("%ov", right), op)
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
				*ov.Salida += "Error: división por cero"
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
				*ov.Salida += "Error: división por cero"
				return nil
			}
			return l / r
		default:
			*ov.Salida += "Error: operador no válido para floats (solo * y / permitidos)"
			return nil
		}
	}

	*ov.Salida += "Error en operador multiplicativo"
	return nil
}

// VisitAddSub
func (ov *OperacionesVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {
	left := ov.Visit(ctx.Expr(0))
	right := ov.Visit(ctx.Expr(1))
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

	*ov.Salida += "Error en operador aritmetic"
	return nil
}

// ------------------------------ FIN DE EXPR ARITMÉTICAS -----------------------------

// ------------------------------- OPERADORES LOGICAS -----------------------------
// VisitEqualsNotEquals
func (ov *OperacionesVisitor) VisitEqualsNotEquals(ctx *gramAntlr.EqualsNotEqualsContext) interface{} {
	left := ov.Visit(ctx.Expr(0))
	right := ov.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	default:
		*ov.Salida += "Operador de comparación desconocido: " + op
		return nil
	}
}

// VisitMinorMajorEqual
func (ov *OperacionesVisitor) VisitMinorMajorEqual(ctx *gramAntlr.MinorMajorEqualContext) interface{} {
	left := ov.Visit(ctx.Expr(0))
	right := ov.Visit(ctx.Expr(1))
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
		*ov.Salida += fmt.Sprintf("Error: comparación no soportada entre %T y %T", left, right)
		return nil
	}

	*ov.Salida += "Operador de comparación no reconocido: " + op
	return nil
}

// VisitLogical
func (ov *OperacionesVisitor) VisitLogical(ctx *gramAntlr.LogicalContext) interface{} {
	left := ov.Visit(ctx.Expr(0))
	right := ov.Visit(ctx.Expr(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	lBool, lok := left.(bool)
	rBool, rok := right.(bool)

	if !lok || !rok {
		*ov.Salida += fmt.Sprintf("Error: operación lógica %s requiere booleanos", op)
		return nil
	}

	switch op {
	case "&&":
		return lBool && rBool
	case "||":
		return lBool || rBool
	default:
		*ov.Salida += "Operador lógico no reconocido: " + op
		return nil
	}
}

// VisitNot
func (ov *OperacionesVisitor) VisitNot(ctx *gramAntlr.NotContext) interface{} {
	val := ov.Visit(ctx.Expr())
	boolVal, ok := val.(bool)
	if !ok {
		*ov.Salida += "Error: operador ! requiere un valor booleano"
		return nil
	}
	return !boolVal
}
