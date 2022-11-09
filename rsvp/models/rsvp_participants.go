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

// RSVPParticipant is an object representing the database table.
type RSVPParticipant struct {
	UserID                  int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RSVPSessionsMessageID   int64     `boil:"rsvp_sessions_message_id" json:"rsvp_sessions_message_id" toml:"rsvp_sessions_message_id" yaml:"rsvp_sessions_message_id"`
	GuildID                 int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	JoinState               int16     `boil:"join_state" json:"join_state" toml:"join_state" yaml:"join_state"`
	ReminderEnabled         bool      `boil:"reminder_enabled" json:"reminder_enabled" toml:"reminder_enabled" yaml:"reminder_enabled"`
	MarkedAsParticipatingAt time.Time `boil:"marked_as_participating_at" json:"marked_as_participating_at" toml:"marked_as_participating_at" yaml:"marked_as_participating_at"`

	R *rsvpParticipantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rsvpParticipantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RSVPParticipantColumns = struct {
	UserID                  string
	RSVPSessionsMessageID   string
	GuildID                 string
	JoinState               string
	ReminderEnabled         string
	MarkedAsParticipatingAt string
}{
	UserID:                  "user_id",
	RSVPSessionsMessageID:   "rsvp_sessions_message_id",
	GuildID:                 "guild_id",
	JoinState:               "join_state",
	ReminderEnabled:         "reminder_enabled",
	MarkedAsParticipatingAt: "marked_as_participating_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

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

var RSVPParticipantWhere = struct {
	UserID                  whereHelperint64
	RSVPSessionsMessageID   whereHelperint64
	GuildID                 whereHelperint64
	JoinState               whereHelperint16
	ReminderEnabled         whereHelperbool
	MarkedAsParticipatingAt whereHelpertime_Time
}{
	UserID:                  whereHelperint64{field: "\"rsvp_participants\".\"user_id\""},
	RSVPSessionsMessageID:   whereHelperint64{field: "\"rsvp_participants\".\"rsvp_sessions_message_id\""},
	GuildID:                 whereHelperint64{field: "\"rsvp_participants\".\"guild_id\""},
	JoinState:               whereHelperint16{field: "\"rsvp_participants\".\"join_state\""},
	ReminderEnabled:         whereHelperbool{field: "\"rsvp_participants\".\"reminder_enabled\""},
	MarkedAsParticipatingAt: whereHelpertime_Time{field: "\"rsvp_participants\".\"marked_as_participating_at\""},
}

// RSVPParticipantRels is where relationship names are stored.
var RSVPParticipantRels = struct {
	RSVPSessionsMessage string
}{
	RSVPSessionsMessage: "RSVPSessionsMessage",
}

// rsvpParticipantR is where relationships are stored.
type rsvpParticipantR struct {
	RSVPSessionsMessage *RSVPSession
}

// NewStruct creates a new relationship struct
func (*rsvpParticipantR) NewStruct() *rsvpParticipantR {
	return &rsvpParticipantR{}
}

// rsvpParticipantL is where Load methods for each relationship are stored.
type rsvpParticipantL struct{}

var (
	rsvpParticipantAllColumns            = []string{"user_id", "rsvp_sessions_message_id", "guild_id", "join_state", "reminder_enabled", "marked_as_participating_at"}
	rsvpParticipantColumnsWithoutDefault = []string{"user_id", "rsvp_sessions_message_id", "guild_id", "join_state", "reminder_enabled", "marked_as_participating_at"}
	rsvpParticipantColumnsWithDefault    = []string{}
	rsvpParticipantPrimaryKeyColumns     = []string{"rsvp_sessions_message_id", "user_id"}
)

type (
	// RSVPParticipantSlice is an alias for a slice of pointers to RSVPParticipant.
	// This should generally be used opposed to []RSVPParticipant.
	RSVPParticipantSlice []*RSVPParticipant

	rsvpParticipantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rsvpParticipantType                 = reflect.TypeOf(&RSVPParticipant{})
	rsvpParticipantMapping              = queries.MakeStructMapping(rsvpParticipantType)
	rsvpParticipantPrimaryKeyMapping, _ = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, rsvpParticipantPrimaryKeyColumns)
	rsvpParticipantInsertCacheMut       sync.RWMutex
	rsvpParticipantInsertCache          = make(map[string]insertCache)
	rsvpParticipantUpdateCacheMut       sync.RWMutex
	rsvpParticipantUpdateCache          = make(map[string]updateCache)
	rsvpParticipantUpsertCacheMut       sync.RWMutex
	rsvpParticipantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single rsvpParticipant record from the query using the global executor.
func (q rsvpParticipantQuery) OneG(ctx context.Context) (*RSVPParticipant, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single rsvpParticipant record from the query.
func (q rsvpParticipantQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RSVPParticipant, error) {
	o := &RSVPParticipant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for rsvp_participants")
	}

	return o, nil
}

// AllG returns all RSVPParticipant records from the query using the global executor.
func (q rsvpParticipantQuery) AllG(ctx context.Context) (RSVPParticipantSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all RSVPParticipant records from the query.
func (q rsvpParticipantQuery) All(ctx context.Context, exec boil.ContextExecutor) (RSVPParticipantSlice, error) {
	var o []*RSVPParticipant

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to RSVPParticipant slice")
	}

	return o, nil
}

// CountG returns the count of all RSVPParticipant records in the query, and panics on error.
func (q rsvpParticipantQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all RSVPParticipant records in the query.
func (q rsvpParticipantQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count rsvp_participants rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q rsvpParticipantQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q rsvpParticipantQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if rsvp_participants exists")
	}

	return count > 0, nil
}

// RSVPSessionsMessage pointed to by the foreign key.
func (o *RSVPParticipant) RSVPSessionsMessage(mods ...qm.QueryMod) rsvpSessionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("message_id=?", o.RSVPSessionsMessageID),
	}

	queryMods = append(queryMods, mods...)

	query := RSVPSessions(queryMods...)
	queries.SetFrom(query.Query, "\"rsvp_sessions\"")

	return query
}

