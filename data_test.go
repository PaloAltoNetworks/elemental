package elemental

import (
	"fmt"
	"time"

    "github.com/globalsign/mgo/bson"
    "github.com/mitchellh/copystructure"
)

//lint:file-ignore U1000 auto generated code.
// ListIdentity represents the Identity of the object.
var ListIdentity = Identity{
	Name:     "list",
	Category: "lists",
	Package:  "todo-list",
	Private:  false,
}

// ListsList represents a list of Lists
type ListsList []*List

// Identity returns the identity of the objects in the list.
func (o ListsList) Identity() Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the ListsList.
func (o ListsList) Copy() Identifiables {

	copy := append(ListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ListsList.
func (o ListsList) Append(objects ...Identifiable) Identifiables {

	out := append(ListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*List))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o ListsList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ListsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ListsList converted to SparseListsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ListsList) ToSparse(fields ...string) Identifiables {

	out := make(SparseListsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseList)
	}

	return out
}

// Version returns the version of the content.
func (o ListsList) Version() int {

	return 1
}

// List represents the model of a list
type List struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly string `json:"creationOnly" msgpack:"creationOnly" bson:"creationonly" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	// The description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly string `json:"readOnly" msgpack:"readOnly" bson:"readonly" mapstructure:"readOnly,omitempty"`

	// This attribute is secret.
	Secret string `json:"secret" msgpack:"secret" bson:"secret" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice []string `json:"slice" msgpack:"slice" bson:"slice" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed string `json:"-" msgpack:"-" bson:"unexposed" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{
		ModelVersion: 1,
		Slice:        []string{},
	}
}

// Identity returns the Identity of the object.
func (o *List) Identity() Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesList{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreationOnly = o.CreationOnly
	s.Date = o.Date
	s.Description = o.Description
	s.Name = o.Name
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.ReadOnly = o.ReadOnly
	s.Secret = o.Secret
	s.Slice = o.Slice
	s.Unexposed = o.Unexposed

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreationOnly = s.CreationOnly
	o.Date = s.Date
	o.Description = s.Description
	o.Name = s.Name
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.ReadOnly = s.ReadOnly
	o.Secret = s.Secret
	o.Slice = s.Slice
	o.Unexposed = s.Unexposed

	return nil
}

// Version returns the hardcoded version of the model.
func (o *List) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *List) BleveType() string {

	return "list"
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

// GetName returns the Name of the receiver.
func (o *List) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *List) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *List) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseList{
			ID:           &o.ID,
			CreationOnly: &o.CreationOnly,
			Date:         &o.Date,
			Description:  &o.Description,
			Name:         &o.Name,
			ParentID:     &o.ParentID,
			ParentType:   &o.ParentType,
			ReadOnly:     &o.ReadOnly,
			Secret:       &o.Secret,
			Slice:        &o.Slice,
			Unexposed:    &o.Unexposed,
		}
	}

	sp := &SparseList{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "creationOnly":
			sp.CreationOnly = &(o.CreationOnly)
		case "date":
			sp.Date = &(o.Date)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "readOnly":
			sp.ReadOnly = &(o.ReadOnly)
		case "secret":
			sp.Secret = &(o.Secret)
		case "slice":
			sp.Slice = &(o.Slice)
		case "unexposed":
			sp.Unexposed = &(o.Unexposed)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseList to the object.
func (o *List) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseList)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreationOnly != nil {
		o.CreationOnly = *so.CreationOnly
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.ReadOnly != nil {
		o.ReadOnly = *so.ReadOnly
	}
	if so.Secret != nil {
		o.Secret = *so.Secret
	}
	if so.Slice != nil {
		o.Slice = *so.Slice
	}
	if so.Unexposed != nil {
		o.Unexposed = *so.Unexposed
	}
}

// DeepCopy returns a deep copy if the List.
func (o *List) DeepCopy() *List {

	if o == nil {
		return nil
	}

	out := &List{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *List.
func (o *List) DeepCopyInto(out *List) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy List: %s", err))
	}

	*out = *target.(*List)
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*List) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := ListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*List) AttributeSpecifications() map[string]AttributeSpecification {

	return ListAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *List) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "creationOnly":
		return o.CreationOnly
	case "date":
		return o.Date
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "readOnly":
		return o.ReadOnly
	case "secret":
		return o.Secret
	case "slice":
		return o.Slice
	case "unexposed":
		return o.Unexposed
	}

	return nil
}

// ListAttributesMap represents the map of attribute for List.
var ListAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreationOnly": {
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ReadOnly": {
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"Slice": {
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Unexposed": {
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// ListLowerCaseAttributesMap represents the map of attribute for List.
var ListLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"creationonly": {
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"readonly": {
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"slice": {
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"unexposed": {
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseListsList represents a list of SparseLists
type SparseListsList []*SparseList

// Identity returns the identity of the objects in the list.
func (o SparseListsList) Identity() Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the SparseListsList.
func (o SparseListsList) Copy() Identifiables {

	copy := append(SparseListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseListsList.
func (o SparseListsList) Append(objects ...Identifiable) Identifiables {

	out := append(SparseListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseList))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseListsList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseListsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseListsList converted to ListsList.
func (o SparseListsList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseListsList) Version() int {

	return 1
}

// SparseList represents the sparse version of a list.
type SparseList struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly *string `json:"creationOnly,omitempty" msgpack:"creationOnly,omitempty" bson:"creationonly,omitempty" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"date,omitempty" mapstructure:"date,omitempty"`

	// The description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The name.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly *string `json:"readOnly,omitempty" msgpack:"readOnly,omitempty" bson:"readonly,omitempty" mapstructure:"readOnly,omitempty"`

	// This attribute is secret.
	Secret *string `json:"secret,omitempty" msgpack:"secret,omitempty" bson:"secret,omitempty" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice *[]string `json:"slice,omitempty" msgpack:"slice,omitempty" bson:"slice,omitempty" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed *string `json:"-" msgpack:"-" bson:"unexposed,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseList returns a new  SparseList.
func NewSparseList() *SparseList {
	return &SparseList{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseList) Identity() Identity {

	return ListIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseList) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseList) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseList{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreationOnly != nil {
		s.CreationOnly = o.CreationOnly
	}
	if o.Date != nil {
		s.Date = o.Date
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.ReadOnly != nil {
		s.ReadOnly = o.ReadOnly
	}
	if o.Secret != nil {
		s.Secret = o.Secret
	}
	if o.Slice != nil {
		s.Slice = o.Slice
	}
	if o.Unexposed != nil {
		s.Unexposed = o.Unexposed
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreationOnly != nil {
		o.CreationOnly = s.CreationOnly
	}
	if s.Date != nil {
		o.Date = s.Date
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.ReadOnly != nil {
		o.ReadOnly = s.ReadOnly
	}
	if s.Secret != nil {
		o.Secret = s.Secret
	}
	if s.Slice != nil {
		o.Slice = s.Slice
	}
	if s.Unexposed != nil {
		o.Unexposed = s.Unexposed
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseList) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseList) ToPlain() PlainIdentifiable {

	out := NewList()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreationOnly != nil {
		out.CreationOnly = *o.CreationOnly
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.ReadOnly != nil {
		out.ReadOnly = *o.ReadOnly
	}
	if o.Secret != nil {
		out.Secret = *o.Secret
	}
	if o.Slice != nil {
		out.Slice = *o.Slice
	}
	if o.Unexposed != nil {
		out.Unexposed = *o.Unexposed
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparseList) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseList) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseList.
func (o *SparseList) DeepCopy() *SparseList {

	if o == nil {
		return nil
	}

	out := &SparseList{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseList.
func (o *SparseList) DeepCopyInto(out *SparseList) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseList: %s", err))
	}

	*out = *target.(*SparseList)
}

type mongoAttributesList struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	CreationOnly string        `bson:"creationonly"`
	Date         time.Time     `bson:"date"`
	Description  string        `bson:"description"`
	Name         string        `bson:"name"`
	ParentID     string        `bson:"parentid"`
	ParentType   string        `bson:"parenttype"`
	ReadOnly     string        `bson:"readonly"`
	Secret       string        `bson:"secret"`
	Slice        []string      `bson:"slice"`
	Unexposed    string        `bson:"unexposed"`
}
type mongoAttributesSparseList struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	CreationOnly *string       `bson:"creationonly,omitempty"`
	Date         *time.Time    `bson:"date,omitempty"`
	Description  *string       `bson:"description,omitempty"`
	Name         *string       `bson:"name,omitempty"`
	ParentID     *string       `bson:"parentid,omitempty"`
	ParentType   *string       `bson:"parenttype,omitempty"`
	ReadOnly     *string       `bson:"readonly,omitempty"`
	Secret       *string       `bson:"secret,omitempty"`
	Slice        *[]string     `bson:"slice,omitempty"`
	Unexposed    *string       `bson:"unexposed,omitempty"`
}
// TaskStatusValue represents the possible values for attribute "status".
type TaskStatusValue string

const (
	// TaskStatusDONE represents the value DONE.
	TaskStatusDONE TaskStatusValue = "DONE"

	// TaskStatusPROGRESS represents the value PROGRESS.
	TaskStatusPROGRESS TaskStatusValue = "PROGRESS"

	// TaskStatusTODO represents the value TODO.
	TaskStatusTODO TaskStatusValue = "TODO"
)

// TaskIdentity represents the Identity of the object.
var TaskIdentity = Identity{
	Name:     "task",
	Category: "tasks",
	Package:  "todo-list",
	Private:  false,
}

// TasksList represents a list of Tasks
type TasksList []*Task

// Identity returns the identity of the objects in the list.
func (o TasksList) Identity() Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the TasksList.
func (o TasksList) Copy() Identifiables {

	copy := append(TasksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TasksList.
func (o TasksList) Append(objects ...Identifiable) Identifiables {

	out := append(TasksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Task))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o TasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TasksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TasksList converted to SparseTasksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TasksList) ToSparse(fields ...string) Identifiables {

	out := make(SparseTasksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTask)
	}

	return out
}

// Version returns the version of the content.
func (o TasksList) Version() int {

	return 1
}

// Task represents the model of a task
type Task struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// The status of the task.
	Status TaskStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTask returns a new *Task
func NewTask() *Task {

	return &Task{
		ModelVersion: 1,
		Status:       TaskStatusTODO,
	}
}

// Identity returns the Identity of the object.
func (o *Task) Identity() Identity {

	return TaskIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Task) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Task) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Task) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTask{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Description = o.Description
	s.Name = o.Name
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.Status = o.Status

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Task) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Description = s.Description
	o.Name = s.Name
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.Status = s.Status

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Task) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Task) BleveType() string {

	return "task"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Task) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Task) Doc() string {

	return `Represent a task to do in a listd.`
}

func (o *Task) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *Task) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Task) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Task) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTask{
			ID:          &o.ID,
			Description: &o.Description,
			Name:        &o.Name,
			ParentID:    &o.ParentID,
			ParentType:  &o.ParentType,
			Status:      &o.Status,
		}
	}

	sp := &SparseTask{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "status":
			sp.Status = &(o.Status)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTask to the object.
func (o *Task) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTask)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
}

// DeepCopy returns a deep copy if the Task.
func (o *Task) DeepCopy() *Task {

	if o == nil {
		return nil
	}

	out := &Task{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Task.
func (o *Task) DeepCopyInto(out *Task) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Task: %s", err))
	}

	*out = *target.(*Task)
}

// Validate valides the current information stored into the structure.
func (o *Task) Validate() error {

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateStringInList("status", string(o.Status), []string{"DONE", "PROGRESS", "TODO"}, false); err != nil {
		errors = errors.Append(err)
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
func (*Task) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := TaskAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TaskLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Task) AttributeSpecifications() map[string]AttributeSpecification {

	return TaskAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Task) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "status":
		return o.Status
	}

	return nil
}

// TaskAttributesMap represents the map of attribute for Task.
var TaskAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"DONE", "PROGRESS", "TODO"},
		ConvertedName:  "Status",
		DefaultValue:   TaskStatusTODO,
		Description:    `The status of the task.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}

// TaskLowerCaseAttributesMap represents the map of attribute for Task.
var TaskLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"DONE", "PROGRESS", "TODO"},
		ConvertedName:  "Status",
		DefaultValue:   TaskStatusTODO,
		Description:    `The status of the task.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}

// SparseTasksList represents a list of SparseTasks
type SparseTasksList []*SparseTask

// Identity returns the identity of the objects in the list.
func (o SparseTasksList) Identity() Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the SparseTasksList.
func (o SparseTasksList) Copy() Identifiables {

	copy := append(SparseTasksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTasksList.
func (o SparseTasksList) Append(objects ...Identifiable) Identifiables {

	out := append(SparseTasksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTask))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseTasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTasksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTasksList converted to TasksList.
func (o SparseTasksList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTasksList) Version() int {

	return 1
}

// SparseTask represents the sparse version of a task.
type SparseTask struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The name.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// The status of the task.
	Status *TaskStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTask returns a new  SparseTask.
func NewSparseTask() *SparseTask {
	return &SparseTask{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTask) Identity() Identity {

	return TaskIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTask) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTask) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTask) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTask{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.Status != nil {
		s.Status = o.Status
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTask) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.Status != nil {
		o.Status = s.Status
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTask) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTask) ToPlain() PlainIdentifiable {

	out := NewTask()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.Status != nil {
		out.Status = *o.Status
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparseTask) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseTask) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseTask.
func (o *SparseTask) DeepCopy() *SparseTask {

	if o == nil {
		return nil
	}

	out := &SparseTask{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTask.
func (o *SparseTask) DeepCopyInto(out *SparseTask) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTask: %s", err))
	}

	*out = *target.(*SparseTask)
}

type mongoAttributesTask struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Description string          `bson:"description"`
	Name        string          `bson:"name"`
	ParentID    string          `bson:"parentid"`
	ParentType  string          `bson:"parenttype"`
	Status      TaskStatusValue `bson:"status"`
}
type mongoAttributesSparseTask struct {
	ID          bson.ObjectId    `bson:"_id,omitempty"`
	Description *string          `bson:"description,omitempty"`
	Name        *string          `bson:"name,omitempty"`
	ParentID    *string          `bson:"parentid,omitempty"`
	ParentType  *string          `bson:"parenttype,omitempty"`
	Status      *TaskStatusValue `bson:"status,omitempty"`
}

// UnmarshalableListIdentity represents the Identity of the object.
var UnmarshalableListIdentity = Identity{Name: "list", Category: "lists"}

// UnmarshalableListsList represents a list of UnmarshalableLists
type UnmarshalableListsList []*UnmarshalableList

// Identity returns the identity of the objects in the list.
func (o UnmarshalableListsList) Identity() Identity {

	return UnmarshalableListIdentity
}

// Copy returns a pointer to a copy the UnmarshalableListsList.
func (o UnmarshalableListsList) Copy() Identifiables {

	copy := append(UnmarshalableListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the UnmarshalableListsList.
func (o UnmarshalableListsList) Append(objects ...Identifiable) Identifiables {

	out := append(UnmarshalableListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*UnmarshalableList))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o UnmarshalableListsList) List() IdentifiablesList {

	out := IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UnmarshalableListsList) DefaultOrder() []string {

	return []string{
		"flagDefaultOrderingKey",
	}
}

// Version returns the version of the content.
func (o UnmarshalableListsList) Version() int {

	return 1
}

// An UnmarshalableList is a List that cannot be marshalled  or unmarshalled.
type UnmarshalableList struct {
	List
}

// NewUnmarshalableList returns a new UnmarshalableList.
func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

// Identity returns the identity.
func (o *UnmarshalableList) Identity() Identity { return UnmarshalableListIdentity }

// UnmarshalJSON makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// Validate validates the data
func (o *UnmarshalableList) Validate() Errors { return nil }

// An UnmarshalableError is a List that cannot be marshalled or unmarshalled.
type UnmarshalableError struct {
	Error
}

// UnmarshalJSON makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}
// UserIdentity represents the Identity of the object.
var UserIdentity = Identity{
	Name:     "user",
	Category: "users",
	Package:  "todo-list",
	Private:  false,
}

// UsersList represents a list of Users
type UsersList []*User

// Identity returns the identity of the objects in the list.
func (o UsersList) Identity() Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the UsersList.
func (o UsersList) Copy() Identifiables {

	copy := append(UsersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the UsersList.
func (o UsersList) Append(objects ...Identifiable) Identifiables {

	out := append(UsersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*User))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o UsersList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UsersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the UsersList converted to SparseUsersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o UsersList) ToSparse(fields ...string) Identifiables {

	out := make(SparseUsersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseUser)
	}

	return out
}

// Version returns the version of the content.
func (o UsersList) Version() int {

	return 1
}

// User represents the model of a user
type User struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The first name.
	FirstName string `json:"firstName" msgpack:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName string `json:"lastName" msgpack:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// the login.
	UserName string `json:"userName" msgpack:"userName" bson:"username" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *User) Identity() Identity {

	return UserIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *User) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *User) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUser{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.FirstName = o.FirstName
	s.LastName = o.LastName
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.UserName = o.UserName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *User) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesUser{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.FirstName = s.FirstName
	o.LastName = s.LastName
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.UserName = s.UserName

	return nil
}

// Version returns the hardcoded version of the model.
func (o *User) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *User) BleveType() string {

	return "user"
}

// DefaultOrder returns the list of default ordering fields.
func (o *User) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *User) Doc() string {

	return `Represent a user.`
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *User) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseUser{
			ID:         &o.ID,
			FirstName:  &o.FirstName,
			LastName:   &o.LastName,
			ParentID:   &o.ParentID,
			ParentType: &o.ParentType,
			UserName:   &o.UserName,
		}
	}

	sp := &SparseUser{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "firstName":
			sp.FirstName = &(o.FirstName)
		case "lastName":
			sp.LastName = &(o.LastName)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "userName":
			sp.UserName = &(o.UserName)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseUser to the object.
func (o *User) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseUser)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.FirstName != nil {
		o.FirstName = *so.FirstName
	}
	if so.LastName != nil {
		o.LastName = *so.LastName
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.UserName != nil {
		o.UserName = *so.UserName
	}
}

