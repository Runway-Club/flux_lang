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
		"get_child", "get_var", "args", "function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 499, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 5, 0, 66, 8, 0, 10, 0, 12, 0, 69, 9, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 79, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 86, 8, 2, 1, 3, 1, 3, 5, 3, 90, 8, 3, 10, 3, 12, 3, 93, 9, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 109, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 124, 8, 8, 10, 8, 12, 8, 127, 9,
		8, 3, 8, 129, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 134, 8, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 148,
		8, 12, 10, 12, 12, 12, 151, 9, 12, 1, 12, 1, 12, 5, 12, 155, 8, 12, 10,
		12, 12, 12, 158, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13,
		166, 8, 13, 10, 13, 12, 13, 169, 9, 13, 1, 13, 1, 13, 5, 13, 173, 8, 13,
		10, 13, 12, 13, 176, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 184, 8, 13, 10, 13, 12, 13, 187, 9, 13, 1, 13, 1, 13, 5, 13, 191, 8,
		13, 10, 13, 12, 13, 194, 9, 13, 1, 13, 1, 13, 3, 13, 198, 8, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 204, 8, 14, 10, 14, 12, 14, 207, 9, 14, 1,
		14, 1, 14, 5, 14, 211, 8, 14, 10, 14, 12, 14, 214, 9, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 222, 8, 14, 10, 14, 12, 14, 225, 9,
		14, 1, 14, 1, 14, 5, 14, 229, 8, 14, 10, 14, 12, 14, 232, 9, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 240, 8, 14, 10, 14, 12, 14, 243,
		9, 14, 1, 14, 1, 14, 5, 14, 247, 8, 14, 10, 14, 12, 14, 250, 9, 14, 1,
		14, 1, 14, 3, 14, 254, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 5, 15, 263, 8, 15, 10, 15, 12, 15, 266, 9, 15, 1, 15, 1, 15, 3,
		15, 270, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 276, 8, 16, 10, 16,
		12, 16, 279, 9, 16, 1, 16, 1, 16, 1, 16, 5, 16, 284, 8, 16, 10, 16, 12,
		16, 287, 9, 16, 1, 16, 5, 16, 290, 8, 16, 10, 16, 12, 16, 293, 9, 16, 1,
		16, 5, 16, 296, 8, 16, 10, 16, 12, 16, 299, 9, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 5, 16, 307, 8, 16, 10, 16, 12, 16, 310, 9, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 315, 8, 16, 10, 16, 12, 16, 318, 9, 16, 1, 16,
		5, 16, 321, 8, 16, 10, 16, 12, 16, 324, 9, 16, 1, 16, 1, 16, 1, 16, 3,
		16, 329, 8, 16, 1, 17, 1, 17, 1, 17, 5, 17, 334, 8, 17, 10, 17, 12, 17,
		337, 9, 17, 1, 17, 1, 17, 5, 17, 341, 8, 17, 10, 17, 12, 17, 344, 9, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 351, 8, 17, 10, 17, 12, 17, 354,
		9, 17, 1, 17, 1, 17, 5, 17, 358, 8, 17, 10, 17, 12, 17, 361, 9, 17, 1,
		17, 1, 17, 3, 17, 365, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 3, 23, 387, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 397, 8, 23, 10, 23, 12, 23, 400, 9,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 412, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 422, 8, 24, 10, 24, 12, 24, 425, 9, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 3, 25, 441, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 447, 8, 25, 10,
		25, 12, 25, 450, 9, 25, 1, 26, 1, 26, 1, 26, 3, 26, 455, 8, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 477,
		8, 28, 1, 29, 1, 29, 1, 29, 3, 29, 482, 8, 29, 1, 30, 1, 30, 1, 30, 5,
		30, 487, 8, 30, 10, 30, 12, 30, 490, 9, 30, 1, 31, 1, 31, 1, 31, 3, 31,
		495, 8, 31, 1, 31, 1, 31, 1, 31, 0, 3, 46, 48, 50, 32, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 0, 7, 4, 0, 2, 2, 4, 4, 6, 6, 43, 43, 4,
		0, 2, 2, 4, 4, 6, 6, 11, 11, 4, 0, 1, 1, 3, 3, 5, 5, 10, 10, 1, 0, 20,
		23, 1, 0, 18, 19, 2, 0, 30, 31, 33, 33, 1, 0, 24, 29, 539, 0, 67, 1, 0,
		0, 0, 2, 78, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 96, 1,
		0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 114, 1, 0, 0, 0,
		16, 116, 1, 0, 0, 0, 18, 137, 1, 0, 0, 0, 20, 139, 1, 0, 0, 0, 22, 141,
		1, 0, 0, 0, 24, 143, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 253, 1, 0, 0,
		0, 30, 269, 1, 0, 0, 0, 32, 328, 1, 0, 0, 0, 34, 364, 1, 0, 0, 0, 36, 366,
		1, 0, 0, 0, 38, 368, 1, 0, 0, 0, 40, 370, 1, 0, 0, 0, 42, 372, 1, 0, 0,
		0, 44, 374, 1, 0, 0, 0, 46, 386, 1, 0, 0, 0, 48, 411, 1, 0, 0, 0, 50, 440,
		1, 0, 0, 0, 52, 454, 1, 0, 0, 0, 54, 456, 1, 0, 0, 0, 56, 476, 1, 0, 0,
		0, 58, 481, 1, 0, 0, 0, 60, 483, 1, 0, 0, 0, 62, 491, 1, 0, 0, 0, 64, 66,
		3, 2, 1, 0, 65, 64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0,
		67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5,
		0, 0, 1, 71, 1, 1, 0, 0, 0, 72, 79, 3, 4, 2, 0, 73, 79, 3, 8, 4, 0, 74,
		79, 3, 10, 5, 0, 75, 79, 3, 16, 8, 0, 76, 79, 3, 12, 6, 0, 77, 79, 5, 45,
		0, 0, 78, 72, 1, 0, 0, 0, 78, 73, 1, 0, 0, 0, 78, 74, 1, 0, 0, 0, 78, 75,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 3, 1, 0, 0, 0,
		80, 86, 3, 62, 31, 0, 81, 86, 3, 34, 17, 0, 82, 86, 3, 30, 15, 0, 83, 86,
		3, 32, 16, 0, 84, 86, 3, 58, 29, 0, 85, 80, 1, 0, 0, 0, 85, 81, 1, 0, 0,
		0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 5, 1,
		0, 0, 0, 87, 91, 5, 34, 0, 0, 88, 90, 3, 2, 1, 0, 89, 88, 1, 0, 0, 0, 90,
		93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 35, 0, 0, 95, 7, 1, 0, 0, 0, 96, 97,
		5, 44, 0, 0, 97, 98, 5, 12, 0, 0, 98, 99, 3, 50, 25, 0, 99, 100, 3, 6,
		3, 0, 100, 9, 1, 0, 0, 0, 101, 102, 5, 44, 0, 0, 102, 103, 5, 13, 0, 0,
		103, 104, 3, 50, 25, 0, 104, 108, 3, 6, 3, 0, 105, 106, 5, 44, 0, 0, 106,
		107, 5, 14, 0, 0, 107, 109, 3, 6, 3, 0, 108, 105, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 11, 1, 0, 0, 0, 110, 111, 5, 44, 0, 0, 111, 112, 5, 16,
		0, 0, 112, 113, 3, 4, 2, 0, 113, 13, 1, 0, 0, 0, 114, 115, 7, 0, 0, 0,
		115, 15, 1, 0, 0, 0, 116, 117, 5, 44, 0, 0, 117, 118, 5, 15, 0, 0, 118,
		119, 5, 42, 0, 0, 119, 128, 5, 36, 0, 0, 120, 125, 5, 42, 0, 0, 121, 122,
		5, 40, 0, 0, 122, 124, 5, 42, 0, 0, 123, 121, 1, 0, 0, 0, 124, 127, 1,
		0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 129, 1, 0, 0,
		0, 127, 125, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		130, 1, 0, 0, 0, 130, 133, 5, 37, 0, 0, 131, 132, 5, 17, 0, 0, 132, 134,
		3, 14, 7, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0,
		0, 0, 135, 136, 3, 6, 3, 0, 136, 17, 1, 0, 0, 0, 137, 138, 7, 1, 0, 0,
		138, 19, 1, 0, 0, 0, 139, 140, 5, 42, 0, 0, 140, 21, 1, 0, 0, 0, 141, 142,
		7, 2, 0, 0, 142, 23, 1, 0, 0, 0, 143, 144, 5, 2, 0, 0, 144, 145, 3, 20,
		10, 0, 145, 149, 5, 34, 0, 0, 146, 148, 5, 45, 0, 0, 147, 146, 1, 0, 0,
		0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 156, 5, 1, 0, 0, 153, 155,
		5, 45, 0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0,
		159, 160, 5, 35, 0, 0, 160, 25, 1, 0, 0, 0, 161, 162, 5, 6, 0, 0, 162,
		163, 3, 20, 10, 0, 163, 167, 5, 34, 0, 0, 164, 166, 5, 45, 0, 0, 165, 164,
		1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0,
		0, 0, 168, 170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 174, 5, 5, 0, 0,
		171, 173, 5, 45, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174,
		172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 174,
		1, 0, 0, 0, 177, 178, 5, 35, 0, 0, 178, 198, 1, 0, 0, 0, 179, 180, 5, 6,
		0, 0, 180, 181, 3, 20, 10, 0, 181, 185, 5, 34, 0, 0, 182, 184, 5, 45, 0,
		0, 183, 182, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 192,
		3, 52, 26, 0, 189, 191, 5, 45, 0, 0, 190, 189, 1, 0, 0, 0, 191, 194, 1,
		0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 195, 196, 5, 35, 0, 0, 196, 198, 1, 0, 0, 0, 197,
		161, 1, 0, 0, 0, 197, 179, 1, 0, 0, 0, 198, 27, 1, 0, 0, 0, 199, 200, 5,
		4, 0, 0, 200, 201, 3, 20, 10, 0, 201, 205, 5, 34, 0, 0, 202, 204, 5, 45,
		0, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0,
		205, 206, 1, 0, 0, 0, 206, 208, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208,
		212, 5, 3, 0, 0, 209, 211, 5, 45, 0, 0, 210, 209, 1, 0, 0, 0, 211, 214,
		1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0,
		0, 0, 214, 212, 1, 0, 0, 0, 215, 216, 5, 35, 0, 0, 216, 254, 1, 0, 0, 0,
		217, 218, 5, 4, 0, 0, 218, 219, 3, 20, 10, 0, 219, 223, 5, 34, 0, 0, 220,
		222, 5, 45, 0, 0, 221, 220, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221,
		1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0, 0, 225, 223, 1, 0,
		0, 0, 226, 230, 3, 48, 24, 0, 227, 229, 5, 45, 0, 0, 228, 227, 1, 0, 0,
		0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231,
		233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 5, 35, 0, 0, 234, 254,
		1, 0, 0, 0, 235, 236, 5, 4, 0, 0, 236, 237, 3, 20, 10, 0, 237, 241, 5,
		34, 0, 0, 238, 240, 5, 45, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0,
		0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0,
		243, 241, 1, 0, 0, 0, 244, 248, 3, 50, 25, 0, 245, 247, 5, 45, 0, 0, 246,
		245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252, 5, 35,
		0, 0, 252, 254, 1, 0, 0, 0, 253, 199, 1, 0, 0, 0, 253, 217, 1, 0, 0, 0,
		253, 235, 1, 0, 0, 0, 254, 29, 1, 0, 0, 0, 255, 270, 3, 24, 12, 0, 256,
		270, 3, 26, 13, 0, 257, 270, 3, 28, 14, 0, 258, 259, 3, 18, 9, 0, 259,
		260, 3, 20, 10, 0, 260, 264, 5, 34, 0, 0, 261, 263, 5, 45, 0, 0, 262, 261,
		1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0,
		0, 0, 265, 267, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 268, 5, 35, 0, 0,
		268, 270, 1, 0, 0, 0, 269, 255, 1, 0, 0, 0, 269, 256, 1, 0, 0, 0, 269,
		257, 1, 0, 0, 0, 269, 258, 1, 0, 0, 0, 270, 31, 1, 0, 0, 0, 271, 272, 3,
		18, 9, 0, 272, 273, 3, 20, 10, 0, 273, 277, 5, 34, 0, 0, 274, 276, 5, 45,
		0, 0, 275, 274, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0,
		277, 278, 1, 0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280,
		291, 3, 22, 11, 0, 281, 285, 5, 40, 0, 0, 282, 284, 5, 45, 0, 0, 283, 282,
		1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0,
		0, 0, 286, 288, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 288, 290, 3, 22, 11,
		0, 289, 281, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 297, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 296,
		5, 45, 0, 0, 295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0,
		0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0,
		300, 301, 5, 35, 0, 0, 301, 329, 1, 0, 0, 0, 302, 303, 3, 18, 9, 0, 303,
		304, 3, 20, 10, 0, 304, 308, 5, 34, 0, 0, 305, 307, 5, 45, 0, 0, 306, 305,
		1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0,
		0, 0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 322, 3, 52, 26,
		0, 312, 316, 5, 40, 0, 0, 313, 315, 5, 45, 0, 0, 314, 313, 1, 0, 0, 0,
		315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317,
		319, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 321, 3, 52, 26, 0, 320, 312,
		1, 0, 0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0,
		0, 0, 323, 325, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 325, 326, 5, 45, 0, 0,
		326, 327, 5, 35, 0, 0, 327, 329, 1, 0, 0, 0, 328, 271, 1, 0, 0, 0, 328,
		302, 1, 0, 0, 0, 329, 33, 1, 0, 0, 0, 330, 331, 3, 20, 10, 0, 331, 335,
		5, 34, 0, 0, 332, 334, 5, 45, 0, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1,
		0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 338, 1, 0, 0,
		0, 337, 335, 1, 0, 0, 0, 338, 342, 3, 22, 11, 0, 339, 341, 5, 45, 0, 0,
		340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342,
		343, 1, 0, 0, 0, 343, 345, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 346,
		5, 35, 0, 0, 346, 365, 1, 0, 0, 0, 347, 348, 3, 20, 10, 0, 348, 352, 5,
		34, 0, 0, 349, 351, 5, 45, 0, 0, 350, 349, 1, 0, 0, 0, 351, 354, 1, 0,
		0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 355, 1, 0, 0, 0,
		354, 352, 1, 0, 0, 0, 355, 359, 3, 52, 26, 0, 356, 358, 5, 45, 0, 0, 357,
		356, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360,
		1, 0, 0, 0, 360, 362, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 363, 5, 35,
		0, 0, 363, 365, 1, 0, 0, 0, 364, 330, 1, 0, 0, 0, 364, 347, 1, 0, 0, 0,
		365, 35, 1, 0, 0, 0, 366, 367, 7, 3, 0, 0, 367, 37, 1, 0, 0, 0, 368, 369,
		7, 4, 0, 0, 369, 39, 1, 0, 0, 0, 370, 371, 7, 5, 0, 0, 371, 41, 1, 0, 0,
		0, 372, 373, 7, 6, 0, 0, 373, 43, 1, 0, 0, 0, 374, 375, 5, 32, 0, 0, 375,
		45, 1, 0, 0, 0, 376, 377, 6, 23, -1, 0, 377, 378, 5, 36, 0, 0, 378, 379,
		3, 46, 23, 0, 379, 380, 5, 37, 0, 0, 380, 387, 1, 0, 0, 0, 381, 387, 5,
		5, 0, 0, 382, 387, 3, 58, 29, 0, 383, 387, 3, 62, 31, 0, 384, 385, 5, 19,
		0, 0, 385, 387, 3, 46, 23, 1, 386, 376, 1, 0, 0, 0, 386, 381, 1, 0, 0,
		0, 386, 382, 1, 0, 0, 0, 386, 383, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387,
		398, 1, 0, 0, 0, 388, 389, 10, 6, 0, 0, 389, 390, 3, 36, 18, 0, 390, 391,
		3, 46, 23, 7, 391, 397, 1, 0, 0, 0, 392, 393, 10, 5, 0, 0, 393, 394, 3,
		38, 19, 0, 394, 395, 3, 46, 23, 6, 395, 397, 1, 0, 0, 0, 396, 388, 1, 0,
		0, 0, 396, 392, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0,
		398, 399, 1, 0, 0, 0, 399, 47, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 402,
		6, 24, -1, 0, 402, 403, 5, 36, 0, 0, 403, 404, 3, 48, 24, 0, 404, 405,
		5, 37, 0, 0, 405, 412, 1, 0, 0, 0, 406, 407, 5, 32, 0, 0, 407, 412, 3,
		48, 24, 4, 408, 412, 5, 3, 0, 0, 409, 412, 3, 58, 29, 0, 410, 412, 3, 62,
		31, 0, 411, 401, 1, 0, 0, 0, 411, 406, 1, 0, 0, 0, 411, 408, 1, 0, 0, 0,
		411, 409, 1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 423, 1, 0, 0, 0, 413,
		414, 10, 6, 0, 0, 414, 415, 3, 40, 20, 0, 415, 416, 3, 48, 24, 7, 416,
		422, 1, 0, 0, 0, 417, 418, 10, 5, 0, 0, 418, 419, 3, 42, 21, 0, 419, 420,
		3, 48, 24, 6, 420, 422, 1, 0, 0, 0, 421, 413, 1, 0, 0, 0, 421, 417, 1,
		0, 0, 0, 422, 425, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0,
		0, 424, 49, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 427, 6, 25, -1, 0, 427,
		428, 3, 46, 23, 0, 428, 429, 3, 42, 21, 0, 429, 430, 3, 46, 23, 0, 430,
		441, 1, 0, 0, 0, 431, 432, 3, 48, 24, 0, 432, 433, 3, 40, 20, 0, 433, 434,
		3, 48, 24, 0, 434, 441, 1, 0, 0, 0, 435, 436, 3, 44, 22, 0, 436, 437, 5,
		36, 0, 0, 437, 438, 3, 50, 25, 0, 438, 439, 5, 37, 0, 0, 439, 441, 1, 0,
		0, 0, 440, 426, 1, 0, 0, 0, 440, 431, 1, 0, 0, 0, 440, 435, 1, 0, 0, 0,
		441, 448, 1, 0, 0, 0, 442, 443, 10, 2, 0, 0, 443, 444, 3, 40, 20, 0, 444,
		445, 3, 48, 24, 0, 445, 447, 1, 0, 0, 0, 446, 442, 1, 0, 0, 0, 447, 450,
		1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 51, 1, 0,
		0, 0, 450, 448, 1, 0, 0, 0, 451, 455, 3, 58, 29, 0, 452, 455, 3, 46, 23,
		0, 453, 455, 3, 48, 24, 0, 454, 451, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0,
		454, 453, 1, 0, 0, 0, 455, 53, 1, 0, 0, 0, 456, 457, 5, 42, 0, 0, 457,
		458, 5, 38, 0, 0, 458, 459, 3, 46, 23, 0, 459, 460, 5, 39, 0, 0, 460, 55,
		1, 0, 0, 0, 461, 462, 5, 42, 0, 0, 462, 463, 5, 41, 0, 0, 463, 477, 5,
		42, 0, 0, 464, 465, 5, 42, 0, 0, 465, 466, 5, 41, 0, 0, 466, 477, 3, 54,
		27, 0, 467, 468, 5, 42, 0, 0, 468, 469, 5, 41, 0, 0, 469, 477, 3, 56, 28,
		0, 470, 471, 3, 54, 27, 0, 471, 472, 5, 41, 0, 0, 472, 473, 3, 56, 28,
		0, 473, 477, 1, 0, 0, 0, 474, 477, 3, 54, 27, 0, 475, 477, 5, 42, 0, 0,
		476, 461, 1, 0, 0, 0, 476, 464, 1, 0, 0, 0, 476, 467, 1, 0, 0, 0, 476,
		470, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476, 475, 1, 0, 0, 0, 477, 57, 1,
		0, 0, 0, 478, 482, 5, 42, 0, 0, 479, 482, 3, 54, 27, 0, 480, 482, 3, 56,
		28, 0, 481, 478, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 481, 480, 1, 0, 0, 0,
		482, 59, 1, 0, 0, 0, 483, 488, 3, 52, 26, 0, 484, 485, 5, 40, 0, 0, 485,
		487, 3, 52, 26, 0, 486, 484, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486,
		1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 61, 1, 0, 0, 0, 490, 488, 1, 0,
		0, 0, 491, 492, 5, 42, 0, 0, 492, 494, 5, 36, 0, 0, 493, 495, 3, 60, 30,
		0, 494, 493, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496,
		497, 5, 37, 0, 0, 497, 63, 1, 0, 0, 0, 50, 67, 78, 85, 91, 108, 125, 128,
		133, 149, 156, 167, 174, 185, 192, 197, 205, 212, 223, 230, 241, 248, 253,
		264, 269, 277, 285, 291, 297, 308, 316, 322, 328, 335, 342, 352, 359, 364,
		386, 396, 398, 411, 421, 423, 440, 448, 454, 476, 481, 488, 494,
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
	FluxRULE_args                    = 30
	FluxRULE_function_call           = 31
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

	// IsProgramContext differentiates from other interfaces.
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57174604646484) != 0 {
		{
			p.SetState(64)
			p.Statement()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
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

	// IsStatementContext differentiates from other interfaces.
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Loop_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Func_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Return_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
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

	// IsExpressionContext differentiates from other interfaces.
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Var_assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Single_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Array_var_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
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
		p.SetState(87)
		p.Match(FluxL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57174604646484) != 0 {
		{
			p.SetState(88)
			p.Statement()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
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

	// IsLoop_statementContext differentiates from other interfaces.
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
		p.SetState(96)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(FluxLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.comparative_expression(0)
	}
	{
		p.SetState(99)
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

	// IsIf_statementContext differentiates from other interfaces.
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
		p.SetState(101)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(FluxIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.comparative_expression(0)
	}
	{
		p.SetState(104)
		p.Block()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(105)
			p.Match(FluxAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(FluxELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
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

	// IsReturn_statementContext differentiates from other interfaces.
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
		p.SetState(110)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(FluxRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
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

	// IsData_typeContext differentiates from other interfaces.
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
		p.SetState(114)
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

	// IsFunc_declarationContext differentiates from other interfaces.
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
		p.SetState(116)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(FluxFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(FluxL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxVAR_IDENTIFIER {
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

		for _la == FluxCOMMA {
			{
				p.SetState(121)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(122)
				p.Match(FluxVAR_IDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(130)
		p.Match(FluxR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxRETURN_TYPE {
		{
			p.SetState(131)
			p.Match(FluxRETURN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Data_type()
		}

	}
	{
		p.SetState(135)
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

	// IsVar_typeContext differentiates from other interfaces.
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
		p.SetState(137)
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

	// IsVar_nameContext differentiates from other interfaces.
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
		p.SetState(139)
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

	// IsVar_valueContext differentiates from other interfaces.
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
		p.SetState(141)
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

	// IsString_var_declarationContext differentiates from other interfaces.
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
		p.SetState(143)
		p.Match(FluxTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Var_name()
	}
	{
		p.SetState(145)
		p.Match(FluxL_BLOCK)
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

	for _la == FluxNEWLINE {
		{
			p.SetState(146)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(FluxTEXT)
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

	for _la == FluxNEWLINE {
		{
			p.SetState(153)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(159)
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

	// IsNumber_var_declarationContext differentiates from other interfaces.
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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Var_name()
		}
		{
			p.SetState(163)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(164)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(170)
			p.Match(FluxNUMBER)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(171)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(177)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Var_name()
		}
		{
			p.SetState(181)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(182)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(187)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(188)
			p.Math_expression()
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(189)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(195)
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

	// IsBoolean_var_declarationContext differentiates from other interfaces.
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

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Var_name()
		}
		{
			p.SetState(201)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(202)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(208)
			p.Match(FluxBOOLEAN)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(209)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(215)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Var_name()
		}
		{
			p.SetState(219)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(220)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.logical_expression(0)
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(227)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(233)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Var_name()
		}
		{
			p.SetState(237)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(238)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(244)
			p.comparative_expression(0)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(245)
				p.Match(FluxNEWLINE)
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
		}
		{
			p.SetState(251)
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

	// IsSingle_var_declarationContext differentiates from other interfaces.
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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.String_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.Number_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(257)
			p.Boolean_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(258)
			p.Var_type()
		}
		{
			p.SetState(259)
			p.Var_name()
		}
		{
			p.SetState(260)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(261)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(267)
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

	// IsArray_var_declarationContext differentiates from other interfaces.
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

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Var_type()
		}
		{
			p.SetState(272)
			p.Var_name()
		}
		{
			p.SetState(273)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(274)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(279)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(280)
			p.Var_value()
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(281)
				p.Match(FluxCOMMA)
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

			for _la == FluxNEWLINE {
				{
					p.SetState(282)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(287)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(288)
				p.Var_value()
			}

			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(294)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(300)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.Var_type()
		}
		{
			p.SetState(303)
			p.Var_name()
		}
		{
			p.SetState(304)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(305)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(311)
			p.Math_expression()
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(312)
				p.Match(FluxCOMMA)
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

			for _la == FluxNEWLINE {
				{
					p.SetState(313)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(318)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(319)
				p.Math_expression()
			}

			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(325)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
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

	// IsVar_assignmentContext differentiates from other interfaces.
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

	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Var_name()
		}
		{
			p.SetState(331)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(332)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(338)
			p.Var_value()
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(339)
				p.Match(FluxNEWLINE)
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
			p.SetState(345)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Var_name()
		}
		{
			p.SetState(348)
			p.Match(FluxL_BLOCK)
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

		for _la == FluxNEWLINE {
			{
				p.SetState(349)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(355)
			p.Math_expression()
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(356)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(362)
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

	// IsOp_level1Context differentiates from other interfaces.
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
		p.SetState(366)
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

	// IsOp_level2Context differentiates from other interfaces.
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
		p.SetState(368)
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

	// IsOp_level3Context differentiates from other interfaces.
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
		p.SetState(370)
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

	// IsOp_level4Context differentiates from other interfaces.
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
		p.SetState(372)
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

	// IsOp_level5Context differentiates from other interfaces.
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
		p.SetState(374)
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

	// IsNumeric_expressionContext differentiates from other interfaces.
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
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(377)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.numeric_expression(0)
		}
		{
			p.SetState(379)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(381)
			p.Match(FluxNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(382)
			p.Get_var()
		}

	case 4:
		{
			p.SetState(383)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(384)
			p.Match(FluxOP_SUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.numeric_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(398)
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
			p.SetState(396)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumeric_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numeric_expression)
				p.SetState(388)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(389)
					p.Op_level1()
				}
				{
					p.SetState(390)
					p.numeric_expression(7)
				}

			case 2:
				localctx = NewNumeric_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numeric_expression)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(393)
					p.Op_level2()
				}
				{
					p.SetState(394)
					p.numeric_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(400)
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

	// IsLogical_expressionContext differentiates from other interfaces.
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
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(402)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.logical_expression(0)
		}
		{
			p.SetState(404)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(406)
			p.Match(FluxOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(408)
			p.Match(FluxBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(409)
			p.Get_var()
		}

	case 5:
		{
			p.SetState(410)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(423)
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
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(414)
					p.Op_level3()
				}
				{
					p.SetState(415)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(418)
					p.Op_level4()
				}
				{
					p.SetState(419)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(425)
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

	// IsComparative_expressionContext differentiates from other interfaces.
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
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(427)
			p.numeric_expression(0)
		}
		{
			p.SetState(428)
			p.Op_level4()
		}
		{
			p.SetState(429)
			p.numeric_expression(0)
		}

	case 2:
		{
			p.SetState(431)
			p.logical_expression(0)
		}
		{
			p.SetState(432)
			p.Op_level3()
		}
		{
			p.SetState(433)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(435)
			p.Op_level5()
		}
		{
			p.SetState(436)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.comparative_expression(0)
		}
		{
			p.SetState(438)
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
	p.SetState(448)
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
			p.SetState(442)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(443)
				p.Op_level3()
			}
			{
				p.SetState(444)
				p.logical_expression(0)
			}

		}
		p.SetState(450)
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

	// IsMath_expressionContext differentiates from other interfaces.
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
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(451)
			p.Get_var()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(452)
			p.numeric_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(453)
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

	// IsGet_array_elementContext differentiates from other interfaces.
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
		p.SetState(456)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(FluxL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.numeric_expression(0)
	}
	{
		p.SetState(459)
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

	// IsGet_childContext differentiates from other interfaces.
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
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(467)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(470)
			p.Get_array_element()
		}
		{
			p.SetState(471)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(474)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(475)
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

	// IsGet_varContext differentiates from other interfaces.
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
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(478)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(479)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(480)
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

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllMath_expression() []IMath_expressionContext {
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

func (s *ArgsContext) Math_expression(i int) IMath_expressionContext {
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

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FluxRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
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
	Args() IArgsContext

	// IsFunction_callContext differentiates from other interfaces.
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

func (s *Function_callContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
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
	p.EnterRule(localctx, 62, FluxRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.Match(FluxL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4471061479464) != 0 {
		{
			p.SetState(493)
			p.Args()
		}

	}
	{
		p.SetState(496)
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
