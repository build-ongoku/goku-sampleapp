package svccore_entjobapplicant_meta

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

	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]{
		DbTableName: naam.New("tb_job_applicant"),
		TypeDALMeta: GetTypeJobApplicantDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeJobApplicantMeta struct {
	types.TypeCommonMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]
}

var _TypeJobApplicantMeta = TypeJobApplicantMeta{
	TypeCommonMeta: types.TypeCommonMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]{
		Name: naam.New("job_applicant"),
		Fields: []svccore_entjobapplicant_typ.JobApplicantField{
			svccore_entjobapplicant_typ.JobApplicantField_ID,
			svccore_entjobapplicant_typ.JobApplicantField_Name,
			svccore_entjobapplicant_typ.JobApplicantField_UserID,
			svccore_entjobapplicant_typ.JobApplicantField_ResumeID,
			svccore_entjobapplicant_typ.JobApplicantField_CreatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_DeletedAt,
		},
	},
}

func GetTypeJobApplicantMeta() *TypeJobApplicantMeta {
	return &_TypeJobApplicantMeta
}

func (meta TypeJobApplicantMeta) SetMetaFieldValues(ctx context.Context, obj svccore_entjobapplicant_typ.JobApplicant, now scalars.Timestamp) svccore_entjobapplicant_typ.JobApplicant {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity JobApplicants already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity JobApplicants already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeJobApplicantMeta) InternalHookSavePre(ctx context.Context, elem svccore_entjobapplicant_typ.JobApplicant, now scalars.Timestamp) (svccore_entjobapplicant_typ.JobApplicant, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity JobApplicants already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity JobApplicants already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeJobApplicantMeta) standardizeTimestamps(ctx context.Context, elem svccore_entjobapplicant_typ.JobApplicant) svccore_entjobapplicant_typ.JobApplicant {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeJobApplicantMeta) SetDefaultFieldValues(obj svccore_entjobapplicant_typ.JobApplicant) svccore_entjobapplicant_typ.JobApplicant {
	return obj
}

type TypeJobApplicantDALMeta struct {
	*TypeJobApplicantMeta
	dalutil.TypeCommonDALMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]
}

var _TypeJobApplicantDALMeta = TypeJobApplicantDALMeta{
	TypeJobApplicantMeta: &_TypeJobApplicantMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField]{
		DatabaseColumnFields: []svccore_entjobapplicant_typ.JobApplicantField{
			svccore_entjobapplicant_typ.JobApplicantField_ID,
			svccore_entjobapplicant_typ.JobApplicantField_Name,
			svccore_entjobapplicant_typ.JobApplicantField_UserID,
			svccore_entjobapplicant_typ.JobApplicantField_ResumeID,
			svccore_entjobapplicant_typ.JobApplicantField_CreatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_DeletedAt,
		},
		DatabaseSubTableFields: []svccore_entjobapplicant_typ.JobApplicantField{},
		SetInternallyByDALFields: []svccore_entjobapplicant_typ.JobApplicantField{
			svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_DeletedAt,
		},
		ImmutableFields: []svccore_entjobapplicant_typ.JobApplicantField{
			svccore_entjobapplicant_typ.JobApplicantField_ID,
			svccore_entjobapplicant_typ.JobApplicantField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svccore_entjobapplicant_typ.JobApplicantField{
			svccore_entjobapplicant_typ.JobApplicantField_CreatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt,
			svccore_entjobapplicant_typ.JobApplicantField_DeletedAt,
		},
		UpdatedAtField: svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt,
	},
}

func GetTypeJobApplicantDALMeta() *TypeJobApplicantDALMeta {
	return &_TypeJobApplicantDALMeta
}

func (meta *TypeJobApplicantDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeJobApplicantDALMeta) GetDirectDBValues(obj svccore_entjobapplicant_typ.JobApplicant) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Name,
		obj.UserID,
		obj.ResumeID,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeJobApplicantDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svccore_entjobapplicant_typ.JobApplicant) (svccore_entjobapplicant_typ.JobApplicant, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeJobApplicantDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svccore_entjobapplicant_typ.JobApplicant, error) {
	var elem svccore_entjobapplicant_typ.JobApplicant
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Name,
		&elem.UserID,
		&elem.ResumeID,
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

func (meta TypeJobApplicantDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svccore_entjobapplicant_typ.JobApplicant) ([]svccore_entjobapplicant_typ.JobApplicant, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeJobApplicantDALMeta) GetChangedFieldsAndValues(old, new svccore_entjobapplicant_typ.JobApplicant, allowedFields []svccore_entjobapplicant_typ.JobApplicantField) ([]svccore_entjobapplicant_typ.JobApplicantField, []interface{}) {

	var colsWithValueChange []svccore_entjobapplicant_typ.JobApplicantField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_UserID, allowedFields) {
		if old.UserID != new.UserID {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_UserID)
			vals = append(vals, new.UserID)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_ResumeID, allowedFields) {
		if old.ResumeID != new.ResumeID {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_ResumeID)
			vals = append(vals, new.ResumeID)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svccore_entjobapplicant_typ.JobApplicantField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svccore_entjobapplicant_typ.JobApplicantField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeJobApplicantDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField], allowedFields []svccore_entjobapplicant_typ.JobApplicantField, elem, oldElem svccore_entjobapplicant_typ.JobApplicant) (svccore_entjobapplicant_typ.JobApplicant, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "JobApplicant", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
