package declarations

import (
	"fmt"
	codeobjects2 "github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/codeobjects/expression"
	"github.com/Runway-Club/flux_lang/exception"
	"strconv"
)

type FluxType string

const (
	FluxTypeNumber FluxType = "num"
	FluxTypeString FluxType = "text"
	FluxTypeBool   FluxType = "bool"
)

type VarDeclaration struct {
	*codeobjects2.BaseStatement
	Name     string
	Type     FluxType
	RawValue string
	Expr     *expression.MathExpression
}

func (v VarDeclaration) Generate(ctx *codeobjects2.GenerateContext) string {
	// to golang type
	varType := "double"
	if v.Type == FluxTypeString {
		varType = "string"
	} else if v.Type == FluxTypeBool {
		varType = "bool"
	} else if v.Type == FluxTypeNumber {
		varType = "double"
	}
	if v.Expr != nil {
		return fmt.Sprintf("%v %v = %v;", varType, v.Name, v.Expr.Generate(ctx))
	} else if v.RawValue != "" {
		return fmt.Sprintf("%v %v = %v;", varType, v.Name, v.RawValue)
	}
	return fmt.Sprintf("%v %v;", v.Name, varType)
}

func (v VarDeclaration) Execute(ctx *codeobjects2.ExecutionContext) *exception.BaseException {
	exprCtx := ctx.Clone()
	if v.Expr != nil {
		err := v.Expr.Execute(exprCtx)
		if err != nil {
			return err
		}
		if v.Type == FluxTypeNumber {
			err := ctx.VarTable.SetNum(v.Name, exprCtx.NumericValue)
			if err != nil {
				return err
			}
		}
		if v.Type == FluxTypeString {
			err := ctx.VarTable.SetText(v.Name, exprCtx.TextValue)
			if err != nil {
				return err
			}
		}
		if v.Type == FluxTypeBool {
			err := ctx.VarTable.SetBool(v.Name, exprCtx.BoolValue)
			if err != nil {
				return err
			}
		}
	} else if v.RawValue != "" {
		if v.Type == FluxTypeNumber {
			value, parseErr := strconv.ParseFloat(v.RawValue, 64)
			if parseErr != nil {
				return &exception.BaseException{
					MessageFmt: "Invalid number value: %s",
					Args:       []interface{}{v.RawValue},
					Line:       v.GetLine(),
					StartPos:   v.GetStartPos(),
					EndPos:     v.GetEndPos(),
				}
			}
			err := ctx.VarTable.SetNum(v.Name, value)
			if err != nil {
				return err
			}
		}
		if v.Type == FluxTypeString {
			err := ctx.VarTable.SetText(v.Name, v.RawValue)
			if err != nil {
				return err
			}
		}
		if v.Type == FluxTypeBool {
			value, parseErr := strconv.ParseBool(v.RawValue)
			if parseErr != nil {
				return &exception.BaseException{
					MessageFmt: "Invalid boolean value: %s",
					Args:       []interface{}{v.RawValue},
					Line:       v.GetLine(),
					StartPos:   v.GetStartPos(),
					EndPos:     v.GetEndPos(),
				}
			}
			err := ctx.VarTable.SetBool(v.Name, value)
			if err != nil {
				return err
			}
		}
	} else {
		// Set default value
		if v.Type == FluxTypeNumber {
			err := ctx.VarTable.SetNum(v.Name, 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewVarDeclaration(line int, startPos int, endPos int, name string, type_ FluxType, rawValue string, expr *expression.MathExpression) *VarDeclaration {
	return &VarDeclaration{BaseStatement: &codeobjects2.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Name: name, Type: type_, RawValue: rawValue, Expr: expr}
}
