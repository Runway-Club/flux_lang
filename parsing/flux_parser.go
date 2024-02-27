// Code generated from Flux.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Flux
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

type Flux struct {
	*antlr.BaseParser
}

var FluxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fluxParserInit() {
	staticData := &FluxParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'bool'", "", "'num'", "'na'", "", "", "", "'ipv4'",
		"'loop'", "'if'", "'else'", "'fun'", "'return'", "'->'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'^'", "'='", "'!='", "'>'", "'<'", "'>='", "'<='",
		"'and'", "'or'", "'not'", "'xor'", "'{'", "'}'", "'('", "')'", "'['",
		"']'", "','", "'.'", "", "", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "LOOP", "IF", "ELSE",
		"FUNC", "RETURN", "RETURN_TYPE", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV",
		"OP_MOD", "OP_POW", "OP_EQ", "OP_NEQ", "OP_GT", "OP_LT", "OP_GTE", "OP_LTE",
		"OP_AND", "OP_OR", "OP_NOT", "OP_XOR", "L_BLOCK", "R_BLOCK", "L_PAREN",
		"R_PAREN", "L_SQUARE", "R_SQUARE", "COMMA", "DOT", "VAR_IDENTIFIER",
		"COMMON_IDENTIFIER", "AT", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expression", "block", "loop_statement", "if_statement",
		"return_statement", "data_type", "func_declaration", "var_type", "var_name",
		"var_value", "string_var_declaration", "number_var_declaration", "boolean_var_declaration",
		"single_var_declaration", "array_var_declaration", "var_assignment",
		"op_level1", "op_level2", "op_level3", "op_level4", "op_level5", "numeric_expression",
		"logical_expression", "comparative_expression", "math_expression", "get_array_element",
		"get_child", "get_var", "function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 496, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 77, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 84, 8,
		2, 1, 3, 1, 3, 5, 3, 88, 8, 3, 10, 3, 12, 3, 91, 9, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		107, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 122, 8, 8, 10, 8, 12, 8, 125, 9, 8, 3, 8, 127,
		8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 132, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 146, 8, 12, 10,
		12, 12, 12, 149, 9, 12, 1, 12, 1, 12, 5, 12, 153, 8, 12, 10, 12, 12, 12,
		156, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 164, 8, 13,
		10, 13, 12, 13, 167, 9, 13, 1, 13, 1, 13, 5, 13, 171, 8, 13, 10, 13, 12,
		13, 174, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 182, 8,
		13, 10, 13, 12, 13, 185, 9, 13, 1, 13, 1, 13, 5, 13, 189, 8, 13, 10, 13,
		12, 13, 192, 9, 13, 1, 13, 1, 13, 3, 13, 196, 8, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 5, 14, 202, 8, 14, 10, 14, 12, 14, 205, 9, 14, 1, 14, 1, 14, 5,
		14, 209, 8, 14, 10, 14, 12, 14, 212, 9, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 220, 8, 14, 10, 14, 12, 14, 223, 9, 14, 1, 14, 1,
		14, 5, 14, 227, 8, 14, 10, 14, 12, 14, 230, 9, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 238, 8, 14, 10, 14, 12, 14, 241, 9, 14, 1,
		14, 1, 14, 5, 14, 245, 8, 14, 10, 14, 12, 14, 248, 9, 14, 1, 14, 1, 14,
		3, 14, 252, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5,
		15, 261, 8, 15, 10, 15, 12, 15, 264, 9, 15, 1, 15, 1, 15, 3, 15, 268, 8,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 274, 8, 16, 10, 16, 12, 16, 277,
		9, 16, 1, 16, 1, 16, 1, 16, 5, 16, 282, 8, 16, 10, 16, 12, 16, 285, 9,
		16, 1, 16, 5, 16, 288, 8, 16, 10, 16, 12, 16, 291, 9, 16, 1, 16, 5, 16,
		294, 8, 16, 10, 16, 12, 16, 297, 9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 305, 8, 16, 10, 16, 12, 16, 308, 9, 16, 1, 16, 1, 16,
		1, 16, 5, 16, 313, 8, 16, 10, 16, 12, 16, 316, 9, 16, 1, 16, 5, 16, 319,
		8, 16, 10, 16, 12, 16, 322, 9, 16, 1, 16, 1, 16, 1, 16, 3, 16, 327, 8,
		16, 1, 17, 1, 17, 1, 17, 5, 17, 332, 8, 17, 10, 17, 12, 17, 335, 9, 17,
		1, 17, 1, 17, 5, 17, 339, 8, 17, 10, 17, 12, 17, 342, 9, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 5, 17, 349, 8, 17, 10, 17, 12, 17, 352, 9, 17,
		1, 17, 1, 17, 5, 17, 356, 8, 17, 10, 17, 12, 17, 359, 9, 17, 1, 17, 1,
		17, 3, 17, 363, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 385, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 5, 23, 395, 8, 23, 10, 23, 12, 23, 398, 9, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24,
		410, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5,
		24, 420, 8, 24, 10, 24, 12, 24, 423, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3,
		25, 439, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 445, 8, 25, 10, 25,
		12, 25, 448, 9, 25, 1, 26, 1, 26, 1, 26, 3, 26, 453, 8, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 475, 8, 28,
		1, 29, 1, 29, 1, 29, 3, 29, 480, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 5, 30, 487, 8, 30, 10, 30, 12, 30, 490, 9, 30, 3, 30, 492, 8, 30, 1,
		30, 1, 30, 1, 30, 0, 3, 46, 48, 50, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 0, 7, 4, 0, 2, 2, 4, 4, 6, 6, 43, 43, 4, 0, 2, 2, 4, 4,
		6, 6, 11, 11, 4, 0, 1, 1, 3, 3, 5, 5, 10, 10, 1, 0, 20, 23, 1, 0, 18, 19,
		2, 0, 30, 31, 33, 33, 1, 0, 24, 29, 537, 0, 65, 1, 0, 0, 0, 2, 76, 1, 0,
		0, 0, 4, 83, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 94, 1, 0, 0, 0, 10, 99,
		1, 0, 0, 0, 12, 108, 1, 0, 0, 0, 14, 112, 1, 0, 0, 0, 16, 114, 1, 0, 0,
		0, 18, 135, 1, 0, 0, 0, 20, 137, 1, 0, 0, 0, 22, 139, 1, 0, 0, 0, 24, 141,
		1, 0, 0, 0, 26, 195, 1, 0, 0, 0, 28, 251, 1, 0, 0, 0, 30, 267, 1, 0, 0,
		0, 32, 326, 1, 0, 0, 0, 34, 362, 1, 0, 0, 0, 36, 364, 1, 0, 0, 0, 38, 366,
		1, 0, 0, 0, 40, 368, 1, 0, 0, 0, 42, 370, 1, 0, 0, 0, 44, 372, 1, 0, 0,
		0, 46, 384, 1, 0, 0, 0, 48, 409, 1, 0, 0, 0, 50, 438, 1, 0, 0, 0, 52, 452,
		1, 0, 0, 0, 54, 454, 1, 0, 0, 0, 56, 474, 1, 0, 0, 0, 58, 479, 1, 0, 0,
		0, 60, 481, 1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0,
		67, 65, 1, 0, 0, 0, 68, 69, 5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 77, 3, 4,
		2, 0, 71, 77, 3, 8, 4, 0, 72, 77, 3, 10, 5, 0, 73, 77, 3, 16, 8, 0, 74,
		77, 3, 12, 6, 0, 75, 77, 5, 45, 0, 0, 76, 70, 1, 0, 0, 0, 76, 71, 1, 0,
		0, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 75,
		1, 0, 0, 0, 77, 3, 1, 0, 0, 0, 78, 84, 3, 60, 30, 0, 79, 84, 3, 34, 17,
		0, 80, 84, 3, 30, 15, 0, 81, 84, 3, 32, 16, 0, 82, 84, 3, 58, 29, 0, 83,
		78, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 83, 81, 1, 0, 0,
		0, 83, 82, 1, 0, 0, 0, 84, 5, 1, 0, 0, 0, 85, 89, 5, 34, 0, 0, 86, 88,
		3, 2, 1, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5,
		35, 0, 0, 93, 7, 1, 0, 0, 0, 94, 95, 5, 44, 0, 0, 95, 96, 5, 12, 0, 0,
		96, 97, 3, 50, 25, 0, 97, 98, 3, 6, 3, 0, 98, 9, 1, 0, 0, 0, 99, 100, 5,
		44, 0, 0, 100, 101, 5, 13, 0, 0, 101, 102, 3, 50, 25, 0, 102, 106, 3, 6,
		3, 0, 103, 104, 5, 44, 0, 0, 104, 105, 5, 14, 0, 0, 105, 107, 3, 6, 3,
		0, 106, 103, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 11, 1, 0, 0, 0, 108,
		109, 5, 44, 0, 0, 109, 110, 5, 16, 0, 0, 110, 111, 3, 4, 2, 0, 111, 13,
		1, 0, 0, 0, 112, 113, 7, 0, 0, 0, 113, 15, 1, 0, 0, 0, 114, 115, 5, 44,
		0, 0, 115, 116, 5, 15, 0, 0, 116, 117, 5, 42, 0, 0, 117, 126, 5, 36, 0,
		0, 118, 123, 5, 42, 0, 0, 119, 120, 5, 40, 0, 0, 120, 122, 5, 42, 0, 0,
		121, 119, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 118,
		1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 131, 5, 37,
		0, 0, 129, 130, 5, 17, 0, 0, 130, 132, 3, 14, 7, 0, 131, 129, 1, 0, 0,
		0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 3, 6, 3, 0, 134,
		17, 1, 0, 0, 0, 135, 136, 7, 1, 0, 0, 136, 19, 1, 0, 0, 0, 137, 138, 5,
		42, 0, 0, 138, 21, 1, 0, 0, 0, 139, 140, 7, 2, 0, 0, 140, 23, 1, 0, 0,
		0, 141, 142, 5, 2, 0, 0, 142, 143, 3, 20, 10, 0, 143, 147, 5, 34, 0, 0,
		144, 146, 5, 45, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 150, 154, 5, 1, 0, 0, 151, 153, 5, 45, 0, 0, 152, 151, 1, 0,
		0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 35, 0, 0, 158,
		25, 1, 0, 0, 0, 159, 160, 5, 6, 0, 0, 160, 161, 3, 20, 10, 0, 161, 165,
		5, 34, 0, 0, 162, 164, 5, 45, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1,
		0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0,
		0, 167, 165, 1, 0, 0, 0, 168, 172, 5, 5, 0, 0, 169, 171, 5, 45, 0, 0, 170,
		169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 5, 35,
		0, 0, 176, 196, 1, 0, 0, 0, 177, 178, 5, 6, 0, 0, 178, 179, 3, 20, 10,
		0, 179, 183, 5, 34, 0, 0, 180, 182, 5, 45, 0, 0, 181, 180, 1, 0, 0, 0,
		182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 190, 3, 52, 26, 0, 187, 189,
		5, 45, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0,
		193, 194, 5, 35, 0, 0, 194, 196, 1, 0, 0, 0, 195, 159, 1, 0, 0, 0, 195,
		177, 1, 0, 0, 0, 196, 27, 1, 0, 0, 0, 197, 198, 5, 4, 0, 0, 198, 199, 3,
		20, 10, 0, 199, 203, 5, 34, 0, 0, 200, 202, 5, 45, 0, 0, 201, 200, 1, 0,
		0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0,
		204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 210, 5, 3, 0, 0, 207,
		209, 5, 45, 0, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 213, 214, 5, 35, 0, 0, 214, 252, 1, 0, 0, 0, 215, 216, 5, 4, 0, 0,
		216, 217, 3, 20, 10, 0, 217, 221, 5, 34, 0, 0, 218, 220, 5, 45, 0, 0, 219,
		218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222,
		1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 228, 3, 48,
		24, 0, 225, 227, 5, 45, 0, 0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0,
		0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 232, 5, 35, 0, 0, 232, 252, 1, 0, 0, 0, 233, 234,
		5, 4, 0, 0, 234, 235, 3, 20, 10, 0, 235, 239, 5, 34, 0, 0, 236, 238, 5,
		45, 0, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242,
		246, 3, 50, 25, 0, 243, 245, 5, 45, 0, 0, 244, 243, 1, 0, 0, 0, 245, 248,
		1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 249, 1, 0,
		0, 0, 248, 246, 1, 0, 0, 0, 249, 250, 5, 35, 0, 0, 250, 252, 1, 0, 0, 0,
		251, 197, 1, 0, 0, 0, 251, 215, 1, 0, 0, 0, 251, 233, 1, 0, 0, 0, 252,
		29, 1, 0, 0, 0, 253, 268, 3, 24, 12, 0, 254, 268, 3, 26, 13, 0, 255, 268,
		3, 28, 14, 0, 256, 257, 3, 18, 9, 0, 257, 258, 3, 20, 10, 0, 258, 262,
		5, 34, 0, 0, 259, 261, 5, 45, 0, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1,
		0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0,
		0, 264, 262, 1, 0, 0, 0, 265, 266, 5, 35, 0, 0, 266, 268, 1, 0, 0, 0, 267,
		253, 1, 0, 0, 0, 267, 254, 1, 0, 0, 0, 267, 255, 1, 0, 0, 0, 267, 256,
		1, 0, 0, 0, 268, 31, 1, 0, 0, 0, 269, 270, 3, 18, 9, 0, 270, 271, 3, 20,
		10, 0, 271, 275, 5, 34, 0, 0, 272, 274, 5, 45, 0, 0, 273, 272, 1, 0, 0,
		0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276,
		278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 289, 3, 22, 11, 0, 279, 283,
		5, 40, 0, 0, 280, 282, 5, 45, 0, 0, 281, 280, 1, 0, 0, 0, 282, 285, 1,
		0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 286, 1, 0, 0,
		0, 285, 283, 1, 0, 0, 0, 286, 288, 3, 22, 11, 0, 287, 279, 1, 0, 0, 0,
		288, 291, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290,
		295, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 294, 5, 45, 0, 0, 293, 292,
		1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0,
		0, 0, 296, 298, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 298, 299, 5, 35, 0, 0,
		299, 327, 1, 0, 0, 0, 300, 301, 3, 18, 9, 0, 301, 302, 3, 20, 10, 0, 302,
		306, 5, 34, 0, 0, 303, 305, 5, 45, 0, 0, 304, 303, 1, 0, 0, 0, 305, 308,
		1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 309, 1, 0,
		0, 0, 308, 306, 1, 0, 0, 0, 309, 320, 3, 52, 26, 0, 310, 314, 5, 40, 0,
		0, 311, 313, 5, 45, 0, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314,
		312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314,
		1, 0, 0, 0, 317, 319, 3, 52, 26, 0, 318, 310, 1, 0, 0, 0, 319, 322, 1,
		0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 323, 1, 0, 0,
		0, 322, 320, 1, 0, 0, 0, 323, 324, 5, 45, 0, 0, 324, 325, 5, 35, 0, 0,
		325, 327, 1, 0, 0, 0, 326, 269, 1, 0, 0, 0, 326, 300, 1, 0, 0, 0, 327,
		33, 1, 0, 0, 0, 328, 329, 3, 20, 10, 0, 329, 333, 5, 34, 0, 0, 330, 332,
		5, 45, 0, 0, 331, 330, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0,
		0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0,
		336, 340, 3, 22, 11, 0, 337, 339, 5, 45, 0, 0, 338, 337, 1, 0, 0, 0, 339,
		342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 343,
		1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343, 344, 5, 35, 0, 0, 344, 363, 1, 0,
		0, 0, 345, 346, 3, 20, 10, 0, 346, 350, 5, 34, 0, 0, 347, 349, 5, 45, 0,
		0, 348, 347, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 353, 357,
		3, 52, 26, 0, 354, 356, 5, 45, 0, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1,
		0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 360, 1, 0, 0,
		0, 359, 357, 1, 0, 0, 0, 360, 361, 5, 35, 0, 0, 361, 363, 1, 0, 0, 0, 362,
		328, 1, 0, 0, 0, 362, 345, 1, 0, 0, 0, 363, 35, 1, 0, 0, 0, 364, 365, 7,
		3, 0, 0, 365, 37, 1, 0, 0, 0, 366, 367, 7, 4, 0, 0, 367, 39, 1, 0, 0, 0,
		368, 369, 7, 5, 0, 0, 369, 41, 1, 0, 0, 0, 370, 371, 7, 6, 0, 0, 371, 43,
		1, 0, 0, 0, 372, 373, 5, 32, 0, 0, 373, 45, 1, 0, 0, 0, 374, 375, 6, 23,
		-1, 0, 375, 376, 5, 36, 0, 0, 376, 377, 3, 46, 23, 0, 377, 378, 5, 37,
		0, 0, 378, 385, 1, 0, 0, 0, 379, 385, 5, 5, 0, 0, 380, 385, 3, 58, 29,
		0, 381, 385, 3, 60, 30, 0, 382, 383, 5, 19, 0, 0, 383, 385, 3, 46, 23,
		1, 384, 374, 1, 0, 0, 0, 384, 379, 1, 0, 0, 0, 384, 380, 1, 0, 0, 0, 384,
		381, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 396, 1, 0, 0, 0, 386, 387,
		10, 6, 0, 0, 387, 388, 3, 36, 18, 0, 388, 389, 3, 46, 23, 7, 389, 395,
		1, 0, 0, 0, 390, 391, 10, 5, 0, 0, 391, 392, 3, 38, 19, 0, 392, 393, 3,
		46, 23, 6, 393, 395, 1, 0, 0, 0, 394, 386, 1, 0, 0, 0, 394, 390, 1, 0,
		0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0,
		397, 47, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 400, 6, 24, -1, 0, 400,
		401, 5, 36, 0, 0, 401, 402, 3, 48, 24, 0, 402, 403, 5, 37, 0, 0, 403, 410,
		1, 0, 0, 0, 404, 405, 5, 32, 0, 0, 405, 410, 3, 48, 24, 4, 406, 410, 5,
		3, 0, 0, 407, 410, 3, 58, 29, 0, 408, 410, 3, 60, 30, 0, 409, 399, 1, 0,
		0, 0, 409, 404, 1, 0, 0, 0, 409, 406, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0,
		409, 408, 1, 0, 0, 0, 410, 421, 1, 0, 0, 0, 411, 412, 10, 6, 0, 0, 412,
		413, 3, 40, 20, 0, 413, 414, 3, 48, 24, 7, 414, 420, 1, 0, 0, 0, 415, 416,
		10, 5, 0, 0, 416, 417, 3, 42, 21, 0, 417, 418, 3, 48, 24, 6, 418, 420,
		1, 0, 0, 0, 419, 411, 1, 0, 0, 0, 419, 415, 1, 0, 0, 0, 420, 423, 1, 0,
		0, 0, 421, 419, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 49, 1, 0, 0, 0,
		423, 421, 1, 0, 0, 0, 424, 425, 6, 25, -1, 0, 425, 426, 3, 46, 23, 0, 426,
		427, 3, 42, 21, 0, 427, 428, 3, 46, 23, 0, 428, 439, 1, 0, 0, 0, 429, 430,
		3, 48, 24, 0, 430, 431, 3, 40, 20, 0, 431, 432, 3, 48, 24, 0, 432, 439,
		1, 0, 0, 0, 433, 434, 3, 44, 22, 0, 434, 435, 5, 36, 0, 0, 435, 436, 3,
		50, 25, 0, 436, 437, 5, 37, 0, 0, 437, 439, 1, 0, 0, 0, 438, 424, 1, 0,
		0, 0, 438, 429, 1, 0, 0, 0, 438, 433, 1, 0, 0, 0, 439, 446, 1, 0, 0, 0,
		440, 441, 10, 2, 0, 0, 441, 442, 3, 40, 20, 0, 442, 443, 3, 48, 24, 0,
		443, 445, 1, 0, 0, 0, 444, 440, 1, 0, 0, 0, 445, 448, 1, 0, 0, 0, 446,
		444, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 51, 1, 0, 0, 0, 448, 446, 1,
		0, 0, 0, 449, 453, 3, 58, 29, 0, 450, 453, 3, 46, 23, 0, 451, 453, 3, 48,
		24, 0, 452, 449, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452, 451, 1, 0, 0, 0,
		453, 53, 1, 0, 0, 0, 454, 455, 5, 42, 0, 0, 455, 456, 5, 38, 0, 0, 456,
		457, 3, 46, 23, 0, 457, 458, 5, 39, 0, 0, 458, 55, 1, 0, 0, 0, 459, 460,
		5, 42, 0, 0, 460, 461, 5, 41, 0, 0, 461, 475, 5, 42, 0, 0, 462, 463, 5,
		42, 0, 0, 463, 464, 5, 41, 0, 0, 464, 475, 3, 54, 27, 0, 465, 466, 5, 42,
		0, 0, 466, 467, 5, 41, 0, 0, 467, 475, 3, 56, 28, 0, 468, 469, 3, 54, 27,
		0, 469, 470, 5, 41, 0, 0, 470, 471, 3, 56, 28, 0, 471, 475, 1, 0, 0, 0,
		472, 475, 3, 54, 27, 0, 473, 475, 5, 42, 0, 0, 474, 459, 1, 0, 0, 0, 474,
		462, 1, 0, 0, 0, 474, 465, 1, 0, 0, 0, 474, 468, 1, 0, 0, 0, 474, 472,
		1, 0, 0, 0, 474, 473, 1, 0, 0, 0, 475, 57, 1, 0, 0, 0, 476, 480, 5, 42,
		0, 0, 477, 480, 3, 54, 27, 0, 478, 480, 3, 56, 28, 0, 479, 476, 1, 0, 0,
		0, 479, 477, 1, 0, 0, 0, 479, 478, 1, 0, 0, 0, 480, 59, 1, 0, 0, 0, 481,
		482, 5, 42, 0, 0, 482, 491, 5, 36, 0, 0, 483, 488, 3, 52, 26, 0, 484, 485,
		5, 40, 0, 0, 485, 487, 3, 52, 26, 0, 486, 484, 1, 0, 0, 0, 487, 490, 1,
		0, 0, 0, 488, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 492, 1, 0, 0,
		0, 490, 488, 1, 0, 0, 0, 491, 483, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492,
		493, 1, 0, 0, 0, 493, 494, 5, 37, 0, 0, 494, 61, 1, 0, 0, 0, 50, 65, 76,
		83, 89, 106, 123, 126, 131, 147, 154, 165, 172, 183, 190, 195, 203, 210,
		221, 228, 239, 246, 251, 262, 267, 275, 283, 289, 295, 306, 314, 320, 326,
		333, 340, 350, 357, 362, 384, 394, 396, 409, 419, 421, 438, 446, 452, 474,
		479, 488, 491,
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

// FluxInit initializes any static state used to implement Flux. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFlux(). You can call this function if you wish to initialize the static state ahead
// of time.
func FluxInit() {
	staticData := &FluxParserStaticData
	staticData.once.Do(fluxParserInit)
}

// NewFlux produces a new parser instance for the optional input antlr.TokenStream.
func NewFlux(input antlr.TokenStream) *Flux {
	FluxInit()
	this := new(Flux)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FluxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Flux.g4"

	return this
}

// Flux tokens.
const (
	FluxEOF               = antlr.TokenEOF
	FluxTEXT              = 1
	FluxTEXT_TYPE         = 2
	FluxBOOLEAN           = 3
	FluxBOOLEAN_TYPE      = 4
	FluxNUMBER            = 5
	FluxNUMBER_TYPE       = 6
	FluxNULL              = 7
	FluxDIGIT             = 8
	FluxOCTET             = 9
	FluxIPV4              = 10
	FluxIPV4_TYPE         = 11
	FluxLOOP              = 12
	FluxIF                = 13
	FluxELSE              = 14
	FluxFUNC              = 15
	FluxRETURN            = 16
	FluxRETURN_TYPE       = 17
	FluxOP_ADD            = 18
	FluxOP_SUB            = 19
	FluxOP_MUL            = 20
	FluxOP_DIV            = 21
	FluxOP_MOD            = 22
	FluxOP_POW            = 23
	FluxOP_EQ             = 24
	FluxOP_NEQ            = 25
	FluxOP_GT             = 26
	FluxOP_LT             = 27
	FluxOP_GTE            = 28
	FluxOP_LTE            = 29
	FluxOP_AND            = 30
	FluxOP_OR             = 31
	FluxOP_NOT            = 32
	FluxOP_XOR            = 33
	FluxL_BLOCK           = 34
	FluxR_BLOCK           = 35
	FluxL_PAREN           = 36
	FluxR_PAREN           = 37
	FluxL_SQUARE          = 38
	FluxR_SQUARE          = 39
	FluxCOMMA             = 40
	FluxDOT               = 41
	FluxVAR_IDENTIFIER    = 42
	FluxCOMMON_IDENTIFIER = 43
	FluxAT                = 44
	FluxNEWLINE           = 45
	FluxWS                = 46
)

// Flux rules.
const (
	FluxRULE_program                 = 0
	FluxRULE_statement               = 1
	FluxRULE_expression              = 2
	FluxRULE_block                   = 3
	FluxRULE_loop_statement          = 4
	FluxRULE_if_statement            = 5
	FluxRULE_return_statement        = 6
	FluxRULE_data_type               = 7
	FluxRULE_func_declaration        = 8
	FluxRULE_var_type                = 9
	FluxRULE_var_name                = 10
	FluxRULE_var_value               = 11
	FluxRULE_string_var_declaration  = 12
	FluxRULE_number_var_declaration  = 13
	FluxRULE_boolean_var_declaration = 14
	FluxRULE_single_var_declaration  = 15
	FluxRULE_array_var_declaration   = 16
	FluxRULE_var_assignment          = 17
	FluxRULE_op_level1               = 18
	FluxRULE_op_level2               = 19
	FluxRULE_op_level3               = 20
	FluxRULE_op_level4               = 21
	FluxRULE_op_level5               = 22
	FluxRULE_numeric_expression      = 23
	FluxRULE_logical_expression      = 24
	FluxRULE_comparative_expression  = 25
	FluxRULE_math_expression         = 26
	FluxRULE_get_array_element       = 27
	FluxRULE_get_child               = 28
	FluxRULE_get_var                 = 29
	FluxRULE_function_call           = 30
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other io.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FluxEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FluxRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57174604646484) != 0 {
		{
			p.SetState(62)
			p.Statement()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(FluxEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Loop_statement() ILoop_statementContext
	If_statement() IIf_statementContext
	Func_declaration() IFunc_declarationContext
	Return_statement() IReturn_statementContext
	NEWLINE() antlr.TerminalNode

	// IsStatementContext differentiates from other io.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) Func_declaration() IFunc_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declarationContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FluxRULE_statement)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Loop_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Func_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.Return_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.Match(FluxNEWLINE)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_call() IFunction_callContext
	Var_assignment() IVar_assignmentContext
	Single_var_declaration() ISingle_var_declarationContext
	Array_var_declaration() IArray_var_declarationContext
	Get_var() IGet_varContext

	// IsExpressionContext differentiates from other io.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Var_assignment() IVar_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignmentContext)
}

func (s *ExpressionContext) Single_var_declaration() ISingle_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_var_declarationContext)
}

func (s *ExpressionContext) Array_var_declaration() IArray_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_var_declarationContext)
}

