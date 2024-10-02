package app_meta

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

	app_typ "sampleapp/backend/.goku/generated/typ"
)

var llog = log.GetLogger().WithHeading("DAL Meta")

type TypeAddressMeta struct {
	types.TypeCommonMeta[app_typ.AddressWithMeta, app_typ.AddressField]
}

var _TypeAddressMeta = TypeAddressMeta{
	TypeCommonMeta: types.TypeCommonMeta[app_typ.AddressWithMeta, app_typ.AddressField]{
		Name: naam.New("address"),
		Fields: []app_typ.AddressField{
			app_typ.AddressField_ParentID,
			app_typ.AddressField_ID,
			app_typ.AddressField_Line1,
			app_typ.AddressField_Line2,
			app_typ.AddressField_City,
			app_typ.AddressField_State,
			app_typ.AddressField_PostalCode,
			app_typ.AddressField_Country,
			app_typ.AddressField_CreatedAt,
			app_typ.AddressField_UpdatedAt,
			app_typ.AddressField_DeletedAt,
		},
	},
}

func GetTypeAddressMeta() *TypeAddressMeta {
	return &_TypeAddressMeta
}

func (meta TypeAddressMeta) SetMetaFieldValues(ctx context.Context, obj app_typ.AddressWithMeta, now time.Time) app_typ.AddressWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Addresses already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity Addresses already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypeAddressMeta) ConvertTimestampColumnsToUTC(obj app_typ.AddressWithMeta) app_typ.AddressWithMeta {
	return obj
}

func (meta TypeAddressMeta) SetDefaultFieldValues(obj app_typ.AddressWithMeta) app_typ.AddressWithMeta {
	return obj
}

type TypeAddressDALMeta struct {
	*TypeAddressMeta
	dalutil.TypeCommonDALMeta[app_typ.AddressWithMeta, app_typ.AddressField]
}

