// compile/utils.go
package compile

import (
	"fmt"
	"strings"
)

// parseSymbolType convierte un string a SymbolType (case‐insensitive).
func parseSymbolType(text string) (SymbolType, error) {
	switch strings.ToUpper(text) {
	case "INT":
		return INT, nil
	case "FLOAT64":
		return FLOAT64, nil
	case "STRING":
		return STRING, nil
	case "BOOL":
		return BOOL, nil
	case "RUNE":
		return RUNE, nil
	case "SLICE":
		return SLICE, nil
	case "STRUCT":
		return STRUCT, nil
	case "VOID":
		return VOID, nil
	default:
		return VOID, fmt.Errorf("tipo no reconocido: %s", text)
	}
}

// IsValidType verifica que el valor sea compatible con el SymbolType dado.
// Ejemplo muy básico; adáptalo a tus necesidades (considerando slices, structs, etc.).
func IsValidType(value interface{}, typ SymbolType) bool {
	switch typ {
	case INT:
		_, ok := value.(int)
		return ok
	case FLOAT64:
		_, okFloat := value.(float64)
		if okInt, ok := value.(int); ok {
			// permite int en float64, aunque esto normalmente se maneje antes.
			_ = okInt
			return true
		}
		return okFloat
	case STRING:
		_, ok := value.(string)
		return ok
	case BOOL:
		_, ok := value.(bool)
		return ok
	case RUNE:
		_, ok := value.(rune)
		return ok
	case SLICE:
		_, ok := value.([]interface{})
		return ok
	case STRUCT:
		// aquí deberías validar tu propia estructura
		return true
	default:
		return false
	}
}

// GetDimensionSlice debería inspeccionar []interface{} y devolver el tipo en string,
// p. ej. "[][]INT" o "[]STRING", según la profundidad. Este es un esqueleto muy básico:
//
// - lista      : el slice de elementos (cada elemento puede ser otro []interface{} o valor primitivo).
// - prefix     : ayuda recursiva para acumular "[]", inicia vacío.
// - infer      : si true, inferir el tipo híbrido más profundo.
//
// Devuelve algo como "SLICE" o, si tú lo quieres más detallado, "INT" / "STRING", etc.
func GetDimensionSlice(lista []interface{}, prefix string, infer bool) (string, error) {
	if len(lista) == 0 {
		// Asumamos que un slice vacío es []interface{}
		return "SLICE", nil
	}
	// Si el primer elemento es otro slice, continuamos recursión
	if inner, ok := lista[0].([]interface{}); ok {
		return GetDimensionSlice(inner, prefix+"[]", infer)
	}
	// Si es un valor primitivo, devolvemos su tipo
	switch lista[0].(type) {
	case int:
		return prefix + "INT", nil
	case float64:
		return prefix + "FLOAT64", nil
	case string:
		return prefix + "STRING", nil
	case bool:
		return prefix + "BOOL", nil
	case rune:
		return prefix + "RUNE", nil
	default:
		return "", fmt.Errorf("tipo no soportado en slice: %T", lista[0])
	}
}