func (s *ExpressionContext) Get_var() IGet_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_varContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FluxRULE_expression)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Var_assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Single_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.Array_var_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(82)
			p.Get_var()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other io.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *BlockContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FluxRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(FluxL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57174604646484) != 0 {
		{
			p.SetState(86)
			p.Statement()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(FluxR_BLOCK)
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

// ILoop_statementContext is an interface to support dynamic dispatch.
type ILoop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	LOOP() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	Block() IBlockContext

	// IsLoop_statementContext differentiates from other io.
	IsLoop_statementContext()
}

type Loop_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_statementContext() *Loop_statementContext {
	var p = new(Loop_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_loop_statement
	return p
}

func InitEmptyLoop_statementContext(p *Loop_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_loop_statement
}

func (*Loop_statementContext) IsLoop_statementContext() {}

func NewLoop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_statementContext {
	var p = new(Loop_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_loop_statement

	return p
}

func (s *Loop_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Loop_statementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(FluxLOOP, 0)
}

func (s *Loop_statementContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Loop_statementContext) Block() IBlockContext {
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

func (s *Loop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterLoop_statement(s)
	}
}

func (s *Loop_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitLoop_statement(s)
	}
}

func (s *Loop_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitLoop_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Loop_statement() (localctx ILoop_statementContext) {
	localctx = NewLoop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FluxRULE_loop_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(FluxLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.comparative_expression(0)
	}
	{
		p.SetState(97)
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAT() []antlr.TerminalNode
	AT(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIf_statementContext differentiates from other io.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(FluxAT)
}

func (s *If_statementContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(FluxAT, i)
}

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(FluxIF, 0)
}

func (s *If_statementContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *If_statementContext) AllBlock() []IBlockContext {
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

func (s *If_statementContext) Block(i int) IBlockContext {
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

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FluxELSE, 0)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FluxRULE_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(FluxIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.comparative_expression(0)
	}
	{
		p.SetState(102)
		p.Block()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(103)
			p.Match(FluxAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(FluxELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Block()
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

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_statementContext differentiates from other io.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(FluxRETURN, 0)
}

func (s *Return_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FluxRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(FluxRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Expression()
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

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	COMMON_IDENTIFIER() antlr.TerminalNode

	// IsData_typeContext differentiates from other io.
	IsData_typeContext()
}

type Data_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *Data_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Data_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Data_typeContext) COMMON_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxCOMMON_IDENTIFIER, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitData_type(s)
	}
}

func (s *Data_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitData_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FluxRULE_data_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796093022292) != 0) {
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

// IFunc_declarationContext is an interface to support dynamic dispatch.
type IFunc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	FUNC() antlr.TerminalNode
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Block() IBlockContext
	RETURN_TYPE() antlr.TerminalNode
	Data_type() IData_typeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunc_declarationContext differentiates from other io.
	IsFunc_declarationContext()
}

type Func_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declarationContext() *Func_declarationContext {
	var p = new(Func_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_func_declaration
	return p
}

func InitEmptyFunc_declarationContext(p *Func_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_func_declaration
}

func (*Func_declarationContext) IsFunc_declarationContext() {}

func NewFunc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declarationContext {
	var p = new(Func_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_func_declaration

	return p
}

func (s *Func_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declarationContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Func_declarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FluxFUNC, 0)
}

func (s *Func_declarationContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FluxVAR_IDENTIFIER)
}

func (s *Func_declarationContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, i)
}

func (s *Func_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Func_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Func_declarationContext) Block() IBlockContext {
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

func (s *Func_declarationContext) RETURN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxRETURN_TYPE, 0)
}

func (s *Func_declarationContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Func_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Func_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Func_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterFunc_declaration(s)
	}
}

func (s *Func_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitFunc_declaration(s)
	}
}

