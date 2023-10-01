// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

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

// AccountAdmin is an object representing the database table.
type AccountAdmin struct {
	AccountAdminID     string             `boil:"account_admin_id" json:"accountAdminID" toml:"accountAdminID" yaml:"accountAdminID"`
	Mail               string             `boil:"mail" json:"mail" toml:"mail" yaml:"mail"`
	Password           string             `boil:"password" json:"password" toml:"password" yaml:"password"`
	AccountAdminRole   AccountAdminRole   `boil:"account_admin_role" json:"accountAdminRole" toml:"accountAdminRole" yaml:"accountAdminRole"`
	AccountAdminStatus AccountAdminStatus `boil:"account_admin_status" json:"accountAdminStatus" toml:"accountAdminStatus" yaml:"accountAdminStatus"`
	CreatedAt          time.Time          `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt          time.Time          `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *accountAdminR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountAdminL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountAdminColumns = struct {
	AccountAdminID     string
	Mail               string
	Password           string
	AccountAdminRole   string
	AccountAdminStatus string
	CreatedAt          string
	UpdatedAt          string
}{
	AccountAdminID:     "account_admin_id",
	Mail:               "mail",
	Password:           "password",
	AccountAdminRole:   "account_admin_role",
	AccountAdminStatus: "account_admin_status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

var AccountAdminTableColumns = struct {
	AccountAdminID     string
	Mail               string
	Password           string
	AccountAdminRole   string
	AccountAdminStatus string
	CreatedAt          string
	UpdatedAt          string
}{
	AccountAdminID:     "account_admins.account_admin_id",
	Mail:               "account_admins.mail",
	Password:           "account_admins.password",
	AccountAdminRole:   "account_admins.account_admin_role",
	AccountAdminStatus: "account_admins.account_admin_status",
	CreatedAt:          "account_admins.created_at",
	UpdatedAt:          "account_admins.updated_at",
}

// Generated where

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

type whereHelperAccountAdminRole struct{ field string }

func (w whereHelperAccountAdminRole) EQ(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperAccountAdminRole) NEQ(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperAccountAdminRole) LT(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperAccountAdminRole) LTE(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperAccountAdminRole) GT(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperAccountAdminRole) GTE(x AccountAdminRole) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperAccountAdminRole) IN(slice []AccountAdminRole) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperAccountAdminRole) NIN(slice []AccountAdminRole) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperAccountAdminStatus struct{ field string }

