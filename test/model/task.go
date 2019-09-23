package testmodel

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

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
var TaskIdentity = elemental.Identity{
	Name:     "task",
	Category: "tasks",
	Package:  "todo-list",
	Private:  false,
}

// TasksList represents a list of Tasks
type TasksList []*Task

// Identity returns the identity of the objects in the list.
func (o TasksList) Identity() elemental.Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the TasksList.
func (o TasksList) Copy() elemental.Identifiables {

	copy := append(TasksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TasksList.
func (o TasksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TasksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Task))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TasksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
func (o TasksList) ToSparse(fields ...string) elemental.Identifiables {

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
func (o *Task) Identity() elemental.Identity {

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

	s := &mongoAttributesTask{}

	s.ID = bson.ObjectIdHex(o.ID)
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
func (o *Task) ToSparse(fields ...string) elemental.SparseIdentifiable {

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
func (o *Task) Patch(sparse elemental.SparseIdentifiable) {
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

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"DONE", "PROGRESS", "TODO"}, false); err != nil {
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
func (*Task) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TaskAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TaskLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Task) AttributeSpecifications() map[string]elemental.AttributeSpecification {

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
var TaskAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Description": elemental.AttributeSpecification{
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
	"Name": elemental.AttributeSpecification{
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
	"Status": elemental.AttributeSpecification{
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
var TaskLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"description": elemental.AttributeSpecification{
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
	"name": elemental.AttributeSpecification{
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
	"status": elemental.AttributeSpecification{
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
func (o SparseTasksList) Identity() elemental.Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the SparseTasksList.
func (o SparseTasksList) Copy() elemental.Identifiables {

	copy := append(SparseTasksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTasksList.
func (o SparseTasksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTasksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTask))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTasksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
func (o SparseTasksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
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
func (o *SparseTask) Identity() elemental.Identity {

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

	o.ID = &id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTask) GetBSON() (interface{}, error) {

	s := &mongoAttributesSparseTask{}

	s.ID = bson.ObjectIdHex(*o.ID)
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
func (o *SparseTask) ToPlain() elemental.PlainIdentifiable {

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
func (o *SparseTask) GetName() string {

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
	ID          bson.ObjectId   `bson:"_id"`
	Description string          `bson:"description"`
	Name        string          `bson:"name"`
	ParentID    string          `bson:"parentid"`
	ParentType  string          `bson:"parenttype"`
	Status      TaskStatusValue `bson:"status"`
}
type mongoAttributesSparseTask struct {
	ID          bson.ObjectId    `bson:"_id"`
	Description *string          `bson:"description,omitempty"`
	Name        *string          `bson:"name,omitempty"`
	ParentID    *string          `bson:"parentid,omitempty"`
	ParentType  *string          `bson:"parenttype,omitempty"`
	Status      *TaskStatusValue `bson:"status,omitempty"`
}
