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

// UserTimezone is an object representing the database table.
type UserTimezone struct {
	UserID       int64  `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TimezoneName string `boil:"timezone_name" json:"timezone_name" toml:"timezone_name" yaml:"timezone_name"`

	R *userTimezoneR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userTimezoneL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserTimezoneColumns = struct {
	UserID       string
	TimezoneName string
}{
	UserID:       "user_id",
	TimezoneName: "timezone_name",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var UserTimezoneWhere = struct {
	UserID       whereHelperint64
	TimezoneName whereHelperstring
}{
	UserID:       whereHelperint64{field: "\"user_timezones\".\"user_id\""},
	TimezoneName: whereHelperstring{field: "\"user_timezones\".\"timezone_name\""},
}

// UserTimezoneRels is where relationship names are stored.
var UserTimezoneRels = struct {
}{}

// userTimezoneR is where relationships are stored.
type userTimezoneR struct {
}

// NewStruct creates a new relationship struct
func (*userTimezoneR) NewStruct() *userTimezoneR {
	return &userTimezoneR{}
}

// userTimezoneL is where Load methods for each relationship are stored.
type userTimezoneL struct{}

var (
	userTimezoneAllColumns            = []string{"user_id", "timezone_name"}
	userTimezoneColumnsWithoutDefault = []string{"user_id", "timezone_name"}
	userTimezoneColumnsWithDefault    = []string{}
	userTimezonePrimaryKeyColumns     = []string{"user_id"}
)

type (
	// UserTimezoneSlice is an alias for a slice of pointers to UserTimezone.
	// This should generally be used opposed to []UserTimezone.
	UserTimezoneSlice []*UserTimezone

	userTimezoneQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userTimezoneType                 = reflect.TypeOf(&UserTimezone{})
	userTimezoneMapping              = queries.MakeStructMapping(userTimezoneType)
	userTimezonePrimaryKeyMapping, _ = queries.BindMapping(userTimezoneType, userTimezoneMapping, userTimezonePrimaryKeyColumns)
	userTimezoneInsertCacheMut       sync.RWMutex
	userTimezoneInsertCache          = make(map[string]insertCache)
	userTimezoneUpdateCacheMut       sync.RWMutex
	userTimezoneUpdateCache          = make(map[string]updateCache)
	userTimezoneUpsertCacheMut       sync.RWMutex
	userTimezoneUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single userTimezone record from the query using the global executor.
func (q userTimezoneQuery) OneG(ctx context.Context) (*UserTimezone, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single userTimezone record from the query.
func (q userTimezoneQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserTimezone, error) {
	o := &UserTimezone{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for user_timezones")
	}

	return o, nil
}

// AllG returns all UserTimezone records from the query using the global executor.
func (q userTimezoneQuery) AllG(ctx context.Context) (UserTimezoneSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all UserTimezone records from the query.
func (q userTimezoneQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserTimezoneSlice, error) {
	var o []*UserTimezone

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to UserTimezone slice")
	}

	return o, nil
}

// CountG returns the count of all UserTimezone records in the query, and panics on error.
func (q userTimezoneQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all UserTimezone records in the query.
func (q userTimezoneQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count user_timezones rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userTimezoneQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q userTimezoneQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if user_timezones exists")
	}

	return count > 0, nil
}

// UserTimezones retrieves all the records using an executor.
func UserTimezones(mods ...qm.QueryMod) userTimezoneQuery {
	mods = append(mods, qm.From("\"user_timezones\""))
	return userTimezoneQuery{NewQuery(mods...)}
}

// FindUserTimezoneG retrieves a single record by ID.
func FindUserTimezoneG(ctx context.Context, userID int64, selectCols ...string) (*UserTimezone, error) {
	return FindUserTimezone(ctx, boil.GetContextDB(), userID, selectCols...)
}

// FindUserTimezone retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserTimezone(ctx context.Context, exec boil.ContextExecutor, userID int64, selectCols ...string) (*UserTimezone, error) {
	userTimezoneObj := &UserTimezone{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_timezones\" where \"user_id\"=$1", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, userTimezoneObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from user_timezones")
	}

	return userTimezoneObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserTimezone) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserTimezone) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_timezones provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(userTimezoneColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userTimezoneInsertCacheMut.RLock()
	cache, cached := userTimezoneInsertCache[key]
	userTimezoneInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userTimezoneAllColumns,
			userTimezoneColumnsWithDefault,
			userTimezoneColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userTimezoneType, userTimezoneMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userTimezoneType, userTimezoneMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_timezones\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_timezones\" %sDEFAULT VALUES%s"
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
		return errors.WrapIf(err, "models: unable to insert into user_timezones")
	}

	if !cached {
		userTimezoneInsertCacheMut.Lock()
		userTimezoneInsertCache[key] = cache
		userTimezoneInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single UserTimezone record using the global executor.
// See Update for more documentation.
func (o *UserTimezone) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the UserTimezone.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserTimezone) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	userTimezoneUpdateCacheMut.RLock()
	cache, cached := userTimezoneUpdateCache[key]
	userTimezoneUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userTimezoneAllColumns,
			userTimezonePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_timezones, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_timezones\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userTimezonePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userTimezoneType, userTimezoneMapping, append(wl, userTimezonePrimaryKeyColumns...))
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
		return 0, errors.WrapIf(err, "models: unable to update user_timezones row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for user_timezones")
	}

	if !cached {
		userTimezoneUpdateCacheMut.Lock()
		userTimezoneUpdateCache[key] = cache
		userTimezoneUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q userTimezoneQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userTimezoneQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for user_timezones")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for user_timezones")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserTimezoneSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserTimezoneSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTimezonePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_timezones\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userTimezonePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in userTimezone slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all userTimezone")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserTimezone) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserTimezone) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_timezones provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(userTimezoneColumnsWithDefault, o)

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

	userTimezoneUpsertCacheMut.RLock()
	cache, cached := userTimezoneUpsertCache[key]
	userTimezoneUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userTimezoneAllColumns,
			userTimezoneColumnsWithDefault,
			userTimezoneColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userTimezoneAllColumns,
			userTimezonePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_timezones, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userTimezonePrimaryKeyColumns))
			copy(conflict, userTimezonePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_timezones\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userTimezoneType, userTimezoneMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userTimezoneType, userTimezoneMapping, ret)
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
		return errors.WrapIf(err, "models: unable to upsert user_timezones")
	}

	if !cached {
		userTimezoneUpsertCacheMut.Lock()
		userTimezoneUpsertCache[key] = cache
		userTimezoneUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single UserTimezone record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserTimezone) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single UserTimezone record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserTimezone) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserTimezone provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userTimezonePrimaryKeyMapping)
	sql := "DELETE FROM \"user_timezones\" WHERE \"user_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from user_timezones")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for user_timezones")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userTimezoneQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userTimezoneQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from user_timezones")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for user_timezones")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserTimezoneSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserTimezoneSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTimezonePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_timezones\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userTimezonePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from userTimezone slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for user_timezones")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserTimezone) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no UserTimezone provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserTimezone) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserTimezone(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserTimezoneSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty UserTimezoneSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserTimezoneSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserTimezoneSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTimezonePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_timezones\".* FROM \"user_timezones\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userTimezonePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in UserTimezoneSlice")
	}

	*o = slice

	return nil
}

// UserTimezoneExistsG checks if the UserTimezone row exists.
func UserTimezoneExistsG(ctx context.Context, userID int64) (bool, error) {
	return UserTimezoneExists(ctx, boil.GetContextDB(), userID)
}

// UserTimezoneExists checks if the UserTimezone row exists.
func UserTimezoneExists(ctx context.Context, exec boil.ContextExecutor, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_timezones\" where \"user_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID)
	}

	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if user_timezones exists")
	}

	return exists, nil
}
