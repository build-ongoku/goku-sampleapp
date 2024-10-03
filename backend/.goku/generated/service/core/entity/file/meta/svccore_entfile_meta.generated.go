package svccore_entfile_meta

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/scalars"
	"github.com/teejays/gokutil/types"

	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]{
		DbTableName: naam.New("tb_file"),
		TypeDALMeta: GetTypeFileDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeFileMeta struct {
	types.TypeCommonMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]
}

var _TypeFileMeta = TypeFileMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]{
		Name: naam.New("file"),
		Fields: []svccore_entfile_typ.FileField{
			svccore_entfile_typ.FileField_ID,
			svccore_entfile_typ.FileField_FileName,
			svccore_entfile_typ.FileField_StorageClient,
			svccore_entfile_typ.FileField_StorageClientIdentifier,
			svccore_entfile_typ.FileField_SizeBytes,
			svccore_entfile_typ.FileField_MimeType,
			svccore_entfile_typ.FileField_FileHash,
			svccore_entfile_typ.FileField_CreatedAt,
			svccore_entfile_typ.FileField_UpdatedAt,
			svccore_entfile_typ.FileField_DeletedAt,
		},
	},
}

func GetTypeFileMeta() *TypeFileMeta {
	return &_TypeFileMeta
}

func (meta TypeFileMeta) SetMetaFieldValues(ctx context.Context, obj svccore_entfile_typ.File, now time.Time) svccore_entfile_typ.File {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Files already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity Files already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypeFileMeta) ConvertTimestampColumnsToUTC(obj svccore_entfile_typ.File) svccore_entfile_typ.File {
	return obj
}

func (meta TypeFileMeta) SetDefaultFieldValues(obj svccore_entfile_typ.File) svccore_entfile_typ.File {
	return obj
}

type TypeFileDALMeta struct {
	*TypeFileMeta
	dalutil.TypeCommonDALMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]
}

var _TypeFileDALMeta = TypeFileDALMeta{
	TypeFileMeta: &_TypeFileMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField]{
		DatabaseColumnFields: []svccore_entfile_typ.FileField{
			svccore_entfile_typ.FileField_ID,
			svccore_entfile_typ.FileField_FileName,
			svccore_entfile_typ.FileField_StorageClient,
			svccore_entfile_typ.FileField_StorageClientIdentifier,
			svccore_entfile_typ.FileField_SizeBytes,
			svccore_entfile_typ.FileField_MimeType,
			svccore_entfile_typ.FileField_FileHash,
			svccore_entfile_typ.FileField_CreatedAt,
			svccore_entfile_typ.FileField_UpdatedAt,
			svccore_entfile_typ.FileField_DeletedAt,
		},
		DatabaseSubTableFields: []svccore_entfile_typ.FileField{},
		SetInternallyByDALFields: []svccore_entfile_typ.FileField{
			svccore_entfile_typ.FileField_UpdatedAt,
			svccore_entfile_typ.FileField_DeletedAt,
		},
		ImmutableFields: []svccore_entfile_typ.FileField{
			svccore_entfile_typ.FileField_ID,
			svccore_entfile_typ.FileField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_entfile_typ.FileField{
			svccore_entfile_typ.FileField_CreatedAt,
			svccore_entfile_typ.FileField_UpdatedAt,
			svccore_entfile_typ.FileField_DeletedAt,
		},
		UpdatedAtField: svccore_entfile_typ.FileField_UpdatedAt,
	},
}

func GetTypeFileDALMeta() *TypeFileDALMeta {
	return &_TypeFileDALMeta
}

func (meta *TypeFileDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_entfile_typ.File, svccore_entfile_typ.FileField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeFileDALMeta) GetDirectDBValues(obj svccore_entfile_typ.File) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.FileName,
		obj.StorageClient,
		obj.StorageClientIdentifier,
		obj.SizeBytes,
		obj.MimeType,
		obj.FileHash,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeFileDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_entfile_typ.File) (svccore_entfile_typ.File, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeFileDALMeta) ScanDBNextRow(rows *sql.Rows) (svccore_entfile_typ.File, error) {
	var elem svccore_entfile_typ.File
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.FileName,
		&elem.StorageClient,
		&elem.StorageClientIdentifier,
		&elem.SizeBytes,
		&elem.MimeType,
		&elem.FileHash,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta TypeFileDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_entfile_typ.File) ([]svccore_entfile_typ.File, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeFileDALMeta) GetChangedFieldsAndValues(old, new svccore_entfile_typ.File, allowedFields []svccore_entfile_typ.FileField) ([]svccore_entfile_typ.FileField, []interface{}) {

	var colsWithValueChange []svccore_entfile_typ.FileField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_entfile_typ.FileField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_FileName, allowedFields) {
		if old.FileName != new.FileName {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_FileName)
			vals = append(vals, new.FileName)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_StorageClient, allowedFields) {
		if old.StorageClient != new.StorageClient {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_StorageClient)
			vals = append(vals, new.StorageClient)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_StorageClientIdentifier, allowedFields) {
		if old.StorageClientIdentifier != new.StorageClientIdentifier {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_StorageClientIdentifier)
			vals = append(vals, new.StorageClientIdentifier)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_SizeBytes, allowedFields) {
		if old.SizeBytes != new.SizeBytes {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_SizeBytes)
			vals = append(vals, new.SizeBytes)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_MimeType, allowedFields) {
		if old.MimeType != new.MimeType {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_MimeType)
			vals = append(vals, new.MimeType)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_FileHash, allowedFields) {
		if old.FileHash != new.FileHash {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_FileHash)
			vals = append(vals, new.FileHash)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entfile_typ.FileField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entfile_typ.FileField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeFileDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_entfile_typ.File, svccore_entfile_typ.FileField], allowedFields []svccore_entfile_typ.FileField, elem svccore_entfile_typ.File) (svccore_entfile_typ.File, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
