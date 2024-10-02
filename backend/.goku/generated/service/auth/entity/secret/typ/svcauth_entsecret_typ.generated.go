package svcauth_entsecret_typ

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

// Secret: <comments>
type Secret struct {
	ID        scalars.ID         `json:"id"`
	UserID    scalars.ID         `json:"userID" validate:"required"`
	Type      Type               `json:"type" validate:"required"`
	Secret    string             `json:"secret" validate:"required"`
	ExpiresAt *scalars.Timestamp `json:"expiresAt"`
	CreatedAt scalars.Timestamp  `json:"createdAt"`
	UpdatedAt scalars.Timestamp  `json:"updatedAt"`
	DeletedAt *scalars.Timestamp `json:"deletedAt"`
}

func (t Secret) GetID() scalars.ID {
	return t.ID
}
func (t Secret) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t Secret) SetUpdatedAt(tim scalars.Timestamp) {
	t.UpdatedAt = tim
}

// SecretInput: <comments>
type SecretInput struct {
	UserID    scalars.ID         `json:"userID" validate:"required"`
	Type      Type               `json:"type" validate:"required"`
	Secret    string             `json:"secret" validate:"required"`
	ExpiresAt *scalars.Timestamp `json:"expiresAt"`
}

func NewSecretWithMetaArrayFromInputs(ctx context.Context, ins []SecretInput) []Secret {
	var outs []Secret
	for _, in := range ins {
		outs = append(outs, NewSecretWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewSecretWithMetaFromInput(ctx context.Context, in SecretInput) Secret {
	return Secret{
		UserID:    in.UserID,
		Type:      in.Type,
		Secret:    in.Secret,
		ExpiresAt: in.ExpiresAt,
	}
}

// SecretField enum: <insert comment>
type SecretField int

const (
	SecretField_INVALID   SecretField = 0
	SecretField_ID        SecretField = 1
	SecretField_UserID    SecretField = 2
	SecretField_Type      SecretField = 3
	SecretField_Secret    SecretField = 4
	SecretField_ExpiresAt SecretField = 5
	SecretField_CreatedAt SecretField = 6
	SecretField_UpdatedAt SecretField = 7
	SecretField_DeletedAt SecretField = 8
)

func NewSecretFieldFromString(s string) SecretField {
	switch s {
	case "INVALID":
		return SecretField_INVALID
	case "ID":
		return SecretField_ID
	case "UserID":
		return SecretField_UserID
	case "Type":
		return SecretField_Type
	case "Secret":
		return SecretField_Secret
	case "ExpiresAt":
		return SecretField_ExpiresAt
	case "CreatedAt":
		return SecretField_CreatedAt
	case "UpdatedAt":
		return SecretField_UpdatedAt
	case "DeletedAt":
		return SecretField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "SecretField"))
	}
}

// String implements the `fmt.Stringer` interface for SecretField. It allows us to print the enum values as strings.
func (f SecretField) String() string {
	switch f {
	case SecretField_INVALID:
		return "INVALID"
	case SecretField_ID:
		return "ID"
	case SecretField_UserID:
		return "UserID"
	case SecretField_Type:
		return "Type"
	case SecretField_Secret:
		return "Secret"
	case SecretField_ExpiresAt:
		return "ExpiresAt"
	case SecretField_CreatedAt:
		return "CreatedAt"
	case SecretField_UpdatedAt:
		return "UpdatedAt"
	case SecretField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "SecretField"))
	}
}

