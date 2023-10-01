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

// Material is an object representing the database table.
type Material struct {
	MaterialID string    `boil:"material_id" json:"materialID" toml:"materialID" yaml:"materialID"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt  time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *materialR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L materialL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MaterialColumns = struct {
	MaterialID string
	Name       string
	CreatedAt  string
	UpdatedAt  string
}{
	MaterialID: "material_id",
	Name:       "name",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var MaterialTableColumns = struct {
	MaterialID string
	Name       string
	CreatedAt  string
	UpdatedAt  string
}{
	MaterialID: "materials.material_id",
	Name:       "materials.name",
	CreatedAt:  "materials.created_at",
	UpdatedAt:  "materials.updated_at",
}

// Generated where

var MaterialWhere = struct {
	MaterialID whereHelperstring
	Name       whereHelperstring
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	MaterialID: whereHelperstring{field: "\"materials\".\"material_id\""},
	Name:       whereHelperstring{field: "\"materials\".\"name\""},
	CreatedAt:  whereHelpertime_Time{field: "\"materials\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"materials\".\"updated_at\""},
}

// MaterialRels is where relationship names are stored.
var MaterialRels = struct {
	Jewelleries string
}{
	Jewelleries: "Jewelleries",
}

// materialR is where relationships are stored.
type materialR struct {
	Jewelleries JewellerySlice `boil:"Jewelleries" json:"Jewelleries" toml:"Jewelleries" yaml:"Jewelleries"`
}

// NewStruct creates a new relationship struct
func (*materialR) NewStruct() *materialR {
	return &materialR{}
}

func (r *materialR) GetJewelleries() JewellerySlice {
	if r == nil {
		return nil
	}
	return r.Jewelleries
}

// materialL is where Load methods for each relationship are stored.
type materialL struct{}

var (
	materialAllColumns            = []string{"material_id", "name", "created_at", "updated_at"}
	materialColumnsWithoutDefault = []string{"name"}
	materialColumnsWithDefault    = []string{"material_id", "created_at", "updated_at"}
	materialPrimaryKeyColumns     = []string{"material_id"}
	materialGeneratedColumns      = []string{}
)

type (
	// MaterialSlice is an alias for a slice of pointers to Material.
	// This should almost always be used instead of []Material.
	MaterialSlice []*Material
	// MaterialHook is the signature for custom Material hook methods
	MaterialHook func(context.Context, boil.ContextExecutor, *Material) error

	materialQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	materialType                 = reflect.TypeOf(&Material{})
	materialMapping              = queries.MakeStructMapping(materialType)
	materialPrimaryKeyMapping, _ = queries.BindMapping(materialType, materialMapping, materialPrimaryKeyColumns)
	materialInsertCacheMut       sync.RWMutex
	materialInsertCache          = make(map[string]insertCache)
	materialUpdateCacheMut       sync.RWMutex
	materialUpdateCache          = make(map[string]updateCache)
	materialUpsertCacheMut       sync.RWMutex
	materialUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var materialAfterSelectHooks []MaterialHook

var materialBeforeInsertHooks []MaterialHook
var materialAfterInsertHooks []MaterialHook

var materialBeforeUpdateHooks []MaterialHook
var materialAfterUpdateHooks []MaterialHook

var materialBeforeDeleteHooks []MaterialHook
var materialAfterDeleteHooks []MaterialHook

var materialBeforeUpsertHooks []MaterialHook
var materialAfterUpsertHooks []MaterialHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Material) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Material) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Material) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Material) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Material) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Material) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Material) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Material) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Material) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range materialAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMaterialHook registers your hook function for all future operations.
func AddMaterialHook(hookPoint boil.HookPoint, materialHook MaterialHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		materialAfterSelectHooks = append(materialAfterSelectHooks, materialHook)
	case boil.BeforeInsertHook:
		materialBeforeInsertHooks = append(materialBeforeInsertHooks, materialHook)
	case boil.AfterInsertHook:
		materialAfterInsertHooks = append(materialAfterInsertHooks, materialHook)
	case boil.BeforeUpdateHook:
		materialBeforeUpdateHooks = append(materialBeforeUpdateHooks, materialHook)
	case boil.AfterUpdateHook:
		materialAfterUpdateHooks = append(materialAfterUpdateHooks, materialHook)
	case boil.BeforeDeleteHook:
		materialBeforeDeleteHooks = append(materialBeforeDeleteHooks, materialHook)
	case boil.AfterDeleteHook:
		materialAfterDeleteHooks = append(materialAfterDeleteHooks, materialHook)
	case boil.BeforeUpsertHook:
		materialBeforeUpsertHooks = append(materialBeforeUpsertHooks, materialHook)
	case boil.AfterUpsertHook:
		materialAfterUpsertHooks = append(materialAfterUpsertHooks, materialHook)
	}
}

// OneG returns a single material record from the query using the global executor.
func (q materialQuery) OneG(ctx context.Context) (*Material, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single material record from the query.
func (q materialQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Material, error) {
	o := &Material{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for materials")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Material records from the query using the global executor.
func (q materialQuery) AllG(ctx context.Context) (MaterialSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Material records from the query.
func (q materialQuery) All(ctx context.Context, exec boil.ContextExecutor) (MaterialSlice, error) {
	var o []*Material

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Material slice")
	}

	if len(materialAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Material records in the query using the global executor
func (q materialQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Material records in the query.
func (q materialQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count materials rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q materialQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q materialQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if materials exists")
	}

	return count > 0, nil
}

// Jewelleries retrieves all the jewellery's Jewelleries with an executor.
func (o *Material) Jewelleries(mods ...qm.QueryMod) jewelleryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"jewelleries\".\"material_id\"=?", o.MaterialID),
	)

	return Jewelleries(queryMods...)
}

// LoadJewelleries allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (materialL) LoadJewelleries(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMaterial interface{}, mods queries.Applicator) error {
	var slice []*Material
	var object *Material

	if singular {
		var ok bool
		object, ok = maybeMaterial.(*Material)
		if !ok {
			object = new(Material)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMaterial)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMaterial))
			}
		}
	} else {
		s, ok := maybeMaterial.(*[]*Material)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMaterial)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMaterial))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &materialR{}
		}
		args = append(args, object.MaterialID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &materialR{}
			}

			for _, a := range args {
				if a == obj.MaterialID {
					continue Outer
				}
			}

			args = append(args, obj.MaterialID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`jewelleries`),
		qm.WhereIn(`jewelleries.material_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load jewelleries")
	}

	var resultSlice []*Jewellery
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice jewelleries")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on jewelleries")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for jewelleries")
	}

	if len(jewelleryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Jewelleries = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.MaterialID == foreign.MaterialID {
				local.R.Jewelleries = append(local.R.Jewelleries, foreign)
				break
			}
		}
	}

	return nil
}

// AddJewelleriesG adds the given related objects to the existing relationships
// of the material, optionally inserting them as new records.
// Appends related to o.R.Jewelleries.
// Uses the global database handle.
func (o *Material) AddJewelleriesG(ctx context.Context, insert bool, related ...*Jewellery) error {
	return o.AddJewelleries(ctx, boil.GetContextDB(), insert, related...)
}

// AddJewelleries adds the given related objects to the existing relationships
// of the material, optionally inserting them as new records.
// Appends related to o.R.Jewelleries.
func (o *Material) AddJewelleries(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Jewellery) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MaterialID = o.MaterialID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"jewelleries\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"material_id"}),
				strmangle.WhereClause("\"", "\"", 2, jewelleryPrimaryKeyColumns),
			)
			values := []interface{}{o.MaterialID, rel.JewelleryID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MaterialID = o.MaterialID
		}
	}

	if o.R == nil {
		o.R = &materialR{
			Jewelleries: related,
		}
	} else {
		o.R.Jewelleries = append(o.R.Jewelleries, related...)
	}

	return nil
}

// Materials retrieves all the records using an executor.
func Materials(mods ...qm.QueryMod) materialQuery {
	mods = append(mods, qm.From("\"materials\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"materials\".*"})
	}

	return materialQuery{q}
}

