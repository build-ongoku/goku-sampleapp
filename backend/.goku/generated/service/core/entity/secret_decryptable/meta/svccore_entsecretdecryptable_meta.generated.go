package svccore_entsecretdecryptable_meta

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/scalars"
	"github.com/teejays/gokutil/types"

	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]{
		DbTableName: naam.New("tb_secret_decryptable"),
		TypeDALMeta: GetTypeSecretDecryptableDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeSecretDecryptableMeta struct {
	types.TypeCommonMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]
}

var _TypeSecretDecryptableMeta = TypeSecretDecryptableMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]{
		Name: naam.New("secret_decryptable"),
		Fields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{
			svccore_entsecretdecryptable_typ.SecretDecryptableField_ID,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_RawValue,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_EncryptedValue,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_Salt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_SecretKeyID,
		},
	},
}

func GetTypeSecretDecryptableMeta() *TypeSecretDecryptableMeta {
	return &_TypeSecretDecryptableMeta
}

func (meta TypeSecretDecryptableMeta) SetMetaFieldValues(ctx context.Context, obj svccore_entsecretdecryptable_typ.SecretDecryptable, now scalars.Timestamp) svccore_entsecretdecryptable_typ.SecretDecryptable {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity SecretDecryptables already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity SecretDecryptables already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeSecretDecryptableMeta) InternalHookSavePre(ctx context.Context, elem svccore_entsecretdecryptable_typ.SecretDecryptable, now scalars.Timestamp) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity SecretDecryptables already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity SecretDecryptables already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeSecretDecryptableMeta) standardizeTimestamps(ctx context.Context, elem svccore_entsecretdecryptable_typ.SecretDecryptable) svccore_entsecretdecryptable_typ.SecretDecryptable {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeSecretDecryptableMeta) SetDefaultFieldValues(obj svccore_entsecretdecryptable_typ.SecretDecryptable) svccore_entsecretdecryptable_typ.SecretDecryptable {
	return obj
}

type TypeSecretDecryptableDALMeta struct {
	*TypeSecretDecryptableMeta
	dalutil.TypeCommonDALMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]
}

var _TypeSecretDecryptableDALMeta = TypeSecretDecryptableDALMeta{
	TypeSecretDecryptableMeta: &_TypeSecretDecryptableMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField]{
		DatabaseColumnFields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{
			svccore_entsecretdecryptable_typ.SecretDecryptableField_ID,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_EncryptedValue,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_Salt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_SecretKeyID,
		},
		DatabaseSubTableFields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{},
		SetInternallyByDALFields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{
			svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt,
		},
		ImmutableFields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{
			svccore_entsecretdecryptable_typ.SecretDecryptableField_ID,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_entsecretdecryptable_typ.SecretDecryptableField{
			svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt,
			svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt,
		},
		UpdatedAtField: svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt,
	},
}

func GetTypeSecretDecryptableDALMeta() *TypeSecretDecryptableDALMeta {
	return &_TypeSecretDecryptableDALMeta
}

func (meta *TypeSecretDecryptableDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeSecretDecryptableDALMeta) GetDirectDBValues(obj svccore_entsecretdecryptable_typ.SecretDecryptable) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.EncryptedValue,
		obj.Salt,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
		obj.SecretKeyID,
	}
	return vals
}

func (meta TypeSecretDecryptableDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeSecretDecryptableDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	var elem svccore_entsecretdecryptable_typ.SecretDecryptable
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.EncryptedValue,
		&elem.Salt,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
		&elem.SecretKeyID,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.standardizeTimestamps(ctx, elem)
	return elem, nil
}

func (meta TypeSecretDecryptableDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_entsecretdecryptable_typ.SecretDecryptable) ([]svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeSecretDecryptableDALMeta) GetChangedFieldsAndValues(old, new svccore_entsecretdecryptable_typ.SecretDecryptable, allowedFields []svccore_entsecretdecryptable_typ.SecretDecryptableField) ([]svccore_entsecretdecryptable_typ.SecretDecryptableField, []interface{}) {

	var colsWithValueChange []svccore_entsecretdecryptable_typ.SecretDecryptableField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_EncryptedValue, allowedFields) {
		if old.EncryptedValue != new.EncryptedValue {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_EncryptedValue)
			vals = append(vals, new.EncryptedValue)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_Salt, allowedFields) {
		if old.Salt != new.Salt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_Salt)
			vals = append(vals, new.Salt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretdecryptable_typ.SecretDecryptableField_SecretKeyID, allowedFields) {
		if old.SecretKeyID != new.SecretKeyID {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretdecryptable_typ.SecretDecryptableField_SecretKeyID)
			vals = append(vals, new.SecretKeyID)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeSecretDecryptableDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField], allowedFields []svccore_entsecretdecryptable_typ.SecretDecryptableField, elem, oldElem svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "SecretDecryptable", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
