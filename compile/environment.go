// compile/environment.go
package compile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Tipos de símbolos
type SymbolType int

const (
	INT SymbolType = iota
	FLOAT64
	STRING
	BOOL
	RUNE
	SLICE
	STRUCT
	VOID
)

// Estructura del símbolo
type Symbol struct {
	Value   interface{}
	Type    SymbolType
	Mutable bool
}

func NewSymbol(value interface{}, typ SymbolType, mutable bool) *Symbol {
	return &Symbol{
		Value:   value,
		Type:    typ,
		Mutable: mutable,
	}
}

// Tupla personalizada para almacenar información del símbolo
type SymbolTableEntry struct {
	ID     string
	Symbol *Symbol
	Line   int
	Col    int
}

// Entorno
type Environment struct {
	tableSymbol []*SymbolTableEntry
	Variables   map[string]*Symbol
	functions   map[string]*MiFunct
	Parent      *Environment
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		tableSymbol: make([]*SymbolTableEntry, 0),
		Variables:   make(map[string]*Symbol),
		functions:   make(map[string]*MiFunct),
		Parent:      parent,
	}
}

func (e *Environment) GetVariable(id string) (*Symbol, error) {
	if sym, ok := e.Variables[id]; ok {
		return sym, nil
	} else if e.Parent != nil {
		return e.Parent.GetVariable(id)
	}
	return nil, errors.New("variable " + id + " not found")
}

func (e *Environment) SetVariable(id string, value interface{}, typ SymbolType, mutable bool, declaracion bool, token antlr.Token) error {
	fmt.Println("entrando a setVariable")

	// Caso 2: Es una declaración
	if declaracion {
		// Verificar si ya existe en cualquier scope superior
		if e.ExistsVariable(id) {
			fmt.Println("cannot redeclare variable '" + id + "'")
			return errors.New("cannot redeclare variable '" + id + "'")
		}

		// Si no existe, declaramos en este entorno
		e.Variables[id] = NewSymbol(value, typ, mutable)
		e.tableSymbol = append(e.tableSymbol, &SymbolTableEntry{
			ID:     id,
			Symbol: NewSymbol(value, typ, mutable),
			Line:   token.GetLine(),
			Col:    token.GetColumn(),
		})
		return nil
	}

	// Caso 1: Es una reasignación
	if sym, ok := e.Variables[id]; ok {
		sym.Value = value
		e.tableSymbol = append(e.tableSymbol, &SymbolTableEntry{
			ID:     id,
			Symbol: NewSymbol(value, typ, mutable),
			Line:   token.GetLine(),
			Col:    token.GetColumn(),
		})
		return nil // Sin error
	}

	// Caso 3: No es reasignación ni declaración, buscar en el padre
	if e.Parent != nil {
		_, err := e.Parent.GetVariable(id)
		if err == nil {
			return e.Parent.SetVariable(id, value, typ, mutable, false, token)
		}
		return err
	}

	// Si llegamos aquí, la variable no existe y no hay padre
	return errors.New("variable " + id + " not found and cannot be assigned without declaration")
}

func (e *Environment) SetFunciones(id string, parametros []*TupleStringSymbol, body interface{}, typeRet SymbolType, token antlr.Token) {
	funcion := NewMiFunct(parametros, body, typeRet)
	if _, ok := e.functions[id]; ok {
		e.functions[id] = funcion
	} else {
		e.functions[id] = funcion
	}
}

func (e *Environment) GetFuncion(id string) (*MiFunct, error) {
	if fn, ok := e.functions[id]; ok {
		return fn, nil
	} else if e.Parent != nil {
		return e.Parent.GetFuncion(id)
	}
	return nil, errors.New("function " + id + " not found")
}

// --------------------------------------------------------
func (e *Environment) ExistsVariable(id string) bool {
	// Verificar en el entorno actual
	if _, ok := e.Variables[id]; ok {
		fmt.Println("Si existe: " + id)
		return true
	}

	// Si hay un entorno padre, buscar allí recursivamente
	if e.Parent != nil {
		return e.Parent.ExistsVariable(id)
	}

	// Si no se encontró en ningún lado
	fmt.Println("No existe: " + id)
	return false
}

// ----------------------------------------------------
// ImprimirScope imprime todas las variables definidas en el entorno actual
func (e *Environment) ImprimirScope() string {
	var output strings.Builder

	output.WriteString("=== Scope Actual ===\n")
	for _, entry := range e.tableSymbol {
		value := entry.Symbol.Value
		valueStr := fmt.Sprintf("%v", value)

		// Para evitar mostrar datos binarios o estructuras complejas sin formato
		if str, ok := value.(string); ok {
			valueStr = fmt.Sprintf("\"%s\"", str)
		}

		output.WriteString(fmt.Sprintf(
			"ID: %-10s | Valor: %-15s | Tipo: %-10s | Línea: %d | Columna: %d\n",
			entry.ID,
			valueStr,
			symbolTypeToString(entry.Symbol.Type),
			entry.Line,
			entry.Col,
		))
	}
	return output.String()
}

// symbolTypeToString convierte un SymbolType a una cadena legible
func symbolTypeToString(t SymbolType) string {
	switch t {
	case INT:
		return "int"
	case FLOAT64:
		return "float64"
	case STRING:
		return "string"
	case BOOL:
		return "bool"
	case RUNE:
		return "rune"
	case SLICE:
		return "slice"
	case STRUCT:
		return "struct"
	case VOID:
		return "void"
	default:
		return "desconocido"
	}
}
