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

// AccountOtp is an object representing the database table.
type AccountOtp struct {
	AccountID string    `boil:"account_id" json:"accountID" toml:"accountID" yaml:"accountID"`
	Otp       string    `boil:"otp" json:"otp" toml:"otp" yaml:"otp"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *accountOtpR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountOtpL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountOtpColumns = struct {
	AccountID string
	Otp       string
	CreatedAt string
	UpdatedAt string
}{
	AccountID: "account_id",
	Otp:       "otp",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var AccountOtpTableColumns = struct {
	AccountID string
	Otp       string
	CreatedAt string
	UpdatedAt string
}{
	AccountID: "account_otps.account_id",
	Otp:       "account_otps.otp",
	CreatedAt: "account_otps.created_at",
	UpdatedAt: "account_otps.updated_at",
}

// Generated where

var AccountOtpWhere = struct {
	AccountID whereHelperstring
	Otp       whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	AccountID: whereHelperstring{field: "\"account_otps\".\"account_id\""},
	Otp:       whereHelperstring{field: "\"account_otps\".\"otp\""},
	CreatedAt: whereHelpertime_Time{field: "\"account_otps\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"account_otps\".\"updated_at\""},
}

// AccountOtpRels is where relationship names are stored.
var AccountOtpRels = struct {
	Account string
}{
	Account: "Account",
}

// accountOtpR is where relationships are stored.
type accountOtpR struct {
	Account *Account `boil:"Account" json:"Account" toml:"Account" yaml:"Account"`
}

// NewStruct creates a new relationship struct
func (*accountOtpR) NewStruct() *accountOtpR {
	return &accountOtpR{}
}

func (r *accountOtpR) GetAccount() *Account {
	if r == nil {
		return nil
	}
	return r.Account
}

// accountOtpL is where Load methods for each relationship are stored.
type accountOtpL struct{}

var (
	accountOtpAllColumns            = []string{"account_id", "otp", "created_at", "updated_at"}
	accountOtpColumnsWithoutDefault = []string{"account_id", "otp"}
	accountOtpColumnsWithDefault    = []string{"created_at", "updated_at"}
	accountOtpPrimaryKeyColumns     = []string{"account_id"}
	accountOtpGeneratedColumns      = []string{}
)

type (
	// AccountOtpSlice is an alias for a slice of pointers to AccountOtp.
	// This should almost always be used instead of []AccountOtp.
	AccountOtpSlice []*AccountOtp
	// AccountOtpHook is the signature for custom AccountOtp hook methods
	AccountOtpHook func(context.Context, boil.ContextExecutor, *AccountOtp) error

	accountOtpQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountOtpType                 = reflect.TypeOf(&AccountOtp{})
	accountOtpMapping              = queries.MakeStructMapping(accountOtpType)
	accountOtpPrimaryKeyMapping, _ = queries.BindMapping(accountOtpType, accountOtpMapping, accountOtpPrimaryKeyColumns)
	accountOtpInsertCacheMut       sync.RWMutex
	accountOtpInsertCache          = make(map[string]insertCache)
	accountOtpUpdateCacheMut       sync.RWMutex
	accountOtpUpdateCache          = make(map[string]updateCache)
	accountOtpUpsertCacheMut       sync.RWMutex
	accountOtpUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accountOtpAfterSelectHooks []AccountOtpHook

var accountOtpBeforeInsertHooks []AccountOtpHook
var accountOtpAfterInsertHooks []AccountOtpHook

var accountOtpBeforeUpdateHooks []AccountOtpHook
var accountOtpAfterUpdateHooks []AccountOtpHook

var accountOtpBeforeDeleteHooks []AccountOtpHook
var accountOtpAfterDeleteHooks []AccountOtpHook

var accountOtpBeforeUpsertHooks []AccountOtpHook
var accountOtpAfterUpsertHooks []AccountOtpHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AccountOtp) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AccountOtp) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AccountOtp) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AccountOtp) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AccountOtp) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AccountOtp) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AccountOtp) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AccountOtp) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AccountOtp) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountOtpAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccountOtpHook registers your hook function for all future operations.
func AddAccountOtpHook(hookPoint boil.HookPoint, accountOtpHook AccountOtpHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		accountOtpAfterSelectHooks = append(accountOtpAfterSelectHooks, accountOtpHook)
	case boil.BeforeInsertHook:
		accountOtpBeforeInsertHooks = append(accountOtpBeforeInsertHooks, accountOtpHook)
	case boil.AfterInsertHook:
		accountOtpAfterInsertHooks = append(accountOtpAfterInsertHooks, accountOtpHook)
	case boil.BeforeUpdateHook:
		accountOtpBeforeUpdateHooks = append(accountOtpBeforeUpdateHooks, accountOtpHook)
	case boil.AfterUpdateHook:
		accountOtpAfterUpdateHooks = append(accountOtpAfterUpdateHooks, accountOtpHook)
	case boil.BeforeDeleteHook:
		accountOtpBeforeDeleteHooks = append(accountOtpBeforeDeleteHooks, accountOtpHook)
	case boil.AfterDeleteHook:
		accountOtpAfterDeleteHooks = append(accountOtpAfterDeleteHooks, accountOtpHook)
	case boil.BeforeUpsertHook:
		accountOtpBeforeUpsertHooks = append(accountOtpBeforeUpsertHooks, accountOtpHook)
	case boil.AfterUpsertHook:
		accountOtpAfterUpsertHooks = append(accountOtpAfterUpsertHooks, accountOtpHook)
	}
}

// OneG returns a single accountOtp record from the query using the global executor.
func (q accountOtpQuery) OneG(ctx context.Context) (*AccountOtp, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single accountOtp record from the query.
func (q accountOtpQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccountOtp, error) {
	o := &AccountOtp{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for account_otps")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all AccountOtp records from the query using the global executor.
func (q accountOtpQuery) AllG(ctx context.Context) (AccountOtpSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AccountOtp records from the query.
func (q accountOtpQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountOtpSlice, error) {
	var o []*AccountOtp

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to AccountOtp slice")
	}

	if len(accountOtpAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all AccountOtp records in the query using the global executor
func (q accountOtpQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AccountOtp records in the query.
func (q accountOtpQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count account_otps rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q accountOtpQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q accountOtpQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if account_otps exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *AccountOtp) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"account_id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	return Accounts(queryMods...)
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (accountOtpL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccountOtp interface{}, mods queries.Applicator) error {
	var slice []*AccountOtp
	var object *AccountOtp

	if singular {
		var ok bool
		object, ok = maybeAccountOtp.(*AccountOtp)
		if !ok {
			object = new(AccountOtp)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAccountOtp)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAccountOtp))
			}
		}
	} else {
		s, ok := maybeAccountOtp.(*[]*AccountOtp)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAccountOtp)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAccountOtp))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountOtpR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountOtpR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`accounts`),
		qm.WhereIn(`accounts.account_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if len(accountAfterSelectHooks) != 0 {
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
		object.R.Account = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.AccountID {
				local.R.Account = foreign
				break
			}
		}
	}

	return nil
}

