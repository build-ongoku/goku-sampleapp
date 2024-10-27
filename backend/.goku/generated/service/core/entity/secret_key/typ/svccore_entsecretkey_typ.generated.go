package svccore_entsecretkey_typ

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

// SecretKey: <comments>
type SecretKey struct {
	ID               scalars.ID         `json:"id"`
	Type             Type               `json:"type" validate:"required"`
	PublicKey        string             `json:"publicKey" validate:"required"`
	PrivateKeyFormat Format             `json:"privateKeyFormat" validate:"required"`
	PublicKeyFormat  Format             `json:"publicKeyFormat" validate:"required"`
	CreatedAt        scalars.Timestamp  `json:"createdAt"`
	UpdatedAt        scalars.Timestamp  `json:"updatedAt"`
	DeletedAt        *scalars.Timestamp `json:"deletedAt"`
}

func (t SecretKey) GetID() scalars.ID {
	return t.ID
}
func (t SecretKey) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t SecretKey) SetUpdatedAt(tim scalars.Timestamp) SecretKey {
	t.UpdatedAt = tim
	return t
}

// SecretKeyInput: <comments>
type SecretKeyInput struct {
	Type Type `json:"type" validate:"required"`
}

func NewSecretKeyWithMetaArrayFromInputs(ctx context.Context, ins []SecretKeyInput) []SecretKey {
	var outs []SecretKey
	for _, in := range ins {
		outs = append(outs, NewSecretKeyWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewSecretKeyWithMetaFromInput(ctx context.Context, in SecretKeyInput) SecretKey {
	return SecretKey{
		Type: in.Type,
	}
}

// Format enum: <insert comment>
type Format int

const (
	Format_INVALID Format = 0
	Format_PEM     Format = 1
	Format_OpenSSH Format = 2
)

func NewFormatFromString(s string) Format {
	switch s {
	case "INVALID":
		return Format_INVALID
	case "PEM":
		return Format_PEM
	case "OpenSSH":
		return Format_OpenSSH

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "Format"))
	}
}

// String implements the `fmt.Stringer` interface for Format. It allows us to print the enum values as strings.
func (f Format) String() string {
	switch f {
	case Format_INVALID:
		return "INVALID"
	case Format_PEM:
		return "PEM"
	case Format_OpenSSH:
		return "OpenSSH"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "Format"))
	}
}

// Name gives a naam representation of the enum value
func (f Format) Name() naam.Name {
	switch f {
	case Format_PEM:
		return naam.New("pem")
	case Format_OpenSSH:
		return naam.New("open_ssh")
	default:
		panics.P("Format.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f Format) Value() (driver.Value, error) {
	switch f {
	case Format_INVALID:
		return nil, nil
	case Format_PEM:
		return "PEM", nil
	case Format_OpenSSH:
		return "OpenSSH", nil

	default:
		return nil, fmt.Errorf("Cannot save enum Format to DB: '%d' is not a valid value for enum Format", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *Format) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewFormatFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "Format")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f Format) ImplementsGraphQLType(name string) bool {
	return name == "Core_SecretKey_Format"
}

func (f *Format) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewFormatFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for Format: %T", input)
	}
	return err
}

func (f *Format) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum Format to a string: %w", err)
	}
	i := NewFormatFromString(enumStr)
	*f = i
	return nil
}

func (f Format) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil Format pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type FormatCondition struct {
	Op     filterlib.Operator
	Values []Format
}

func (c FormatCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c FormatCondition) Len() int {
	return len(c.Values)
}
func (c FormatCondition) GetValue(i int) interface{} {
	return c.Values[i]
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
	MethodName_HookInit      MethodName = 6
	MethodName_HookCreatePre MethodName = 7
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
	case "HookInit":
		return MethodName_HookInit
	case "HookCreatePre":
		return MethodName_HookCreatePre

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
	case MethodName_HookInit:
		return "HookInit"
	case MethodName_HookCreatePre:
		return "HookCreatePre"

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
	case MethodName_HookInit:
		return naam.New("hook_init")
	case MethodName_HookCreatePre:
		return naam.New("hook_create_pre")
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
	case MethodName_HookInit:
		return "HookInit", nil
	case MethodName_HookCreatePre:
		return "HookCreatePre", nil

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
	return name == "Core_SecretKey_MethodName"
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

// SecretKeyField enum: <insert comment>
type SecretKeyField int

const (
	SecretKeyField_INVALID          SecretKeyField = 0
	SecretKeyField_ID               SecretKeyField = 1
	SecretKeyField_Type             SecretKeyField = 2
	SecretKeyField_PublicKey        SecretKeyField = 3
	SecretKeyField_PrivateKeyFormat SecretKeyField = 4
	SecretKeyField_PublicKeyFormat  SecretKeyField = 5
	SecretKeyField_CreatedAt        SecretKeyField = 6
	SecretKeyField_UpdatedAt        SecretKeyField = 7
	SecretKeyField_DeletedAt        SecretKeyField = 8
)

func NewSecretKeyFieldFromString(s string) SecretKeyField {
	switch s {
	case "INVALID":
		return SecretKeyField_INVALID
	case "ID":
		return SecretKeyField_ID
	case "Type":
		return SecretKeyField_Type
	case "PublicKey":
		return SecretKeyField_PublicKey
	case "PrivateKeyFormat":
		return SecretKeyField_PrivateKeyFormat
	case "PublicKeyFormat":
		return SecretKeyField_PublicKeyFormat
	case "CreatedAt":
		return SecretKeyField_CreatedAt
	case "UpdatedAt":
		return SecretKeyField_UpdatedAt
	case "DeletedAt":
		return SecretKeyField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "SecretKeyField"))
	}
}

