package testmodel

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// ListIdentity represents the Identity of the object
var ListIdentity = elemental.Identity{
	Name:     "list",
	Category: "lists",
}

// ListsList represents a list of Lists
type ListsList []*List

// ContentIdentity returns the identity of the objects in the list.
func (o ListsList) ContentIdentity() elemental.Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the ListsList.
func (o ListsList) Copy() elemental.ContentIdentifiable {

	copy := append(ListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ListsList.
func (o ListsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*List))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ListsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ListsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ListsList) Version() int {

	return 1.0
}

// List represents the model of a list
type List struct {
	// The identifier
	ID string `json:"ID" bson:"_id"`

	// A creation only only attribute
	CreationOnly string `json:"creationOnly" bson:"creationonly"`

	// A Date
	Date time.Time `json:"-" bson:"date"`

	// The description
	Description string `json:"description" bson:"description"`

	// The name
	Name string `json:"name" bson:"name"`

	// The identifier of the parent of the object
	ParentID string `json:"parentID" bson:"parentid"`

	// The type of the parent of the object
	ParentType string `json:"parentType" bson:"parenttype"`

	// A read only attribute
	ReadOnly string `json:"readOnly" bson:"readonly"`

	// A Slice
	Slice []string `json:"-" bson:"slice"`

	// An unexposed attribute
	Unexposed string `json:"-" bson:"unexposed"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewList returns a new *List
func NewList() *List {

	return &List{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *List) Identity() elemental.Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *List) Version() int {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *List) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *List) Doc() string {
	return `Represent a a list of task to do.`
}

func (o *List) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreationOnly returns the creationOnly of the receiver
func (o *List) GetCreationOnly() string {
	return o.CreationOnly
}

// GetDate returns the date of the receiver
func (o *List) GetDate() time.Time {
	return o.Date
}

// GetName returns the name of the receiver
func (o *List) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *List) SetName(name string) {
	o.Name = name
}

// GetReadOnly returns the readOnly of the receiver
func (o *List) GetReadOnly() string {
	return o.ReadOnly
}

// GetSlice returns the slice of the receiver
func (o *List) GetSlice() []string {
	return o.Slice
}

// GetUnexposed returns the unexposed of the receiver
func (o *List) GetUnexposed() string {
	return o.Unexposed
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("creationOnly", o.CreationOnly); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("creationOnly", o.CreationOnly); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("readOnly", o.ReadOnly); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("readOnly", o.ReadOnly); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*List) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*List) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ListAttributesMap
}

// ListAttributesMap represents the map of attribute for List.
var ListAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The identifier`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"CreationOnly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A creation only only attribute`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "creationOnly",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A Date`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "date",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
		Unique:         true,
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The description`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The identifier of the parent of the object`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The type of the parent of the object`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentType",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"ReadOnly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `A read only attribute`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Slice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A Slice`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "slice",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
		Unique:         true,
	},
	"Unexposed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `An unexposed attribute`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "unexposed",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
}

// ListLowerCaseAttributesMap represents the map of attribute for List.
var ListLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The identifier`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"creationonly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A creation only only attribute`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "creationOnly",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A Date`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "date",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
		Unique:         true,
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The description`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The identifier of the parent of the object`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"parenttype": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The type of the parent of the object`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentType",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"readonly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `A read only attribute`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"slice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `A Slice`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "slice",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
		Unique:         true,
	},
	"unexposed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `An unexposed attribute`,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "unexposed",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
}
