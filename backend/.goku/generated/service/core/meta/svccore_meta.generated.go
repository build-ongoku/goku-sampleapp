package svccore_meta

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

	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type TypeCronExpressionMeta struct {
	types.TypeCommonMeta[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField]
}

var _TypeCronExpressionMeta = TypeCronExpressionMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField]{
		Name: naam.New("cron_expression"),
		Fields: []svccore_typ.CronExpressionField{
			svccore_typ.CronExpressionField_ParentID,
			svccore_typ.CronExpressionField_ID,
			svccore_typ.CronExpressionField_Second,
			svccore_typ.CronExpressionField_Minute,
			svccore_typ.CronExpressionField_Hour,
			svccore_typ.CronExpressionField_DayOfMonth,
			svccore_typ.CronExpressionField_Month,
			svccore_typ.CronExpressionField_DayOfWeek,
			svccore_typ.CronExpressionField_CreatedAt,
			svccore_typ.CronExpressionField_UpdatedAt,
			svccore_typ.CronExpressionField_DeletedAt,
		},
	},
}

func GetTypeCronExpressionMeta() *TypeCronExpressionMeta {
	return &_TypeCronExpressionMeta
}

func (meta TypeCronExpressionMeta) SetMetaFieldValues(ctx context.Context, obj svccore_typ.CronExpressionWithMeta, now scalars.Timestamp) svccore_typ.CronExpressionWithMeta {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity CronExpressions already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity CronExpressions already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeCronExpressionMeta) InternalHookSavePre(ctx context.Context, elem svccore_typ.CronExpressionWithMeta, now scalars.Timestamp) (svccore_typ.CronExpressionWithMeta, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity CronExpressions already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity CronExpressions already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeCronExpressionMeta) standardizeTimestamps(ctx context.Context, elem svccore_typ.CronExpressionWithMeta) svccore_typ.CronExpressionWithMeta {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeCronExpressionMeta) SetDefaultFieldValues(obj svccore_typ.CronExpressionWithMeta) svccore_typ.CronExpressionWithMeta {
	return obj
}

type TypeCronExpressionDALMeta struct {
	*TypeCronExpressionMeta
	dalutil.TypeCommonDALMeta[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField]
}

var _TypeCronExpressionDALMeta = TypeCronExpressionDALMeta{
	TypeCronExpressionMeta: &_TypeCronExpressionMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField]{
		DatabaseColumnFields: []svccore_typ.CronExpressionField{
			svccore_typ.CronExpressionField_ParentID,
			svccore_typ.CronExpressionField_ID,
			svccore_typ.CronExpressionField_Second,
			svccore_typ.CronExpressionField_Minute,
			svccore_typ.CronExpressionField_Hour,
			svccore_typ.CronExpressionField_DayOfMonth,
			svccore_typ.CronExpressionField_Month,
			svccore_typ.CronExpressionField_DayOfWeek,
			svccore_typ.CronExpressionField_CreatedAt,
			svccore_typ.CronExpressionField_UpdatedAt,
			svccore_typ.CronExpressionField_DeletedAt,
		},
		DatabaseSubTableFields: []svccore_typ.CronExpressionField{},
		SetInternallyByDALFields: []svccore_typ.CronExpressionField{
			svccore_typ.CronExpressionField_UpdatedAt,
			svccore_typ.CronExpressionField_DeletedAt,
		},
		ImmutableFields: []svccore_typ.CronExpressionField{
			svccore_typ.CronExpressionField_ParentID,
			svccore_typ.CronExpressionField_ID,
			svccore_typ.CronExpressionField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_typ.CronExpressionField{
			svccore_typ.CronExpressionField_CreatedAt,
			svccore_typ.CronExpressionField_UpdatedAt,
			svccore_typ.CronExpressionField_DeletedAt,
		},
		UpdatedAtField: svccore_typ.CronExpressionField_UpdatedAt,
	},
}

func GetTypeCronExpressionDALMeta() *TypeCronExpressionDALMeta {
	return &_TypeCronExpressionDALMeta
}

func (meta *TypeCronExpressionDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeCronExpressionDALMeta) GetDirectDBValues(obj svccore_typ.CronExpressionWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.Second,
		obj.Minute,
		obj.Hour,
		obj.DayOfMonth,
		obj.Month,
		obj.DayOfWeek,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeCronExpressionDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_typ.CronExpressionWithMeta) (svccore_typ.CronExpressionWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeCronExpressionDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_typ.CronExpressionWithMeta, error) {
	var elem svccore_typ.CronExpressionWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.Second,
		&elem.Minute,
		&elem.Hour,
		&elem.DayOfMonth,
		&elem.Month,
		&elem.DayOfWeek,
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

func (meta TypeCronExpressionDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_typ.CronExpressionWithMeta) ([]svccore_typ.CronExpressionWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeCronExpressionDALMeta) GetChangedFieldsAndValues(old, new svccore_typ.CronExpressionWithMeta, allowedFields []svccore_typ.CronExpressionField) ([]svccore_typ.CronExpressionField, []interface{}) {

	var colsWithValueChange []svccore_typ.CronExpressionField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_typ.CronExpressionField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_Second, allowedFields) {
		if old.Second != new.Second {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_Second)
			vals = append(vals, new.Second)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_Minute, allowedFields) {
		if old.Minute != new.Minute {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_Minute)
			vals = append(vals, new.Minute)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_Hour, allowedFields) {
		if old.Hour != new.Hour {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_Hour)
			vals = append(vals, new.Hour)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_DayOfMonth, allowedFields) {
		if old.DayOfMonth != new.DayOfMonth {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_DayOfMonth)
			vals = append(vals, new.DayOfMonth)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_Month, allowedFields) {
		if old.Month != new.Month {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_Month)
			vals = append(vals, new.Month)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_DayOfWeek, allowedFields) {
		if old.DayOfWeek != new.DayOfWeek {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_DayOfWeek)
			vals = append(vals, new.DayOfWeek)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_typ.CronExpressionField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_typ.CronExpressionField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeCronExpressionDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_typ.CronExpressionWithMeta, svccore_typ.CronExpressionField], allowedFields []svccore_typ.CronExpressionField, elem, oldElem svccore_typ.CronExpressionWithMeta) (svccore_typ.CronExpressionWithMeta, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "CronExpression", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
