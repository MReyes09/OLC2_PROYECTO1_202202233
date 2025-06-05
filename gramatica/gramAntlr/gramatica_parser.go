// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gramAntlr // gramatica
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type gramaticaParser struct {
	*antlr.BaseParser
}

var GramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'println('", "','", "')'", "';'", "'print('", "'if'",
		"'else'", "'switch'", "'case'", "':'", "'default:'", "'for'", "':='",
		"'range'", "'mut'", "'='", "'var'", "'[]'", "'type'", "'struct'", "'+='",
		"'-='", "'++'", "'--'", "'['", "']'", "'.'", "'-'", "'!'", "'*'", "'/'",
		"'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'nil'", "'('", "'slices.Index('", "'strings.Join('", "'len('",
		"'append('", "'strconv.Atoi('", "'strconv.ParseFloat('", "'reflect.TypeOf('",
		"'int'", "'float64'", "'string'", "'bool'", "'rune'", "'break'", "'continue'",
		"'func'", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "INT", "DOUBLE", "CHAR",
		"STRING", "BOOL", "BLANCOS", "ID_VARIABLE", "COMENTARIOLINEA", "COMENTARIOMULTILINEA",
	}
	staticData.RuleNames = []string{
		"inicio", "instrucciones", "imprimir", "sIf", "sSwitch", "cases", "block",
		"sFor", "varDcl", "varDclSlice", "assign", "nuevoSlice", "contenidoSlice",
		"varDclStruct", "varStructDcl", "varAsign", "expr", "posicion", "type",
		"break", "continue", "functions", "functionStruct", "defParams", "varCallStatement",
		"varCallFuncStruct", "valRet", "retorno",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 553, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 68, 8, 1, 10, 1, 12, 1, 71, 9, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 87, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 93, 8, 2, 10, 2, 12, 2,
		96, 9, 2, 3, 2, 98, 8, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 108, 8, 2, 10, 2, 12, 2, 111, 9, 2, 3, 2, 113, 8, 2, 1, 2,
		1, 2, 3, 2, 117, 8, 2, 3, 2, 119, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 126, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 134, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 146, 8, 5,
		10, 5, 12, 5, 149, 9, 5, 1, 5, 3, 5, 152, 8, 5, 1, 5, 1, 5, 5, 5, 156,
		8, 5, 10, 5, 12, 5, 159, 9, 5, 3, 5, 161, 8, 5, 1, 6, 1, 6, 5, 6, 165,
		8, 6, 10, 6, 12, 6, 168, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 192, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 206, 8, 8, 1, 9, 1, 9, 1, 9,
		4, 9, 211, 8, 9, 11, 9, 12, 9, 212, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 4, 9, 223, 8, 9, 11, 9, 12, 9, 224, 1, 9, 1, 9, 3, 9, 229,
		8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 238, 8, 12,
		10, 12, 12, 12, 241, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 251, 8, 12, 5, 12, 253, 8, 12, 10, 12, 12, 12, 256, 9,
		12, 3, 12, 258, 8, 12, 1, 13, 3, 13, 261, 8, 13, 1, 13, 3, 13, 264, 8,
		13, 1, 13, 1, 13, 3, 13, 268, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		274, 8, 13, 4, 13, 276, 8, 13, 11, 13, 12, 13, 277, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4,
		14, 293, 8, 14, 11, 14, 12, 14, 294, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 310, 8,
		14, 4, 14, 312, 8, 14, 11, 14, 12, 14, 313, 1, 14, 1, 14, 3, 14, 318, 8,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 4, 15, 333, 8, 15, 11, 15, 12, 15, 334, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 343, 8, 15, 11, 15, 12, 15, 344,
		1, 15, 1, 15, 3, 15, 349, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 372, 8, 16, 11, 16, 12, 16, 373,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 391, 8, 16, 10, 16, 12, 16, 394,
		9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 4, 16, 418, 8, 16, 11, 16, 12, 16, 419, 1, 16, 3, 16, 423,
		8, 16, 1, 16, 1, 16, 3, 16, 427, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		5, 16, 444, 8, 16, 10, 16, 12, 16, 447, 9, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 457, 8, 19, 1, 20, 1, 20, 3, 20,
		461, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5,
		21, 471, 8, 21, 10, 21, 12, 21, 474, 9, 21, 3, 21, 476, 8, 21, 1, 21, 1,
		21, 3, 21, 480, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 492, 8, 22, 1, 22, 1, 22, 3, 22, 496, 8, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 505, 8, 23, 10,
		23, 12, 23, 508, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 515,
		8, 24, 10, 24, 12, 24, 518, 9, 24, 3, 24, 520, 8, 24, 1, 24, 1, 24, 3,
		24, 524, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25,
		533, 8, 25, 10, 25, 12, 25, 536, 9, 25, 3, 25, 538, 8, 25, 1, 25, 1, 25,
		3, 25, 542, 8, 25, 1, 26, 1, 26, 1, 27, 1, 27, 3, 27, 548, 8, 27, 1, 27,
		3, 27, 551, 8, 27, 1, 27, 0, 1, 32, 28, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 0, 9, 2, 0, 15, 15, 18, 18, 1, 0, 23, 24, 1, 0, 25, 26, 1, 0, 32, 34,
		2, 0, 30, 30, 35, 35, 1, 0, 36, 39, 1, 0, 40, 41, 1, 0, 42, 43, 3, 0, 22,
		22, 53, 57, 68, 68, 627, 0, 59, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 118,
		1, 0, 0, 0, 6, 133, 1, 0, 0, 0, 8, 135, 1, 0, 0, 0, 10, 160, 1, 0, 0, 0,
		12, 162, 1, 0, 0, 0, 14, 191, 1, 0, 0, 0, 16, 205, 1, 0, 0, 0, 18, 228,
		1, 0, 0, 0, 20, 230, 1, 0, 0, 0, 22, 232, 1, 0, 0, 0, 24, 257, 1, 0, 0,
		0, 26, 260, 1, 0, 0, 0, 28, 317, 1, 0, 0, 0, 30, 348, 1, 0, 0, 0, 32, 426,
		1, 0, 0, 0, 34, 448, 1, 0, 0, 0, 36, 452, 1, 0, 0, 0, 38, 454, 1, 0, 0,
		0, 40, 458, 1, 0, 0, 0, 42, 462, 1, 0, 0, 0, 44, 483, 1, 0, 0, 0, 46, 499,
		1, 0, 0, 0, 48, 509, 1, 0, 0, 0, 50, 525, 1, 0, 0, 0, 52, 543, 1, 0, 0,
		0, 54, 545, 1, 0, 0, 0, 56, 58, 3, 2, 1, 0, 57, 56, 1, 0, 0, 0, 58, 61,
		1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 1, 1, 0, 0, 0,
		61, 59, 1, 0, 0, 0, 62, 87, 3, 4, 2, 0, 63, 87, 3, 6, 3, 0, 64, 87, 3,
		8, 4, 0, 65, 69, 5, 1, 0, 0, 66, 68, 3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 68,
		71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0,
		0, 71, 69, 1, 0, 0, 0, 72, 87, 5, 2, 0, 0, 73, 87, 3, 14, 7, 0, 74, 87,
		3, 18, 9, 0, 75, 87, 3, 30, 15, 0, 76, 87, 3, 16, 8, 0, 77, 87, 3, 26,
		13, 0, 78, 87, 3, 28, 14, 0, 79, 87, 3, 38, 19, 0, 80, 87, 3, 40, 20, 0,
		81, 87, 3, 42, 21, 0, 82, 87, 3, 44, 22, 0, 83, 87, 3, 48, 24, 0, 84, 87,
		3, 50, 25, 0, 85, 87, 3, 54, 27, 0, 86, 62, 1, 0, 0, 0, 86, 63, 1, 0, 0,
		0, 86, 64, 1, 0, 0, 0, 86, 65, 1, 0, 0, 0, 86, 73, 1, 0, 0, 0, 86, 74,
		1, 0, 0, 0, 86, 75, 1, 0, 0, 0, 86, 76, 1, 0, 0, 0, 86, 77, 1, 0, 0, 0,
		86, 78, 1, 0, 0, 0, 86, 79, 1, 0, 0, 0, 86, 80, 1, 0, 0, 0, 86, 81, 1,
		0, 0, 0, 86, 82, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86,
		85, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 97, 5, 3, 0, 0, 89, 94, 3, 32, 16,
		0, 90, 91, 5, 4, 0, 0, 91, 93, 3, 32, 16, 0, 92, 90, 1, 0, 0, 0, 93, 96,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 101, 5, 5, 0, 0, 100, 102, 5, 6, 0, 0, 101, 100, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 119, 1, 0, 0, 0, 103, 112, 5, 7, 0, 0, 104,
		109, 3, 32, 16, 0, 105, 106, 5, 4, 0, 0, 106, 108, 3, 32, 16, 0, 107, 105,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 104, 1, 0, 0, 0,
		112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 5, 5, 0, 0, 115,
		117, 5, 6, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119,
		1, 0, 0, 0, 118, 88, 1, 0, 0, 0, 118, 103, 1, 0, 0, 0, 119, 5, 1, 0, 0,
		0, 120, 121, 5, 8, 0, 0, 121, 122, 3, 32, 16, 0, 122, 125, 3, 12, 6, 0,
		123, 124, 5, 9, 0, 0, 124, 126, 3, 12, 6, 0, 125, 123, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 134, 1, 0, 0, 0, 127, 128, 5, 8, 0, 0, 128, 129,
		3, 32, 16, 0, 129, 130, 3, 12, 6, 0, 130, 131, 5, 9, 0, 0, 131, 132, 3,
		6, 3, 0, 132, 134, 1, 0, 0, 0, 133, 120, 1, 0, 0, 0, 133, 127, 1, 0, 0,
		0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 10, 0, 0, 136, 137, 3, 32, 16, 0, 137,
		138, 5, 1, 0, 0, 138, 139, 3, 10, 5, 0, 139, 140, 5, 2, 0, 0, 140, 9, 1,
		0, 0, 0, 141, 142, 5, 11, 0, 0, 142, 143, 3, 32, 16, 0, 143, 147, 5, 12,
		0, 0, 144, 146, 3, 2, 1, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0,
		147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 150, 152, 3, 10, 5, 0, 151, 150, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 161, 1, 0, 0, 0, 153, 157, 5, 13, 0, 0, 154, 156, 3, 2,
		1, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160,
		141, 1, 0, 0, 0, 160, 153, 1, 0, 0, 0, 161, 11, 1, 0, 0, 0, 162, 166, 5,
		1, 0, 0, 163, 165, 3, 2, 1, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 169, 170, 5, 2, 0, 0, 170, 13, 1, 0, 0, 0, 171, 172, 5,
		14, 0, 0, 172, 173, 3, 32, 16, 0, 173, 174, 3, 12, 6, 0, 174, 192, 1, 0,
		0, 0, 175, 176, 5, 14, 0, 0, 176, 177, 3, 16, 8, 0, 177, 178, 5, 6, 0,
		0, 178, 179, 3, 32, 16, 0, 179, 180, 5, 6, 0, 0, 180, 181, 3, 30, 15, 0,
		181, 182, 3, 12, 6, 0, 182, 192, 1, 0, 0, 0, 183, 184, 5, 14, 0, 0, 184,
		185, 5, 68, 0, 0, 185, 186, 5, 4, 0, 0, 186, 187, 5, 68, 0, 0, 187, 188,
		5, 15, 0, 0, 188, 189, 5, 16, 0, 0, 189, 190, 5, 68, 0, 0, 190, 192, 3,
		12, 6, 0, 191, 171, 1, 0, 0, 0, 191, 175, 1, 0, 0, 0, 191, 183, 1, 0, 0,
		0, 192, 15, 1, 0, 0, 0, 193, 194, 5, 17, 0, 0, 194, 195, 5, 68, 0, 0, 195,
		196, 3, 36, 18, 0, 196, 197, 5, 18, 0, 0, 197, 198, 3, 32, 16, 0, 198,
		206, 1, 0, 0, 0, 199, 200, 5, 17, 0, 0, 200, 201, 5, 68, 0, 0, 201, 206,
		3, 36, 18, 0, 202, 203, 5, 68, 0, 0, 203, 204, 5, 15, 0, 0, 204, 206, 3,
		32, 16, 0, 205, 193, 1, 0, 0, 0, 205, 199, 1, 0, 0, 0, 205, 202, 1, 0,
		0, 0, 206, 17, 1, 0, 0, 0, 207, 208, 5, 68, 0, 0, 208, 210, 3, 20, 10,
		0, 209, 211, 3, 22, 11, 0, 210, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0,
		212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		215, 3, 36, 18, 0, 215, 216, 5, 1, 0, 0, 216, 217, 3, 24, 12, 0, 217, 218,
		5, 2, 0, 0, 218, 229, 1, 0, 0, 0, 219, 220, 5, 19, 0, 0, 220, 222, 5, 68,
		0, 0, 221, 223, 3, 22, 11, 0, 222, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0,
		0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226,
		227, 3, 36, 18, 0, 227, 229, 1, 0, 0, 0, 228, 207, 1, 0, 0, 0, 228, 219,
		1, 0, 0, 0, 229, 19, 1, 0, 0, 0, 230, 231, 7, 0, 0, 0, 231, 21, 1, 0, 0,
		0, 232, 233, 5, 20, 0, 0, 233, 23, 1, 0, 0, 0, 234, 239, 3, 32, 16, 0,
		235, 236, 5, 4, 0, 0, 236, 238, 3, 32, 16, 0, 237, 235, 1, 0, 0, 0, 238,
		241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 258,
		1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 243, 5, 1, 0, 0, 243, 244, 3, 24,
		12, 0, 244, 254, 5, 2, 0, 0, 245, 250, 5, 4, 0, 0, 246, 247, 5, 1, 0, 0,
		247, 248, 3, 24, 12, 0, 248, 249, 5, 2, 0, 0, 249, 251, 1, 0, 0, 0, 250,
		246, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 245,
		1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0,
		0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 234, 1, 0, 0, 0,
		257, 242, 1, 0, 0, 0, 258, 25, 1, 0, 0, 0, 259, 261, 5, 21, 0, 0, 260,
		259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 263, 1, 0, 0, 0, 262, 264,
		5, 22, 0, 0, 263, 262, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 1, 0,
		0, 0, 265, 267, 5, 68, 0, 0, 266, 268, 5, 22, 0, 0, 267, 266, 1, 0, 0,
		0, 267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 275, 5, 1, 0, 0, 270,
		271, 5, 68, 0, 0, 271, 273, 3, 36, 18, 0, 272, 274, 5, 6, 0, 0, 273, 272,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 270, 1, 0,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0,
		278, 279, 1, 0, 0, 0, 279, 280, 5, 2, 0, 0, 280, 27, 1, 0, 0, 0, 281, 282,
		5, 68, 0, 0, 282, 283, 5, 68, 0, 0, 283, 284, 5, 18, 0, 0, 284, 285, 5,
		1, 0, 0, 285, 286, 5, 68, 0, 0, 286, 287, 5, 12, 0, 0, 287, 292, 3, 32,
		16, 0, 288, 289, 5, 4, 0, 0, 289, 290, 5, 68, 0, 0, 290, 291, 5, 12, 0,
		0, 291, 293, 3, 32, 16, 0, 292, 288, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0,
		294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		297, 5, 2, 0, 0, 297, 318, 1, 0, 0, 0, 298, 299, 5, 68, 0, 0, 299, 300,
		5, 15, 0, 0, 300, 301, 5, 68, 0, 0, 301, 302, 5, 1, 0, 0, 302, 303, 5,
		68, 0, 0, 303, 304, 5, 12, 0, 0, 304, 311, 3, 32, 16, 0, 305, 309, 5, 4,
		0, 0, 306, 307, 5, 68, 0, 0, 307, 308, 5, 12, 0, 0, 308, 310, 3, 32, 16,
		0, 309, 306, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 1, 0, 0, 0, 311,
		305, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314,
		1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 5, 2, 0, 0, 316, 318, 1, 0,
		0, 0, 317, 281, 1, 0, 0, 0, 317, 298, 1, 0, 0, 0, 318, 29, 1, 0, 0, 0,
		319, 320, 5, 68, 0, 0, 320, 321, 5, 18, 0, 0, 321, 349, 3, 32, 16, 0, 322,
		323, 5, 68, 0, 0, 323, 324, 7, 1, 0, 0, 324, 349, 3, 32, 16, 0, 325, 326,
		5, 68, 0, 0, 326, 349, 7, 2, 0, 0, 327, 332, 5, 68, 0, 0, 328, 329, 5,
		27, 0, 0, 329, 330, 3, 32, 16, 0, 330, 331, 5, 28, 0, 0, 331, 333, 1, 0,
		0, 0, 332, 328, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0,
		334, 335, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 5, 18, 0, 0, 337,
		338, 3, 32, 16, 0, 338, 349, 1, 0, 0, 0, 339, 342, 5, 68, 0, 0, 340, 341,
		5, 29, 0, 0, 341, 343, 5, 68, 0, 0, 342, 340, 1, 0, 0, 0, 343, 344, 1,
		0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0, 0,
		0, 346, 347, 5, 18, 0, 0, 347, 349, 3, 32, 16, 0, 348, 319, 1, 0, 0, 0,
		348, 322, 1, 0, 0, 0, 348, 325, 1, 0, 0, 0, 348, 327, 1, 0, 0, 0, 348,
		339, 1, 0, 0, 0, 349, 31, 1, 0, 0, 0, 350, 351, 6, 16, -1, 0, 351, 352,
		5, 30, 0, 0, 352, 427, 3, 32, 16, 26, 353, 354, 5, 31, 0, 0, 354, 427,
		3, 32, 16, 25, 355, 427, 5, 62, 0, 0, 356, 427, 5, 63, 0, 0, 357, 427,
		5, 65, 0, 0, 358, 427, 5, 66, 0, 0, 359, 427, 5, 68, 0, 0, 360, 427, 5,
		64, 0, 0, 361, 427, 5, 44, 0, 0, 362, 363, 5, 45, 0, 0, 363, 364, 3, 32,
		16, 0, 364, 365, 5, 5, 0, 0, 365, 427, 1, 0, 0, 0, 366, 371, 5, 68, 0,
		0, 367, 368, 5, 27, 0, 0, 368, 369, 3, 32, 16, 0, 369, 370, 5, 28, 0, 0,
		370, 372, 1, 0, 0, 0, 371, 367, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373,
		371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 427, 1, 0, 0, 0, 375, 376,
		5, 46, 0, 0, 376, 377, 5, 68, 0, 0, 377, 378, 5, 4, 0, 0, 378, 379, 3,
		32, 16, 0, 379, 380, 5, 5, 0, 0, 380, 427, 1, 0, 0, 0, 381, 382, 5, 47,
		0, 0, 382, 383, 5, 68, 0, 0, 383, 384, 5, 4, 0, 0, 384, 385, 3, 32, 16,
		0, 385, 386, 5, 5, 0, 0, 386, 427, 1, 0, 0, 0, 387, 388, 5, 48, 0, 0, 388,
		392, 5, 68, 0, 0, 389, 391, 3, 34, 17, 0, 390, 389, 1, 0, 0, 0, 391, 394,
		1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 1, 0,
		0, 0, 394, 392, 1, 0, 0, 0, 395, 427, 5, 5, 0, 0, 396, 397, 5, 49, 0, 0,
		397, 398, 5, 68, 0, 0, 398, 399, 5, 4, 0, 0, 399, 400, 3, 32, 16, 0, 400,
		401, 5, 5, 0, 0, 401, 427, 1, 0, 0, 0, 402, 403, 5, 50, 0, 0, 403, 404,
		3, 32, 16, 0, 404, 405, 5, 5, 0, 0, 405, 427, 1, 0, 0, 0, 406, 407, 5,
		51, 0, 0, 407, 408, 3, 32, 16, 0, 408, 409, 5, 5, 0, 0, 409, 427, 1, 0,
		0, 0, 410, 411, 5, 52, 0, 0, 411, 412, 3, 32, 16, 0, 412, 413, 5, 5, 0,
		0, 413, 427, 1, 0, 0, 0, 414, 417, 5, 68, 0, 0, 415, 416, 5, 29, 0, 0,
		416, 418, 5, 68, 0, 0, 417, 415, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419,
		417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 422, 1, 0, 0, 0, 421, 423,
		5, 6, 0, 0, 422, 421, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 427, 1, 0,
		0, 0, 424, 427, 3, 48, 24, 0, 425, 427, 3, 50, 25, 0, 426, 350, 1, 0, 0,
		0, 426, 353, 1, 0, 0, 0, 426, 355, 1, 0, 0, 0, 426, 356, 1, 0, 0, 0, 426,
		357, 1, 0, 0, 0, 426, 358, 1, 0, 0, 0, 426, 359, 1, 0, 0, 0, 426, 360,
		1, 0, 0, 0, 426, 361, 1, 0, 0, 0, 426, 362, 1, 0, 0, 0, 426, 366, 1, 0,
		0, 0, 426, 375, 1, 0, 0, 0, 426, 381, 1, 0, 0, 0, 426, 387, 1, 0, 0, 0,
		426, 396, 1, 0, 0, 0, 426, 402, 1, 0, 0, 0, 426, 406, 1, 0, 0, 0, 426,
		410, 1, 0, 0, 0, 426, 414, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 425,
		1, 0, 0, 0, 427, 445, 1, 0, 0, 0, 428, 429, 10, 24, 0, 0, 429, 430, 7,
		3, 0, 0, 430, 444, 3, 32, 16, 25, 431, 432, 10, 23, 0, 0, 432, 433, 7,
		4, 0, 0, 433, 444, 3, 32, 16, 24, 434, 435, 10, 22, 0, 0, 435, 436, 7,
		5, 0, 0, 436, 444, 3, 32, 16, 23, 437, 438, 10, 21, 0, 0, 438, 439, 7,
		6, 0, 0, 439, 444, 3, 32, 16, 22, 440, 441, 10, 20, 0, 0, 441, 442, 7,
		7, 0, 0, 442, 444, 3, 32, 16, 21, 443, 428, 1, 0, 0, 0, 443, 431, 1, 0,
		0, 0, 443, 434, 1, 0, 0, 0, 443, 437, 1, 0, 0, 0, 443, 440, 1, 0, 0, 0,
		444, 447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446,
		33, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 448, 449, 5, 27, 0, 0, 449, 450,
		3, 32, 16, 0, 450, 451, 5, 28, 0, 0, 451, 35, 1, 0, 0, 0, 452, 453, 7,
		8, 0, 0, 453, 37, 1, 0, 0, 0, 454, 456, 5, 58, 0, 0, 455, 457, 5, 6, 0,
		0, 456, 455, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 39, 1, 0, 0, 0, 458,
		460, 5, 59, 0, 0, 459, 461, 5, 6, 0, 0, 460, 459, 1, 0, 0, 0, 460, 461,
		1, 0, 0, 0, 461, 41, 1, 0, 0, 0, 462, 463, 5, 60, 0, 0, 463, 464, 5, 68,
		0, 0, 464, 475, 5, 45, 0, 0, 465, 466, 5, 68, 0, 0, 466, 472, 3, 36, 18,
		0, 467, 468, 5, 4, 0, 0, 468, 469, 5, 68, 0, 0, 469, 471, 3, 36, 18, 0,
		470, 467, 1, 0, 0, 0, 471, 474, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472,
		473, 1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0, 475, 465,
		1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 479, 5, 5,
		0, 0, 478, 480, 3, 52, 26, 0, 479, 478, 1, 0, 0, 0, 479, 480, 1, 0, 0,
		0, 480, 481, 1, 0, 0, 0, 481, 482, 3, 12, 6, 0, 482, 43, 1, 0, 0, 0, 483,
		484, 5, 60, 0, 0, 484, 485, 5, 45, 0, 0, 485, 486, 5, 68, 0, 0, 486, 487,
		5, 68, 0, 0, 487, 488, 5, 5, 0, 0, 488, 489, 5, 68, 0, 0, 489, 491, 5,
		45, 0, 0, 490, 492, 3, 46, 23, 0, 491, 490, 1, 0, 0, 0, 491, 492, 1, 0,
		0, 0, 492, 493, 1, 0, 0, 0, 493, 495, 5, 5, 0, 0, 494, 496, 3, 52, 26,
		0, 495, 494, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497,
		498, 3, 12, 6, 0, 498, 45, 1, 0, 0, 0, 499, 500, 5, 68, 0, 0, 500, 506,
		3, 36, 18, 0, 501, 502, 5, 4, 0, 0, 502, 503, 5, 68, 0, 0, 503, 505, 3,
		36, 18, 0, 504, 501, 1, 0, 0, 0, 505, 508, 1, 0, 0, 0, 506, 504, 1, 0,
		0, 0, 506, 507, 1, 0, 0, 0, 507, 47, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0,
		509, 510, 5, 68, 0, 0, 510, 519, 5, 45, 0, 0, 511, 516, 3, 32, 16, 0, 512,
		513, 5, 4, 0, 0, 513, 515, 3, 32, 16, 0, 514, 512, 1, 0, 0, 0, 515, 518,
		1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 520, 1, 0,
		0, 0, 518, 516, 1, 0, 0, 0, 519, 511, 1, 0, 0, 0, 519, 520, 1, 0, 0, 0,
		520, 521, 1, 0, 0, 0, 521, 523, 5, 5, 0, 0, 522, 524, 5, 6, 0, 0, 523,
		522, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 49, 1, 0, 0, 0, 525, 526, 5,
		68, 0, 0, 526, 527, 5, 29, 0, 0, 527, 528, 5, 68, 0, 0, 528, 537, 5, 45,
		0, 0, 529, 534, 3, 32, 16, 0, 530, 531, 5, 4, 0, 0, 531, 533, 3, 32, 16,
		0, 532, 530, 1, 0, 0, 0, 533, 536, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534,
		535, 1, 0, 0, 0, 535, 538, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 537, 529,
		1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 539, 541, 5, 5,
		0, 0, 540, 542, 5, 6, 0, 0, 541, 540, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0,
		542, 51, 1, 0, 0, 0, 543, 544, 3, 36, 18, 0, 544, 53, 1, 0, 0, 0, 545,
		547, 5, 61, 0, 0, 546, 548, 3, 32, 16, 0, 547, 546, 1, 0, 0, 0, 547, 548,
		1, 0, 0, 0, 548, 550, 1, 0, 0, 0, 549, 551, 5, 6, 0, 0, 550, 549, 1, 0,
		0, 0, 550, 551, 1, 0, 0, 0, 551, 55, 1, 0, 0, 0, 61, 59, 69, 86, 94, 97,
		101, 109, 112, 116, 118, 125, 133, 147, 151, 157, 160, 166, 191, 205, 212,
		224, 228, 239, 250, 254, 257, 260, 263, 267, 273, 277, 294, 309, 313, 317,
		334, 344, 348, 373, 392, 419, 422, 426, 443, 445, 456, 460, 472, 475, 479,
		491, 495, 506, 516, 519, 523, 534, 537, 541, 547, 550,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// gramaticaParserInit initializes any static state used to implement gramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewgramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgramaticaParser(input antlr.TokenStream) *gramaticaParser {
	GramaticaParserInit()
	this := new(gramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gramatica.g4"

	return this
}

// gramaticaParser tokens.
const (
	gramaticaParserEOF                  = antlr.TokenEOF
	gramaticaParserT__0                 = 1
	gramaticaParserT__1                 = 2
	gramaticaParserT__2                 = 3
	gramaticaParserT__3                 = 4
	gramaticaParserT__4                 = 5
	gramaticaParserT__5                 = 6
	gramaticaParserT__6                 = 7
	gramaticaParserT__7                 = 8
	gramaticaParserT__8                 = 9
	gramaticaParserT__9                 = 10
	gramaticaParserT__10                = 11
	gramaticaParserT__11                = 12
	gramaticaParserT__12                = 13
	gramaticaParserT__13                = 14
	gramaticaParserT__14                = 15
	gramaticaParserT__15                = 16
	gramaticaParserT__16                = 17
	gramaticaParserT__17                = 18
	gramaticaParserT__18                = 19
	gramaticaParserT__19                = 20
	gramaticaParserT__20                = 21
	gramaticaParserT__21                = 22
	gramaticaParserT__22                = 23
	gramaticaParserT__23                = 24
	gramaticaParserT__24                = 25
	gramaticaParserT__25                = 26
	gramaticaParserT__26                = 27
	gramaticaParserT__27                = 28
	gramaticaParserT__28                = 29
	gramaticaParserT__29                = 30
	gramaticaParserT__30                = 31
	gramaticaParserT__31                = 32
	gramaticaParserT__32                = 33
	gramaticaParserT__33                = 34
	gramaticaParserT__34                = 35
	gramaticaParserT__35                = 36
	gramaticaParserT__36                = 37
	gramaticaParserT__37                = 38
	gramaticaParserT__38                = 39
	gramaticaParserT__39                = 40
	gramaticaParserT__40                = 41
	gramaticaParserT__41                = 42
	gramaticaParserT__42                = 43
	gramaticaParserT__43                = 44
	gramaticaParserT__44                = 45
	gramaticaParserT__45                = 46
	gramaticaParserT__46                = 47
	gramaticaParserT__47                = 48
	gramaticaParserT__48                = 49
	gramaticaParserT__49                = 50
	gramaticaParserT__50                = 51
	gramaticaParserT__51                = 52
	gramaticaParserT__52                = 53
	gramaticaParserT__53                = 54
	gramaticaParserT__54                = 55
	gramaticaParserT__55                = 56
	gramaticaParserT__56                = 57
	gramaticaParserT__57                = 58
	gramaticaParserT__58                = 59
	gramaticaParserT__59                = 60
	gramaticaParserT__60                = 61
	gramaticaParserINT                  = 62
	gramaticaParserDOUBLE               = 63
	gramaticaParserCHAR                 = 64
	gramaticaParserSTRING               = 65
	gramaticaParserBOOL                 = 66
	gramaticaParserBLANCOS              = 67
	gramaticaParserID_VARIABLE          = 68
	gramaticaParserCOMENTARIOLINEA      = 69
	gramaticaParserCOMENTARIOMULTILINEA = 70
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_inicio            = 0
	gramaticaParserRULE_instrucciones     = 1
	gramaticaParserRULE_imprimir          = 2
	gramaticaParserRULE_sIf               = 3
	gramaticaParserRULE_sSwitch           = 4
	gramaticaParserRULE_cases             = 5
	gramaticaParserRULE_block             = 6
	gramaticaParserRULE_sFor              = 7
	gramaticaParserRULE_varDcl            = 8
	gramaticaParserRULE_varDclSlice       = 9
	gramaticaParserRULE_assign            = 10
	gramaticaParserRULE_nuevoSlice        = 11
	gramaticaParserRULE_contenidoSlice    = 12
	gramaticaParserRULE_varDclStruct      = 13
	gramaticaParserRULE_varStructDcl      = 14
	gramaticaParserRULE_varAsign          = 15
	gramaticaParserRULE_expr              = 16
	gramaticaParserRULE_posicion          = 17
	gramaticaParserRULE_type              = 18
	gramaticaParserRULE_break             = 19
	gramaticaParserRULE_continue          = 20
	gramaticaParserRULE_functions         = 21
	gramaticaParserRULE_functionStruct    = 22
	gramaticaParserRULE_defParams         = 23
	gramaticaParserRULE_varCallStatement  = 24
	gramaticaParserRULE_varCallFuncStruct = 25
	gramaticaParserRULE_valRet            = 26
	gramaticaParserRULE_retorno           = 27
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstrucciones() []IInstruccionesContext
	Instrucciones(i int) IInstruccionesContext

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_inicio
	return p
}

func InitEmptyInicioContext(p *InicioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_inicio
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *InicioContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInicio(s)
	}
}

