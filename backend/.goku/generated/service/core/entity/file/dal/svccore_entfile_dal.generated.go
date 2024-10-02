package svccore_entfile_dal

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

	svccore_dal "sampleapp/backend/.goku/generated/service/core/dal"
	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_meta "sampleapp/backend/.goku/generated/service/core/entity/file/meta"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

var llog = log.GetLogger().WithHeading("DAL").WithHeading("?.service[Core].entity[File]")

/* Entity Client */

type Client struct {
	conn db.Connection
	dal  DAL
}

func NewClient(ctx context.Context) (svccore_entfile_client.CRUDClient, error) {
	// Get DB Connection
	databaseName := "core"
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

// Add handles insertion of a single File entity in the database.
func (c Client) Add(ctx context.Context, req svccore_entfile_typ.AddRequest) (svccore_entfile_typ.File, error) {
	elems, err := c.AddBatch(ctx, req)
	if err != nil {
		return svccore_entfile_typ.File{}, err
	}
	panics.If(len(elems) != 1, "Add%s: Expected 1 element to be inserted, but got %d", "File", len(elems))
	return elems[0], nil
}

// AddBatch handles insertion of multiple Files in the database.
func (c Client) AddBatch(ctx context.Context, reqs ...svccore_entfile_typ.AddRequest) ([]svccore_entfile_typ.File, error) {
	llog.Debug(ctx, "AddFiles: Adding entities...", "request", reqs)

	// Meta info
	meta := svccore_entfile_meta.GetEntityFileDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := c.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var elems []svccore_entfile_typ.File
	for _, r := range reqs {
		obj := svccore_entfile_typ.NewFileWithMetaFromInput(ctx, r.Object)
		elems = append(elems, obj)
	}

	resp, err := dalutil.BatchAddType[svccore_entfile_typ.File, svccore_entfile_typ.FileField](ctx, c.conn, params, meta.GetTypeDALMeta(), elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Files", err)
	}

	err = c.conn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get fetches the entity File based on the ID provided.
func (c Client) Get(ctx context.Context, req svccore_entfile_typ.GetRequest) (svccore_entfile_typ.File, error) {
	id := req.ID
	var elem svccore_entfile_typ.File

	listReq := svccore_entfile_typ.ListRequest{
		Filter: svccore_entfile_typ.FileFilter{
			ID: &filterlib.IDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := c.List(ctx, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Files by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No File found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Files found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

// QueryByText searches the entity File by the given text query.
func (c Client) QueryByText(ctx context.Context, req svccore_entfile_typ.QueryByTextRequest) (svccore_entfile_typ.ListResponse, error) {

	var resp svccore_entfile_typ.ListResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := c.dal.GetQueryByTextFilterForTypeFile(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		llog.Error(ctx, "Text search is not set up", "type", "File")
		return resp, nil
	}

	listReq := svccore_entfile_typ.ListRequest{
		Filter: *filter,
	}

	listResp, err := c.List(ctx, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Files by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// List fetches a list of File entities based on the given parameters.
func (c Client) List(ctx context.Context, req svccore_entfile_typ.ListRequest) (svccore_entfile_typ.ListResponse, error) {
	var resp svccore_entfile_typ.ListResponse

	// Meta info
	meta := svccore_entfile_meta.GetEntityFileDALMeta()

	// Begin a Transaction
	err := c.conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique File IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := c.listIDs(ctx, req, params)
	if err != nil {
		return resp, err
	}

	llog.Debug(ctx, "Listing File", "ids", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[svccore_entfile_typ.File, svccore_entfile_typ.FileField](ctx, c.conn, listTypeByIDsParams, meta.GetTypeDALMeta())
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

// listIDs fetches a list of unique IDs for entity File that match the filter provided.
func (c Client) listIDs(ctx context.Context, req svccore_entfile_typ.ListRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := c.listUniqueIDsByFilter(ctx, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// listUniqueIDsByFilter fetches a list of unique IDs for entity File that match the filter provided.
func (c Client) listUniqueIDsByFilter(ctx context.Context, filter svccore_entfile_typ.FileFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

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

// listUniqueIDsByBaseFilter fetches a list of unique IDs for entity File that match the filter provided, ignoring the OR and AND filters.
func (c Client) listUniqueIDsByBaseFilter(ctx context.Context, filter svccore_entfile_typ.FileFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := dalutil.ListTypeRequest[svccore_entfile_typ.FileFilter]{Filter: filter}
	ds, err := c.dal.ConstructListQueryBuilderForTypeFile(ctx, c.conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique File IDs that satisfy the filters: %w", err)
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

// Update handles the update of File entity.
func (c Client) Update(ctx context.Context, req svccore_entfile_typ.UpdateRequest) (dalutil.UpdateEntityResponse[svccore_entfile_typ.File], error) {

	var resp dalutil.UpdateEntityResponse[svccore_entfile_typ.File]

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := svccore_entfile_meta.GetEntityFileDALMeta()

	subReq := dalutil.UpdateTypeRequest[svccore_entfile_typ.File, svccore_entfile_typ.FileField]{
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
		return resp, fmt.Errorf("Updating File: %w", err)
	}

	err = c.conn.Commit(ctx)
	if err != nil {
		c.conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	var getReq svccore_entfile_typ.GetRequest
	getReq.ID = req.Object.ID
	entity, err := c.Get(ctx, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating File: failed to get File entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}

/* Type Methods */

// DAL encapsulates DAL methods for types that fall under File
type DAL struct {
	svccore_dal.DAL

	conn db.Connection
}

func NewDAL(ctx context.Context, conn db.Connection) (DAL, error) {
	dal := DAL{conn: conn}
	parentDAL, err := svccore_dal.NewDAL(ctx, conn)
	if err != nil {
		return DAL{}, err
	}
	dal.DAL = parentDAL
	return dal, nil
}

// ConstructListQueryBuilderForTypeFile provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all File items that belong to `params.TableName` and match the filter.
func (d DAL) ConstructListQueryBuilderForTypeFile(ctx context.Context, conn db.Connection, req dalutil.ListTypeRequest[svccore_entfile_typ.FileFilter], params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	llog.Debug(ctx, "Constructing query builder for File", "request", req)
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
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.FileName != nil {
		// inject a Where filter for FileName
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.FileName, ds, "file_name")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.StorageClient != nil {
		// inject a Where filter for StorageClient
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.StorageClient, ds, "storage_client")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.StorageClientIdentifier != nil {
		// inject a Where filter for StorageClientIdentifier
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.StorageClientIdentifier, ds, "storage_client_identifier")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.SizeBytes != nil {
		// inject a Where filter for SizeBytes
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.SizeBytes, ds, "size_bytes")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.MimeType != nil {
		// inject a Where filter for MimeType
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.MimeType, ds, "mime_type")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.FileHash != nil {
		// inject a Where filter for FileHash
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.FileHash, ds, "file_hash")
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

func (d DAL) GetQueryByTextFilterForTypeFile(ctx context.Context, queryText string) (*svccore_entfile_typ.FileFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []svccore_entfile_typ.FileFilter
	{
		filter := svccore_entfile_typ.FileFilter{FileName: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_entfile_typ.FileFilter{StorageClientIdentifier: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_entfile_typ.FileFilter{MimeType: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := svccore_entfile_typ.FileFilter{FileHash: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := svccore_entfile_typ.FileFilter{
		Or: orFilters,
	}

	return &filter, nil
}
