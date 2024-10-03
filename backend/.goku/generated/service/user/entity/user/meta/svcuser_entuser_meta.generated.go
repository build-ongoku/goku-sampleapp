package svcuser_entuser_meta

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/scalars"
	"github.com/teejays/gokutil/types"

	app_meta "sampleapp/backend/.goku/generated/meta"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type EntityDALMeta struct {
	dalutil.EntityCommonDALMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]
}

// _EntityDALMeta implements the singleton pattern.
var _EntityDALMeta = EntityDALMeta{
	dalutil.EntityCommonDALMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]{
		DbTableName: naam.New("tb_user"),
		TypeDALMeta: GetTypeUserDALMeta(),
	},
}

func GetEntityDALMeta() *EntityDALMeta {
	return &_EntityDALMeta
}

type TypeUserMeta struct {
	types.TypeCommonMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]
}

var _TypeUserMeta = TypeUserMeta{
	TypeCommonMeta: types.TypeCommonMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]{
		Name: naam.New("user"),
		Fields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_ID,
			svcuser_entuser_typ.UserField_Name,
			svcuser_entuser_typ.UserField_Name_ParentID,
			svcuser_entuser_typ.UserField_Name_ID,
			svcuser_entuser_typ.UserField_Name_FirstName,
			svcuser_entuser_typ.UserField_Name_MiddleInitial,
			svcuser_entuser_typ.UserField_Name_LastName,
			svcuser_entuser_typ.UserField_Name_CreatedAt,
			svcuser_entuser_typ.UserField_Name_UpdatedAt,
			svcuser_entuser_typ.UserField_Name_DeletedAt,
			svcuser_entuser_typ.UserField_Email,
			svcuser_entuser_typ.UserField_Addresses,
			svcuser_entuser_typ.UserField_Addresses_ParentID,
			svcuser_entuser_typ.UserField_Addresses_ID,
			svcuser_entuser_typ.UserField_Addresses_Line1,
			svcuser_entuser_typ.UserField_Addresses_Line2,
			svcuser_entuser_typ.UserField_Addresses_City,
			svcuser_entuser_typ.UserField_Addresses_State,
			svcuser_entuser_typ.UserField_Addresses_PostalCode,
			svcuser_entuser_typ.UserField_Addresses_Country,
			svcuser_entuser_typ.UserField_Addresses_CreatedAt,
			svcuser_entuser_typ.UserField_Addresses_UpdatedAt,
			svcuser_entuser_typ.UserField_Addresses_DeletedAt,
			svcuser_entuser_typ.UserField_TeamID,
			svcuser_entuser_typ.UserField_PastTeamIDs,
			svcuser_entuser_typ.UserField_CreatedAt,
			svcuser_entuser_typ.UserField_UpdatedAt,
			svcuser_entuser_typ.UserField_DeletedAt,
		},
	},
}

func GetTypeUserMeta() *TypeUserMeta {
	return &_TypeUserMeta
}

func (meta TypeUserMeta) SetMetaFieldValues(ctx context.Context, obj svcuser_entuser_typ.User, now time.Time) svcuser_entuser_typ.User {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Users already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity Users already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypeUserMeta) ConvertTimestampColumnsToUTC(obj svcuser_entuser_typ.User) svcuser_entuser_typ.User {
	return obj
}

func (meta TypeUserMeta) SetDefaultFieldValues(obj svcuser_entuser_typ.User) svcuser_entuser_typ.User {
	return obj
}

type TypeUserDALMeta struct {
	*TypeUserMeta
	dalutil.TypeCommonDALMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]
}

var _TypeUserDALMeta = TypeUserDALMeta{
	TypeUserMeta: &_TypeUserMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField]{
		DatabaseColumnFields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_ID,
			svcuser_entuser_typ.UserField_Email,
			svcuser_entuser_typ.UserField_TeamID,
			svcuser_entuser_typ.UserField_PastTeamIDs,
			svcuser_entuser_typ.UserField_CreatedAt,
			svcuser_entuser_typ.UserField_UpdatedAt,
			svcuser_entuser_typ.UserField_DeletedAt,
		},
		DatabaseSubTableFields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_Name,
			svcuser_entuser_typ.UserField_Addresses,
		},
		SetInternallyByDALFields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_Name_UpdatedAt,
			svcuser_entuser_typ.UserField_Name_DeletedAt,
			svcuser_entuser_typ.UserField_Addresses_UpdatedAt,
			svcuser_entuser_typ.UserField_Addresses_DeletedAt,
			svcuser_entuser_typ.UserField_UpdatedAt,
			svcuser_entuser_typ.UserField_DeletedAt,
		},
		ImmutableFields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_ID,
			svcuser_entuser_typ.UserField_Name_ParentID,
			svcuser_entuser_typ.UserField_Name_ID,
			svcuser_entuser_typ.UserField_Name_CreatedAt,
			svcuser_entuser_typ.UserField_Addresses_ParentID,
			svcuser_entuser_typ.UserField_Addresses_ID,
			svcuser_entuser_typ.UserField_Addresses_CreatedAt,
			svcuser_entuser_typ.UserField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []svcuser_entuser_typ.UserField{
			svcuser_entuser_typ.UserField_CreatedAt,
			svcuser_entuser_typ.UserField_UpdatedAt,
			svcuser_entuser_typ.UserField_DeletedAt,
		},
		UpdatedAtField: svcuser_entuser_typ.UserField_UpdatedAt,
	},
}

