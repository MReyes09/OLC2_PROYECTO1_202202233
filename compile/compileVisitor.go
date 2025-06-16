package compile

import (
	"OLC2CLIENTE/compile/expresiones/nativas"
	"OLC2CLIENTE/compile/expresiones/operaciones"
	"OLC2CLIENTE/compile/print"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"
	"reflect"
	"regexp"
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
	funcionesNativasVisitor         *nativas.NativasVisitor         // Visitor para funciones nativas
	currentEnv                      *Environment                    // scope
	conditionExpr                   interface{}                     // Added for switch statement
	StructRelational                map[string][]string
	Errores                         []ErrorReport
	TablaSimbolos                   *TablaSimbolos
}

// Constructor opcional
func NewCompilerVisitor() *CompilerVisitor {
	v := &CompilerVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Salida:               "",
		currentEnv:           NewEnvironment(nil),       // Entorno raíz
		StructRelational:     make(map[string][]string), // <--- NUEVO campo agregado
		TablaSimbolos:        NewTablaSimbolos(),
	}
	// Inicializa el printVisitor pasándole la función Visit y la referencia a la salida
	v.printVisitor = print.NewPrintVisitor(&v.Salida, v.Visit)
	v.operacionesVisitor = operaciones.NewOperacionesVisitor(&v.Salida, v.Visit)
	v.funcionesNativasVisitor = nativas.NewNativasVisitor(&v.Salida, v.Visit)
	return v
}

// ReportScope imprime el contenido del scope actual
func (v *CompilerVisitor) ReportScope() string {
	return v.currentEnv.ImprimirScope()
}

// ReportFunctions imprime las funciones globales
func (v *CompilerVisitor) ReportFunctions() string {
	return v.currentEnv.ImprimirScopeFunc()
}

// -----------------------
//------------------------------------------------------------------------------------------

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
	v.handlePrintExprs(ctx.AllExpr(), false)
	return nil
}

func (v *CompilerVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	v.handlePrintExprs(ctx.AllExpr(), true)
	return nil
}

