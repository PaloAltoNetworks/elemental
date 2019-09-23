package testmodel

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RootIdentity represents the Identity of the object.
var RootIdentity = elemental.Identity{
	Name:     "root",
	Category: "root",
	Package:  "todo-list",
	Private:  false,
}

// Root represents the model of a root
type Root struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRoot returns a new *Root
func NewRoot() *Root {

	return &Root{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Root) Identity() elemental.Identity {

	return RootIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Root) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Root) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Root) GetBSON() (interface{}, error) {

	s := &mongoAttributesRoot{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Root) SetBSON(raw bson.Raw) error {

	s := &mongoAttributesRoot{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Root) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Root) BleveType() string {

	return "root"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Root) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Root) Doc() string {

	return `Root object of the API.`
}

func (o *Root) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// DeepCopy returns a deep copy if the Root.
func (o *Root) DeepCopy() *Root {

	if o == nil {
		return nil
	}

	out := &Root{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Root.
func (o *Root) DeepCopyInto(out *Root) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Root: %s", err))
	}

	*out = *target.(*Root)
}

// Validate valides the current information stored into the structure.
func (o *Root) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Root) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RootAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RootLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Root) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RootAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Root) ValueForAttribute(name string) interface{} {

	switch name {
	}

	return nil
}

// RootAttributesMap represents the map of attribute for Root.
var RootAttributesMap = map[string]elemental.AttributeSpecification{}

// RootLowerCaseAttributesMap represents the map of attribute for Root.
var RootLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}

type mongoAttributesRoot struct {
}
