// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"emperror.dev/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ServerStatsMemberPeriod is an object representing the database table.
type ServerStatsMemberPeriod struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID    int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	NumMembers int64     `boil:"num_members" json:"num_members" toml:"num_members" yaml:"num_members"`
	Joins      int64     `boil:"joins" json:"joins" toml:"joins" yaml:"joins"`
	Leaves     int64     `boil:"leaves" json:"leaves" toml:"leaves" yaml:"leaves"`
	MaxOnline  int64     `boil:"max_online" json:"max_online" toml:"max_online" yaml:"max_online"`

	R *serverStatsMemberPeriodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serverStatsMemberPeriodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServerStatsMemberPeriodColumns = struct {
	ID         string
	GuildID    string
	CreatedAt  string
	NumMembers string
	Joins      string
	Leaves     string
	MaxOnline  string
}{
	ID:         "id",
	GuildID:    "guild_id",
	CreatedAt:  "created_at",
	NumMembers: "num_members",
	Joins:      "joins",
	Leaves:     "leaves",
	MaxOnline:  "max_online",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ServerStatsMemberPeriodWhere = struct {
	ID         whereHelperint64
	GuildID    whereHelperint64
	CreatedAt  whereHelpertime_Time
	NumMembers whereHelperint64
	Joins      whereHelperint64
	Leaves     whereHelperint64
	MaxOnline  whereHelperint64
}{
	ID:         whereHelperint64{field: "\"server_stats_member_periods\".\"id\""},
	GuildID:    whereHelperint64{field: "\"server_stats_member_periods\".\"guild_id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"server_stats_member_periods\".\"created_at\""},
	NumMembers: whereHelperint64{field: "\"server_stats_member_periods\".\"num_members\""},
	Joins:      whereHelperint64{field: "\"server_stats_member_periods\".\"joins\""},
	Leaves:     whereHelperint64{field: "\"server_stats_member_periods\".\"leaves\""},
	MaxOnline:  whereHelperint64{field: "\"server_stats_member_periods\".\"max_online\""},
}

// ServerStatsMemberPeriodRels is where relationship names are stored.
var ServerStatsMemberPeriodRels = struct {
}{}

// serverStatsMemberPeriodR is where relationships are stored.
type serverStatsMemberPeriodR struct {
}

// NewStruct creates a new relationship struct
func (*serverStatsMemberPeriodR) NewStruct() *serverStatsMemberPeriodR {
	return &serverStatsMemberPeriodR{}
}

// serverStatsMemberPeriodL is where Load methods for each relationship are stored.
type serverStatsMemberPeriodL struct{}

var (
	serverStatsMemberPeriodAllColumns            = []string{"id", "guild_id", "created_at", "num_members", "joins", "leaves", "max_online"}
	serverStatsMemberPeriodColumnsWithoutDefault = []string{"guild_id", "created_at", "num_members", "joins", "leaves", "max_online"}
	serverStatsMemberPeriodColumnsWithDefault    = []string{"id"}
	serverStatsMemberPeriodPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServerStatsMemberPeriodSlice is an alias for a slice of pointers to ServerStatsMemberPeriod.
	// This should generally be used opposed to []ServerStatsMemberPeriod.
	ServerStatsMemberPeriodSlice []*ServerStatsMemberPeriod

	serverStatsMemberPeriodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serverStatsMemberPeriodType                 = reflect.TypeOf(&ServerStatsMemberPeriod{})
	serverStatsMemberPeriodMapping              = queries.MakeStructMapping(serverStatsMemberPeriodType)
	serverStatsMemberPeriodPrimaryKeyMapping, _ = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, serverStatsMemberPeriodPrimaryKeyColumns)
	serverStatsMemberPeriodInsertCacheMut       sync.RWMutex
	serverStatsMemberPeriodInsertCache          = make(map[string]insertCache)
	serverStatsMemberPeriodUpdateCacheMut       sync.RWMutex
	serverStatsMemberPeriodUpdateCache          = make(map[string]updateCache)
	serverStatsMemberPeriodUpsertCacheMut       sync.RWMutex
	serverStatsMemberPeriodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single serverStatsMemberPeriod record from the query using the global executor.
func (q serverStatsMemberPeriodQuery) OneG(ctx context.Context) (*ServerStatsMemberPeriod, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single serverStatsMemberPeriod record from the query.
func (q serverStatsMemberPeriodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServerStatsMemberPeriod, error) {
	o := &ServerStatsMemberPeriod{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for server_stats_member_periods")
	}

	return o, nil
}

// AllG returns all ServerStatsMemberPeriod records from the query using the global executor.
func (q serverStatsMemberPeriodQuery) AllG(ctx context.Context) (ServerStatsMemberPeriodSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ServerStatsMemberPeriod records from the query.
func (q serverStatsMemberPeriodQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServerStatsMemberPeriodSlice, error) {
	var o []*ServerStatsMemberPeriod

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to ServerStatsMemberPeriod slice")
	}

	return o, nil
}

// CountG returns the count of all ServerStatsMemberPeriod records in the query, and panics on error.
func (q serverStatsMemberPeriodQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ServerStatsMemberPeriod records in the query.
func (q serverStatsMemberPeriodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count server_stats_member_periods rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q serverStatsMemberPeriodQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q serverStatsMemberPeriodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if server_stats_member_periods exists")
	}

	return count > 0, nil
}

// ServerStatsMemberPeriods retrieves all the records using an executor.
func ServerStatsMemberPeriods(mods ...qm.QueryMod) serverStatsMemberPeriodQuery {
	mods = append(mods, qm.From("\"server_stats_member_periods\""))
	return serverStatsMemberPeriodQuery{NewQuery(mods...)}
}

// FindServerStatsMemberPeriodG retrieves a single record by ID.
func FindServerStatsMemberPeriodG(ctx context.Context, iD int64, selectCols ...string) (*ServerStatsMemberPeriod, error) {
	return FindServerStatsMemberPeriod(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindServerStatsMemberPeriod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServerStatsMemberPeriod(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ServerStatsMemberPeriod, error) {
	serverStatsMemberPeriodObj := &ServerStatsMemberPeriod{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"server_stats_member_periods\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serverStatsMemberPeriodObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from server_stats_member_periods")
	}

	return serverStatsMemberPeriodObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ServerStatsMemberPeriod) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServerStatsMemberPeriod) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no server_stats_member_periods provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(serverStatsMemberPeriodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serverStatsMemberPeriodInsertCacheMut.RLock()
	cache, cached := serverStatsMemberPeriodInsertCache[key]
	serverStatsMemberPeriodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serverStatsMemberPeriodAllColumns,
			serverStatsMemberPeriodColumnsWithDefault,
			serverStatsMemberPeriodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"server_stats_member_periods\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"server_stats_member_periods\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.WrapIf(err, "models: unable to insert into server_stats_member_periods")
	}

	if !cached {
		serverStatsMemberPeriodInsertCacheMut.Lock()
		serverStatsMemberPeriodInsertCache[key] = cache
		serverStatsMemberPeriodInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ServerStatsMemberPeriod record using the global executor.
// See Update for more documentation.
func (o *ServerStatsMemberPeriod) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ServerStatsMemberPeriod.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServerStatsMemberPeriod) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	serverStatsMemberPeriodUpdateCacheMut.RLock()
	cache, cached := serverStatsMemberPeriodUpdateCache[key]
	serverStatsMemberPeriodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serverStatsMemberPeriodAllColumns,
			serverStatsMemberPeriodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update server_stats_member_periods, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"server_stats_member_periods\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serverStatsMemberPeriodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, append(wl, serverStatsMemberPeriodPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update server_stats_member_periods row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for server_stats_member_periods")
	}

	if !cached {
		serverStatsMemberPeriodUpdateCacheMut.Lock()
		serverStatsMemberPeriodUpdateCache[key] = cache
		serverStatsMemberPeriodUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q serverStatsMemberPeriodQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q serverStatsMemberPeriodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for server_stats_member_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for server_stats_member_periods")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ServerStatsMemberPeriodSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServerStatsMemberPeriodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsMemberPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"server_stats_member_periods\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serverStatsMemberPeriodPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in serverStatsMemberPeriod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all serverStatsMemberPeriod")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ServerStatsMemberPeriod) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServerStatsMemberPeriod) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no server_stats_member_periods provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(serverStatsMemberPeriodColumnsWithDefault, o)

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

	serverStatsMemberPeriodUpsertCacheMut.RLock()
	cache, cached := serverStatsMemberPeriodUpsertCache[key]
	serverStatsMemberPeriodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serverStatsMemberPeriodAllColumns,
			serverStatsMemberPeriodColumnsWithDefault,
			serverStatsMemberPeriodColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serverStatsMemberPeriodAllColumns,
			serverStatsMemberPeriodPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert server_stats_member_periods, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serverStatsMemberPeriodPrimaryKeyColumns))
			copy(conflict, serverStatsMemberPeriodPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"server_stats_member_periods\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serverStatsMemberPeriodType, serverStatsMemberPeriodMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.WrapIf(err, "models: unable to upsert server_stats_member_periods")
	}

	if !cached {
		serverStatsMemberPeriodUpsertCacheMut.Lock()
		serverStatsMemberPeriodUpsertCache[key] = cache
		serverStatsMemberPeriodUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ServerStatsMemberPeriod record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ServerStatsMemberPeriod) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ServerStatsMemberPeriod record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServerStatsMemberPeriod) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServerStatsMemberPeriod provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serverStatsMemberPeriodPrimaryKeyMapping)
	sql := "DELETE FROM \"server_stats_member_periods\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from server_stats_member_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for server_stats_member_periods")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serverStatsMemberPeriodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serverStatsMemberPeriodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from server_stats_member_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for server_stats_member_periods")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ServerStatsMemberPeriodSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServerStatsMemberPeriodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsMemberPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"server_stats_member_periods\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsMemberPeriodPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from serverStatsMemberPeriod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for server_stats_member_periods")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ServerStatsMemberPeriod) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ServerStatsMemberPeriod provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ServerStatsMemberPeriod) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServerStatsMemberPeriod(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsMemberPeriodSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ServerStatsMemberPeriodSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsMemberPeriodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServerStatsMemberPeriodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsMemberPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"server_stats_member_periods\".* FROM \"server_stats_member_periods\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsMemberPeriodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in ServerStatsMemberPeriodSlice")
	}

	*o = slice

	return nil
}

// ServerStatsMemberPeriodExistsG checks if the ServerStatsMemberPeriod row exists.
func ServerStatsMemberPeriodExistsG(ctx context.Context, iD int64) (bool, error) {
	return ServerStatsMemberPeriodExists(ctx, boil.GetContextDB(), iD)
}

// ServerStatsMemberPeriodExists checks if the ServerStatsMemberPeriod row exists.
func ServerStatsMemberPeriodExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"server_stats_member_periods\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if server_stats_member_periods exists")
	}

	return exists, nil
}
