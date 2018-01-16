// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
	"reflect"
	"strings"
)

// internalPatchIdentity is the Identity of an patch.
var internalPatchIdentity = Identity{
	Name:     "__internal_patch__",
	Category: "__internal_patch__",
}

// PatchType represents the mode of an operation.
// This is informational. It is up to your implementation to honor or not
// the PatchType.
type PatchType uint16

const (

	// PatchTypeSet sets the patch data no matter what are their current values.
	PatchTypeSet PatchType = iota

	// PatchTypeSetIfZero sets the patch data if their current values are empty.
	PatchTypeSetIfZero
)

// PatchData contains data to apply.
type PatchData map[string]interface{}

// An Patch represents an additive operation to apply to one or more Identifiables.
type Patch struct {
	Data PatchData `json:"data"`
	Type PatchType `json:"type"`
}

// NewPatch returns a new Patch.
func NewPatch(typ PatchType, data PatchData) *Patch {

	return &Patch{
		Data: data,
		Type: typ,
	}
}

// Identity returns the Identity of the of the receiver.
//
// In that case it will return a the private type indentity.
func (a *Patch) Identity() Identity {

	return internalPatchIdentity
}

// Identifier returns the unique identifier of the of the receiver.
//
// In that case, it will return the string "__internal__".
func (a *Patch) Identifier() string {

	return "__internal__"
}

// Version returns the version of the patch.
func (a *Patch) Version() int {

	return 1
}

// SetIdentifier sets the unique identifier of the of the receiver.
//
// In that case it does nothing.
func (a *Patch) SetIdentifier(string) {}

func (a *Patch) String() string {

	return fmt.Sprintf("<patch type:%d data:%v>", a.Type, a.Data)
}

// Validate validates the current information stored into the Patch.
func (a *Patch) Validate() error {

	errors := Errors{}

	if err := ValidateIntInList("type", int(a.Type), []int{0, 1}); err != nil {
		errors = append(errors, err.(Error))
	}

	return errors
}

// Apply applies the patch the the given AttributeSpecifiable.
// obj must be an pointer to an AttributeSpecifiable or Apply will panic.
func (a *Patch) Apply(obj AttributeSpecifiable) error {

	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr {
		panic("A pointer to elemental.AttributeSpecifiable must be passed to Apply")
	}
	objValue = objValue.Elem()

	for k, v := range a.Data {

		f := objValue.FieldByName(obj.SpecificationForAttribute(strings.ToLower(k)).ConvertedName)

		if !f.IsValid() {
			return fmt.Errorf("field '%s' is invalid", k)
		}

		if !f.CanSet() {
			return fmt.Errorf("field '%s' cannot be set", k)
		}

		vValue := reflect.ValueOf(v)
		f.Set(vValue)
	}

	return nil
}
