package compile

import (
	"OLC2CLIENTE/compile/expresiones/nativas"
	"OLC2CLIENTE/compile/expresiones/operaciones"
	"OLC2CLIENTE/compile/print"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"
	"reflect"
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
}

// Constructor opcional
func NewCompilerVisitor() *CompilerVisitor {
	v := &CompilerVisitor{
		BasegramaticaVisitor: &gramAntlr.BasegramaticaVisitor{},
		Salida:               "",
		currentEnv:           NewEnvironment(nil), // Entorno raíz
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

// -------------------------------- DECLARACIONES --------------------------------------
// VisitVarDcl
func (v *CompilerVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	return v.Visit(ctx.VarDcl())
}

// 'var' ID type '=' expr ';'
func (v *CompilerVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	expr := ctx.Expr()
	value := v.Visit(expr)

	// Convertir tipo string a SymbolType
	symbolType, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		return nil
	}

	// Manejar valores nulos
	if value == nil {
		switch symbolType {
		case INT:
			value = 0
		case FLOAT64:
			value = 0.0
		case STRING:
			value = ""
		case BOOL:
			value = false
		case RUNE:
			value = '\000'
		}
	}

	// Convertir int a float64 si necesario
	if num, ok := value.(int); ok && symbolType == FLOAT64 {
		value = float64(num)
	} else if !isValidType(value, symbolType) {
		v.Salida += fmt.Sprintf("Error semántico: tipos incompatibles para variable %s\n", id)
		return nil
	}

	v.currentEnv.SetVariable(id, value, symbolType, false, true, ctx.GetStart())
	return nil
}

// 'var' ID type ';'
func (v *CompilerVisitor) VisitVarDclWithTypeOnly(ctx *gramAntlr.VarDclWithTypeOnlyContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()

	symbolType, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
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
	default:
		v.Salida += fmt.Sprintf("Error: tipo %s no soporta valor por defecto\n", typeStr)
		return nil
	}

	v.currentEnv.SetVariable(id, defaultValue, symbolType, false, true, ctx.GetStart())
	return nil
}

// 'var' ID ':=' expr ';'
func (v *CompilerVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())

	// Determinar tipo basado en valor
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
		return nil
	}

	v.currentEnv.SetVariable(id, value, symbolType, true, true, ctx.GetStart())
	return nil
}

// ----------------------------- Acceso a arreglos -----------------------------
// VisitArrayAccess - 'ID_VARIABLE' '[' expr ']' '=' expr
func (v *CompilerVisitor) VisitArrayAccess(ctx *gramAntlr.ArrayAccessContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		return nil
	}

	listaBase, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		return nil
	}

	for i := 0; i < len(ctx.AllExpr()); i++ {
		index := v.Visit(ctx.AllExpr()[i])
		if idx, ok := index.(int); ok {
			if i == len(ctx.AllExpr())-2 {
				// Asignar valor en la posición final
				value := v.Visit(ctx.AllExpr()[i+1])
				listaBase[idx] = value
				return nil
			} else {
				// Ir al siguiente nivel del arreglo anidado
				nextLevel, ok := listaBase[idx].([]interface{})
				if !ok {
					v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la posición %d no es un arreglo.\n", idx)
					return nil
				}
				listaBase = nextLevel
			}
		} else {
			v.Salida += "Error-semántico: índice debe ser un entero.\n"
			return nil
		}
	}

	return nil
}

// VisitArrayAccessSimple - 'ID_VARIABLE' '[' expr ']'
func (v *CompilerVisitor) VisitArrayAccessSimple(ctx *gramAntlr.ArrayAccessSimpleContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		return nil
	}

	listaBase, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		return nil
	}

	contador := 0
	for _, expr := range ctx.AllExpr() {
		index := v.Visit(expr)
		if idx, ok := index.(int); ok {
			dataList := listaBase[idx]
			if contador == len(ctx.AllExpr())-1 {
				return dataList
			}

			if nextList, ok := dataList.([]interface{}); ok {
				listaBase = nextList
			} else {
				return dataList
			}
			contador++
		} else {
			v.Salida += "Error-semántico: índice debe ser un entero.\n"
			return nil
		}
	}

	return nil
}

// VisitArrayFindIndex - 'ID_VARIABLE' '.' 'findIndex' '(' expr ')'
func (v *CompilerVisitor) VisitArrayFindIndex(ctx *gramAntlr.ArrayFindIndexContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	value := v.Visit(ctx.Expr())

	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
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

	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
		return nil
	}

	if strValue, ok := value.(string); ok {
		result := strings.Join(convertToStringSlice(tempList), strValue)
		return result
	} else {
		v.Salida += "Error-semántico: al unir el arreglo, el valor no es un string.\n"
		return nil
	}
}

