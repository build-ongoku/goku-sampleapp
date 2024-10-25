package svccore_enttask_typ

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/teejays/gokutil/dalutil"
	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/panics"
	"github.com/teejays/gokutil/scalars"

	app_typ "sampleapp/backend/.goku/generated/typ"
)

// Task: <comments>
type Task struct {
	ID                scalars.ID          `json:"id"`
	Name              string              `json:"name" validate:"required"`
	Description       string              `json:"description" validate:"required"`
	Method            app_typ.MethodName  `json:"method" validate:"required"`
	MethodRequestData scalars.GenericData `json:"methodRequestData" validate:"required"`
	Enabled           bool                `json:"enabled"`
	CreatedAt         scalars.Timestamp   `json:"createdAt"`
	UpdatedAt         scalars.Timestamp   `json:"updatedAt"`
	DeletedAt         *scalars.Timestamp  `json:"deletedAt"`
}

func (t Task) GetID() scalars.ID {
	return t.ID
}
func (t Task) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t Task) SetUpdatedAt(tim scalars.Timestamp) Task {
	t.UpdatedAt = tim
	return t
}

// TaskInput: <comments>
type TaskInput struct {
	Name              string              `json:"name" validate:"required"`
	Description       string              `json:"description" validate:"required"`
	Method            app_typ.MethodName  `json:"method" validate:"required"`
	MethodRequestData scalars.GenericData `json:"methodRequestData" validate:"required"`
	Enabled           bool                `json:"enabled"`
}

func NewTaskWithMetaArrayFromInputs(ctx context.Context, ins []TaskInput) []Task {
	var outs []Task
	for _, in := range ins {
		outs = append(outs, NewTaskWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewTaskWithMetaFromInput(ctx context.Context, in TaskInput) Task {
	return Task{
		Name:              in.Name,
		Description:       in.Description,
		Method:            in.Method,
		MethodRequestData: in.MethodRequestData,
		Enabled:           in.Enabled,
	}
}

// MethodName enum: <insert comment>
type MethodName int

const (
	MethodName_INVALID     MethodName = 0
	MethodName_Add         MethodName = 1
	MethodName_Update      MethodName = 2
	MethodName_Get         MethodName = 3
	MethodName_List        MethodName = 4
	MethodName_QueryByText MethodName = 5
	MethodName_HookSavePre MethodName = 6
	MethodName_ActionRun   MethodName = 7
)

func NewMethodNameFromString(s string) MethodName {
	switch s {
	case "INVALID":
		return MethodName_INVALID
	case "Add":
		return MethodName_Add
	case "Update":
		return MethodName_Update
	case "Get":
		return MethodName_Get
	case "List":
		return MethodName_List
	case "QueryByText":
		return MethodName_QueryByText
	case "HookSavePre":
		return MethodName_HookSavePre
	case "ActionRun":
		return MethodName_ActionRun

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "MethodName"))
	}
}

// String implements the `fmt.Stringer` interface for MethodName. It allows us to print the enum values as strings.
func (f MethodName) String() string {
	switch f {
	case MethodName_INVALID:
		return "INVALID"
	case MethodName_Add:
		return "Add"
	case MethodName_Update:
		return "Update"
	case MethodName_Get:
		return "Get"
	case MethodName_List:
		return "List"
	case MethodName_QueryByText:
		return "QueryByText"
	case MethodName_HookSavePre:
		return "HookSavePre"
	case MethodName_ActionRun:
		return "ActionRun"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "MethodName"))
	}
}

// Name gives a naam representation of the enum value
func (f MethodName) Name() naam.Name {
	switch f {
	case MethodName_Add:
		return naam.New("add")
	case MethodName_Update:
		return naam.New("update")
	case MethodName_Get:
		return naam.New("get")
	case MethodName_List:
		return naam.New("list")
	case MethodName_QueryByText:
		return naam.New("query_by_text")
	case MethodName_HookSavePre:
		return naam.New("hook_save_pre")
	case MethodName_ActionRun:
		return naam.New("action_run")
	default:
		panics.P("MethodName.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f MethodName) Value() (driver.Value, error) {
	switch f {
	case MethodName_INVALID:
		return nil, nil
	case MethodName_Add:
		return "Add", nil
	case MethodName_Update:
		return "Update", nil
	case MethodName_Get:
		return "Get", nil
	case MethodName_List:
		return "List", nil
	case MethodName_QueryByText:
		return "QueryByText", nil
	case MethodName_HookSavePre:
		return "HookSavePre", nil
	case MethodName_ActionRun:
		return "ActionRun", nil

	default:
		return nil, fmt.Errorf("Cannot save enum MethodName to DB: '%d' is not a valid value for enum MethodName", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *MethodName) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewMethodNameFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "MethodName")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f MethodName) ImplementsGraphQLType(name string) bool {
	return name == "Core_Task_MethodName"
}

func (f *MethodName) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewMethodNameFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for MethodName: %T", input)
	}
	return err
}

func (f *MethodName) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum MethodName to a string: %w", err)
	}
	i := NewMethodNameFromString(enumStr)
	*f = i
	return nil
}

func (f MethodName) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil MethodName pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type MethodNameCondition struct {
	Op     filterlib.Operator
	Values []MethodName
}

func (c MethodNameCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c MethodNameCondition) Len() int {
	return len(c.Values)
}
func (c MethodNameCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskAction enum: <insert comment>
type TaskAction int

const (
	TaskAction_INVALID TaskAction = 0
	TaskAction_Run     TaskAction = 1
)

func NewTaskActionFromString(s string) TaskAction {
	switch s {
	case "INVALID":
		return TaskAction_INVALID
	case "Run":
		return TaskAction_Run

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TaskAction"))
	}
}

// String implements the `fmt.Stringer` interface for TaskAction. It allows us to print the enum values as strings.
func (f TaskAction) String() string {
	switch f {
	case TaskAction_INVALID:
		return "INVALID"
	case TaskAction_Run:
		return "Run"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TaskAction"))
	}
}

// Name gives a naam representation of the enum value
func (f TaskAction) Name() naam.Name {
	switch f {
	case TaskAction_Run:
		return naam.New("run")
	default:
		panics.P("TaskAction.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TaskAction) Value() (driver.Value, error) {
	switch f {
	case TaskAction_INVALID:
		return nil, nil
	case TaskAction_Run:
		return "Run", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TaskAction to DB: '%d' is not a valid value for enum TaskAction", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TaskAction) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTaskActionFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TaskAction")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TaskAction) ImplementsGraphQLType(name string) bool {
	return name == "Core_Task_TaskAction"
}

func (f *TaskAction) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTaskActionFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TaskAction: %T", input)
	}
	return err
}

func (f *TaskAction) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TaskAction to a string: %w", err)
	}
	i := NewTaskActionFromString(enumStr)
	*f = i
	return nil
}

func (f TaskAction) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TaskAction pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TaskActionCondition struct {
	Op     filterlib.Operator
	Values []TaskAction
}

