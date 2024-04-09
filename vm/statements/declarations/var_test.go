package declarations

import (
	"context"
	"github.com/Runway-Club/flux_lang/vm/core"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/Runway-Club/flux_lang/vm/statements/expression"
	"github.com/Runway-Club/flux_lang/vm/statements/expression/operators"
	"testing"
)

func TestVarDeclaration_Execute(t *testing.T) {
	varTable := core.NewVarTable()
	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", FluxTypeNumber, "",
			expression.NewMathExpression(1, 1, 1,
				expression.NewNumericExpression(1, 1, 1,
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(2)),
					operators.NewAdd(1, 1, 1, nil, nil),
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(3)), 0),
			),
		)
		ctx := statements.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 5 {
			t.Errorf("Expected 5, got %v", testVar)
		}
	})
}
