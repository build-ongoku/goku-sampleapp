package svccore_entfile_typ

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

// File: <comments>
type File struct {
	ID                      scalars.ID         `json:"id"`
	FileName                string             `json:"fileName" validate:"required"`
	StorageClient           StorageClient      `json:"storageClient" validate:"required"`
	StorageClientIdentifier string             `json:"storageClientIdentifier" validate:"required"`
	SizeBytes               int                `json:"sizeBytes" validate:"required"`
	MimeType                string             `json:"mimeType" validate:"required"`
	FileHash                string             `json:"fileHash" validate:"required"`
	CreatedAt               scalars.Timestamp  `json:"createdAt"`
	UpdatedAt               scalars.Timestamp  `json:"updatedAt"`
	DeletedAt               *scalars.Timestamp `json:"deletedAt"`
}

func (t File) GetID() scalars.ID {
	return t.ID
}
func (t File) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t File) SetUpdatedAt(tim scalars.Timestamp) {
	t.UpdatedAt = tim
}

// FileInput: <comments>
type FileInput struct {
	FileName                string        `json:"fileName" validate:"required"`
	StorageClient           StorageClient `json:"storageClient" validate:"required"`
	StorageClientIdentifier string        `json:"storageClientIdentifier" validate:"required"`
}

func NewFileWithMetaArrayFromInputs(ctx context.Context, ins []FileInput) []File {
	var outs []File
	for _, in := range ins {
		outs = append(outs, NewFileWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewFileWithMetaFromInput(ctx context.Context, in FileInput) File {
	return File{
		FileName:                in.FileName,
		StorageClient:           in.StorageClient,
		StorageClientIdentifier: in.StorageClientIdentifier,
	}
}

// FileField enum: <insert comment>
type FileField int

const (
	FileField_INVALID                 FileField = 0
	FileField_ID                      FileField = 1
	FileField_FileName                FileField = 2
	FileField_StorageClient           FileField = 3
	FileField_StorageClientIdentifier FileField = 4
	FileField_SizeBytes               FileField = 5
	FileField_MimeType                FileField = 6
	FileField_FileHash                FileField = 7
	FileField_CreatedAt               FileField = 8
	FileField_UpdatedAt               FileField = 9
	FileField_DeletedAt               FileField = 10
)

func NewFileFieldFromString(s string) FileField {
	switch s {
	case "INVALID":
		return FileField_INVALID
	case "ID":
		return FileField_ID
	case "FileName":
		return FileField_FileName
	case "StorageClient":
		return FileField_StorageClient
	case "StorageClientIdentifier":
		return FileField_StorageClientIdentifier
	case "SizeBytes":
		return FileField_SizeBytes
	case "MimeType":
		return FileField_MimeType
	case "FileHash":
		return FileField_FileHash
	case "CreatedAt":
		return FileField_CreatedAt
	case "UpdatedAt":
		return FileField_UpdatedAt
	case "DeletedAt":
		return FileField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "FileField"))
	}
}

// String implements the `fmt.Stringer` interface for FileField. It allows us to print the enum values as strings.
func (f FileField) String() string {
	switch f {
	case FileField_INVALID:
		return "INVALID"
	case FileField_ID:
		return "ID"
	case FileField_FileName:
		return "FileName"
	case FileField_StorageClient:
		return "StorageClient"
	case FileField_StorageClientIdentifier:
		return "StorageClientIdentifier"
	case FileField_SizeBytes:
		return "SizeBytes"
	case FileField_MimeType:
		return "MimeType"
	case FileField_FileHash:
		return "FileHash"
	case FileField_CreatedAt:
		return "CreatedAt"
	case FileField_UpdatedAt:
		return "UpdatedAt"
	case FileField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "FileField"))
	}
}