func (w whereHelperAccountAdminStatus) EQ(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperAccountAdminStatus) NEQ(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperAccountAdminStatus) LT(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperAccountAdminStatus) LTE(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperAccountAdminStatus) GT(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperAccountAdminStatus) GTE(x AccountAdminStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperAccountAdminStatus) IN(slice []AccountAdminStatus) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperAccountAdminStatus) NIN(slice []AccountAdminStatus) qm.QueryMod {
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

var AccountAdminWhere = struct {
	AccountAdminID     whereHelperstring
	Mail               whereHelperstring
	Password           whereHelperstring
	AccountAdminRole   whereHelperAccountAdminRole
	AccountAdminStatus whereHelperAccountAdminStatus
	CreatedAt          whereHelpertime_Time
	UpdatedAt          whereHelpertime_Time
}{
	AccountAdminID:     whereHelperstring{field: "\"account_admins\".\"account_admin_id\""},
	Mail:               whereHelperstring{field: "\"account_admins\".\"mail\""},
	Password:           whereHelperstring{field: "\"account_admins\".\"password\""},
	AccountAdminRole:   whereHelperAccountAdminRole{field: "\"account_admins\".\"account_admin_role\""},
	AccountAdminStatus: whereHelperAccountAdminStatus{field: "\"account_admins\".\"account_admin_status\""},
	CreatedAt:          whereHelpertime_Time{field: "\"account_admins\".\"created_at\""},
	UpdatedAt:          whereHelpertime_Time{field: "\"account_admins\".\"updated_at\""},
}

// AccountAdminRels is where relationship names are stored.
var AccountAdminRels = struct {
}{}

// accountAdminR is where relationships are stored.
type accountAdminR struct {
}

// NewStruct creates a new relationship struct
func (*accountAdminR) NewStruct() *accountAdminR {
	return &accountAdminR{}
}

// accountAdminL is where Load methods for each relationship are stored.
type accountAdminL struct{}

var (
	accountAdminAllColumns            = []string{"account_admin_id", "mail", "password", "account_admin_role", "account_admin_status", "created_at", "updated_at"}
	accountAdminColumnsWithoutDefault = []string{"mail", "password", "account_admin_role", "account_admin_status"}
	accountAdminColumnsWithDefault    = []string{"account_admin_id", "created_at", "updated_at"}
	accountAdminPrimaryKeyColumns     = []string{"account_admin_id"}
	accountAdminGeneratedColumns      = []string{}
)

type (
	// AccountAdminSlice is an alias for a slice of pointers to AccountAdmin.
	// This should almost always be used instead of []AccountAdmin.
	AccountAdminSlice []*AccountAdmin
	// AccountAdminHook is the signature for custom AccountAdmin hook methods
	AccountAdminHook func(context.Context, boil.ContextExecutor, *AccountAdmin) error

	accountAdminQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountAdminType                 = reflect.TypeOf(&AccountAdmin{})
	accountAdminMapping              = queries.MakeStructMapping(accountAdminType)
	accountAdminPrimaryKeyMapping, _ = queries.BindMapping(accountAdminType, accountAdminMapping, accountAdminPrimaryKeyColumns)
	accountAdminInsertCacheMut       sync.RWMutex
	accountAdminInsertCache          = make(map[string]insertCache)
	accountAdminUpdateCacheMut       sync.RWMutex
	accountAdminUpdateCache          = make(map[string]updateCache)
	accountAdminUpsertCacheMut       sync.RWMutex
	accountAdminUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accountAdminAfterSelectHooks []AccountAdminHook

var accountAdminBeforeInsertHooks []AccountAdminHook
var accountAdminAfterInsertHooks []AccountAdminHook

var accountAdminBeforeUpdateHooks []AccountAdminHook
var accountAdminAfterUpdateHooks []AccountAdminHook

var accountAdminBeforeDeleteHooks []AccountAdminHook
var accountAdminAfterDeleteHooks []AccountAdminHook

var accountAdminBeforeUpsertHooks []AccountAdminHook
var accountAdminAfterUpsertHooks []AccountAdminHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AccountAdmin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AccountAdmin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AccountAdmin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AccountAdmin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AccountAdmin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AccountAdmin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AccountAdmin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AccountAdmin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AccountAdmin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAdminAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccountAdminHook registers your hook function for all future operations.
func AddAccountAdminHook(hookPoint boil.HookPoint, accountAdminHook AccountAdminHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		accountAdminAfterSelectHooks = append(accountAdminAfterSelectHooks, accountAdminHook)
	case boil.BeforeInsertHook:
		accountAdminBeforeInsertHooks = append(accountAdminBeforeInsertHooks, accountAdminHook)
	case boil.AfterInsertHook:
		accountAdminAfterInsertHooks = append(accountAdminAfterInsertHooks, accountAdminHook)
	case boil.BeforeUpdateHook:
		accountAdminBeforeUpdateHooks = append(accountAdminBeforeUpdateHooks, accountAdminHook)
	case boil.AfterUpdateHook:
		accountAdminAfterUpdateHooks = append(accountAdminAfterUpdateHooks, accountAdminHook)
	case boil.BeforeDeleteHook:
		accountAdminBeforeDeleteHooks = append(accountAdminBeforeDeleteHooks, accountAdminHook)
	case boil.AfterDeleteHook:
		accountAdminAfterDeleteHooks = append(accountAdminAfterDeleteHooks, accountAdminHook)
	case boil.BeforeUpsertHook:
		accountAdminBeforeUpsertHooks = append(accountAdminBeforeUpsertHooks, accountAdminHook)
	case boil.AfterUpsertHook:
		accountAdminAfterUpsertHooks = append(accountAdminAfterUpsertHooks, accountAdminHook)
	}
}

// OneG returns a single accountAdmin record from the query using the global executor.
func (q accountAdminQuery) OneG(ctx context.Context) (*AccountAdmin, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single accountAdmin record from the query.
func (q accountAdminQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccountAdmin, error) {
	o := &AccountAdmin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for account_admins")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all AccountAdmin records from the query using the global executor.
func (q accountAdminQuery) AllG(ctx context.Context) (AccountAdminSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AccountAdmin records from the query.
func (q accountAdminQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountAdminSlice, error) {
	var o []*AccountAdmin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to AccountAdmin slice")
	}

	if len(accountAdminAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all AccountAdmin records in the query using the global executor
func (q accountAdminQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AccountAdmin records in the query.
func (q accountAdminQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count account_admins rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q accountAdminQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q accountAdminQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if account_admins exists")
	}

	return count > 0, nil
}

// AccountAdmins retrieves all the records using an executor.
func AccountAdmins(mods ...qm.QueryMod) accountAdminQuery {
	mods = append(mods, qm.From("\"account_admins\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"account_admins\".*"})
	}

	return accountAdminQuery{q}
}

// FindAccountAdminG retrieves a single record by ID.
func FindAccountAdminG(ctx context.Context, accountAdminID string, selectCols ...string) (*AccountAdmin, error) {
	return FindAccountAdmin(ctx, boil.GetContextDB(), accountAdminID, selectCols...)
}

// FindAccountAdmin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccountAdmin(ctx context.Context, exec boil.ContextExecutor, accountAdminID string, selectCols ...string) (*AccountAdmin, error) {
	accountAdminObj := &AccountAdmin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"account_admins\" where \"account_admin_id\"=$1", sel,
	)

	q := queries.Raw(query, accountAdminID)

	err := q.Bind(ctx, exec, accountAdminObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from account_admins")
	}

	if err = accountAdminObj.doAfterSelectHooks(ctx, exec); err != nil {
		return accountAdminObj, err
	}

	return accountAdminObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AccountAdmin) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccountAdmin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no account_admins provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountAdminColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountAdminInsertCacheMut.RLock()
	cache, cached := accountAdminInsertCache[key]
	accountAdminInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountAdminAllColumns,
			accountAdminColumnsWithDefault,
			accountAdminColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountAdminType, accountAdminMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountAdminType, accountAdminMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"account_admins\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"account_admins\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into account_admins")
	}

	if !cached {
		accountAdminInsertCacheMut.Lock()
		accountAdminInsertCache[key] = cache
		accountAdminInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single AccountAdmin record using the global executor.
// See Update for more documentation.
func (o *AccountAdmin) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AccountAdmin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccountAdmin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accountAdminUpdateCacheMut.RLock()
	cache, cached := accountAdminUpdateCache[key]
	accountAdminUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountAdminAllColumns,
			accountAdminPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update account_admins, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"account_admins\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accountAdminPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountAdminType, accountAdminMapping, append(wl, accountAdminPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update account_admins row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for account_admins")
	}

	if !cached {
		accountAdminUpdateCacheMut.Lock()
		accountAdminUpdateCache[key] = cache
		accountAdminUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q accountAdminQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q accountAdminQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for account_admins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for account_admins")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AccountAdminSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountAdminSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountAdminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"account_admins\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accountAdminPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in accountAdmin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all accountAdmin")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AccountAdmin) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccountAdmin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no account_admins provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountAdminColumnsWithDefault, o)

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

	accountAdminUpsertCacheMut.RLock()
	cache, cached := accountAdminUpsertCache[key]
	accountAdminUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountAdminAllColumns,
			accountAdminColumnsWithDefault,
			accountAdminColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			accountAdminAllColumns,
			accountAdminPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert account_admins, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accountAdminPrimaryKeyColumns))
			copy(conflict, accountAdminPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"account_admins\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accountAdminType, accountAdminMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountAdminType, accountAdminMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert account_admins")
	}

	if !cached {
		accountAdminUpsertCacheMut.Lock()
		accountAdminUpsertCache[key] = cache
		accountAdminUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single AccountAdmin record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AccountAdmin) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AccountAdmin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccountAdmin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no AccountAdmin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountAdminPrimaryKeyMapping)
	sql := "DELETE FROM \"account_admins\" WHERE \"account_admin_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from account_admins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for account_admins")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q accountAdminQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q accountAdminQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no accountAdminQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from account_admins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for account_admins")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AccountAdminSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountAdminSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(accountAdminBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountAdminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"account_admins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountAdminPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from accountAdmin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for account_admins")
	}

	if len(accountAdminAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AccountAdmin) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: no AccountAdmin provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AccountAdmin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccountAdmin(ctx, exec, o.AccountAdminID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountAdminSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: empty AccountAdminSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountAdminSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountAdminSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountAdminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"account_admins\".* FROM \"account_admins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountAdminPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in AccountAdminSlice")
	}

	*o = slice

	return nil
}

// AccountAdminExistsG checks if the AccountAdmin row exists.
func AccountAdminExistsG(ctx context.Context, accountAdminID string) (bool, error) {
	return AccountAdminExists(ctx, boil.GetContextDB(), accountAdminID)
}

// AccountAdminExists checks if the AccountAdmin row exists.
func AccountAdminExists(ctx context.Context, exec boil.ContextExecutor, accountAdminID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"account_admins\" where \"account_admin_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, accountAdminID)
	}
	row := exec.QueryRowContext(ctx, sql, accountAdminID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if account_admins exists")
	}

	return exists, nil
}

// Exists checks if the AccountAdmin row exists.
func (o *AccountAdmin) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AccountAdminExists(ctx, exec, o.AccountAdminID)
}
