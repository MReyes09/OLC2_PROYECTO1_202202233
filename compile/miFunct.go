// compile/miFunct.go
package compile

// TupleStringSymbol simula una tupla (string, *Symbol)
type TupleStringSymbol struct {
	Key   string
	Value *Symbol
}

func NewTupleStringSymbol(key string, value *Symbol) *TupleStringSymbol {
	return &TupleStringSymbol{
		Key:   key,
		Value: value,
	}
}

// MiFunct representa una función definida por el usuario
type MiFunct struct {
	Parameters []*TupleStringSymbol
	Body       interface{} // gramaticaParser.BlockContext en tu parser
	ValRet     SymbolType
	IsSlice    bool
}

func NewMiFunct(parameters []*TupleStringSymbol, body interface{}, valRet SymbolType) *MiFunct {
	return &MiFunct{
		Parameters: parameters,
		Body:       body,
		ValRet:     valRet,
	}
}
