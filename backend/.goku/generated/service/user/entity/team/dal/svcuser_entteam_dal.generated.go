package svcuser_entteam_dal

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/dalutil"
	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/panics"
	"github.com/teejays/gokutil/scalars"

	svcuser_dal "sampleapp/backend/.goku/generated/service/user/dal"
	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entteam_meta "sampleapp/backend/.goku/generated/service/user/entity/team/meta"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
)

var llog = log.GetLogger().WithHeading("DAL").WithHeading("$.service[User].entity[Team]")

/* Entity Client */

type Client struct {
	conn db.Connection
	dal  DAL
}

func NewClient(ctx context.Context) (svcuser_entteam_client.CRUDClient, error) {
	// Get DB Connection
	databaseName := "user"
	conn, err := db.NewConnection(databaseName)
	if err != nil {
		return nil, fmt.Errorf("Getting DB connection for database [%s]: %w", databaseName, err)
	}

	dal, err := NewDAL(ctx, conn)
	if err != nil {
		return nil, err
	}
	return Client{conn: conn, dal: dal}, nil
}

// Add handles insertion of a single Team entity in the database.
func (c Client) Add(ctx context.Context, req svcuser_entteam_typ.AddRequest) (svcuser_entteam_typ.Team, error) {
	llog.Info(ctx, "Add entity", "entity", "Team", "request", req)

	elems, err := c.AddBatch(ctx, req)
	if err != nil {
		return svcuser_entteam_typ.Team{}, err
	}
	panics.If(len(elems) != 1, "Add%s: Expected 1 element to be inserted, but got %d", "Team", len(elems))
	return elems[0], nil
}

