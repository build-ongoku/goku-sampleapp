package svccore_enttask_meta

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

	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]{
		DbTableName: naam.New("tb_task"),
		TypeDALMeta: GetTypeTaskDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeTaskMeta struct {
	types.TypeCommonMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]
}

var _TypeTaskMeta = TypeTaskMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]{
		Name: naam.New("task"),
		Fields: []svccore_enttask_typ.TaskField{
			svccore_enttask_typ.TaskField_ID,
			svccore_enttask_typ.TaskField_Name,
			svccore_enttask_typ.TaskField_Description,
			svccore_enttask_typ.TaskField_Method,
			svccore_enttask_typ.TaskField_MethodRequestData,
			svccore_enttask_typ.TaskField_Enabled,
			svccore_enttask_typ.TaskField_CreatedAt,
			svccore_enttask_typ.TaskField_UpdatedAt,
			svccore_enttask_typ.TaskField_DeletedAt,
		},
	},
}

func GetTypeTaskMeta() *TypeTaskMeta {
	return &_TypeTaskMeta
}

func (meta TypeTaskMeta) SetMetaFieldValues(ctx context.Context, obj svccore_enttask_typ.Task, now scalars.Timestamp) svccore_enttask_typ.Task {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Tasks already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Tasks already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeTaskMeta) InternalHookSavePre(ctx context.Context, elem svccore_enttask_typ.Task, now scalars.Timestamp) (svccore_enttask_typ.Task, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Tasks already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Tasks already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeTaskMeta) standardizeTimestamps(ctx context.Context, elem svccore_enttask_typ.Task) svccore_enttask_typ.Task {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeTaskMeta) SetDefaultFieldValues(obj svccore_enttask_typ.Task) svccore_enttask_typ.Task {
	return obj
}

type TypeTaskDALMeta struct {
	*TypeTaskMeta
	dalutil.TypeCommonDALMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]
}

var _TypeTaskDALMeta = TypeTaskDALMeta{
	TypeTaskMeta: &_TypeTaskMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField]{
		DatabaseColumnFields: []svccore_enttask_typ.TaskField{
			svccore_enttask_typ.TaskField_ID,
			svccore_enttask_typ.TaskField_Name,
			svccore_enttask_typ.TaskField_Description,
			svccore_enttask_typ.TaskField_Method,
			svccore_enttask_typ.TaskField_MethodRequestData,
			svccore_enttask_typ.TaskField_Enabled,
			svccore_enttask_typ.TaskField_CreatedAt,
			svccore_enttask_typ.TaskField_UpdatedAt,
			svccore_enttask_typ.TaskField_DeletedAt,
		},
		DatabaseSubTableFields: []svccore_enttask_typ.TaskField{},
		SetInternallyByDALFields: []svccore_enttask_typ.TaskField{
			svccore_enttask_typ.TaskField_UpdatedAt,
			svccore_enttask_typ.TaskField_DeletedAt,
		},
		ImmutableFields: []svccore_enttask_typ.TaskField{
			svccore_enttask_typ.TaskField_ID,
			svccore_enttask_typ.TaskField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_enttask_typ.TaskField{
			svccore_enttask_typ.TaskField_CreatedAt,
			svccore_enttask_typ.TaskField_UpdatedAt,
			svccore_enttask_typ.TaskField_DeletedAt,
		},
		UpdatedAtField: svccore_enttask_typ.TaskField_UpdatedAt,
	},
}

func GetTypeTaskDALMeta() *TypeTaskDALMeta {
	return &_TypeTaskDALMeta
}

func (meta *TypeTaskDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeTaskDALMeta) GetDirectDBValues(obj svccore_enttask_typ.Task) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Name,
		obj.Description,
		obj.Method,
		obj.MethodRequestData,
		obj.Enabled,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeTaskDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeTaskDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_enttask_typ.Task, error) {
	var elem svccore_enttask_typ.Task
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Name,
		&elem.Description,
		&elem.Method,
		&elem.MethodRequestData,
		&elem.Enabled,
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

func (meta TypeTaskDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_enttask_typ.Task) ([]svccore_enttask_typ.Task, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeTaskDALMeta) GetChangedFieldsAndValues(old, new svccore_enttask_typ.Task, allowedFields []svccore_enttask_typ.TaskField) ([]svccore_enttask_typ.TaskField, []interface{}) {

	var colsWithValueChange []svccore_enttask_typ.TaskField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_Description, allowedFields) {
		if old.Description != new.Description {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_Description)
			vals = append(vals, new.Description)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_Method, allowedFields) {
		if old.Method != new.Method {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_Method)
			vals = append(vals, new.Method)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_MethodRequestData, allowedFields) {
		if old.MethodRequestData != new.MethodRequestData {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_MethodRequestData)
			vals = append(vals, new.MethodRequestData)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_Enabled, allowedFields) {
		if old.Enabled != new.Enabled {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_Enabled)
			vals = append(vals, new.Enabled)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttask_typ.TaskField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttask_typ.TaskField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeTaskDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField], allowedFields []svccore_enttask_typ.TaskField, elem, oldElem svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "Task", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
