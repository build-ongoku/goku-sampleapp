package svccore_entjobapplicant_typ

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

// JobApplicant: <comments>
type JobApplicant struct {
	ID        scalars.ID         `json:"id"`
	Name      string             `json:"name" validate:"required"`
	UserID    scalars.ID         `json:"userID" validate:"required"`
	ResumeID  scalars.ID         `json:"resumeID" validate:"required"`
	CreatedAt scalars.Timestamp  `json:"createdAt"`
	UpdatedAt scalars.Timestamp  `json:"updatedAt"`
	DeletedAt *scalars.Timestamp `json:"deletedAt"`
}

func (t JobApplicant) GetID() scalars.ID {
	return t.ID
}
func (t JobApplicant) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t JobApplicant) SetUpdatedAt(tim scalars.Timestamp) JobApplicant {
	t.UpdatedAt = tim
	return t
}

// JobApplicantInput: <comments>
type JobApplicantInput struct {
	Name     string     `json:"name" validate:"required"`
	UserID   scalars.ID `json:"userID" validate:"required"`
	ResumeID scalars.ID `json:"resumeID" validate:"required"`
}

func NewJobApplicantWithMetaArrayFromInputs(ctx context.Context, ins []JobApplicantInput) []JobApplicant {
	var outs []JobApplicant
	for _, in := range ins {
		outs = append(outs, NewJobApplicantWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewJobApplicantWithMetaFromInput(ctx context.Context, in JobApplicantInput) JobApplicant {
	return JobApplicant{
		Name:     in.Name,
		UserID:   in.UserID,
		ResumeID: in.ResumeID,
	}
}

// JobApplicantField enum: <insert comment>
type JobApplicantField int

const (
	JobApplicantField_INVALID   JobApplicantField = 0
	JobApplicantField_ID        JobApplicantField = 1
	JobApplicantField_Name      JobApplicantField = 2
	JobApplicantField_UserID    JobApplicantField = 3
	JobApplicantField_ResumeID  JobApplicantField = 4
	JobApplicantField_CreatedAt JobApplicantField = 5
	JobApplicantField_UpdatedAt JobApplicantField = 6
	JobApplicantField_DeletedAt JobApplicantField = 7
)

func NewJobApplicantFieldFromString(s string) JobApplicantField {
	switch s {
	case "INVALID":
		return JobApplicantField_INVALID
	case "ID":
		return JobApplicantField_ID
	case "Name":
		return JobApplicantField_Name
	case "UserID":
		return JobApplicantField_UserID
	case "ResumeID":
		return JobApplicantField_ResumeID
	case "CreatedAt":
		return JobApplicantField_CreatedAt
	case "UpdatedAt":
		return JobApplicantField_UpdatedAt
	case "DeletedAt":
		return JobApplicantField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "JobApplicantField"))
	}
}

// String implements the `fmt.Stringer` interface for JobApplicantField. It allows us to print the enum values as strings.
func (f JobApplicantField) String() string {
	switch f {
	case JobApplicantField_INVALID:
		return "INVALID"
	case JobApplicantField_ID:
		return "ID"
	case JobApplicantField_Name:
		return "Name"
	case JobApplicantField_UserID:
		return "UserID"
	case JobApplicantField_ResumeID:
		return "ResumeID"
	case JobApplicantField_CreatedAt:
		return "CreatedAt"
	case JobApplicantField_UpdatedAt:
		return "UpdatedAt"
	case JobApplicantField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "JobApplicantField"))
	}
}

// Name gives a naam representation of the enum value
func (f JobApplicantField) Name() naam.Name {
	switch f {
	case JobApplicantField_ID:
		return naam.New("id")
	case JobApplicantField_Name:
		return naam.New("name")
	case JobApplicantField_UserID:
		return naam.New("user_id")
	case JobApplicantField_ResumeID:
		return naam.New("resume_id")
	case JobApplicantField_CreatedAt:
		return naam.New("created_at")
	case JobApplicantField_UpdatedAt:
		return naam.New("updated_at")
	case JobApplicantField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("JobApplicantField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f JobApplicantField) Value() (driver.Value, error) {
	switch f {
	case JobApplicantField_INVALID:
		return nil, nil
	case JobApplicantField_ID:
		return "ID", nil
	case JobApplicantField_Name:
		return "Name", nil
	case JobApplicantField_UserID:
		return "UserID", nil
	case JobApplicantField_ResumeID:
		return "ResumeID", nil
	case JobApplicantField_CreatedAt:
		return "CreatedAt", nil
	case JobApplicantField_UpdatedAt:
		return "UpdatedAt", nil
	case JobApplicantField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum JobApplicantField to DB: '%d' is not a valid value for enum JobApplicantField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *JobApplicantField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewJobApplicantFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "JobApplicantField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f JobApplicantField) ImplementsGraphQLType(name string) bool {
	return name == "Core_JobApplicant_JobApplicantField"
}

func (f *JobApplicantField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewJobApplicantFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for JobApplicantField: %T", input)
	}
	return err
}

func (f *JobApplicantField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum JobApplicantField to a string: %w", err)
	}
	i := NewJobApplicantFieldFromString(enumStr)
	*f = i
	return nil
}

func (f JobApplicantField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil JobApplicantField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type JobApplicantFieldCondition struct {
	Op     filterlib.Operator
	Values []JobApplicantField
}

func (c JobApplicantFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c JobApplicantFieldCondition) Len() int {
	return len(c.Values)
}
func (c JobApplicantFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
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
	return name == "Core_JobApplicant_MethodName"
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

// JobApplicantFilter: <comments>
type JobApplicantFilter struct {
	ID        *filterlib.IDCondition        `json:"id"`
	Name      *filterlib.StringCondition    `json:"name"`
	UserID    *filterlib.IDCondition        `json:"userID"`
	ResumeID  *filterlib.IDCondition        `json:"resumeID"`
	CreatedAt *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt *filterlib.TimestampCondition `json:"deletedAt"`
	And       []JobApplicantFilter          `json:"and"`
	Or        []JobApplicantFilter          `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object JobApplicant `json:"object"`
}

// AddRequest:
type AddRequest = dalutil.EntityAddRequest[JobApplicantInput]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[JobApplicant]

// JobApplicant: Main type for entity JobApplicant
// JobApplicantFieldCondition:
// JobApplicantFilter:
// ListRequest:
type ListRequest = dalutil.ListEntityRequest[JobApplicantFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[JobApplicant]

// MethodNameCondition:
// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[JobApplicant]

// StandardEntityRequest:
// StandardEntityResponse:
// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[JobApplicant, JobApplicantField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[JobApplicant]