// String implements the `fmt.Stringer` interface for SecretKeyField. It allows us to print the enum values as strings.
func (f SecretKeyField) String() string {
	switch f {
	case SecretKeyField_INVALID:
		return "INVALID"
	case SecretKeyField_ID:
		return "ID"
	case SecretKeyField_Type:
		return "Type"
	case SecretKeyField_PublicKey:
		return "PublicKey"
	case SecretKeyField_PrivateKeyFormat:
		return "PrivateKeyFormat"
	case SecretKeyField_PublicKeyFormat:
		return "PublicKeyFormat"
	case SecretKeyField_CreatedAt:
		return "CreatedAt"
	case SecretKeyField_UpdatedAt:
		return "UpdatedAt"
	case SecretKeyField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "SecretKeyField"))
	}
}

// Name gives a naam representation of the enum value
func (f SecretKeyField) Name() naam.Name {
	switch f {
	case SecretKeyField_ID:
		return naam.New("id")
	case SecretKeyField_Type:
		return naam.New("type")
	case SecretKeyField_PublicKey:
		return naam.New("public_key")
	case SecretKeyField_PrivateKeyFormat:
		return naam.New("private_key_format")
	case SecretKeyField_PublicKeyFormat:
		return naam.New("public_key_format")
	case SecretKeyField_CreatedAt:
		return naam.New("created_at")
	case SecretKeyField_UpdatedAt:
		return naam.New("updated_at")
	case SecretKeyField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("SecretKeyField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f SecretKeyField) Value() (driver.Value, error) {
	switch f {
	case SecretKeyField_INVALID:
		return nil, nil
	case SecretKeyField_ID:
		return "ID", nil
	case SecretKeyField_Type:
		return "Type", nil
	case SecretKeyField_PublicKey:
		return "PublicKey", nil
	case SecretKeyField_PrivateKeyFormat:
		return "PrivateKeyFormat", nil
	case SecretKeyField_PublicKeyFormat:
		return "PublicKeyFormat", nil
	case SecretKeyField_CreatedAt:
		return "CreatedAt", nil
	case SecretKeyField_UpdatedAt:
		return "UpdatedAt", nil
	case SecretKeyField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum SecretKeyField to DB: '%d' is not a valid value for enum SecretKeyField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *SecretKeyField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewSecretKeyFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "SecretKeyField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f SecretKeyField) ImplementsGraphQLType(name string) bool {
	return name == "Core_SecretKey_SecretKeyField"
}

func (f *SecretKeyField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewSecretKeyFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for SecretKeyField: %T", input)
	}
	return err
}

func (f *SecretKeyField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum SecretKeyField to a string: %w", err)
	}
	i := NewSecretKeyFieldFromString(enumStr)
	*f = i
	return nil
}

func (f SecretKeyField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil SecretKeyField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type SecretKeyFieldCondition struct {
	Op     filterlib.Operator
	Values []SecretKeyField
}

func (c SecretKeyFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c SecretKeyFieldCondition) Len() int {
	return len(c.Values)
}
func (c SecretKeyFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// Type enum: <insert comment>
type Type int

const (
	Type_INVALID Type = 0
	Type_Ed25519 Type = 1
	Type_RSA     Type = 2
)

func NewTypeFromString(s string) Type {
	switch s {
	case "INVALID":
		return Type_INVALID
	case "Ed25519":
		return Type_Ed25519
	case "RSA":
		return Type_RSA

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "Type"))
	}
}

// String implements the `fmt.Stringer` interface for Type. It allows us to print the enum values as strings.
func (f Type) String() string {
	switch f {
	case Type_INVALID:
		return "INVALID"
	case Type_Ed25519:
		return "Ed25519"
	case Type_RSA:
		return "RSA"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "Type"))
	}
}

// Name gives a naam representation of the enum value
func (f Type) Name() naam.Name {
	switch f {
	case Type_Ed25519:
		return naam.New("ed_25519")
	case Type_RSA:
		return naam.New("rsa")
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
	case Type_Ed25519:
		return "Ed25519", nil
	case Type_RSA:
		return "RSA", nil

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
	return name == "Core_SecretKey_Type"
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

// SecretKeyFilter: <comments>
type SecretKeyFilter struct {
	ID               *filterlib.IDCondition        `json:"id"`
	Type             *TypeCondition                `json:"type"`
	PublicKey        *filterlib.StringCondition    `json:"publicKey"`
	PrivateKeyFormat *FormatCondition              `json:"privateKeyFormat"`
	PublicKeyFormat  *FormatCondition              `json:"publicKeyFormat"`
	CreatedAt        *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt        *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt        *filterlib.TimestampCondition `json:"deletedAt"`
	And              []SecretKeyFilter             `json:"and"`
	Or               []SecretKeyFilter             `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object SecretKey `json:"object"`
}

// AddRequest:
type AddRequest = dalutil.EntityAddRequest[SecretKeyInput]

// FormatCondition:
// GetRequest:
type GetRequest = dalutil.GetEntityRequest[SecretKey]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[SecretKeyFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[SecretKey]

// MethodNameCondition:
// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[SecretKey]

// SecretKey: Main type for entity SecretKey
// SecretKeyFieldCondition:
// SecretKeyFilter:
// StandardEntityRequest:
// StandardEntityResponse:
// TypeCondition:
// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[SecretKey, SecretKeyField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[SecretKey]