func (s *Func_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitFunc_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Func_declaration() (localctx IFunc_declarationContext) {
	localctx = NewFunc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FluxRULE_func_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(FluxFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(FluxL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxVAR_IDENTIFIER {
		{
			p.SetState(118)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(119)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.Match(FluxVAR_IDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(128)
		p.Match(FluxR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxRETURN_TYPE {
		{
			p.SetState(129)
			p.Match(FluxRETURN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Data_type()
		}

	}
	{
		p.SetState(133)
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	IPV4_TYPE() antlr.TerminalNode

	// IsVar_typeContext differentiates from other io.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *Var_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Var_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Var_typeContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxIPV4_TYPE, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FluxRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2132) != 0) {
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

// IVar_nameContext is an interface to support dynamic dispatch.
type IVar_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode

	// IsVar_nameContext differentiates from other io.
	IsVar_nameContext()
}

type Var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_nameContext() *Var_nameContext {
	var p = new(Var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (s *Var_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FluxRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(FluxVAR_IDENTIFIER)
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

// IVar_valueContext is an interface to support dynamic dispatch.
type IVar_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IPV4() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsVar_valueContext differentiates from other io.
	IsVar_valueContext()
}

type Var_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_valueContext() *Var_valueContext {
	var p = new(Var_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_value
	return p
}

func InitEmptyVar_valueContext(p *Var_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_value
}

func (*Var_valueContext) IsVar_valueContext() {}

func NewVar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_valueContext {
	var p = new(Var_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_value

	return p
}

func (s *Var_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_valueContext) IPV4() antlr.TerminalNode {
	return s.GetToken(FluxIPV4, 0)
}

func (s *Var_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FluxTEXT, 0)
}

func (s *Var_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Var_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Var_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_value(s)
	}
}

func (s *Var_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_value(s)
	}
}

func (s *Var_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_value() (localctx IVar_valueContext) {
	localctx = NewVar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FluxRULE_var_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1066) != 0) {
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

// IString_var_declarationContext is an interface to support dynamic dispatch.
type IString_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsString_var_declarationContext differentiates from other io.
	IsString_var_declarationContext()
}

type String_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_var_declarationContext() *String_var_declarationContext {
	var p = new(String_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_string_var_declaration
	return p
}

func InitEmptyString_var_declarationContext(p *String_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_string_var_declaration
}

func (*String_var_declarationContext) IsString_var_declarationContext() {}

func NewString_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_var_declarationContext {
	var p = new(String_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_string_var_declaration

	return p
}

func (s *String_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *String_var_declarationContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *String_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *String_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *String_var_declarationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FluxTEXT, 0)
}

func (s *String_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *String_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *String_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *String_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitString_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) String_var_declaration() (localctx IString_var_declarationContext) {
	localctx = NewString_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FluxRULE_string_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FluxTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Var_name()
	}
	{
		p.SetState(143)
		p.Match(FluxL_BLOCK)
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

	for _la == FluxNEWLINE {
		{
			p.SetState(144)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(FluxTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FluxNEWLINE {
		{
			p.SetState(151)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(FluxR_BLOCK)
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

// INumber_var_declarationContext is an interface to support dynamic dispatch.
type INumber_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsNumber_var_declarationContext differentiates from other io.
	IsNumber_var_declarationContext()
}

type Number_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_var_declarationContext() *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_number_var_declaration
	return p
}

func InitEmptyNumber_var_declarationContext(p *Number_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_number_var_declaration
}

func (*Number_var_declarationContext) IsNumber_var_declarationContext() {}

func NewNumber_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_number_var_declaration

	return p
}

func (s *Number_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_var_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Number_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Number_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Number_var_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Number_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Number_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Number_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Number_var_declarationContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Number_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitNumber_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Number_var_declaration() (localctx INumber_var_declarationContext) {
	localctx = NewNumber_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FluxRULE_number_var_declaration)
	var _la int

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Var_name()
		}
		{
			p.SetState(161)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(162)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(168)
			p.Match(FluxNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(169)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(175)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Var_name()
		}
		{
			p.SetState(179)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(180)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(186)
			p.Math_expression()
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(187)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(193)
			p.Match(FluxR_BLOCK)
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

// IBoolean_var_declarationContext is an interface to support dynamic dispatch.
type IBoolean_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Logical_expression() ILogical_expressionContext
	Comparative_expression() IComparative_expressionContext

	// IsBoolean_var_declarationContext differentiates from other io.
	IsBoolean_var_declarationContext()
}

type Boolean_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_var_declarationContext() *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_boolean_var_declaration
	return p
}

func InitEmptyBoolean_var_declarationContext(p *Boolean_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_boolean_var_declaration
}

func (*Boolean_var_declarationContext) IsBoolean_var_declarationContext() {}

func NewBoolean_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_boolean_var_declaration

	return p
}

func (s *Boolean_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_var_declarationContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Boolean_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Boolean_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Boolean_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Boolean_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Boolean_var_declarationContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Boolean_var_declarationContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Boolean_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitBoolean_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Boolean_var_declaration() (localctx IBoolean_var_declarationContext) {
	localctx = NewBoolean_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FluxRULE_boolean_var_declaration)
	var _la int

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Var_name()
		}
		{
			p.SetState(199)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(200)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(206)
			p.Match(FluxBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(207)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(213)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Var_name()
		}
		{
			p.SetState(217)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(218)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(224)
			p.logical_expression(0)
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(225)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(231)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Var_name()
		}
		{
			p.SetState(235)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(236)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(242)
			p.comparative_expression(0)
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(243)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(249)
			p.Match(FluxR_BLOCK)
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

// ISingle_var_declarationContext is an interface to support dynamic dispatch.
type ISingle_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_var_declaration() IString_var_declarationContext
	Number_var_declaration() INumber_var_declarationContext
	Boolean_var_declaration() IBoolean_var_declarationContext
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsSingle_var_declarationContext differentiates from other io.
	IsSingle_var_declarationContext()
}

type Single_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_var_declarationContext() *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_single_var_declaration
	return p
}

func InitEmptySingle_var_declarationContext(p *Single_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_single_var_declaration
}

func (*Single_var_declarationContext) IsSingle_var_declarationContext() {}

func NewSingle_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_single_var_declaration

	return p
}

func (s *Single_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_var_declarationContext) String_var_declaration() IString_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_var_declarationContext)
}

func (s *Single_var_declarationContext) Number_var_declaration() INumber_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_var_declarationContext)
}

func (s *Single_var_declarationContext) Boolean_var_declaration() IBoolean_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_var_declarationContext)
}

func (s *Single_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Single_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Single_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Single_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Single_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Single_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Single_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitSingle_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Single_var_declaration() (localctx ISingle_var_declarationContext) {
	localctx = NewSingle_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FluxRULE_single_var_declaration)
	var _la int

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.String_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Number_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Boolean_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.Var_type()
		}
		{
			p.SetState(257)
			p.Var_name()
		}
		{
			p.SetState(258)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(259)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(264)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(265)
			p.Match(FluxR_BLOCK)
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

// IArray_var_declarationContext is an interface to support dynamic dispatch.
type IArray_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	AllVar_value() []IVar_valueContext
	Var_value(i int) IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext

	// IsArray_var_declarationContext differentiates from other io.
	IsArray_var_declarationContext()
}

type Array_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_var_declarationContext() *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_array_var_declaration
	return p
}

func InitEmptyArray_var_declarationContext(p *Array_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_array_var_declaration
}

func (*Array_var_declarationContext) IsArray_var_declarationContext() {}

func NewArray_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_array_var_declaration

	return p
}

func (s *Array_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Array_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Array_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllVar_value() []IVar_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_valueContext); ok {
			len++
		}
	}

	tst := make([]IVar_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_valueContext); ok {
			tst[i] = t.(IVar_valueContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Var_value(i int) IVar_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
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

	return t.(IVar_valueContext)
}

func (s *Array_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Array_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Array_var_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Array_var_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Array_var_declarationContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
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

	return t.(IMath_expressionContext)
}

func (s *Array_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitArray_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Array_var_declaration() (localctx IArray_var_declarationContext) {
	localctx = NewArray_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FluxRULE_array_var_declaration)
	var _la int

	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			p.Var_type()
		}
		{
			p.SetState(270)
			p.Var_name()
		}
		{
			p.SetState(271)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(272)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
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
			p.SetState(278)
			p.Var_value()
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(279)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FluxNEWLINE {
				{
					p.SetState(280)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(285)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(286)
				p.Var_value()
			}

			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(292)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(298)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Var_type()
		}
		{
			p.SetState(301)
			p.Var_name()
		}
		{
			p.SetState(302)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(303)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(308)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(309)
			p.Math_expression()
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(310)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FluxNEWLINE {
				{
					p.SetState(311)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(316)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(317)
				p.Math_expression()
			}

			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(323)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(FluxR_BLOCK)
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

// IVar_assignmentContext is an interface to support dynamic dispatch.
type IVar_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	Var_value() IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsVar_assignmentContext differentiates from other io.
	IsVar_assignmentContext()
}

type Var_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignmentContext() *Var_assignmentContext {
	var p = new(Var_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_assignment
	return p
}

func InitEmptyVar_assignmentContext(p *Var_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_assignment
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_assignment

	return p
}

func (s *Var_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignmentContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Var_assignmentContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Var_assignmentContext) Var_value() IVar_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_valueContext)
}

func (s *Var_assignmentContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Var_assignmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Var_assignmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Var_assignmentContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Var_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (s *Var_assignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_assignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FluxRULE_var_assignment)
	var _la int

	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Var_name()
		}
		{
			p.SetState(329)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(330)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(335)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(336)
			p.Var_value()
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(337)
				p.Match(FluxNEWLINE)
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
		}
		{
			p.SetState(343)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Var_name()
		}
		{
			p.SetState(346)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(347)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(352)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(353)
			p.Math_expression()
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(354)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(360)
			p.Match(FluxR_BLOCK)
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

// IOp_level1Context is an interface to support dynamic dispatch.
type IOp_level1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_MUL() antlr.TerminalNode
	OP_DIV() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode
	OP_POW() antlr.TerminalNode

	// IsOp_level1Context differentiates from other io.
	IsOp_level1Context()
}

type Op_level1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level1Context() *Op_level1Context {
	var p = new(Op_level1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level1
	return p
}

func InitEmptyOp_level1Context(p *Op_level1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level1
}

func (*Op_level1Context) IsOp_level1Context() {}

func NewOp_level1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level1Context {
	var p = new(Op_level1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level1

	return p
}

func (s *Op_level1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level1Context) OP_MUL() antlr.TerminalNode {
	return s.GetToken(FluxOP_MUL, 0)
}

func (s *Op_level1Context) OP_DIV() antlr.TerminalNode {
	return s.GetToken(FluxOP_DIV, 0)
}

func (s *Op_level1Context) OP_MOD() antlr.TerminalNode {
	return s.GetToken(FluxOP_MOD, 0)
}

func (s *Op_level1Context) OP_POW() antlr.TerminalNode {
	return s.GetToken(FluxOP_POW, 0)
}

func (s *Op_level1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level1(s)
	}
}

func (s *Op_level1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level1(s)
	}
}

func (s *Op_level1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level1() (localctx IOp_level1Context) {
	localctx = NewOp_level1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FluxRULE_op_level1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
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

// IOp_level2Context is an interface to support dynamic dispatch.
type IOp_level2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_ADD() antlr.TerminalNode
	OP_SUB() antlr.TerminalNode

	// IsOp_level2Context differentiates from other io.
	IsOp_level2Context()
}

type Op_level2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level2Context() *Op_level2Context {
	var p = new(Op_level2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level2
	return p
}

func InitEmptyOp_level2Context(p *Op_level2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level2
}

func (*Op_level2Context) IsOp_level2Context() {}

func NewOp_level2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level2Context {
	var p = new(Op_level2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level2

	return p
}

func (s *Op_level2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level2Context) OP_ADD() antlr.TerminalNode {
	return s.GetToken(FluxOP_ADD, 0)
}

func (s *Op_level2Context) OP_SUB() antlr.TerminalNode {
	return s.GetToken(FluxOP_SUB, 0)
}

func (s *Op_level2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level2(s)
	}
}

func (s *Op_level2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level2(s)
	}
}

func (s *Op_level2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level2() (localctx IOp_level2Context) {
	localctx = NewOp_level2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FluxRULE_op_level2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FluxOP_ADD || _la == FluxOP_SUB) {
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

// IOp_level3Context is an interface to support dynamic dispatch.
type IOp_level3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode
	OP_XOR() antlr.TerminalNode

	// IsOp_level3Context differentiates from other io.
	IsOp_level3Context()
}

type Op_level3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level3Context() *Op_level3Context {
	var p = new(Op_level3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level3
	return p
}

func InitEmptyOp_level3Context(p *Op_level3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level3
}

func (*Op_level3Context) IsOp_level3Context() {}

func NewOp_level3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level3Context {
	var p = new(Op_level3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level3

	return p
}

func (s *Op_level3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level3Context) OP_AND() antlr.TerminalNode {
	return s.GetToken(FluxOP_AND, 0)
}

func (s *Op_level3Context) OP_OR() antlr.TerminalNode {
	return s.GetToken(FluxOP_OR, 0)
}

func (s *Op_level3Context) OP_XOR() antlr.TerminalNode {
	return s.GetToken(FluxOP_XOR, 0)
}

func (s *Op_level3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level3(s)
	}
}

func (s *Op_level3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level3(s)
	}
}

func (s *Op_level3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level3() (localctx IOp_level3Context) {
	localctx = NewOp_level3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FluxRULE_op_level3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&11811160064) != 0) {
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

// IOp_level4Context is an interface to support dynamic dispatch.
type IOp_level4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_EQ() antlr.TerminalNode
	OP_NEQ() antlr.TerminalNode
	OP_LT() antlr.TerminalNode
	OP_GT() antlr.TerminalNode
	OP_GTE() antlr.TerminalNode
	OP_LTE() antlr.TerminalNode

	// IsOp_level4Context differentiates from other io.
	IsOp_level4Context()
}

type Op_level4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level4Context() *Op_level4Context {
	var p = new(Op_level4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level4
	return p
}

func InitEmptyOp_level4Context(p *Op_level4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level4
}

func (*Op_level4Context) IsOp_level4Context() {}

func NewOp_level4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level4Context {
	var p = new(Op_level4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level4

	return p
}

func (s *Op_level4Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level4Context) OP_EQ() antlr.TerminalNode {
	return s.GetToken(FluxOP_EQ, 0)
}

func (s *Op_level4Context) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(FluxOP_NEQ, 0)
}

func (s *Op_level4Context) OP_LT() antlr.TerminalNode {
	return s.GetToken(FluxOP_LT, 0)
}

func (s *Op_level4Context) OP_GT() antlr.TerminalNode {
	return s.GetToken(FluxOP_GT, 0)
}

func (s *Op_level4Context) OP_GTE() antlr.TerminalNode {
	return s.GetToken(FluxOP_GTE, 0)
}

func (s *Op_level4Context) OP_LTE() antlr.TerminalNode {
	return s.GetToken(FluxOP_LTE, 0)
}

func (s *Op_level4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level4(s)
	}
}

func (s *Op_level4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level4(s)
	}
}

func (s *Op_level4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level4(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level4() (localctx IOp_level4Context) {
	localctx = NewOp_level4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FluxRULE_op_level4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1056964608) != 0) {
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

// IOp_level5Context is an interface to support dynamic dispatch.
type IOp_level5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_NOT() antlr.TerminalNode

	// IsOp_level5Context differentiates from other io.
	IsOp_level5Context()
}

type Op_level5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level5Context() *Op_level5Context {
	var p = new(Op_level5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level5
	return p
}

func InitEmptyOp_level5Context(p *Op_level5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level5
}

func (*Op_level5Context) IsOp_level5Context() {}

func NewOp_level5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level5Context {
	var p = new(Op_level5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level5

	return p
}

func (s *Op_level5Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level5Context) OP_NOT() antlr.TerminalNode {
	return s.GetToken(FluxOP_NOT, 0)
}

func (s *Op_level5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level5(s)
	}
}

func (s *Op_level5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level5(s)
	}
}

func (s *Op_level5Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level5(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level5() (localctx IOp_level5Context) {
	localctx = NewOp_level5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FluxRULE_op_level5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(FluxOP_NOT)
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

// INumeric_expressionContext is an interface to support dynamic dispatch.
type INumeric_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllNumeric_expression() []INumeric_expressionContext
	Numeric_expression(i int) INumeric_expressionContext
	R_PAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Get_var() IGet_varContext
	Function_call() IFunction_callContext
	OP_SUB() antlr.TerminalNode
	Op_level1() IOp_level1Context
	Op_level2() IOp_level2Context

	// IsNumeric_expressionContext differentiates from other io.
	IsNumeric_expressionContext()
}

type Numeric_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_expressionContext() *Numeric_expressionContext {
	var p = new(Numeric_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_numeric_expression
	return p
}

func InitEmptyNumeric_expressionContext(p *Numeric_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_numeric_expression
}

func (*Numeric_expressionContext) IsNumeric_expressionContext() {}

func NewNumeric_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_expressionContext {
	var p = new(Numeric_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_numeric_expression

	return p
}

func (s *Numeric_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Numeric_expressionContext) AllNumeric_expression() []INumeric_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumeric_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumeric_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumeric_expressionContext); ok {
			tst[i] = t.(INumeric_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Numeric_expressionContext) Numeric_expression(i int) INumeric_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_expressionContext); ok {
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

	return t.(INumeric_expressionContext)
}

func (s *Numeric_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Numeric_expressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Numeric_expressionContext) Get_var() IGet_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_varContext)
}

func (s *Numeric_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Numeric_expressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(FluxOP_SUB, 0)
}

func (s *Numeric_expressionContext) Op_level1() IOp_level1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level1Context)
}

func (s *Numeric_expressionContext) Op_level2() IOp_level2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level2Context)
}

func (s *Numeric_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterNumeric_expression(s)
	}
}

func (s *Numeric_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitNumeric_expression(s)
	}
}

func (s *Numeric_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitNumeric_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Numeric_expression() (localctx INumeric_expressionContext) {
	return p.numeric_expression(0)
}

func (p *Flux) numeric_expression(_p int) (localctx INumeric_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNumeric_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumeric_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, FluxRULE_numeric_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(375)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.numeric_expression(0)
		}
		{
			p.SetState(377)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(379)
			p.Match(FluxNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(380)
			p.Get_var()
		}

	case 4:
		{
			p.SetState(381)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(382)
			p.Match(FluxOP_SUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.numeric_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumeric_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numeric_expression)
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(387)
					p.Op_level1()
				}
				{
					p.SetState(388)
					p.numeric_expression(7)
				}

			case 2:
				localctx = NewNumeric_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numeric_expression)
				p.SetState(390)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(391)
					p.Op_level2()
				}
				{
					p.SetState(392)
					p.numeric_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
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

// ILogical_expressionContext is an interface to support dynamic dispatch.
type ILogical_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	R_PAREN() antlr.TerminalNode
	OP_NOT() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	Get_var() IGet_varContext
	Function_call() IFunction_callContext
	Op_level3() IOp_level3Context
	Op_level4() IOp_level4Context

	// IsLogical_expressionContext differentiates from other io.
	IsLogical_expressionContext()
}

type Logical_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_expressionContext() *Logical_expressionContext {
	var p = new(Logical_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_logical_expression
	return p
}

func InitEmptyLogical_expressionContext(p *Logical_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_logical_expression
}

func (*Logical_expressionContext) IsLogical_expressionContext() {}

func NewLogical_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_expressionContext {
	var p = new(Logical_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_logical_expression

	return p
}

func (s *Logical_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Logical_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Logical_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
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

	return t.(ILogical_expressionContext)
}

func (s *Logical_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Logical_expressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(FluxOP_NOT, 0)
}

func (s *Logical_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Logical_expressionContext) Get_var() IGet_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_varContext)
}

func (s *Logical_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Logical_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Logical_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Logical_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterLogical_expression(s)
	}
}

func (s *Logical_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitLogical_expression(s)
	}
}

