package vm

import (
	"context"
	"github.com/Runway-Club/flux_lang/parsing"
	"github.com/Runway-Club/flux_lang/shared"
	"github.com/Runway-Club/flux_lang/vm/core"
	"github.com/Runway-Club/flux_lang/vm/io"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"time"
)

type FluxVirtualMachine struct {
	parser   *FluxProgramParser
	varTable *core.VarTable
}

func NewFluxVirtualMachine() *FluxVirtualMachine {
	return &FluxVirtualMachine{}
}

func (f *FluxVirtualMachine) Execute(params *shared.ExecutionParams) shared.ExecutionResult {
	// create error collector
	errorCollector := io.NewBaseErrorCollector()

	elapsedTime := time.Now().UnixMilli()
	// read file at params.EntryPoint
	// get absolute path of entry point from relative path
	absPath, err := os.Getwd()
	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			ErrorCollector: errorCollector,
			ElapsedTime:    elapsedTime,
		}
	}
	mainFileData := ""
	if params.EntryPoint != "" {
		params.EntryPoint = absPath + "/" + params.EntryPoint
		// read file

		mainFileBytes, err := os.ReadFile(params.EntryPoint)
		if err != nil {
			elapsedTime = time.Now().UnixMilli() - elapsedTime
			return shared.ExecutionResult{
				ErrorCollector: errorCollector,
				ElapsedTime:    elapsedTime,
			}
		}
		mainFileData = string(mainFileBytes)
	} else {
		mainFileData = params.SourceCode
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

	varTable := core.NewVarTable()
	f.varTable = varTable
	f.parser = NewFluxProgramParser(logger, errorCollector, varTable)
	// add traverser to parser
	parser.AddParseListener(f.parser)
	// start parsing
	parser.Program()

	executionCtx := statements.NewExecutionContext(context.TODO(), varTable)

	except := f.parser.GetProgram().Execute(executionCtx)
	// return result
	elapsedTime = time.Now().UnixMilli() - elapsedTime

	return shared.ExecutionResult{
		ErrorCollector:   errorCollector,
		ElapsedTime:      elapsedTime,
		RuntimeException: except,
	}
}

func (f *FluxVirtualMachine) GetVarTable() *core.VarTable {
	return f.varTable
}