// VisitArrayLength - 'ID_VARIABLE' '.' 'length' '(' posicion* ')'
func (v *CompilerVisitor) VisitArrayLength(ctx *gramAntlr.ArrayLengthContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	variable, err := v.currentEnv.GetVariable(id)
	if err != nil {
		v.Salida += fmt.Sprintf("Error semántico: variable %s no declarada\n", id)
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
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
					return nil
				}
				nextItem := listaBase[idx]
				if nextList, ok := nextItem.([]interface{}); ok {
					listaBase = nextList
				} else {
					v.Salida += "Error-semántico: al acceder al arreglo, la posición no es un arreglo.\n"
					return nil
				}
			} else {
				v.Salida += "Error-semántico: índice debe ser un entero.\n"
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
		return nil
	}

	tempList, ok := variable.Value.([]interface{})
	if !ok {
		v.Salida += fmt.Sprintf("Error-semántico: al acceder al arreglo, la variable %s no es un arreglo.\n", id)
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
	return nil
}

// ------------------------------ DECLARACION SLICE ----------------------------------

// VisitVarDeclSliceStmt - 'var' ID_VARIABLE slice '=' sliceValores ';'
func (v *CompilerVisitor) VisitVarDeclSliceStmt(ctx *gramAntlr.VarDeclSliceStmtContext) interface{} {
	return v.Visit(ctx.VarDclSlice())
}