// AddBatch handles insertion of multiple Teams in the database.
func (c Client) AddBatch(ctx context.Context, reqs ...svcuser_entteam_typ.AddRequest) ([]svcuser_entteam_typ.Team, error) {
	llog.Info(ctx, "AddTeams: Adding entities...", "entity", "Team", "request", reqs)

	// Meta info
	meta := svcuser_entteam_meta.GetEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := c.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var elems []svcuser_entteam_typ.Team
	for _, r := range reqs {
		obj := svcuser_entteam_typ.NewTeamWithMetaFromInput(ctx, r.Object)
		elems = append(elems, obj)
	}

	resp, err := dalutil.BatchAddType[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField](ctx, c.conn, params, meta.GetTypeDALMeta(), elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Teams", err)
	}

	err = c.conn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get fetches the entity Team based on the ID provided.
func (c Client) Get(ctx context.Context, req svcuser_entteam_typ.GetRequest) (svcuser_entteam_typ.Team, error) {
	llog.Info(ctx, "Get entity", "entity", "Team", "request", req)
	id := req.ID
	var elem svcuser_entteam_typ.Team

	listReq := svcuser_entteam_typ.ListRequest{
		Filter: svcuser_entteam_typ.TeamFilter{
			ID: &filterlib.IDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := c.List(ctx, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Teams by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No Team found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Teams found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

// QueryByText searches the entity Team by the given text query.
func (c Client) QueryByText(ctx context.Context, req svcuser_entteam_typ.QueryByTextRequest) (svcuser_entteam_typ.ListResponse, error) {

	var resp svcuser_entteam_typ.ListResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := c.dal.GetQueryByTextFilterForTypeTeam(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		llog.Error(ctx, "Text search is not set up", "type", "Team")
		return resp, nil
	}

	listReq := svcuser_entteam_typ.ListRequest{
		Filter: *filter,
	}

	listResp, err := c.List(ctx, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Teams by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// List fetches a list of Team entities based on the given parameters.
func (c Client) List(ctx context.Context, req svcuser_entteam_typ.ListRequest) (svcuser_entteam_typ.ListResponse, error) {
	llog.Info(ctx, "List entities", "entity", "Team", "request", req)
	var resp svcuser_entteam_typ.ListResponse

	// Meta info
	meta := svcuser_entteam_meta.GetEntityDALMeta()

	// Begin a Transaction
	err := c.conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique Team IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := c.listIDs(ctx, req, params)
	if err != nil {
		return resp, err
	}

	llog.Debug(ctx, "Listing Team", "ids", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField](ctx, c.conn, listTypeByIDsParams, meta.GetTypeDALMeta())
	if err != nil {
		return resp, err
	}

	resp.Items = typeResp.Items
	resp.Count = len(typeResp.Items)

	// Commit the transaction
	err = c.conn.Commit(ctx)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

// listIDs fetches a list of unique IDs for entity Team that match the filter provided.
func (c Client) listIDs(ctx context.Context, req svcuser_entteam_typ.ListRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := c.listUniqueIDsByFilter(ctx, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// listUniqueIDsByFilter fetches a list of unique IDs for entity Team that match the filter provided.
func (c Client) listUniqueIDsByFilter(ctx context.Context, filter svcuser_entteam_typ.TeamFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := c.listUniqueIDsByBaseFilter(ctx, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := c.listUniqueIDsByFilter(ctx, orFilter, params)
			if err != nil {
				return nil, err
			}
			orIDLists[i] = orFilterIDs
		}
		orIDs = dalutil.GetUUIDsUnion(orIDLists...)
		finalIDs = dalutil.GetUUIDsIntersection(finalIDs, orIDs)
	}

	var andIDs []scalars.ID
	if len(filter.And) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := c.listUniqueIDsByFilter(ctx, orFilter, params)
			if err != nil {
				return nil, err
			}
			orIDLists[i] = orFilterIDs
		}
		andIDs = dalutil.GetUUIDsIntersection(orIDLists...)
		finalIDs = dalutil.GetUUIDsIntersection(finalIDs, andIDs)
	}

	return finalIDs, nil
}

// listUniqueIDsByBaseFilter fetches a list of unique IDs for entity Team that match the filter provided, ignoring the OR and AND filters.
func (c Client) listUniqueIDsByBaseFilter(ctx context.Context, filter svcuser_entteam_typ.TeamFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := dalutil.ListTypeRequest[svcuser_entteam_typ.TeamFilter]{Filter: filter}
	ds, err := c.dal.ConstructListQueryBuilderForTypeTeam(ctx, c.conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique Team IDs that satisfy the filters: %w", err)
	}

	// Construct the query
	query, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}

	// Run the query
	rows, err := c.conn.QueryRows(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids, err := db.SqlRowsToUUIDs(ctx, rows)
	if err != nil {
		return nil, fmt.Errorf("Parsing sql.Rows into []ID: %w", err)
	}

	return ids, nil
}

// Update handles the update of Team entity.
func (c Client) Update(ctx context.Context, req svcuser_entteam_typ.UpdateRequest) (dalutil.UpdateEntityResponse[svcuser_entteam_typ.Team], error) {
	llog.Info(ctx, "Update entity", "entity", "Team", "request", req)

	var resp dalutil.UpdateEntityResponse[svcuser_entteam_typ.Team]

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := svcuser_entteam_meta.GetEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField]{
		TableName:     meta.DbTableName.FormatSQLTable(),
		Object:        req.Object,
		Fields:        req.Fields,
		ExcludeFields: req.ExcludeFields,
	}

	// Begin a Transaction
	err := c.conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	_, err = dalutil.UpdateType(ctx, c.conn, subReq, meta.GetTypeDALMeta())
	if err != nil {
		c.conn.MustRollback(ctx)
		return resp, fmt.Errorf("Updating Team: %w", err)
	}

	err = c.conn.Commit(ctx)
	if err != nil {
		c.conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	var getReq svcuser_entteam_typ.GetRequest
	getReq.ID = req.Object.ID
	entity, err := c.Get(ctx, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating Team: failed to get Team entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}

/* Type Methods */

// DAL encapsulates DAL methods for types that fall under Team
type DAL struct {
	svcuser_dal.DAL

	conn db.Connection
}

func NewDAL(ctx context.Context, conn db.Connection) (DAL, error) {
	dal := DAL{conn: conn}
	parentDAL, err := svcuser_dal.NewDAL(ctx, conn)
	if err != nil {
		return DAL{}, err
	}
	dal.DAL = parentDAL
	return dal, nil
}

// ConstructListQueryBuilderForTypeTeam provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Team items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypeTeam(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[svcuser_entteam_typ.TeamFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for Team", "request", req)
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

	if req.Filter.ID != nil { // inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id", false)
		if err != nil {
			return resp, err
		}
	}

	if req.Filter.Name != nil { // inject a Where filter for Name
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Name, ds, "name", false)
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

func (d DAL) GetQueryByTextFilterForTypeTeam(ctx context.Context, queryText string) (*svcuser_entteam_typ.TeamFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []svcuser_entteam_typ.TeamFilter
	{
		filter := svcuser_entteam_typ.TeamFilter{Name: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := svcuser_entteam_typ.TeamFilter{
		Or: orFilters,
	}

	return &filter, nil
}
