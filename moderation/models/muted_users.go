// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// MutedUser is an object representing the database table.
type MutedUser struct {
	ID           int              `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt    time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time        `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ExpiresAt    null.Time        `boil:"expires_at" json:"expires_at,omitempty" toml:"expires_at" yaml:"expires_at,omitempty"`
	GuildID      int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	UserID       int64            `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AuthorID     int64            `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	Reason       string           `boil:"reason" json:"reason" toml:"reason" yaml:"reason"`
	RemovedRoles types.Int64Array `boil:"removed_roles" json:"removed_roles,omitempty" toml:"removed_roles" yaml:"removed_roles,omitempty"`

	R *mutedUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mutedUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MutedUserColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	ExpiresAt    string
	GuildID      string
	UserID       string
	AuthorID     string
	Reason       string
	RemovedRoles string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ExpiresAt:    "expires_at",
	GuildID:      "guild_id",
	UserID:       "user_id",
	AuthorID:     "author_id",
	Reason:       "reason",
	RemovedRoles: "removed_roles",
}

var MutedUserTableColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	ExpiresAt    string
	GuildID      string
	UserID       string
	AuthorID     string
	Reason       string
	RemovedRoles string
}{
	ID:           "muted_users.id",
	CreatedAt:    "muted_users.created_at",
	UpdatedAt:    "muted_users.updated_at",
	ExpiresAt:    "muted_users.expires_at",
	GuildID:      "muted_users.guild_id",
	UserID:       "muted_users.user_id",
	AuthorID:     "muted_users.author_id",
	Reason:       "muted_users.reason",
	RemovedRoles: "muted_users.removed_roles",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var MutedUserWhere = struct {
	ID           whereHelperint
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
	ExpiresAt    whereHelpernull_Time
	GuildID      whereHelperint64
	UserID       whereHelperint64
	AuthorID     whereHelperint64
	Reason       whereHelperstring
	RemovedRoles whereHelpertypes_Int64Array
}{
	ID:           whereHelperint{field: "\"muted_users\".\"id\""},
	CreatedAt:    whereHelpertime_Time{field: "\"muted_users\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"muted_users\".\"updated_at\""},
	ExpiresAt:    whereHelpernull_Time{field: "\"muted_users\".\"expires_at\""},
	GuildID:      whereHelperint64{field: "\"muted_users\".\"guild_id\""},
	UserID:       whereHelperint64{field: "\"muted_users\".\"user_id\""},
	AuthorID:     whereHelperint64{field: "\"muted_users\".\"author_id\""},
	Reason:       whereHelperstring{field: "\"muted_users\".\"reason\""},
	RemovedRoles: whereHelpertypes_Int64Array{field: "\"muted_users\".\"removed_roles\""},
}

// MutedUserRels is where relationship names are stored.
var MutedUserRels = struct {
}{}

// mutedUserR is where relationships are stored.
type mutedUserR struct {
}

// NewStruct creates a new relationship struct
func (*mutedUserR) NewStruct() *mutedUserR {
	return &mutedUserR{}
}

// mutedUserL is where Load methods for each relationship are stored.
type mutedUserL struct{}

var (
	mutedUserAllColumns            = []string{"id", "created_at", "updated_at", "expires_at", "guild_id", "user_id", "author_id", "reason", "removed_roles"}
	mutedUserColumnsWithoutDefault = []string{"created_at", "updated_at", "guild_id", "user_id", "author_id", "reason"}
	mutedUserColumnsWithDefault    = []string{"id", "expires_at", "removed_roles"}
	mutedUserPrimaryKeyColumns     = []string{"id"}
	mutedUserGeneratedColumns      = []string{}
)

type (
	// MutedUserSlice is an alias for a slice of pointers to MutedUser.
	// This should almost always be used instead of []MutedUser.
	MutedUserSlice []*MutedUser

	mutedUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mutedUserType                 = reflect.TypeOf(&MutedUser{})
	mutedUserMapping              = queries.MakeStructMapping(mutedUserType)
	mutedUserPrimaryKeyMapping, _ = queries.BindMapping(mutedUserType, mutedUserMapping, mutedUserPrimaryKeyColumns)
	mutedUserInsertCacheMut       sync.RWMutex
	mutedUserInsertCache          = make(map[string]insertCache)
	mutedUserUpdateCacheMut       sync.RWMutex
	mutedUserUpdateCache          = make(map[string]updateCache)
	mutedUserUpsertCacheMut       sync.RWMutex
	mutedUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single mutedUser record from the query using the global executor.
func (q mutedUserQuery) OneG(ctx context.Context) (*MutedUser, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single mutedUser record from the query.
func (q mutedUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MutedUser, error) {
	o := &MutedUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for muted_users")
	}

	return o, nil
}

// AllG returns all MutedUser records from the query using the global executor.
func (q mutedUserQuery) AllG(ctx context.Context) (MutedUserSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all MutedUser records from the query.
func (q mutedUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (MutedUserSlice, error) {
	var o []*MutedUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MutedUser slice")
	}

	return o, nil
}

// CountG returns the count of all MutedUser records in the query using the global executor
func (q mutedUserQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all MutedUser records in the query.
func (q mutedUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count muted_users rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q mutedUserQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q mutedUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if muted_users exists")
	}

	return count > 0, nil
}

// MutedUsers retrieves all the records using an executor.
func MutedUsers(mods ...qm.QueryMod) mutedUserQuery {
	mods = append(mods, qm.From("\"muted_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"muted_users\".*"})
	}

	return mutedUserQuery{q}
}

// FindMutedUserG retrieves a single record by ID.
func FindMutedUserG(ctx context.Context, iD int, selectCols ...string) (*MutedUser, error) {
	return FindMutedUser(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindMutedUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMutedUser(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*MutedUser, error) {
	mutedUserObj := &MutedUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"muted_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mutedUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from muted_users")
	}

	return mutedUserObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *MutedUser) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MutedUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no muted_users provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(mutedUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mutedUserInsertCacheMut.RLock()
	cache, cached := mutedUserInsertCache[key]
	mutedUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mutedUserAllColumns,
			mutedUserColumnsWithDefault,
			mutedUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mutedUserType, mutedUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mutedUserType, mutedUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"muted_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"muted_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into muted_users")
	}

	if !cached {
		mutedUserInsertCacheMut.Lock()
		mutedUserInsertCache[key] = cache
		mutedUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single MutedUser record using the global executor.
// See Update for more documentation.
func (o *MutedUser) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the MutedUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MutedUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	mutedUserUpdateCacheMut.RLock()
	cache, cached := mutedUserUpdateCache[key]
	mutedUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mutedUserAllColumns,
			mutedUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update muted_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"muted_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mutedUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mutedUserType, mutedUserMapping, append(wl, mutedUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update muted_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for muted_users")
	}

	if !cached {
		mutedUserUpdateCacheMut.Lock()
		mutedUserUpdateCache[key] = cache
		mutedUserUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q mutedUserQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q mutedUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for muted_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for muted_users")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o MutedUserSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MutedUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mutedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"muted_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mutedUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mutedUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mutedUser")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *MutedUser) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MutedUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no muted_users provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(mutedUserColumnsWithDefault, o)

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

	mutedUserUpsertCacheMut.RLock()
	cache, cached := mutedUserUpsertCache[key]
	mutedUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			mutedUserAllColumns,
			mutedUserColumnsWithDefault,
			mutedUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mutedUserAllColumns,
			mutedUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert muted_users, could not build update column list")
		}

		ret := strmangle.SetComplement(mutedUserAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(mutedUserPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert muted_users, could not build conflict column list")
			}

			conflict = make([]string, len(mutedUserPrimaryKeyColumns))
			copy(conflict, mutedUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"muted_users\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(mutedUserType, mutedUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mutedUserType, mutedUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert muted_users")
	}

	if !cached {
		mutedUserUpsertCacheMut.Lock()
		mutedUserUpsertCache[key] = cache
		mutedUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single MutedUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *MutedUser) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single MutedUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MutedUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MutedUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mutedUserPrimaryKeyMapping)
	sql := "DELETE FROM \"muted_users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from muted_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for muted_users")
	}

	return rowsAff, nil
}

func (q mutedUserQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q mutedUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mutedUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from muted_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for muted_users")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o MutedUserSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MutedUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mutedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"muted_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mutedUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mutedUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for muted_users")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *MutedUser) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no MutedUser provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MutedUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMutedUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MutedUserSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty MutedUserSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MutedUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MutedUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mutedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"muted_users\".* FROM \"muted_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mutedUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MutedUserSlice")
	}

	*o = slice

	return nil
}

// MutedUserExistsG checks if the MutedUser row exists.
func MutedUserExistsG(ctx context.Context, iD int) (bool, error) {
	return MutedUserExists(ctx, boil.GetContextDB(), iD)
}

// MutedUserExists checks if the MutedUser row exists.
func MutedUserExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"muted_users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if muted_users exists")
	}

	return exists, nil
}

// Exists checks if the MutedUser row exists.
func (o *MutedUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MutedUserExists(ctx, exec, o.ID)
}
