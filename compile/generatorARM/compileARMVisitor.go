package generatorARM

import (
	fragmentvisitor "OLC2CLIENTE/compile/generatorARM/fragmentVisitor"
	"OLC2CLIENTE/compile/generatorARM/registros"
	"OLC2CLIENTE/compile/generatorARM/traductor"
	"OLC2CLIENTE/gramatica/gramAntlr"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type CompileARMVisitor struct {
	*gramAntlr.BasegramaticaVisitor
	MetaData              map[string]DataFuncion
	C                     traductor.GeneratorARMInstructions
	InFunction            string
	FragmentPointerOffSet int
	ReturnLabels          string
}

type DataFuncion struct {
	FrameSize int
	TypeRet   traductor.TypeObject
}

func NewCompileARMVisitor() *CompileARMVisitor {
	v := &CompileARMVisitor{
		BasegramaticaVisitor:  &gramAntlr.BasegramaticaVisitor{},
		MetaData:              make(map[string]DataFuncion),
		C:                     *traductor.NewGeneratorARMInstructions(),
		InFunction:            "",
		FragmentPointerOffSet: 0,
		ReturnLabels:          "",
	}
	return v
}

func (v *CompileARMVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *CompileARMVisitor) VisitInicio(ctx *gramAntlr.InicioContext) interface{} {
	for _, stmtCtx := range ctx.AllInstrucciones() {
		v.Visit(stmtCtx)
	}
	return nil
}

// --------------------------------- DECLARACIONES ----------------------------------
func (v *CompileARMVisitor) VisitIdentifier(ctx *gramAntlr.IdentifierContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	offSet, object := v.C.GetObject(id)
	fmt.Println("Encontramos el objeto:", object.Id_, "con tipo:", object.Type_, "y offset:", offSet)
	if v.InFunction != "" {
		v.C.Mov(registros.X0, offSet)
		v.C.Sub(registros.X0, registros.FP, registros.X0)
		v.C.LDR(registros.X0, registros.X0, 0) //0 es el por defecto
		v.C.Push(registros.X0)
		CloneObject := v.C.CloneObject(object)
		CloneObject.Id_ = ""
		v.C.PushObjectStack(CloneObject)
		return nil

	}
	v.C.Mov(registros.X0, offSet)
	v.C.Add(registros.X0, registros.SP, registros.X0)
	v.C.LDR(registros.X0, registros.X0, 0)
	v.C.Push(registros.X0)
	CloneObject := v.C.CloneObject(object)
	CloneObject.Id_ = ""
	v.C.PushObjectStack(CloneObject)
	return nil
}

// VisitVarDcl
func (v *CompileARMVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	return v.Visit(ctx.VarDcl())
}

// 'mut' ID type '=' expr ';'
func (v *CompileARMVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	expr := ctx.Expr()
	v.Visit(expr)
	v.C.COMENT(fmt.Sprintf("Declaracion explicita: %s con tipo %s", id, typeStr))

	if v.InFunction != "" {
		LocalObjecto := v.C.GetFrameLocal(v.FragmentPointerOffSet)
		ValorObjecto := v.C.POPOBJECT(registros.X0)
		v.C.Mov(registros.X1, v.FragmentPointerOffSet*8)
		v.C.Sub(registros.X1, registros.FP, registros.X1)
		v.C.Str(registros.X0, registros.X1)
		LocalObjecto.Type_ = ValorObjecto.Type_
		v.FragmentPointerOffSet++
		return nil
	}

	v.C.TagObjecto(id)
	return nil
}

// 'mut' ID ':=' expr ';'
func (v *CompileARMVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	expr := ctx.Expr()
	v.Visit(expr)
	v.C.COMENT(fmt.Sprintf("Declaracion implicita: %s", id))

	if v.InFunction != "" {
		LocalObjecto := v.C.GetFrameLocal(v.FragmentPointerOffSet)
		ValorObjecto := v.C.POPOBJECT(registros.X0)
		v.C.Mov(registros.X1, v.FragmentPointerOffSet*8)
		v.C.Sub(registros.X1, registros.FP, registros.X1)
		v.C.Str(registros.X0, registros.X1)
		LocalObjecto.Type_ = ValorObjecto.Type_
		v.FragmentPointerOffSet++
		return nil
	}

	v.C.TagObjecto(id)
	return nil
}

// 'mut' ID type ';'
func (v *CompileARMVisitor) VisitVarDclWithTypeOnly(ctx *gramAntlr.VarDclWithTypeOnlyContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()
	v.C.COMENT(fmt.Sprintf("Declaracion no inicializada: %s con tipo %s", id, typeStr))
	switch typeStr {
	case "int":
		var IntObject = v.C.IntObject()
		v.C.PushConst(IntObject, 0)
		break
	case "float64":
		var FloatObject = v.C.FloatObject()
		v.C.PushConst(FloatObject, 0.0)
		break
	case "string":
		var StringObject = v.C.StrObject()
		v.C.PushConst(StringObject, "")
		break
	case "bool":
		var BoolObject = v.C.BoolObject()
		v.C.PushConst(BoolObject, 0)
		break
	default:
		panic(fmt.Sprintf("Tipo no soportado: %s", typeStr))
	}

	if v.InFunction != "" {
		LocalObjecto := v.C.GetFrameLocal(v.FragmentPointerOffSet)
		fmt.Println("Obtuve var:", LocalObjecto.Id_, "tipo:", LocalObjecto.Type_)
		ValorObjecto := v.C.POPOBJECT(registros.X0)
		fmt.Println("En ultima posicion:", ValorObjecto.Id_, "tipo:", ValorObjecto.Type_)
		v.C.Mov(registros.X1, v.FragmentPointerOffSet*8)
		v.C.Sub(registros.X1, registros.FP, registros.X1)
		v.C.Str(registros.X0, registros.X1)
		LocalObjecto.Type_ = ValorObjecto.Type_
		fmt.Println("Verificamos si actualizo")
		temp := v.C.GetFrameLocal(v.FragmentPointerOffSet)
		fmt.Println("Obtuve var:", temp.Id_, "tipo:", temp.Type_)
		v.FragmentPointerOffSet++
		return nil
	}

	v.C.TagObjecto(id)
	return nil
}

// --------------------------------- EXPRESIONES -----------------------------------

// --------------------------------- TIPO DE DATOS -----------------------------------
func (v *CompileARMVisitor) VisitString(ctx *gramAntlr.StringContext) interface{} {
	v.C.COMENT(fmt.Sprintf("cadena: %s", ctx.STRING().GetText()))

	texto := ctx.GetText()
	if strings.HasPrefix(texto, "\"") && strings.HasSuffix(texto, "\"") {
		texto = texto[1 : len(texto)-1] // eliminar comillas
	}

	texto = strings.ReplaceAll(texto, `\\`, `\`)
	texto = strings.ReplaceAll(texto, `\n`, "\n")
	texto = strings.ReplaceAll(texto, `\t`, "\t")
	texto = strings.ReplaceAll(texto, `\r`, "\r")
	texto = strings.ReplaceAll(texto, `\"`, `"`)

	strObject := v.C.StrObject()
	v.C.PushConst(strObject, texto)

	return nil
}

func (v *CompileARMVisitor) VisitInteger(ctx *gramAntlr.IntegerContext) interface{} {
	value, err := strconv.Atoi(ctx.GetText())

	if err != nil {
		panic(fmt.Sprintf("Error al convertir a entero: %s", ctx.GetText()))
	}

	v.C.COMENT(fmt.Sprintf("Entero: %d", value))

	IntObject := v.C.IntObject()
	v.C.PushConst(IntObject, value)

	return nil
}

func (v *CompileARMVisitor) VisitDouble(ctx *gramAntlr.DoubleContext) interface{} {
	value, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(fmt.Sprintf("Error al convertir a float: %s", ctx.GetText()))
	}
	v.C.COMENT(fmt.Sprintf("Float: %f", value))
	floatObject := v.C.FloatObject()
	v.C.PushConst(floatObject, value)

	return nil
}

func (v *CompileARMVisitor) VisitBoolean(ctx *gramAntlr.BooleanContext) interface{} {

	value := 0
	valueStr := strings.ToLower(ctx.GetText())
	if valueStr == "true" {
		value = 1
	} else {
		value = 0
	}

	v.C.COMENT(fmt.Sprintf("Booleano: %d", value))
	boolObject := v.C.BoolObject()
	v.C.PushConst(boolObject, value)
	return nil
}

// -------------------------------- OPERADORES ARITMETICOS --------------------------------
func (v *CompileARMVisitor) VisitParens(ctx *gramAntlr.ParensContext) interface{} {
	v.Visit(ctx.Expr())
	return nil
}

func (v *CompileARMVisitor) VisitNegate(ctx *gramAntlr.NegateContext) interface{} {
	v.Visit(ctx.Expr())

	// Obtener el tipo de la cima de la pila
	top := v.C.GetTopObjectStack()
	var reg string

	// Elegir el registro adecuado
	if top.Type_ == traductor.Float {
		reg = registros.D0
	} else {
		reg = registros.X0
	}

	// Hacer POP con el registro correcto
	value := v.C.POPOBJECT(reg)

	// Aplicar la negación según el tipo
	if value.Type_ == traductor.Int {
		v.C.Neg(registros.X0, registros.X0)
		v.C.Push(registros.X0)
		v.C.PushObjectStack(v.C.CloneObject(value))
	} else if value.Type_ == traductor.Float {
		v.C.COMENT("Float64 Negativo")
		v.C.FNeg(registros.D0, registros.D0)
		v.C.Push(registros.D0)
		v.C.PushObjectStack(v.C.CloneObject(value))
	}

	return nil
}

func (v *CompileARMVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {

	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))

	operador := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	DerechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float

	var Derecha traductor.ObjectStack
	if DerechaEsFloat {
		Derecha = v.C.POPOBJECT(registros.D1)
	} else {
		Derecha = v.C.POPOBJECT(registros.X1)
	}

	IzquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	var Izquierda traductor.ObjectStack

	if IzquierdaEsFloat {
		Izquierda = v.C.POPOBJECT(registros.D0)
	} else {
		Izquierda = v.C.POPOBJECT(registros.X0)
	}

	switch operador {
	case "+":
		// int + int
		if Derecha.Type_ == traductor.Int && Izquierda.Type_ == traductor.Int {
			v.C.Add(registros.X0, registros.X0, registros.X1)
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(Izquierda))
		} else if Derecha.Type_ == traductor.StringType && Izquierda.Type_ == traductor.StringType {
			// string + string
			v.C.ConcatString()
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(Izquierda))
		} else if Derecha.Type_ == traductor.Float || Izquierda.Type_ == traductor.Float {
			// float64 + ...
			if !DerechaEsFloat {
				v.C.COMENT("Convertir Derecha a float64")
				v.C.Scvtf(registros.D1, registros.X1)
			}
			if !IzquierdaEsFloat {
				v.C.COMENT("Convertir Izquierda a float64")
				v.C.Scvtf(registros.D0, registros.X0)
			}
			v.C.COMENT("Suma de float64")
			v.C.Fadd(registros.D0, registros.D0, registros.D1)
			v.C.Push(registros.D0)
			v.C.PushObjectStack(v.C.FloatObject())
		}
	case "-":
		// int - int
		if Derecha.Type_ == traductor.Int && Izquierda.Type_ == traductor.Int {
			v.C.Sub(registros.X0, registros.X0, registros.X1)
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(Izquierda))
		} else if Derecha.Type_ == traductor.Float || Izquierda.Type_ == traductor.Float {
			// float64 - ...
			if !DerechaEsFloat {
				v.C.COMENT("Convertir Derecha a float64")
				v.C.Scvtf(registros.D1, registros.X1)
			}
			if !IzquierdaEsFloat {
				v.C.COMENT("Convertir Izquierda a float64")
				v.C.Scvtf(registros.D0, registros.X0)
			}
			v.C.Fsub(registros.D0, registros.D0, registros.D1)
			v.C.Push(registros.D0)
			v.C.PushObjectStack(v.C.FloatObject())
		}
	}

	return nil
}

func (v *CompileARMVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))

	operador := ctx.GetOp().GetText()
	derechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float

	var derecha traductor.ObjectStack
	if derechaEsFloat {
		derecha = v.C.POPOBJECT(registros.D1)
	} else {
		derecha = v.C.POPOBJECT(registros.X1)
	}

	izquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	var izquierda traductor.ObjectStack
	if izquierdaEsFloat {
		izquierda = v.C.POPOBJECT(registros.D0)
	} else {
		izquierda = v.C.POPOBJECT(registros.X0)
	}

	switch operador {
	case "*":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.Mul(registros.X0, registros.X0, registros.X1)
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(izquierda))
		} else if derecha.Type_ == traductor.Float || izquierda.Type_ == traductor.Float {
			if !izquierdaEsFloat {
				v.C.Scvtf(registros.D0, registros.X0)
			}
			if !derechaEsFloat {
				v.C.Scvtf(registros.D1, registros.X1)
			}
			v.C.FMul(registros.D0, registros.D0, registros.D1)
			v.C.Push(registros.D0)
			if izquierdaEsFloat {
				v.C.PushObjectStack(v.C.CloneObject(izquierda))
			} else {
				v.C.PushObjectStack(v.C.CloneObject(derecha))
			}
		}

	case "/":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.Div(registros.X0, registros.X0, registros.X1)
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(izquierda))
		} else if derecha.Type_ == traductor.Float || izquierda.Type_ == traductor.Float {
			if !izquierdaEsFloat {
				v.C.Scvtf(registros.D0, registros.X0)
			}
			if !derechaEsFloat {
				v.C.Scvtf(registros.D1, registros.X1)
			}
			v.C.FDiv(registros.D0, registros.D0, registros.D1)
			v.C.Push(registros.D0)
			if izquierdaEsFloat {
				v.C.PushObjectStack(v.C.CloneObject(izquierda))
			} else {
				v.C.PushObjectStack(v.C.CloneObject(derecha))
			}
		}

	case "%":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.Mod(registros.X0, registros.X0, registros.X1)
			v.C.Push(registros.X0)
			v.C.PushObjectStack(v.C.CloneObject(izquierda))
		}
	}
	return nil
}

// instrucciones: imprimir         # PrintStmt
func (v *CompileARMVisitor) VisitPrintStmt(ctx *gramAntlr.PrintStmtContext) interface{} {
	v.Visit(ctx.Imprimir())
	return nil
}

// imprimir: 'println(' (expr (',' expr)*)?')' ';'? #Println
func (v *CompileARMVisitor) VisitPrintln(ctx *gramAntlr.PrintlnContext) interface{} {
	v.C.COMENT("Función embebida: println")

	for _, exprCtx := range ctx.AllExpr() {
		v.Visit(exprCtx)
		isFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
		var value traductor.ObjectStack
		if isFloat {
			value = v.C.POPOBJECT(registros.D0)
		} else {
			value = v.C.POPOBJECT(registros.X0)
		}
		switch value.Type_ {
		case traductor.Int:
			v.C.ImprimirEntero(registros.X0)
		case traductor.Float:
			v.C.ImprimirDecimal()
		case traductor.StringType:
			v.C.ImprimirCadena(registros.X0)
		case traductor.Rune:
			//v.C.ImprimirCaracter(registros.X0)
		case traductor.Bool:
			v.C.ImprimirBooleano(registros.X0)
		case traductor.Slice:
			//v.C.ImprimirArreglo(value.TipoElemento_)
		}
		v.C.Espaciado()
	}
	v.C.SaltoLinea()
	return nil
}

// | functions                 # FunctionStmt
func (v *CompileARMVisitor) VisitFunctionStmt(ctx *gramAntlr.FunctionStmtContext) interface{} {
	return v.Visit(ctx.Functions())
}

// functions: 'func' ID_VARIABLE '(' (ID_VARIABLE type (',' ID_VARIABLE type)*)? ')' valRet? block # Funciones
func (v *CompileARMVisitor) VisitFunciones(ctx *gramAntlr.FuncionesContext) interface{} {
	//Manjamos estados de la pila
	baseOffSet := 2
	paramsOffSet := 0

	//Validamos los parametros
	if len(ctx.AllID_VARIABLE()) > 1 {
		paramsOffSet = len(ctx.AllID_VARIABLE()) - 1
	}

	// Creamos el visitor para manejar los fragmentos de código y el manejo de offsets (desplazamientos)
	fragmentVisitor := fragmentvisitor.NewFragmentVisitor(baseOffSet + paramsOffSet)

	for _, child := range ctx.Block().GetChildren() {
		// Debemos verificar que los hijos sean del tipo ParseTree
		// y luego visitar cada uno de ellos con el fragmentVisitor
		if node, ok := child.(antlr.ParseTree); ok {
			fragmentVisitor.Visit(node)
		}
	}

	fragment := fragmentVisitor.Fragment

	localOffSet := len(fragment)
	returnOffSet := 1
	sizeFragment := baseOffSet + paramsOffSet + localOffSet + returnOffSet
	nameFunction := ctx.ID_VARIABLE(0).GetText()

	typeReturn := traductor.Void

	if ctx.ValRet() != nil {
		typeReturn = TypeOfReturn(ctx.ValRet().GetText())
	}

	v.MetaData[nameFunction] = DataFuncion{
		FrameSize: sizeFragment,
		TypeRet:   typeReturn,
	}

	// Mantenemos las instrucciones previas y las reiniciamos
	instructionsPrev := v.C.Instrucciones
	v.C.Instrucciones = []string{}

	//Validamos los parametros
	if len(ctx.AllID_VARIABLE()) > 1 {

		for i := 0; i < len(ctx.AllID_VARIABLE()); i++ {
			v.C.PushObjectStack(traductor.ObjectStack{
				Type_:   TypeOfReturn(ctx.Type_(i).GetText()),
				Id_:     ctx.ID_VARIABLE(i + 1).GetText(),
				Offset_: i + 1 + baseOffSet,
				Length_: 8,
			})
		}
	}

	for _, instruction := range fragment {
		v.C.PushObjectStack(traductor.ObjectStack{
			Type_:   traductor.Void,
			Id_:     instruction.Name,
			Offset_: instruction.Offset,
			Length_: 8,
		})
	}

	v.InFunction = nameFunction
	v.FragmentPointerOffSet = 0
	v.ReturnLabels = v.C.GetLabel()
	v.C.SetLabel(nameFunction)

	if nameFunction == "main" {
		v.C.Instrucciones = append(v.C.Instrucciones, "stp x29, x30, [sp, #-16]!")
		v.C.Instrucciones = append(v.C.Instrucciones, "mov x29, sp")
	}

	//Recorremos el bloque de instrucciones
	for _, child := range ctx.Block().GetChildren() {
		// Debemos verificar que los hijos sean del tipo ParseTree
		// y luego visitar cada uno de ellos con el fragmentVisitor
		if node, ok := child.(antlr.ParseTree); ok {
			v.Visit(node)
		}
	}

	v.C.SetLabel(v.ReturnLabels)
	if nameFunction != "main" {
		v.C.Add(registros.X0, registros.FP, registros.XZR)
		v.C.LDR(registros.LR, registros.X0, 0)
		v.C.Br(registros.LR)
	} else {
		v.C.Instrucciones = append(v.C.Instrucciones, "LDP x29, x30, [sp], #16")
		v.C.Instrucciones = append(v.C.Instrucciones, "RET")
	}

	if paramsOffSet > 0 {
		for i := 0; i < paramsOffSet; i++ {
			v.C.POPOBJECT2()
		}
	}

	v.C.FuncionesInstrucciones = append(v.C.FuncionesInstrucciones, v.C.Instrucciones...)
	v.C.Instrucciones = instructionsPrev
	v.InFunction = ""

	return nil
}

// --------------------- FUNCIONES COMPLEMENTO ---------------------
func TypeOfReturn(tipo string) traductor.TypeObject {
	switch strings.ToLower(tipo) {
	case "int":
		return traductor.Int
	case "float64":
		return traductor.Float
	case "string":
		return traductor.StringType
	case "bool":
		return traductor.Bool
	case "rune":
		return traductor.Rune
	default:
		panic(fmt.Sprintf("Tipo no soportado: %s", tipo))
	}
}