func (s *InicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gramaticaParserRULE_inicio)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4323455642282640778) != 0) || _la == gramaticaParserID_VARIABLE {
		{
			p.SetState(56)
			p.Instrucciones()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) CopyAll(ctx *InstruccionesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SeccionInstruccionContext struct {
	InstruccionesContext
}

func NewSeccionInstruccionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SeccionInstruccionContext {
	var p = new(SeccionInstruccionContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *SeccionInstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeccionInstruccionContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *SeccionInstruccionContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *SeccionInstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSeccionInstruccion(s)
	}
}

func (s *SeccionInstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSeccionInstruccion(s)
	}
}

func (s *SeccionInstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSeccionInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionStmtContext struct {
	InstruccionesContext
}

func NewCallFunctionStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionStmtContext {
	var p = new(CallFunctionStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *CallFunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionStmtContext) VarCallStatement() IVarCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCallStatementContext)
}

func (s *CallFunctionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunctionStmt(s)
	}
}

func (s *CallFunctionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunctionStmt(s)
	}
}

func (s *CallFunctionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunctionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStmtContext struct {
	InstruccionesContext
}

func NewPrintStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtContext {
	var p = new(PrintStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) Imprimir() IImprimirContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImprimirContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImprimirContext)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarStructDclStmtContext struct {
	InstruccionesContext
}

func NewVarStructDclStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarStructDclStmtContext {
	var p = new(VarStructDclStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *VarStructDclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarStructDclStmtContext) VarStructDcl() IVarStructDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarStructDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarStructDclContext)
}

func (s *VarStructDclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarStructDclStmt(s)
	}
}

func (s *VarStructDclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarStructDclStmt(s)
	}
}

func (s *VarStructDclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarStructDclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignStmtContext struct {
	InstruccionesContext
}

func NewAsignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignStmtContext {
	var p = new(AsignStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *AsignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignStmtContext) VarAsign() IVarAsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAsignContext)
}

func (s *AsignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignStmt(s)
	}
}

func (s *AsignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignStmt(s)
	}
}

func (s *AsignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionStructStmtContext struct {
	InstruccionesContext
}

func NewCallFunctionStructStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionStructStmtContext {
	var p = new(CallFunctionStructStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *CallFunctionStructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionStructStmtContext) VarCallFuncStruct() IVarCallFuncStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCallFuncStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCallFuncStructContext)
}

func (s *CallFunctionStructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunctionStructStmt(s)
	}
}

func (s *CallFunctionStructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunctionStructStmt(s)
	}
}

func (s *CallFunctionStructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunctionStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStmtContext struct {
	InstruccionesContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) Continue_() IContinueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueContext)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclStructStmtContext struct {
	InstruccionesContext
}

func NewVarDeclStructStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclStructStmtContext {
	var p = new(VarDeclStructStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *VarDeclStructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStructStmtContext) VarDclStruct() IVarDclStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclStructContext)
}

