// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRouters(t *testing.T) {
	t.Parallel()

	query := Routers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoutersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Routers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RouterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Router exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RouterExists to return true, but got false.")
	}
}

func testRoutersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routerFound, err := FindRouter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if routerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoutersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Routers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoutersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Routers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoutersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routerOne := &Router{}
	routerTwo := &Router{}
	if err = randomize.Struct(seed, routerOne, routerDBTypes, false, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}
	if err = randomize.Struct(seed, routerTwo, routerDBTypes, false, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Routers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoutersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routerOne := &Router{}
	routerTwo := &Router{}
	if err = randomize.Struct(seed, routerOne, routerDBTypes, false, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}
	if err = randomize.Struct(seed, routerTwo, routerDBTypes, false, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func routerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Router) error {
	*o = Router{}
	return nil
}

func testRoutersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Router{}
	o := &Router{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Router object: %s", err)
	}

	AddRouterHook(boil.BeforeInsertHook, routerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routerBeforeInsertHooks = []RouterHook{}

	AddRouterHook(boil.AfterInsertHook, routerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routerAfterInsertHooks = []RouterHook{}

	AddRouterHook(boil.AfterSelectHook, routerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routerAfterSelectHooks = []RouterHook{}

	AddRouterHook(boil.BeforeUpdateHook, routerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routerBeforeUpdateHooks = []RouterHook{}

	AddRouterHook(boil.AfterUpdateHook, routerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routerAfterUpdateHooks = []RouterHook{}

	AddRouterHook(boil.BeforeDeleteHook, routerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routerBeforeDeleteHooks = []RouterHook{}

	AddRouterHook(boil.AfterDeleteHook, routerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routerAfterDeleteHooks = []RouterHook{}

	AddRouterHook(boil.BeforeUpsertHook, routerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routerBeforeUpsertHooks = []RouterHook{}

	AddRouterHook(boil.AfterUpsertHook, routerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routerAfterUpsertHooks = []RouterHook{}
}

func testRoutersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoutersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouterToOnePeerUsingClient(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Router
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routerDBTypes, false, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ClientID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Client().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RouterSlice{&local}
	if err = local.L.LoadClient(ctx, tx, false, (*[]*Router)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Client = nil
	if err = local.L.LoadClient(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRouterToOneSetOpPeerUsingClient(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Router
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routerDBTypes, false, strmangle.SetComplement(routerPrimaryKeyColumns, routerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Peer{&b, &c} {
		err = a.SetClient(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Client != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClientRouters[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClientID != x.ID {
			t.Error("foreign key was wrong value", a.ClientID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClientID))
		reflect.Indirect(reflect.ValueOf(&a.ClientID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClientID != x.ID {
			t.Error("foreign key was wrong value", a.ClientID, x.ID)
		}
	}
}

func testRoutersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoutersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoutersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Routers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routerDBTypes = map[string]string{`ID`: `integer`, `ClientID`: `integer`, `HTML`: `text`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testRoutersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routerAllColumns) == len(routerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routerDBTypes, true, routerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoutersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routerAllColumns) == len(routerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Router{}
	if err = randomize.Struct(seed, o, routerDBTypes, true, routerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routerDBTypes, true, routerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routerAllColumns, routerPrimaryKeyColumns) {
		fields = routerAllColumns
	} else {
		fields = strmangle.SetComplement(
			routerAllColumns,
			routerPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, routerGeneratedColumns)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RouterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoutersUpsert(t *testing.T) {
	t.Parallel()

	if len(routerAllColumns) == len(routerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Router{}
	if err = randomize.Struct(seed, &o, routerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Router: %s", err)
	}

	count, err := Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routerDBTypes, false, routerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Router struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Router: %s", err)
	}

	count, err = Routers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
