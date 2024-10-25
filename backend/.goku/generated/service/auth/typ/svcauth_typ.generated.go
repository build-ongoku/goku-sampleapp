package svcauth_typ

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/panics"
	"github.com/teejays/gokutil/scalars"

	app_typ "sampleapp/backend/.goku/generated/typ"
)

// EntityName enum: <insert comment>
type EntityName int

const (
	EntityName_INVALID EntityName = 0
	EntityName_Secret  EntityName = 1
)

func NewEntityNameFromString(s string) EntityName {
	switch s {
	case "INVALID":
		return EntityName_INVALID
	case "Secret":
		return EntityName_Secret

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "EntityName"))
	}
}

// String implements the `fmt.Stringer` interface for EntityName. It allows us to print the enum values as strings.
func (f EntityName) String() string {
	switch f {
	case EntityName_INVALID:
		return "INVALID"
	case EntityName_Secret:
		return "Secret"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "EntityName"))
	}
}

// Name gives a naam representation of the enum value
func (f EntityName) Name() naam.Name {
	switch f {
	case EntityName_Secret:
		return naam.New("secret")
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
	case EntityName_Secret:
		return "Secret", nil

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
	return name == "Auth_EntityName"
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
	MethodName_INVALID           MethodName = 0
	MethodName_LoginUser         MethodName = 1
	MethodName_RegisterUser      MethodName = 2
	MethodName_AuthenticateToken MethodName = 3
)

func NewMethodNameFromString(s string) MethodName {
	switch s {
	case "INVALID":
		return MethodName_INVALID
	case "LoginUser":
		return MethodName_LoginUser
	case "RegisterUser":
		return MethodName_RegisterUser
	case "AuthenticateToken":
		return MethodName_AuthenticateToken

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "MethodName"))
	}
}

// String implements the `fmt.Stringer` interface for MethodName. It allows us to print the enum values as strings.
func (f MethodName) String() string {
	switch f {
	case MethodName_INVALID:
		return "INVALID"
	case MethodName_LoginUser:
		return "LoginUser"
	case MethodName_RegisterUser:
		return "RegisterUser"
	case MethodName_AuthenticateToken:
		return "AuthenticateToken"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "MethodName"))
	}
}

// Name gives a naam representation of the enum value
func (f MethodName) Name() naam.Name {
	switch f {
	case MethodName_LoginUser:
		return naam.New("login_user")
	case MethodName_RegisterUser:
		return naam.New("register_user")
	case MethodName_AuthenticateToken:
		return naam.New("authenticate_token")
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
	case MethodName_LoginUser:
		return "LoginUser", nil
	case MethodName_RegisterUser:
		return "RegisterUser", nil
	case MethodName_AuthenticateToken:
		return "AuthenticateToken", nil

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
	return name == "Auth_MethodName"
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

// AuthenticateResponse: <comments>
type AuthenticateResponse struct {
	Token string `json:"token" validate:"required"`
}

// AuthenticateTokenRequest: <comments>
type AuthenticateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

// LoginRequest: <comments>
type LoginRequest struct {
	Email    scalars.Email `json:"email" validate:"required"`
	Password string        `json:"password" validate:"required"`
}

// RegisterUserRequest: <comments>
type RegisterUserRequest struct {
	Email    scalars.Email           `json:"email" validate:"required"`
	Name     app_typ.PersonNameInput `json:"name" validate:"required"`
	Password string                  `json:"password" validate:"required"`
	TeamID   scalars.ID              `json:"teamID" validate:"required,uuid"`
}

// AuthenticateResponse:
// AuthenticateTokenRequest:
// EntityNameCondition:
// LoginRequest:
// MethodNameCondition:
// RegisterUserRequest:
