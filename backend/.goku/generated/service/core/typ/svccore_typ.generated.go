package svccore_typ

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/panics"
	"github.com/teejays/gokutil/scalars"
)

// CronExpression: <comments>
type CronExpression struct {
	Second     string `json:"second" validate:"required"`
	Minute     string `json:"minute" validate:"required"`
	Hour       string `json:"hour" validate:"required"`
	DayOfMonth string `json:"dayOfMonth" validate:"required"`
	Month      string `json:"month" validate:"required"`
	DayOfWeek  string `json:"dayOfWeek" validate:"required"`
}

// CronExpressionWithMeta: <comments>
type CronExpressionWithMeta struct {
	ParentID   scalars.ID         `json:"parentID" json:"-"`
	ID         scalars.ID         `json:"id"`
	Second     string             `json:"second" validate:"required"`
	Minute     string             `json:"minute" validate:"required"`
	Hour       string             `json:"hour" validate:"required"`
	DayOfMonth string             `json:"dayOfMonth" validate:"required"`
	Month      string             `json:"month" validate:"required"`
	DayOfWeek  string             `json:"dayOfWeek" validate:"required"`
	CreatedAt  scalars.Timestamp  `json:"createdAt"`
	UpdatedAt  scalars.Timestamp  `json:"updatedAt"`
	DeletedAt  *scalars.Timestamp `json:"deletedAt"`
}

func (t CronExpressionWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t CronExpressionWithMeta) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t *CronExpressionWithMeta) SetUpdatedAt(tim scalars.Timestamp) {
	t.UpdatedAt = tim
}

// CronExpressionInput: <comments>
type CronExpressionInput struct {
	Second     string `json:"second" validate:"required"`
	Minute     string `json:"minute" validate:"required"`
	Hour       string `json:"hour" validate:"required"`
	DayOfMonth string `json:"dayOfMonth" validate:"required"`
	Month      string `json:"month" validate:"required"`
	DayOfWeek  string `json:"dayOfWeek" validate:"required"`
}

func NewCronExpressionWithMetaArrayFromInputs(ctx context.Context, ins []CronExpressionInput) []CronExpressionWithMeta {
	var outs []CronExpressionWithMeta
	for _, in := range ins {
		outs = append(outs, NewCronExpressionWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewCronExpressionWithMetaFromInput(ctx context.Context, in CronExpressionInput) CronExpressionWithMeta {
	return CronExpressionWithMeta{
		Second:     in.Second,
		Minute:     in.Minute,
		Hour:       in.Hour,
		DayOfMonth: in.DayOfMonth,
		Month:      in.Month,
		DayOfWeek:  in.DayOfWeek,
	}
}

// CronExpressionField enum: <insert comment>
type CronExpressionField int

const (
	CronExpressionField_INVALID    CronExpressionField = 0
	CronExpressionField_ParentID   CronExpressionField = 1
	CronExpressionField_ID         CronExpressionField = 2
	CronExpressionField_Second     CronExpressionField = 3
	CronExpressionField_Minute     CronExpressionField = 4
	CronExpressionField_Hour       CronExpressionField = 5
	CronExpressionField_DayOfMonth CronExpressionField = 6
	CronExpressionField_Month      CronExpressionField = 7
	CronExpressionField_DayOfWeek  CronExpressionField = 8
	CronExpressionField_CreatedAt  CronExpressionField = 9
	CronExpressionField_UpdatedAt  CronExpressionField = 10
	CronExpressionField_DeletedAt  CronExpressionField = 11
)

func NewCronExpressionFieldFromString(s string) CronExpressionField {
	switch s {
	case "INVALID":
		return CronExpressionField_INVALID
	case "ParentID":
		return CronExpressionField_ParentID
	case "ID":
		return CronExpressionField_ID
	case "Second":
		return CronExpressionField_Second
	case "Minute":
		return CronExpressionField_Minute
	case "Hour":
		return CronExpressionField_Hour
	case "DayOfMonth":
		return CronExpressionField_DayOfMonth
	case "Month":
		return CronExpressionField_Month
	case "DayOfWeek":
		return CronExpressionField_DayOfWeek
	case "CreatedAt":
		return CronExpressionField_CreatedAt
	case "UpdatedAt":
		return CronExpressionField_UpdatedAt
	case "DeletedAt":
		return CronExpressionField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "CronExpressionField"))
	}
}

// String implements the `fmt.Stringer` interface for CronExpressionField. It allows us to print the enum values as strings.
func (f CronExpressionField) String() string {
	switch f {
	case CronExpressionField_INVALID:
		return "INVALID"
	case CronExpressionField_ParentID:
		return "ParentID"
	case CronExpressionField_ID:
		return "ID"
	case CronExpressionField_Second:
		return "Second"
	case CronExpressionField_Minute:
		return "Minute"
	case CronExpressionField_Hour:
		return "Hour"
	case CronExpressionField_DayOfMonth:
		return "DayOfMonth"
	case CronExpressionField_Month:
		return "Month"
	case CronExpressionField_DayOfWeek:
		return "DayOfWeek"
	case CronExpressionField_CreatedAt:
		return "CreatedAt"
	case CronExpressionField_UpdatedAt:
		return "UpdatedAt"
	case CronExpressionField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "CronExpressionField"))
	}
}

