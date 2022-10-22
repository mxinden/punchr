// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Router is an object representing the database table.
type Router struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClientID  int       `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	HTML      string    `boil:"html" json:"html" toml:"html" yaml:"html"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *routerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L routerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RouterColumns = struct {
	ID        string
	ClientID  string
	HTML      string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	ClientID:  "client_id",
	HTML:      "html",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

var RouterTableColumns = struct {
	ID        string
	ClientID  string
	HTML      string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "routers.id",
	ClientID:  "routers.client_id",
	HTML:      "routers.html",
	UpdatedAt: "routers.updated_at",
	CreatedAt: "routers.created_at",
}

// Generated where

var RouterWhere = struct {
	ID        whereHelperint
	ClientID  whereHelperint
	HTML      whereHelperstring
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"routers\".\"id\""},
	ClientID:  whereHelperint{field: "\"routers\".\"client_id\""},
	HTML:      whereHelperstring{field: "\"routers\".\"html\""},
	UpdatedAt: whereHelpertime_Time{field: "\"routers\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"routers\".\"created_at\""},
}

// RouterRels is where relationship names are stored.
var RouterRels = struct {
	Client string
}{
	Client: "Client",
}

// routerR is where relationships are stored.
type routerR struct {
	Client *Peer `boil:"Client" json:"Client" toml:"Client" yaml:"Client"`
}

// NewStruct creates a new relationship struct
func (*routerR) NewStruct() *routerR {
	return &routerR{}
}

func (r *routerR) GetClient() *Peer {
	if r == nil {
		return nil
	}
	return r.Client
}

// routerL is where Load methods for each relationship are stored.
type routerL struct{}

var (
	routerAllColumns            = []string{"id", "client_id", "html", "updated_at", "created_at"}
	routerColumnsWithoutDefault = []string{"client_id", "html", "updated_at", "created_at"}
	routerColumnsWithDefault    = []string{"id"}
	routerPrimaryKeyColumns     = []string{"id"}
	routerGeneratedColumns      = []string{"id"}
)

type (
	// RouterSlice is an alias for a slice of pointers to Router.
	// This should almost always be used instead of []Router.
	RouterSlice []*Router
	// RouterHook is the signature for custom Router hook methods
	RouterHook func(context.Context, boil.ContextExecutor, *Router) error

	routerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	routerType                 = reflect.TypeOf(&Router{})
	routerMapping              = queries.MakeStructMapping(routerType)
	routerPrimaryKeyMapping, _ = queries.BindMapping(routerType, routerMapping, routerPrimaryKeyColumns)
	routerInsertCacheMut       sync.RWMutex
	routerInsertCache          = make(map[string]insertCache)
	routerUpdateCacheMut       sync.RWMutex
	routerUpdateCache          = make(map[string]updateCache)
	routerUpsertCacheMut       sync.RWMutex
	routerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var routerAfterSelectHooks []RouterHook

var routerBeforeInsertHooks []RouterHook
var routerAfterInsertHooks []RouterHook

var routerBeforeUpdateHooks []RouterHook
var routerAfterUpdateHooks []RouterHook

var routerBeforeDeleteHooks []RouterHook
var routerAfterDeleteHooks []RouterHook

var routerBeforeUpsertHooks []RouterHook
var routerAfterUpsertHooks []RouterHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Router) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Router) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Router) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Router) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Router) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Router) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Router) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Router) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Router) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRouterHook registers your hook function for all future operations.
func AddRouterHook(hookPoint boil.HookPoint, routerHook RouterHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		routerAfterSelectHooks = append(routerAfterSelectHooks, routerHook)
	case boil.BeforeInsertHook:
		routerBeforeInsertHooks = append(routerBeforeInsertHooks, routerHook)
	case boil.AfterInsertHook:
		routerAfterInsertHooks = append(routerAfterInsertHooks, routerHook)
	case boil.BeforeUpdateHook:
		routerBeforeUpdateHooks = append(routerBeforeUpdateHooks, routerHook)
	case boil.AfterUpdateHook:
		routerAfterUpdateHooks = append(routerAfterUpdateHooks, routerHook)
	case boil.BeforeDeleteHook:
		routerBeforeDeleteHooks = append(routerBeforeDeleteHooks, routerHook)
	case boil.AfterDeleteHook:
		routerAfterDeleteHooks = append(routerAfterDeleteHooks, routerHook)
	case boil.BeforeUpsertHook:
		routerBeforeUpsertHooks = append(routerBeforeUpsertHooks, routerHook)
	case boil.AfterUpsertHook:
		routerAfterUpsertHooks = append(routerAfterUpsertHooks, routerHook)
	}
}