func GetTypeUserDALMeta() *TypeUserDALMeta {
	return &_TypeUserDALMeta
}

func (meta *TypeUserDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeUserDALMeta) GetDirectDBValues(obj svcuser_entuser_typ.User) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Email,
		obj.TeamID, pq.Array(obj.PastTeamIDs),
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeUserDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj svcuser_entuser_typ.User) (svcuser_entuser_typ.User, error) {

	// Insert 1:1 sub-tables
	// Insert Name
	{
		// Populate the ParentID for nested object, in case it is not populated
		if !obj.Name.ParentID.IsEmpty() && obj.Name.ParentID != obj.ID {
			return obj, fmt.Errorf("Adding User: nested object ParentID (%s) is different than the User ID (%s)", obj.Name.ParentID, obj.ID)
		}
		obj.Name.ParentID = obj.ID
		req := db.InsertTypeParams{
			TableName: params.TableName + "_" + svcuser_entuser_typ.UserField_Name.Name().FormatSQLTable(),
		}
		subMeta := app_meta.GetTypePersonNameDALMeta()
		subElems, err := dalutil.BatchAddType[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, req, subMeta, obj.Name)
		if err != nil {
			return obj, fmt.Errorf("inserting Name: %w", err)
		}
		obj.Name = subElems[0]
	}

	// Insert 1:Many sub-tables
	// Insert Addresses
	if len(obj.Addresses) > 0 {
		for j := range obj.Addresses {
			obj.Addresses[j].ParentID = obj.ID
		}
		req := db.InsertTypeParams{
			TableName: params.TableName + "_" + svcuser_entuser_typ.UserField_Addresses.Name().FormatSQLTable(),
		}
		subMeta := app_meta.GetTypeAddressDALMeta()
		subElems, err := dalutil.BatchAddType[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, req, subMeta, obj.Addresses...)
		if err != nil {
			return obj, fmt.Errorf("Inserting Addresses: %w", err)
		}
		obj.Addresses = subElems
	}

	return obj, nil
}

