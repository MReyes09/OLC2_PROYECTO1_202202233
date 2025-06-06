package compile

import (
	"OLC2CLIENTE/compile/expresiones/operaciones"
	"OLC2CLIENTE/compile/print"
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
	printVisitor                    *print.PrintVisitor
	operacionesVisitor              *operaciones.OperacionesVisitor // Visitor para operaciones aritméticas
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

	/*
		// Manejar slices
		if slice, ok := value.([]interface{}); ok {
			dataType := getSliceType(slice)
			symbolType, _ := parseSymbolType(dataType)
			v.currentEnv.SetVariable(id, value, symbolType, true, true, ctx.GetStart())
			return nil
		}
	*/

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
	sym, _ := v.currentEnv.GetVariable(id)
	mutabilidad := sym.Mutable

	context_start := ctx.GetStart()

	if value == nil {
		v.Salida += fmt.Sprintf("Error-semántico: al asignar el valor a la variable, el valor es nulo.")
		return nil
	}

	// validar cuando es un slice
	if !isValidType(value, sym.Type) {

		if mutabilidad {
			v.currentEnv.SetVariable(id, value, sym.Type, mutabilidad, false, context_start)
			return nil
		}

		v.Salida += fmt.Sprintf("Error-semántico: la variable %s no es mutable y no se puede asignar un nuevo valor.\n", id)
		return nil
	}

	v.currentEnv.SetVariable(id, value, sym.Type, mutabilidad, false, context_start)

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

// ------------
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
