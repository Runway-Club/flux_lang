package statements

import (
	"errors"
	"github.com/Runway-Club/flux_lang/vm/common"
)

const (
	VarTypeNum  = "num"
	VarTypeText = "text"
	VarTypeBool = "bool"
)

var (
	ErrVarNotFound = errors.New("variable not found")
	ErrVarExists   = errors.New("variable already exists")
)

type VarTableEntry struct {
	Name     string
	Type     string
	RawValue string
}

type VarTableGeneric[T any] map[string]common.Identifiable[T]

type VarTable struct {
	mainTable map[string]*VarTableEntry
	numTable  *VarTableGeneric[int]
	strTable  *VarTableGeneric[string]
	boolTable *VarTableGeneric[bool]
}

func NewVarTable() *VarTable {
	return &VarTable{
		mainTable: make(map[string]*VarTableEntry),
		numTable:  &VarTableGeneric[int]{},
		strTable:  &VarTableGeneric[string]{},
		boolTable: &VarTableGeneric[bool]{},
	}
}

func (vt *VarTable) Set(name string, type_ string, rawValue string) *common.InternalError {
	if vt.Exists(name) {
		return &common.InternalError{Err: ErrVarExists, Factor: name}
	}
	vt.mainTable[name] = &VarTableEntry{name, type_, rawValue}
	return nil
}

func (vt *VarTable) Get(name string) (*VarTableEntry, *common.InternalError) {
	if !vt.Exists(name) {
		return nil, &common.InternalError{Err: ErrVarNotFound, Factor: name}
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

func (vt *VarTable) SetNum(name string, value int) *common.InternalError {
	// set to main table as well
	err := vt.Set(name, VarTypeNum, "")
	if err != nil {
		return err
	}
	(*vt.numTable)[name] = common.NewNumber(name, value)
	return nil
}

func (vt *VarTable) GetNumValue(name string) (int, *common.InternalError) {
	if !vt.Exists(name) {
		return 0, &common.InternalError{Err: ErrVarNotFound, Factor: name}
	}
	return (*vt.numTable)[name].GetValue(), nil
}

func (vt *VarTable) SetText(name string, value string) *common.InternalError {
	// set to main table as well
	err := vt.Set(name, VarTypeText, "")
	if err != nil {
		return err

	}
	(*vt.strTable)[name] = common.NewText(name, value)
	return nil
}

func (vt *VarTable) GetTextValue(name string) (string, *common.InternalError) {
	if !vt.Exists(name) {
		return "", &common.InternalError{Err: ErrVarNotFound, Factor: name}
	}
	return (*vt.strTable)[name].GetValue(), nil
}

func (vt *VarTable) SetBool(name string, value bool) *common.InternalError {
	// set to main table as well
	err := vt.Set(name, VarTypeBool, "")
	if err != nil {
		return err
	}
	(*vt.boolTable)[name] = common.NewBoolean(name, value)
	return nil
}

func (vt *VarTable) GetBoolValue(name string) (bool, *common.InternalError) {
	if !vt.Exists(name) {
		return false, &common.InternalError{Err: ErrVarNotFound, Factor: name}
	}
	return (*vt.boolTable)[name].GetValue(), nil
}