// One returns a single router record from the query.
func (q routerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Router, error) {
	o := &Router{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for routers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Router records from the query.
func (q routerQuery) All(ctx context.Context, exec boil.ContextExecutor) (RouterSlice, error) {
	var o []*Router

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Router slice")
	}

	if len(routerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Router records in the query.
func (q routerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count routers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q routerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if routers exists")
	}

	return count > 0, nil
}

// Client pointed to by the foreign key.
func (o *Router) Client(mods ...qm.QueryMod) peerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ClientID),
	}

	queryMods = append(queryMods, mods...)

	return Peers(queryMods...)
}

// LoadClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (routerL) LoadClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRouter interface{}, mods queries.Applicator) error {
	var slice []*Router
	var object *Router

	if singular {
		var ok bool
		object, ok = maybeRouter.(*Router)
		if !ok {
			object = new(Router)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRouter)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRouter))
			}
		}
	} else {
		s, ok := maybeRouter.(*[]*Router)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRouter)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRouter))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &routerR{}
		}
		args = append(args, object.ClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &routerR{}
			}

			for _, a := range args {
				if a == obj.ClientID {
					continue Outer
				}
			}

			args = append(args, obj.ClientID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`peers`),
		qm.WhereIn(`peers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Peer")
	}

	var resultSlice []*Peer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Peer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for peers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for peers")
	}

	if len(routerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Client = foreign
		if foreign.R == nil {
			foreign.R = &peerR{}
		}
		foreign.R.ClientRouters = append(foreign.R.ClientRouters, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClientID == foreign.ID {
				local.R.Client = foreign
				if foreign.R == nil {
					foreign.R = &peerR{}
				}
				foreign.R.ClientRouters = append(foreign.R.ClientRouters, local)
				break
			}
		}
	}

	return nil
}

// SetClient of the router to the related item.
// Sets o.R.Client to related.
// Adds o to related.R.ClientRouters.
func (o *Router) SetClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Peer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"routers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"client_id"}),
		strmangle.WhereClause("\"", "\"", 2, routerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ClientID = related.ID
	if o.R == nil {
		o.R = &routerR{
			Client: related,
		}
	} else {
		o.R.Client = related
	}

	if related.R == nil {
		related.R = &peerR{
			ClientRouters: RouterSlice{o},
		}
	} else {
		related.R.ClientRouters = append(related.R.ClientRouters, o)
	}

	return nil
}

// Routers retrieves all the records using an executor.
func Routers(mods ...qm.QueryMod) routerQuery {
	mods = append(mods, qm.From("\"routers\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"routers\".*"})
	}

	return routerQuery{q}
}

// FindRouter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRouter(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Router, error) {
	routerObj := &Router{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"routers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, routerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from routers")
	}

	if err = routerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return routerObj, err
	}

	return routerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Router) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no routers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(routerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	routerInsertCacheMut.RLock()
	cache, cached := routerInsertCache[key]
	routerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			routerAllColumns,
			routerColumnsWithDefault,
			routerColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, routerGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(routerType, routerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(routerType, routerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"routers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"routers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into routers")
	}

	if !cached {
		routerInsertCacheMut.Lock()
		routerInsertCache[key] = cache
		routerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Router.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Router) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	routerUpdateCacheMut.RLock()
	cache, cached := routerUpdateCache[key]
	routerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			routerAllColumns,
			routerPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, routerGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update routers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"routers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, routerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(routerType, routerMapping, append(wl, routerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update routers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for routers")
	}

	if !cached {
		routerUpdateCacheMut.Lock()
		routerUpdateCache[key] = cache
		routerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q routerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for routers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for routers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RouterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"routers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, routerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in router slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all router")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Router) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no routers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(routerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	routerUpsertCacheMut.RLock()
	cache, cached := routerUpsertCache[key]
	routerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			routerAllColumns,
			routerColumnsWithDefault,
			routerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			routerAllColumns,
			routerPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, routerGeneratedColumns)
		update = strmangle.SetComplement(update, routerGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert routers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(routerPrimaryKeyColumns))
			copy(conflict, routerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"routers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(routerType, routerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(routerType, routerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert routers")
	}

	if !cached {
		routerUpsertCacheMut.Lock()
		routerUpsertCache[key] = cache
		routerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Router record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Router) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Router provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), routerPrimaryKeyMapping)
	sql := "DELETE FROM \"routers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from routers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for routers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q routerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no routerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from routers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for routers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RouterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(routerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"routers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, routerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from router slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for routers")
	}

	if len(routerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Router) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRouter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RouterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RouterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"routers\".* FROM \"routers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, routerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RouterSlice")
	}

	*o = slice

	return nil
}

// RouterExists checks if the Router row exists.
func RouterExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"routers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if routers exists")
	}

	return exists, nil
}
