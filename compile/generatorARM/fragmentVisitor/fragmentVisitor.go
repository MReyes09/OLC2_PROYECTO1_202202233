package fragmentvisitor

import (
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Parte del código que maneja los fragmentos de código en ARM
type FragmentElement struct {
	Name   string
	Offset int
}

type FragmentVisitor struct {
	*gramAntlr.BasegramaticaVisitor
	Fragment      []FragmentElement
	LocalOffSet   int
	BaseOffSet    int
	Depth         int // Profundidad de anidamiento, si es necesario
	UsedPositions []int
	MaxPosition   int
}

func NewFragmentVisitor(baseOffset int) *FragmentVisitor {
	f := &FragmentVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Fragment:             make([]FragmentElement, 0),
		LocalOffSet:          0,
		BaseOffSet:           baseOffset,
		Depth:                -1, // Inicialmente no hay anidamiento
		UsedPositions:        make([]int, 0),
		MaxPosition:          0,
	}
	return f
}

// Implementación del método Visit para el FragmentVisitor
func (f *FragmentVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(f)
}

// ------------------ DECLARACIONES Y ASIGNACIONES ---------------------------------

// VisitVarDcl
func (f *FragmentVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	f.Visit(ctx.VarDcl())
	//f.cleanOffset()
	return nil
}

// varDcl: 'mut'? ID_VARIABLE ':=' expr           # VarDclWithInference
func (f *FragmentVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	varName := ctx.ID_VARIABLE().GetText()
	fmt.Println("Visitando VarDclWithInference:", varName)
	// Primero visitamos la expresión, para que incremente el offset si es necesario
	f.Visit(ctx.Expr())

	// Luego asignamos el offset de la variable (después de todo lo que la expresión haya usado)
	f.Fragment = append(f.Fragment, FragmentElement{
		Name:   varName,
		Offset: f.LocalOffSet + f.BaseOffSet,
	})
	f.LocalOffSet += 1 // para la propia variable
	return nil
}

func (f *FragmentVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {

	varName := ctx.ID_VARIABLE().GetText()

	f.Visit(ctx.Expr())

	f.Fragment = append(f.Fragment, FragmentElement{
		Name:   varName,
		Offset: f.LocalOffSet + f.BaseOffSet,
	})
	f.LocalOffSet += 1
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

// ------------------------- PRINT Y PRINTLN ---------------------------------

func (f *FragmentVisitor) VisitPrintStmt(ctx *gramAntlr.PrintStmtContext) interface{} {
	//Visitamos imprimir
	f.Visit(ctx.Imprimir())
	//f.ResetUsedPositions()
	return nil
}

func (f *FragmentVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	// recorremos todas las expresiones dentro del println
	for _, expr := range ctx.AllExpr() {
		// recuerda aumentar el offset de la base
		f.Visit(expr)
		f.UsedPositions = append(f.UsedPositions, f.LocalOffSet)
	}
	return nil
}

// --------------------------------- TIPOS PRIMITIVOS ---------------------------------
func (f *FragmentVisitor) VisitString(ctx *gramAntlr.StringContext) interface{} {
	if f.Depth >= 0 {
		f.LocalOffSet += 1
	}
	return nil
}

func (f *FragmentVisitor) VisitInteger(ctx *gramAntlr.IntegerContext) interface{} {
	if f.Depth >= 0 {
		f.LocalOffSet += 1
	}
	return nil
}

func (f *FragmentVisitor) VisitBoolean(ctx *gramAntlr.BooleanContext) interface{} {
	if f.Depth >= 0 {
		f.LocalOffSet += 1
	}
	return nil
}

func (f *FragmentVisitor) VisitDouble(ctx *gramAntlr.DoubleContext) interface{} {
	if f.Depth >= 0 {
		f.LocalOffSet += 1
	}
	return nil
}

// --------------------------------- IDENTIFICADORES ---------------------------------
func (f *FragmentVisitor) VisitIdentifier(ctx *gramAntlr.IdentifierContext) interface{} {
	if f.Depth >= 0 {
		f.LocalOffSet += 1
	}
	return nil
}

// --------------------------------- OPERACIONES ---------------------------------

func (f *FragmentVisitor) VisitParens(ctx *gramAntlr.ParensContext) interface{} {
	f.Visit(ctx.Expr())
	return nil
}

func (f *FragmentVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {
	f.Depth += 1

	f.Visit(ctx.Expr(0)) // izquierda
	f.Visit(ctx.Expr(1)) // derecha

	// Se agrega un frame porque es una operacion y la suma de las expresiones debe guardarse
	f.LocalOffSet += 1

	f.Depth -= 1
	return nil
}

func (f *FragmentVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	f.Depth += 1

	f.Visit(ctx.Expr(0)) // izquierda
	f.Visit(ctx.Expr(1)) // derecha

	// Se agrega un frame porque es una operacion y la multiplicacion de las expresiones debe guardarse
	f.LocalOffSet += 1

	f.Depth -= 1
	return nil
}

func (f *FragmentVisitor) VisitMinorMajorEqual(ctx *gramAntlr.MinorMajorEqualContext) interface{} {
	f.Depth += 1

	f.Visit(ctx.Expr(0)) // izquierda
	f.Visit(ctx.Expr(1)) // derecha

	// Se agrega un frame porque es una operacion y la suma de las expresiones debe guardarse
	f.LocalOffSet += 1

	f.Depth -= 1
	return nil
}

func (f *FragmentVisitor) VisitEqualsNotEquals(ctx *gramAntlr.EqualsNotEqualsContext) interface{} {
	f.Depth += 1

	f.Visit(ctx.Expr(0)) // izquierda
	f.Visit(ctx.Expr(1)) // derecha

	// Se agrega un frame porque es una operacion y la suma de las expresiones debe guardarse
	f.LocalOffSet += 1

	f.Depth -= 1
	return nil
}

// Auxiliar

/*func (f *FragmentVisitor) cleanOffset() {
	if len(f.Fragment) == 0 {
		panic("No hay fragmentos para limpiar el offset")
	}
	f.LocalOffSet = f.LocalOffSet - f.MaxPosition
	f.Fragment[len(f.Fragment)-1].Offset = f.LocalOffSet
}

func (f *FragmentVisitor) ResetUsedPositions() {
	if len(f.UsedPositions) > 0 {
		f.LocalOffSet = f.LocalOffSet - len(f.UsedPositions) + 1
		f.UsedPositions = make([]int, 0)
	}
}
*/
