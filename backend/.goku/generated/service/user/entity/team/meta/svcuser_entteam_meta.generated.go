package svcuser_entteam_meta

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

	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]{
		DbTableName: naam.New("tb_team"),
		TypeDALMeta: GetTypeTeamDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeTeamMeta struct {
	types.TypeCommonMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]
}

var _TypeTeamMeta = TypeTeamMeta{
	TypeCommonMeta: types.TypeCommonMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]{
		Name: naam.New("team"),
		Fields: []svcuser_entteam_typ.TeamField{
			svcuser_entteam_typ.TeamField_ID,
			svcuser_entteam_typ.TeamField_Name,
			svcuser_entteam_typ.TeamField_CreatedAt,
			svcuser_entteam_typ.TeamField_UpdatedAt,
			svcuser_entteam_typ.TeamField_DeletedAt,
		},
	},
}

func GetTypeTeamMeta() *TypeTeamMeta {
	return &_TypeTeamMeta
}

func (meta TypeTeamMeta) SetMetaFieldValues(ctx context.Context, obj svcuser_entteam_typ.Team, now scalars.Timestamp) svcuser_entteam_typ.Team {

	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Teams already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsEmpty() {
		obj.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Teams already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = now
	return obj
}

func (meta TypeTeamMeta) InternalHookSavePre(ctx context.Context, elem svcuser_entteam_typ.Team, now scalars.Timestamp) (svcuser_entteam_typ.Team, error) {
	// Set meta field values, if not already set

	// ID
	if elem.ID.IsEmpty() {
		elem.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Teams already has ID set", "value", elem.ID)
	}

	// CreatedAt
	if elem.CreatedAt.IsEmpty() {
		elem.CreatedAt = now
	} else {
		llog.Warn(ctx, "Entity Teams already has field CreatedAt set", "value", elem.CreatedAt)
	}

	// UpdatedAt
	elem.UpdatedAt = now

	// Standardize any timestamp fields.
	elem = meta.standardizeTimestamps(ctx, elem)

	return elem, nil
}

func (meta TypeTeamMeta) standardizeTimestamps(ctx context.Context, elem svcuser_entteam_typ.Team) svcuser_entteam_typ.Team {
	// Standardize any timestamp fields.

	elem.CreatedAt.Standardize()

	elem.UpdatedAt.Standardize()

	if elem.DeletedAt != nil {
		elem.DeletedAt.Standardize()
	}
	return elem
}

func (meta TypeTeamMeta) SetDefaultFieldValues(obj svcuser_entteam_typ.Team) svcuser_entteam_typ.Team {
	return obj
}

type TypeTeamDALMeta struct {
	*TypeTeamMeta
	dalutil.TypeCommonDALMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]
}

var _TypeTeamDALMeta = TypeTeamDALMeta{
	TypeTeamMeta: &_TypeTeamMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]{
		DatabaseColumnFields: []svcuser_entteam_typ.TeamField{
			svcuser_entteam_typ.TeamField_ID,
			svcuser_entteam_typ.TeamField_Name,
			svcuser_entteam_typ.TeamField_CreatedAt,
			svcuser_entteam_typ.TeamField_UpdatedAt,
			svcuser_entteam_typ.TeamField_DeletedAt,
		},
		DatabaseSubTableFields: []svcuser_entteam_typ.TeamField{},
		SetInternallyByDALFields: []svcuser_entteam_typ.TeamField{
			svcuser_entteam_typ.TeamField_UpdatedAt,
			svcuser_entteam_typ.TeamField_DeletedAt,
		},
		ImmutableFields: []svcuser_entteam_typ.TeamField{
			svcuser_entteam_typ.TeamField_ID,
			svcuser_entteam_typ.TeamField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svcuser_entteam_typ.TeamField{
			svcuser_entteam_typ.TeamField_CreatedAt,
			svcuser_entteam_typ.TeamField_UpdatedAt,
			svcuser_entteam_typ.TeamField_DeletedAt,
		},
		UpdatedAtField: svcuser_entteam_typ.TeamField_UpdatedAt,
	},
}

func GetTypeTeamDALMeta() *TypeTeamDALMeta {
	return &_TypeTeamDALMeta
}

func (meta *TypeTeamDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeTeamDALMeta) GetDirectDBValues(obj svcuser_entteam_typ.Team) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Name,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeTeamDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svcuser_entteam_typ.Team) (svcuser_entteam_typ.Team, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeTeamDALMeta) ScanDBNextRow(ctx context.Context, rows *sql.Rows) (svcuser_entteam_typ.Team, error) {
	var elem svcuser_entteam_typ.Team
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Name,
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

func (meta TypeTeamDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svcuser_entteam_typ.Team) ([]svcuser_entteam_typ.Team, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeTeamDALMeta) GetChangedFieldsAndValues(old, new svcuser_entteam_typ.Team, allowedFields []svcuser_entteam_typ.TeamField) ([]svcuser_entteam_typ.TeamField, []interface{}) {

	var colsWithValueChange []svcuser_entteam_typ.TeamField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svcuser_entteam_typ.TeamField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svcuser_entteam_typ.TeamField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svcuser_entteam_typ.TeamField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, svcuser_entteam_typ.TeamField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(svcuser_entteam_typ.TeamField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entteam_typ.TeamField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svcuser_entteam_typ.TeamField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entteam_typ.TeamField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svcuser_entteam_typ.TeamField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entteam_typ.TeamField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeTeamDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField], allowedFields []svcuser_entteam_typ.TeamField, elem, oldElem svcuser_entteam_typ.Team) (svcuser_entteam_typ.Team, error) {
	llog.Debug(ctx, "Updating sub-types", "parentType", "Team", "ID", elem.ID)

	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
