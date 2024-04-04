package tests

import (
	"github.com/Runway-Club/flux_lang/shared"
	"github.com/Runway-Club/flux_lang/vm"
	"testing"
)

func TestVM_VarDeclaration(t *testing.T) {
	t.Run("create number variables", func(t *testing.T) {
		myVM := vm.NewFluxVirtualMachine()
		result := myVM.Execute(&shared.ExecutionParams{
			EntryPoint: "var_declaration_test.flux",
		})

		if result.Error != "" {
			t.Errorf("Error: %v", result.Error)
		}

		if result.ElapsedTime > 1000 {
			t.Errorf("Execution time too long: %v", result.ElapsedTime)
		}

		if !myVM.GetVarTable().Exists("a") {
			t.Errorf("Variable a not found")
		}

		if !myVM.GetVarTable().Exists("pi") {
			t.Errorf("Variable pi not found")
		}

		aValue, err := myVM.GetVarTable().GetNumValue("a")
		if err != nil {
			t.Errorf("Error getting variable a: %v", err)
		}

		if aValue != 2 {
			t.Errorf("Variable a has wrong value: %v", aValue)
		}

		piValue, err := myVM.GetVarTable().GetNumValue("pi")
		if err != nil {
			t.Errorf("Error getting variable pi: %v", err)
		}

		if piValue != 3 {
			t.Errorf("Variable pi has wrong value: %v", piValue)
		}
	})
}
