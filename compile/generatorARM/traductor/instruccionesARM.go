package traductor

import "OLC2CLIENTE/compile/generatorARM/funciones"

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

type ObjectoStack struct {
	Type_         TypeObject
	Length_       int
	Depth_        int
	Id_           *string // puntero para permitir nil
	TipoElemento_ TypeObject
	Offset_       int
}

type generatorARMIntruccions struct {
	Instrucciones          []string
	FuncionesInstrucciones []string
	Estandar               *funciones.EstandarFunc
	Stack                  []ObjectoStack

	depth             int
	contadorEtiquetas int
}

func NewGeneratorARMInstructions() *generatorARMIntruccions {
	return &generatorARMIntruccions{
		Instrucciones:          []string{},
		FuncionesInstrucciones: []string{},
		Estandar:               funciones.NewEstandarFunc(),
		Stack:                  []ObjectoStack{},
		depth:                  0,
		contadorEtiquetas:      0,
	}
}
