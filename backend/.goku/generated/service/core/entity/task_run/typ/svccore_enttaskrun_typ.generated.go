package svccore_enttaskrun_typ

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
)

// TaskRun: <comments>
type TaskRun struct {
	ID                 scalars.ID          `json:"id"`
	StartedAt          scalars.Timestamp   `json:"startedAt" validate:"required"`
	CompletedAt        scalars.Timestamp   `json:"completedAt" validate:"required"`
	Status             Status              `json:"status" validate:"required"`
	MethodRequestData  scalars.GenericData `json:"methodRequestData" validate:"required"`
	MethodResponseData scalars.GenericData `json:"methodResponseData" validate:"required"`
	TriggeredBy        TriggeredBy         `json:"triggeredBy" validate:"required"`
	CreatedAt          scalars.Timestamp   `json:"createdAt"`
	UpdatedAt          scalars.Timestamp   `json:"updatedAt"`
	DeletedAt          *scalars.Timestamp  `json:"deletedAt"`
	TaskID             scalars.ID          `json:"taskID"`
}

func (t TaskRun) GetID() scalars.ID {
	return t.ID
}
func (t TaskRun) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t TaskRun) SetUpdatedAt(tim scalars.Timestamp) TaskRun {
	t.UpdatedAt = tim
	return t
}

// TaskRunInput: <comments>
type TaskRunInput struct {
	MethodRequestData scalars.GenericData `json:"methodRequestData" validate:"required"`
	TriggeredBy       TriggeredBy         `json:"triggeredBy" validate:"required"`
	TaskID            scalars.ID          `json:"taskID"`
}

func NewTaskRunWithMetaArrayFromInputs(ctx context.Context, ins []TaskRunInput) []TaskRun {
	var outs []TaskRun
	for _, in := range ins {
		outs = append(outs, NewTaskRunWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewTaskRunWithMetaFromInput(ctx context.Context, in TaskRunInput) TaskRun {
	return TaskRun{
		MethodRequestData: in.MethodRequestData,
		TriggeredBy:       in.TriggeredBy,
		TaskID:            in.TaskID,
	}
}

// MethodName enum: <insert comment>
type MethodName int

const (
	MethodName_INVALID       MethodName = 0
	MethodName_Add           MethodName = 1
	MethodName_Update        MethodName = 2
	MethodName_Get           MethodName = 3
	MethodName_List          MethodName = 4
	MethodName_QueryByText   MethodName = 5
	MethodName_HookCreatePre MethodName = 6
	MethodName_ActionRun     MethodName = 7
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
	case "HookCreatePre":
		return MethodName_HookCreatePre
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
	case MethodName_HookCreatePre:
		return "HookCreatePre"
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
	case MethodName_HookCreatePre:
		return naam.New("hook_create_pre")
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
	case MethodName_HookCreatePre:
		return "HookCreatePre", nil
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
	return name == "Core_TaskRun_MethodName"
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

// Status enum: <insert comment>
type Status int

const (
	Status_INVALID Status = 0
	Status_Pending Status = 1
	Status_Running Status = 2
	Status_Success Status = 3
	Status_Failed  Status = 4
)

func NewStatusFromString(s string) Status {
	switch s {
	case "INVALID":
		return Status_INVALID
	case "Pending":
		return Status_Pending
	case "Running":
		return Status_Running
	case "Success":
		return Status_Success
	case "Failed":
		return Status_Failed

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "Status"))
	}
}

// String implements the `fmt.Stringer` interface for Status. It allows us to print the enum values as strings.
func (f Status) String() string {
	switch f {
	case Status_INVALID:
		return "INVALID"
	case Status_Pending:
		return "Pending"
	case Status_Running:
		return "Running"
	case Status_Success:
		return "Success"
	case Status_Failed:
		return "Failed"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "Status"))
	}
}

