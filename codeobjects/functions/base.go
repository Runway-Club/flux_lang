package functions

import (
	"github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/common"
	"github.com/Runway-Club/flux_lang/exception"
)

type BaseFunction struct {
	*codeobjects.BaseStatement
	Args       *Args
	ReturnType common.FluxType
}

func (b BaseFunction) Generate(ctx *codeobjects.GenerateContext) string {
	return ""
}

func (b BaseFunction) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	return nil
}