func (s *Logical_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitLogical_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Logical_expression() (localctx ILogical_expressionContext) {
	return p.logical_expression(0)
}

func (p *Flux) logical_expression(_p int) (localctx ILogical_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, FluxRULE_logical_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(400)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.logical_expression(0)
		}
		{
			p.SetState(402)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(404)
			p.Match(FluxOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(406)
			p.Match(FluxBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(407)
			p.Get_var()
		}

	case 5:
		{
			p.SetState(408)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(419)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(412)
					p.Op_level3()
				}
				{
					p.SetState(413)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(415)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(416)
					p.Op_level4()
				}
				{
					p.SetState(417)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
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

// IComparative_expressionContext is an interface to support dynamic dispatch.
type IComparative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumeric_expression() []INumeric_expressionContext
	Numeric_expression(i int) INumeric_expressionContext
	Op_level4() IOp_level4Context
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	Op_level3() IOp_level3Context
	Op_level5() IOp_level5Context
	L_PAREN() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	R_PAREN() antlr.TerminalNode

	// IsComparative_expressionContext differentiates from other io.
	IsComparative_expressionContext()
}

type Comparative_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparative_expressionContext() *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_comparative_expression
	return p
}

func InitEmptyComparative_expressionContext(p *Comparative_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_comparative_expression
}

func (*Comparative_expressionContext) IsComparative_expressionContext() {}

func NewComparative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_comparative_expression

	return p
}

func (s *Comparative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparative_expressionContext) AllNumeric_expression() []INumeric_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumeric_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumeric_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumeric_expressionContext); ok {
			tst[i] = t.(INumeric_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Numeric_expression(i int) INumeric_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_expressionContext); ok {
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

	return t.(INumeric_expressionContext)
}

func (s *Comparative_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Comparative_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
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

	return t.(ILogical_expressionContext)
}

func (s *Comparative_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Comparative_expressionContext) Op_level5() IOp_level5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level5Context)
}

