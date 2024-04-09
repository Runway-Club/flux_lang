package core

import (
	"github.com/Runway-Club/flux_lang/vm/common"
	"github.com/Runway-Club/flux_lang/vm/exception"
)

type VarTableEntry struct {
	Name     string
	Type     string
	RawValue string
}

type VarTableGeneric[T any] map[string]common.Identifiable[T]

type VarTable struct {
	mainTable map[string]*VarTableEntry
	numTable  *VarTableGeneric[float64]
	strTable  *VarTableGeneric[string]
	boolTable *VarTableGeneric[bool]
}

func NewVarTable() *VarTable {
	return &VarTable{
		mainTable: make(map[string]*VarTableEntry),
		numTable:  &VarTableGeneric[float64]{},
		strTable:  &VarTableGeneric[string]{},
		boolTable: &VarTableGeneric[bool]{},
	}
}

func (vt *VarTable) Set(name string, type_ string, rawValue string) *exception.BaseException {
	if vt.Exists(name) {
		return &exception.BaseException{Err: ErrVarExists, MessageFmt: "Variable %s already exists", Args: []interface{}{name}}
	}
	vt.mainTable[name] = &VarTableEntry{name, type_, rawValue}
	return nil
}

func (vt *VarTable) Get(name string) (*VarTableEntry, *exception.BaseException) {
	if !vt.Exists(name) {
		return nil, &exception.BaseException{
			Err:        ErrVarNotFound,
			MessageFmt: "Variable %s not found",
			Args:       []interface{}{name},
		}
	}
	return vt.mainTable[name], nil
}

func (vt *VarTable) Exists(name string) bool {
	_, exists := vt.mainTable[name]
	return exists
}

func (vt *VarTable) Remove(name string) {
	delete(vt.mainTable, name)
}

func (vt *VarTable) SetNum(name string, value float64) *exception.BaseException {
	// set to main table as well
	err := vt.Set(name, VarTypeNum, "")
	if err != nil {
		return err
	}
	(*vt.numTable)[name] = common.NewNumber(name, value)
	return nil
}

func (vt *VarTable) GetNumValue(name string) (float64, *exception.BaseException) {
	if !vt.Exists(name) {
		return 0, &exception.BaseException{Err: ErrVarNotFound, MessageFmt: "Variable %s not found", Args: []interface{}{name}}
	}
	return (*vt.numTable)[name].GetValue(), nil
}

func (vt *VarTable) SetText(name string, value string) *exception.BaseException {
	// set to main table as well
	err := vt.Set(name, VarTypeText, "")
	if err != nil {
		return err

	}
	(*vt.strTable)[name] = common.NewText(name, value)
	return nil
}

func (vt *VarTable) GetTextValue(name string) (string, *exception.BaseException) {
	if !vt.Exists(name) {
		return "", &exception.BaseException{Err: ErrVarNotFound, MessageFmt: "Variable %s not found", Args: []interface{}{name}}
	}
	return (*vt.strTable)[name].GetValue(), nil
}

func (vt *VarTable) SetBool(name string, value bool) *exception.BaseException {
	// set to main table as well
	err := vt.Set(name, VarTypeBool, "")
	if err != nil {
		return err
	}
	(*vt.boolTable)[name] = common.NewBoolean(name, value)
	return nil
}

func (vt *VarTable) GetBoolValue(name string) (bool, *exception.BaseException) {
	if !vt.Exists(name) {
		return false,
			&exception.BaseException{Err: ErrVarNotFound, MessageFmt: "Variable %s not found", Args: []interface{}{name}}
	}
	return (*vt.boolTable)[name].GetValue(), nil
}
