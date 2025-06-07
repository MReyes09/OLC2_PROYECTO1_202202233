package nativas

import (
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type NativasVisitor struct {
	Salida *string // Puntero a salida para modificar directamente
	Visit  func(antlr.ParseTree) interface{}
}

func NewNativasVisitor(salida *string, visitFunc func(antlr.ParseTree) interface{}) *NativasVisitor {
	return &NativasVisitor{
		Salida: salida,
		Visit:  visitFunc,
	}
}

func (nv *NativasVisitor) VisitIntToString(ctx *gramAntlr.IntToStringContext) interface{} {
	value := nv.Visit(ctx.Expr())

	if strVal, ok := value.(string); ok && strVal != "" {
		numero, err := strconv.Atoi(strVal)
		if err != nil {
			if _, ok := err.(*strconv.NumError); ok {
				if err.(*strconv.NumError).Err == strconv.ErrRange {
					*nv.Salida += "Error-semántico: el valor está fuera del rango permitido para un entero.\n"
				} else {
					*nv.Salida += "Error-semántico: el valor no tiene un formato válido para ser convertido a entero.\n"
				}
			} else {
				*nv.Salida += fmt.Sprintf("Error inesperado al convertir a entero: %v\n", err)
			}
			return nil
		}
		return numero
	}

	*nv.Salida += "Error-semántico: el valor no es un string válido.\n"
	return nil
}

func (nv *NativasVisitor) VisitFloatToString(ctx *gramAntlr.FloatToStringContext) interface{} {
	value := nv.Visit(ctx.Expr())

	if strVal, ok := value.(string); ok && strVal != "" {
		numero, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok {
				if numErr.Err == strconv.ErrRange {
					*nv.Salida += "Error-semántico: el valor está fuera del rango permitido para un flotante.\n"
				} else {
					*nv.Salida += "Error-semántico: el valor no tiene un formato válido para ser convertido a flotante.\n"
				}
			} else {
				*nv.Salida += fmt.Sprintf("Error inesperado al convertir a flotante: %v\n", err)
			}
			return nil
		}
		return numero
	}

	*nv.Salida += "Error-semántico: el valor no es un string válido.\n"
	return nil
}

func (nv *NativasVisitor) VisitReflectType(ctx *gramAntlr.ReflectTypeContext) interface{} {
	value := nv.Visit(ctx.Expr())

	switch v := value.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case rune:
		return "rune"
	case string:
		return "string"
	//case map[string]interface{}:
	//	return "struct"
	case []interface{}:
		return v
	default:
		return "desconocido"
	}
}