func (s *Comparative_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Comparative_expressionContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Comparative_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Comparative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitComparative_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Comparative_expression() (localctx IComparative_expressionContext) {
	return p.comparative_expression(0)
}

func (p *Flux) comparative_expression(_p int) (localctx IComparative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComparative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, FluxRULE_comparative_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(425)
			p.numeric_expression(0)
		}
		{
			p.SetState(426)
			p.Op_level4()
		}
		{
			p.SetState(427)
			p.numeric_expression(0)
		}

	case 2:
		{
			p.SetState(429)
			p.logical_expression(0)
		}
		{
			p.SetState(430)
			p.Op_level3()
		}
		{
			p.SetState(431)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(433)
			p.Op_level5()
		}
		{
			p.SetState(434)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.comparative_expression(0)
		}
		{
			p.SetState(436)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(446)
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
			localctx = NewComparative_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FluxRULE_comparative_expression)
			p.SetState(440)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(441)
				p.Op_level3()
			}
			{
				p.SetState(442)
				p.logical_expression(0)
			}

		}
		p.SetState(448)
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

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Get_var() IGet_varContext
	Numeric_expression() INumeric_expressionContext
	Logical_expression() ILogical_expressionContext

	// IsMath_expressionContext differentiates from other io.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) Get_var() IGet_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_varContext)
}

