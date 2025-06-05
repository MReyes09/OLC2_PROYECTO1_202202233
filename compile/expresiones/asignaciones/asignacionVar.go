package asignaciones

import (
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type AsignacionesVisitor struct {
	Salida *string
	Visit  func(antlr.ParseTree) interface{}
}

func NewAsignacionesVisitor(salida *string, visitFunc func(antlr.ParseTree) interface{}) *AsignacionesVisitor {
	return &AsignacionesVisitor{
		Salida: salida,
		Visit:  visitFunc,
	}
}

func (av *AsignacionesVisitor) VisitVarExpr(ctx *gramAntlr.VarExprContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := av.Visit(ctx.Expr())
	fmt.Println("Asignando valor a la variable:", id, "con valor:", value)

	return nil
}

/*
public override object VisitVarExpr(gramaticaParser.VarExprContext context)
    {
        string id = context.ID_VARIABLE().GetText();
        object value = Visit(context.expr());
        SymbolType type = currentEnvironment.GetVariable(id).Type;
        bool mutabilidad = currentEnvironment.GetVariable(id).Mutable;

        var context_Start = context.Start;

        if (value is null)
        {
            throw new ErrorSemantico("Error-semántico: al asignar el valor a la variable, el valor es nulo.", context_Start);
        }

        if (value is List<Object> tempList)
        {
            if (IsValidType(tempList[0], type) || value is List<Object>)
            {
                currentEnvironment.SetVariable(id, value, type, mutabilidad, false, context_Start);
                return null;
            }
            else
            {
                throw new ErrorSemantico("Error-semántico: al asignar el valor a la variable, los tipos no son compatibles.", context.Start);
            }
        }

        if (!IsValidType(value, type))
        {
            if (mutabilidad)
            {
                currentEnvironment.SetVariable(id, value, type, mutabilidad, false, context_Start);
                return null;
            }
            throw new ErrorSemantico("Error-semántico: al asignar el valor a la variable, los tipos no son compatibles.", context.Start);
        }

        currentEnvironment.SetVariable(id, value, type, mutabilidad, false, context_Start);

        return null;
    }
*/
