package app_dal

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	filterlib "github.com/teejays/gokutil/filter"

	app_typ "sampleapp/backend/.goku/generated/typ"
)

var llog = log.GetLogger().WithHeading("DAL").WithHeading("?")

/* Type Methods */

// DAL encapsulates DAL methods for types that fall under Sampleapp
type DAL struct {
	conn db.Connection
}

func NewDAL(ctx context.Context, conn db.Connection) (DAL, error) {
	dal := DAL{conn: conn}
	return dal, nil
}

// ConstructListQueryBuilderForTypeAddress provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Address items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypeAddress(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[app_typ.AddressFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for Address", "request", req)
	var resp db.SelectQueryBuilder
	var err error

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables

	// Handle Nested objects 1:1

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ParentID != nil {
		// inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Line1 != nil {
		// inject a Where filter for Line1
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Line1, ds, "line_1")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Line2 != nil {
		// inject a Where filter for Line2
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Line2, ds, "line_2")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.City != nil {
		// inject a Where filter for City
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.City, ds, "city")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.State != nil {
		// inject a Where filter for State
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.State, ds, "state")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.PostalCode != nil {
		// inject a Where filter for PostalCode
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PostalCode, ds, "postal_code")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Country != nil {
		// inject a Where filter for Country
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Country, ds, "country")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d DAL) GetQueryByTextFilterForTypeAddress(ctx context.Context, queryText string) (*app_typ.AddressFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []app_typ.AddressFilter
	{
		filter := app_typ.AddressFilter{Line1: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.AddressFilter{Line2: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.AddressFilter{City: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.AddressFilter{State: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.AddressFilter{PostalCode: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := app_typ.AddressFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ConstructListQueryBuilderForTypeContact provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Contact items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypeContact(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[app_typ.ContactFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for Contact", "request", req)
	var resp db.SelectQueryBuilder
	var err error

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables

	// Handle Nested objects 1:1
	if req.Filter.Name != nil {
		subTableName := params.TableName + "_" + "name"
		subReq := dalutil.ListTypeRequest[app_typ.PersonNameFilter]{
			Filter: *req.Filter.Name,
		}
		subParams := db.UniqueIDsQueryBuilderParams{
			TableName:     subTableName,
			SelectColumns: []string{"id", "parent_id"},
		}
		subQuery, err := d.ConstructListQueryBuilderForTypePersonName(ctx, d.conn, subReq, subParams)
		if err != nil {
			return resp, err
		}
		// Add With commands
		for _, with := range subQuery.Withs {
			resp.Withs = append(resp.Withs, with)
		}
		// Add the Select command as with too
		resp.Withs = append(resp.Withs, db.With{Alias: "cte_" + subTableName, Select: subQuery.Select})

		// Add Join to the new With Table
		ds = ds.Join(
			goqu.T("cte_"+subTableName), goqu.On(goqu.I("cte_"+subTableName+".parent_id").Eq(goqu.I(params.TableName+".id"))),
		)
	}
	if req.Filter.Address != nil {
		subTableName := params.TableName + "_" + "address"
		subReq := dalutil.ListTypeRequest[app_typ.AddressFilter]{
			Filter: *req.Filter.Address,
		}
		subParams := db.UniqueIDsQueryBuilderParams{
			TableName:     subTableName,
			SelectColumns: []string{"id", "parent_id"},
		}
		subQuery, err := d.ConstructListQueryBuilderForTypeAddress(ctx, d.conn, subReq, subParams)
		if err != nil {
			return resp, err
		}
		// Add With commands
		for _, with := range subQuery.Withs {
			resp.Withs = append(resp.Withs, with)
		}
		// Add the Select command as with too
		resp.Withs = append(resp.Withs, db.With{Alias: "cte_" + subTableName, Select: subQuery.Select})

		// Add Join to the new With Table
		ds = ds.Join(
			goqu.T("cte_"+subTableName), goqu.On(goqu.I("cte_"+subTableName+".parent_id").Eq(goqu.I(params.TableName+".id"))),
		)
	}

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ParentID != nil {
		// inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Email != nil {
		// inject a Where filter for Email
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Email, ds, "email")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d DAL) GetQueryByTextFilterForTypeContact(ctx context.Context, queryText string) (*app_typ.ContactFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []app_typ.ContactFilter
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypePersonName(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing PersonName query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, app_typ.ContactFilter{Name: &nestedFilter.Or[i]})
			}
		}
	}
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypeAddress(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing Address query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, app_typ.ContactFilter{Address: &nestedFilter.Or[i]})
			}
		}
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := app_typ.ContactFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ConstructListQueryBuilderForTypePersonName provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all PersonName items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypePersonName(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[app_typ.PersonNameFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for PersonName", "request", req)
	var resp db.SelectQueryBuilder
	var err error

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables

	// Handle Nested objects 1:1

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ParentID != nil {
		// inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.FirstName != nil {
		// inject a Where filter for FirstName
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.FirstName, ds, "first_name")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.MiddleInitial != nil {
		// inject a Where filter for MiddleInitial
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.MiddleInitial, ds, "middle_initial")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.LastName != nil {
		// inject a Where filter for LastName
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.LastName, ds, "last_name")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d DAL) GetQueryByTextFilterForTypePersonName(ctx context.Context, queryText string) (*app_typ.PersonNameFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []app_typ.PersonNameFilter
	{
		filter := app_typ.PersonNameFilter{FirstName: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.PersonNameFilter{MiddleInitial: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.PersonNameFilter{LastName: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := app_typ.PersonNameFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ConstructListQueryBuilderForTypePhoneNumber provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all PhoneNumber items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypePhoneNumber(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[app_typ.PhoneNumberFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for PhoneNumber", "request", req)
	var resp db.SelectQueryBuilder
	var err error

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables

	// Handle Nested objects 1:1

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ParentID != nil {
		// inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CountryCode != nil {
		// inject a Where filter for CountryCode
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CountryCode, ds, "country_code")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Number != nil {
		// inject a Where filter for Number
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Number, ds, "number")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Extension != nil {
		// inject a Where filter for Extension
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Extension, ds, "extension")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d DAL) GetQueryByTextFilterForTypePhoneNumber(ctx context.Context, queryText string) (*app_typ.PhoneNumberFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []app_typ.PhoneNumberFilter
	{
		filter := app_typ.PhoneNumberFilter{Number: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := app_typ.PhoneNumberFilter{Extension: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := app_typ.PhoneNumberFilter{
		Or: orFilters,
	}

	return &filter, nil
}