func (s *Math_expressionContext) Numeric_expression() INumeric_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumeric_expressionContext)
}

func (s *Math_expressionContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (s *Math_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitMath_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FluxRULE_math_expression)
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(449)
			p.Get_var()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			p.numeric_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(451)
			p.logical_expression(0)
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

// IGet_array_elementContext is an interface to support dynamic dispatch.
type IGet_array_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_SQUARE() antlr.TerminalNode
	Numeric_expression() INumeric_expressionContext
	R_SQUARE() antlr.TerminalNode

	// IsGet_array_elementContext differentiates from other io.
	IsGet_array_elementContext()
}

type Get_array_elementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_array_elementContext() *Get_array_elementContext {
	var p = new(Get_array_elementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_array_element
	return p
}

func InitEmptyGet_array_elementContext(p *Get_array_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_array_element
}

func (*Get_array_elementContext) IsGet_array_elementContext() {}

func NewGet_array_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_array_elementContext {
	var p = new(Get_array_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_array_element

	return p
}

func (s *Get_array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_array_elementContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Get_array_elementContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(FluxL_SQUARE, 0)
}

func (s *Get_array_elementContext) Numeric_expression() INumeric_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumeric_expressionContext)
}

func (s *Get_array_elementContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(FluxR_SQUARE, 0)
}