// Name gives a naam representation of the enum value
func (f Status) Name() naam.Name {
	switch f {
	case Status_Pending:
		return naam.New("pending")
	case Status_Running:
		return naam.New("running")
	case Status_Success:
		return naam.New("success")
	case Status_Failed:
		return naam.New("failed")
	default:
		panics.P("Status.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f Status) Value() (driver.Value, error) {
	switch f {
	case Status_INVALID:
		return nil, nil
	case Status_Pending:
		return "Pending", nil
	case Status_Running:
		return "Running", nil
	case Status_Success:
		return "Success", nil
	case Status_Failed:
		return "Failed", nil

	default:
		return nil, fmt.Errorf("Cannot save enum Status to DB: '%d' is not a valid value for enum Status", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *Status) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewStatusFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "Status")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f Status) ImplementsGraphQLType(name string) bool {
	return name == "Core_TaskRun_Status"
}

func (f *Status) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewStatusFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for Status: %T", input)
	}
	return err
}

func (f *Status) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum Status to a string: %w", err)
	}
	i := NewStatusFromString(enumStr)
	*f = i
	return nil
}

func (f Status) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil Status pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type StatusCondition struct {
	Op     filterlib.Operator
	Values []Status
}

func (c StatusCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c StatusCondition) Len() int {
	return len(c.Values)
}
func (c StatusCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskRunAction enum: <insert comment>
type TaskRunAction int

const (
	TaskRunAction_INVALID TaskRunAction = 0
	TaskRunAction_Run     TaskRunAction = 1
)

func NewTaskRunActionFromString(s string) TaskRunAction {
	switch s {
	case "INVALID":
		return TaskRunAction_INVALID
	case "Run":
		return TaskRunAction_Run

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TaskRunAction"))
	}
}

// String implements the `fmt.Stringer` interface for TaskRunAction. It allows us to print the enum values as strings.
func (f TaskRunAction) String() string {
	switch f {
	case TaskRunAction_INVALID:
		return "INVALID"
	case TaskRunAction_Run:
		return "Run"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TaskRunAction"))
	}
}

// Name gives a naam representation of the enum value
func (f TaskRunAction) Name() naam.Name {
	switch f {
	case TaskRunAction_Run:
		return naam.New("run")
	default:
		panics.P("TaskRunAction.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TaskRunAction) Value() (driver.Value, error) {
	switch f {
	case TaskRunAction_INVALID:
		return nil, nil
	case TaskRunAction_Run:
		return "Run", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TaskRunAction to DB: '%d' is not a valid value for enum TaskRunAction", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TaskRunAction) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTaskRunActionFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TaskRunAction")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TaskRunAction) ImplementsGraphQLType(name string) bool {
	return name == "Core_TaskRun_TaskRunAction"
}

func (f *TaskRunAction) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTaskRunActionFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TaskRunAction: %T", input)
	}
	return err
}

func (f *TaskRunAction) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TaskRunAction to a string: %w", err)
	}
	i := NewTaskRunActionFromString(enumStr)
	*f = i
	return nil
}

func (f TaskRunAction) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TaskRunAction pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TaskRunActionCondition struct {
	Op     filterlib.Operator
	Values []TaskRunAction
}

