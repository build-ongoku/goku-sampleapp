package svccore_entsecretdecryptable_typ

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

// SecretDecryptable: <comments>
type SecretDecryptable struct {
	ID             scalars.ID         `json:"id"`
	RawValue       string             `json:"rawValue"`
	EncryptedValue string             `json:"encryptedValue" validate:"required"`
	Salt           string             `json:"salt" validate:"required"`
	CreatedAt      scalars.Timestamp  `json:"createdAt"`
	UpdatedAt      scalars.Timestamp  `json:"updatedAt"`
	DeletedAt      *scalars.Timestamp `json:"deletedAt"`
	SecretKeyID    scalars.ID         `json:"secretKeyID"`
}

func (t SecretDecryptable) GetID() scalars.ID {
	return t.ID
}
func (t SecretDecryptable) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t SecretDecryptable) SetUpdatedAt(tim scalars.Timestamp) SecretDecryptable {
	t.UpdatedAt = tim
	return t
}

// SecretDecryptableInput: <comments>
type SecretDecryptableInput struct {
	RawValue    string     `json:"rawValue"`
	SecretKeyID scalars.ID `json:"secretKeyID"`
}

func NewSecretDecryptableWithMetaArrayFromInputs(ctx context.Context, ins []SecretDecryptableInput) []SecretDecryptable {
	var outs []SecretDecryptable
	for _, in := range ins {
		outs = append(outs, NewSecretDecryptableWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewSecretDecryptableWithMetaFromInput(ctx context.Context, in SecretDecryptableInput) SecretDecryptable {
	return SecretDecryptable{
		RawValue:    in.RawValue,
		SecretKeyID: in.SecretKeyID,
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
	MethodName_ActionDecrypt MethodName = 7
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
	case "ActionDecrypt":
		return MethodName_ActionDecrypt

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
	case MethodName_ActionDecrypt:
		return "ActionDecrypt"

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
	case MethodName_ActionDecrypt:
		return naam.New("action_decrypt")
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
	case MethodName_ActionDecrypt:
		return "ActionDecrypt", nil

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
	return name == "Core_SecretDecryptable_MethodName"
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

// SecretDecryptableAction enum: <insert comment>
type SecretDecryptableAction int

const (
	SecretDecryptableAction_INVALID SecretDecryptableAction = 0
	SecretDecryptableAction_Decrypt SecretDecryptableAction = 1
)

func NewSecretDecryptableActionFromString(s string) SecretDecryptableAction {
	switch s {
	case "INVALID":
		return SecretDecryptableAction_INVALID
	case "Decrypt":
		return SecretDecryptableAction_Decrypt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "SecretDecryptableAction"))
	}
}

// String implements the `fmt.Stringer` interface for SecretDecryptableAction. It allows us to print the enum values as strings.
func (f SecretDecryptableAction) String() string {
	switch f {
	case SecretDecryptableAction_INVALID:
		return "INVALID"
	case SecretDecryptableAction_Decrypt:
		return "Decrypt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "SecretDecryptableAction"))
	}
}

// Name gives a naam representation of the enum value
func (f SecretDecryptableAction) Name() naam.Name {
	switch f {
	case SecretDecryptableAction_Decrypt:
		return naam.New("decrypt")
	default:
		panics.P("SecretDecryptableAction.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f SecretDecryptableAction) Value() (driver.Value, error) {
	switch f {
	case SecretDecryptableAction_INVALID:
		return nil, nil
	case SecretDecryptableAction_Decrypt:
		return "Decrypt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum SecretDecryptableAction to DB: '%d' is not a valid value for enum SecretDecryptableAction", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *SecretDecryptableAction) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewSecretDecryptableActionFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "SecretDecryptableAction")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f SecretDecryptableAction) ImplementsGraphQLType(name string) bool {
	return name == "Core_SecretDecryptable_SecretDecryptableAction"
}

func (f *SecretDecryptableAction) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewSecretDecryptableActionFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for SecretDecryptableAction: %T", input)
	}
	return err
}

func (f *SecretDecryptableAction) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum SecretDecryptableAction to a string: %w", err)
	}
	i := NewSecretDecryptableActionFromString(enumStr)
	*f = i
	return nil
}

func (f SecretDecryptableAction) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil SecretDecryptableAction pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type SecretDecryptableActionCondition struct {
	Op     filterlib.Operator
	Values []SecretDecryptableAction
}

func (c SecretDecryptableActionCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c SecretDecryptableActionCondition) Len() int {
	return len(c.Values)
}
func (c SecretDecryptableActionCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// SecretDecryptableField enum: <insert comment>
type SecretDecryptableField int

const (
	SecretDecryptableField_INVALID        SecretDecryptableField = 0
	SecretDecryptableField_ID             SecretDecryptableField = 1
	SecretDecryptableField_RawValue       SecretDecryptableField = 2
	SecretDecryptableField_EncryptedValue SecretDecryptableField = 3
	SecretDecryptableField_Salt           SecretDecryptableField = 4
	SecretDecryptableField_CreatedAt      SecretDecryptableField = 5
	SecretDecryptableField_UpdatedAt      SecretDecryptableField = 6
	SecretDecryptableField_DeletedAt      SecretDecryptableField = 7
	SecretDecryptableField_SecretKeyID    SecretDecryptableField = 8
)

func NewSecretDecryptableFieldFromString(s string) SecretDecryptableField {
	switch s {
	case "INVALID":
		return SecretDecryptableField_INVALID
	case "ID":
		return SecretDecryptableField_ID
	case "RawValue":
		return SecretDecryptableField_RawValue
	case "EncryptedValue":
		return SecretDecryptableField_EncryptedValue
	case "Salt":
		return SecretDecryptableField_Salt
	case "CreatedAt":
		return SecretDecryptableField_CreatedAt
	case "UpdatedAt":
		return SecretDecryptableField_UpdatedAt
	case "DeletedAt":
		return SecretDecryptableField_DeletedAt
	case "SecretKeyID":
		return SecretDecryptableField_SecretKeyID

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "SecretDecryptableField"))
	}
}

// String implements the `fmt.Stringer` interface for SecretDecryptableField. It allows us to print the enum values as strings.
func (f SecretDecryptableField) String() string {
	switch f {
	case SecretDecryptableField_INVALID:
		return "INVALID"
	case SecretDecryptableField_ID:
		return "ID"
	case SecretDecryptableField_RawValue:
		return "RawValue"
	case SecretDecryptableField_EncryptedValue:
		return "EncryptedValue"
	case SecretDecryptableField_Salt:
		return "Salt"
	case SecretDecryptableField_CreatedAt:
		return "CreatedAt"
	case SecretDecryptableField_UpdatedAt:
		return "UpdatedAt"
	case SecretDecryptableField_DeletedAt:
		return "DeletedAt"
	case SecretDecryptableField_SecretKeyID:
		return "SecretKeyID"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "SecretDecryptableField"))
	}
}

// Name gives a naam representation of the enum value
func (f SecretDecryptableField) Name() naam.Name {
	switch f {
	case SecretDecryptableField_ID:
		return naam.New("id")
	case SecretDecryptableField_RawValue:
		return naam.New("raw_value")
	case SecretDecryptableField_EncryptedValue:
		return naam.New("encrypted_value")
	case SecretDecryptableField_Salt:
		return naam.New("salt")
	case SecretDecryptableField_CreatedAt:
		return naam.New("created_at")
	case SecretDecryptableField_UpdatedAt:
		return naam.New("updated_at")
	case SecretDecryptableField_DeletedAt:
		return naam.New("deleted_at")
	case SecretDecryptableField_SecretKeyID:
		return naam.New("secret_key_id")
	default:
		panics.P("SecretDecryptableField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f SecretDecryptableField) Value() (driver.Value, error) {
	switch f {
	case SecretDecryptableField_INVALID:
		return nil, nil
	case SecretDecryptableField_ID:
		return "ID", nil
	case SecretDecryptableField_RawValue:
		return "RawValue", nil
	case SecretDecryptableField_EncryptedValue:
		return "EncryptedValue", nil
	case SecretDecryptableField_Salt:
		return "Salt", nil
	case SecretDecryptableField_CreatedAt:
		return "CreatedAt", nil
	case SecretDecryptableField_UpdatedAt:
		return "UpdatedAt", nil
	case SecretDecryptableField_DeletedAt:
		return "DeletedAt", nil
	case SecretDecryptableField_SecretKeyID:
		return "SecretKeyID", nil

	default:
		return nil, fmt.Errorf("Cannot save enum SecretDecryptableField to DB: '%d' is not a valid value for enum SecretDecryptableField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *SecretDecryptableField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewSecretDecryptableFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "SecretDecryptableField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f SecretDecryptableField) ImplementsGraphQLType(name string) bool {
	return name == "Core_SecretDecryptable_SecretDecryptableField"
}

func (f *SecretDecryptableField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewSecretDecryptableFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for SecretDecryptableField: %T", input)
	}
	return err
}

func (f *SecretDecryptableField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum SecretDecryptableField to a string: %w", err)
	}
	i := NewSecretDecryptableFieldFromString(enumStr)
	*f = i
	return nil
}

func (f SecretDecryptableField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil SecretDecryptableField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type SecretDecryptableFieldCondition struct {
	Op     filterlib.Operator
	Values []SecretDecryptableField
}

func (c SecretDecryptableFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c SecretDecryptableFieldCondition) Len() int {
	return len(c.Values)
}
func (c SecretDecryptableFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// SecretDecryptableFilter: <comments>
type SecretDecryptableFilter struct {
	ID             *filterlib.IDCondition        `json:"id"`
	RawValue       *filterlib.StringCondition    `json:"rawValue"`
	EncryptedValue *filterlib.StringCondition    `json:"encryptedValue"`
	Salt           *filterlib.StringCondition    `json:"salt"`
	CreatedAt      *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt      *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt      *filterlib.TimestampCondition `json:"deletedAt"`
	SecretKeyID    *filterlib.IDCondition        `json:"secretKeyID"`
	And            []SecretDecryptableFilter     `json:"and"`
	Or             []SecretDecryptableFilter     `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object SecretDecryptable `json:"object"`
}

// AddRequest:
type AddRequest = dalutil.EntityAddRequest[SecretDecryptableInput]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[SecretDecryptable]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[SecretDecryptableFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[SecretDecryptable]

// MethodNameCondition:
// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[SecretDecryptable]

// SecretDecryptable: Main type for entity SecretDecryptable
// SecretDecryptableActionCondition:
// SecretDecryptableFieldCondition:
// SecretDecryptableFilter:
// StandardEntityRequest:
// StandardEntityResponse:
// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[SecretDecryptable, SecretDecryptableField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[SecretDecryptable]