// Name gives a naam representation of the enum value
func (f FileField) Name() naam.Name {
	switch f {
	case FileField_ID:
		return naam.New("id")
	case FileField_FileName:
		return naam.New("file_name")
	case FileField_StorageClient:
		return naam.New("storage_client")
	case FileField_StorageClientIdentifier:
		return naam.New("storage_client_identifier")
	case FileField_SizeBytes:
		return naam.New("size_bytes")
	case FileField_MimeType:
		return naam.New("mime_type")
	case FileField_FileHash:
		return naam.New("file_hash")
	case FileField_CreatedAt:
		return naam.New("created_at")
	case FileField_UpdatedAt:
		return naam.New("updated_at")
	case FileField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("FileField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f FileField) Value() (driver.Value, error) {
	switch f {
	case FileField_INVALID:
		return nil, nil
	case FileField_ID:
		return "ID", nil
	case FileField_FileName:
		return "FileName", nil
	case FileField_StorageClient:
		return "StorageClient", nil
	case FileField_StorageClientIdentifier:
		return "StorageClientIdentifier", nil
	case FileField_SizeBytes:
		return "SizeBytes", nil
	case FileField_MimeType:
		return "MimeType", nil
	case FileField_FileHash:
		return "FileHash", nil
	case FileField_CreatedAt:
		return "CreatedAt", nil
	case FileField_UpdatedAt:
		return "UpdatedAt", nil
	case FileField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum FileField to DB: '%d' is not a valid value for enum FileField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *FileField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewFileFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "FileField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f FileField) ImplementsGraphQLType(name string) bool {
	return name == "Core_File_FileField"
}

func (f *FileField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewFileFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for FileField: %T", input)
	}
	return err
}

func (f *FileField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum FileField to a string: %w", err)
	}
	i := NewFileFieldFromString(enumStr)
	*f = i
	return nil
}

func (f FileField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil FileField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type FileFieldCondition struct {
	Op     filterlib.Operator
	Values []FileField
}

func (c FileFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c FileFieldCondition) Len() int {
	return len(c.Values)
}
func (c FileFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// StorageClient enum: <insert comment>
type StorageClient int

const (
	StorageClient_INVALID       StorageClient = 0
	StorageClient_LocalFile     StorageClient = 1
	StorageClient_Database      StorageClient = 2
	StorageClient_CloudAmazonS3 StorageClient = 3
)

func NewStorageClientFromString(s string) StorageClient {
	switch s {
	case "INVALID":
		return StorageClient_INVALID
	case "LocalFile":
		return StorageClient_LocalFile
	case "Database":
		return StorageClient_Database
	case "CloudAmazonS3":
		return StorageClient_CloudAmazonS3

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "StorageClient"))
	}
}

// String implements the `fmt.Stringer` interface for StorageClient. It allows us to print the enum values as strings.
func (f StorageClient) String() string {
	switch f {
	case StorageClient_INVALID:
		return "INVALID"
	case StorageClient_LocalFile:
		return "LocalFile"
	case StorageClient_Database:
		return "Database"
	case StorageClient_CloudAmazonS3:
		return "CloudAmazonS3"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "StorageClient"))
	}
}

// Name gives a naam representation of the enum value
func (f StorageClient) Name() naam.Name {
	switch f {
	case StorageClient_LocalFile:
		return naam.New("local_file")
	case StorageClient_Database:
		return naam.New("database")
	case StorageClient_CloudAmazonS3:
		return naam.New("cloud_amazon_s_3")
	default:
		panics.P("StorageClient.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f StorageClient) Value() (driver.Value, error) {
	switch f {
	case StorageClient_INVALID:
		return nil, nil
	case StorageClient_LocalFile:
		return "LocalFile", nil
	case StorageClient_Database:
		return "Database", nil
	case StorageClient_CloudAmazonS3:
		return "CloudAmazonS3", nil

	default:
		return nil, fmt.Errorf("Cannot save enum StorageClient to DB: '%d' is not a valid value for enum StorageClient", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *StorageClient) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewStorageClientFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "StorageClient")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f StorageClient) ImplementsGraphQLType(name string) bool {
	return name == "Core_File_StorageClient"
}

func (f *StorageClient) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewStorageClientFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for StorageClient: %T", input)
	}
	return err
}

func (f *StorageClient) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum StorageClient to a string: %w", err)
	}
	i := NewStorageClientFromString(enumStr)
	*f = i
	return nil
}

func (f StorageClient) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil StorageClient pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type StorageClientCondition struct {
	Op     filterlib.Operator
	Values []StorageClient
}

func (c StorageClientCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c StorageClientCondition) Len() int {
	return len(c.Values)
}
func (c StorageClientCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// FileFilter: <comments>
type FileFilter struct {
	ID                      *filterlib.IDCondition        `json:"id"`
	FileName                *filterlib.StringCondition    `json:"fileName"`
	StorageClient           *StorageClientCondition       `json:"storageClient"`
	StorageClientIdentifier *filterlib.StringCondition    `json:"storageClientIdentifier"`
	SizeBytes               *filterlib.NumberCondition    `json:"sizeBytes"`
	MimeType                *filterlib.StringCondition    `json:"mimeType"`
	FileHash                *filterlib.StringCondition    `json:"fileHash"`
	CreatedAt               *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt               *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt               *filterlib.TimestampCondition `json:"deletedAt"`
	And                     []FileFilter                  `json:"and"`
	Or                      []FileFilter                  `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object File `json:"object"`
}

// File:
// FileFieldCondition:
// FileFilter:
// StandardEntityRequest:
// StandardEntityResponse:
// StorageClientCondition:
// AddRequest:
type AddRequest = dalutil.EntityAddRequest[FileInput]

// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[File, FileField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[File]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[File]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[FileFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[File]

// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[File]