func (c TaskRunActionCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TaskRunActionCondition) Len() int {
	return len(c.Values)
}
func (c TaskRunActionCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskRunField enum: <insert comment>
type TaskRunField int

const (
	TaskRunField_INVALID            TaskRunField = 0
	TaskRunField_ID                 TaskRunField = 1
	TaskRunField_StartedAt          TaskRunField = 2
	TaskRunField_CompletedAt        TaskRunField = 3
	TaskRunField_Status             TaskRunField = 4
	TaskRunField_MethodRequestData  TaskRunField = 5
	TaskRunField_MethodResponseData TaskRunField = 6
	TaskRunField_TriggeredBy        TaskRunField = 7
	TaskRunField_CreatedAt          TaskRunField = 8
	TaskRunField_UpdatedAt          TaskRunField = 9
	TaskRunField_DeletedAt          TaskRunField = 10
	TaskRunField_TaskID             TaskRunField = 11
)

func NewTaskRunFieldFromString(s string) TaskRunField {
	switch s {
	case "INVALID":
		return TaskRunField_INVALID
	case "ID":
		return TaskRunField_ID
	case "StartedAt":
		return TaskRunField_StartedAt
	case "CompletedAt":
		return TaskRunField_CompletedAt
	case "Status":
		return TaskRunField_Status
	case "MethodRequestData":
		return TaskRunField_MethodRequestData
	case "MethodResponseData":
		return TaskRunField_MethodResponseData
	case "TriggeredBy":
		return TaskRunField_TriggeredBy
	case "CreatedAt":
		return TaskRunField_CreatedAt
	case "UpdatedAt":
		return TaskRunField_UpdatedAt
	case "DeletedAt":
		return TaskRunField_DeletedAt
	case "TaskID":
		return TaskRunField_TaskID

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TaskRunField"))
	}
}

// String implements the `fmt.Stringer` interface for TaskRunField. It allows us to print the enum values as strings.
func (f TaskRunField) String() string {
	switch f {
	case TaskRunField_INVALID:
		return "INVALID"
	case TaskRunField_ID:
		return "ID"
	case TaskRunField_StartedAt:
		return "StartedAt"
	case TaskRunField_CompletedAt:
		return "CompletedAt"
	case TaskRunField_Status:
		return "Status"
	case TaskRunField_MethodRequestData:
		return "MethodRequestData"
	case TaskRunField_MethodResponseData:
		return "MethodResponseData"
	case TaskRunField_TriggeredBy:
		return "TriggeredBy"
	case TaskRunField_CreatedAt:
		return "CreatedAt"
	case TaskRunField_UpdatedAt:
		return "UpdatedAt"
	case TaskRunField_DeletedAt:
		return "DeletedAt"
	case TaskRunField_TaskID:
		return "TaskID"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TaskRunField"))
	}
}

// Name gives a naam representation of the enum value
func (f TaskRunField) Name() naam.Name {
	switch f {
	case TaskRunField_ID:
		return naam.New("id")
	case TaskRunField_StartedAt:
		return naam.New("started_at")
	case TaskRunField_CompletedAt:
		return naam.New("completed_at")
	case TaskRunField_Status:
		return naam.New("status")
	case TaskRunField_MethodRequestData:
		return naam.New("method_request_data")
	case TaskRunField_MethodResponseData:
		return naam.New("method_response_data")
	case TaskRunField_TriggeredBy:
		return naam.New("triggered_by")
	case TaskRunField_CreatedAt:
		return naam.New("created_at")
	case TaskRunField_UpdatedAt:
		return naam.New("updated_at")
	case TaskRunField_DeletedAt:
		return naam.New("deleted_at")
	case TaskRunField_TaskID:
		return naam.New("task_id")
	default:
		panics.P("TaskRunField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TaskRunField) Value() (driver.Value, error) {
	switch f {
	case TaskRunField_INVALID:
		return nil, nil
	case TaskRunField_ID:
		return "ID", nil
	case TaskRunField_StartedAt:
		return "StartedAt", nil
	case TaskRunField_CompletedAt:
		return "CompletedAt", nil
	case TaskRunField_Status:
		return "Status", nil
	case TaskRunField_MethodRequestData:
		return "MethodRequestData", nil
	case TaskRunField_MethodResponseData:
		return "MethodResponseData", nil
	case TaskRunField_TriggeredBy:
		return "TriggeredBy", nil
	case TaskRunField_CreatedAt:
		return "CreatedAt", nil
	case TaskRunField_UpdatedAt:
		return "UpdatedAt", nil
	case TaskRunField_DeletedAt:
		return "DeletedAt", nil
	case TaskRunField_TaskID:
		return "TaskID", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TaskRunField to DB: '%d' is not a valid value for enum TaskRunField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TaskRunField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTaskRunFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TaskRunField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TaskRunField) ImplementsGraphQLType(name string) bool {
	return name == "Core_TaskRun_TaskRunField"
}

func (f *TaskRunField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTaskRunFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TaskRunField: %T", input)
	}
	return err
}

func (f *TaskRunField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TaskRunField to a string: %w", err)
	}
	i := NewTaskRunFieldFromString(enumStr)
	*f = i
	return nil
}

func (f TaskRunField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TaskRunField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TaskRunFieldCondition struct {
	Op     filterlib.Operator
	Values []TaskRunField
}

func (c TaskRunFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TaskRunFieldCondition) Len() int {
	return len(c.Values)
}
func (c TaskRunFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TriggeredBy enum: <insert comment>
type TriggeredBy int

const (
	TriggeredBy_INVALID TriggeredBy = 0
	TriggeredBy_Cron    TriggeredBy = 1
	TriggeredBy_Manual  TriggeredBy = 2
)

func NewTriggeredByFromString(s string) TriggeredBy {
	switch s {
	case "INVALID":
		return TriggeredBy_INVALID
	case "Cron":
		return TriggeredBy_Cron
	case "Manual":
		return TriggeredBy_Manual

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TriggeredBy"))
	}
}

// String implements the `fmt.Stringer` interface for TriggeredBy. It allows us to print the enum values as strings.
func (f TriggeredBy) String() string {
	switch f {
	case TriggeredBy_INVALID:
		return "INVALID"
	case TriggeredBy_Cron:
		return "Cron"
	case TriggeredBy_Manual:
		return "Manual"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TriggeredBy"))
	}
}

// Name gives a naam representation of the enum value
func (f TriggeredBy) Name() naam.Name {
	switch f {
	case TriggeredBy_Cron:
		return naam.New("cron")
	case TriggeredBy_Manual:
		return naam.New("manual")
	default:
		panics.P("TriggeredBy.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TriggeredBy) Value() (driver.Value, error) {
	switch f {
	case TriggeredBy_INVALID:
		return nil, nil
	case TriggeredBy_Cron:
		return "Cron", nil
	case TriggeredBy_Manual:
		return "Manual", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TriggeredBy to DB: '%d' is not a valid value for enum TriggeredBy", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TriggeredBy) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTriggeredByFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TriggeredBy")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TriggeredBy) ImplementsGraphQLType(name string) bool {
	return name == "Core_TaskRun_TriggeredBy"
}

func (f *TriggeredBy) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTriggeredByFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TriggeredBy: %T", input)
	}
	return err
}

func (f *TriggeredBy) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TriggeredBy to a string: %w", err)
	}
	i := NewTriggeredByFromString(enumStr)
	*f = i
	return nil
}

func (f TriggeredBy) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TriggeredBy pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TriggeredByCondition struct {
	Op     filterlib.Operator
	Values []TriggeredBy
}

func (c TriggeredByCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TriggeredByCondition) Len() int {
	return len(c.Values)
}
func (c TriggeredByCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TaskRunFilter: <comments>
type TaskRunFilter struct {
	ID                 *filterlib.IDCondition          `json:"id"`
	StartedAt          *filterlib.TimestampCondition   `json:"startedAt"`
	CompletedAt        *filterlib.TimestampCondition   `json:"completedAt"`
	Status             *StatusCondition                `json:"status"`
	MethodRequestData  *filterlib.GenericDataCondition `json:"methodRequestData"`
	MethodResponseData *filterlib.GenericDataCondition `json:"methodResponseData"`
	TriggeredBy        *TriggeredByCondition           `json:"triggeredBy"`
	CreatedAt          *filterlib.TimestampCondition   `json:"createdAt"`
	UpdatedAt          *filterlib.TimestampCondition   `json:"updatedAt"`
	DeletedAt          *filterlib.TimestampCondition   `json:"deletedAt"`
	TaskID             *filterlib.IDCondition          `json:"taskID"`
	And                []TaskRunFilter                 `json:"and"`
	Or                 []TaskRunFilter                 `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object TaskRun `json:"object"`
}

// AddRequest:
type AddRequest = dalutil.EntityAddRequest[TaskRunInput]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[TaskRun]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[TaskRunFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[TaskRun]

// MethodNameCondition:
// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[TaskRun]

// StandardEntityRequest:
// StandardEntityResponse:
// StatusCondition:
// TaskRun: Main type for entity TaskRun
// TaskRunActionCondition:
// TaskRunFieldCondition:
// TaskRunFilter:
// TriggeredByCondition:
// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[TaskRun, TaskRunField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[TaskRun]