// Name gives a naam representation of the enum value
func (f CronExpressionField) Name() naam.Name {
	switch f {
	case CronExpressionField_ParentID:
		return naam.New("parent_id")
	case CronExpressionField_ID:
		return naam.New("id")
	case CronExpressionField_Second:
		return naam.New("second")
	case CronExpressionField_Minute:
		return naam.New("minute")
	case CronExpressionField_Hour:
		return naam.New("hour")
	case CronExpressionField_DayOfMonth:
		return naam.New("day_of_month")
	case CronExpressionField_Month:
		return naam.New("month")
	case CronExpressionField_DayOfWeek:
		return naam.New("day_of_week")
	case CronExpressionField_CreatedAt:
		return naam.New("created_at")
	case CronExpressionField_UpdatedAt:
		return naam.New("updated_at")
	case CronExpressionField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("CronExpressionField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f CronExpressionField) Value() (driver.Value, error) {
	switch f {
	case CronExpressionField_INVALID:
		return nil, nil
	case CronExpressionField_ParentID:
		return "ParentID", nil
	case CronExpressionField_ID:
		return "ID", nil
	case CronExpressionField_Second:
		return "Second", nil
	case CronExpressionField_Minute:
		return "Minute", nil
	case CronExpressionField_Hour:
		return "Hour", nil
	case CronExpressionField_DayOfMonth:
		return "DayOfMonth", nil
	case CronExpressionField_Month:
		return "Month", nil
	case CronExpressionField_DayOfWeek:
		return "DayOfWeek", nil
	case CronExpressionField_CreatedAt:
		return "CreatedAt", nil
	case CronExpressionField_UpdatedAt:
		return "UpdatedAt", nil
	case CronExpressionField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum CronExpressionField to DB: '%d' is not a valid value for enum CronExpressionField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *CronExpressionField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewCronExpressionFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "CronExpressionField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f CronExpressionField) ImplementsGraphQLType(name string) bool {
	return name == "Core_CronExpressionField"
}

func (f *CronExpressionField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewCronExpressionFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for CronExpressionField: %T", input)
	}
	return err
}

func (f *CronExpressionField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum CronExpressionField to a string: %w", err)
	}
	i := NewCronExpressionFieldFromString(enumStr)
	*f = i
	return nil
}

func (f CronExpressionField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil CronExpressionField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type CronExpressionFieldCondition struct {
	Op     filterlib.Operator
	Values []CronExpressionField
}

func (c CronExpressionFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c CronExpressionFieldCondition) Len() int {
	return len(c.Values)
}
func (c CronExpressionFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// EntityName enum: <insert comment>
type EntityName int

const (
	EntityName_INVALID EntityName = 0
	EntityName_File    EntityName = 1
	EntityName_Task    EntityName = 2
	EntityName_TaskRun EntityName = 3
)

func NewEntityNameFromString(s string) EntityName {
	switch s {
	case "INVALID":
		return EntityName_INVALID
	case "File":
		return EntityName_File
	case "Task":
		return EntityName_Task
	case "TaskRun":
		return EntityName_TaskRun

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "EntityName"))
	}
}

// String implements the `fmt.Stringer` interface for EntityName. It allows us to print the enum values as strings.
func (f EntityName) String() string {
	switch f {
	case EntityName_INVALID:
		return "INVALID"
	case EntityName_File:
		return "File"
	case EntityName_Task:
		return "Task"
	case EntityName_TaskRun:
		return "TaskRun"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "EntityName"))
	}
}

