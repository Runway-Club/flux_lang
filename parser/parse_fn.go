package parser

import (
	fluxIO "github.com/Runway-Club/flux_lang/io"
	"github.com/Runway-Club/flux_lang/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func Parse(data string, errorCollector fluxIO.ErrorCollector, logger fluxIO.Logger) *FluxProgramParser {

	input := antlr.NewInputStream(string(data))
	// create lexer
	lexer := parsing.NewPrimitives(input)
	// create parser
	parser := parsing.NewFlux(antlr.NewCommonTokenStream(lexer, 0))
	// create traverser
	fluxProgramParser := NewFluxProgramParser(logger, errorCollector)
	// add traverser to parser
	parser.AddParseListener(fluxProgramParser)
	// start parsing
	parser.Program()

	return fluxProgramParser
}