// FindMaterialG retrieves a single record by ID.
func FindMaterialG(ctx context.Context, materialID string, selectCols ...string) (*Material, error) {
	return FindMaterial(ctx, boil.GetContextDB(), materialID, selectCols...)
}

// FindMaterial retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMaterial(ctx context.Context, exec boil.ContextExecutor, materialID string, selectCols ...string) (*Material, error) {
	materialObj := &Material{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"materials\" where \"material_id\"=$1", sel,
	)

	q := queries.Raw(query, materialID)

	err := q.Bind(ctx, exec, materialObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from materials")
	}

	if err = materialObj.doAfterSelectHooks(ctx, exec); err != nil {
		return materialObj, err
	}

	return materialObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Material) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Material) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no materials provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(materialColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	materialInsertCacheMut.RLock()
	cache, cached := materialInsertCache[key]
	materialInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			materialAllColumns,
			materialColumnsWithDefault,
			materialColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(materialType, materialMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(materialType, materialMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"materials\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"materials\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into materials")
	}

	if !cached {
		materialInsertCacheMut.Lock()
		materialInsertCache[key] = cache
		materialInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Material record using the global executor.
// See Update for more documentation.
func (o *Material) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Material.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Material) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	materialUpdateCacheMut.RLock()
	cache, cached := materialUpdateCache[key]
	materialUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			materialAllColumns,
			materialPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update materials, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"materials\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, materialPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(materialType, materialMapping, append(wl, materialPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update materials row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for materials")
	}

	if !cached {
		materialUpdateCacheMut.Lock()
		materialUpdateCache[key] = cache
		materialUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q materialQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q materialQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for materials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for materials")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o MaterialSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MaterialSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), materialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"materials\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, materialPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in material slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all material")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Material) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Material) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no materials provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(materialColumnsWithDefault, o)

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

	materialUpsertCacheMut.RLock()
	cache, cached := materialUpsertCache[key]
	materialUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			materialAllColumns,
			materialColumnsWithDefault,
			materialColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			materialAllColumns,
			materialPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert materials, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(materialPrimaryKeyColumns))
			copy(conflict, materialPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"materials\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(materialType, materialMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(materialType, materialMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert materials")
	}

	if !cached {
		materialUpsertCacheMut.Lock()
		materialUpsertCache[key] = cache
		materialUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Material record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Material) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Material record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Material) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Material provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), materialPrimaryKeyMapping)
	sql := "DELETE FROM \"materials\" WHERE \"material_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from materials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for materials")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q materialQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q materialQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no materialQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from materials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for materials")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o MaterialSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MaterialSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(materialBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), materialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"materials\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, materialPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from material slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for materials")
	}

	if len(materialAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Material) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: no Material provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Material) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMaterial(ctx, exec, o.MaterialID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MaterialSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("orm: empty MaterialSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MaterialSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MaterialSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), materialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"materials\".* FROM \"materials\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, materialPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in MaterialSlice")
	}

	*o = slice

	return nil
}

// MaterialExistsG checks if the Material row exists.
func MaterialExistsG(ctx context.Context, materialID string) (bool, error) {
	return MaterialExists(ctx, boil.GetContextDB(), materialID)
}

// MaterialExists checks if the Material row exists.
func MaterialExists(ctx context.Context, exec boil.ContextExecutor, materialID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"materials\" where \"material_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, materialID)
	}
	row := exec.QueryRowContext(ctx, sql, materialID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if materials exists")
	}

	return exists, nil
}

// Exists checks if the Material row exists.
func (o *Material) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MaterialExists(ctx, exec, o.MaterialID)
}
