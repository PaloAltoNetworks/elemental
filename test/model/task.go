package testmodel

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
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
	Private:  false,
}

// TasksList represents a list of Tasks
type TasksList []*Task

// ContentIdentity returns the identity of the objects in the list.
func (o TasksList) ContentIdentity() elemental.Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the TasksList.
func (o TasksList) Copy() elemental.ContentIdentifiable {

	copy := append(TasksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TasksList.
func (o TasksList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(TasksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Task))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TasksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TasksList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o TasksList) Version() int {

	return 1
}

// Task represents the model of a task
type Task struct {
	// The identifier.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// The description.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// The status of the task.
	Status TaskStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTask returns a new *Task
func NewTask() *Task {

	return &Task{
		ModelVersion: 1,
		Status:       "TODO",
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

// Version returns the hardcoded version of the model.
func (o *Task) Version() int {

	return 1
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

// SetName sets the given Name of the receiver.
func (o *Task) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *Task) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"DONE", "PROGRESS", "TODO"}, false); err != nil {
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

// TaskAttributesMap represents the map of attribute for Task.
var TaskAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
