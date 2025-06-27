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
	PositionFramePointer  int
	Depth                 int   // Profundidad de anidamiento, si es necesario
	UsedPositions         []int // Posiciones usadas en el frame pointer
	LabelCompareCount     int   // Contador de etiquetas de comparación
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
		PositionFramePointer:  1,
		Depth:                 -1,             // Inicialmente no hay anidamiento
		UsedPositions:         make([]int, 0), // Inicializamos el slice de posiciones usadas
		LabelCompareCount:     0,
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
	// Verificamos si el identificador existe en nuestra tabla de símbolos
	offSet, object := v.C.GetObject(id)
	fmt.Println("Mostrando el resultado")
	fmt.Println("Offset:", offSet, "Object.Id:", object.Id_, "Object.Type:", object.Type_)

	if v.InFunction != "" {
		// Caso especial de identificador, posibilidad de uso en expresion anidada con profundidad
		if v.Depth >= 0 {
			v.C.COMENT("Identificador: " + id + " con profundidad: " + strconv.Itoa(v.Depth))
			v.LoadIdentifierAndSave(object)
			return nil
		}
		v.C.COMENT("Identificador: " + id)
		//ALTERACION CON IDENTIFIER!!!
		// MOVILIZAMOS EL VALOR AL REGISTRO SEGUN EL TIPO DE DATO!
		switch object.Type_ {
		case traductor.Int, traductor.Bool:
			// Los tipos enteros van dirigidos al registro x1
			v.C.LDR(registros.X1, registros.X29, -offSet*8)
		case traductor.StringType:
			// Los tipos string van dirigidos al registro x0
			v.C.LDR(registros.X0, registros.X29, -offSet*8)
			// Movemos el registro x0 al x11 para que sea manipulado
			v.C.MovReg(registros.X11, registros.X0)
		case traductor.Float:
			// Los tipos float van dirigidos al registro d0
			v.C.LDR(registros.D0, registros.X29, -offSet*8)
			v.C.Fmov(registros.D0, registros.X0)
		}

		// Copiamos el objeto para manipularlo
		CloneObject := v.C.CloneObject(object)
		CloneObject.Id_ = ""
		// Colocamos el objeto clonado en la pila para que sea manipulado
		v.C.PushObjectStack(CloneObject)
		return nil

	}

	v.C.Mov(registros.X0, offSet)
	v.C.Add(registros.X0, registros.SP, registros.X0)
	v.C.LDR(registros.X0, registros.X0, 0)

	CloneObject := v.C.CloneObject(object)
	v.C.PushObjectStack(CloneObject)
	return nil
}

// VisitVarDcl
func (v *CompileARMVisitor) VisitVarDeclStmt(ctx *gramAntlr.VarDeclStmtContext) interface{} {
	return v.Visit(ctx.VarDcl())
}

// 'mut' ID type '=' expr ';'
func (v *CompileARMVisitor) VisitVarDclWithTypeAndValue(ctx *gramAntlr.VarDclWithTypeAndValueContext) interface{} {
	fmt.Println(" ------------------------ VisitVarDclWithTypeAndValue ------------------------")
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()

	expr := ctx.Expr()
	v.Visit(expr)
	v.C.COMENT(fmt.Sprintf("Declaracion explicita: %s con tipo %s", id, typeStr))

	v.SaveFrameLocalObject(id)

	return nil
}

// 'mut' ID ':=' expr ';'
func (v *CompileARMVisitor) VisitVarDclWithInference(ctx *gramAntlr.VarDclWithInferenceContext) interface{} {
	id := ctx.ID_VARIABLE().GetText()
	expr := ctx.Expr()
	// Generamos el codigo para la expresion a asignar
	// Se asume que en ese visit expresion se hara un push del objeto
	v.Visit(expr)

	v.C.COMENT(fmt.Sprintf("Declaracion inferida: %s", id))

	// Guardamos el objeto anterior en la variable que corresponda a la posicion del frame pointer
	v.SaveFrameLocalObject(id)
	return nil
}