// Name gives a naam representation of the enum value
func (f SecretField) Name() naam.Name {
	switch f {
	case SecretField_ID:
		return naam.New("id")
	case SecretField_UserID:
		return naam.New("user_id")
	case SecretField_Type:
		return naam.New("type")
	case SecretField_Secret:
		return naam.New("secret")
	case SecretField_ExpiresAt:
		return naam.New("expires_at")
	case SecretField_CreatedAt:
		return naam.New("created_at")
	case SecretField_UpdatedAt:
		return naam.New("updated_at")
	case SecretField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("SecretField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f SecretField) Value() (driver.Value, error) {
	switch f {
	case SecretField_INVALID:
		return nil, nil
	case SecretField_ID:
		return "ID", nil
	case SecretField_UserID:
		return "UserID", nil
	case SecretField_Type:
		return "Type", nil
	case SecretField_Secret:
		return "Secret", nil
	case SecretField_ExpiresAt:
		return "ExpiresAt", nil
	case SecretField_CreatedAt:
		return "CreatedAt", nil
	case SecretField_UpdatedAt:
		return "UpdatedAt", nil
	case SecretField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum SecretField to DB: '%d' is not a valid value for enum SecretField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *SecretField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewSecretFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "SecretField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f SecretField) ImplementsGraphQLType(name string) bool {
	return name == "Auth_Secret_SecretField"
}

func (f *SecretField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewSecretFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for SecretField: %T", input)
	}
	return err
}

func (f *SecretField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum SecretField to a string: %w", err)
	}
	i := NewSecretFieldFromString(enumStr)
	*f = i
	return nil
}

func (f SecretField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil SecretField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type SecretFieldCondition struct {
	Op     filterlib.Operator
	Values []SecretField
}

func (c SecretFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c SecretFieldCondition) Len() int {
	return len(c.Values)
}
func (c SecretFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// Type enum: <insert comment>
type Type int

const (
	Type_INVALID     Type = 0
	Type_Password    Type = 1
	Type_GithubToken Type = 2
)

func NewTypeFromString(s string) Type {
	switch s {
	case "INVALID":
		return Type_INVALID
	case "Password":
		return Type_Password
	case "GithubToken":
		return Type_GithubToken

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "Type"))
	}
}

// String implements the `fmt.Stringer` interface for Type. It allows us to print the enum values as strings.
func (f Type) String() string {
	switch f {
	case Type_INVALID:
		return "INVALID"
	case Type_Password:
		return "Password"
	case Type_GithubToken:
		return "GithubToken"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "Type"))
	}
}

// Name gives a naam representation of the enum value
func (f Type) Name() naam.Name {
	switch f {
	case Type_Password:
		return naam.New("password")
	case Type_GithubToken:
		return naam.New("github_token")
	default:
		panics.P("Type.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f Type) Value() (driver.Value, error) {
	switch f {
	case Type_INVALID:
		return nil, nil
	case Type_Password:
		return "Password", nil
	case Type_GithubToken:
		return "GithubToken", nil

	default:
		return nil, fmt.Errorf("Cannot save enum Type to DB: '%d' is not a valid value for enum Type", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *Type) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewTypeFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "Type")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f Type) ImplementsGraphQLType(name string) bool {
	return name == "Auth_Secret_Type"
}

func (f *Type) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewTypeFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for Type: %T", input)
	}
	return err
}

func (f *Type) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum Type to a string: %w", err)
	}
	i := NewTypeFromString(enumStr)
	*f = i
	return nil
}

func (f Type) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil Type pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type TypeCondition struct {
	Op     filterlib.Operator
	Values []Type
}

func (c TypeCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c TypeCondition) Len() int {
	return len(c.Values)
}
func (c TypeCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// SecretFilter: <comments>
type SecretFilter struct {
	ID        *filterlib.IDCondition        `json:"id"`
	UserID    *filterlib.IDCondition        `json:"userID"`
	Type      *TypeCondition                `json:"type"`
	Secret    *filterlib.StringCondition    `json:"secret"`
	ExpiresAt *filterlib.TimestampCondition `json:"expiresAt"`
	CreatedAt *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt *filterlib.TimestampCondition `json:"deletedAt"`
	And       []SecretFilter                `json:"and"`
	Or        []SecretFilter                `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object Secret `json:"object"`
}

// Secret:
// SecretFieldCondition:
// SecretFilter:
// StandardEntityRequest:
// StandardEntityResponse:
// TypeCondition:
// AddRequest:
type AddRequest = dalutil.EntityAddRequest[SecretInput]

// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[Secret, SecretField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[Secret]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[Secret]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[SecretFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[Secret]

// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[Secret]