// DeepCopy returns a deep copy if the User.
func (o *User) DeepCopy() *User {

	if o == nil {
		return nil
	}

	out := &User{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *User.
func (o *User) DeepCopyInto(out *User) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy User: %s", err))
	}

	*out = *target.(*User)
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() error {

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("firstName", o.FirstName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateRequiredString("lastName", o.LastName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateRequiredString("userName", o.UserName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*User) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := UserAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UserLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*User) AttributeSpecifications() map[string]AttributeSpecification {

	return UserAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *User) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "firstName":
		return o.FirstName
	case "lastName":
		return o.LastName
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "userName":
		return o.UserName
	}

	return nil
}

// UserAttributesMap represents the map of attribute for User.
var UserAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"FirstName": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `The first name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"LastName": {
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `The last name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"UserName": {
		AllowedChoices: []string{},
		ConvertedName:  "UserName",
		Description:    `the login.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "userName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// UserLowerCaseAttributesMap represents the map of attribute for User.
var UserLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"firstname": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `The first name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"lastname": {
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `The last name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"username": {
		AllowedChoices: []string{},
		ConvertedName:  "UserName",
		Description:    `the login.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "userName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseUsersList represents a list of SparseUsers
type SparseUsersList []*SparseUser

// Identity returns the identity of the objects in the list.
func (o SparseUsersList) Identity() Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the SparseUsersList.
func (o SparseUsersList) Copy() Identifiables {

	copy := append(SparseUsersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseUsersList.
func (o SparseUsersList) Append(objects ...Identifiable) Identifiables {

	out := append(SparseUsersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseUser))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseUsersList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseUsersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseUsersList converted to UsersList.
func (o SparseUsersList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseUsersList) Version() int {

	return 1
}

// SparseUser represents the sparse version of a user.
type SparseUser struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The first name.
	FirstName *string `json:"firstName,omitempty" msgpack:"firstName,omitempty" bson:"firstname,omitempty" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName *string `json:"lastName,omitempty" msgpack:"lastName,omitempty" bson:"lastname,omitempty" mapstructure:"lastName,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// the login.
	UserName *string `json:"userName,omitempty" msgpack:"userName,omitempty" bson:"username,omitempty" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseUser returns a new  SparseUser.
func NewSparseUser() *SparseUser {
	return &SparseUser{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseUser) Identity() Identity {

	return UserIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseUser) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseUser) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUser) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseUser{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.FirstName != nil {
		s.FirstName = o.FirstName
	}
	if o.LastName != nil {
		s.LastName = o.LastName
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.UserName != nil {
		s.UserName = o.UserName
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUser) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseUser{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.FirstName != nil {
		o.FirstName = s.FirstName
	}
	if s.LastName != nil {
		o.LastName = s.LastName
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.UserName != nil {
		o.UserName = s.UserName
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseUser) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseUser) ToPlain() PlainIdentifiable {

	out := NewUser()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.FirstName != nil {
		out.FirstName = *o.FirstName
	}
	if o.LastName != nil {
		out.LastName = *o.LastName
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.UserName != nil {
		out.UserName = *o.UserName
	}

	return out
}

// DeepCopy returns a deep copy if the SparseUser.
func (o *SparseUser) DeepCopy() *SparseUser {

	if o == nil {
		return nil
	}

	out := &SparseUser{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseUser.
func (o *SparseUser) DeepCopyInto(out *SparseUser) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseUser: %s", err))
	}

	*out = *target.(*SparseUser)
}

type mongoAttributesUser struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	FirstName  string        `bson:"firstname"`
	LastName   string        `bson:"lastname"`
	ParentID   string        `bson:"parentid"`
	ParentType string        `bson:"parenttype"`
	UserName   string        `bson:"username"`
}
type mongoAttributesSparseUser struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	FirstName  *string       `bson:"firstname,omitempty"`
	LastName   *string       `bson:"lastname,omitempty"`
	ParentID   *string       `bson:"parentid,omitempty"`
	ParentType *string       `bson:"parenttype,omitempty"`
	UserName   *string       `bson:"username,omitempty"`
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
func (o *Root) Identity() Identity {

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

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRoot{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Root) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

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

	errors := Errors{}
	requiredErrors := Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Root) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := RootAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RootLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Root) AttributeSpecifications() map[string]AttributeSpecification {

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
var RootAttributesMap = map[string]AttributeSpecification{}

// RootLowerCaseAttributesMap represents the map of attribute for Root.
var RootLowerCaseAttributesMap = map[string]AttributeSpecification{}

type mongoAttributesRoot struct {
}
var (
	identityNamesMap = map[string]Identity{
		"list": ListIdentity,
		"root": RootIdentity,
		"task": TaskIdentity,
		"user": UserIdentity,
	}

	identitycategoriesMap = map[string]Identity{
		"lists": ListIdentity,
		"root":  RootIdentity,
		"tasks": TaskIdentity,
		"users": UserIdentity,
	}

	aliasesMap = map[string]Identity{
		"lst": ListIdentity,
		"tsk": TaskIdentity,
		"usr": UserIdentity,
	}

	indexesMap = map[string][][]string{
		"list": nil,
		"root": nil,
		"task": nil,
		"user": nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity Identity) Identifiable {

	switch identity {

	case ListIdentity:
		return NewList()
	case RootIdentity:
		return NewRoot()
	case TaskIdentity:
		return NewTask()
	case UserIdentity:
		return NewUser()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity Identity) SparseIdentifiable {

	switch identity {

	case ListIdentity:
		return NewSparseList()
	case TaskIdentity:
		return NewSparseTask()
	case UserIdentity:
		return NewSparseUser()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity Identity) Identifiables {

	switch identity {

	case ListIdentity:
		return &ListsList{}
	case TaskIdentity:
		return &TasksList{}
	case UserIdentity:
		return &UsersList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity Identity) SparseIdentifiables {

	switch identity {

	case ListIdentity:
		return &SparseListsList{}
	case TaskIdentity:
		return &SparseTasksList{}
	case UserIdentity:
		return &SparseUsersList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() RelationshipsRegistry {

	return relationshipsRegistry
}

var manager = modelManager{}

// Manager returns the model ModelManager.
func Manager() ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []Identity {

	return []Identity{
		ListIdentity,
		RootIdentity,
		TaskIdentity,
		UserIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity Identity) []string {

	switch identity {
	case ListIdentity:
		return []string{
			"lst",
		}
	case RootIdentity:
		return []string{}
	case TaskIdentity:
		return []string{
			"tsk",
		}
	case UserIdentity:
		return []string{
			"usr",
		}
	}

	return nil
}

var relationshipsRegistry RelationshipsRegistry

func init() {

	relationshipsRegistry = RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlcp1",
						Type: "string",
					},
					{
						Name: "rlcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "ldp1",
						Type: "string",
					},
					{
						Name: "ldp2",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lgp1",
						Type: "string",
					},
					{
						Name: "lgp2",
						Type: "boolean",
					},
					{
						Name: "sAp1",
						Type: "string",
					},
					{
						Name: "sAp2",
						Type: "boolean",
					},
					{
						Name: "sBp1",
						Type: "string",
					},
					{
						Name: "sBp2",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RootIdentity] = &Relationship{}

	relationshipsRegistry[TaskIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltcp1",
						Type: "string",
					},
					{
						Name: "ltcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[UserIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rucp1",
						Type: "string",
					},
					{
						Name: "rucp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {
				RequiredParameters: NewParametersRequirement(
					[][][]string{
						{
							{
								"confirm",
							},
						},
					},
				),
				Parameters: []ParameterDefinition{
					{
						Name: "confirm",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

}