// SetAccountG of the accountOtp to the related item.
// Sets o.R.Account to related.
// Uses the global database handle.
func (o *AccountOtp) SetAccountG(ctx context.Context, insert bool, related *Account) error {
	return o.SetAccount(ctx, boil.GetContextDB(), insert, related)
}

// SetAccount of the accountOtp to the related item.
// Sets o.R.Account to related.
func (o *AccountOtp) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"account_otps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 2, accountOtpPrimaryKeyColumns),
	)
	values := []interface{}{related.AccountID, o.AccountID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.AccountID
	if o.R == nil {
		o.R = &accountOtpR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	return nil
}

// AccountOtps retrieves all the records using an executor.
func AccountOtps(mods ...qm.QueryMod) accountOtpQuery {
	mods = append(mods, qm.From("\"account_otps\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"account_otps\".*"})
	}

	return accountOtpQuery{q}
}

// FindAccountOtpG retrieves a single record by ID.
func FindAccountOtpG(ctx context.Context, accountID string, selectCols ...string) (*AccountOtp, error) {
	return FindAccountOtp(ctx, boil.GetContextDB(), accountID, selectCols...)
}

// FindAccountOtp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccountOtp(ctx context.Context, exec boil.ContextExecutor, accountID string, selectCols ...string) (*AccountOtp, error) {
	accountOtpObj := &AccountOtp{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"account_otps\" where \"account_id\"=$1", sel,
	)

	q := queries.Raw(query, accountID)

	err := q.Bind(ctx, exec, accountOtpObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from account_otps")
	}

	if err = accountOtpObj.doAfterSelectHooks(ctx, exec); err != nil {
		return accountOtpObj, err
	}

	return accountOtpObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AccountOtp) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccountOtp) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no account_otps provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(accountOtpColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountOtpInsertCacheMut.RLock()
	cache, cached := accountOtpInsertCache[key]
	accountOtpInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountOtpAllColumns,
			accountOtpColumnsWithDefault,
			accountOtpColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountOtpType, accountOtpMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountOtpType, accountOtpMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"account_otps\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"account_otps\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into account_otps")
	}

	if !cached {
		accountOtpInsertCacheMut.Lock()
		accountOtpInsertCache[key] = cache
		accountOtpInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single AccountOtp record using the global executor.
// See Update for more documentation.
func (o *AccountOtp) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AccountOtp.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccountOtp) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accountOtpUpdateCacheMut.RLock()
	cache, cached := accountOtpUpdateCache[key]
	accountOtpUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountOtpAllColumns,
			accountOtpPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update account_otps, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"account_otps\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accountOtpPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountOtpType, accountOtpMapping, append(wl, accountOtpPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update account_otps row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for account_otps")
	}

	if !cached {
		accountOtpUpdateCacheMut.Lock()
		accountOtpUpdateCache[key] = cache
		accountOtpUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q accountOtpQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q accountOtpQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for account_otps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for account_otps")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AccountOtpSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountOtpSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountOtpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"account_otps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accountOtpPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in accountOtp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all accountOtp")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AccountOtp) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccountOtp) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no account_otps provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(accountOtpColumnsWithDefault, o)

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

	accountOtpUpsertCacheMut.RLock()
	cache, cached := accountOtpUpsertCache[key]
	accountOtpUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountOtpAllColumns,
			accountOtpColumnsWithDefault,
			accountOtpColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			accountOtpAllColumns,
			accountOtpPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert account_otps, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accountOtpPrimaryKeyColumns))
			copy(conflict, accountOtpPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"account_otps\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accountOtpType, accountOtpMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountOtpType, accountOtpMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert account_otps")
	}

	if !cached {
		accountOtpUpsertCacheMut.Lock()
		accountOtpUpsertCache[key] = cache
		accountOtpUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single AccountOtp record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AccountOtp) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AccountOtp record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccountOtp) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no AccountOtp provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountOtpPrimaryKeyMapping)
	sql := "DELETE FROM \"account_otps\" WHERE \"account_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from account_otps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for account_otps")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q accountOtpQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q accountOtpQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no accountOtpQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from account_otps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for account_otps")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AccountOtpSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountOtpSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(accountOtpBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountOtpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"account_otps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountOtpPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from accountOtp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for account_otps")
	}

	if len(accountOtpAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AccountOtp) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: no AccountOtp provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AccountOtp) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccountOtp(ctx, exec, o.AccountID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountOtpSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: empty AccountOtpSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountOtpSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountOtpSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountOtpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"account_otps\".* FROM \"account_otps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountOtpPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in AccountOtpSlice")
	}

	*o = slice

	return nil
}

// AccountOtpExistsG checks if the AccountOtp row exists.
func AccountOtpExistsG(ctx context.Context, accountID string) (bool, error) {
	return AccountOtpExists(ctx, boil.GetContextDB(), accountID)
}

// AccountOtpExists checks if the AccountOtp row exists.
func AccountOtpExists(ctx context.Context, exec boil.ContextExecutor, accountID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"account_otps\" where \"account_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, accountID)
	}
	row := exec.QueryRowContext(ctx, sql, accountID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if account_otps exists")
	}

	return exists, nil
}

// Exists checks if the AccountOtp row exists.
func (o *AccountOtp) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AccountOtpExists(ctx, exec, o.AccountID)
}
