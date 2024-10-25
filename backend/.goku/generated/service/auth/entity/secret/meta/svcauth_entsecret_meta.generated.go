package svcauth_entsecret_meta

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

	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]{
		DbTableName: naam.New("tb_secret"),
		TypeDALMeta: GetTypeSecretDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeSecretMeta struct {
	types.TypeCommonMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]
}

var _TypeSecretMeta = TypeSecretMeta{
	TypeCommonMeta: types.TypeCommonMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]{
		Name: naam.New("secret"),
		Fields: []svcauth_entsecret_typ.SecretField{
			svcauth_entsecret_typ.SecretField_ID,
			svcauth_entsecret_typ.SecretField_UserID,
			svcauth_entsecret_typ.SecretField_Type,
			svcauth_entsecret_typ.SecretField_Secret,
			svcauth_entsecret_typ.SecretField_ExpiresAt,
			svcauth_entsecret_typ.SecretField_CreatedAt,
			svcauth_entsecret_typ.SecretField_UpdatedAt,
			svcauth_entsecret_typ.SecretField_DeletedAt,
		},
	},
}

func GetTypeSecretMeta() *TypeSecretMeta {
	return &_TypeSecretMeta
}

func (meta TypeSecretMeta) SetMetaFieldValues(ctx context.Context, obj svcauth_entsecret_typ.Secret, now scalars.Timestamp) svcauth_entsecret_typ.Secret {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Secrets already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Secrets already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeSecretMeta) InternalHookSavePre(ctx context.Context, elem svcauth_entsecret_typ.Secret, now scalars.Timestamp) (svcauth_entsecret_typ.Secret, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Secrets already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Secrets already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeSecretMeta) standardizeTimestamps(ctx context.Context, elem svcauth_entsecret_typ.Secret) svcauth_entsecret_typ.Secret {
	// Standardize any timestamp fields.
	if elem.ExpiresAt != nil {
		elem.ExpiresAt.Standardize()
	}

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeSecretMeta) SetDefaultFieldValues(obj svcauth_entsecret_typ.Secret) svcauth_entsecret_typ.Secret {
	return obj
}

type TypeSecretDALMeta struct {
	*TypeSecretMeta
	dalutil.TypeCommonDALMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]
}

var _TypeSecretDALMeta = TypeSecretDALMeta{
	TypeSecretMeta: &_TypeSecretMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField]{
		DatabaseColumnFields: []svcauth_entsecret_typ.SecretField{
			svcauth_entsecret_typ.SecretField_ID,
			svcauth_entsecret_typ.SecretField_UserID,
			svcauth_entsecret_typ.SecretField_Type,
			svcauth_entsecret_typ.SecretField_Secret,
			svcauth_entsecret_typ.SecretField_ExpiresAt,
			svcauth_entsecret_typ.SecretField_CreatedAt,
			svcauth_entsecret_typ.SecretField_UpdatedAt,
			svcauth_entsecret_typ.SecretField_DeletedAt,
		},
		DatabaseSubTableFields: []svcauth_entsecret_typ.SecretField{},
		SetInternallyByDALFields: []svcauth_entsecret_typ.SecretField{
			svcauth_entsecret_typ.SecretField_UpdatedAt,
			svcauth_entsecret_typ.SecretField_DeletedAt,
		},
		ImmutableFields: []svcauth_entsecret_typ.SecretField{
			svcauth_entsecret_typ.SecretField_ID,
			svcauth_entsecret_typ.SecretField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svcauth_entsecret_typ.SecretField{
			svcauth_entsecret_typ.SecretField_ExpiresAt,
			svcauth_entsecret_typ.SecretField_CreatedAt,
			svcauth_entsecret_typ.SecretField_UpdatedAt,
			svcauth_entsecret_typ.SecretField_DeletedAt,
		},
		UpdatedAtField: svcauth_entsecret_typ.SecretField_UpdatedAt,
	},
}

func GetTypeSecretDALMeta() *TypeSecretDALMeta {
	return &_TypeSecretDALMeta
}

func (meta *TypeSecretDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeSecretDALMeta) GetDirectDBValues(obj svcauth_entsecret_typ.Secret) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.UserID,
		obj.Type,
		obj.Secret,
		obj.ExpiresAt,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeSecretDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svcauth_entsecret_typ.Secret) (svcauth_entsecret_typ.Secret, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeSecretDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svcauth_entsecret_typ.Secret, error) {
	var elem svcauth_entsecret_typ.Secret
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.UserID,
		&elem.Type,
		&elem.Secret,
		&elem.ExpiresAt,
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

func (meta TypeSecretDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svcauth_entsecret_typ.Secret) ([]svcauth_entsecret_typ.Secret, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeSecretDALMeta) GetChangedFieldsAndValues(old, new svcauth_entsecret_typ.Secret, allowedFields []svcauth_entsecret_typ.SecretField) ([]svcauth_entsecret_typ.SecretField, []interface{}) {

	var colsWithValueChange []svcauth_entsecret_typ.SecretField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_UserID, allowedFields) {
		if old.UserID != new.UserID {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_UserID)
			vals = append(vals, new.UserID)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_Type, allowedFields) {
		if old.Type != new.Type {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_Type)
			vals = append(vals, new.Type)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_Secret, allowedFields) {
		if old.Secret != new.Secret {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_Secret)
			vals = append(vals, new.Secret)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_ExpiresAt, allowedFields) {
		if old.ExpiresAt != new.ExpiresAt {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_ExpiresAt)
			vals = append(vals, new.ExpiresAt)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svcauth_entsecret_typ.SecretField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svcauth_entsecret_typ.SecretField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeSecretDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField], allowedFields []svcauth_entsecret_typ.SecretField, elem, oldElem svcauth_entsecret_typ.Secret) (svcauth_entsecret_typ.Secret, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "Secret", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