// 'mut' ID type ';'
func (v *CompileARMVisitor) VisitVarDclWithTypeOnly(ctx *gramAntlr.VarDclWithTypeOnlyContext) interface{} {
	fmt.Println(" ------------------------ VisitVarDclWithTypeOnly ------------------------")
	// Obtenemos el ID y el tipo de dato
	id := ctx.ID_VARIABLE().GetText()
	typeStr := ctx.Type_().GetText()

	// Dependiendo del tipo de dato, generamos el objeto correspondiente
	// y lo empujamos a la pila como constante para luego ser utilizado de inmediato.
	switch typeStr {
	case "int":
		v.C.COMENT("Valor por defecto para entero")
		var IntObject = v.C.IntObject()
		v.C.PushConst(IntObject, 0)
	case "float64":
		v.C.COMENT("Valor por defecto para float64")
		var FloatObject = v.C.FloatObject()
		v.C.PushConst(FloatObject, 0.0)
	case "string":
		v.C.COMENT("Valor por defecto para string")
		var StringObject = v.C.StrObject()
		v.C.PushConst(StringObject, "")
	case "bool":
		v.C.COMENT("Valor por defecto para booleano")
		var BoolObject = v.C.BoolObject()
		v.C.PushConst(BoolObject, 0)
	default:
		panic(fmt.Sprintf("Tipo no soportado: %s", typeStr))
	}

	// Generamos un comentario para debuguear en ARM
	v.C.COMENT(fmt.Sprintf("Declaracion no inicializada: %s con tipo %s", id, typeStr))

	// Guardamos el objeto anterior en la variable que corresponda a la posicion del frame pointer
	v.SaveFrameLocalObject(id)

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
	v.C.MovReg(registros.X11, registros.X10)

	if v.Depth >= 0 {
		strObject.Offset_ = v.PositionFramePointer
		v.C.PositionFramePointer = v.PositionFramePointer // Guardamos la posicion del frame pointer
		v.C.PushConst(strObject, texto)                   // Guardamos el string en el stack frame
		v.C.Str(registros.X11)                            // Guardamos el string en el stack frame
		fmt.Println("String pusheado -> offset:", strObject.Offset_, "Profundidad:", v.Depth, "valor:", texto)
		v.PositionFramePointer += 1 // Aumentamos el offset del frame pointer
		return nil
	}

	// Manipulamos el string para unificarlo
	v.C.PushConst(strObject, texto)
	// Actualizamos la posicion del frame pointer 1 byte mas
	//v.PositionFramePointer += 1
	return nil
}

func (v *CompileARMVisitor) VisitInteger(ctx *gramAntlr.IntegerContext) interface{} {
	value, err := strconv.Atoi(ctx.GetText())

	if err != nil {
		panic(fmt.Sprintf("Error al convertir a entero: %s", ctx.GetText()))
	}

	v.C.COMENT(fmt.Sprintf("Entero: %d", value))

	// Indispensable para reconocimiento de tipo entero en print o cualquier otra operacion!
	IntObject := v.C.IntObject()

	// Si el entero esta en una profundidad anidada, lo manejamos de forma especial
	// Esto es para manejar el caso de anidamiento de expresiones
	// y que el valor se guarde en la pila correctamente.
	if v.Depth >= 0 {
		// Cambio el offest del objeto para saber donde se guarda y obtenerlo
		IntObject.Offset_ = v.PositionFramePointer
		// Guardo el entero en stack frame
		v.C.PositionFramePointer = v.PositionFramePointer
		v.C.PushConst(IntObject, value)
		v.C.Str(registros.X1)
		fmt.Println("Integer pusheado -> offset:", IntObject.Offset_, "Profundidad:", v.Depth, "valor:", value)
		v.PositionFramePointer += 1 // Aumentamos el offset del frame pointer
		return nil
	}
	// no hubo necesidad de anidamiento, por lo que simplemente guardamos el objeto
	v.C.PushConst(IntObject, value)
	fmt.Println("Integer pusheado -> offset:", IntObject.Offset_, "Profundidad:", v.Depth, "valor:", value)

	return nil
}