func (meta TypeUserDALMeta) ScanDBNextRow(rows *sql.Rows) (svcuser_entuser_typ.User, error) {
	var elem svcuser_entuser_typ.User
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Email,
		&elem.TeamID,
		pq.Array(&elem.PastTeamIDs),
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

func (meta TypeUserDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []svcuser_entuser_typ.User) ([]svcuser_entuser_typ.User, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields
	{
		// Fetch Name (type PersonNameWithMeta)
		subParams := db.ListTypeByIDsParams{
			TableName: params.TableName + "_" + "name",
			IDColumn:  "parent_id",
			IDs:       ids,
		}
		subMeta := app_meta.GetTypePersonNameDALMeta()
		subResp, err := dalutil.ListTypeByIDs[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, subParams, subMeta)
		if err != nil {
			return nil, fmt.Errorf("fetching PersonName: %w", err)
		}
		// assign the right PersonNameWithMeta to the right elem
		var subElemMap = make(map[scalars.ID]app_typ.PersonNameWithMeta)
		for _, subElem := range subResp.Items {
			if _, alreadyExists := subElemMap[subElem.ParentID]; alreadyExists {
				return nil, fmt.Errorf("fetching PersonName: got multiple elements for User.Name field but we expect only one element")
			}
			subElemMap[subElem.ParentID] = subElem
		}

		for i := range elems {
			subElem, exists := subElemMap[elems[i].ID]
			if !exists {
				return nil, fmt.Errorf("fetching PersonName: did not find any element for User.Name field for User with ID %s", elems[i].ID)
			}
			elems[i].Name = subElem
		}
	}

	// Step 2: Get the Nested (1:Many) fields
	{
		// Fetch Addresses (type AddressWithMeta)
		subParams := db.ListTypeByIDsParams{
			TableName: params.TableName + "_" + "addresses",
			IDColumn:  "parent_id",
			IDs:       ids,
		}
		subMeta := app_meta.GetTypeAddressDALMeta()
		subResp, err := dalutil.ListTypeByIDs[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, subParams, subMeta)
		if err != nil {
			return nil, fmt.Errorf("fetching Address: %w", err)
		}

		// assign the right Address to the right elem
		var subElemMap = make(map[scalars.ID][]app_typ.AddressWithMeta)
		for _, subElem := range subResp.Items {
			subElemMap[subElem.ParentID] = append(subElemMap[subElem.ParentID], subElem)
		}

		for i := range elems {
			subElems := subElemMap[elems[i].ID]
			elems[i].Addresses = subElems
		}
	}

	return elems, nil
}

func (meta TypeUserDALMeta) GetChangedFieldsAndValues(old, new svcuser_entuser_typ.User, allowedFields []svcuser_entuser_typ.UserField) ([]svcuser_entuser_typ.UserField, []interface{}) {

	var colsWithValueChange []svcuser_entuser_typ.UserField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_Email, allowedFields) {
		if old.Email != new.Email {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_Email)
			vals = append(vals, new.Email)
		}
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_TeamID, allowedFields) {
		if old.TeamID != new.TeamID {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_TeamID)
			vals = append(vals, new.TeamID)
		}
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_PastTeamIDs, allowedFields) {
		vals = append(vals, pq.Array(new.PastTeamIDs))
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(svcuser_entuser_typ.UserField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, svcuser_entuser_typ.UserField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeUserDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField], allowedFields []svcuser_entuser_typ.UserField, elem svcuser_entuser_typ.User) (svcuser_entuser_typ.User, error) {
	// Update Nested (1:1 & 1:Many)
	// Update Name
	{
		subElem := elem.Name
		// Populate subElem.ParentID in case it is missing
		if !subElem.ParentID.IsEmpty() && subElem.ParentID != elem.ID {
			return elem, fmt.Errorf("Updating User: Nested object Name's ParentID (%s) is different than the parent User's ID (%s)", subElem.ParentID, elem.ID)
		}
		subElem.ParentID = elem.ID
		// TODO: If the 1:1 subobject does not have an ID, that's okay since we don't always expect users to provide
		// IDs for subobjects, so fetch it.
		subElemTableName := req.TableName + "_" + svcuser_entuser_typ.UserField_Name.Name().FormatSQLTable()
		if subElem.ID.IsEmpty() {
			subElemsListParams := db.ListTypeByIDsParams{
				TableName: subElemTableName,
				IDColumn:  "parent_id",
				IDs:       []scalars.ID{elem.ID},
			}
			subElemsListResp, err := dalutil.ListTypeByIDs[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField](ctx, conn, subElemsListParams, meta)
			if err != nil {
				return elem, err
			}
			if len(subElemsListResp.Items) < 1 {
				return elem, fmt.Errorf("Updating User.Name: entity ID %s has no existing Name, don't know what to update", elem.ID)
			}
			if len(subElemsListResp.Items) > 1 {
				return elem, fmt.Errorf("Updating User.Name: entity ID %s has %d existing Name, we expected only one", elem.ID, len(subElemsListResp.Items))
			}
			subElem.ID = subElemsListResp.Items[0].ID
		}
		subReq := dalutil.UpdateTypeRequest[app_typ.PersonNameWithMeta, app_typ.PersonNameField]{
			TableName: subElemTableName,
			Object:    subElem,
			// Fields: <empty> : include all fields as we can't yet handle nested fields
			// ExcludeFields: <empty> : include all fields as we can't yet handle nested fields
		}
		subMeta := app_meta.GetTypePersonNameDALMeta()
		subResp, err := dalutil.UpdateType[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, subReq, subMeta)
		if err != nil {
			return elem, fmt.Errorf("Updating Name: %w", err)
		}
		elem.Name = subResp.Object

	}
	// Update Addresses
	{
		for i := range elem.Addresses {
			if elem.Addresses != nil {
				subElem := elem.Addresses[i]
				// Populate subElem.ParentID in case it is missing
				if !subElem.ParentID.IsEmpty() && subElem.ParentID != elem.ID {
					return elem, fmt.Errorf("Updating User: Nested object Addresses's ParentID (%s) is different than the parent User's ID (%s)", subElem.ParentID, elem.ID)
				}
				subElem.ParentID = elem.ID
				// TODO: If the 1:1 subobject does not have an ID, that's okay since we don't always expect users to provide
				// IDs for subobjects, so fetch it.
				subElemTableName := req.TableName + "_" + svcuser_entuser_typ.UserField_Addresses.Name().FormatSQLTable()
				subReq := dalutil.UpdateTypeRequest[app_typ.AddressWithMeta, app_typ.AddressField]{
					TableName: subElemTableName,
					Object:    subElem,
					// Fields: <empty> : include all fields as we can't yet handle nested fields
					// ExcludeFields: <empty> : include all fields as we can't yet handle nested fields
				}
				subMeta := app_meta.GetTypeAddressDALMeta()
				subResp, err := dalutil.UpdateType[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, subReq, subMeta)
				if err != nil {
					return elem, fmt.Errorf("Updating Addresses: %w", err)
				}
				elem.Addresses[i] = subResp.Object
			}
		}
	}

	return elem, nil
}
