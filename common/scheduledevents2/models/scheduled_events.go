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

// ScheduledEvent is an object representing the database table.
type ScheduledEvent struct {
	ID           int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	TriggersAt   time.Time   `boil:"triggers_at" json:"triggers_at" toml:"triggers_at" yaml:"triggers_at"`
	RetryOnError bool        `boil:"retry_on_error" json:"retry_on_error" toml:"retry_on_error" yaml:"retry_on_error"`
	GuildID      int64       `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	EventName    string      `boil:"event_name" json:"event_name" toml:"event_name" yaml:"event_name"`
	Data         types.JSON  `boil:"data" json:"data" toml:"data" yaml:"data"`
	Processed    bool        `boil:"processed" json:"processed" toml:"processed" yaml:"processed"`
	Error        null.String `boil:"error" json:"error,omitempty" toml:"error" yaml:"error,omitempty"`

	R *scheduledEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scheduledEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ScheduledEventColumns = struct {
	ID           string
	CreatedAt    string
	TriggersAt   string
	RetryOnError string
	GuildID      string
	EventName    string
	Data         string
	Processed    string
	Error        string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	TriggersAt:   "triggers_at",
	RetryOnError: "retry_on_error",
	GuildID:      "guild_id",
	EventName:    "event_name",
	Data:         "data",
	Processed:    "processed",
	Error:        "error",
}

var ScheduledEventTableColumns = struct {
	ID           string
	CreatedAt    string
	TriggersAt   string
	RetryOnError string
	GuildID      string
	EventName    string
	Data         string
	Processed    string
	Error        string
}{
	ID:           "scheduled_events.id",
	CreatedAt:    "scheduled_events.created_at",
	TriggersAt:   "scheduled_events.triggers_at",
	RetryOnError: "scheduled_events.retry_on_error",
	GuildID:      "scheduled_events.guild_id",
	EventName:    "scheduled_events.event_name",
	Data:         "scheduled_events.data",
	Processed:    "scheduled_events.processed",
	Error:        "scheduled_events.error",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

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

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ScheduledEventWhere = struct {
	ID           whereHelperint64
	CreatedAt    whereHelpertime_Time
	TriggersAt   whereHelpertime_Time
	RetryOnError whereHelperbool
	GuildID      whereHelperint64
	EventName    whereHelperstring
	Data         whereHelpertypes_JSON
	Processed    whereHelperbool
	Error        whereHelpernull_String
}{
	ID:           whereHelperint64{field: "\"scheduled_events\".\"id\""},
	CreatedAt:    whereHelpertime_Time{field: "\"scheduled_events\".\"created_at\""},
	TriggersAt:   whereHelpertime_Time{field: "\"scheduled_events\".\"triggers_at\""},
	RetryOnError: whereHelperbool{field: "\"scheduled_events\".\"retry_on_error\""},
	GuildID:      whereHelperint64{field: "\"scheduled_events\".\"guild_id\""},
	EventName:    whereHelperstring{field: "\"scheduled_events\".\"event_name\""},
	Data:         whereHelpertypes_JSON{field: "\"scheduled_events\".\"data\""},
	Processed:    whereHelperbool{field: "\"scheduled_events\".\"processed\""},
	Error:        whereHelpernull_String{field: "\"scheduled_events\".\"error\""},
}

// ScheduledEventRels is where relationship names are stored.
var ScheduledEventRels = struct {
}{}

// scheduledEventR is where relationships are stored.
type scheduledEventR struct {
}

// NewStruct creates a new relationship struct
func (*scheduledEventR) NewStruct() *scheduledEventR {
	return &scheduledEventR{}
}

// scheduledEventL is where Load methods for each relationship are stored.
type scheduledEventL struct{}

var (
	scheduledEventAllColumns            = []string{"id", "created_at", "triggers_at", "retry_on_error", "guild_id", "event_name", "data", "processed", "error"}
	scheduledEventColumnsWithoutDefault = []string{"created_at", "triggers_at", "retry_on_error", "guild_id", "event_name", "data", "processed"}
	scheduledEventColumnsWithDefault    = []string{"id", "error"}
	scheduledEventPrimaryKeyColumns     = []string{"id"}
	scheduledEventGeneratedColumns      = []string{}
)

type (
	// ScheduledEventSlice is an alias for a slice of pointers to ScheduledEvent.
	// This should almost always be used instead of []ScheduledEvent.
	ScheduledEventSlice []*ScheduledEvent

	scheduledEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	scheduledEventType                 = reflect.TypeOf(&ScheduledEvent{})
	scheduledEventMapping              = queries.MakeStructMapping(scheduledEventType)
	scheduledEventPrimaryKeyMapping, _ = queries.BindMapping(scheduledEventType, scheduledEventMapping, scheduledEventPrimaryKeyColumns)
	scheduledEventInsertCacheMut       sync.RWMutex
	scheduledEventInsertCache          = make(map[string]insertCache)
	scheduledEventUpdateCacheMut       sync.RWMutex
	scheduledEventUpdateCache          = make(map[string]updateCache)
	scheduledEventUpsertCacheMut       sync.RWMutex
	scheduledEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single scheduledEvent record from the query using the global executor.
func (q scheduledEventQuery) OneG(ctx context.Context) (*ScheduledEvent, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single scheduledEvent record from the query.
func (q scheduledEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ScheduledEvent, error) {
	o := &ScheduledEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for scheduled_events")
	}

	return o, nil
}

// AllG returns all ScheduledEvent records from the query using the global executor.
func (q scheduledEventQuery) AllG(ctx context.Context) (ScheduledEventSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ScheduledEvent records from the query.
func (q scheduledEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (ScheduledEventSlice, error) {
	var o []*ScheduledEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ScheduledEvent slice")
	}

	return o, nil
}

// CountG returns the count of all ScheduledEvent records in the query using the global executor
func (q scheduledEventQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ScheduledEvent records in the query.
func (q scheduledEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count scheduled_events rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q scheduledEventQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q scheduledEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if scheduled_events exists")
	}

	return count > 0, nil
}

// ScheduledEvents retrieves all the records using an executor.
func ScheduledEvents(mods ...qm.QueryMod) scheduledEventQuery {
	mods = append(mods, qm.From("\"scheduled_events\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"scheduled_events\".*"})
	}

	return scheduledEventQuery{q}
}

// FindScheduledEventG retrieves a single record by ID.
func FindScheduledEventG(ctx context.Context, iD int64, selectCols ...string) (*ScheduledEvent, error) {
	return FindScheduledEvent(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindScheduledEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindScheduledEvent(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ScheduledEvent, error) {
	scheduledEventObj := &ScheduledEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"scheduled_events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, scheduledEventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from scheduled_events")
	}

	return scheduledEventObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ScheduledEvent) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ScheduledEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no scheduled_events provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(scheduledEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	scheduledEventInsertCacheMut.RLock()
	cache, cached := scheduledEventInsertCache[key]
	scheduledEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			scheduledEventAllColumns,
			scheduledEventColumnsWithDefault,
			scheduledEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(scheduledEventType, scheduledEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(scheduledEventType, scheduledEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"scheduled_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"scheduled_events\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into scheduled_events")
	}

	if !cached {
		scheduledEventInsertCacheMut.Lock()
		scheduledEventInsertCache[key] = cache
		scheduledEventInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ScheduledEvent record using the global executor.
// See Update for more documentation.
func (o *ScheduledEvent) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ScheduledEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ScheduledEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	scheduledEventUpdateCacheMut.RLock()
	cache, cached := scheduledEventUpdateCache[key]
	scheduledEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			scheduledEventAllColumns,
			scheduledEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update scheduled_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"scheduled_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, scheduledEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(scheduledEventType, scheduledEventMapping, append(wl, scheduledEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update scheduled_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for scheduled_events")
	}

	if !cached {
		scheduledEventUpdateCacheMut.Lock()
		scheduledEventUpdateCache[key] = cache
		scheduledEventUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q scheduledEventQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q scheduledEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for scheduled_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for scheduled_events")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ScheduledEventSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ScheduledEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"scheduled_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, scheduledEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in scheduledEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all scheduledEvent")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ScheduledEvent) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ScheduledEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no scheduled_events provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(scheduledEventColumnsWithDefault, o)

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

	scheduledEventUpsertCacheMut.RLock()
	cache, cached := scheduledEventUpsertCache[key]
	scheduledEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			scheduledEventAllColumns,
			scheduledEventColumnsWithDefault,
			scheduledEventColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			scheduledEventAllColumns,
			scheduledEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert scheduled_events, could not build update column list")
		}

		ret := strmangle.SetComplement(scheduledEventAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(scheduledEventPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert scheduled_events, could not build conflict column list")
			}

			conflict = make([]string, len(scheduledEventPrimaryKeyColumns))
			copy(conflict, scheduledEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"scheduled_events\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(scheduledEventType, scheduledEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(scheduledEventType, scheduledEventMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert scheduled_events")
	}

	if !cached {
		scheduledEventUpsertCacheMut.Lock()
		scheduledEventUpsertCache[key] = cache
		scheduledEventUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ScheduledEvent record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ScheduledEvent) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ScheduledEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ScheduledEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ScheduledEvent provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), scheduledEventPrimaryKeyMapping)
	sql := "DELETE FROM \"scheduled_events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from scheduled_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for scheduled_events")
	}

	return rowsAff, nil
}

func (q scheduledEventQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q scheduledEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no scheduledEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from scheduled_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for scheduled_events")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ScheduledEventSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ScheduledEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"scheduled_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scheduledEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from scheduledEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for scheduled_events")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ScheduledEvent) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ScheduledEvent provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ScheduledEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindScheduledEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScheduledEventSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ScheduledEventSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScheduledEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ScheduledEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"scheduled_events\".* FROM \"scheduled_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scheduledEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ScheduledEventSlice")
	}

	*o = slice

	return nil
}

// ScheduledEventExistsG checks if the ScheduledEvent row exists.
func ScheduledEventExistsG(ctx context.Context, iD int64) (bool, error) {
	return ScheduledEventExists(ctx, boil.GetContextDB(), iD)
}

// ScheduledEventExists checks if the ScheduledEvent row exists.
func ScheduledEventExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"scheduled_events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if scheduled_events exists")
	}

	return exists, nil
}

// Exists checks if the ScheduledEvent row exists.
func (o *ScheduledEvent) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ScheduledEventExists(ctx, exec, o.ID)
}
