package svccore_dal

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	filterlib "github.com/teejays/gokutil/filter"

	app_dal "sampleapp/backend/.goku/generated/dal"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
)

var llog = log.GetLogger().WithHeading("DAL").WithHeading("$.service[Core]")

/* Type Methods */

// DAL encapsulates DAL methods for types that fall under Core
type DAL struct {
	app_dal.DAL

	conn db.Connection
}

func NewDAL(ctx context.Context, conn db.Connection) (DAL, error) {
	dal := DAL{conn: conn}
	parentDAL, err := app_dal.NewDAL(ctx, conn)
	if err != nil {
		return DAL{}, err
	}
	dal.DAL = parentDAL
	return dal, nil
}

// ConstructListQueryBuilderForTypeCronExpression provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all CronExpression items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypeCronExpression(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[svccore_typ.CronExpressionFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for CronExpression", "request", req)
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

	if req.Filter.ParentID != nil { // inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.ID != nil { // inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.Second != nil { // inject a Where filter for Second
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Second, ds, "second", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.Minute != nil { // inject a Where filter for Minute
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Minute, ds, "minute", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.Hour != nil { // inject a Where filter for Hour
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Hour, ds, "hour", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.DayOfMonth != nil { // inject a Where filter for DayOfMonth
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DayOfMonth, ds, "day_of_month", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.Month != nil { // inject a Where filter for Month
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Month, ds, "month", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.DayOfWeek != nil { // inject a Where filter for DayOfWeek
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DayOfWeek, ds, "day_of_week", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.CreatedAt != nil { // inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.UpdatedAt != nil { // inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.DeletedAt != nil { // inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at", false)
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d DAL) GetQueryByTextFilterForTypeCronExpression(ctx context.Context, queryText string) (*svccore_typ.CronExpressionFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []svccore_typ.CronExpressionFilter
	{
		filter := svccore_typ.CronExpressionFilter{Second: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_typ.CronExpressionFilter{Minute: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_typ.CronExpressionFilter{Hour: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_typ.CronExpressionFilter{DayOfMonth: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_typ.CronExpressionFilter{Month: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_typ.CronExpressionFilter{DayOfWeek: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := svccore_typ.CronExpressionFilter{
		Or: orFilters,
	}

	return &filter, nil
}