func (s *VarDeclStructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDeclStructStmt(s)
	}
}

func (s *VarDeclStructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDeclStructStmt(s)
	}
}

func (s *VarDeclStructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDeclStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStmtContext struct {
	InstruccionesContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) SIf() ISIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISIfContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionStmtContext struct {
	InstruccionesContext
}

func NewFunctionStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionStmtContext {
	var p = new(FunctionStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *FunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStmtContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *FunctionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFunctionStmt(s)
	}
}

func (s *FunctionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFunctionStmt(s)
	}
}

func (s *FunctionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFunctionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionStructStmtContext struct {
	InstruccionesContext
}

func NewFunctionStructStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionStructStmtContext {
	var p = new(FunctionStructStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *FunctionStructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStructStmtContext) FunctionStruct() IFunctionStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStructContext)
}

func (s *FunctionStructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFunctionStructStmt(s)
	}
}

func (s *FunctionStructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFunctionStructStmt(s)
	}
}

func (s *FunctionStructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFunctionStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclSliceStmtContext struct {
	InstruccionesContext
}

func NewVarDeclSliceStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclSliceStmtContext {
	var p = new(VarDeclSliceStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *VarDeclSliceStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclSliceStmtContext) VarDclSlice() IVarDclSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclSliceContext)
}

func (s *VarDeclSliceStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDeclSliceStmt(s)
	}
}

func (s *VarDeclSliceStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDeclSliceStmt(s)
	}
}

func (s *VarDeclSliceStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDeclSliceStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclStmtContext struct {
	InstruccionesContext
}

func NewVarDeclStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *VarDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStmtContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *VarDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	InstruccionesContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) Break_() IBreakContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakContext)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchInstruccionContext struct {
	InstruccionesContext
}

func NewSwitchInstruccionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchInstruccionContext {
	var p = new(SwitchInstruccionContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *SwitchInstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchInstruccionContext) SSwitch() ISSwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISSwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISSwitchContext)
}

func (s *SwitchInstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSwitchInstruccion(s)
	}
}

func (s *SwitchInstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSwitchInstruccion(s)
	}
}

func (s *SwitchInstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSwitchInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStmtContext struct {
	InstruccionesContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) SFor() ISForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISForContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	InstruccionesContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyInstruccionesContext(&p.InstruccionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstruccionesContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) Retorno() IRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetornoContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_instrucciones)
	var _la int

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Imprimir()
		}

	case 2:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.SIf()
		}

	case 3:
		localctx = NewSwitchInstruccionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.SSwitch()
		}

	case 4:
		localctx = NewSeccionInstruccionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Match(gramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4323455642282640778) != 0) || _la == gramaticaParserID_VARIABLE {
			{
				p.SetState(66)
				p.Instrucciones()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(72)
			p.Match(gramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.SFor()
		}

	case 6:
		localctx = NewVarDeclSliceStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.VarDclSlice()
		}

	case 7:
		localctx = NewAsignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(75)
			p.VarAsign()
		}

	case 8:
		localctx = NewVarDeclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(76)
			p.VarDcl()
		}

	case 9:
		localctx = NewVarDeclStructStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(77)
			p.VarDclStruct()
		}

	case 10:
		localctx = NewVarStructDclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(78)
			p.VarStructDcl()
		}

	case 11:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(79)
			p.Break_()
		}

	case 12:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(80)
			p.Continue_()
		}

	case 13:
		localctx = NewFunctionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(81)
			p.Functions()
		}

	case 14:
		localctx = NewFunctionStructStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(82)
			p.FunctionStruct()
		}

	case 15:
		localctx = NewCallFunctionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(83)
			p.VarCallStatement()
		}

	case 16:
		localctx = NewCallFunctionStructStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(84)
			p.VarCallFuncStruct()
		}

	case 17:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(85)
			p.Retorno()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImprimirContext is an interface to support dynamic dispatch.
type IImprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsImprimirContext differentiates from other interfaces.
	IsImprimirContext()
}

type ImprimirContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImprimirContext() *ImprimirContext {
	var p = new(ImprimirContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_imprimir
	return p
}

func InitEmptyImprimirContext(p *ImprimirContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_imprimir
}

func (*ImprimirContext) IsImprimirContext() {}

func NewImprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimirContext {
	var p = new(ImprimirContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_imprimir

	return p
}

func (s *ImprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimirContext) CopyAll(ctx *ImprimirContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ImprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	ImprimirContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	InitEmptyImprimirContext(&p.ImprimirContext)
	p.parser = parser
	p.CopyAll(ctx.(*ImprimirContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnContext struct {
	ImprimirContext
}

func NewPrintlnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnContext {
	var p = new(PrintlnContext)

	InitEmptyImprimirContext(&p.ImprimirContext)
	p.parser = parser
	p.CopyAll(ctx.(*ImprimirContext))

	return p
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintlnContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrintln(s)
	}
}

func (s *PrintlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrintln(s)
	}
}

func (s *PrintlnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPrintln(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Imprimir() (localctx IImprimirContext) {
	localctx = NewImprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gramaticaParserRULE_imprimir)
	var _la int

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserT__2:
		localctx = NewPrintlnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(gramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&408030265347) != 0 {
			{
				p.SetState(89)
				p.expr(0)
			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == gramaticaParserT__3 {
				{
					p.SetState(90)
					p.Match(gramaticaParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(91)
					p.expr(0)
				}

				p.SetState(96)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(99)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__5 {
			{
				p.SetState(100)
				p.Match(gramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case gramaticaParserT__6:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(gramaticaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&408030265347) != 0 {
			{
				p.SetState(104)
				p.expr(0)
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == gramaticaParserT__3 {
				{
					p.SetState(105)
					p.Match(gramaticaParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expr(0)
				}

				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(114)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__5 {
			{
				p.SetState(115)
				p.Match(gramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISIfContext is an interface to support dynamic dispatch.
type ISIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSIfContext differentiates from other interfaces.
	IsSIfContext()
}

type SIfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySIfContext() *SIfContext {
	var p = new(SIfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sIf
	return p
}

func InitEmptySIfContext(p *SIfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sIf
}

func (*SIfContext) IsSIfContext() {}

func NewSIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SIfContext {
	var p = new(SIfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sIf

	return p
}

func (s *SIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SIfContext) CopyAll(ctx *SIfContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfAnidadoContext struct {
	SIfContext
}

func NewIfAnidadoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfAnidadoContext {
	var p = new(IfAnidadoContext)

	InitEmptySIfContext(&p.SIfContext)
	p.parser = parser
	p.CopyAll(ctx.(*SIfContext))

	return p
}

func (s *IfAnidadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfAnidadoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfAnidadoContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfAnidadoContext) SIf() ISIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISIfContext)
}

func (s *IfAnidadoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIfAnidado(s)
	}
}

func (s *IfAnidadoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIfAnidado(s)
	}
}

func (s *IfAnidadoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIfAnidado(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfOnlyContext struct {
	SIfContext
}

func NewIfOnlyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfOnlyContext {
	var p = new(IfOnlyContext)

	InitEmptySIfContext(&p.SIfContext)
	p.parser = parser
	p.CopyAll(ctx.(*SIfContext))

	return p
}

func (s *IfOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfOnlyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfOnlyContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfOnlyContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIfOnly(s)
	}
}

func (s *IfOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIfOnly(s)
	}
}

func (s *IfOnlyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIfOnly(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SIf() (localctx ISIfContext) {
	localctx = NewSIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_sIf)
	var _la int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfOnlyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(gramaticaParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.expr(0)
		}
		{
			p.SetState(122)
			p.Block()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__8 {
			{
				p.SetState(123)
				p.Match(gramaticaParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.Block()
			}

		}

	case 2:
		localctx = NewIfAnidadoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(gramaticaParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		{
			p.SetState(129)
			p.Block()
		}
		{
			p.SetState(130)
			p.Match(gramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.SIf()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISSwitchContext is an interface to support dynamic dispatch.
type ISSwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSSwitchContext differentiates from other interfaces.
	IsSSwitchContext()
}

type SSwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySSwitchContext() *SSwitchContext {
	var p = new(SSwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sSwitch
	return p
}

func InitEmptySSwitchContext(p *SSwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sSwitch
}

func (*SSwitchContext) IsSSwitchContext() {}

func NewSSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SSwitchContext {
	var p = new(SSwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sSwitch

	return p
}

func (s *SSwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SSwitchContext) CopyAll(ctx *SSwitchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	SSwitchContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySSwitchContext(&p.SSwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SSwitchContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SSwitch() (localctx ISSwitchContext) {
	localctx = NewSSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_sSwitch)
	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(gramaticaParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.expr(0)
	}
	{
		p.SetState(137)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Cases()
	}
	{
		p.SetState(139)
		p.Match(gramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) CopyAll(ctx *CasesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultContext struct {
	CasesContext
}

func NewDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultContext {
	var p = new(DefaultContext)

	InitEmptyCasesContext(&p.CasesContext)
	p.parser = parser
	p.CopyAll(ctx.(*CasesContext))

	return p
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *DefaultContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

type CaseContext struct {
	CasesContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	InitEmptyCasesContext(&p.CasesContext)
	p.parser = parser
	p.CopyAll(ctx.(*CasesContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *CaseContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *CaseContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gramaticaParserRULE_cases)
	var _la int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserT__10:
		localctx = NewCaseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(gramaticaParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.expr(0)
		}
		{
			p.SetState(143)
			p.Match(gramaticaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4323455642282640778) != 0) || _la == gramaticaParserID_VARIABLE {
			{
				p.SetState(144)
				p.Instrucciones()
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__10 || _la == gramaticaParserT__12 {
			{
				p.SetState(150)
				p.Cases()
			}

		}

	case gramaticaParserT__12:
		localctx = NewDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(gramaticaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4323455642282640778) != 0) || _la == gramaticaParserID_VARIABLE {
			{
				p.SetState(154)
				p.Instrucciones()
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockStmtContext struct {
	BlockContext
}

func NewBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmtContext {
	var p = new(BlockStmtContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_block)
	var _la int

	localctx = NewBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4323455642282640778) != 0) || _la == gramaticaParserID_VARIABLE {
		{
			p.SetState(163)
			p.Instrucciones()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(gramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISForContext is an interface to support dynamic dispatch.
type ISForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSForContext differentiates from other interfaces.
	IsSForContext()
}

type SForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySForContext() *SForContext {
	var p = new(SForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sFor
	return p
}

func InitEmptySForContext(p *SForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sFor
}

func (*SForContext) IsSForContext() {}

func NewSForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SForContext {
	var p = new(SForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sFor

	return p
}

func (s *SForContext) GetParser() antlr.Parser { return s.parser }

func (s *SForContext) CopyAll(ctx *SForContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForCondicionContext struct {
	SForContext
}

func NewForCondicionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondicionContext {
	var p = new(ForCondicionContext)

	InitEmptySForContext(&p.SForContext)
	p.parser = parser
	p.CopyAll(ctx.(*SForContext))

	return p
}

func (s *ForCondicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondicionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCondicionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCondicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForCondicion(s)
	}
}

func (s *ForCondicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForCondicion(s)
	}
}

func (s *ForCondicionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForCondicion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeContext struct {
	SForContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	InitEmptySForContext(&p.SForContext)
	p.parser = parser
	p.CopyAll(ctx.(*SForContext))

	return p
}

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *ForRangeContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *ForRangeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForAsignacionContext struct {
	SForContext
}

func NewForAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForAsignacionContext {
	var p = new(ForAsignacionContext)

	InitEmptySForContext(&p.SForContext)
	p.parser = parser
	p.CopyAll(ctx.(*SForContext))

	return p
}

func (s *ForAsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForAsignacionContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *ForAsignacionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForAsignacionContext) VarAsign() IVarAsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAsignContext)
}

func (s *ForAsignacionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForAsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForAsignacion(s)
	}
}

func (s *ForAsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForAsignacion(s)
	}
}

func (s *ForAsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SFor() (localctx ISForContext) {
	localctx = NewSForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gramaticaParserRULE_sFor)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForCondicionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Match(gramaticaParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.expr(0)
		}
		{
			p.SetState(173)
			p.Block()
		}

	case 2:
		localctx = NewForAsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(gramaticaParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.VarDcl()
		}
		{
			p.SetState(177)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.expr(0)
		}
		{
			p.SetState(179)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.VarAsign()
		}
		{
			p.SetState(181)
			p.Block()
		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.Match(gramaticaParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(gramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(gramaticaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(gramaticaParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDclContext is an interface to support dynamic dispatch.
type IVarDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDclContext differentiates from other interfaces.
	IsVarDclContext()
}

type VarDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDclContext() *VarDclContext {
	var p = new(VarDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDcl
	return p
}

func InitEmptyVarDclContext(p *VarDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDcl
}

func (*VarDclContext) IsVarDclContext() {}

func NewVarDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclContext {
	var p = new(VarDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varDcl

	return p
}

func (s *VarDclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclContext) CopyAll(ctx *VarDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarDclWithTypeAndValueContext struct {
	VarDclContext
}

func NewVarDclWithTypeAndValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDclWithTypeAndValueContext {
	var p = new(VarDclWithTypeAndValueContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VarDclWithTypeAndValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclWithTypeAndValueContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarDclWithTypeAndValueContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDclWithTypeAndValueContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDclWithTypeAndValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDclWithTypeAndValue(s)
	}
}

func (s *VarDclWithTypeAndValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDclWithTypeAndValue(s)
	}
}

func (s *VarDclWithTypeAndValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDclWithTypeAndValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDclWithInferenceContext struct {
	VarDclContext
}

func NewVarDclWithInferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDclWithInferenceContext {
	var p = new(VarDclWithInferenceContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VarDclWithInferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclWithInferenceContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarDclWithInferenceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDclWithInferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDclWithInference(s)
	}
}

func (s *VarDclWithInferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDclWithInference(s)
	}
}

func (s *VarDclWithInferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDclWithInference(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDclWithTypeOnlyContext struct {
	VarDclContext
}

func NewVarDclWithTypeOnlyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDclWithTypeOnlyContext {
	var p = new(VarDclWithTypeOnlyContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VarDclWithTypeOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclWithTypeOnlyContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarDclWithTypeOnlyContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDclWithTypeOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarDclWithTypeOnly(s)
	}
}

func (s *VarDclWithTypeOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarDclWithTypeOnly(s)
	}
}

func (s *VarDclWithTypeOnlyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarDclWithTypeOnly(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarDcl() (localctx IVarDclContext) {
	localctx = NewVarDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gramaticaParserRULE_varDcl)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDclWithTypeAndValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(gramaticaParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Type_()
		}
		{
			p.SetState(196)
			p.Match(gramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.expr(0)
		}

	case 2:
		localctx = NewVarDclWithTypeOnlyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(gramaticaParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Type_()
		}

	case 3:
		localctx = NewVarDclWithInferenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(gramaticaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDclSliceContext is an interface to support dynamic dispatch.
type IVarDclSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDclSliceContext differentiates from other interfaces.
	IsVarDclSliceContext()
}

type VarDclSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDclSliceContext() *VarDclSliceContext {
	var p = new(VarDclSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDclSlice
	return p
}

func InitEmptyVarDclSliceContext(p *VarDclSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDclSlice
}

func (*VarDclSliceContext) IsVarDclSliceContext() {}

func NewVarDclSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclSliceContext {
	var p = new(VarDclSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varDclSlice

	return p
}

func (s *VarDclSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclSliceContext) CopyAll(ctx *VarDclSliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDclSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceValoresContext struct {
	VarDclSliceContext
}

func NewSliceValoresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceValoresContext {
	var p = new(SliceValoresContext)

	InitEmptyVarDclSliceContext(&p.VarDclSliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclSliceContext))

	return p
}

func (s *SliceValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValoresContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *SliceValoresContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *SliceValoresContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SliceValoresContext) ContenidoSlice() IContenidoSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoSliceContext)
}

func (s *SliceValoresContext) AllNuevoSlice() []INuevoSliceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INuevoSliceContext); ok {
			len++
		}
	}

	tst := make([]INuevoSliceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INuevoSliceContext); ok {
			tst[i] = t.(INuevoSliceContext)
			i++
		}
	}

	return tst
}

func (s *SliceValoresContext) NuevoSlice(i int) INuevoSliceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INuevoSliceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INuevoSliceContext)
}

func (s *SliceValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceValores(s)
	}
}

func (s *SliceValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceValores(s)
	}
}

func (s *SliceValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceValores(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceVacioContext struct {
	VarDclSliceContext
}

func NewSliceVacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceVacioContext {
	var p = new(SliceVacioContext)

	InitEmptyVarDclSliceContext(&p.VarDclSliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclSliceContext))

	return p
}

func (s *SliceVacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceVacioContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *SliceVacioContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SliceVacioContext) AllNuevoSlice() []INuevoSliceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INuevoSliceContext); ok {
			len++
		}
	}

	tst := make([]INuevoSliceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INuevoSliceContext); ok {
			tst[i] = t.(INuevoSliceContext)
			i++
		}
	}

	return tst
}

func (s *SliceVacioContext) NuevoSlice(i int) INuevoSliceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INuevoSliceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INuevoSliceContext)
}

func (s *SliceVacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceVacio(s)
	}
}

func (s *SliceVacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceVacio(s)
	}
}

func (s *SliceVacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceVacio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarDclSlice() (localctx IVarDclSliceContext) {
	localctx = NewVarDclSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gramaticaParserRULE_varDclSlice)
	var _la int

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserID_VARIABLE:
		localctx = NewSliceValoresContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Assign()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__19 {
			{
				p.SetState(209)
				p.NuevoSlice()
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(214)
			p.Type_()
		}
		{
			p.SetState(215)
			p.Match(gramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.ContenidoSlice()
		}
		{
			p.SetState(217)
			p.Match(gramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case gramaticaParserT__18:
		localctx = NewSliceVacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(gramaticaParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__19 {
			{
				p.SetState(221)
				p.NuevoSlice()
			}

			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Type_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gramaticaParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gramaticaParserT__14 || _la == gramaticaParserT__17) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INuevoSliceContext is an interface to support dynamic dispatch.
type INuevoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNuevoSliceContext differentiates from other interfaces.
	IsNuevoSliceContext()
}

type NuevoSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNuevoSliceContext() *NuevoSliceContext {
	var p = new(NuevoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_nuevoSlice
	return p
}

func InitEmptyNuevoSliceContext(p *NuevoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_nuevoSlice
}

func (*NuevoSliceContext) IsNuevoSliceContext() {}

func NewNuevoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NuevoSliceContext {
	var p = new(NuevoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_nuevoSlice

	return p
}

func (s *NuevoSliceContext) GetParser() antlr.Parser { return s.parser }
func (s *NuevoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NuevoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NuevoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterNuevoSlice(s)
	}
}

func (s *NuevoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitNuevoSlice(s)
	}
}

func (s *NuevoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitNuevoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) NuevoSlice() (localctx INuevoSliceContext) {
	localctx = NewNuevoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gramaticaParserRULE_nuevoSlice)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(gramaticaParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContenidoSliceContext is an interface to support dynamic dispatch.
type IContenidoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContenidoSliceContext differentiates from other interfaces.
	IsContenidoSliceContext()
}

type ContenidoSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContenidoSliceContext() *ContenidoSliceContext {
	var p = new(ContenidoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_contenidoSlice
	return p
}

func InitEmptyContenidoSliceContext(p *ContenidoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_contenidoSlice
}

func (*ContenidoSliceContext) IsContenidoSliceContext() {}

func NewContenidoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContenidoSliceContext {
	var p = new(ContenidoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_contenidoSlice

	return p
}

func (s *ContenidoSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *ContenidoSliceContext) CopyAll(ctx *ContenidoSliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ContenidoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContenidoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceContenidoSliceContext struct {
	ContenidoSliceContext
}

func NewSliceContenidoSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceContenidoSliceContext {
	var p = new(SliceContenidoSliceContext)

	InitEmptyContenidoSliceContext(&p.ContenidoSliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContenidoSliceContext))

	return p
}

func (s *SliceContenidoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContenidoSliceContext) AllContenidoSlice() []IContenidoSliceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContenidoSliceContext); ok {
			len++
		}
	}

	tst := make([]IContenidoSliceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContenidoSliceContext); ok {
			tst[i] = t.(IContenidoSliceContext)
			i++
		}
	}

	return tst
}

func (s *SliceContenidoSliceContext) ContenidoSlice(i int) IContenidoSliceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoSliceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoSliceContext)
}

func (s *SliceContenidoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceContenidoSlice(s)
	}
}

func (s *SliceContenidoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceContenidoSlice(s)
	}
}

func (s *SliceContenidoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceContenidoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceContenidoContext struct {
	ContenidoSliceContext
}

func NewSliceContenidoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceContenidoContext {
	var p = new(SliceContenidoContext)

	InitEmptyContenidoSliceContext(&p.ContenidoSliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContenidoSliceContext))

	return p
}

func (s *SliceContenidoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContenidoContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceContenidoContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceContenidoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceContenido(s)
	}
}

func (s *SliceContenidoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceContenido(s)
	}
}

func (s *SliceContenidoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceContenido(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ContenidoSlice() (localctx IContenidoSliceContext) {
	localctx = NewContenidoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gramaticaParserRULE_contenidoSlice)
	var _la int

	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserT__29, gramaticaParserT__30, gramaticaParserT__43, gramaticaParserT__44, gramaticaParserT__45, gramaticaParserT__46, gramaticaParserT__47, gramaticaParserT__48, gramaticaParserT__49, gramaticaParserT__50, gramaticaParserT__51, gramaticaParserINT, gramaticaParserDOUBLE, gramaticaParserCHAR, gramaticaParserSTRING, gramaticaParserBOOL, gramaticaParserID_VARIABLE:
		localctx = NewSliceContenidoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.expr(0)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__3 {
			{
				p.SetState(235)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(236)
				p.expr(0)
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case gramaticaParserT__0:
		localctx = NewSliceContenidoSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(gramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.ContenidoSlice()
		}
		{
			p.SetState(244)
			p.Match(gramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__3 {
			{
				p.SetState(245)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == gramaticaParserT__0 {
				{
					p.SetState(246)
					p.Match(gramaticaParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(247)
					p.ContenidoSlice()
				}
				{
					p.SetState(248)
					p.Match(gramaticaParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDclStructContext is an interface to support dynamic dispatch.
type IVarDclStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDclStructContext differentiates from other interfaces.
	IsVarDclStructContext()
}

type VarDclStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDclStructContext() *VarDclStructContext {
	var p = new(VarDclStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDclStruct
	return p
}

func InitEmptyVarDclStructContext(p *VarDclStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varDclStruct
}

func (*VarDclStructContext) IsVarDclStructContext() {}

func NewVarDclStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclStructContext {
	var p = new(VarDclStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varDclStruct

	return p
}

func (s *VarDclStructContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclStructContext) CopyAll(ctx *VarDclStructContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDclStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclStructDataContext struct {
	VarDclStructContext
}

func NewDeclStructDataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclStructDataContext {
	var p = new(DeclStructDataContext)

	InitEmptyVarDclStructContext(&p.VarDclStructContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclStructContext))

	return p
}

func (s *DeclStructDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclStructDataContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *DeclStructDataContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *DeclStructDataContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *DeclStructDataContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclStructDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclStructData(s)
	}
}

func (s *DeclStructDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclStructData(s)
	}
}

func (s *DeclStructDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclStructData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarDclStruct() (localctx IVarDclStructContext) {
	localctx = NewVarDclStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gramaticaParserRULE_varDclStruct)
	var _la int

	localctx = NewDeclStructDataContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__20 {
		{
			p.SetState(259)
			p.Match(gramaticaParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__21 {
		{
			p.SetState(262)
			p.Match(gramaticaParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(265)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__21 {
		{
			p.SetState(266)
			p.Match(gramaticaParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(269)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gramaticaParserID_VARIABLE {
		{
			p.SetState(270)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Type_()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__5 {
			{
				p.SetState(272)
				p.Match(gramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(gramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarStructDclContext is an interface to support dynamic dispatch.
type IVarStructDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarStructDclContext differentiates from other interfaces.
	IsVarStructDclContext()
}

type VarStructDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarStructDclContext() *VarStructDclContext {
	var p = new(VarStructDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varStructDcl
	return p
}

func InitEmptyVarStructDclContext(p *VarStructDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varStructDcl
}

func (*VarStructDclContext) IsVarStructDclContext() {}

func NewVarStructDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarStructDclContext {
	var p = new(VarStructDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varStructDcl

	return p
}

func (s *VarStructDclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarStructDclContext) CopyAll(ctx *VarStructDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarStructDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarStructDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVarTypeContext struct {
	VarStructDclContext
}

func NewStructVarTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVarTypeContext {
	var p = new(StructVarTypeContext)

	InitEmptyVarStructDclContext(&p.VarStructDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarStructDclContext))

	return p
}

func (s *StructVarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVarTypeContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *StructVarTypeContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *StructVarTypeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StructVarTypeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructVarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructVarType(s)
	}
}

func (s *StructVarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructVarType(s)
	}
}

func (s *StructVarTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructVarType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructVarTypeInferenceContext struct {
	VarStructDclContext
}

func NewStructVarTypeInferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVarTypeInferenceContext {
	var p = new(StructVarTypeInferenceContext)

	InitEmptyVarStructDclContext(&p.VarStructDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarStructDclContext))

	return p
}

func (s *StructVarTypeInferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVarTypeInferenceContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *StructVarTypeInferenceContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *StructVarTypeInferenceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StructVarTypeInferenceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructVarTypeInferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructVarTypeInference(s)
	}
}

func (s *StructVarTypeInferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructVarTypeInference(s)
	}
}

func (s *StructVarTypeInferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructVarTypeInference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarStructDcl() (localctx IVarStructDclContext) {
	localctx = NewVarStructDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gramaticaParserRULE_varStructDcl)
	var _la int

	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructVarTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(gramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(gramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(gramaticaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.expr(0)
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__3 {
			{
				p.SetState(288)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(289)
				p.Match(gramaticaParserID_VARIABLE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(290)
				p.Match(gramaticaParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(291)
				p.expr(0)
			}

			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(296)
			p.Match(gramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStructVarTypeInferenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(gramaticaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(gramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(gramaticaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.expr(0)
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__3 {
			{
				p.SetState(305)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == gramaticaParserID_VARIABLE {
				{
					p.SetState(306)
					p.Match(gramaticaParserID_VARIABLE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(307)
					p.Match(gramaticaParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(308)
					p.expr(0)
				}

			}

			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(315)
			p.Match(gramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarAsignContext is an interface to support dynamic dispatch.
type IVarAsignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarAsignContext differentiates from other interfaces.
	IsVarAsignContext()
}

type VarAsignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAsignContext() *VarAsignContext {
	var p = new(VarAsignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varAsign
	return p
}

func InitEmptyVarAsignContext(p *VarAsignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varAsign
}

func (*VarAsignContext) IsVarAsignContext() {}

func NewVarAsignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAsignContext {
	var p = new(VarAsignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varAsign

	return p
}

func (s *VarAsignContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAsignContext) CopyAll(ctx *VarAsignContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarAsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAsignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarIncContext struct {
	VarAsignContext
	op antlr.Token
}

func NewVarIncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarIncContext {
	var p = new(VarIncContext)

	InitEmptyVarAsignContext(&p.VarAsignContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAsignContext))

	return p
}

func (s *VarIncContext) GetOp() antlr.Token { return s.op }

func (s *VarIncContext) SetOp(v antlr.Token) { s.op = v }

func (s *VarIncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarIncContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarIncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarInc(s)
	}
}

func (s *VarIncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarInc(s)
	}
}

func (s *VarIncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarInc(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExprContext struct {
	VarAsignContext
}

func NewVarExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprContext {
	var p = new(VarExprContext)

	InitEmptyVarAsignContext(&p.VarAsignContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAsignContext))

	return p
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayAccessContext struct {
	VarAsignContext
}

func NewArrayAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessContext {
	var p = new(ArrayAccessContext)

	InitEmptyVarAsignContext(&p.VarAsignContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAsignContext))

	return p
}

func (s *ArrayAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayAccessContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayAccessContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayAccess(s)
	}
}

func (s *ArrayAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayAccess(s)
	}
}

func (s *ArrayAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAccessAsignContext struct {
	VarAsignContext
}

func NewStructAccessAsignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAccessAsignContext {
	var p = new(StructAccessAsignContext)

	InitEmptyVarAsignContext(&p.VarAsignContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAsignContext))

	return p
}

func (s *StructAccessAsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessAsignContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *StructAccessAsignContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *StructAccessAsignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructAccessAsignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructAccessAsign(s)
	}
}

func (s *StructAccessAsignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructAccessAsign(s)
	}
}

func (s *StructAccessAsignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructAccessAsign(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarAddContext struct {
	VarAsignContext
	op antlr.Token
}

func NewVarAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarAddContext {
	var p = new(VarAddContext)

	InitEmptyVarAsignContext(&p.VarAsignContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAsignContext))

	return p
}

func (s *VarAddContext) GetOp() antlr.Token { return s.op }

func (s *VarAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *VarAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAddContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *VarAddContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarAdd(s)
	}
}

func (s *VarAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarAdd(s)
	}
}

func (s *VarAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitVarAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarAsign() (localctx IVarAsignContext) {
	localctx = NewVarAsignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gramaticaParserRULE_varAsign)
	var _la int

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(gramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.expr(0)
		}

	case 2:
		localctx = NewVarAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VarAddContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == gramaticaParserT__22 || _la == gramaticaParserT__23) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VarAddContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(324)
			p.expr(0)
		}

	case 3:
		localctx = NewVarIncContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VarIncContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == gramaticaParserT__24 || _la == gramaticaParserT__25) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VarIncContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		localctx = NewArrayAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(327)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__26 {
			{
				p.SetState(328)
				p.Match(gramaticaParserT__26)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(329)
				p.expr(0)
			}
			{
				p.SetState(330)
				p.Match(gramaticaParserT__27)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(334)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(336)
			p.Match(gramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.expr(0)
		}

	case 5:
		localctx = NewStructAccessAsignContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(339)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == gramaticaParserT__28 {
			{
				p.SetState(340)
				p.Match(gramaticaParserT__28)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(341)
				p.Match(gramaticaParserID_VARIABLE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(346)
			p.Match(gramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionValueContext struct {
	ExprContext
}

func NewCallFunctionValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionValueContext {
	var p = new(CallFunctionValueContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallFunctionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionValueContext) VarCallStatement() IVarCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCallStatementContext)
}

func (s *CallFunctionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunctionValue(s)
	}
}

func (s *CallFunctionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunctionValue(s)
	}
}

func (s *CallFunctionValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunctionValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalContext struct {
	ExprContext
	op antlr.Token
}

func NewLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalContext {
	var p = new(LogicalContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalContext) GetOp() antlr.Token { return s.op }

func (s *LogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (s *LogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAccessContext struct {
	ExprContext
}

func NewStructAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAccessContext {
	var p = new(StructAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *StructAccessContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *StructAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructAccess(s)
	}
}

func (s *StructAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructAccess(s)
	}
}

func (s *StructAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharContext struct {
	ExprContext
}

func NewCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharContext {
	var p = new(CharContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharContext) CHAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCHAR, 0)
}

func (s *CharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterChar(s)
	}
}

func (s *CharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitChar(s)
	}
}

func (s *CharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitChar(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExprContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserBOOL, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionStructValueContext struct {
	ExprContext
}

func NewCallFunctionStructValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionStructValueContext {
	var p = new(CallFunctionStructValueContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallFunctionStructValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionStructValueContext) VarCallFuncStruct() IVarCallFuncStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCallFuncStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCallFuncStructContext)
}

func (s *CallFunctionStructValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunctionStructValue(s)
	}
}

func (s *CallFunctionStructValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunctionStructValue(s)
	}
}

func (s *CallFunctionStructValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunctionStructValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayFindIndexContext struct {
	ExprContext
}

func NewArrayFindIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayFindIndexContext {
	var p = new(ArrayFindIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayFindIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayFindIndexContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayFindIndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayFindIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayFindIndex(s)
	}
}

func (s *ArrayFindIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayFindIndex(s)
	}
}

func (s *ArrayFindIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayFindIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayAppendContext struct {
	ExprContext
}

func NewArrayAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAppendContext {
	var p = new(ArrayAppendContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAppendContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayAppendContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayAppend(s)
	}
}

func (s *ArrayAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayAppend(s)
	}
}

func (s *ArrayAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualsNotEqualsContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualsNotEqualsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualsNotEqualsContext {
	var p = new(EqualsNotEqualsContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualsNotEqualsContext) GetOp() antlr.Token { return s.op }

func (s *EqualsNotEqualsContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualsNotEqualsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsNotEqualsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualsNotEqualsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualsNotEqualsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterEqualsNotEquals(s)
	}
}

func (s *EqualsNotEqualsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitEqualsNotEquals(s)
	}
}

func (s *EqualsNotEqualsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitEqualsNotEquals(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntToStringContext struct {
	ExprContext
}

func NewIntToStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntToStringContext {
	var p = new(IntToStringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntToStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntToStringContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IntToStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIntToString(s)
	}
}

func (s *IntToStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIntToString(s)
	}
}

func (s *IntToStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIntToString(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayAccessSimpleContext struct {
	ExprContext
}

func NewArrayAccessSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessSimpleContext {
	var p = new(ArrayAccessSimpleContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayAccessSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessSimpleContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayAccessSimpleContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayAccessSimpleContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayAccessSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayAccessSimple(s)
	}
}

func (s *ArrayAccessSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayAccessSimple(s)
	}
}

func (s *ArrayAccessSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayAccessSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLengthContext struct {
	ExprContext
}

func NewArrayLengthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLengthContext {
	var p = new(ArrayLengthContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLengthContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayLengthContext) AllPosicion() []IPosicionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPosicionContext); ok {
			len++
		}
	}

	tst := make([]IPosicionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPosicionContext); ok {
			tst[i] = t.(IPosicionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLengthContext) Posicion(i int) IPosicionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPosicionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPosicionContext)
}

func (s *ArrayLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayLength(s)
	}
}

func (s *ArrayLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayLength(s)
	}
}

func (s *ArrayLengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayLength(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModuloContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivModuloContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModuloContext {
	var p = new(MulDivModuloContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivModuloContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModuloContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModuloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModuloContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModuloContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivModuloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterMulDivModulo(s)
	}
}

func (s *MulDivModuloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitMulDivModulo(s)
	}
}

func (s *MulDivModuloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitMulDivModulo(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	ExprContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOUBLE, 0)
}

func (s *DoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDouble(s)
	}
}

func (s *DoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDouble(s)
	}
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerContext struct {
	ExprContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINT, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilContext struct {
	ExprContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinorMajorEqualContext struct {
	ExprContext
	op antlr.Token
}

func NewMinorMajorEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinorMajorEqualContext {
	var p = new(MinorMajorEqualContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MinorMajorEqualContext) GetOp() antlr.Token { return s.op }

func (s *MinorMajorEqualContext) SetOp(v antlr.Token) { s.op = v }

func (s *MinorMajorEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinorMajorEqualContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MinorMajorEqualContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinorMajorEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterMinorMajorEqual(s)
	}
}

func (s *MinorMajorEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitMinorMajorEqual(s)
	}
}

func (s *MinorMajorEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitMinorMajorEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContext struct {
	ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReflectTypeContext struct {
	ExprContext
}

func NewReflectTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReflectTypeContext {
	var p = new(ReflectTypeContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ReflectTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReflectTypeContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReflectTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterReflectType(s)
	}
}

func (s *ReflectTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitReflectType(s)
	}
}

func (s *ReflectTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitReflectType(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateContext struct {
	ExprContext
}

func NewNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateContext {
	var p = new(NegateContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterNegate(s)
	}
}

func (s *NegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitNegate(s)
	}
}

func (s *NegateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitNegate(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayJoinContext struct {
	ExprContext
}

func NewArrayJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayJoinContext {
	var p = new(ArrayJoinContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayJoinContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *ArrayJoinContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArrayJoin(s)
	}
}

func (s *ArrayJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArrayJoin(s)
	}
}

func (s *ArrayJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitArrayJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatToStringContext struct {
	ExprContext
}

func NewFloatToStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatToStringContext {
	var p = new(FloatToStringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatToStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatToStringContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloatToStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFloatToString(s)
	}
}

func (s *FloatToStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFloatToString(s)
	}
}

func (s *FloatToStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFloatToString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *gramaticaParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, gramaticaParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(351)
			p.Match(gramaticaParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.expr(26)
		}

	case 2:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(353)
			p.Match(gramaticaParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expr(25)
		}

	case 3:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(355)
			p.Match(gramaticaParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewDoubleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(356)
			p.Match(gramaticaParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(357)
			p.Match(gramaticaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(358)
			p.Match(gramaticaParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(359)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewCharContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(360)
			p.Match(gramaticaParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(361)
			p.Match(gramaticaParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(362)
			p.Match(gramaticaParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.expr(0)
		}
		{
			p.SetState(364)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewArrayAccessSimpleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(366)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(367)
					p.Match(gramaticaParserT__26)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(368)
					p.expr(0)
				}
				{
					p.SetState(369)
					p.Match(gramaticaParserT__27)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 12:
		localctx = NewArrayFindIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(375)
			p.Match(gramaticaParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(gramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.expr(0)
		}
		{
			p.SetState(379)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewArrayJoinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(381)
			p.Match(gramaticaParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(gramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.expr(0)
		}
		{
			p.SetState(385)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewArrayLengthContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(387)
			p.Match(gramaticaParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__26 {
			{
				p.SetState(389)
				p.Posicion()
			}

			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(395)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewArrayAppendContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(396)
			p.Match(gramaticaParserT__48)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(gramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expr(0)
		}
		{
			p.SetState(400)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewIntToStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(402)
			p.Match(gramaticaParserT__49)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.expr(0)
		}
		{
			p.SetState(404)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewFloatToStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(406)
			p.Match(gramaticaParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.expr(0)
		}
		{
			p.SetState(408)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		localctx = NewReflectTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(410)
			p.Match(gramaticaParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.expr(0)
		}
		{
			p.SetState(412)
			p.Match(gramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		localctx = NewStructAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(415)
					p.Match(gramaticaParserT__28)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(416)
					p.Match(gramaticaParserID_VARIABLE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(419)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(422)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(421)
				p.Match(gramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 20:
		localctx = NewCallFunctionValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(424)
			p.VarCallStatement()
		}

	case 21:
		localctx = NewCallFunctionStructValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(425)
			p.VarCallFuncStruct()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(443)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModuloContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expr)
				p.SetState(428)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(429)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModuloContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModuloContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(430)
					p.expr(25)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expr)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(432)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserT__29 || _la == gramaticaParserT__34) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(433)
					p.expr(24)
				}

			case 3:
				localctx = NewMinorMajorEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expr)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(435)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MinorMajorEqualContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030792151040) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MinorMajorEqualContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(436)
					p.expr(23)
				}

			case 4:
				localctx = NewEqualsNotEqualsContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expr)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(438)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualsNotEqualsContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserT__39 || _la == gramaticaParserT__40) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualsNotEqualsContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(439)
					p.expr(22)
				}

			case 5:
				localctx = NewLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expr)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(441)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserT__41 || _la == gramaticaParserT__42) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(442)
					p.expr(21)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPosicionContext is an interface to support dynamic dispatch.
type IPosicionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPosicionContext differentiates from other interfaces.
	IsPosicionContext()
}

type PosicionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPosicionContext() *PosicionContext {
	var p = new(PosicionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_posicion
	return p
}

func InitEmptyPosicionContext(p *PosicionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_posicion
}

func (*PosicionContext) IsPosicionContext() {}

func NewPosicionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PosicionContext {
	var p = new(PosicionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_posicion

	return p
}

func (s *PosicionContext) GetParser() antlr.Parser { return s.parser }

func (s *PosicionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PosicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosicionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PosicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPosicion(s)
	}
}

func (s *PosicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPosicion(s)
	}
}

func (s *PosicionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPosicion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Posicion() (localctx IPosicionContext) {
	localctx = NewPosicionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gramaticaParserRULE_posicion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(gramaticaParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.expr(0)
	}
	{
		p.SetState(450)
		p.Match(gramaticaParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID_VARIABLE() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gramaticaParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-22)) & ^0x3f) == 0 && ((int64(1)<<(_la-22))&70435316170753) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakContext is an interface to support dynamic dispatch.
type IBreakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakContext differentiates from other interfaces.
	IsBreakContext()
}

type BreakContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakContext() *BreakContext {
	var p = new(BreakContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_break
	return p
}

func InitEmptyBreakContext(p *BreakContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_break
}

func (*BreakContext) IsBreakContext() {}

func NewBreakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakContext {
	var p = new(BreakContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_break

	return p
}

func (s *BreakContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBreak(s)
	}
}

func (s *BreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Break_() (localctx IBreakContext) {
	localctx = NewBreakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gramaticaParserRULE_break)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(gramaticaParserT__57)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__5 {
		{
			p.SetState(455)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueContext is an interface to support dynamic dispatch.
type IContinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueContext differentiates from other interfaces.
	IsContinueContext()
}

type ContinueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueContext() *ContinueContext {
	var p = new(ContinueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_continue
	return p
}

func InitEmptyContinueContext(p *ContinueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_continue
}

func (*ContinueContext) IsContinueContext() {}

func NewContinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueContext {
	var p = new(ContinueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_continue

	return p
}

func (s *ContinueContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitContinue(s)
	}
}

func (s *ContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Continue_() (localctx IContinueContext) {
	localctx = NewContinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gramaticaParserRULE_continue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(gramaticaParserT__58)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__5 {
		{
			p.SetState(459)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_functions
	return p
}

func InitEmptyFunctionsContext(p *FunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_functions
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) CopyAll(ctx *FunctionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionesContext struct {
	FunctionsContext
}

func NewFuncionesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionesContext {
	var p = new(FuncionesContext)

	InitEmptyFunctionsContext(&p.FunctionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionsContext))

	return p
}

func (s *FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *FuncionesContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *FuncionesContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncionesContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *FuncionesContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncionesContext) ValRet() IValRetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValRetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValRetContext)
}

func (s *FuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFunciones(s)
	}
}

func (s *FuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFunciones(s)
	}
}

func (s *FuncionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFunciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gramaticaParserRULE_functions)
	var _la int

	localctx = NewFuncionesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(gramaticaParserT__59)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.Match(gramaticaParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserID_VARIABLE {
		{
			p.SetState(465)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.Type_()
		}
		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__3 {
			{
				p.SetState(467)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(468)
				p.Match(gramaticaParserID_VARIABLE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(469)
				p.Type_()
			}

			p.SetState(474)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(477)
		p.Match(gramaticaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-22)) & ^0x3f) == 0 && ((int64(1)<<(_la-22))&70435316170753) != 0 {
		{
			p.SetState(478)
			p.ValRet()
		}

	}
	{
		p.SetState(481)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionStructContext is an interface to support dynamic dispatch.
type IFunctionStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionStructContext differentiates from other interfaces.
	IsFunctionStructContext()
}

type FunctionStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStructContext() *FunctionStructContext {
	var p = new(FunctionStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_functionStruct
	return p
}

func InitEmptyFunctionStructContext(p *FunctionStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_functionStruct
}

func (*FunctionStructContext) IsFunctionStructContext() {}

func NewFunctionStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStructContext {
	var p = new(FunctionStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_functionStruct

	return p
}

func (s *FunctionStructContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStructContext) CopyAll(ctx *FunctionStructContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionesStructsNativasContext struct {
	FunctionStructContext
}

func NewFuncionesStructsNativasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionesStructsNativasContext {
	var p = new(FuncionesStructsNativasContext)

	InitEmptyFunctionStructContext(&p.FunctionStructContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionStructContext))

	return p
}

func (s *FuncionesStructsNativasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesStructsNativasContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *FuncionesStructsNativasContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *FuncionesStructsNativasContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncionesStructsNativasContext) DefParams() IDefParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefParamsContext)
}

func (s *FuncionesStructsNativasContext) ValRet() IValRetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValRetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValRetContext)
}

func (s *FuncionesStructsNativasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncionesStructsNativas(s)
	}
}

func (s *FuncionesStructsNativasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncionesStructsNativas(s)
	}
}

func (s *FuncionesStructsNativasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFuncionesStructsNativas(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FunctionStruct() (localctx IFunctionStructContext) {
	localctx = NewFunctionStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gramaticaParserRULE_functionStruct)
	var _la int

	localctx = NewFuncionesStructsNativasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(gramaticaParserT__59)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
		p.Match(gramaticaParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(485)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.Match(gramaticaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Match(gramaticaParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserID_VARIABLE {
		{
			p.SetState(490)
			p.DefParams()
		}

	}
	{
		p.SetState(493)
		p.Match(gramaticaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-22)) & ^0x3f) == 0 && ((int64(1)<<(_la-22))&70435316170753) != 0 {
		{
			p.SetState(494)
			p.ValRet()
		}

	}
	{
		p.SetState(497)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefParamsContext is an interface to support dynamic dispatch.
type IDefParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID_VARIABLE() []antlr.TerminalNode
	ID_VARIABLE(i int) antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext

	// IsDefParamsContext differentiates from other interfaces.
	IsDefParamsContext()
}

type DefParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefParamsContext() *DefParamsContext {
	var p = new(DefParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_defParams
	return p
}

func InitEmptyDefParamsContext(p *DefParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_defParams
}

func (*DefParamsContext) IsDefParamsContext() {}

func NewDefParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefParamsContext {
	var p = new(DefParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_defParams

	return p
}

func (s *DefParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefParamsContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *DefParamsContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *DefParamsContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *DefParamsContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DefParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDefParams(s)
	}
}

func (s *DefParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDefParams(s)
	}
}

func (s *DefParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDefParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DefParams() (localctx IDefParamsContext) {
	localctx = NewDefParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gramaticaParserRULE_defParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.Type_()
	}
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserT__3 {
		{
			p.SetState(501)
			p.Match(gramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(gramaticaParserID_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.Type_()
		}

		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarCallStatementContext is an interface to support dynamic dispatch.
type IVarCallStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarCallStatementContext differentiates from other interfaces.
	IsVarCallStatementContext()
}

type VarCallStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarCallStatementContext() *VarCallStatementContext {
	var p = new(VarCallStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varCallStatement
	return p
}

func InitEmptyVarCallStatementContext(p *VarCallStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varCallStatement
}

func (*VarCallStatementContext) IsVarCallStatementContext() {}

func NewVarCallStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCallStatementContext {
	var p = new(VarCallStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varCallStatement

	return p
}

func (s *VarCallStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VarCallStatementContext) CopyAll(ctx *VarCallStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCallStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallFunctionContext struct {
	VarCallStatementContext
}

func NewCallFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionContext {
	var p = new(CallFunctionContext)

	InitEmptyVarCallStatementContext(&p.VarCallStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarCallStatementContext))

	return p
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ID_VARIABLE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, 0)
}

func (s *CallFunctionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallFunctionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunction(s)
	}
}

func (s *CallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunction(s)
	}
}

func (s *CallFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarCallStatement() (localctx IVarCallStatementContext) {
	localctx = NewVarCallStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gramaticaParserRULE_varCallStatement)
	var _la int

	localctx = NewCallFunctionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)
		p.Match(gramaticaParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&408030265347) != 0 {
		{
			p.SetState(511)
			p.expr(0)
		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__3 {
			{
				p.SetState(512)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(513)
				p.expr(0)
			}

			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(521)
		p.Match(gramaticaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(522)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarCallFuncStructContext is an interface to support dynamic dispatch.
type IVarCallFuncStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarCallFuncStructContext differentiates from other interfaces.
	IsVarCallFuncStructContext()
}

type VarCallFuncStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarCallFuncStructContext() *VarCallFuncStructContext {
	var p = new(VarCallFuncStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varCallFuncStruct
	return p
}

func InitEmptyVarCallFuncStructContext(p *VarCallFuncStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varCallFuncStruct
}

func (*VarCallFuncStructContext) IsVarCallFuncStructContext() {}

func NewVarCallFuncStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCallFuncStructContext {
	var p = new(VarCallFuncStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varCallFuncStruct

	return p
}

func (s *VarCallFuncStructContext) GetParser() antlr.Parser { return s.parser }

func (s *VarCallFuncStructContext) CopyAll(ctx *VarCallFuncStructContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarCallFuncStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCallFuncStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallFunctionStructContext struct {
	VarCallFuncStructContext
}

func NewCallFunctionStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionStructContext {
	var p = new(CallFunctionStructContext)

	InitEmptyVarCallFuncStructContext(&p.VarCallFuncStructContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarCallFuncStructContext))

	return p
}

func (s *CallFunctionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionStructContext) AllID_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserID_VARIABLE)
}

func (s *CallFunctionStructContext) ID_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserID_VARIABLE, i)
}

func (s *CallFunctionStructContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallFunctionStructContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallFunctionStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCallFunctionStruct(s)
	}
}

func (s *CallFunctionStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCallFunctionStruct(s)
	}
}

func (s *CallFunctionStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCallFunctionStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) VarCallFuncStruct() (localctx IVarCallFuncStructContext) {
	localctx = NewVarCallFuncStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gramaticaParserRULE_varCallFuncStruct)
	var _la int

	localctx = NewCallFunctionStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(525)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(526)
		p.Match(gramaticaParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Match(gramaticaParserID_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(gramaticaParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&408030265347) != 0 {
		{
			p.SetState(529)
			p.expr(0)
		}
		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserT__3 {
			{
				p.SetState(530)
				p.Match(gramaticaParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(531)
				p.expr(0)
			}

			p.SetState(536)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(539)
		p.Match(gramaticaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(541)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(540)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValRetContext is an interface to support dynamic dispatch.
type IValRetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsValRetContext differentiates from other interfaces.
	IsValRetContext()
}

type ValRetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValRetContext() *ValRetContext {
	var p = new(ValRetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_valRet
	return p
}

func InitEmptyValRetContext(p *ValRetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_valRet
}

func (*ValRetContext) IsValRetContext() {}

func NewValRetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValRetContext {
	var p = new(ValRetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_valRet

	return p
}

func (s *ValRetContext) GetParser() antlr.Parser { return s.parser }

func (s *ValRetContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ValRetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValRetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValRetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterValRet(s)
	}
}

func (s *ValRetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitValRet(s)
	}
}

func (s *ValRetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitValRet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ValRet() (localctx IValRetContext) {
	localctx = NewValRetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gramaticaParserRULE_valRet)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(543)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetornoContext is an interface to support dynamic dispatch.
type IRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsRetornoContext differentiates from other interfaces.
	IsRetornoContext()
}

type RetornoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetornoContext() *RetornoContext {
	var p = new(RetornoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_retorno
	return p
}

func InitEmptyRetornoContext(p *RetornoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_retorno
}

func (*RetornoContext) IsRetornoContext() {}

func NewRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornoContext {
	var p = new(RetornoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_retorno

	return p
}

func (s *RetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterRetorno(s)
	}
}

func (s *RetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitRetorno(s)
	}
}

func (s *RetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Retorno() (localctx IRetornoContext) {
	localctx = NewRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gramaticaParserRULE_retorno)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		p.Match(gramaticaParserT__60)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(546)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserT__5 {
		{
			p.SetState(549)
			p.Match(gramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gramaticaParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 20)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
