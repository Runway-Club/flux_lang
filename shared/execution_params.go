package shared

import (
	"github.com/Runway-Club/flux_lang/exception"
	"github.com/Runway-Club/flux_lang/io"
)

type ExecutionParams struct {
	EntryPoint string `json:"entryPoint"`
	Verbose    bool   `json:"verbose"`
	SourceCode string `json:"sourceCode"`
}

type ExecutionResult struct {
	ErrorCollector   io.ErrorCollector
	ElapsedTime      int64                    `json:"elapsedTime"`
	RuntimeException *exception.BaseException `json:"runtimeException"`
}