// Name gives a naam representation of the enum value
func (f EntityName) Name() naam.Name {
	switch f {
	case EntityName_File:
		return naam.New("file")
	case EntityName_Task:
		return naam.New("task")
	case EntityName_TaskRun:
		return naam.New("task_run")
	default:
		panics.P("EntityName.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f EntityName) Value() (driver.Value, error) {
	switch f {
	case EntityName_INVALID:
		return nil, nil
	case EntityName_File:
		return "File", nil
	case EntityName_Task:
		return "Task", nil
	case EntityName_TaskRun:
		return "TaskRun", nil

	default:
		return nil, fmt.Errorf("Cannot save enum EntityName to DB: '%d' is not a valid value for enum EntityName", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *EntityName) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewEntityNameFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "EntityName")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f EntityName) ImplementsGraphQLType(name string) bool {
	return name == "Core_EntityName"
}

func (f *EntityName) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewEntityNameFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for EntityName: %T", input)
	}
	return err
}

func (f *EntityName) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum EntityName to a string: %w", err)
	}
	i := NewEntityNameFromString(enumStr)
	*f = i
	return nil
}

func (f EntityName) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil EntityName pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type EntityNameCondition struct {
	Op     filterlib.Operator
	Values []EntityName
}

func (c EntityNameCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c EntityNameCondition) Len() int {
	return len(c.Values)
}
func (c EntityNameCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// MethodName enum: <insert comment>
type MethodName int

const (
	MethodName_INVALID    MethodName = 0
	MethodName_FileUpload MethodName = 1
)

func NewMethodNameFromString(s string) MethodName {
	switch s {
	case "INVALID":
		return MethodName_INVALID
	case "FileUpload":
		return MethodName_FileUpload

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "MethodName"))
	}
}

// String implements the `fmt.Stringer` interface for MethodName. It allows us to print the enum values as strings.
func (f MethodName) String() string {
	switch f {
	case MethodName_INVALID:
		return "INVALID"
	case MethodName_FileUpload:
		return "FileUpload"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "MethodName"))
	}
}

// Name gives a naam representation of the enum value
func (f MethodName) Name() naam.Name {
	switch f {
	case MethodName_FileUpload:
		return naam.New("file_upload")
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
	case MethodName_FileUpload:
		return "FileUpload", nil

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
	return name == "Core_MethodName"
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

// CronExpressionFilter: <comments>
type CronExpressionFilter struct {
	ParentID   *filterlib.IDCondition        `json:"parentID"`
	ID         *filterlib.IDCondition        `json:"id"`
	Second     *filterlib.StringCondition    `json:"second"`
	Minute     *filterlib.StringCondition    `json:"minute"`
	Hour       *filterlib.StringCondition    `json:"hour"`
	DayOfMonth *filterlib.StringCondition    `json:"dayOfMonth"`
	Month      *filterlib.StringCondition    `json:"month"`
	DayOfWeek  *filterlib.StringCondition    `json:"dayOfWeek"`
	CreatedAt  *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt  *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt  *filterlib.TimestampCondition `json:"deletedAt"`
	And        []CronExpressionFilter        `json:"and"`
	Or         []CronExpressionFilter        `json:"or"`
}

// CronExpression:
// CronExpressionFieldCondition:
// CronExpressionFilter:
// EntityNameCondition:
// MethodNameCondition:
