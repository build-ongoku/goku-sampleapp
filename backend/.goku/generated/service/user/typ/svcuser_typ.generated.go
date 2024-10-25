package svcuser_typ

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/panics"
)

// EntityName enum: <insert comment>
type EntityName int

const (
	EntityName_INVALID EntityName = 0
	EntityName_User    EntityName = 1
	EntityName_Team    EntityName = 2
)

func NewEntityNameFromString(s string) EntityName {
	switch s {
	case "INVALID":
		return EntityName_INVALID
	case "User":
		return EntityName_User
	case "Team":
		return EntityName_Team

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "EntityName"))
	}
}

// String implements the `fmt.Stringer` interface for EntityName. It allows us to print the enum values as strings.
func (f EntityName) String() string {
	switch f {
	case EntityName_INVALID:
		return "INVALID"
	case EntityName_User:
		return "User"
	case EntityName_Team:
		return "Team"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "EntityName"))
	}
}

// Name gives a naam representation of the enum value
func (f EntityName) Name() naam.Name {
	switch f {
	case EntityName_User:
		return naam.New("user")
	case EntityName_Team:
		return naam.New("team")
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
	case EntityName_User:
		return "User", nil
	case EntityName_Team:
		return "Team", nil

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
	return name == "User_EntityName"
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

// EntityNameCondition:
