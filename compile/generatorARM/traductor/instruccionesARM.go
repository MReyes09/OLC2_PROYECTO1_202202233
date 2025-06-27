package traductor

import (
	"OLC2CLIENTE/compile/generatorARM/funciones"
	primitvos "OLC2CLIENTE/compile/generatorARM/primitivos"
	"OLC2CLIENTE/compile/generatorARM/registros"
	"fmt"
	"math"
	"strings"
)

type TypeObject int

const (
	Int TypeObject = iota
	Float
	StringType
	Bool
	Rune
	Void
	Slice
)

type ObjectStack struct {
	Type_         TypeObject
	Length_       int
	Depth_        int
	Id_           string
	TipoElemento_ TypeObject
	Offset_       int
}

type GeneratorARMInstructions struct {
	Instrucciones          []string
	FuncionesInstrucciones []string
	Estandar               *funciones.EstandarFunc
	Stack                  []ObjectStack
	PositionFramePointer   int
	depth                  int
	contadorEtiquetas      int
}

func NewGeneratorARMInstructions() *GeneratorARMInstructions {
	return &GeneratorARMInstructions{
		Instrucciones:          []string{},
		FuncionesInstrucciones: []string{},
		Estandar:               funciones.NewEstandarFunc(),
		Stack:                  []ObjectStack{},
		PositionFramePointer:   0,
		depth:                  0,
		contadorEtiquetas:      0,
	}
}

// ------------------------ INSTRUCCIONES ARM ------------------------
func (g *GeneratorARMInstructions) Add(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) Sub(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) Mul(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MUL %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) FMul(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FMUL %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) Div(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("SDIV %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) FDiv(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FDIV %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) Mod(rd string, dividend string, divisor string) {
	temp := registros.X3 // O cualquier otro temporal no usado
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("SDIV %s, %s, %s", temp, dividend, divisor))
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MSUB %s, %s, %s, %s", rd, temp, divisor, dividend))
}

// Nueva implemetacion para LDR

func (g *GeneratorARMInstructions) LDR(rd string, base string, offset int) {
	if offset >= -255 && offset%8 == 0 {
		g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("LDR %s, [%s, #%d]", rd, base, offset))
	} else {
		// Versión usando registro temporal para offsets grandes
		g.LdrByTemp(rd, offset)
	}
}

func (g *GeneratorARMInstructions) LdrByTemp(rd string, offset int) {
	g.Mov("x9", offset)
	g.Add("x9", "x29", "x9")
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("LDR %s, [x9]", rd))
}

func (g *GeneratorARMInstructions) Adr(rd string, label string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("ADR %s, %s", rd, label))
}

func (g *GeneratorARMInstructions) Br(destino string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("BR %s", destino))
}

func (g *GeneratorARMInstructions) Mov(rd string, valueDirect int) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOV %s, #%d", rd, valueDirect))
}

func (g *GeneratorARMInstructions) MovReg(rd string, rs string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOV %s, %s", rd, rs))
}

func (g *GeneratorARMInstructions) Push(rs string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STR %s, [SP, #-8]!", rs))
}

func (g *GeneratorARMInstructions) PushInFramePointer(rs string, PositionFramePointer int) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STR %s, [x29, #-%d]", rs, PositionFramePointer))
}

func (g *GeneratorARMInstructions) Pop(rd string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("LDR %s, [SP], #8", rd))
}

func (g *GeneratorARMInstructions) PopInStackPointer(PositionFramePointer int) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("ADD sp, sp, #%d", PositionFramePointer))
}

func (g *GeneratorARMInstructions) Srtb(rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STRB %s, [%s]", rs1, rs2))
}

// Implementacion nueva de STR
func (g *GeneratorARMInstructions) Str(rs string) {
	offset := g.PositionFramePointer * 8
	if offset <= 255 {
		g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STR %s, [x29, #-%d]", rs, offset))
	} else {
		g.StrByTemp(rs, offset)
	}
}

// Guardando registros offset > 255
func (g *GeneratorARMInstructions) StrByTemp(rs string, offset int) {
	g.Mov("x9", -offset)                                                       // offset negativo (hacia abajo en el stack)
	g.Add("x9", "x29", "x9")                                                   // dirección efectiva = x29 - offset
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STR %s, [x9]", rs)) // guardar valor
}

func (g *GeneratorARMInstructions) StrFromReg(rs string, rd string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("STR %s, [%s]", rs, rd))
}