func (pv *CompilerVisitor) handlePrintExprs(exprs []gramAntlr.IExprContext, newline bool) {
	for _, expr := range exprs {
		value := pv.Visit(expr)
		if value == nil {
			continue
		}
		// Si es slice, lo mostramos bonito
		if slice, ok := value.([]interface{}); ok {
			pv.Salida += "{" + getStringSlice(slice, "") + "}"
		} else if valStruct, ok := value.(map[string]Symbol); ok {
			pv.Salida += "{ "
			first := true
			for key, sym := range valStruct {
				if !first {
					pv.Salida += ", "
				}
				pv.Salida += fmt.Sprintf("%s: %v", key, sym.Value)
				first = false
			}
			pv.Salida += "}"
		} else {
			pv.Salida += fmt.Sprintf("%v", value)
		}
		pv.Salida += " "
	}
	if newline {
		pv.Salida += "\n"
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

// ----------------------------- Final de print --------------.---------------

// ----------------------------- Manejo de Errores ---------------------------
func (v *CompilerVisitor) AgregarError(mensaje string, linea, columna int, tipo string) {
	v.Errores = append(v.Errores, ErrorReport{
		No:      len(v.Errores) + 1,
		Mensaje: mensaje,
		Linea:   linea,
		Columna: columna,
		Tipo:    tipo,
	})
}

// ----------------------------- Expresiones -----------------------------
// VisitParens
func (v *CompilerVisitor) VisitParens(ctx *gramAntlr.ParensContext) interface{} {
	return v.Visit(ctx.Expr()) // Visita la expresión dentro de los paréntesis
}

// ----------------------------- TIPOS DE DATOS -----------------------------
// VisitNumber
func (v *CompilerVisitor) VisitInteger(ctx *gramAntlr.IntegerContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "int"           // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
	numero, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		v.Salida += "Error en la conversión de número: " + ctx.GetText()
		v.AgregarError(
			"Error en la conversión de número",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	} else {
		return numero
	}
}

// VisitFloat
func (v *CompilerVisitor) VisitDouble(ctx *gramAntlr.DoubleContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "float"         // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
	texto := ctx.GetText()
	numero, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		v.Salida += "Error en la conversión de double: " + texto
		v.AgregarError(
			"Error en la conversión de double",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	return numero
}

// VisitString
func (v *CompilerVisitor) VisitString(ctx *gramAntlr.StringContext) interface{} {

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "string"        // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
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

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "char"          // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, nil)
	texto := ctx.GetText() // e.g. `'a'`
	texto = strings.Trim(texto, "'")
	r, _ := utf8.DecodeRuneInString(texto)
	fmt.Println("Rune:", r)
	return r
}

// VisitNull
func (v *CompilerVisitor) VisitNil(ctx *gramAntlr.NilContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "nil"           // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
	return "nil"
}

// VisitBoolean
func (v *CompilerVisitor) VisitBoolean(ctx *gramAntlr.BooleanContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "bool"          // Obtén el tipo real
	ambito := "global"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
	texto := ctx.GetText() // e.g. "true" o "false"
	if texto == "true" {
		return true
	} else if texto == "false" {
		return false
	}
	v.Salida += "Error en el valor booleano: " + texto
	v.AgregarError(
		"Error en el valor booleano",
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"semántico",
	)
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

// -------------------------------- DECLARACIONES --------------------------------------
// VisitVarDcl
func (v *CompilerVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	return v.Visit(ctx.VarDcl())
}

// 'mut' ID type '=' expr ';'
func (v *CompilerVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	expr := ctx.Expr()
	value := v.Visit(expr)

	// Convertir tipo string a SymbolType
	symbolType, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		v.AgregarError(
			"Error tipo de dato no valido"+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	// Manejar valores nulos
	if value == nil {
		switch symbolType {
		case INT:
			nombre := ctx.GetText() // Obtén el nombre correctamente
			tipo := "int"           // Obtén el tipo real
			ambito := "global"      // Obtén el ámbito real
			linea := ctx.GetStart().GetLine()
			v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
			value = 0
		case FLOAT64:
			nombre := ctx.GetText() // Obtén el nombre correctamente
			tipo := "float"         // Obtén el tipo real
			ambito := "global"      // Obtén el ámbito real
			linea := ctx.GetStart().GetLine()
			v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
			value = 0.0
		case STRING:
			nombre := ctx.GetText() // Obtén el nombre correctamente
			tipo := "string"        // Obtén el tipo real
			ambito := "global"      // Obtén el ámbito real
			linea := ctx.GetStart().GetLine()
			v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
			value = ""
		case BOOL:
			nombre := ctx.GetText() // Obtén el nombre correctamente
			tipo := "bool"          // Obtén el tipo real
			ambito := "global"      // Obtén el ámbito real
			linea := ctx.GetStart().GetLine()
			v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
			value = false
		case RUNE:
			nombre := ctx.GetText() // Obtén el nombre correctamente
			tipo := "nil"           // Obtén el tipo real
			ambito := "global"      // Obtén el ámbito real
			linea := ctx.GetStart().GetLine()
			v.TablaSimbolos.Insertar(nombre, tipo, ambito, linea, "variable")
			value = '\000'
		}
	}

	// Convertir int a float64 si necesario
	if num, ok := value.(int); ok && symbolType == FLOAT64 {
		value = float64(num)
	} else if !isValidType(value, symbolType) {
		v.Salida += fmt.Sprintf("Error semántico: tipos incompatibles para variable %s\n", id)
		v.AgregarError(
			"Error tipo de dato no valido"+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	miError := v.currentEnv.SetVariable(id, value, symbolType, false, true, ctx.GetStart())
	if miError != nil {
		v.Salida += fmt.Sprintf("Error-semántico: Intento redeclarar la variable %s", id)
		v.AgregarError(
			"Error Intento redeclarar la variable "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	return nil
}

// 'mut' ID type ';'
func (v *CompilerVisitor) VisitVarDclWithTypeOnly(ctx *gramAntlr.VarDclWithTypeOnlyContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	symbolType, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		v.AgregarError(
			"Error tipo no valido"+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	var defaultValue interface{}
	switch symbolType {
	case INT:
		defaultValue = 0
	case FLOAT64:
		defaultValue = 0.0
	case STRING:
		defaultValue = ""
	case BOOL:
		defaultValue = false
	case RUNE:
		defaultValue = '\000'
	case STRUCT:
		return nil
	default:
		v.Salida += fmt.Sprintf("Error: tipo %s no soporta valor por defecto\n", typeStr)
		v.AgregarError(
			"Error tipo no soporta valor por defecto"+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	miError := v.currentEnv.SetVariable(id, defaultValue, symbolType, false, true, ctx.GetStart())
	if miError != nil {
		v.Salida += fmt.Sprintf("Error-semántico: Intento redeclarar la variable %s", id)
		v.AgregarError(
			"Error Intento redeclarar una variable "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	return nil
}

// 'mut' ID ':=' expr ';'
func (v *CompilerVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())
	var symbolType SymbolType
	switch value.(type) {
	case int:
		symbolType = INT
	case float64:
		symbolType = FLOAT64
	case string:
		symbolType = STRING
	case bool:
		symbolType = BOOL
	case rune:
		symbolType = RUNE
	default:
		v.Salida += fmt.Sprintf("Error: tipo no reconocido para variable %s\n", id)
		v.AgregarError(
			"Error tipo no reconocido para variable "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	miError := v.currentEnv.SetVariable(id, value, symbolType, false, true, ctx.GetStart())
	if miError != nil {
		v.Salida += fmt.Sprintf("Error-semántico: Intento redeclarar la variable %s", id)
		v.AgregarError(
			"Error intento redeclarar la variable "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	return nil
}

// ----------------------------- Acceso a arreglos -----------------------------
// VisitArrayAccess - 'ID_VARIABLE' '[' expr ']' '=' expr
func (v *CompilerVisitor) VisitArrayAccess(ctx *gramAntlr.ArrayAccessContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Arreglos"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error variable no declarada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	listaBase, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder a arreglo, variable no es de un arreglo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	exprs := ctx.AllExpr()
	for i := 0; i < len(exprs); i++ {
		index := v.Visit(exprs[i])
		idx, ok := index.(int)
		if !ok {
			v.Salida += "Error-semántico: índice debe ser un entero.\n"
			v.AgregarError(
				"Error el índice debe ser un entero",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}

		// Verificación de rango
		if idx < 0 || idx >= len(listaBase) {
			v.Salida += fmt.Sprintf("Error-semántico: índice %d fuera de rango (tamaño del arreglo: %d)\n", idx, len(listaBase))
			v.AgregarError(
				"Error indice fuera de rango",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}

		if i == len(exprs)-2 {
			// Asignar valor en la posición final
			value := v.Visit(exprs[i+1])
			listaBase[idx] = value
			return nil
		} else {
			// Ir al siguiente nivel del arreglo anidado
			nextLevel, ok := listaBase[idx].([]interface{})
			if !ok {
				v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la posición %d no es un arreglo.\n", idx)
				v.AgregarError(
					"Error al acceder al arreglo, posicion no es un arreglo",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			listaBase = nextLevel
		}
	}

	return nil
}

// VisitArrayAccessSimple - 'ID_VARIABLE' '[' expr ']'
func (v *CompilerVisitor) VisitArrayAccessSimple(ctx *gramAntlr.ArrayAccessSimpleContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Arreglos"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error  variable no declarada",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	listaBase, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, variable no es un arreglo",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	exprs := ctx.AllExpr()
	for i, expr := range exprs {
		index := v.Visit(expr)
		idx, ok := index.(int)
		if !ok {
			v.Salida += "Error-semántico: índice debe ser un entero.\n"
			v.AgregarError(
				"Error el índice debe ser un entero",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}

		// Verificación de rango
		if idx < 0 || idx >= len(listaBase) {
			v.Salida += fmt.Sprintf("Error-semántico: índice %d fuera de rango (tamaño del arreglo: %d)\n", idx, len(listaBase))
			v.AgregarError(
				"Error indice fuera de rango",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}

		if i == len(exprs)-1 {
			// Retornar el valor final del acceso
			return listaBase[idx]
		}

		// Continuar con el siguiente nivel del arreglo anidado
		nextLevel, ok := listaBase[idx].([]interface{})
		if !ok {
			v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la posición %d no es un arreglo.\n", idx)
			v.AgregarError(
				"Error al acceder a arreglo, posicion no es un arreglo",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
		listaBase = nextLevel
	}

	return nil
}

// VisitArrayFindIndex - 'ID_VARIABLE' '.' 'findIndex' '(' expr ')'
func (v *CompilerVisitor) VisitArrayFindIndex(ctx *gramAntlr.ArrayFindIndexContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Arreglos"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error variable no declarada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, variable no es un arreglo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	valReturn := -1
	for i, item := range tempList {
		if reflect.DeepEqual(item, value) {
			valReturn = i
			break
		}
	}

	return valReturn
}

// VisitArrayJoin - 'ID_VARIABLE' '.' 'join' '(' expr ')'
func (v *CompilerVisitor) VisitArrayJoin(ctx *gramAntlr.ArrayJoinContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Arreglos"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error varible no declarda "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, variable no es un arreglo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	if strValue, ok := value.(string); ok {
		result := strings.Join(convertToStringSlice(tempList), strValue)
		return result
	} else {
		v.Salida += "Error-semántico: al unir el arreglo, el valor no es un string.\n"
		v.AgregarError(
			"Error al unir el arreglo, el valor no es un string",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
}

// VisitArrayLength - 'ID_VARIABLE' '.' 'length' '(' posicion* ')'
func (v *CompilerVisitor) VisitArrayLength(ctx *gramAntlr.ArrayLengthContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)

	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error variable no declarada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, variable no es un arreglo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	var listaBase []interface{} = tempList
	lenResult := 0

	if len(ctx.AllPosicion()) == 0 {
		lenResult = len(listaBase)
	} else {
		for _, pos := range ctx.AllPosicion() {
			index := v.Visit(pos.Expr())
			if idx, ok := index.(int); ok {
				if idx >= len(listaBase) {
					v.Salida += "Error-semántico: índice fuera de rango.\n"
					v.AgregarError(
						"Error índice fuera de rango",
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"semántico",
					)
					return nil
				}
				nextItem := listaBase[idx]
				if nextList, ok := nextItem.([]interface{}); ok {
					listaBase = nextList
				} else {
					v.Salida += "Error-semántico: al acceder al arreglo, la posición no es un arreglo.\n"
					v.AgregarError(
						"Error al acceder al arreglo, la posición no es un arreglo",
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"semántico",
					)
					return nil
				}
			} else {
				v.Salida += "Error-semántico: índice debe ser un entero.\n"
				v.AgregarError(
					"Error índice debe ser un entero",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
		}
		lenResult = len(listaBase)
	}

	return lenResult
}

// VisitArrayAppend - 'ID_VARIABLE' '.' 'append' '(' expr ')'
func (v *CompilerVisitor) VisitArrayAppend(ctx *gramAntlr.ArrayAppendContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())
	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error variable no declarada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, variable no es un arreglo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	if listValue, ok := value.([]interface{}); ok {
		newList := append(tempList, listValue...)
		return newList
	}

	if isValidType(value, variable.Type) {
		newList := append(tempList, value)
		return newList
	}

	v.Salida += "Error-semántico: al agregar valor al arreglo, los tipos no son compatibles.\n"
	v.AgregarError(
		"Error al agregar valor al arreglo, los tipos no son compatibles",
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"semántico",
	)
	return nil
}

// ------------------------------ DECLARACION SLICE ----------------------------------

// VisitVarDeclSliceStmt - 'var' ID_VARIABLE slice '=' sliceValores ';'
func (v *CompilerVisitor) VisitVarDeclSliceStmt(ctx *gramAntlr.VarDeclSliceStmtContext) interface{} {

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Arreglos"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "slice")

	return v.Visit(ctx.VarDclSlice())
}

// VisitSliceValores - 'ID_VARIABLE' ':' type '=' sliceValores
func (v *CompilerVisitor) VisitSliceValores(ctx *gramAntlr.SliceValoresContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	typo, err := parseSymbolType(typeStr)

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := typeStr         // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		v.AgregarError(
			"Error tipo no valido "+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	numDimensiones := 0
	for range ctx.AllNuevoSlice() {
		numDimensiones++
	}

	sliceFinal := v.Visit(ctx.ContenidoSlice())

	v.currentEnv.SetVariable(id, sliceFinal, typo, false, true, ctx.GetStart())

	return nil
}

// VisitSliceDcl_Asign
func (v *CompilerVisitor) VisitSliceDcl_Asign(ctx *gramAntlr.SliceDcl_AsignContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	typeVar, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		return nil
	}

	numDimensiones := 0
	for range ctx.AllNuevoSlice() {
		numDimensiones++
	}

	valueSlice := v.Visit(ctx.Expr())

	slice, ok := valueSlice.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: se intenta asignar %v a un slice.\n", valueSlice)
	}

	dimensionSliceExpr := v.funcionesNativasVisitor.GetDimensionSlice(slice, "[]", false)

	typeExpr, dimensionExpr, err := getDimension_type(dimensionSliceExpr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error-semántico: al obtener la dimensión del slice %s, %v\n", id, err)
		return nil
	}

	typeVarString := SymbolTypeToString(typeVar)

	if numDimensiones != dimensionExpr || typeVarString != typeExpr {
		v.Salida += fmt.Sprintf("Error-semántico: al asignar el slice a la variable %s, las dimensiones o tipos no coinciden.\n", id)
		return nil
	}

	v.currentEnv.SetVariable(id, slice, typeVar, false, true, ctx.GetStart())
	return nil
}

// VisitSliceVacio - 'ID_VARIABLE' ':' type '=' sliceVacio
func (v *CompilerVisitor) VisitSliceVacio(ctx *gramAntlr.SliceVacioContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	symbolType, err := parseSymbolType(typeStr)

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := typeStr         // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, id, linea, "slice")

	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		v.AgregarError(
			"Error tipo de dato no valido "+typeStr,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	v.currentEnv.SetVariable(id, make([]interface{}, 0), symbolType, false, true, ctx.GetStart())

	return nil
}

// VisitSliceContenido - contenido de un slice con expresiones
func (v *CompilerVisitor) VisitSliceContenido(ctx *gramAntlr.SliceContenidoContext) interface{} {
	arrayTemp := []interface{}{}

	nombre := ctx.GetText()    // Obtén el nombre correctamente
	tipo := "Contendio Slices" // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "slice")

	for _, expr := range ctx.AllExpr() {
		arrayTemp = append(arrayTemp, v.Visit(expr))
	}
	return arrayTemp
}

// VisitSliceContenidoSlice - contenido de un slice anidado
func (v *CompilerVisitor) VisitSliceContenidoSlice(ctx *gramAntlr.SliceContenidoSliceContext) interface{} {
	arrayTemp := []interface{}{}

	nombre := ctx.GetText()    // Obtén el nombre correctamente
	tipo := "Contendio Slices" // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "slice")

	for _, slice := range ctx.AllContenidoSlice() {
		sliceValue := v.Visit(slice)
		if sliceList, ok := sliceValue.([]interface{}); ok {
			arrayTemp = append(arrayTemp, sliceList)
		}
	}
	return arrayTemp
}

// ------------------------------ ASIGNACIONES --------------------------------------
// VisitAsignStmt
func (v *CompilerVisitor) VisitAsignStmt(ctx *gramAntlr.AsignStmtContext) interface{} {
	// ctx.VarExpr() accede al nodo 'VarExpr' dentro de la asignación.
	return v.Visit(ctx.VarAsign())
}

// VisitVarExpr
func (v *CompilerVisitor) VisitVarExpr(ctx *gramAntlr.VarExprContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())

	sym, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error-semántico: variable %s no encontrada.\n", id)
		v.AgregarError(
			"Error variable no encontrada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	mutabilidad := sym.Mutable
	tipo := sym.Type
	contextStart := ctx.GetStart()

	if value == nil {
		v.Salida += fmt.Sprintf("Error-semántico: al asignar el valor a la variable '%s', el valor es nulo.\n", id)
		v.AgregarError(
			"Error al acceder al arreglo, valor es nulo "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	// Validar si es un slice
	if slice, ok := value.([]interface{}); ok {
		if len(slice) == 0 || isValidType(slice[0], tipo) {
			v.currentEnv.SetVariable(id, value, tipo, mutabilidad, false, contextStart)
			return nil
		}
		v.Salida += fmt.Sprintf("Error-semántico: tipos no compatibles para el slice asignado a %s.\n", id)
		v.AgregarError(
			"Error tipo no compatible slice "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	// Validación normal
	if !isValidType(value, tipo) {
		if mutabilidad {
			v.currentEnv.SetVariable(id, value, tipo, mutabilidad, false, contextStart)
			return nil
		}
		v.Salida += fmt.Sprintf("Error-semántico: la variable %s no es mutable y no se puede asignar un nuevo valor.\n", id)
		v.AgregarError(
			"Error variable no mutable "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	v.currentEnv.SetVariable(id, value, tipo, mutabilidad, false, contextStart)
	return nil
}

// VisitvarAdd
func (v *CompilerVisitor) VisitVarAdd(ctx *gramAntlr.VarAddContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())
	token := ctx.GetStart()

	sym, _ := v.currentEnv.GetVariable(id)
	typ := sym.Type

	switch typ {
	case 1: // FLOAT64
		switch val := value.(type) {
		case int:
			switch ctx.GetOp().GetText() {
			case "+=":
				v.currentEnv.SetVariable(id, sym.Value.(float64)+float64(val), typ, sym.Mutable, false, token)
			case "-=":
				v.currentEnv.SetVariable(id, sym.Value.(float64)-float64(val), typ, sym.Mutable, false, token)
			}
			return nil
		case float64:
			switch ctx.GetOp().GetText() {
			case "+=":
				v.currentEnv.SetVariable(id, sym.Value.(float64)+val, typ, sym.Mutable, false, token)
			case "-=":
				v.currentEnv.SetVariable(id, sym.Value.(float64)-val, typ, sym.Mutable, false, token)
			}
			return nil
		}
	case 2: // STRING
		if ctx.GetOp().GetText() == "-=" {
			v.Salida += "Error-semántico: al operar -= no se pueden operar strings."
			v.AgregarError(
				"Error al operar -= no se pueden operar strings.",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
		}
		if val, ok := value.(string); ok {
			v.currentEnv.SetVariable(id, sym.Value.(string)+val, typ, sym.Mutable, false, token)
			return nil
		}
	case 0: // INT
		if val, ok := value.(int); ok {
			switch ctx.GetOp().GetText() {
			case "+=":
				v.currentEnv.SetVariable(id, sym.Value.(int)+val, typ, sym.Mutable, false, token)
			case "-=":
				v.currentEnv.SetVariable(id, sym.Value.(int)-val, typ, sym.Mutable, false, token)
			}
			return nil
		}
	}

	v.Salida = "Error-semántico: al asignar el valor a la variable, los tipos no son compatibles."
	v.AgregarError(
		"Error asignacion de valor a la variable, no compatible "+id,
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"semántico",
	)
	return nil
}

// VarInc
func (v *CompilerVisitor) VisitVarInc(ctx *gramAntlr.VarIncContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	token := ctx.GetStart()

	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "variable"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")

	sym, _ := v.currentEnv.GetVariable(id)
	typ := sym.Type

	// Solo se permite ++ y -- para INT (0) y FLOAT64 (1)
	if typ != 0 && typ != 1 {
		v.Salida += "Error-semántico: el tipo de variable no acepta operador ++ o --."
		v.AgregarError(
			"Error el tipo de variable no acepta operador ++ o --.",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	switch ctx.GetOp().GetText() {
	case "++":
		if typ == 0 { // INT
			v.currentEnv.SetVariable(id, sym.Value.(int)+1, typ, sym.Mutable, false, token)
		} else { // FLOAT64
			v.currentEnv.SetVariable(id, sym.Value.(float64)+1, typ, sym.Mutable, false, token)
		}
	case "--":
		if typ == 0 { // INT
			v.currentEnv.SetVariable(id, sym.Value.(int)-1, typ, sym.Mutable, false, token)
		} else { // FLOAT64
			v.currentEnv.SetVariable(id, sym.Value.(float64)-1, typ, sym.Mutable, false, token)
		}
	}

	return nil
}

// ---------------------------- CONDICIONAL IF ---------------------------------------
// VisitIfStmt actúa como un puente para las reglas anidadas de 'sIf'.
func (v *CompilerVisitor) VisitIfStmt(ctx *gramAntlr.IfStmtContext) interface{} {
	// ctx.SIf() accede al nodo 'sIf' dentro de la instrucción 'if'.
	// Al visitarlo, ANTLR llamará al método correcto: VisitIfOnly o VisitIfAnidado.
	return v.Visit(ctx.SIf())
}

// -------------------- Produccion IF ELSE --------------------
func (v *CompilerVisitor) VisitIfOnly(ctx *gramAntlr.IfOnlyContext) interface{} {

	// Evaluar la condición
	value := v.Visit(ctx.Expr())
	if value == nil {
		v.Salida += "Error semántico: condición del if es nil\n"
		v.AgregarError(
			"Error condición del if es nil",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	cond, ok := value.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del if, no es un booleano.\n"
		v.AgregarError(
			"Error al evaluar la condición del if, no es un booleano.",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	// Si la condición es verdadera, ejecutar el bloque[0]
	if cond {
		// Nuevo entorno anidado
		newEnv := NewEnvironment(v.currentEnv)
		v.currentEnv = newEnv

		// Visitar el bloque del 'if'
		result := v.Visit(ctx.Block(0))

		//fmt.Println("SCOPE DEL IF antes de restaurar padre \n", v.currentEnv.ImprimirScope())
		// Restaurar el entorno
		v.currentEnv = v.currentEnv.Parent

		// Propagar break / continue / return void
		if str, ok := result.(string); ok {
			if str == "break" || str == "continue" || str == "Excepcion___Return_Void" {
				return str
			}
		}
		// Si fue una expresión con valor, devolverla
		if result != nil {
			return result
		}
	} else {
		// Si hay un bloque 'else' (ctx.Block() retorna []IBlockContext)
		blocks := ctx.AllBlock()
		if len(blocks) > 1 {
			// Visitar el bloque del 'else'
			// Nuevo entorno anidado
			newEnv := NewEnvironment(v.currentEnv)
			v.currentEnv = newEnv

			result := v.Visit(blocks[1])
			// Restaurar el entorno
			v.currentEnv = v.currentEnv.Parent
			if str, ok := result.(string); ok {
				if str == "break" || str == "continue" || str == "Excepcion___Return_Void" {
					return str
				}
			}
			if result != nil {
				return result
			}
		}
	}

	return nil
}

// -------------------- Produccion IF ELSE IF --------------------
func (v *CompilerVisitor) VisitIfAnidado(ctx *gramAntlr.IfAnidadoContext) interface{} {
	// Evaluar la condición
	value := v.Visit(ctx.Expr())
	cond, ok := value.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del if, no es un booleano.\n"
		v.AgregarError(
			"Error al evaluar la condición del if, no es un booleano.",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	if cond {
		// Nuevo entorno anidado
		newEnv := NewEnvironment(v.currentEnv)
		v.currentEnv = newEnv

		// Visitar el bloque del 'if'
		result := v.Visit(ctx.Block())

		// Restaurar el entorno
		v.currentEnv = v.currentEnv.Parent

		// Propagar break / continue / return void
		if str, ok := result.(string); ok {
			if str == "break" || str == "continue" || str == "Excepcion___Return_Void" {
				return str
			}
		}
		if result != nil {
			return result
		}
	} else {
		// Si la condición es falsa, ejecutar el 'else if' en ctx.SIf()
		result := v.Visit(ctx.SIf())
		if str, ok := result.(string); ok {
			if str == "break" || str == "continue" || str == "Excepcion___Return_Void" {
				return str
			}
		}
		if result != nil {
			return result
		}
	}
	return nil
}

// -------------------- VisitBreakStmt --------------------
func (v *CompilerVisitor) VisitBreakStmt(ctx *gramAntlr.BreakStmtContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "break"         // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")
	return "break"
}

// -------------------- VisitContinue --------------------
func (v *CompilerVisitor) VisitContinueStmt(ctx *gramAntlr.ContinueStmtContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "continue"      // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")
	return "continue"
}

// -------------------- VisitReturnStmt --------------------
func (v *CompilerVisitor) VisitReturnStmt(ctx *gramAntlr.ReturnStmtContext) interface{} {
	var valueRet = ctx.Retorno()
	if valueRet.Expr() != nil {
		nombre := ctx.GetText() // Obtén el nombre correctamente
		tipo := "return"        // Obtén el ámbito real
		linea := ctx.GetStart().GetLine()
		v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")
		return v.Visit(valueRet.Expr())
	} else {
		nombre := ctx.GetText() // Obtén el nombre correctamente
		tipo := "return"        // Obtén el ámbito real
		linea := ctx.GetStart().GetLine()
		v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")
		return "Excepcion___Return_Void" // Retorno vacío
	}
}

// -------------------- VisitBlockStmt --------------------
func (v *CompilerVisitor) VisitBlockStmt(ctx *gramAntlr.BlockStmtContext) interface{} {
	for _, instrCtx := range ctx.AllInstrucciones() {
		value := v.Visit(instrCtx)
		// Solo propagar control de flujo especial
		if str, ok := value.(string); ok {
			if str == "break" || str == "continue" {
				return str
			}
			if str == "Excepcion___Return_Void" {
				return str // Retorno vacío
			}
		}
		if value != nil {
			// Si hay un valor de retorno, devolverlo
			return value
		}
	}
	return nil
}

// ----------------------- VisitIdentifier -----------------------
func (v *CompilerVisitor) VisitIdentifier(ctx *gramAntlr.IdentifierContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	sym, err := v.currentEnv.GetVariable(id)

	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		v.AgregarError(
			"Error variable no declarada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	return sym.Value
}

// ---------------------------- SENTENCIAS FOR -----------------------------------
func (v *CompilerVisitor) VisitForStmt(ctx *gramAntlr.ForStmtContext) interface{} {
	return v.Visit(ctx.SFor())
}

// VisitForCondicion
func (v *CompilerVisitor) VisitForCondicion(ctx *gramAntlr.ForCondicionContext) interface{} {
	condition := v.Visit(ctx.Expr())
	if condBool, ok := condition.(bool); !ok {
		v.Salida += "Error-semántico: al evaluar la condición del for, no es un booleano."
		v.AgregarError(
			"Error al evaluar la condición del if, no es un booleano.",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
	} else {
		for condBool {
			newEnv := NewEnvironment(v.currentEnv)
			v.currentEnv = newEnv
			result := v.Visit(ctx.Block())

			if str, ok := result.(string); ok {
				if str == "break" {
					break
				} else if str == "continue" {
					condition = v.Visit(ctx.Expr())
					if condBool, ok = condition.(bool); !ok {
						v.Salida += "Error-semántico: al reevaluar la condición del for tras un 'continue', no es un booleano."
						v.AgregarError(
							"Error al reevaluar la condición del for tras un 'continue', no es un booleano.",
							ctx.GetStart().GetLine(),
							ctx.GetStart().GetColumn(),
							"semántico",
						)
					}
					continue
				} else if str == "Excepcion___Return_Void" {
					return str
				} else {
					return str
				}
			}
			if result != nil {
				return result
			}

			// Recalcular la condición
			condition = v.Visit(ctx.Expr())
			if condBool, ok = condition.(bool); !ok {
				v.Salida += "Error-semántico: al reevaluar la condición del for, no es un booleano."
				v.AgregarError(
					"Error al reevaluar la condición del for, no es un booleano.",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
			}
			v.currentEnv = newEnv.Parent // Restaurar el entorno
		}
	}
	return nil
}

// VisitForAsignacion
// VisitForAsignacion
func (v *CompilerVisitor) VisitForAsignacion(ctx *gramAntlr.ForAsignacionContext) interface{} {
	// Crear nuevo entorno para el for
	newEnv := NewEnvironment(v.currentEnv)
	v.currentEnv = newEnv

	// Declaración inicial del for
	v.Visit(ctx.VarDcl())

	// Evaluación inicial de la condición
	condition := v.Visit(ctx.Expr())
	condBool, ok := condition.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del for, no es un booleano.\n"
		// Restaurar entorno padre antes de salir
		v.currentEnv = newEnv.Parent
		return nil
	}

	for condBool {
		// Ejecutar bloque
		result := v.Visit(ctx.Block())

		// Control de flujo: break, continue, return
		if str, ok := result.(string); ok {
			if str == "break" {
				break
			} else if str == "continue" {
				// Ejecutar asignación y reevaluar la condición
				v.Visit(ctx.VarAsign())
				condition = v.Visit(ctx.Expr())
				if condBool, ok = condition.(bool); !ok {
					v.Salida += "Error-semántico: al reevaluar la condición del for tras un 'continue', no es un booleano."
				}
				continue
			} else if str == "Excepcion___Return_Void" {
				v.currentEnv = newEnv.Parent
				return str
			} else {
				v.currentEnv = newEnv.Parent
				return str
			}
		}

		if result != nil {
			v.currentEnv = newEnv.Parent
			return result
		}

		// Ejecutar asignación final
		v.Visit(ctx.VarAsign())

		// Reevaluar condición
		condition = v.Visit(ctx.Expr())
		if condBool, ok = condition.(bool); !ok {
			v.Salida += "Error-semántico: al reevaluar la condición del for, no es un booleano."
		}
	}

	// Restaurar entorno
	v.currentEnv = newEnv.Parent
	return nil
}

// VisitForRange
func (v *CompilerVisitor) VisitForRange(ctx *gramAntlr.ForRangeContext) interface{} {
	// Crear nuevo entorno para el for
	newEnv := NewEnvironment(v.currentEnv)
	v.currentEnv = newEnv

	// Obtenemos el arreglo a iterar
	idSlice := ctx.ID_VARIABLE(2).GetText()

	// Verificamos que el idSlice sea un []interface{}
	sliceVar, err := v.currentEnv.GetVariable(idSlice)
	// Si no existe, mostramos un error
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", idSlice)
		v.currentEnv = newEnv.Parent // Restaurar entorno padre
		return nil
	}
	// si no es un slice, mostramos un error
	sliceValue, ok := sliceVar.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al iterar, la variable %s no es un slice.\n", idSlice)
		v.currentEnv = newEnv.Parent // Restaurar entorno padre
		return nil
	}

	// Inicializamos variable que recibe el valor del slice
	switch sliceVar.Type {
	case 0:
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), 0, INT, false, true, ctx.GetStart())
	case 1:
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), 0.0, FLOAT64, false, true, ctx.GetStart())
	case 2:
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), "", STRING, false, true, ctx.GetStart())
	case 3:
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), false, BOOL, false, true, ctx.GetStart())
	case 4:
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), rune(0), RUNE, false, true, ctx.GetStart())
	default:
		v.Salida += fmt.Sprintf("Error-semántico: tipo de variable %s no soportado para iteración.\n", ctx.ID_VARIABLE(1).GetText())
		v.currentEnv = newEnv.Parent // Restaurar entorno padre
		return nil
	}

	// Inicializamos la variable de índice
	v.currentEnv.SetVariable(ctx.ID_VARIABLE(0).GetText(), 0, INT, false, true, ctx.GetStart())

	// Iteramos con range en el slice
	for i, item := range sliceValue {
		// Actualizamos tanto el indice como el valor del slice
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(0).GetText(), i, INT, false, false, ctx.GetStart())
		v.currentEnv.SetVariable(ctx.ID_VARIABLE(1).GetText(), item, sliceVar.Type, false, false, ctx.GetStart())

		// Ejecutamos el bloque del for
		result := v.Visit(ctx.Block())

		// Control de flujo: break, continue, return
		if str, ok := result.(string); ok {
			if str == "break" {
				break // Salir del for
			} else if str == "continue" {
				continue // Continuar con la siguiente iteración
			} else if str == "Excepcion___Return_Void" {
				v.currentEnv = newEnv.Parent // Restaurar entorno padre
				return str                   // Retorno vacío
			} else {
				v.currentEnv = newEnv.Parent // Restaurar entorno padre
				return str                   // Retornar cualquier otro valor
			}
		}

		// Si hay un valor de retorno, devolverlo
		if result != nil {
			v.currentEnv = newEnv.Parent // Restaurar entorno padre
			return result
		}
	}
	// Recuperar el entorno padre
	v.currentEnv = newEnv.Parent

	return nil
}

// ---------------------------------------------------- switch ----------------------------------------------------

// Produccion de instrucciones de switch
func (v *CompilerVisitor) VisitSwitchInstruccion(ctx *gramAntlr.SwitchInstruccionContext) interface{} {
	return v.Visit(ctx.SSwitch())
}

// Produccion de switch
func (v *CompilerVisitor) VisitSwitchStmt(ctx *gramAntlr.SwitchStmtContext) interface{} {
	return v.Visit(ctx.Cases())
}

// Produccion cases
func (v *CompilerVisitor) VisitCase(ctx *gramAntlr.CaseContext) interface{} {

	caseCondition := v.Visit(ctx.Expr())

	// Implement isEqualType functionality
	if (caseCondition == nil && v.conditionExpr == nil) ||
		(caseCondition != nil && v.conditionExpr != nil &&
			fmt.Sprintf("%T", caseCondition) == fmt.Sprintf("%T", v.conditionExpr)) {

		if caseCondition == v.conditionExpr { // Direct comparison for equality
			newEnv := NewEnvironment(v.currentEnv)
			v.currentEnv = newEnv

			for _, instruccion := range ctx.AllInstrucciones() {
				dato := v.Visit(instruccion)
				if str, ok := dato.(string); ok && str == "break" {
					break
				}
			}
			v.currentEnv = newEnv.Parent
			return nil
		}
	} else {
		v.Salida += fmt.Sprintf("Error semántico: al evaluar la condición del case, los tipos no son compatibles. Linea: %d, Columna: %d\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		v.AgregarError(
			"Error al evaluar la condicion case",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil // Or propagate an error
	}

	if ctx.Cases() != nil { // Check if there are more cases or a default
		v.Visit(ctx.Cases())
	}

	return nil
}

// Produccion default
func (v *CompilerVisitor) VisitDefault(ctx *gramAntlr.DefaultContext) interface{} {
	newEnv := NewEnvironment(v.currentEnv)
	v.currentEnv = newEnv
	nombre := ctx.GetText()       // Obtén el nombre correctamente
	tipo := "switch case default" // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "funcion")
	for _, instruccion := range ctx.AllInstrucciones() {
		v.Visit(instruccion)
	}

	v.currentEnv = newEnv.Parent

	return nil
}

// ---------------------------- FUNCIONES NATIVAS -----------------------------
// VisitIntToString
func (v *CompilerVisitor) VisitIntToString(ctx *gramAntlr.IntToStringContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Nativa"        // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, "global", linea, "funcion")
	return v.funcionesNativasVisitor.VisitIntToString(ctx)
}

// VisitfloatToString
func (v *CompilerVisitor) VisitFloatToString(ctx *gramAntlr.FloatToStringContext) interface{} {
	nombre := ctx.GetText() // Obtén el nombre correctamente
	tipo := "Nativa"        // Obtén el ámbito real
	linea := ctx.GetStart().GetLine()
	v.TablaSimbolos.Insertar(nombre, tipo, "global", linea, "funcion")
	return v.funcionesNativasVisitor.VisitFloatToString(ctx)
}

// VisitReflectType
func (v *CompilerVisitor) VisitReflectType(ctx *gramAntlr.ReflectTypeContext) interface{} {
	value := v.Visit(ctx.Expr())
	if _, ok := value.(map[string]Symbol); ok {

		return "struct"
	}

	return v.funcionesNativasVisitor.VisitReflectType(ctx)
}

// -------------------------------- STRUCTS -------------------------------------
func (v *CompilerVisitor) VisitVarDeclStructStmt(ctx *gramAntlr.VarDeclStructStmtContext) interface{} {

	return v.Visit(ctx.VarDclStruct())
}

// VisitDeclStructData - Definición de estructura del struct
func (v *CompilerVisitor) VisitDeclStructData(ctx *gramAntlr.DeclStructDataContext) interface{} {
	var firstTime bool = true
	var tipoValor int = 0
	var id string = ""
	variableStruct := make(map[string]Symbol)

	for _, variable := range ctx.AllID_VARIABLE() {
		if firstTime {
			firstTime = false
			id = variable.GetText() // Nombre del struct
			continue
		} else {
			varStruct, ok := v.currentEnv.GetVariable(ctx.AllType_()[tipoValor].GetText())
			if ok == nil {
				// Agregando variable al struct
				variableStruct[variable.GetText()] = *varStruct
				continue
			}

			if varStruct == nil && ctx.AllType_()[tipoValor].GetText() == id {
				variableStruct[variable.GetText()] = Symbol{nil, STRUCT, true}
				continue
			}
			//agregar aqui
			symbolType, _ := StringToSymbolType(ctx.AllType_()[tipoValor].GetText())
			switch symbolType {
			case INT:
				nombre := ctx.GetText() // Obtén el nombre correctamente
				tipo := "int"           // Obtén el ámbito real
				linea := ctx.GetStart().GetLine()
				v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "struct")
				variableStruct[variable.GetText()] = Symbol{0, INT, false}
			case FLOAT64:
				nombre := ctx.GetText() // Obtén el nombre correctamente
				tipo := "float"         // Obtén el ámbito real
				linea := ctx.GetStart().GetLine()
				v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "struct")
				variableStruct[variable.GetText()] = Symbol{0.0, FLOAT64, false}
			case STRING:
				nombre := ctx.GetText() // Obtén el nombre correctamente
				tipo := "string"        // Obtén el ámbito real
				linea := ctx.GetStart().GetLine()
				v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "struct")
				variableStruct[variable.GetText()] = Symbol{"", STRING, false}
			case BOOL:
				nombre := ctx.GetText() // Obtén el nombre correctamente
				tipo := "bool"          // Obtén el ámbito real
				linea := ctx.GetStart().GetLine()
				v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "struct")
				variableStruct[variable.GetText()] = Symbol{false, BOOL, false}
			case RUNE:
				nombre := ctx.GetText() // Obtén el nombre correctamente
				tipo := "nil"           // Obtén el ámbito real
				linea := ctx.GetStart().GetLine()
				v.TablaSimbolos.Insertar(nombre, tipo, nombre, linea, "struct")
				variableStruct[variable.GetText()] = Symbol{rune(0), RUNE, false}
			default:
				v.Salida += fmt.Sprintf("Error semántico: tipo no compatible para variable: %s", ctx.AllType_()[tipoValor].GetText())
				v.AgregarError(
					"Error tipo no compatible para variable "+id,
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
			}
			tipoValor++
		}

	}
	v.currentEnv.SetVariable(id, variableStruct, STRUCT, false, true, ctx.GetStart())
	return nil
}

// VisitVarStructDclStmt - Declaración de variable de tipo struct
func (v *CompilerVisitor) VisitVarStructDclStmt(ctx *gramAntlr.VarStructDclStmtContext) interface{} {
	return v.Visit(ctx.VarStructDcl())
}

func (v *CompilerVisitor) VisitStructVarTypeInference(ctx *gramAntlr.StructVarTypeInferenceContext) interface{} {
	idStruct := ctx.ID_VARIABLE(1).GetText()
	baseStruct, _ := v.currentEnv.GetVariable(idStruct)

	if baseStruct.Type != 6 {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no es un struct\n", idStruct)
		v.AgregarError(
			"Error variable no es Struct "+idStruct,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	bodyStruct, ok := baseStruct.Value.(map[string]Symbol)
	if !ok {
		v.Salida += fmt.Sprintf("Error semántico: el contenido de %s no es un struct válido (map[string]Symbol)\n", idStruct)
		v.AgregarError(
			"Error no es un struct valido "+idStruct,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	if len(bodyStruct) != len(ctx.AllID_VARIABLE())-2 {
		v.Salida += fmt.Sprintf("Error semántico: número de campos en la declaración de struct %s no coincide con el número de campos definidos\n", idStruct)
		v.AgregarError(
			"Error numero de campos en la declaracion Struct ",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
	}
	//Agregar copia de datoStructBase a una nueva instancia de struct
	// ✅ Deep Copy del struct base
	copiaDeep := make(map[string]Symbol)

	for k, vSimbolo := range bodyStruct {
		var copiedValue interface{}

		// Verificar si el value es otro struct (map[string]Symbol)
		if innerMap, ok := vSimbolo.Value.(map[string]Symbol); ok {
			innerCopy := make(map[string]Symbol)
			for innerKey, innerSymbol := range innerMap {
				innerCopy[innerKey] = Symbol{
					Value:   innerSymbol.Value, // puedes profundizar más aquí si se necesita
					Type:    innerSymbol.Type,
					Mutable: innerSymbol.Mutable,
				}
			}
			copiedValue = innerCopy
		} else {
			// Si no es un map, asumimos que es un tipo básico o valor por copia
			copiedValue = vSimbolo.Value
		}

		// Crear la copia del símbolo
		copiaDeep[k] = Symbol{
			Value:   copiedValue,
			Type:    vSimbolo.Type,
			Mutable: vSimbolo.Mutable,
		}
	}

	for i := 2; i < len(ctx.AllID_VARIABLE()); i++ {
		varBaseStruct, exist := copiaDeep[ctx.ID_VARIABLE(i).GetText()]
		if !exist {
			v.Salida += fmt.Sprintf("Error semántico: campo %s no existe en struct base %s\n",
				ctx.ID_VARIABLE(i).GetText(), idStruct)
			v.AgregarError(
				"Error campo no existe en Struct",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
		expVisit := v.Visit(ctx.Expr(i - 2))
		if expVisit == "nil" {
			temp := copiaDeep[ctx.ID_VARIABLE(i).GetText()]
			temp.Value = "nil"
			copiaDeep[ctx.ID_VARIABLE(i).GetText()] = temp
			continue
		}
		if isValidType(expVisit, varBaseStruct.Type) {
			temp := copiaDeep[ctx.ID_VARIABLE(i).GetText()]
			temp.Value = expVisit
			copiaDeep[ctx.ID_VARIABLE(i).GetText()] = temp
			continue
		} else {
			v.Salida += fmt.Sprintf("Error semántico: tipo incompatible para campo %s\n", ctx.ID_VARIABLE(i).GetText())
			v.AgregarError(
				"Error tipo incompatible para campo",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	}

	v.currentEnv.SetVariable(ctx.ID_VARIABLE(0).GetText(), copiaDeep, STRUCT, false, true, ctx.GetStart())

	if _, exists := v.StructRelational[idStruct]; exists {
		v.StructRelational[idStruct] = append(v.StructRelational[idStruct], ctx.ID_VARIABLE(0).GetText())
	} else {
		v.StructRelational[idStruct] = []string{ctx.ID_VARIABLE(0).GetText()}
	}
	return nil
}

// VisitStructAccess - Acceso a campo de struct
func (v *CompilerVisitor) VisitStructAccess(ctx *gramAntlr.StructAccessContext) interface{} {
	var varStruct Symbol
	var datosStruct map[string]Symbol

	first := true

	for i := 0; i < len(ctx.AllID_VARIABLE())-1; i++ {
		idStruct := ctx.ID_VARIABLE(i).GetText()
		idVar := ctx.ID_VARIABLE(i + 1).GetText()

		if first {
			first = false
			// Buscar la variable struct en el entorno
			symbol, found := v.currentEnv.GetVariable(idStruct)
			if found != nil || symbol.Type != STRUCT {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no es un struct o no existe.\n", idStruct)
				v.AgregarError(
					"Error variable no es Struct o no existe",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			varStruct = *symbol

			// Convertir el valor del struct a map[string]Symbol
			castedMap, ok := varStruct.Value.(map[string]Symbol)
			if !ok {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no contiene un struct válido.\n", idStruct)
				v.AgregarError(
					"Error la variable no contiene struct",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			datosStruct = castedMap

			if val, exists := datosStruct[idVar]; exists {
				if i+1 == len(ctx.AllID_VARIABLE())-1 {
					return val.Value
				} else {
					varStruct = val
					if nestedMap, ok := varStruct.Value.(map[string]Symbol); ok {
						datosStruct = nestedMap
					} else {
						v.Salida += fmt.Sprintf("Error semántico: la variable %s no es un struct (anidado).\n", idVar)
						v.AgregarError(
							"Error variable no es un struct",
							ctx.GetStart().GetLine(),
							ctx.GetStart().GetColumn(),
							"semántico",
						)
						return nil
					}
				}
			} else {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no existe en el struct %s.\n", idVar, idStruct)
				v.AgregarError(
					"Error el Struct no existe",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
		} else {
			idVar := ctx.ID_VARIABLE(i + 1).GetText()

			if val, exists := datosStruct[idVar]; exists {
				if i+1 == len(ctx.AllID_VARIABLE())-1 {
					return val.Value
				}
				varStruct = val
				if nestedMap, ok := varStruct.Value.(map[string]Symbol); ok {
					datosStruct = nestedMap
				} else {
					v.Salida += fmt.Sprintf("Error semántico: la variable %s no es un struct (anidado).\n", idVar)
					v.AgregarError(
						"Error la variable no es un Struct",
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"semántico",
					)
					return nil
				}
			} else {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no existe en el struct %s.\n", idVar, ctx.ID_VARIABLE(i).GetText())
				v.AgregarError(
					"Error la variable no existe en el Struct",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
		}
	}

	return nil
}

// VisitStructAccessAsign - Asignación a campo de struct
func (v *CompilerVisitor) VisitStructAccessAsign(ctx *gramAntlr.StructAccessAsignContext) interface{} {
	datosStruct := make(map[string]Symbol)
	first := true

	for i := 0; i < len(ctx.AllID_VARIABLE())-1; i++ {
		idStruct := ctx.ID_VARIABLE(i).GetText()
		idVar := ctx.ID_VARIABLE(i + 1).GetText()

		if first {
			first = false
			varStruct, err := v.currentEnv.GetVariable(idStruct)
			if err != nil {
				return nil
			}
			if varStruct.Type != 6 {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no es un struct.\n", idStruct)
				v.AgregarError(
					"Error la variable no es un struct",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			datosStruct = make(map[string]Symbol)
			castedMap, ok := varStruct.Value.(map[string]Symbol)
			if !ok {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no contiene un struct válido.\n", idStruct)
				v.AgregarError(
					"Error la variable no contiene struct",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			datosStruct = castedMap
			sym, exists := datosStruct[idVar]
			if exists {
				if i+1 == len(ctx.AllID_VARIABLE())-1 {
					temp := datosStruct[idVar]
					temp.Value = v.Visit(ctx.Expr())
					datosStruct[idVar] = temp
					return nil
				} else {
					varStruct = &sym
					if nestedMap, ok := varStruct.Value.(map[string]Symbol); ok {
						datosStruct = nestedMap
					} else {
						v.Salida += fmt.Sprintf("Error semántico: la variable %s no es un struct (anidado).\n", idVar)
						v.AgregarError(
							"Error la variable no es un struct",
							ctx.GetStart().GetLine(),
							ctx.GetStart().GetColumn(),
							"semántico",
						)
						return nil
					}
				}
			} else {
				v.Salida += fmt.Sprintf("Error semántico: la variable %s no existe en el struct %s.\n", idVar, idStruct)
				v.AgregarError(
					"Error variable no existe en el Struct",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}

		}

	}

	return nil
}

// -------------------------------- FUNCIONES ------------------------------------
// VisitFuncionStmt
func (v *CompilerVisitor) VisitFunctionStmt(ctx *gramAntlr.FunctionStmtContext) interface{} {

	if ctx.Functions().GetChild(1).(antlr.ParseTree).GetText() == "main" {
		v.Visit(ctx.Functions())
		varFunc, _ := v.currentEnv.GetFuncion("main")
		body := varFunc.Body
		return v.Visit(body)
	}
	return v.Visit(ctx.Functions())
}

// VisitFunciones
func (v *CompilerVisitor) VisitFunciones(ctx *gramAntlr.FuncionesContext) interface{} {

	id := ctx.ID_VARIABLE(0).GetText() // Nombre de la función

	// Mapa de parámetros: id -> Symbol
	parametros := []*TupleStringSymbol{}

	for i := 1; i < len(ctx.AllID_VARIABLE()); i++ {
		idParam := ctx.ID_VARIABLE(i).GetText()
		typeText := ctx.Type_(i - 1).GetText()

		var tipo SymbolType
		switch strings.ToLower(typeText) {
		case "int":
			tipo = INT
			parametros = append(parametros, &TupleStringSymbol{idParam, &Symbol{0, tipo, true}})
		case "float64":
			tipo = FLOAT64
			parametros = append(parametros, &TupleStringSymbol{idParam, &Symbol{0.0, tipo, true}})
		case "string":
			tipo = STRING
			parametros = append(parametros, &TupleStringSymbol{idParam, &Symbol{"", tipo, true}})
		case "bool":
			tipo = BOOL
			parametros = append(parametros, &TupleStringSymbol{idParam, &Symbol{false, tipo, true}})
		case "rune":
			tipo = RUNE
			parametros = append(parametros, &TupleStringSymbol{idParam, &Symbol{'\x00', tipo, true}})
		default:
			v.Salida += fmt.Sprintf("Error semántico: tipo de parámetro no reconocido: %s\n", typeText)
			v.AgregarError(
				"Error tipo de parametro no reconocido",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	}

	// Tipo de retorno
	var tipoRet SymbolType = VOID
	if ctx.ValRet() != nil {
		typeText := ctx.ValRet().Type_().GetText()
		switch strings.ToLower(typeText) {
		case "int":
			tipoRet = INT
		case "float64":
			tipoRet = FLOAT64
		case "string":
			tipoRet = STRING
		case "bool":
			tipoRet = BOOL
		case "rune":
			tipoRet = RUNE
		default:
			v.Salida += fmt.Sprintf("Error semántico: tipo de retorno no reconocido: %s\n", typeText)
			v.AgregarError(
				"Error tipo de retorno no reconocido",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	}

	// Guardar la función en el entorno actual
	v.currentEnv.SetFunciones(id, parametros, ctx.Block(), tipoRet, ctx.GetStart())
	return nil
}

// VisitCallFunctionStmt
func (v *CompilerVisitor) VisitCallFunctionStmt(ctx *gramAntlr.CallFunctionStmtContext) interface{} {

	return v.Visit(ctx.VarCallStatement())
}

// VisitCallFunction
func (v *CompilerVisitor) VisitCallFunction(ctx *gramAntlr.CallFunctionContext) interface{} {

	id := ctx.ID_VARIABLE().GetText()
	parametros := []*TupleStringSymbol{}
	for i := 0; i < len(ctx.AllExpr()); i++ {
		value := v.Visit(ctx.Expr(i))

		if value == nil {
			v.Salida += fmt.Sprintf("Error-semántico: al llamar a la función %s, el parámetro %d es nulo.\n", id, i+1)
			v.AgregarError(
				"Error al llamar funcion",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
		switch value.(type) {
		case int:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, INT, false}})
		case float64:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, FLOAT64, false}})
		case string:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, STRING, false}})
		case bool:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, BOOL, false}})
		case rune:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, RUNE, false}})
		default:
			v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro no reconocido para la función %s.\n", id)
			v.AgregarError(
				"Error topo de parametro no reconocido "+id,
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	}
	funcion, error := v.currentEnv.GetFuncion(id)
	if error != nil {
		v.Salida += fmt.Sprintf("Error-semántico: función %s no encontrada.\n", id)
		v.AgregarError(
			"Error funcion no encontrada "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	parameters := funcion.Parameters
	var body gramAntlr.IBlockContext = funcion.Body
	tipoReturn := funcion.ValRet

	if len(parameters) != len(parametros) {
		v.Salida += fmt.Sprintf("Error-semántico: número de parámetros incorrecto al llamar a la función %s. Se esperaban %d, pero se recibieron %d.\n", id, len(parameters), len(parametros))
		v.AgregarError(
			"Error numero de parametros incorrecto "+id,
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	newEnv := NewEnvironment(v.currentEnv)
	v.currentEnv = newEnv

	if tipoReturn != 7 {
		// Verificar tipos de parámetros
		for i := 0; i < len(parametros); i++ {
			if parametros[i].Value.Type != parameters[i].Value.Type {
				v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro %v no coincide con el tipo esperado %v en la función %s.\n",
					parametros[i].Value.Type, parameters[i].Value.Type, id)
				v.AgregarError(
					"Error tipo de parametro no coincide "+id,
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			v.currentEnv.SetVariable(parameters[i].Key, parametros[i].Value.Value, parametros[i].Value.Type, false, true, ctx.GetStart())
		}
		valRet := v.Visit(body)
		fmt.Println("VALRET:", valRet, "TIPORETURN:", tipoReturn)
		fmt.Println("isValidType:", isValidType(valRet, tipoReturn))
		if isValidType(valRet, tipoReturn) {
			v.currentEnv = v.currentEnv.Parent // Restaurar el entorno anterior
			return valRet
		} else {
			v.Salida += fmt.Sprintf("Error-semántico: tipo de retorno %v no coincide con el tipo esperado %v en la función %s.\n",
				tipoReturn, fmt.Sprintf("%T", valRet), id)
			v.AgregarError(
				"Error tipo de retorno no coincide "+id,
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	} else {
		for i := 0; i < len(parametros); i++ {
			if parametros[i].Value.Type != parameters[i].Value.Type {
				v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro %v no coincide con el tipo esperado %v en la función %s.\n",
					parametros[i].Value.Type, parameters[i].Value.Type, id)
				v.AgregarError(
					"Error tipo de parametro no coincide "+id,
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			v.currentEnv.SetVariable(parameters[i].Key, parametros[i].Value.Value, parameters[i].Value.Type, parameters[i].Value.Mutable, true, ctx.GetStart())
		}
		v.Visit(body) // Ejecutar el cuerpo de la función
	}
	v.currentEnv = v.currentEnv.Parent // Restaurar el entorno anterior
	return nil
}

// VisitCallFunctionValue
func (v *CompilerVisitor) VisitCallFunctionValue(ctx *gramAntlr.CallFunctionValueContext) interface{} {

	return v.Visit(ctx.VarCallStatement())
}

// ----------------------------- FUNCION DE STRUCT --------------------------------
func (v *CompilerVisitor) VisitFunctionStructStmt(ctx *gramAntlr.FunctionStructStmtContext) interface{} {

	return v.Visit(ctx.FunctionStruct())
}

// Declaracion de función de struct
func (v *CompilerVisitor) VisitFuncionesStructsNativas(ctx *gramAntlr.FuncionesStructsNativasContext) interface{} {
	nombreFunc := ctx.ID_VARIABLE(2).GetText()
	var parametroDef gramAntlr.IDefParamsContext = ctx.DefParams()
	parametros := []*TupleStringSymbol{}

	if parametroDef != nil {
		for i := 0; i < len(parametroDef.AllID_VARIABLE()); i++ {
			id := parametroDef.ID_VARIABLE(i).GetText()
			tipo := parametroDef.Type_(i).GetText()
			switch tipo {
			case "int":
				parametros = append(parametros, &TupleStringSymbol{id, &Symbol{0, INT, true}})
			case "float64":
				parametros = append(parametros, &TupleStringSymbol{id, &Symbol{0.0, FLOAT64, true}})
			case "string":
				parametros = append(parametros, &TupleStringSymbol{id, &Symbol{"", STRING, true}})
			case "bool":
				parametros = append(parametros, &TupleStringSymbol{id, &Symbol{false, BOOL, true}})
			case "rune":
				parametros = append(parametros, &TupleStringSymbol{id, &Symbol{rune(0), RUNE, true}})
			default:
				v.Salida += fmt.Sprintf("Error semántico: tipo de parámetro no reconocido: %s\n", tipo)
				v.AgregarError(
					"Error tipo de parametro no coincide "+id,
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
		}
	}
	var tipoRetorno SymbolType
	var err error
	if ctx.ValRet() != nil {
		valRet := ctx.ValRet().GetText()
		tipoRetorno, err = StringToSymbolType(valRet)
		if err != nil {
			v.Salida += fmt.Sprintf("Error semántico: tipo de retorno no reconocido: %s\n", valRet)
			v.AgregarError(
				"Error tipo de retorno no reconocido",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	} else {
		tipoRetorno = VOID // Si no hay tipo de retorno, es void
	}
	// ------- Variable definida en el struct -------
	idVar := ctx.ID_VARIABLE(0).GetText()
	idStruct := ctx.ID_VARIABLE(1).GetText()

	structVar, _ := v.currentEnv.GetVariable(idStruct)
	if structVar == nil || structVar.Type != STRUCT {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no es un struct o no existe.\n", idStruct)
		v.AgregarError(
			"Error variable no es un struct",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}
	parametros = append(parametros, &TupleStringSymbol{idVar, structVar}) // Agregar el struct como parámetro
	var body gramAntlr.IBlockContext = ctx.Block()
	v.currentEnv.SetFunciones(nombreFunc, parametros, body, tipoRetorno, ctx.GetStart())

	if v.StructRelational[idStruct] != nil {
		v.StructRelational[idStruct] = append(v.StructRelational[idStruct], nombreFunc)
	} else {
		v.StructRelational[idStruct] = []string{nombreFunc}
	}

	return nil
}

// Llamada a función de struct
func (v *CompilerVisitor) VisitCallFunctionStructStmt(ctx *gramAntlr.CallFunctionStructStmtContext) interface{} {

	return v.Visit(ctx.VarCallFuncStruct())
}

func (v *CompilerVisitor) VisitCallFunctionStructValue(ctx *gramAntlr.CallFunctionStructValueContext) interface{} {

	return v.Visit(ctx.VarCallFuncStruct())
}

func (v *CompilerVisitor) VisitCallFunctionStruct(ctx *gramAntlr.CallFunctionStructContext) interface{} {
	nameStruct := ctx.ID_VARIABLE(0).GetText()     // Nombre del struct
	nameFuncStruct := ctx.ID_VARIABLE(1).GetText() // Nombre de la función del struct

	structVar, _ := v.currentEnv.GetVariable(nameStruct)
	if structVar == nil || structVar.Type != STRUCT {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no es un struct o no existe.\n", nameStruct)
		v.AgregarError(
			"Error variable no es un struct",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	var structVarBase string
	var findStructVar bool = false

	for key, value := range v.StructRelational {
		structVarBase = key

		for _, value2 := range value {
			if value2 == nameStruct {
				findStructVar = true
				break
			}
		}

	}

	if !findStructVar {
		v.Salida += fmt.Sprintf("Error semántico: función %s no encontrada en el struct %s.\n", nameFuncStruct, nameStruct)
		v.AgregarError(
			"Error funcion no encontrado en el struct",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	findStructVar = false

	var funcionesStruct []string = v.StructRelational[structVarBase]

	for _, funcName := range funcionesStruct {
		if funcName == nameFuncStruct {
			findStructVar = true
			break
		}
	}

	if !findStructVar {
		v.Salida += fmt.Sprintf("Error semántico: función %s no encontrada en el struct %s.\n", nameFuncStruct, nameStruct)
		v.AgregarError(
			"Error funcion no encontrada",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	parametros := []*TupleStringSymbol{}

	for i := 0; i < len(ctx.AllExpr()); i++ {
		value := v.Visit(ctx.Expr(i))

		if value == nil {
			v.Salida += fmt.Sprintf("Error-semántico: al llamar a la función %s del struct %s, el parámetro %d es nulo.\n", nameFuncStruct, nameStruct, i+1)
			v.AgregarError(
				"Error llamar a la funcion",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}

		switch value.(type) {
		case int:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, INT, false}})
		case float64:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, FLOAT64, false}})
		case string:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, STRING, false}})
		case bool:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, BOOL, false}})
		case rune:
			parametros = append(parametros, &TupleStringSymbol{ctx.Expr(i).GetText(), &Symbol{value, RUNE, false}})
		default:
			v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro no reconocido para la función %s del struct %s.\n", nameFuncStruct, nameStruct)
			v.AgregarError(
				"Error tipo de parametro no reconocido",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	}

	funcion, err := v.currentEnv.GetFuncion(nameFuncStruct)
	if err != nil {
		v.Salida += err.Error()
		return nil
	}

	parameters := funcion.Parameters
	var body gramAntlr.IBlockContext = funcion.Body
	tipoReturn := funcion.ValRet

	parametros = append(parametros, &TupleStringSymbol{parameters[len(parameters)-1].Key, structVar}) // Agregar el struct como parámetro
	if len(parameters) != len(parametros) {
		v.Salida += fmt.Sprintf("Error-semántico: número de parámetros incorrecto al llamar a la función %s del struct %s. Se esperaban %d, pero se recibieron %d.\n", nameFuncStruct, nameStruct, len(parameters), len(parametros))
		v.AgregarError(
			"Error numeor de parametros incorrecto",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"semántico",
		)
		return nil
	}

	enviroment := NewEnvironment(v.currentEnv)
	v.currentEnv = enviroment

	if tipoReturn != 7 {
		for i := 0; i < len(parametros); i++ {
			if parametros[i].Value.Type != parameters[i].Value.Type {
				v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro %v no coincide con el tipo esperado %v en la función %s del struct %s.\n",
					parametros[i].Value.Type, parameters[i].Value.Type, nameFuncStruct, nameStruct)
				v.AgregarError(
					"Error tipo de parametros no coincide",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			v.currentEnv.SetVariable(parameters[i].Key, parametros[i].Value.Value, parametros[i].Value.Type, parametros[i].Value.Mutable, true, ctx.GetStart())
		}
		posicionFinal := len(parameters) - 1
		v.currentEnv.SetVariable(parameters[posicionFinal].Key, structVar.Value, structVar.Type, structVar.Mutable, true, ctx.GetStart())

		valRet := v.Visit(body)

		if isValidType(valRet, tipoReturn) {
			v.currentEnv = v.currentEnv.Parent // Restaurar el entorno anterior
			return valRet
		} else {
			v.Salida += fmt.Sprintf("Error-semántico: tipo de retorno %v no coincide con el tipo esperado %v en la función %s del struct %s.\n",
				tipoReturn, fmt.Sprintf("%T", valRet), nameFuncStruct, nameStruct)
			v.AgregarError(
				"Error tipo de retorno no coincide",
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"semántico",
			)
			return nil
		}
	} else {
		for i := 0; i < len(parametros); i++ {
			if parametros[i].Value.Type != parameters[i].Value.Type {
				v.Salida += fmt.Sprintf("Error-semántico: tipo de parámetro %v no coincide con el tipo esperado %v en la función %s del struct %s.\n",
					parametros[i].Value.Type, parameters[i].Value.Type, nameFuncStruct, nameStruct)
				v.AgregarError(
					"Error tipo de parametros no coincide",
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"semántico",
				)
				return nil
			}
			v.currentEnv.SetVariable(parameters[i].Key, parametros[i].Value.Value, parametros[i].Value.Type, parametros[i].Value.Mutable, true, ctx.GetStart())
		}
		posicionFinal := len(parameters) - 1
		v.currentEnv.SetVariable(parameters[posicionFinal].Key, structVar.Value, structVar.Type, structVar.Mutable, true, ctx.GetStart())

		v.Visit(body) // Ejecutar el cuerpo de la función
	}
	v.currentEnv = v.currentEnv.Parent // Restaurar el entorno anterior
	return nil
}

//----------------------------- FUNCIONES AUXILIARES -----------------------------

// Validación de tipos básicos
func isValidType(value interface{}, typ SymbolType) bool {
	switch typ {
	case INT:
		_, ok := value.(int)
		return ok
	case FLOAT64:
		_, ok := value.(float64)
		return ok
	case STRING:
		_, ok := value.(string)
		return ok
	case BOOL:
		_, ok := value.(bool)
		return ok
	case RUNE:
		_, ok := value.(rune)
		return ok
	case STRUCT:
		_, ok := value.(*StructInstance)
		return ok
	}
	return false
}

func convertToStringSlice(list []interface{}) []string {
	result := make([]string, 0, len(list))
	for _, item := range list {
		if str, ok := item.(string); ok {
			result = append(result, str)
		} else {
			result = append(result, fmt.Sprintf("%v", item))
		}
	}
	return result
}

// Función para analizar el tipo, como "[][]string", "[]int", etc.
func getDimension_type(typeStr string) (string, int, error) {
	// Contar las apariciones de "[]"
	re := regexp.MustCompile(`\[\]`)
	dimensions := len(re.FindAllString(typeStr, -1))

	// Eliminar todos los "[]" para obtener el tipo base
	baseType := re.ReplaceAllString(typeStr, "")
	baseType = strings.TrimSpace(baseType)

	// Validación opcional de tipos conocidos (puedes expandir esta lista)
	validTypes := map[string]bool{
		"int": true, "float": true, "string": true, "bool": true,
	}

	if _, ok := validTypes[baseType]; !ok {
		return "", 0, fmt.Errorf("tipo base '%s' no es válido", baseType)
	}

	return baseType, dimensions, nil
}
