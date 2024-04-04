package vm

import (
	"github.com/Runway-Club/flux_lang/parsing"
	"github.com/Runway-Club/flux_lang/shared"
	"github.com/Runway-Club/flux_lang/vm/io"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"time"
)

type FluxVirtualMachine struct {
	traverser parsing.FluxListener
	varTable  *statements.VarTable
}

func NewFluxVirtualMachine() *FluxVirtualMachine {
	return &FluxVirtualMachine{}
}

func (f *FluxVirtualMachine) Execute(params *shared.ExecutionParams) shared.ExecutionResult {
	elapsedTime := time.Now().UnixMilli()
	// read file at params.EntryPoint
	// get absolute path of entry point from relative path
	absPath, err := os.Getwd()
	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			Error:       err.Error(),
			ElapsedTime: elapsedTime,
		}
	}
	params.EntryPoint = absPath + "/" + params.EntryPoint
	// read file

	mainFileData, err := os.ReadFile(params.EntryPoint)
	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			Error:       err.Error(),
			ElapsedTime: elapsedTime,
		}
	}
	input := antlr.NewInputStream(string(mainFileData))
	// create lexer
	lexer := parsing.NewPrimitives(input)
	// create parser
	parser := parsing.NewFlux(antlr.NewCommonTokenStream(lexer, 0))
	// create traverser
	// create logger
	var logger io.Logger = io.NewEmptyLogger()
	if params.Verbose {
		logger = io.NewBaseLogger()
	}

	varTable := statements.NewVarTable()
	f.varTable = varTable
	f.traverser = NewFluxTraverser(logger, varTable)
	// add traverser to parser
	parser.AddParseListener(f.traverser)
	// start parsing
	parser.Program()
	// return result
	elapsedTime = time.Now().UnixMilli() - elapsedTime
	return shared.ExecutionResult{
		Error:       "",
		ElapsedTime: elapsedTime,
	}
}

func (f *FluxVirtualMachine) GetVarTable() *statements.VarTable {
	return f.varTable
}
