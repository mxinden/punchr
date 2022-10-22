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

func testAuthorizations(t *testing.T) {
	t.Parallel()

	query := Authorizations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthorizationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
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

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorizationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Authorizations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorizationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthorizationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorizationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthorizationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Authorization exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthorizationExists to return true, but got false.")
	}
}

func testAuthorizationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authorizationFound, err := FindAuthorization(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authorizationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthorizationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Authorizations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthorizationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Authorizations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthorizationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authorizationOne := &Authorization{}
	authorizationTwo := &Authorization{}
	if err = randomize.Struct(seed, authorizationOne, authorizationDBTypes, false, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}
	if err = randomize.Struct(seed, authorizationTwo, authorizationDBTypes, false, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authorizationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authorizationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Authorizations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthorizationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authorizationOne := &Authorization{}
	authorizationTwo := &Authorization{}
	if err = randomize.Struct(seed, authorizationOne, authorizationDBTypes, false, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}
	if err = randomize.Struct(seed, authorizationTwo, authorizationDBTypes, false, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authorizationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authorizationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authorizationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func authorizationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Authorization) error {
	*o = Authorization{}
	return nil
}

func testAuthorizationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Authorization{}
	o := &Authorization{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authorizationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Authorization object: %s", err)
	}

	AddAuthorizationHook(boil.BeforeInsertHook, authorizationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authorizationBeforeInsertHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.AfterInsertHook, authorizationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authorizationAfterInsertHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.AfterSelectHook, authorizationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authorizationAfterSelectHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.BeforeUpdateHook, authorizationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authorizationBeforeUpdateHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.AfterUpdateHook, authorizationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authorizationAfterUpdateHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.BeforeDeleteHook, authorizationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authorizationBeforeDeleteHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.AfterDeleteHook, authorizationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authorizationAfterDeleteHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.BeforeUpsertHook, authorizationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authorizationBeforeUpsertHooks = []AuthorizationHook{}

	AddAuthorizationHook(boil.AfterUpsertHook, authorizationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authorizationAfterUpsertHooks = []AuthorizationHook{}
}

func testAuthorizationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthorizationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authorizationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthorizationToManyClients(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Authorization
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AuthorizationID = a.ID
	c.AuthorizationID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Clients().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AuthorizationID == b.AuthorizationID {
			bFound = true
		}
		if v.AuthorizationID == c.AuthorizationID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthorizationSlice{&a}
	if err = a.L.LoadClients(ctx, tx, false, (*[]*Authorization)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Clients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Clients = nil
	if err = a.L.LoadClients(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Clients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthorizationToManyAddOpClients(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Authorization
	var b, c, d, e Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorizationDBTypes, false, strmangle.SetComplement(authorizationPrimaryKeyColumns, authorizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Client{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Client{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddClients(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AuthorizationID {
			t.Error("foreign key was wrong value", a.ID, first.AuthorizationID)
		}
		if a.ID != second.AuthorizationID {
			t.Error("foreign key was wrong value", a.ID, second.AuthorizationID)
		}

		if first.R.Authorization != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Authorization != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Clients[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Clients[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Clients().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthorizationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
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

func testAuthorizationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthorizationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthorizationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Authorizations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authorizationDBTypes = map[string]string{`ID`: `integer`, `APIKey`: `character varying`, `Username`: `text`, `CreatedAt`: `timestamp with time zone`}
	_                    = bytes.MinRead
)

func testAuthorizationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authorizationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authorizationAllColumns) == len(authorizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthorizationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authorizationAllColumns) == len(authorizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Authorization{}
	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authorizationDBTypes, true, authorizationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authorizationAllColumns, authorizationPrimaryKeyColumns) {
		fields = authorizationAllColumns
	} else {
		fields = strmangle.SetComplement(
			authorizationAllColumns,
			authorizationPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, authorizationGeneratedColumns)
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

	slice := AuthorizationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthorizationsUpsert(t *testing.T) {
	t.Parallel()

	if len(authorizationAllColumns) == len(authorizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Authorization{}
	if err = randomize.Struct(seed, &o, authorizationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Authorization: %s", err)
	}

	count, err := Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authorizationDBTypes, false, authorizationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Authorization struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Authorization: %s", err)
	}

	count, err = Authorizations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