// LoadRSVPSessionsMessage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (rsvpParticipantL) LoadRSVPSessionsMessage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRSVPParticipant interface{}, mods queries.Applicator) error {
	var slice []*RSVPParticipant
	var object *RSVPParticipant

	if singular {
		object = maybeRSVPParticipant.(*RSVPParticipant)
	} else {
		slice = *maybeRSVPParticipant.(*[]*RSVPParticipant)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rsvpParticipantR{}
		}
		args = append(args, object.RSVPSessionsMessageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rsvpParticipantR{}
			}

			for _, a := range args {
				if a == obj.RSVPSessionsMessageID {
					continue Outer
				}
			}

			args = append(args, obj.RSVPSessionsMessageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`rsvp_sessions`), qm.WhereIn(`message_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.WrapIf(err, "failed to eager load RSVPSession")
	}

	var resultSlice []*RSVPSession
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.WrapIf(err, "failed to bind eager loaded slice RSVPSession")
	}

	if err = results.Close(); err != nil {
		return errors.WrapIf(err, "failed to close results of eager load for rsvp_sessions")
	}
	if err = results.Err(); err != nil {
		return errors.WrapIf(err, "error occurred during iteration of eager loaded relations for rsvp_sessions")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.RSVPSessionsMessage = foreign
		if foreign.R == nil {
			foreign.R = &rsvpSessionR{}
		}
		foreign.R.RSVPSessionsMessageRSVPParticipants = append(foreign.R.RSVPSessionsMessageRSVPParticipants, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RSVPSessionsMessageID == foreign.MessageID {
				local.R.RSVPSessionsMessage = foreign
				if foreign.R == nil {
					foreign.R = &rsvpSessionR{}
				}
				foreign.R.RSVPSessionsMessageRSVPParticipants = append(foreign.R.RSVPSessionsMessageRSVPParticipants, local)
				break
			}
		}
	}

	return nil
}

// SetRSVPSessionsMessageG of the rsvpParticipant to the related item.
// Sets o.R.RSVPSessionsMessage to related.
// Adds o to related.R.RSVPSessionsMessageRSVPParticipants.
// Uses the global database handle.
func (o *RSVPParticipant) SetRSVPSessionsMessageG(ctx context.Context, insert bool, related *RSVPSession) error {
	return o.SetRSVPSessionsMessage(ctx, boil.GetContextDB(), insert, related)
}

// SetRSVPSessionsMessage of the rsvpParticipant to the related item.
// Sets o.R.RSVPSessionsMessage to related.
// Adds o to related.R.RSVPSessionsMessageRSVPParticipants.
func (o *RSVPParticipant) SetRSVPSessionsMessage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RSVPSession) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.WrapIf(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"rsvp_participants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"rsvp_sessions_message_id"}),
		strmangle.WhereClause("\"", "\"", 2, rsvpParticipantPrimaryKeyColumns),
	)
	values := []interface{}{related.MessageID, o.RSVPSessionsMessageID, o.UserID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.WrapIf(err, "failed to update local table")
	}

	o.RSVPSessionsMessageID = related.MessageID
	if o.R == nil {
		o.R = &rsvpParticipantR{
			RSVPSessionsMessage: related,
		}
	} else {
		o.R.RSVPSessionsMessage = related
	}

	if related.R == nil {
		related.R = &rsvpSessionR{
			RSVPSessionsMessageRSVPParticipants: RSVPParticipantSlice{o},
		}
	} else {
		related.R.RSVPSessionsMessageRSVPParticipants = append(related.R.RSVPSessionsMessageRSVPParticipants, o)
	}

	return nil
}

// RSVPParticipants retrieves all the records using an executor.
func RSVPParticipants(mods ...qm.QueryMod) rsvpParticipantQuery {
	mods = append(mods, qm.From("\"rsvp_participants\""))
	return rsvpParticipantQuery{NewQuery(mods...)}
}

// FindRSVPParticipantG retrieves a single record by ID.
func FindRSVPParticipantG(ctx context.Context, rSVPSessionsMessageID int64, userID int64, selectCols ...string) (*RSVPParticipant, error) {
	return FindRSVPParticipant(ctx, boil.GetContextDB(), rSVPSessionsMessageID, userID, selectCols...)
}

// FindRSVPParticipant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRSVPParticipant(ctx context.Context, exec boil.ContextExecutor, rSVPSessionsMessageID int64, userID int64, selectCols ...string) (*RSVPParticipant, error) {
	rsvpParticipantObj := &RSVPParticipant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rsvp_participants\" where \"rsvp_sessions_message_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, rSVPSessionsMessageID, userID)

	err := q.Bind(ctx, exec, rsvpParticipantObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from rsvp_participants")
	}

	return rsvpParticipantObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *RSVPParticipant) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RSVPParticipant) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rsvp_participants provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(rsvpParticipantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rsvpParticipantInsertCacheMut.RLock()
	cache, cached := rsvpParticipantInsertCache[key]
	rsvpParticipantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rsvpParticipantAllColumns,
			rsvpParticipantColumnsWithDefault,
			rsvpParticipantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rsvp_participants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rsvp_participants\" %sDEFAULT VALUES%s"
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
		return errors.WrapIf(err, "models: unable to insert into rsvp_participants")
	}

	if !cached {
		rsvpParticipantInsertCacheMut.Lock()
		rsvpParticipantInsertCache[key] = cache
		rsvpParticipantInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single RSVPParticipant record using the global executor.
// See Update for more documentation.
func (o *RSVPParticipant) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the RSVPParticipant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RSVPParticipant) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	rsvpParticipantUpdateCacheMut.RLock()
	cache, cached := rsvpParticipantUpdateCache[key]
	rsvpParticipantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rsvpParticipantAllColumns,
			rsvpParticipantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rsvp_participants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rsvp_participants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rsvpParticipantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, append(wl, rsvpParticipantPrimaryKeyColumns...))
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
		return 0, errors.WrapIf(err, "models: unable to update rsvp_participants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for rsvp_participants")
	}

	if !cached {
		rsvpParticipantUpdateCacheMut.Lock()
		rsvpParticipantUpdateCache[key] = cache
		rsvpParticipantUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q rsvpParticipantQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q rsvpParticipantQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for rsvp_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for rsvp_participants")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RSVPParticipantSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RSVPParticipantSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rsvp_participants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rsvpParticipantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in rsvpParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all rsvpParticipant")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *RSVPParticipant) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RSVPParticipant) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rsvp_participants provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(rsvpParticipantColumnsWithDefault, o)

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

	rsvpParticipantUpsertCacheMut.RLock()
	cache, cached := rsvpParticipantUpsertCache[key]
	rsvpParticipantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rsvpParticipantAllColumns,
			rsvpParticipantColumnsWithDefault,
			rsvpParticipantColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rsvpParticipantAllColumns,
			rsvpParticipantPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rsvp_participants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rsvpParticipantPrimaryKeyColumns))
			copy(conflict, rsvpParticipantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rsvp_participants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rsvpParticipantType, rsvpParticipantMapping, ret)
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
		return errors.WrapIf(err, "models: unable to upsert rsvp_participants")
	}

	if !cached {
		rsvpParticipantUpsertCacheMut.Lock()
		rsvpParticipantUpsertCache[key] = cache
		rsvpParticipantUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single RSVPParticipant record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *RSVPParticipant) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single RSVPParticipant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RSVPParticipant) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RSVPParticipant provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rsvpParticipantPrimaryKeyMapping)
	sql := "DELETE FROM \"rsvp_participants\" WHERE \"rsvp_sessions_message_id\"=$1 AND \"user_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from rsvp_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for rsvp_participants")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rsvpParticipantQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rsvpParticipantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from rsvp_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for rsvp_participants")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RSVPParticipantSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RSVPParticipantSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rsvp_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rsvpParticipantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from rsvpParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for rsvp_participants")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *RSVPParticipant) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no RSVPParticipant provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RSVPParticipant) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRSVPParticipant(ctx, exec, o.RSVPSessionsMessageID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RSVPParticipantSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty RSVPParticipantSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RSVPParticipantSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RSVPParticipantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rsvp_participants\".* FROM \"rsvp_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rsvpParticipantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in RSVPParticipantSlice")
	}

	*o = slice

	return nil
}

// RSVPParticipantExistsG checks if the RSVPParticipant row exists.
func RSVPParticipantExistsG(ctx context.Context, rSVPSessionsMessageID int64, userID int64) (bool, error) {
	return RSVPParticipantExists(ctx, boil.GetContextDB(), rSVPSessionsMessageID, userID)
}

// RSVPParticipantExists checks if the RSVPParticipant row exists.
func RSVPParticipantExists(ctx context.Context, exec boil.ContextExecutor, rSVPSessionsMessageID int64, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rsvp_participants\" where \"rsvp_sessions_message_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, rSVPSessionsMessageID, userID)
	}

	row := exec.QueryRowContext(ctx, sql, rSVPSessionsMessageID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if rsvp_participants exists")
	}

	return exists, nil
}