func (v *CompileARMVisitor) VisitDouble(ctx *gramAntlr.DoubleContext) interface{} {
	value, err := strconv.ParseFloat(ctx.GetText(), 64)
	//Verificar que se pueda convertir el texto a float64
	if err != nil {
		panic(fmt.Sprintf("Error al convertir a float: %s", ctx.GetText()))
	}
	// Comentario para debuguar el valor del float
	v.C.COMENT(fmt.Sprintf("Float: %f", value))

	//Guardamos el objeto float64 en la pila para su proxima recuperacion
	floatObject := v.C.FloatObject()

	// verificamos profundidad de anidamiento
	if v.Depth >= 0 {
		floatObject.Offset_ = v.PositionFramePointer
		v.C.PushConst(floatObject, value)
		v.C.PositionFramePointer = v.PositionFramePointer // Guardamos la posicion del frame pointer
		v.C.Str(registros.D0)                             // Guardamos el valor en el stack frame
		fmt.Println("Float pusheado -> offset:", floatObject.Offset_, "Profundidad:", v.Depth, "valor:", value)
		v.PositionFramePointer += 1 // Aumentamos el offset del frame pointer
		return nil
	}

	// Realizamos el push del objeto float64 con su valor
	v.C.PushConst(floatObject, value)
	fmt.Println("Float pusheado -> offset:", floatObject.Offset_, "Profundidad:", v.Depth, "valor:", value)

	return nil
}

func (v *CompileARMVisitor) VisitBoolean(ctx *gramAntlr.BooleanContext) interface{} {

	value := 0
	valueStr := strings.ToLower(ctx.GetText())
	// Pasamos el dato en formato String a un entero 0 false | 1 true segun corresponda
	if valueStr == "true" {
		value = 1
	} else {
		value = 0
	}
	// Agregamos un comentario para el booleano y facilidad de debug
	// Recordemos que el booleano se maneja como un entero en ARM
	v.C.COMENT(fmt.Sprintf("Booleano: %d", value))

	// Creamos el objeto booleano y lo empujamos a la pila para su proxima recuperacion
	boolObject := v.C.BoolObject()

	if v.Depth >= 0 {
		boolObject.Offset_ = v.PositionFramePointer
		v.C.PositionFramePointer = v.PositionFramePointer // Guardamos la posicion del frame pointer
		v.C.PushConst(boolObject, value)
		v.C.Str(registros.X1)
		fmt.Println("Boolean pusheado -> offset:", boolObject.Offset_, "Profundidad:", v.Depth, "valor:", value)
		v.PositionFramePointer += 1 // Aumentamos el offset del frame pointer
		return nil
	}
	v.C.PushConst(boolObject, value)
	fmt.Println("Boolean pusheado -> offset:", boolObject.Offset_, "Profundidad:", v.Depth, "valor:", value)
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
	//var reg string

	// Elegir el registro adecuado
	if top.Type_ == traductor.Float {
		//reg = registros.D0
	} else {
		//reg = registros.X0
	}

	// Hacer POP con el registro correcto
	value := v.C.POPOBJECT()

	// Aplicar la negación según el tipo
	if value.Type_ == traductor.Int {
		v.C.Neg(registros.X0, registros.X0)

		v.C.PushObjectStack(v.C.CloneObject(value))
	} else if value.Type_ == traductor.Float {
		v.C.COMENT("Float64 Negativo")
		v.C.FNeg(registros.D0, registros.D0)

		v.C.PushObjectStack(v.C.CloneObject(value))
	}

	return nil
}