func (s *Get_array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_array_element(s)
	}
}

func (s *Get_array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_array_element(s)
	}
}

func (s *Get_array_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_array_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_array_element() (localctx IGet_array_elementContext) {
	localctx = NewGet_array_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FluxRULE_get_array_element)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(FluxL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.numeric_expression(0)
	}
	{
		p.SetState(457)
		p.Match(FluxR_SQUARE)
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

// IGet_childContext is an interface to support dynamic dispatch.
type IGet_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_childContext differentiates from other io.
	IsGet_childContext()
}

type Get_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_childContext() *Get_childContext {
	var p = new(Get_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_child
	return p
}

func InitEmptyGet_childContext(p *Get_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_child
}

func (*Get_childContext) IsGet_childContext() {}

func NewGet_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_childContext {
	var p = new(Get_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_child

	return p
}

func (s *Get_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_childContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FluxVAR_IDENTIFIER)
}

func (s *Get_childContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, i)
}

func (s *Get_childContext) DOT() antlr.TerminalNode {
	return s.GetToken(FluxDOT, 0)
}

func (s *Get_childContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_childContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_child(s)
	}
}

func (s *Get_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_child(s)
	}
}

func (s *Get_childContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_child(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_child() (localctx IGet_childContext) {
	localctx = NewGet_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FluxRULE_get_child)
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(468)
			p.Get_array_element()
		}
		{
			p.SetState(469)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(472)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(473)
			p.Match(FluxVAR_IDENTIFIER)
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

// IGet_varContext is an interface to support dynamic dispatch.
type IGet_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_varContext differentiates from other io.
	IsGet_varContext()
}

type Get_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_varContext() *Get_varContext {
	var p = new(Get_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_var
	return p
}

func InitEmptyGet_varContext(p *Get_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_var
}

func (*Get_varContext) IsGet_varContext() {}

func NewGet_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_varContext {
	var p = new(Get_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_var

	return p
}

func (s *Get_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_varContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Get_varContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_varContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_var(s)
	}
}

func (s *Get_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_var(s)
	}
}

func (s *Get_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_var() (localctx IGet_varContext) {
	localctx = NewGet_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FluxRULE_get_var)
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(478)
			p.Get_child()
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

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other io.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Function_callContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
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

	return t.(IMath_expressionContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FluxRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.Match(FluxL_PAREN)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4471061479464) != 0 {
		{
			p.SetState(483)
			p.Math_expression()
		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(484)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(485)
				p.Math_expression()
			}

			p.SetState(490)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(493)
		p.Match(FluxR_PAREN)
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

func (p *Flux) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 23:
		var t *Numeric_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Numeric_expressionContext)
		}
		return p.Numeric_expression_Sempred(t, predIndex)

	case 24:
		var t *Logical_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_expressionContext)
		}
		return p.Logical_expression_Sempred(t, predIndex)

	case 25:
		var t *Comparative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Comparative_expressionContext)
		}
		return p.Comparative_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Flux) Numeric_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Flux) Logical_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Flux) Comparative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