// VisitSliceValores - 'ID_VARIABLE' ':' type '=' sliceValores
func (v *CompilerVisitor) VisitSliceValores(ctx *gramAntlr.SliceValoresContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	typo, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
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

// VisitSliceVacio - 'ID_VARIABLE' ':' type '=' sliceVacio
func (v *CompilerVisitor) VisitSliceVacio(ctx *gramAntlr.SliceVacioContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	symbolType, err := parseSymbolType(typeStr)
	if err != nil {
		v.Salida += fmt.Sprintf("Error: tipo %s no válido\n", typeStr)
		return nil
	}

	v.currentEnv.SetVariable(id, make([]interface{}, 0), symbolType, false, true, ctx.GetStart())

	return nil
}

// VisitSliceContenido - contenido de un slice con expresiones
func (v *CompilerVisitor) VisitSliceContenido(ctx *gramAntlr.SliceContenidoContext) interface{} {
	arrayTemp := []interface{}{}
	for _, expr := range ctx.AllExpr() {
		arrayTemp = append(arrayTemp, v.Visit(expr))
	}
	return arrayTemp
}

// VisitSliceContenidoSlice - contenido de un slice anidado
func (v *CompilerVisitor) VisitSliceContenidoSlice(ctx *gramAntlr.SliceContenidoSliceContext) interface{} {
	arrayTemp := []interface{}{}
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
		return nil
	}

	mutabilidad := sym.Mutable
	tipo := sym.Type
	contextStart := ctx.GetStart()

	if value == nil {
		v.Salida += fmt.Sprintf("Error-semántico: al asignar el valor a la variable '%s', el valor es nulo.\n", id)
		return nil
	}

	// Validar si es un slice
	if slice, ok := value.([]interface{}); ok {
		if len(slice) == 0 || isValidType(slice[0], tipo) {
			v.currentEnv.SetVariable(id, value, tipo, mutabilidad, false, contextStart)
			return nil
		}
		v.Salida += fmt.Sprintf("Error-semántico: tipos no compatibles para el slice asignado a %s.\n", id)
		return nil
	}

	// Validación normal
	if !isValidType(value, tipo) {
		if mutabilidad {
			v.currentEnv.SetVariable(id, value, tipo, mutabilidad, false, contextStart)
			return nil
		}
		v.Salida += fmt.Sprintf("Error-semántico: la variable %s no es mutable y no se puede asignar un nuevo valor.\n", id)
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
	return nil
}

// VarInc
func (v *CompilerVisitor) VisitVarInc(ctx *gramAntlr.VarIncContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	token := ctx.GetStart()

	sym, _ := v.currentEnv.GetVariable(id)
	typ := sym.Type

	// Solo se permite ++ y -- para INT (0) y FLOAT64 (1)
	if typ != 0 && typ != 1 {
		v.Salida += "Error-semántico: el tipo de variable no acepta operador ++ o --."
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
	fmt.Println("ENTRE EN VISIT IF ONLY")

	// Evaluar la condición
	value := v.Visit(ctx.Expr())
	if value == nil {
		v.Salida += "Error semántico: condición del if es nil\n"
		return nil
	}
	cond, ok := value.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del if, no es un booleano.\n"
		return nil
	}

	// Si la condición es verdadera, ejecutar el bloque[0]
	if cond {
		// Nuevo entorno anidado
		newEnv := NewEnvironment(v.currentEnv)
		v.currentEnv = newEnv

		// Visitar el bloque del 'if'
		result := v.Visit(ctx.Block(0))
		fmt.Println("RESULTADO DEL IF:", result)

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
			result := v.Visit(blocks[1])
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

	fmt.Println("SALI NIL EN VISIT IF ONLY")
	return nil
}

// -------------------- Produccion IF ELSE IF --------------------
func (v *CompilerVisitor) VisitIfAnidado(ctx *gramAntlr.IfAnidadoContext) interface{} {
	// Evaluar la condición
	value := v.Visit(ctx.Expr())
	cond, ok := value.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del if, no es un booleano.\n"
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
	return "break"
}

// -------------------- VisitContinue --------------------
func (v *CompilerVisitor) VisitContinueStmt(ctx *gramAntlr.ContinueStmtContext) interface{} {
	return "continue"
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
	} else {
		for condBool {
			result := v.Visit(ctx.Block())

			if str, ok := result.(string); ok {
				if str == "break" {
					break
				} else if str == "continue" {
					condition = v.Visit(ctx.Expr())
					if condBool, ok = condition.(bool); !ok {
						v.Salida += "Error-semántico: al reevaluar la condición del for tras un 'continue', no es un booleano."
					}
					continue
				} else if str == "Excepcion___Return_Void" {
					return str
				} else {
					return str
				}
			}

			// Recalcular la condición
			condition = v.Visit(ctx.Expr())
			if condBool, ok = condition.(bool); !ok {
				v.Salida += "Error-semántico: al reevaluar la condición del for, no es un booleano."
			}
		}
	}

	return nil
}

// VisitForAsignacion
func (v *CompilerVisitor) VisitForAsignacion(ctx *gramAntlr.ForAsignacionContext) interface{} {
	// Declaración inicial
	v.Visit(ctx.VarDcl())

	// Evaluación inicial de la condición
	condition := v.Visit(ctx.Expr())

	condBool, ok := condition.(bool)
	if !ok {
		v.Salida += "Error-semántico: al evaluar la condición del for, no es un booleano."
	}

	for condBool {
		// Ejecutar bloque del for
		result := v.Visit(ctx.Block())
		fmt.Println("Resultado del bloque del for:", result)

		if str, ok := result.(string); ok {
			if str == "break" {
				break
			} else if str == "continue" {
				// Evaluar la asignación y la nueva condición
				v.Visit(ctx.VarAsign())
				condition = v.Visit(ctx.Expr())
				if condBool, ok = condition.(bool); !ok {
					v.Salida += "Error-semántico: al reevaluar la condición del for tras un 'continue', no es un booleano."
				}
				continue
			} else if str == "Excepcion___Return_Void" {
				return str
			} else {
				return str
			}
		}

		// Ejecutar asignación final del for
		v.Visit(ctx.VarAsign())

		// Reevaluar condición
		condition = v.Visit(ctx.Expr())
		if condBool, ok = condition.(bool); !ok {
			v.Salida += "Error-semántico: al reevaluar la condición del for, no es un booleano."
		}
	}

	return nil
}

// VisitForRange
func (v *CompilerVisitor) VisitForRange(ctx *gramAntlr.ForRangeContext) interface{} {
	sym, _ := v.currentEnv.GetVariable(ctx.ID_VARIABLE(2).GetText())

	slice, err := sym.Value.([]interface{})
	if !err {
		v.Salida += fmt.Sprintf("Error-semántico: al iterar sobre el rango, la variable %s no es un slice.\n", ctx.ID_VARIABLE(2).GetText())
		return nil
	}

	index := ctx.ID_VARIABLE(0).GetText()
	value := ctx.ID_VARIABLE(1).GetText()

	v.currentEnv.SetVariable(index, 0, sym.Type, false, true, ctx.GetStart())         // Inicializar índice
	v.currentEnv.SetVariable(value, nil, sym.Type, sym.Mutable, true, ctx.GetStart()) // Inicializar valor

	newEnv := NewEnvironment(v.currentEnv)
	v.currentEnv = newEnv

	for i, val := range slice {
		v.currentEnv.SetVariable(index, i, INT, false, false, ctx.GetStart())
		v.currentEnv.SetVariable(value, val, sym.Type, false, false, ctx.GetStart())

		result := v.Visit(ctx.Block())

		if resStr, ok := result.(string); ok {
			if resStr == "break" {
				break
			} else if resStr == "continue" {
				continue
			} else if resStr == "Excepcion___Return_Void" {
				return result
			} else {
				return result
			}
		}
	}
	// Restaurar el entorno anterior
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
	v.conditionExpr = v.Visit(ctx.Expr()) // Evaluar la condicion-tipo del switch
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

	for _, instruccion := range ctx.AllInstrucciones() {
		v.Visit(instruccion)
	}

	v.currentEnv = newEnv.Parent

	return nil
}

// ---------------------------- FUNCIONES NATIVAS -----------------------------
// VisitIntToString
func (v *CompilerVisitor) VisitIntToString(ctx *gramAntlr.IntToStringContext) interface{} {
	return v.funcionesNativasVisitor.VisitIntToString(ctx)
}

// VisitfloatToString
func (v *CompilerVisitor) VisitFloatToString(ctx *gramAntlr.FloatToStringContext) interface{} {
	return v.funcionesNativasVisitor.VisitFloatToString(ctx)
}

// VisitReflectType
func (v *CompilerVisitor) VisitReflectType(ctx *gramAntlr.ReflectTypeContext) interface{} {
	return v.funcionesNativasVisitor.VisitReflectType(ctx)
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