func (g *GeneratorARMInstructions) Fmov(rs string, rd string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FMOV %s, %s", rd, rs))
}

func (g *GeneratorARMInstructions) Svc() {
	g.Instrucciones = append(g.Instrucciones, "SVC #0")
}

func (g *GeneratorARMInstructions) Scvtf(rd string, rs string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("SCVTF %s, %s", rd, rs))
}

func (g *GeneratorARMInstructions) Fadd(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FADD %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) Fsub(rd string, rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FSUB %s, %s, %s", rd, rs1, rs2))
}

func (g *GeneratorARMInstructions) BCond(condition string, label string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("B.%s %s", condition, label))
}

func (g *GeneratorARMInstructions) Cmp(rs1 string, rs2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("CMP %s, %s", rs1, rs2))
}

func (g *GeneratorARMInstructions) B(label string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("B %s", label))
}

func (g *GeneratorARMInstructions) Eor(rd string, rn string, operand2 string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("EOR %s, %s, %s", rd, rn, operand2))
}

func (g *GeneratorARMInstructions) ImprimirCadena(rs string) {
	g.Estandar.Usar("print_cadena")
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOV X0, %s", rs))
	g.Instrucciones = append(g.Instrucciones, "BL print_cadena")
}

func (g *GeneratorARMInstructions) ImprimirEntero(rs string) {
	// Agregamos a la lista de llamadas a la función estándar de impresión de enteros
	g.Estandar.Usar("print_entero")

	// Preparamos el entero en el registro x0 para su impresion con print_entero
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOV X0, %s", rs))
	g.Instrucciones = append(g.Instrucciones, "BL print_entero")
}

func (g *GeneratorARMInstructions) ImprimirDecimal() {
	g.Estandar.Usar("print_entero")
	g.Estandar.Usar("print_decimal")
	g.Instrucciones = append(g.Instrucciones, "BL print_decimal")
}

func (g *GeneratorARMInstructions) ImprimirBooleano(rs string) {
	g.Estandar.Usar("print_booleano")
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOV X0, %s", rs))
	g.Instrucciones = append(g.Instrucciones, "BL print_booleano")
}

func (g *GeneratorARMInstructions) ConcatString() {
	g.Estandar.Usar("concatenar_cadena")
	g.Instrucciones = append(g.Instrucciones, "BL concatenar_cadena")
}

func (g *GeneratorARMInstructions) SaltoLinea() {
	g.Instrucciones = append(g.Instrucciones, "MOV X0, #1")
	g.Instrucciones = append(g.Instrucciones, "ADR X1, salto_linea_str")
	g.Instrucciones = append(g.Instrucciones, "MOV X2, #1")
	g.Instrucciones = append(g.Instrucciones, "MOV W8, #64")
	g.Instrucciones = append(g.Instrucciones, "SVC #0")
}

func (g *GeneratorARMInstructions) Espaciado() {
	g.Instrucciones = append(g.Instrucciones, "MOV X0, #1")
	g.Instrucciones = append(g.Instrucciones, "ADR X1, espacio_str")
	g.Instrucciones = append(g.Instrucciones, "MOV X2, #1")
	g.Instrucciones = append(g.Instrucciones, "MOV W8, #64")
	g.Instrucciones = append(g.Instrucciones, "SVC #0")
}

func (g *GeneratorARMInstructions) SetAsideSizeFrameFunction(size int) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("SUB SP, SP, #%d", size))
}

func (g *GeneratorARMInstructions) Neg(rd string, rs string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("NEG %s, %s", rd, rs))
}

func (g *GeneratorARMInstructions) FNeg(rd string, rs string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FNEG %s, %s", rd, rs))
}