func (c TaskActionCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TaskActionCondition) Len() int {
	return len(c.Values)
}
func (c TaskActionCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskField enum: <insert comment>
type TaskField int

const (
	TaskField_INVALID           TaskField = 0
	TaskField_ID                TaskField = 1
	TaskField_Name              TaskField = 2
	TaskField_Description       TaskField = 3
	TaskField_Method            TaskField = 4
	TaskField_MethodRequestData TaskField = 5
	TaskField_Enabled           TaskField = 6
	TaskField_CreatedAt         TaskField = 7
	TaskField_UpdatedAt         TaskField = 8
	TaskField_DeletedAt         TaskField = 9
)

func NewTaskFieldFromString(s string) TaskField {
	switch s {
	case "INVALID":
		return TaskField_INVALID
	case "ID":
		return TaskField_ID
	case "Name":
		return TaskField_Name
	case "Description":
		return TaskField_Description
	case "Method":
		return TaskField_Method
	case "MethodRequestData":
		return TaskField_MethodRequestData
	case "Enabled":
		return TaskField_Enabled
	case "CreatedAt":
		return TaskField_CreatedAt
	case "UpdatedAt":
		return TaskField_UpdatedAt
	case "DeletedAt":
		return TaskField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TaskField"))
	}
}

// String implements the `fmt.Stringer` interface for TaskField. It allows us to print the enum values as strings.
func (f TaskField) String() string {
	switch f {
	case TaskField_INVALID:
		return "INVALID"
	case TaskField_ID:
		return "ID"
	case TaskField_Name:
		return "Name"
	case TaskField_Description:
		return "Description"
	case TaskField_Method:
		return "Method"
	case TaskField_MethodRequestData:
		return "MethodRequestData"
	case TaskField_Enabled:
		return "Enabled"
	case TaskField_CreatedAt:
		return "CreatedAt"
	case TaskField_UpdatedAt:
		return "UpdatedAt"
	case TaskField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TaskField"))
	}
}

// Name gives a naam representation of the enum value
func (f TaskField) Name() naam.Name {
	switch f {
	case TaskField_ID:
		return naam.New("id")
	case TaskField_Name:
		return naam.New("name")
	case TaskField_Description:
		return naam.New("description")
	case TaskField_Method:
		return naam.New("method")
	case TaskField_MethodRequestData:
		return naam.New("method_request_data")
	case TaskField_Enabled:
		return naam.New("enabled")
	case TaskField_CreatedAt:
		return naam.New("created_at")
	case TaskField_UpdatedAt:
		return naam.New("updated_at")
	case TaskField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("TaskField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TaskField) Value() (driver.Value, error) {
	switch f {
	case TaskField_INVALID:
		return nil, nil
	case TaskField_ID:
		return "ID", nil
	case TaskField_Name:
		return "Name", nil
	case TaskField_Description:
		return "Description", nil
	case TaskField_Method:
		return "Method", nil
	case TaskField_MethodRequestData:
		return "MethodRequestData", nil
	case TaskField_Enabled:
		return "Enabled", nil
	case TaskField_CreatedAt:
		return "CreatedAt", nil
	case TaskField_UpdatedAt:
		return "UpdatedAt", nil
	case TaskField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TaskField to DB: '%d' is not a valid value for enum TaskField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TaskField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTaskFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TaskField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TaskField) ImplementsGraphQLType(name string) bool {
	return name == "Core_Task_TaskField"
}

func (f *TaskField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTaskFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TaskField: %T", input)
	}
	return err
}

func (f *TaskField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TaskField to a string: %w", err)
	}
	i := NewTaskFieldFromString(enumStr)
	*f = i
	return nil
}

func (f TaskField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TaskField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TaskFieldCondition struct {
	Op     filterlib.Operator
	Values []TaskField
}

func (c TaskFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TaskFieldCondition) Len() int {
	return len(c.Values)
}
func (c TaskFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskFilter: <comments>
type TaskFilter struct {
	ID                *filterlib.IDCondition          `json:"id"`
	Name              *filterlib.StringCondition      `json:"name"`
	Description       *filterlib.StringCondition      `json:"description"`
	Method            *app_typ.MethodNameCondition    `json:"method"`
	MethodRequestData *filterlib.GenericDataCondition `json:"methodRequestData"`
	Enabled           *filterlib.BoolCondition        `json:"enabled"`
	CreatedAt         *filterlib.TimestampCondition   `json:"createdAt"`
	UpdatedAt         *filterlib.TimestampCondition   `json:"updatedAt"`
	DeletedAt         *filterlib.TimestampCondition   `json:"deletedAt"`
	And               []TaskFilter                    `json:"and"`
	Or                []TaskFilter                    `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object Task `json:"object"`
}

// AddRequest:
type AddRequest = dalutil.EntityAddRequest[TaskInput]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[Task]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[TaskFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[Task]

// MethodNameCondition:
// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[Task]

// StandardEntityRequest:
// StandardEntityResponse:
// Task: Main type for entity Task
// TaskActionCondition:
// TaskFieldCondition:
// TaskFilter:
// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[Task, TaskField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[Task]
