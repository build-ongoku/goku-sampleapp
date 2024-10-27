package svccore_entsecretkey_meta

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

	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]{
		DbTableName: naam.New("tb_secret_key"),
		TypeDALMeta: GetTypeSecretKeyDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeSecretKeyMeta struct {
	types.TypeCommonMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]
}

var _TypeSecretKeyMeta = TypeSecretKeyMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]{
		Name: naam.New("secret_key"),
		Fields: []svccore_entsecretkey_typ.SecretKeyField{
			svccore_entsecretkey_typ.SecretKeyField_ID,
			svccore_entsecretkey_typ.SecretKeyField_Type,
			svccore_entsecretkey_typ.SecretKeyField_PublicKey,
			svccore_entsecretkey_typ.SecretKeyField_PrivateKeyFormat,
			svccore_entsecretkey_typ.SecretKeyField_PublicKeyFormat,
			svccore_entsecretkey_typ.SecretKeyField_CreatedAt,
			svccore_entsecretkey_typ.SecretKeyField_UpdatedAt,
			svccore_entsecretkey_typ.SecretKeyField_DeletedAt,
		},
	},
}

func GetTypeSecretKeyMeta() *TypeSecretKeyMeta {
	return &_TypeSecretKeyMeta
}

func (meta TypeSecretKeyMeta) SetMetaFieldValues(ctx context.Context, obj svccore_entsecretkey_typ.SecretKey, now scalars.Timestamp) svccore_entsecretkey_typ.SecretKey {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity SecretKeies already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity SecretKeies already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeSecretKeyMeta) InternalHookSavePre(ctx context.Context, elem svccore_entsecretkey_typ.SecretKey, now scalars.Timestamp) (svccore_entsecretkey_typ.SecretKey, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity SecretKeies already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity SecretKeies already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeSecretKeyMeta) standardizeTimestamps(ctx context.Context, elem svccore_entsecretkey_typ.SecretKey) svccore_entsecretkey_typ.SecretKey {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeSecretKeyMeta) SetDefaultFieldValues(obj svccore_entsecretkey_typ.SecretKey) svccore_entsecretkey_typ.SecretKey {
	return obj
}

type TypeSecretKeyDALMeta struct {
	*TypeSecretKeyMeta
	dalutil.TypeCommonDALMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]
}

var _TypeSecretKeyDALMeta = TypeSecretKeyDALMeta{
	TypeSecretKeyMeta: &_TypeSecretKeyMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField]{
		DatabaseColumnFields: []svccore_entsecretkey_typ.SecretKeyField{
			svccore_entsecretkey_typ.SecretKeyField_ID,
			svccore_entsecretkey_typ.SecretKeyField_Type,
			svccore_entsecretkey_typ.SecretKeyField_PublicKey,
			svccore_entsecretkey_typ.SecretKeyField_PrivateKeyFormat,
			svccore_entsecretkey_typ.SecretKeyField_PublicKeyFormat,
			svccore_entsecretkey_typ.SecretKeyField_CreatedAt,
			svccore_entsecretkey_typ.SecretKeyField_UpdatedAt,
			svccore_entsecretkey_typ.SecretKeyField_DeletedAt,
		},
		DatabaseSubTableFields: []svccore_entsecretkey_typ.SecretKeyField{},
		SetInternallyByDALFields: []svccore_entsecretkey_typ.SecretKeyField{
			svccore_entsecretkey_typ.SecretKeyField_UpdatedAt,
			svccore_entsecretkey_typ.SecretKeyField_DeletedAt,
		},
		ImmutableFields: []svccore_entsecretkey_typ.SecretKeyField{
			svccore_entsecretkey_typ.SecretKeyField_ID,
			svccore_entsecretkey_typ.SecretKeyField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_entsecretkey_typ.SecretKeyField{
			svccore_entsecretkey_typ.SecretKeyField_CreatedAt,
			svccore_entsecretkey_typ.SecretKeyField_UpdatedAt,
			svccore_entsecretkey_typ.SecretKeyField_DeletedAt,
		},
		UpdatedAtField: svccore_entsecretkey_typ.SecretKeyField_UpdatedAt,
	},
}

func GetTypeSecretKeyDALMeta() *TypeSecretKeyDALMeta {
	return &_TypeSecretKeyDALMeta
}

func (meta *TypeSecretKeyDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeSecretKeyDALMeta) GetDirectDBValues(obj svccore_entsecretkey_typ.SecretKey) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Type,
		obj.PublicKey,
		obj.PrivateKeyFormat,
		obj.PublicKeyFormat,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeSecretKeyDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_entsecretkey_typ.SecretKey) (svccore_entsecretkey_typ.SecretKey, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeSecretKeyDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_entsecretkey_typ.SecretKey, error) {
	var elem svccore_entsecretkey_typ.SecretKey
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Type,
		&elem.PublicKey,
		&elem.PrivateKeyFormat,
		&elem.PublicKeyFormat,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.standardizeTimestamps(ctx, elem)
	return elem, nil
}

func (meta TypeSecretKeyDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_entsecretkey_typ.SecretKey) ([]svccore_entsecretkey_typ.SecretKey, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeSecretKeyDALMeta) GetChangedFieldsAndValues(old, new svccore_entsecretkey_typ.SecretKey, allowedFields []svccore_entsecretkey_typ.SecretKeyField) ([]svccore_entsecretkey_typ.SecretKeyField, []interface{}) {

	var colsWithValueChange []svccore_entsecretkey_typ.SecretKeyField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_Type, allowedFields) {
		if old.Type != new.Type {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_Type)
			vals = append(vals, new.Type)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_PublicKey, allowedFields) {
		if old.PublicKey != new.PublicKey {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_PublicKey)
			vals = append(vals, new.PublicKey)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_PrivateKeyFormat, allowedFields) {
		if old.PrivateKeyFormat != new.PrivateKeyFormat {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_PrivateKeyFormat)
			vals = append(vals, new.PrivateKeyFormat)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_PublicKeyFormat, allowedFields) {
		if old.PublicKeyFormat != new.PublicKeyFormat {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_PublicKeyFormat)
			vals = append(vals, new.PublicKeyFormat)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entsecretkey_typ.SecretKeyField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entsecretkey_typ.SecretKeyField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeSecretKeyDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField], allowedFields []svccore_entsecretkey_typ.SecretKeyField, elem, oldElem svccore_entsecretkey_typ.SecretKey) (svccore_entsecretkey_typ.SecretKey, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "SecretKey", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
