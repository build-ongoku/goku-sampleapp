package svccore_enttaskrun_meta

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

	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]{
		DbTableName: naam.New("tb_task_run"),
		TypeDALMeta: GetTypeTaskRunDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeTaskRunMeta struct {
	types.TypeCommonMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]
}

var _TypeTaskRunMeta = TypeTaskRunMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]{
		Name: naam.New("task_run"),
		Fields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_ID,
			svccore_enttaskrun_typ.TaskRunField_StartedAt,
			svccore_enttaskrun_typ.TaskRunField_CompletedAt,
			svccore_enttaskrun_typ.TaskRunField_Status,
			svccore_enttaskrun_typ.TaskRunField_MethodRequestData,
			svccore_enttaskrun_typ.TaskRunField_MethodResponseData,
			svccore_enttaskrun_typ.TaskRunField_TriggeredBy,
			svccore_enttaskrun_typ.TaskRunField_CreatedAt,
			svccore_enttaskrun_typ.TaskRunField_UpdatedAt,
			svccore_enttaskrun_typ.TaskRunField_DeletedAt,
			svccore_enttaskrun_typ.TaskRunField_TaskID,
		},
	},
}

func GetTypeTaskRunMeta() *TypeTaskRunMeta {
	return &_TypeTaskRunMeta
}

func (meta TypeTaskRunMeta) SetMetaFieldValues(ctx context.Context, obj svccore_enttaskrun_typ.TaskRun, now scalars.Timestamp) svccore_enttaskrun_typ.TaskRun {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity TaskRuns already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity TaskRuns already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeTaskRunMeta) InternalHookSavePre(ctx context.Context, elem svccore_enttaskrun_typ.TaskRun, now scalars.Timestamp) (svccore_enttaskrun_typ.TaskRun, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity TaskRuns already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity TaskRuns already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeTaskRunMeta) standardizeTimestamps(ctx context.Context, elem svccore_enttaskrun_typ.TaskRun) svccore_enttaskrun_typ.TaskRun {
	// Standardize any timestamp fields.

	elem.StartedAt.Standardize()

	elem.CompletedAt.Standardize()

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeTaskRunMeta) SetDefaultFieldValues(obj svccore_enttaskrun_typ.TaskRun) svccore_enttaskrun_typ.TaskRun {
	return obj
}

type TypeTaskRunDALMeta struct {
	*TypeTaskRunMeta
	dalutil.TypeCommonDALMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]
}

var _TypeTaskRunDALMeta = TypeTaskRunDALMeta{
	TypeTaskRunMeta: &_TypeTaskRunMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField]{
		DatabaseColumnFields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_ID,
			svccore_enttaskrun_typ.TaskRunField_StartedAt,
			svccore_enttaskrun_typ.TaskRunField_CompletedAt,
			svccore_enttaskrun_typ.TaskRunField_Status,
			svccore_enttaskrun_typ.TaskRunField_MethodRequestData,
			svccore_enttaskrun_typ.TaskRunField_MethodResponseData,
			svccore_enttaskrun_typ.TaskRunField_TriggeredBy,
			svccore_enttaskrun_typ.TaskRunField_CreatedAt,
			svccore_enttaskrun_typ.TaskRunField_UpdatedAt,
			svccore_enttaskrun_typ.TaskRunField_DeletedAt,
			svccore_enttaskrun_typ.TaskRunField_TaskID,
		},
		DatabaseSubTableFields: []svccore_enttaskrun_typ.TaskRunField{},
		SetInternallyByDALFields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_UpdatedAt,
			svccore_enttaskrun_typ.TaskRunField_DeletedAt,
		},
		ImmutableFields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_ID,
			svccore_enttaskrun_typ.TaskRunField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_StartedAt,
			svccore_enttaskrun_typ.TaskRunField_CompletedAt,
			svccore_enttaskrun_typ.TaskRunField_CreatedAt,
			svccore_enttaskrun_typ.TaskRunField_UpdatedAt,
			svccore_enttaskrun_typ.TaskRunField_DeletedAt,
		},
		UpdatedAtField: svccore_enttaskrun_typ.TaskRunField_UpdatedAt,
	},
}

func GetTypeTaskRunDALMeta() *TypeTaskRunDALMeta {
	return &_TypeTaskRunDALMeta
}

func (meta *TypeTaskRunDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeTaskRunDALMeta) GetDirectDBValues(obj svccore_enttaskrun_typ.TaskRun) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.StartedAt,
		obj.CompletedAt,
		obj.Status,
		obj.MethodRequestData,
		obj.MethodResponseData,
		obj.TriggeredBy,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
		obj.TaskID,
	}
	return vals
}

func (meta TypeTaskRunDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_enttaskrun_typ.TaskRun) (svccore_enttaskrun_typ.TaskRun, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeTaskRunDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_enttaskrun_typ.TaskRun, error) {
	var elem svccore_enttaskrun_typ.TaskRun
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.StartedAt,
		&elem.CompletedAt,
		&elem.Status,
		&elem.MethodRequestData,
		&elem.MethodResponseData,
		&elem.TriggeredBy,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
		&elem.TaskID,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.standardizeTimestamps(ctx, elem)
	return elem, nil
}

func (meta TypeTaskRunDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_enttaskrun_typ.TaskRun) ([]svccore_enttaskrun_typ.TaskRun, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeTaskRunDALMeta) GetChangedFieldsAndValues(old, new svccore_enttaskrun_typ.TaskRun, allowedFields []svccore_enttaskrun_typ.TaskRunField) ([]svccore_enttaskrun_typ.TaskRunField, []interface{}) {

	var colsWithValueChange []svccore_enttaskrun_typ.TaskRunField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_StartedAt, allowedFields) {
		if old.StartedAt != new.StartedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_StartedAt)
			vals = append(vals, new.StartedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_CompletedAt, allowedFields) {
		if old.CompletedAt != new.CompletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_CompletedAt)
			vals = append(vals, new.CompletedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_Status, allowedFields) {
		if old.Status != new.Status {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_Status)
			vals = append(vals, new.Status)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_MethodRequestData, allowedFields) {
		if old.MethodRequestData != new.MethodRequestData {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_MethodRequestData)
			vals = append(vals, new.MethodRequestData)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_MethodResponseData, allowedFields) {
		if old.MethodResponseData != new.MethodResponseData {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_MethodResponseData)
			vals = append(vals, new.MethodResponseData)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_TriggeredBy, allowedFields) {
		if old.TriggeredBy != new.TriggeredBy {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_TriggeredBy)
			vals = append(vals, new.TriggeredBy)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}
	if types.IsFieldInFields(svccore_enttaskrun_typ.TaskRunField_TaskID, allowedFields) {
		if old.TaskID != new.TaskID {
			colsWithValueChange = append(colsWithValueChange, svccore_enttaskrun_typ.TaskRunField_TaskID)
			vals = append(vals, new.TaskID)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeTaskRunDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField], allowedFields []svccore_enttaskrun_typ.TaskRunField, elem, oldElem svccore_enttaskrun_typ.TaskRun) (svccore_enttaskrun_typ.TaskRun, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "TaskRun", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
