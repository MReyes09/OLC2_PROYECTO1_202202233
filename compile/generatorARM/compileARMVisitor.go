package generatorARM

import (
	"OLC2CLIENTE/compile/generatorARM/traductor"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CompileARMVisitor struct {
	*gramAntlr.BasegramaticaVisitor
}

type DataFuncion struct {
	frameSize int
	typeRet   *traductor.TypeObject
}

func NewCompileARMVisitor() *CompileARMVisitor {
	v := &CompileARMVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
	}
	return v
}

func (v *CompileARMVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *CompileARMVisitor) VisitInicio(context *gramAntlr.InicioContext) interface{} {
	fmt.Println("Visitando Inicio desde CompileARMVisitor")
	return nil
}
