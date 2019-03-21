package testmodel

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UserIdentity represents the Identity of the object.
var UserIdentity = elemental.Identity{
	Name:     "user",
	Category: "users",
	Package:  "todo-list",
	Private:  false,
}

// UsersList represents a list of Users
type UsersList []*User

// Identity returns the identity of the objects in the list.
func (o UsersList) Identity() elemental.Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the UsersList.
func (o UsersList) Copy() elemental.Identifiables {

	copy := append(UsersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the UsersList.
func (o UsersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(UsersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*User))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o UsersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
func (o UsersList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
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
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// The first name.
	FirstName string `json:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName string `json:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// the login.
	UserName string `json:"userName" bson:"username" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *User) Identity() elemental.Identity {

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

// Version returns the hardcoded version of the model.
func (o *User) Version() int {

	return 1
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
func (o *User) ToSparse(fields ...string) elemental.SparseIdentifiable {

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
func (o *User) Patch(sparse elemental.SparseIdentifiable) {
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

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("firstName", o.FirstName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("lastName", o.LastName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("userName", o.UserName); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*User) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := UserAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UserLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*User) AttributeSpecifications() map[string]elemental.AttributeSpecification {

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
var UserAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
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
	"FirstName": elemental.AttributeSpecification{
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
	"LastName": elemental.AttributeSpecification{
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
	"ParentID": elemental.AttributeSpecification{
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
	"ParentType": elemental.AttributeSpecification{
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
	"UserName": elemental.AttributeSpecification{
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
var UserLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
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
	"firstname": elemental.AttributeSpecification{
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
	"lastname": elemental.AttributeSpecification{
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
	"parentid": elemental.AttributeSpecification{
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
	"parenttype": elemental.AttributeSpecification{
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
	"username": elemental.AttributeSpecification{
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
func (o SparseUsersList) Identity() elemental.Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the SparseUsersList.
func (o SparseUsersList) Copy() elemental.Identifiables {

	copy := append(SparseUsersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseUsersList.
func (o SparseUsersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseUsersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseUser))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseUsersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
func (o SparseUsersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// The first name.
	FirstName *string `json:"firstName,omitempty" bson:"firstname" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName *string `json:"lastName,omitempty" bson:"lastname" mapstructure:"lastName,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// the login.
	UserName *string `json:"userName,omitempty" bson:"username" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseUser returns a new  SparseUser.
func NewSparseUser() *SparseUser {
	return &SparseUser{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseUser) Identity() elemental.Identity {

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

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseUser) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseUser) ToPlain() elemental.PlainIdentifiable {

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
