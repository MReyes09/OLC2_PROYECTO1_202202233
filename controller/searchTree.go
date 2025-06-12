package controller

import (
	"OLC2CLIENTE/gramatica/gramAntlr"

	"github.com/antlr4-go/antlr/v4"
)

type SearchTree struct {
	*gramAntlr.BasegramaticaVisitor // <-- importante
	FunctMain                       []*gramAntlr.FunctionStmtContext
	FunctStruct                     []*gramAntlr.FunctionStructStmtContext
	Funciones                       []*gramAntlr.FunctionStmtContext
	DeclaracionesSimples            []*gramAntlr.VarDeclStmtContext
	DeclaracionesArreglos           []*gramAntlr.VarDeclSliceStmtContext
	DeclaracionesStructs            []*gramAntlr.VarDeclStructStmtContext
	DeclaracionesStructs2           []*gramAntlr.VarStructDclStmtContext
	Asignaciones                    []*gramAntlr.AsignStmtContext
}

// Constructor opcional
func NewSearchTree() *SearchTree {
	return &SearchTree{
		BasegramaticaVisitor:  &gramAntlr.BasegramaticaVisitor{},
		FunctMain:             []*gramAntlr.FunctionStmtContext{},
		FunctStruct:           []*gramAntlr.FunctionStructStmtContext{},
		Funciones:             []*gramAntlr.FunctionStmtContext{},
		DeclaracionesSimples:  []*gramAntlr.VarDeclStmtContext{},
		DeclaracionesArreglos: []*gramAntlr.VarDeclSliceStmtContext{},
		DeclaracionesStructs:  []*gramAntlr.VarDeclStructStmtContext{},
		DeclaracionesStructs2: []*gramAntlr.VarStructDclStmtContext{},
		Asignaciones:          []*gramAntlr.AsignStmtContext{},
	}
}

// Implementaciones de los métodos Visit

func (s *SearchTree) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(s)
}

// Sobrescribiendo un método del visitor
func (s *SearchTree) VisitInicio(ctx *gramAntlr.InicioContext) interface{} {
	for _, stmtCtx := range ctx.AllInstrucciones() {
		s.Visit(stmtCtx)
	}
	return nil
}

func (s *SearchTree) VisitFunctionStmt(ctx *gramAntlr.FunctionStmtContext) interface{} {
	funcion := ctx.Functions()
	nombre := funcion.GetChild(1).(antlr.ParseTree).GetText()
	if nombre == "main" {
		s.FunctMain = append(s.FunctMain, ctx)
	} else {
		s.Funciones = append(s.Funciones, ctx)
	}
	return nil
}

func (s *SearchTree) VisitFunctionStructStmt(ctx *gramAntlr.FunctionStructStmtContext) interface{} {
	s.FunctStruct = append(s.FunctStruct, ctx)
	return nil
}

func (s *SearchTree) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	s.DeclaracionesSimples = append(s.DeclaracionesSimples, ctx)
	return nil
}

func (s *SearchTree) VisitVarDeclSliceStmt(ctx *gramAntlr.VarDeclSliceStmtContext) interface{} {
	s.DeclaracionesArreglos = append(s.DeclaracionesArreglos, ctx)
	return nil
}

func (s *SearchTree) VisitVarDeclStructStmt(ctx *gramAntlr.VarDeclStructStmtContext) interface{} {
	s.DeclaracionesStructs = append(s.DeclaracionesStructs, ctx)
	return nil
}

func (s *SearchTree) VisitVarStructDclStmt(ctx *gramAntlr.VarStructDclStmtContext) interface{} {
	s.DeclaracionesStructs2 = append(s.DeclaracionesStructs2, ctx)
	return nil
}

func (s *SearchTree) VisitAsignStmt(ctx *gramAntlr.AsignStmtContext) interface{} {
	s.Asignaciones = append(s.Asignaciones, ctx)
	return nil
}