var _TypeAddressDALMeta = TypeAddressDALMeta{
	TypeAddressMeta: &_TypeAddressMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[app_typ.AddressWithMeta, app_typ.AddressField]{
		DatabaseColumnFields: []app_typ.AddressField{
			app_typ.AddressField_ParentID,
			app_typ.AddressField_ID,
			app_typ.AddressField_Line1,
			app_typ.AddressField_Line2,
			app_typ.AddressField_City,
			app_typ.AddressField_State,
			app_typ.AddressField_PostalCode,
			app_typ.AddressField_Country,
			app_typ.AddressField_CreatedAt,
			app_typ.AddressField_UpdatedAt,
			app_typ.AddressField_DeletedAt,
		},
		DatabaseSubTableFields: []app_typ.AddressField{},
		SetInternallyByDALFields: []app_typ.AddressField{
			app_typ.AddressField_UpdatedAt,
			app_typ.AddressField_DeletedAt,
		},
		ImmutableFields: []app_typ.AddressField{
			app_typ.AddressField_ParentID,
			app_typ.AddressField_ID,
			app_typ.AddressField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []app_typ.AddressField{
			app_typ.AddressField_CreatedAt,
			app_typ.AddressField_UpdatedAt,
			app_typ.AddressField_DeletedAt,
		},
		UpdatedAtField: app_typ.AddressField_UpdatedAt,
	},
}

func GetTypeAddressDALMeta() *TypeAddressDALMeta {
	return &_TypeAddressDALMeta
}

func (meta *TypeAddressDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[app_typ.AddressWithMeta, app_typ.AddressField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeAddressDALMeta) GetDirectDBValues(obj app_typ.AddressWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.Line1,
		obj.Line2,
		obj.City,
		obj.State,
		obj.PostalCode,
		obj.Country,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeAddressDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj app_typ.AddressWithMeta) (app_typ.AddressWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeAddressDALMeta) ScanDBNextRow(rows *sql.Rows) (app_typ.AddressWithMeta, error) {
	var elem app_typ.AddressWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.Line1,
		&elem.Line2,
		&elem.City,
		&elem.State,
		&elem.PostalCode,
		&elem.Country,
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

func (meta TypeAddressDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []app_typ.AddressWithMeta) ([]app_typ.AddressWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeAddressDALMeta) GetChangedFieldsAndValues(old, new app_typ.AddressWithMeta, allowedFields []app_typ.AddressField) ([]app_typ.AddressField, []interface{}) {

	var colsWithValueChange []app_typ.AddressField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(app_typ.AddressField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_Line1, allowedFields) {
		if old.Line1 != new.Line1 {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_Line1)
			vals = append(vals, new.Line1)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_Line2, allowedFields) {
		if old.Line2 != new.Line2 {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_Line2)
			vals = append(vals, new.Line2)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_City, allowedFields) {
		if old.City != new.City {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_City)
			vals = append(vals, new.City)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_State, allowedFields) {
		if old.State != new.State {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_State)
			vals = append(vals, new.State)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_PostalCode, allowedFields) {
		if old.PostalCode != new.PostalCode {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_PostalCode)
			vals = append(vals, new.PostalCode)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_Country, allowedFields) {
		if old.Country != new.Country {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_Country)
			vals = append(vals, new.Country)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.AddressField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.AddressField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeAddressDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[app_typ.AddressWithMeta, app_typ.AddressField], allowedFields []app_typ.AddressField, elem app_typ.AddressWithMeta) (app_typ.AddressWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type TypeContactMeta struct {
	types.TypeCommonMeta[app_typ.ContactWithMeta, app_typ.ContactField]
}

var _TypeContactMeta = TypeContactMeta{
	TypeCommonMeta: types.TypeCommonMeta[app_typ.ContactWithMeta, app_typ.ContactField]{
		Name: naam.New("contact"),
		Fields: []app_typ.ContactField{
			app_typ.ContactField_ParentID,
			app_typ.ContactField_ID,
			app_typ.ContactField_Name,
			app_typ.ContactField_Name_ParentID,
			app_typ.ContactField_Name_ID,
			app_typ.ContactField_Name_FirstName,
			app_typ.ContactField_Name_MiddleInitial,
			app_typ.ContactField_Name_LastName,
			app_typ.ContactField_Name_CreatedAt,
			app_typ.ContactField_Name_UpdatedAt,
			app_typ.ContactField_Name_DeletedAt,
			app_typ.ContactField_Email,
			app_typ.ContactField_Address,
			app_typ.ContactField_Address_ParentID,
			app_typ.ContactField_Address_ID,
			app_typ.ContactField_Address_Line1,
			app_typ.ContactField_Address_Line2,
			app_typ.ContactField_Address_City,
			app_typ.ContactField_Address_State,
			app_typ.ContactField_Address_PostalCode,
			app_typ.ContactField_Address_Country,
			app_typ.ContactField_Address_CreatedAt,
			app_typ.ContactField_Address_UpdatedAt,
			app_typ.ContactField_Address_DeletedAt,
			app_typ.ContactField_CreatedAt,
			app_typ.ContactField_UpdatedAt,
			app_typ.ContactField_DeletedAt,
		},
	},
}

func GetTypeContactMeta() *TypeContactMeta {
	return &_TypeContactMeta
}

func (meta TypeContactMeta) SetMetaFieldValues(ctx context.Context, obj app_typ.ContactWithMeta, now time.Time) app_typ.ContactWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity Contacts already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity Contacts already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypeContactMeta) ConvertTimestampColumnsToUTC(obj app_typ.ContactWithMeta) app_typ.ContactWithMeta {
	return obj
}

func (meta TypeContactMeta) SetDefaultFieldValues(obj app_typ.ContactWithMeta) app_typ.ContactWithMeta {
	return obj
}

type TypeContactDALMeta struct {
	*TypeContactMeta
	dalutil.TypeCommonDALMeta[app_typ.ContactWithMeta, app_typ.ContactField]
}

var _TypeContactDALMeta = TypeContactDALMeta{
	TypeContactMeta: &_TypeContactMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[app_typ.ContactWithMeta, app_typ.ContactField]{
		DatabaseColumnFields: []app_typ.ContactField{
			app_typ.ContactField_ParentID,
			app_typ.ContactField_ID,
			app_typ.ContactField_Email,
			app_typ.ContactField_CreatedAt,
			app_typ.ContactField_UpdatedAt,
			app_typ.ContactField_DeletedAt,
		},
		DatabaseSubTableFields: []app_typ.ContactField{
			app_typ.ContactField_Name,
			app_typ.ContactField_Address,
		},
		SetInternallyByDALFields: []app_typ.ContactField{
			app_typ.ContactField_Name_UpdatedAt,
			app_typ.ContactField_Name_DeletedAt,
			app_typ.ContactField_Address_UpdatedAt,
			app_typ.ContactField_Address_DeletedAt,
			app_typ.ContactField_UpdatedAt,
			app_typ.ContactField_DeletedAt,
		},
		ImmutableFields: []app_typ.ContactField{
			app_typ.ContactField_ParentID,
			app_typ.ContactField_ID,
			app_typ.ContactField_Name_ParentID,
			app_typ.ContactField_Name_ID,
			app_typ.ContactField_Name_CreatedAt,
			app_typ.ContactField_Address_ParentID,
			app_typ.ContactField_Address_ID,
			app_typ.ContactField_Address_CreatedAt,
			app_typ.ContactField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []app_typ.ContactField{
			app_typ.ContactField_CreatedAt,
			app_typ.ContactField_UpdatedAt,
			app_typ.ContactField_DeletedAt,
		},
		UpdatedAtField: app_typ.ContactField_UpdatedAt,
	},
}

func GetTypeContactDALMeta() *TypeContactDALMeta {
	return &_TypeContactDALMeta
}

func (meta *TypeContactDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[app_typ.ContactWithMeta, app_typ.ContactField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypeContactDALMeta) GetDirectDBValues(obj app_typ.ContactWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.Email,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypeContactDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj app_typ.ContactWithMeta) (app_typ.ContactWithMeta, error) {

	// Insert 1:1 sub-tables
	// Insert Name
	{
		// Populate the ParentID for nested object, in case it is not populated
		if !obj.Name.ParentID.IsEmpty() && obj.Name.ParentID != obj.ID {
			return obj, fmt.Errorf("Adding Contact: nested object ParentID (%s) is different than the Contact ID (%s)", obj.Name.ParentID, obj.ID)
		}
		obj.Name.ParentID = obj.ID
		req := db.InsertTypeParams{
			TableName: params.TableName + "_" + app_typ.ContactField_Name.Name().FormatSQLTable(),
		}
		subMeta := GetTypePersonNameDALMeta()
		subElems, err := dalutil.BatchAddType[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, req, subMeta, obj.Name)
		if err != nil {
			return obj, fmt.Errorf("inserting Name: %w", err)
		}
		obj.Name = subElems[0]
	}
	// Insert Address
	{
		// Populate the ParentID for nested object, in case it is not populated
		if !obj.Address.ParentID.IsEmpty() && obj.Address.ParentID != obj.ID {
			return obj, fmt.Errorf("Adding Contact: nested object ParentID (%s) is different than the Contact ID (%s)", obj.Address.ParentID, obj.ID)
		}
		obj.Address.ParentID = obj.ID
		req := db.InsertTypeParams{
			TableName: params.TableName + "_" + app_typ.ContactField_Address.Name().FormatSQLTable(),
		}
		subMeta := GetTypeAddressDALMeta()
		subElems, err := dalutil.BatchAddType[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, req, subMeta, obj.Address)
		if err != nil {
			return obj, fmt.Errorf("inserting Address: %w", err)
		}
		obj.Address = subElems[0]
	}

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypeContactDALMeta) ScanDBNextRow(rows *sql.Rows) (app_typ.ContactWithMeta, error) {
	var elem app_typ.ContactWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.Email,
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

func (meta TypeContactDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []app_typ.ContactWithMeta) ([]app_typ.ContactWithMeta, error) {
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
		subMeta := GetTypePersonNameDALMeta()
		subResp, err := dalutil.ListTypeByIDs[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, subParams, subMeta)
		if err != nil {
			return nil, fmt.Errorf("fetching PersonName: %w", err)
		}
		// assign the right PersonNameWithMeta to the right elem
		var subElemMap = make(map[scalars.ID]app_typ.PersonNameWithMeta)
		for _, subElem := range subResp.Items {
			if _, alreadyExists := subElemMap[subElem.ParentID]; alreadyExists {
				return nil, fmt.Errorf("fetching PersonName: got multiple elements for Contact.Name field but we expect only one element")
			}
			subElemMap[subElem.ParentID] = subElem
		}

		for i := range elems {
			subElem, exists := subElemMap[elems[i].ID]
			if !exists {
				return nil, fmt.Errorf("fetching PersonName: did not find any element for Contact.Name field for Contact with ID %s", elems[i].ID)
			}
			elems[i].Name = subElem
		}
	}
	{
		// Fetch Address (type AddressWithMeta)
		subParams := db.ListTypeByIDsParams{
			TableName: params.TableName + "_" + "address",
			IDColumn:  "parent_id",
			IDs:       ids,
		}
		subMeta := GetTypeAddressDALMeta()
		subResp, err := dalutil.ListTypeByIDs[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, subParams, subMeta)
		if err != nil {
			return nil, fmt.Errorf("fetching Address: %w", err)
		}
		// assign the right AddressWithMeta to the right elem
		var subElemMap = make(map[scalars.ID]app_typ.AddressWithMeta)
		for _, subElem := range subResp.Items {
			if _, alreadyExists := subElemMap[subElem.ParentID]; alreadyExists {
				return nil, fmt.Errorf("fetching Address: got multiple elements for Contact.Address field but we expect only one element")
			}
			subElemMap[subElem.ParentID] = subElem
		}

		for i := range elems {
			subElem, exists := subElemMap[elems[i].ID]
			if !exists {
				return nil, fmt.Errorf("fetching Address: did not find any element for Contact.Address field for Contact with ID %s", elems[i].ID)
			}
			elems[i].Address = subElem
		}
	}

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypeContactDALMeta) GetChangedFieldsAndValues(old, new app_typ.ContactWithMeta, allowedFields []app_typ.ContactField) ([]app_typ.ContactField, []interface{}) {

	var colsWithValueChange []app_typ.ContactField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(app_typ.ContactField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(app_typ.ContactField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(app_typ.ContactField_Email, allowedFields) {
		if old.Email != new.Email {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_Email)
			vals = append(vals, new.Email)
		}
	}
	if types.IsFieldInFields(app_typ.ContactField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.ContactField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.ContactField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.ContactField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypeContactDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[app_typ.ContactWithMeta, app_typ.ContactField], allowedFields []app_typ.ContactField, elem app_typ.ContactWithMeta) (app_typ.ContactWithMeta, error) {
	// Update Nested (1:1 & 1:Many)
	// Update Name
	{
		subElem := elem.Name
		// Populate subElem.ParentID in case it is missing
		if !subElem.ParentID.IsEmpty() && subElem.ParentID != elem.ID {
			return elem, fmt.Errorf("Updating Contact: Nested object Name's ParentID (%s) is different than the parent Contact's ID (%s)", subElem.ParentID, elem.ID)
		}
		subElem.ParentID = elem.ID
		// TODO: If the 1:1 subobject does not have an ID, that's okay since we don't always expect users to provide
		// IDs for subobjects, so fetch it.
		subElemTableName := req.TableName + "_" + app_typ.ContactField_Name.Name().FormatSQLTable()
		if subElem.ID.IsEmpty() {
			subElemsListParams := db.ListTypeByIDsParams{
				TableName: subElemTableName,
				IDColumn:  "parent_id",
				IDs:       []scalars.ID{elem.ID},
			}
			subElemsListResp, err := dalutil.ListTypeByIDs[app_typ.ContactWithMeta, app_typ.ContactField](ctx, conn, subElemsListParams, meta)
			if err != nil {
				return elem, err
			}
			if len(subElemsListResp.Items) < 1 {
				return elem, fmt.Errorf("Updating Contact.Name: entity ID %s has no existing Name, don't know what to update", elem.ID)
			}
			if len(subElemsListResp.Items) > 1 {
				return elem, fmt.Errorf("Updating Contact.Name: entity ID %s has %d existing Name, we expected only one", elem.ID, len(subElemsListResp.Items))
			}
			subElem.ID = subElemsListResp.Items[0].ID
		}
		subReq := dalutil.UpdateTypeRequest[app_typ.PersonNameWithMeta, app_typ.PersonNameField]{
			TableName: subElemTableName,
			Object:    subElem,
			// Fields: <empty> : include all fields as we can't yet handle nested fields
			// ExcludeFields: <empty> : include all fields as we can't yet handle nested fields
		}
		subMeta := GetTypePersonNameDALMeta()
		subResp, err := dalutil.UpdateType[app_typ.PersonNameWithMeta, app_typ.PersonNameField](ctx, conn, subReq, subMeta)
		if err != nil {
			return elem, fmt.Errorf("Updating Name: %w", err)
		}
		elem.Name = subResp.Object

	}
	// Update Address
	{
		subElem := elem.Address
		// Populate subElem.ParentID in case it is missing
		if !subElem.ParentID.IsEmpty() && subElem.ParentID != elem.ID {
			return elem, fmt.Errorf("Updating Contact: Nested object Address's ParentID (%s) is different than the parent Contact's ID (%s)", subElem.ParentID, elem.ID)
		}
		subElem.ParentID = elem.ID
		// TODO: If the 1:1 subobject does not have an ID, that's okay since we don't always expect users to provide
		// IDs for subobjects, so fetch it.
		subElemTableName := req.TableName + "_" + app_typ.ContactField_Address.Name().FormatSQLTable()
		if subElem.ID.IsEmpty() {
			subElemsListParams := db.ListTypeByIDsParams{
				TableName: subElemTableName,
				IDColumn:  "parent_id",
				IDs:       []scalars.ID{elem.ID},
			}
			subElemsListResp, err := dalutil.ListTypeByIDs[app_typ.ContactWithMeta, app_typ.ContactField](ctx, conn, subElemsListParams, meta)
			if err != nil {
				return elem, err
			}
			if len(subElemsListResp.Items) < 1 {
				return elem, fmt.Errorf("Updating Contact.Address: entity ID %s has no existing Address, don't know what to update", elem.ID)
			}
			if len(subElemsListResp.Items) > 1 {
				return elem, fmt.Errorf("Updating Contact.Address: entity ID %s has %d existing Address, we expected only one", elem.ID, len(subElemsListResp.Items))
			}
			subElem.ID = subElemsListResp.Items[0].ID
		}
		subReq := dalutil.UpdateTypeRequest[app_typ.AddressWithMeta, app_typ.AddressField]{
			TableName: subElemTableName,
			Object:    subElem,
			// Fields: <empty> : include all fields as we can't yet handle nested fields
			// ExcludeFields: <empty> : include all fields as we can't yet handle nested fields
		}
		subMeta := GetTypeAddressDALMeta()
		subResp, err := dalutil.UpdateType[app_typ.AddressWithMeta, app_typ.AddressField](ctx, conn, subReq, subMeta)
		if err != nil {
			return elem, fmt.Errorf("Updating Address: %w", err)
		}
		elem.Address = subResp.Object

	}

	return elem, nil
}

type TypePersonNameMeta struct {
	types.TypeCommonMeta[app_typ.PersonNameWithMeta, app_typ.PersonNameField]
}

var _TypePersonNameMeta = TypePersonNameMeta{
	TypeCommonMeta: types.TypeCommonMeta[app_typ.PersonNameWithMeta, app_typ.PersonNameField]{
		Name: naam.New("person_name"),
		Fields: []app_typ.PersonNameField{
			app_typ.PersonNameField_ParentID,
			app_typ.PersonNameField_ID,
			app_typ.PersonNameField_FirstName,
			app_typ.PersonNameField_MiddleInitial,
			app_typ.PersonNameField_LastName,
			app_typ.PersonNameField_CreatedAt,
			app_typ.PersonNameField_UpdatedAt,
			app_typ.PersonNameField_DeletedAt,
		},
	},
}

func GetTypePersonNameMeta() *TypePersonNameMeta {
	return &_TypePersonNameMeta
}

func (meta TypePersonNameMeta) SetMetaFieldValues(ctx context.Context, obj app_typ.PersonNameWithMeta, now time.Time) app_typ.PersonNameWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity PersonNames already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity PersonNames already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypePersonNameMeta) ConvertTimestampColumnsToUTC(obj app_typ.PersonNameWithMeta) app_typ.PersonNameWithMeta {
	return obj
}

func (meta TypePersonNameMeta) SetDefaultFieldValues(obj app_typ.PersonNameWithMeta) app_typ.PersonNameWithMeta {
	return obj
}

type TypePersonNameDALMeta struct {
	*TypePersonNameMeta
	dalutil.TypeCommonDALMeta[app_typ.PersonNameWithMeta, app_typ.PersonNameField]
}

var _TypePersonNameDALMeta = TypePersonNameDALMeta{
	TypePersonNameMeta: &_TypePersonNameMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[app_typ.PersonNameWithMeta, app_typ.PersonNameField]{
		DatabaseColumnFields: []app_typ.PersonNameField{
			app_typ.PersonNameField_ParentID,
			app_typ.PersonNameField_ID,
			app_typ.PersonNameField_FirstName,
			app_typ.PersonNameField_MiddleInitial,
			app_typ.PersonNameField_LastName,
			app_typ.PersonNameField_CreatedAt,
			app_typ.PersonNameField_UpdatedAt,
			app_typ.PersonNameField_DeletedAt,
		},
		DatabaseSubTableFields: []app_typ.PersonNameField{},
		SetInternallyByDALFields: []app_typ.PersonNameField{
			app_typ.PersonNameField_UpdatedAt,
			app_typ.PersonNameField_DeletedAt,
		},
		ImmutableFields: []app_typ.PersonNameField{
			app_typ.PersonNameField_ParentID,
			app_typ.PersonNameField_ID,
			app_typ.PersonNameField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []app_typ.PersonNameField{
			app_typ.PersonNameField_CreatedAt,
			app_typ.PersonNameField_UpdatedAt,
			app_typ.PersonNameField_DeletedAt,
		},
		UpdatedAtField: app_typ.PersonNameField_UpdatedAt,
	},
}

func GetTypePersonNameDALMeta() *TypePersonNameDALMeta {
	return &_TypePersonNameDALMeta
}

func (meta *TypePersonNameDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[app_typ.PersonNameWithMeta, app_typ.PersonNameField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypePersonNameDALMeta) GetDirectDBValues(obj app_typ.PersonNameWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.FirstName,
		obj.MiddleInitial,
		obj.LastName,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypePersonNameDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj app_typ.PersonNameWithMeta) (app_typ.PersonNameWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypePersonNameDALMeta) ScanDBNextRow(rows *sql.Rows) (app_typ.PersonNameWithMeta, error) {
	var elem app_typ.PersonNameWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.FirstName,
		&elem.MiddleInitial,
		&elem.LastName,
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

func (meta TypePersonNameDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []app_typ.PersonNameWithMeta) ([]app_typ.PersonNameWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypePersonNameDALMeta) GetChangedFieldsAndValues(old, new app_typ.PersonNameWithMeta, allowedFields []app_typ.PersonNameField) ([]app_typ.PersonNameField, []interface{}) {

	var colsWithValueChange []app_typ.PersonNameField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(app_typ.PersonNameField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_FirstName, allowedFields) {
		if old.FirstName != new.FirstName {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_FirstName)
			vals = append(vals, new.FirstName)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_MiddleInitial, allowedFields) {
		if old.MiddleInitial != new.MiddleInitial {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_MiddleInitial)
			vals = append(vals, new.MiddleInitial)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_LastName, allowedFields) {
		if old.LastName != new.LastName {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_LastName)
			vals = append(vals, new.LastName)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.PersonNameField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PersonNameField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypePersonNameDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[app_typ.PersonNameWithMeta, app_typ.PersonNameField], allowedFields []app_typ.PersonNameField, elem app_typ.PersonNameWithMeta) (app_typ.PersonNameWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type TypePhoneNumberMeta struct {
	types.TypeCommonMeta[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField]
}

var _TypePhoneNumberMeta = TypePhoneNumberMeta{
	TypeCommonMeta: types.TypeCommonMeta[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField]{
		Name: naam.New("phone_number"),
		Fields: []app_typ.PhoneNumberField{
			app_typ.PhoneNumberField_ParentID,
			app_typ.PhoneNumberField_ID,
			app_typ.PhoneNumberField_CountryCode,
			app_typ.PhoneNumberField_Number,
			app_typ.PhoneNumberField_Extension,
			app_typ.PhoneNumberField_CreatedAt,
			app_typ.PhoneNumberField_UpdatedAt,
			app_typ.PhoneNumberField_DeletedAt,
		},
	},
}

func GetTypePhoneNumberMeta() *TypePhoneNumberMeta {
	return &_TypePhoneNumberMeta
}

func (meta TypePhoneNumberMeta) SetMetaFieldValues(ctx context.Context, obj app_typ.PhoneNumberWithMeta, now time.Time) app_typ.PhoneNumberWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		llog.Warn(ctx, "Entity PhoneNumbers already has ID set", "value", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		llog.Warn(ctx, "Entity PhoneNumbers already has field CreatedAt set", "value", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta TypePhoneNumberMeta) ConvertTimestampColumnsToUTC(obj app_typ.PhoneNumberWithMeta) app_typ.PhoneNumberWithMeta {
	return obj
}

func (meta TypePhoneNumberMeta) SetDefaultFieldValues(obj app_typ.PhoneNumberWithMeta) app_typ.PhoneNumberWithMeta {
	return obj
}

type TypePhoneNumberDALMeta struct {
	*TypePhoneNumberMeta
	dalutil.TypeCommonDALMeta[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField]
}

var _TypePhoneNumberDALMeta = TypePhoneNumberDALMeta{
	TypePhoneNumberMeta: &_TypePhoneNumberMeta,
	TypeCommonDALMeta: dalutil.TypeCommonDALMeta[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField]{
		DatabaseColumnFields: []app_typ.PhoneNumberField{
			app_typ.PhoneNumberField_ParentID,
			app_typ.PhoneNumberField_ID,
			app_typ.PhoneNumberField_CountryCode,
			app_typ.PhoneNumberField_Number,
			app_typ.PhoneNumberField_Extension,
			app_typ.PhoneNumberField_CreatedAt,
			app_typ.PhoneNumberField_UpdatedAt,
			app_typ.PhoneNumberField_DeletedAt,
		},
		DatabaseSubTableFields: []app_typ.PhoneNumberField{},
		SetInternallyByDALFields: []app_typ.PhoneNumberField{
			app_typ.PhoneNumberField_UpdatedAt,
			app_typ.PhoneNumberField_DeletedAt,
		},
		ImmutableFields: []app_typ.PhoneNumberField{
			app_typ.PhoneNumberField_ParentID,
			app_typ.PhoneNumberField_ID,
			app_typ.PhoneNumberField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []app_typ.PhoneNumberField{
			app_typ.PhoneNumberField_CreatedAt,
			app_typ.PhoneNumberField_UpdatedAt,
			app_typ.PhoneNumberField_DeletedAt,
		},
		UpdatedAtField: app_typ.PhoneNumberField_UpdatedAt,
	},
}

func GetTypePhoneNumberDALMeta() *TypePhoneNumberDALMeta {
	return &_TypePhoneNumberDALMeta
}

func (meta *TypePhoneNumberDALMeta) GetCommonDALMeta() *dalutil.TypeCommonDALMeta[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField] {
	return &meta.TypeCommonDALMeta
}

func (meta TypePhoneNumberDALMeta) GetDirectDBValues(obj app_typ.PhoneNumberWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.CountryCode,
		obj.Number,
		obj.Extension,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta TypePhoneNumberDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj app_typ.PhoneNumberWithMeta) (app_typ.PhoneNumberWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta TypePhoneNumberDALMeta) ScanDBNextRow(rows *sql.Rows) (app_typ.PhoneNumberWithMeta, error) {
	var elem app_typ.PhoneNumberWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.CountryCode,
		&elem.Number,
		&elem.Extension,
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

func (meta TypePhoneNumberDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []app_typ.PhoneNumberWithMeta) ([]app_typ.PhoneNumberWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta TypePhoneNumberDALMeta) GetChangedFieldsAndValues(old, new app_typ.PhoneNumberWithMeta, allowedFields []app_typ.PhoneNumberField) ([]app_typ.PhoneNumberField, []interface{}) {

	var colsWithValueChange []app_typ.PhoneNumberField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(app_typ.PhoneNumberField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_CountryCode, allowedFields) {
		if old.CountryCode != new.CountryCode {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_CountryCode)
			vals = append(vals, new.CountryCode)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_Number, allowedFields) {
		if old.Number != new.Number {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_Number)
			vals = append(vals, new.Number)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_Extension, allowedFields) {
		if old.Extension != new.Extension {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_Extension)
			vals = append(vals, new.Extension)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(app_typ.PhoneNumberField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, app_typ.PhoneNumberField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta *TypePhoneNumberDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[app_typ.PhoneNumberWithMeta, app_typ.PhoneNumberField], allowedFields []app_typ.PhoneNumberField, elem app_typ.PhoneNumberWithMeta) (app_typ.PhoneNumberWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