func (g *GeneratorARMInstructions) Fcmp(rs string, rd string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("FCMP %s, %s", rs, rd))
}

// ------------------------ FUNCIONES DE TERMINACION ------------------------

func (g *GeneratorARMInstructions) EndProgram() {
	g.Mov(registros.X0, 0)
	g.Mov(registros.X8, 93)
	g.Svc()
}

// ------------------------ FUNCIONES EXTRA ------------------------

func (g *GeneratorARMInstructions) GetFrameLocal(index int) *ObjectStack {

	if index < 0 || index >= len(g.Stack) {
		panic(fmt.Sprintf("Índice %d fuera de rango para objetos de tipo Void", index))
	}

	return &g.Stack[index]

}

func (g *GeneratorARMInstructions) PushObjectStack(objectStack ObjectStack) {
	g.Stack = append(g.Stack, objectStack)
}

func (g *GeneratorARMInstructions) GetLabel() string {
	g.contadorEtiquetas++
	return fmt.Sprintf("L%d", g.contadorEtiquetas)
}

func (g *GeneratorARMInstructions) SetLabel(label string) {
	g.Instrucciones = append(g.Instrucciones, label+":")
}

func (g *GeneratorARMInstructions) COMENT(comentario string) {
	g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("// %s", comentario))
}

func (g *GeneratorARMInstructions) POPOBJECT() ObjectStack {
	object := g.Stack[len(g.Stack)-1]
	g.POPOBJECT2()
	return object

}

func (g *GeneratorARMInstructions) POPOBJECT2() {
	g.Stack = g.Stack[:len(g.Stack)-1]
}

func (g *GeneratorARMInstructions) PushConst(object ObjectStack, valor interface{}) {
	switch object.Type_ {
	case Int, Bool, Rune:
		g.Mov(registros.X1, valor.(int))
		//g.Push(registros.X0)

	case Float:
		// Convertimos el float64 a sus bits de representación
		floatBits := math.Float64bits(valor.(float64))

		// Extraemos los 4 bloques de 16 bits
		floatParts := make([]uint16, 4)

		// Descomponemos el floatBits en 4 partes de 16 bits
		for i := 0; i < 4; i++ {
			floatParts[i] = uint16((floatBits >> (i * 16)) & 0xFFFF)
		}

		// Generamos las instrucciones ARM para cargar el float en X0
		g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOVZ X0, #%d, LSL #0", floatParts[0]))

		// Usamos MOVK para cargar los siguientes 3 bloques de 16 bits
		for i := 1; i < 4; i++ {
			g.Instrucciones = append(g.Instrucciones, fmt.Sprintf("MOVK X0, #%d, LSL #%d", floatParts[i], i*16))
		}
		//Mover bits a registro D0
		g.Fmov(registros.X0, registros.D0)

	case StringType:
		// Agregamos STR registro, [x29, #-posicionFramePointer]   // guardar dirección del string en la posicion correcta
		//g.PushInFramePointer(registros.X0, g.PositionFramePointer*8)
		cadena := primitvos.StringToByte(valor.(string))

		for _, charCode := range cadena {
			g.COMENT(fmt.Sprintf("Byte: %d uso de Heap: %q", charCode, charCode))
			g.Mov("w0", int(charCode))
			g.Srtb("w0", registros.HP)
			g.Mov(registros.X0, 1)
			g.Add(registros.HP, registros.HP, registros.X0)
		}
	default:
		panic(fmt.Sprintf("Tipo de objeto no soportado: %v", object.Type_))
	}
	g.PushObjectStack(object)
}

func (g *GeneratorARMInstructions) GetTopObjectStack() ObjectStack {
	if len(g.Stack) == 0 {
		panic("No hay objetos en el stack")
	}
	topObjectStack := g.Stack[len(g.Stack)-1]
	//fmt.Println("id_Top:", topObjectStack.Id_, "tipo:", topObjectStack.Type_, "offset:", topObjectStack.Offset_)
	return topObjectStack
}

func (g *GeneratorARMInstructions) ShowStack() {
	fmt.Println("Contenido del Stack:", len(g.Stack))
	for i, obj := range g.Stack {
		fmt.Printf("Objeto %d: Id: %s, Tipo: %d, Offset: %d, Longitud: %d, Profundidad: %d\n",
			i, obj.Id_, obj.Type_, obj.Offset_, obj.Length_, obj.Depth_)
	}
}

// GetObject busca un objeto en el Stack por su Id_ y devuelve el desplazamiento (Offset_) y el objeto encontrado
func (g *GeneratorARMInstructions) GetObject(id string) (int, ObjectStack) {
	byteOffset := 0

	for i := 0; i < len(g.Stack); i++ {
		if g.Stack[i].Id_ == id {
			return g.Stack[i].Offset_, g.Stack[i]
		}
		byteOffset += g.Stack[i].Length_
	}

	panic(fmt.Sprintf("No se encontró el objeto %s", id))
}

func (g *GeneratorARMInstructions) TagObjecto(id string) {
	if len(g.Stack) == 0 {
		panic("No hay objetos en el stack para etiquetar")
	}
	fmt.Println("g.Stack[len(g.Stack)-1].Id_:", g.Stack[len(g.Stack)-1].Id_, "con id:", id, "y tipo:", g.Stack[len(g.Stack)-1].Type_)
	g.Stack[len(g.Stack)-1].Id_ = id
}

func (g *GeneratorARMInstructions) StrObject() ObjectStack {
	return ObjectStack{
		Type_:   StringType,
		Length_: 8,
		Depth_:  g.depth,
		Id_:     "",
	}
}

func (g *GeneratorARMInstructions) IntObject() ObjectStack {
	return ObjectStack{
		Type_:   Int,
		Length_: 8,
		Depth_:  g.depth,
		Id_:     "",
	}
}

func (g *GeneratorARMInstructions) FloatObject() ObjectStack {
	return ObjectStack{
		Type_:   Float,
		Length_: 8,
		Depth_:  g.depth,
		Id_:     "",
	}
}

func (g *GeneratorARMInstructions) BoolObject() ObjectStack {
	return ObjectStack{
		Type_:   Bool,
		Length_: 8,
		Depth_:  g.depth,
		Id_:     "",
	}
}

func (g *GeneratorARMInstructions) CloneObject(object ObjectStack) ObjectStack {
	return ObjectStack{
		Type_:         object.Type_,
		Length_:       object.Length_,
		Depth_:        g.depth,
		Id_:           object.Id_,
		TipoElemento_: object.TipoElemento_,
		Offset_:       object.Offset_,
	}
}

func (g *GeneratorARMInstructions) GenerateCodeARM() string {
	var sb strings.Builder

	sb.WriteString(".data\n")
	sb.WriteString("buffer: .space 256\n")
	sb.WriteString("salto_linea_str: .ascii \"\\n\"\n")
	sb.WriteString("espacio_str: .ascii \" \"\n")
	sb.WriteString("corchete_apertura: .ascii \"[\"\n")
	sb.WriteString("corchete_cierre: .ascii \"]\"\n")
	sb.WriteString("coma_espacio: .ascii \", \"\n")
	sb.WriteString("heap: .space 4096\n")
	sb.WriteString(".text\n")
	sb.WriteString(".global _start\n")
	sb.WriteString("_start:\n")
	sb.WriteString(" mov x29, sp\n") // Inicializar el frame pointer (quitar si da error en main)
	sb.WriteString(" adr x10, heap\n")
	//sb.WriteString(" mov sp, x10\n") // Inicializar el stack pointer
	sb.WriteString(" bl main\n") // Llamar a main

	g.EndProgram()

	for _, instr := range g.Instrucciones {
		sb.WriteString(instr + "\n")
	}

	sb.WriteString("\n//Funciones Foraneas:\n")
	for _, funcion := range g.FuncionesInstrucciones {
		sb.WriteString(funcion + "\n")
	}

	sb.WriteString("\n//Funciones De Impresion:\n")
	sb.WriteString(g.Estandar.ObtenerDefinicionFuncion())

	return sb.String()
}