func (v *CompileARMVisitor) VisitAddSub(ctx *gramAntlr.AddSubContext) interface{} {
	v.Depth += 1
	// Visitamos las expresiones del lado izquierdo y derecho para preparar la pila
	v.C.COMENT(" ------------- VisitAddSub -------------")
	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))
	// Vemos estado de la pila
	// Obtenemos el operador que puede ser '+' o '-'
	operador := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	/*
		Hacemos un pop del objeto y lo asignamos a Derecha.
		Sabemos que lo ultimo en la pila es la expresion del lado derecho
		porque el visit de la expresion del lado derecho se ejecuta despues del izquierdo.
		Por lo tanto, el ultimo objeto en la pila es el de la derecha.
	*/

	//Obtenemos lo que esta en la cima de la pila y validamos su tipo
	DerechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	Derecha := v.C.POPOBJECT()

	IzquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	Izquierda := v.C.POPOBJECT()

	switch operador {
	case "+":
		// int + int
		if Derecha.Type_ == traductor.Int && Izquierda.Type_ == traductor.Int {
			v.C.COMENT("SUMA DE ENTEROS")
			v.LoadIntObjet(Izquierda, Derecha)
			v.C.Add(registros.X0, registros.X2, registros.X1)
			// Hacer push de la suma realizada a la pila
			v.SaveResult(registros.X0, traductor.Int)

		} else if Derecha.Type_ == traductor.StringType && Izquierda.Type_ == traductor.StringType {
			v.C.COMENT("CONCATENACION DE CADENAS")
			// string + string
			// Cargamos el lado izquierdo y derecho
			v.C.LDR(registros.X0, registros.X29, -Izquierda.Offset_*8)
			v.C.LDR(registros.X1, registros.X29, -Derecha.Offset_*8)
			// Cargamos el buffer para la concatenacion
			v.C.Adr(registros.X10, "buffer")
			v.C.ConcatString()

			v.SaveResult(registros.X0, traductor.StringType)
			// Reiniciamos el buffer para uso futuro
			v.C.Adr(registros.X10, "buffer")
		} else if Derecha.Type_ == traductor.Float || Izquierda.Type_ == traductor.Float {
			v.C.COMENT("SUMA DE FLOAT64")
			// float64 + ...
			v.ChangeTypeToFloat(Izquierda.Offset_, IzquierdaEsFloat, registros.X0, registros.D0)
			v.ChangeTypeToFloat(Derecha.Offset_, DerechaEsFloat, registros.X1, registros.D1)
			v.C.Fadd(registros.D0, registros.D0, registros.D1)
			v.SaveResult(registros.D0, traductor.Float)
		}
	case "-":
		if Izquierda.Type_ == traductor.Int && Derecha.Type_ == traductor.Int {
			v.C.COMENT("RESTA DE ENTEROS")
			v.LoadIntObjet(Izquierda, Derecha)
			v.C.Sub(registros.X0, registros.X2, registros.X1)
			v.SaveResult(registros.X0, traductor.Int)

		} else if Izquierda.Type_ == traductor.Float || Derecha.Type_ == traductor.Float {
			v.C.COMENT("RESTA DE FLOAT64")
			v.ChangeTypeToFloat(Izquierda.Offset_, IzquierdaEsFloat, registros.X0, registros.D0)
			v.ChangeTypeToFloat(Derecha.Offset_, DerechaEsFloat, registros.X1, registros.D1)
			v.C.Fsub(registros.D0, registros.D0, registros.D1)
			v.SaveResult(registros.D0, traductor.Float)
		}
	}
	v.Depth -= 1
	return nil
}

func (v *CompileARMVisitor) VisitMulDivModulo(ctx *gramAntlr.MulDivModuloContext) interface{} {
	v.Depth += 1
	v.C.COMENT(" ------------- VisitMulDivModulo -------------")
	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))

	operador := ctx.GetOp().GetText()

	derechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	derecha := v.C.POPOBJECT()

	izquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	izquierda := v.C.POPOBJECT()

	switch operador {
	case "*":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.COMENT("MULTIPLICACION DE ENTEROS")
			v.LoadIntObjet(izquierda, derecha)
			v.C.Mul(registros.X0, registros.X2, registros.X1)
			v.SaveResult(registros.X0, traductor.Int)

		} else if derecha.Type_ == traductor.Float || izquierda.Type_ == traductor.Float {
			v.C.COMENT("MULTIPLICACION DE FLOAT64")
			v.ChangeTypeToFloat(izquierda.Offset_, izquierdaEsFloat, registros.X0, registros.D0)
			v.ChangeTypeToFloat(derecha.Offset_, derechaEsFloat, registros.X1, registros.D1)
			v.C.FMul(registros.D0, registros.D0, registros.D1)
			v.SaveResult(registros.D0, traductor.Float)
		}

	case "/":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.COMENT("DIVISION DE ENTEROS")
			v.LoadIntObjet(izquierda, derecha)
			v.C.Div(registros.X0, registros.X2, registros.X1)
			v.SaveResult(registros.X0, traductor.Int)

		} else if derecha.Type_ == traductor.Float || izquierda.Type_ == traductor.Float {
			v.C.COMENT("DIVISION DE FLOAT64")
			v.ChangeTypeToFloat(izquierda.Offset_, izquierdaEsFloat, registros.X0, registros.D0)
			v.ChangeTypeToFloat(derecha.Offset_, derechaEsFloat, registros.X1, registros.D1)
			v.C.FDiv(registros.D0, registros.D0, registros.D1)
			v.SaveResult(registros.D0, traductor.Float)

		}

	case "%":
		if derecha.Type_ == traductor.Int && izquierda.Type_ == traductor.Int {
			v.C.COMENT("MODULO DE ENTEROS")
			v.LoadIntObjet(izquierda, derecha)
			v.C.Mod(registros.X0, registros.X2, registros.X1)
			v.SaveResult(registros.X0, traductor.Int)
		}
	}
	v.Depth -= 1
	return nil
}

// --------------------------------- OPERADORES DE COMPARACION ---------------------------------
// VisitMinorMajorEqual
func (v *CompileARMVisitor) VisitMinorMajorEqual(ctx *gramAntlr.MinorMajorEqualContext) interface{} {
	v.C.COMENT(" ------------- Operación Relacional -------------")
	v.Depth += 1
	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))
	op := ctx.GetOp().GetText()

	derechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	derecha := v.C.POPOBJECT()

	izquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	izquierda := v.C.POPOBJECT()

	if izquierdaEsFloat || derechaEsFloat {
		v.C.COMENT("Comparación de Float64")
		v.ChangeTypeToFloat(izquierda.Offset_, izquierdaEsFloat, registros.X0, registros.D0)
		v.ChangeTypeToFloat(derecha.Offset_, derechaEsFloat, registros.X1, registros.D1)
		v.C.Fcmp(registros.D0, registros.D1)
	} else {
		v.C.COMENT("Comparación de Enteros")
		v.LoadIntObjet(izquierda, derecha)
		v.C.Cmp(registros.X2, registros.X1)
	}
	etiquetaFalse := "Lfalse" + strconv.Itoa(v.LabelCompareCount)
	etiquetaTrue := "Ltrue" + strconv.Itoa(v.LabelCompareCount)

	switch op {
	case "<":
		v.C.BCond("LT", etiquetaTrue)
	case ">":
		v.C.BCond("GT", etiquetaTrue)
	case "<=":
		v.C.BCond("LE", etiquetaTrue)
	case ">=":
		v.C.BCond("GE", etiquetaTrue)
	}

	v.C.Mov(registros.X0, 0)
	v.C.B(etiquetaFalse)
	v.C.SetLabel(etiquetaTrue)
	v.C.Mov(registros.X0, 1)
	v.C.SetLabel(etiquetaFalse)
	v.SaveResult(registros.X0, traductor.Bool)
	// Incrementamos el contador de etiquetas de comparación
	v.LabelCompareCount++
	v.Depth -= 1
	return nil
}

// expr: expr op = ('==' | '!=') expr                          # EqualsNotEquals
func (v *CompileARMVisitor) VisitEqualsNotEquals(ctx *gramAntlr.EqualsNotEqualsContext) interface{} {
	fmt.Println("----------------------------VisitEqualsNotEquals")
	v.C.COMENT(" ------------- Operación de Igualdad -------------")
	v.Depth += 1
	v.Visit(ctx.Expr(0))
	v.Visit(ctx.Expr(1))
	op := ctx.GetOp().GetText()

	derechaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	derecha := v.C.POPOBJECT()
	fmt.Println("----------------------------VisitEqualsNotEquals pop der")
	izquierdaEsFloat := v.C.GetTopObjectStack().Type_ == traductor.Float
	izquierda := v.C.POPOBJECT()
	fmt.Println("----------------------------VisitEqualsNotEquals por izq")
	if izquierdaEsFloat || derechaEsFloat {
		v.C.COMENT("Comparación de Float64")
		v.ChangeTypeToFloat(izquierda.Offset_, izquierdaEsFloat, registros.X0, registros.D0)
		v.ChangeTypeToFloat(derecha.Offset_, derechaEsFloat, registros.X1, registros.D1)
		v.C.Fcmp(registros.D0, registros.D1)
	} else {
		v.C.COMENT("Comparación de Enteros")
		v.LoadIntObjet(izquierda, derecha)
		v.C.Cmp(registros.X2, registros.X1)
	}

	etiquetaFalse := "Lfalse" + strconv.Itoa(v.LabelCompareCount)
	etiquetaTrue := "Ltrue" + strconv.Itoa(v.LabelCompareCount)

	switch op {
	case "==":
		v.C.BCond("EQ", etiquetaTrue)
	case "!=":
		v.C.BCond("NE", etiquetaTrue)
	}

	v.C.Mov(registros.X0, 0)
	v.C.B(etiquetaFalse)
	v.C.SetLabel(etiquetaTrue)
	v.C.Mov(registros.X0, 1)
	v.C.SetLabel(etiquetaFalse)
	v.SaveResult(registros.X0, traductor.Bool)
	// Incrementamos el contador de etiquetas de comparación
	v.LabelCompareCount++
	v.Depth -= 1
	return nil
}

// expr    | '!' expr                                              # Not
func (v *CompileARMVisitor) VisitNot(ctx *gramAntlr.NotContext) interface{} {
	v.Depth += 1
	v.C.COMENT(" ------------- Operación NOT -------------")
	v.Visit(ctx.Expr()) // Evaluar la expresión interna

	// Obtener el valor de la cima de la pila
	value := v.C.POPOBJECT()

	// Verificar tipo del valor: debe ser booleano
	if value.Type_ != traductor.Bool {
		panic("Operador '!' solo puede aplicarse a booleanos")
	}
	// Cargar el valor en el registro X0
	v.C.LDR(registros.X0, registros.X29, -value.Offset_*8)

	// Aplicar NOT lógico: XOR con 1 (0 -> 1, 1 -> 0)
	v.C.Eor(registros.X0, registros.X0, "#1")

	v.SaveResult(registros.X0, traductor.Bool)
	v.Depth -= 1
	return nil
}

// IMPRESIONES

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
		// Extraemos el primer elemento de la pila
		miExpr := v.C.GetTopObjectStack()
		isFloat := miExpr.Type_ == traductor.Float

		var value traductor.ObjectStack
		if isFloat {
			// CAPTURAMOS EL VALOR FLOAT
			value = v.C.POPOBJECT()
		} else {
			value = v.C.POPOBJECT()
		}

		switch value.Type_ {
		case traductor.Int:
			if miExpr.Offset_ != 0 {
				v.C.COMENT("Preparando entero para imprimir")
				v.C.LDR(registros.X1, registros.X29, -miExpr.Offset_*8)
			}
			// Mandamos el registro x1 pues anteriormente hicimos un mov x1 con el valor entero
			v.C.ImprimirEntero(registros.X1)

		case traductor.Bool:
			if miExpr.Offset_ != 0 {
				v.C.LDR(registros.X1, registros.X29, -miExpr.Offset_*8)
			}
			v.C.ImprimirBooleano(registros.X1)
			v.C.ImprimirCadena(registros.X1)

		case traductor.Float:
			if miExpr.Offset_ != 0 {
				v.C.COMENT("Preparando float para imprimir")
				v.C.LDR(registros.D0, registros.X29, -miExpr.Offset_*8)
			}
			v.C.ImprimirDecimal()

		case traductor.StringType:
			if miExpr.Offset_ != 0 {
				v.C.COMENT("Preparando string para imprimir")
				v.C.LDR(registros.X11, registros.X29, -miExpr.Offset_*8)
			}
			v.C.MovReg(registros.X0, registros.X11)
			v.C.ImprimirCadena(registros.X0)

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
	// Inicializamos el PositionFramePointer con baseOffSet
	v.PositionFramePointer = baseOffSet

	// Creamos el visitor para manejar los fragmentos de código y el manejo de offsets (desplazamientos)
	fragmentVisitor := fragmentvisitor.NewFragmentVisitor(baseOffSet + paramsOffSet)

	for _, child := range ctx.Block().GetChildren() {
		// Debemos verificar que los hijos sean del tipo ParseTree
		// y luego visitar cada uno de ellos con el fragmentVisitor
		if node, ok := child.(antlr.ParseTree); ok {
			fmt.Println("Visitando nodo:", node.GetText())
			fragmentVisitor.Visit(node)
			fmt.Println("LocalOffset:", fragmentVisitor.LocalOffSet)
		}
	}

	fmt.Println("")

	fragment := fragmentVisitor.Fragment

	localOffSet := fragmentVisitor.LocalOffSet
	returnOffSet := 1
	nameFunction := ctx.ID_VARIABLE(0).GetText()
	fmt.Println("SizeFragment: baseOffSet", baseOffSet, "paramsOffSet", paramsOffSet, "localOffSet", localOffSet, "returnOffSet", returnOffSet)
	sizeFragment := baseOffSet + paramsOffSet + localOffSet + returnOffSet
	fmt.Println("SizeFragment (bytes):", sizeFragment)
	fmt.Println("")

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

	fmt.Println("Generando objetos de fragmento: ", len(fragment))
	for _, instruction := range fragment {
		fmt.Println("Pushing fragment object:", instruction.Name, "Offset:", instruction.Offset)
		v.C.PushObjectStack(traductor.ObjectStack{
			Type_:   traductor.Void,
			Id_:     instruction.Name,
			Offset_: instruction.Offset,
			Length_: 8,
		})
	}
	fmt.Println("")

	v.InFunction = nameFunction
	v.FragmentPointerOffSet = 0
	v.ReturnLabels = v.C.GetLabel()
	v.C.SetLabel(nameFunction)

	if nameFunction == "main" {
		// Agregamos los punteros de x29 y x30 al stack
		v.C.Instrucciones = append(v.C.Instrucciones, "stp x29, x30, [sp, #-16]!")
		// Establecemos el frame pointer
		v.C.Instrucciones = append(v.C.Instrucciones, "mov x29, sp")
		//Reservamos el espacio necesario para operar en x29, recordemos pasar de bytes a bits
		v.C.SetAsideSizeFrameFunction(sizeFragment * 8)
	}
	//Recorremos el bloque de instrucciones
	fmt.Println("Instrucciones dentro del bloque")
	for _, child := range ctx.Block().GetChildren() {
		// Debemos verificar que los hijos sean del tipo ParseTree
		// y luego visitar cada uno de ellos con el fragmentVisitor
		if node, ok := child.(antlr.ParseTree); ok {
			fmt.Println()
			fmt.Println("Visitando nodo en bloque:", node.GetText())
			v.Visit(node)
			fmt.Println("Posicion del Frame Pointer:", v.PositionFramePointer)
		}
	}
	fmt.Println("")
	v.C.SetLabel(v.ReturnLabels)
	if nameFunction != "main" {
		v.C.Add(registros.X0, registros.FP, registros.XZR)
		v.C.LDR(registros.LR, registros.X0, 0)
		v.C.Br(registros.LR)
	} else {
		// Regresamos al stackPointer lo que reservamos para la funcion
		v.C.PopInStackPointer(sizeFragment * 8)
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

// --------------------- FUNCIONES AUXILIARES ---------------------
// Funcion para obtener pasar el tipo de string a TypeObject
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

// Función equivalente para los 3 tipos de declaración de variables
func (v *CompileARMVisitor) SaveFrameLocalObject(id string) {
	/*
		Para usar correctamente esta función, se asume que anteriormente guardaste en la pila el objeto
		el cual aqui se asignara a una variable con la que hacemos referencia mediante PositionFramePointer.
	*/

	// Si estamos dentro de una funcion, guardamos el objeto en el frame local
	if v.InFunction != "" {

		LocalObjecto := v.C.GetFrameLocal(v.FragmentPointerOffSet)
		// Evaluamos el ultimo objeto en la pila
		ValorObjecto := v.C.POPOBJECT()

		// Movemos el valor anterior al registro x0 que es mi variable que se guardara en frame pointer
		// Validamos que tipo de dato  es para saber que registro usar
		switch ValorObjecto.Type_ {
		case traductor.StringType:
			// Caso especial, el objeto valorObject pudo haber sino una operacion anterior
			if ValorObjecto.Offset_ != 0 {
				v.C.LDR(registros.X11, registros.X29, -ValorObjecto.Offset_*8)
			}
			v.C.MovReg(registros.X0, registros.X11) // x11 es el registro para strings y el inicio de la cadena
			// Reservamos el espacio en el frame pointer
			v.C.PositionFramePointer = v.PositionFramePointer
			// Restamos el frame pointer al offset del frame pointer
			v.C.Str(registros.X0)
		case traductor.Int, traductor.Bool:
			// Caso especial, el objeto valorObject pudo haber sino una operacion anterior
			if ValorObjecto.Offset_ != 0 {
				v.C.LDR(registros.X1, registros.X29, -ValorObjecto.Offset_*8)
			}
			// Movemos el registro x1 cone el valor al registro x0
			v.C.MovReg(registros.X0, registros.X1) // x1 es el registro para enteros
			// Reservamos el espacio en el frame pointer
			v.C.PositionFramePointer = v.PositionFramePointer
			// Restamos el frame pointer al offset del frame pointer
			v.C.Str(registros.X0)
		case traductor.Float:
			// Caso especial, el objeto valorObject pudo haber sino una operacion anterior
			if ValorObjecto.Offset_ != 0 {
				v.C.LDR(registros.D0, registros.X29, -ValorObjecto.Offset_*8)
			}
			// Se guarda el valor inmediato en D0
			v.C.PositionFramePointer = v.PositionFramePointer
			v.C.Str(registros.D0)
		}
		// Cambiamos el tipo del objeto local al tipo del valor del objeto
		LocalObjecto.Type_ = ValorObjecto.Type_
		// Aumentamos el offset del frame pointer
		v.FragmentPointerOffSet++
		// Aumentamos el frame pointer para la siguiente variable
		v.PositionFramePointer += 1
		return
	}

	v.C.TagObjecto(id)
}

func (v *CompileARMVisitor) LoadIdentifierAndSave(object traductor.ObjectStack) {

	switch object.Type_ {
	case traductor.Int, traductor.StringType, traductor.Bool:
		// Carga en x0 el valor del objeto desde el stack frame
		v.C.LDR(registros.X0, registros.X29, -object.Offset_*8)

	case traductor.Float:
		// Carga en d0 el valor float del objeto desde el stack frame
		v.C.LDR(registros.D0, registros.X29, -object.Offset_*8)

	default:
		panic("Tipo no soportado en LoadIdentifier")
	}
	// Clonamos el objeto para manipularlo
	objectClone := v.C.CloneObject(object)
	// Le cambiamos su offset al del frame pointer
	objectClone.Offset_ = v.PositionFramePointer
	// Lo empujamos a la pila para que sea manipulado
	v.C.PushObjectStack(objectClone)

	// Guardamos en stack frame según tipo
	v.C.PositionFramePointer = v.PositionFramePointer
	switch object.Type_ {
	case traductor.Int, traductor.StringType, traductor.Bool:
		v.C.Str(registros.X0)
	case traductor.Float:
		v.C.Str(registros.D0)
	}
	v.PositionFramePointer += 1 // Aumentamos el offset del frame pointer
}

// --------------------- FUNCIONES AUXILIARES EXPRESIONES ---------------------
func (v *CompileARMVisitor) LoadIntObjet(Izquierda traductor.ObjectStack, Derecha traductor.ObjectStack) {
	// Cargamos los valores de izquierda y derecha al registro x0 y x1
	v.C.LDR(registros.X2, registros.X29, -Izquierda.Offset_*8)
	v.C.LDR(registros.X1, registros.X29, -Derecha.Offset_*8)
}

func (v *CompileARMVisitor) ChangeTypeToFloat(offset int, esFloat bool, regInt, regFloat string) {
	if !esFloat {
		v.C.LDR(regInt, registros.X29, -offset*8)
		v.C.Scvtf(regFloat, regInt)
	} else {
		v.C.LDR(regFloat, registros.X29, -offset*8)
	}
}

func (v *CompileARMVisitor) SaveResult(registro string, tipo traductor.TypeObject) {
	var result traductor.ObjectStack
	switch tipo {
	case traductor.Int:
		result = v.C.IntObject()
	case traductor.Float:
		result = v.C.FloatObject()
	case traductor.StringType:
		result = v.C.StrObject()
	case traductor.Bool:
		result = v.C.BoolObject()
	default:
		panic("Tipo desconocido")
	}
	result.Offset_ = v.PositionFramePointer
	v.C.PushObjectStack(result)
	v.C.PositionFramePointer = v.PositionFramePointer // Guardamos la posicion del frame pointer
	v.C.Str(registro)
	v.PositionFramePointer += 1
}

// --------------------- DEPURACION DE OFFSETS BASURA ---------------------
func CleanUpStack(usedPositions []int) []int {
	return nil
}

// --------------------- FUNCIONES AUXILIARES PARA POSICION MAYOR A 255 ---------------------
func (v *CompileARMVisitor) StrOffset(reg string, offset int) {
	v.C.Mov("x9", offset)
	v.C.Add("x9", "x29", "x9")
	//v.C.Str(reg, "[x9]")
}
