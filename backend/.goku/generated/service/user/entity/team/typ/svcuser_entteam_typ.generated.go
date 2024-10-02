package svcuser_entteam_typ

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

// Team: <comments>
type Team struct {
	ID        scalars.ID         `json:"id"`
	Name      string             `json:"name" validate:"required"`
	CreatedAt scalars.Timestamp  `json:"createdAt"`
	UpdatedAt scalars.Timestamp  `json:"updatedAt"`
	DeletedAt *scalars.Timestamp `json:"deletedAt"`
}

func (t Team) GetID() scalars.ID {
	return t.ID
}
func (t Team) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t Team) SetUpdatedAt(tim scalars.Timestamp) {
	t.UpdatedAt = tim
}

// TeamInput: <comments>
type TeamInput struct {
	Name string `json:"name" validate:"required"`
}

func NewTeamWithMetaArrayFromInputs(ctx context.Context, ins []TeamInput) []Team {
	var outs []Team
	for _, in := range ins {
		outs = append(outs, NewTeamWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewTeamWithMetaFromInput(ctx context.Context, in TeamInput) Team {
	return Team{
		Name: in.Name,
	}
}

// TeamField enum: <insert comment>
type TeamField int

const (
	TeamField_INVALID   TeamField = 0
	TeamField_ID        TeamField = 1
	TeamField_Name      TeamField = 2
	TeamField_CreatedAt TeamField = 3
	TeamField_UpdatedAt TeamField = 4
	TeamField_DeletedAt TeamField = 5
)

func NewTeamFieldFromString(s string) TeamField {
	switch s {
	case "INVALID":
		return TeamField_INVALID
	case "ID":
		return TeamField_ID
	case "Name":
		return TeamField_Name
	case "CreatedAt":
		return TeamField_CreatedAt
	case "UpdatedAt":
		return TeamField_UpdatedAt
	case "DeletedAt":
		return TeamField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "TeamField"))
	}
}

// String implements the `fmt.Stringer` interface for TeamField. It allows us to print the enum values as strings.
func (f TeamField) String() string {
	switch f {
	case TeamField_INVALID:
		return "INVALID"
	case TeamField_ID:
		return "ID"
	case TeamField_Name:
		return "Name"
	case TeamField_CreatedAt:
		return "CreatedAt"
	case TeamField_UpdatedAt:
		return "UpdatedAt"
	case TeamField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "TeamField"))
	}
}

// Name gives a naam representation of the enum value
func (f TeamField) Name() naam.Name {
	switch f {
	case TeamField_ID:
		return naam.New("id")
	case TeamField_Name:
		return naam.New("name")
	case TeamField_CreatedAt:
		return naam.New("created_at")
	case TeamField_UpdatedAt:
		return naam.New("updated_at")
	case TeamField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("TeamField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f TeamField) Value() (driver.Value, error) {
	switch f {
	case TeamField_INVALID:
		return nil, nil
	case TeamField_ID:
		return "ID", nil
	case TeamField_Name:
		return "Name", nil
	case TeamField_CreatedAt:
		return "CreatedAt", nil
	case TeamField_UpdatedAt:
		return "UpdatedAt", nil
	case TeamField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum TeamField to DB: '%d' is not a valid value for enum TeamField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *TeamField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTeamFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "TeamField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f TeamField) ImplementsGraphQLType(name string) bool {
	return name == "User_Team_TeamField"
}

func (f *TeamField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTeamFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for TeamField: %T", input)
	}
	return err
}

func (f *TeamField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum TeamField to a string: %w", err)
	}
	i := NewTeamFieldFromString(enumStr)
	*f = i
	return nil
}

func (f TeamField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil TeamField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TeamFieldCondition struct {
	Op     filterlib.Operator
	Values []TeamField
}

func (c TeamFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TeamFieldCondition) Len() int {
	return len(c.Values)
}
func (c TeamFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// TeamFilter: <comments>
type TeamFilter struct {
	ID        *filterlib.IDCondition        `json:"id"`
	Name      *filterlib.StringCondition    `json:"name"`
	CreatedAt *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt *filterlib.TimestampCondition `json:"deletedAt"`
	And       []TeamFilter                  `json:"and"`
	Or        []TeamFilter                  `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object Team `json:"object"`
}

// StandardEntityRequest:
// StandardEntityResponse:
// Team:
// TeamFieldCondition:
// TeamFilter:
// AddRequest:
type AddRequest = dalutil.EntityAddRequest[TeamInput]

// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[Team, TeamField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[Team]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[Team]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[TeamFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[Team]

// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[Team]
