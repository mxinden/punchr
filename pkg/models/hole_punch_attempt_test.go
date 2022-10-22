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

func testHolePunchAttempts(t *testing.T) {
	t.Parallel()

	query := HolePunchAttempts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHolePunchAttemptsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
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

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchAttemptsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := HolePunchAttempts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchAttemptsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HolePunchAttemptSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchAttemptsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HolePunchAttemptExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if HolePunchAttempt exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HolePunchAttemptExists to return true, but got false.")
	}
}

func testHolePunchAttemptsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	holePunchAttemptFound, err := FindHolePunchAttempt(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if holePunchAttemptFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHolePunchAttemptsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = HolePunchAttempts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHolePunchAttemptsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := HolePunchAttempts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHolePunchAttemptsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	holePunchAttemptOne := &HolePunchAttempt{}
	holePunchAttemptTwo := &HolePunchAttempt{}
	if err = randomize.Struct(seed, holePunchAttemptOne, holePunchAttemptDBTypes, false, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}
	if err = randomize.Struct(seed, holePunchAttemptTwo, holePunchAttemptDBTypes, false, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holePunchAttemptOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holePunchAttemptTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := HolePunchAttempts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHolePunchAttemptsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	holePunchAttemptOne := &HolePunchAttempt{}
	holePunchAttemptTwo := &HolePunchAttempt{}
	if err = randomize.Struct(seed, holePunchAttemptOne, holePunchAttemptDBTypes, false, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}
	if err = randomize.Struct(seed, holePunchAttemptTwo, holePunchAttemptDBTypes, false, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holePunchAttemptOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holePunchAttemptTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func holePunchAttemptBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func holePunchAttemptAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchAttempt) error {
	*o = HolePunchAttempt{}
	return nil
}

func testHolePunchAttemptsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &HolePunchAttempt{}
	o := &HolePunchAttempt{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, false); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt object: %s", err)
	}

	AddHolePunchAttemptHook(boil.BeforeInsertHook, holePunchAttemptBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptBeforeInsertHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.AfterInsertHook, holePunchAttemptAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptAfterInsertHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.AfterSelectHook, holePunchAttemptAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptAfterSelectHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.BeforeUpdateHook, holePunchAttemptBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptBeforeUpdateHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.AfterUpdateHook, holePunchAttemptAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptAfterUpdateHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.BeforeDeleteHook, holePunchAttemptBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptBeforeDeleteHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.AfterDeleteHook, holePunchAttemptAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptAfterDeleteHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.BeforeUpsertHook, holePunchAttemptBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptBeforeUpsertHooks = []HolePunchAttemptHook{}

	AddHolePunchAttemptHook(boil.AfterUpsertHook, holePunchAttemptAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	holePunchAttemptAfterUpsertHooks = []HolePunchAttemptHook{}
}

func testHolePunchAttemptsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHolePunchAttemptsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(holePunchAttemptColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHolePunchAttemptToManyMultiAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchAttempt
	var b, c MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, multiAddressDBTypes, false, multiAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, multiAddressDBTypes, false, multiAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"hole_punch_attempt_x_multi_addresses\" (\"hole_punch_attempt\", \"multi_address_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"hole_punch_attempt_x_multi_addresses\" (\"hole_punch_attempt\", \"multi_address_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.MultiAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HolePunchAttemptSlice{&a}
	if err = a.L.LoadMultiAddresses(ctx, tx, false, (*[]*HolePunchAttempt)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MultiAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MultiAddresses = nil
	if err = a.L.LoadMultiAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MultiAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHolePunchAttemptToManyAddOpMultiAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchAttempt
	var b, c, d, e MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchAttemptDBTypes, false, strmangle.SetComplement(holePunchAttemptPrimaryKeyColumns, holePunchAttemptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MultiAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*MultiAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMultiAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.HolePunchAttempts[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.HolePunchAttempts[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.MultiAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MultiAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MultiAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testHolePunchAttemptToManySetOpMultiAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchAttempt
	var b, c, d, e MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchAttemptDBTypes, false, strmangle.SetComplement(holePunchAttemptPrimaryKeyColumns, holePunchAttemptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MultiAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetMultiAddresses(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.MultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetMultiAddresses(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.MultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.HolePunchAttempts) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.HolePunchAttempts) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.HolePunchAttempts[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.HolePunchAttempts[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.MultiAddresses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.MultiAddresses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testHolePunchAttemptToManyRemoveOpMultiAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchAttempt
	var b, c, d, e MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchAttemptDBTypes, false, strmangle.SetComplement(holePunchAttemptPrimaryKeyColumns, holePunchAttemptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MultiAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddMultiAddresses(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.MultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveMultiAddresses(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.MultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.HolePunchAttempts) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.HolePunchAttempts) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.HolePunchAttempts[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.HolePunchAttempts[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.MultiAddresses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.MultiAddresses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.MultiAddresses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testHolePunchAttemptToOneHolePunchResultUsingHolePunchResult(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local HolePunchAttempt
	var foreign HolePunchResult

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, holePunchAttemptDBTypes, false, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResult struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.HolePunchResultID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.HolePunchResult().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := HolePunchAttemptSlice{&local}
	if err = local.L.LoadHolePunchResult(ctx, tx, false, (*[]*HolePunchAttempt)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.HolePunchResult = nil
	if err = local.L.LoadHolePunchResult(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testHolePunchAttemptToOneSetOpHolePunchResultUsingHolePunchResult(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchAttempt
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchAttemptDBTypes, false, strmangle.SetComplement(holePunchAttemptPrimaryKeyColumns, holePunchAttemptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*HolePunchResult{&b, &c} {
		err = a.SetHolePunchResult(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.HolePunchResult != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.HolePunchAttempts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.HolePunchResultID != x.ID {
			t.Error("foreign key was wrong value", a.HolePunchResultID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.HolePunchResultID))
		reflect.Indirect(reflect.ValueOf(&a.HolePunchResultID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.HolePunchResultID != x.ID {
			t.Error("foreign key was wrong value", a.HolePunchResultID, x.ID)
		}
	}
}

func testHolePunchAttemptsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
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

func testHolePunchAttemptsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HolePunchAttemptSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHolePunchAttemptsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := HolePunchAttempts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	holePunchAttemptDBTypes = map[string]string{`ID`: `integer`, `HolePunchResultID`: `integer`, `OpenedAt`: `timestamp with time zone`, `StartedAt`: `timestamp with time zone`, `EndedAt`: `timestamp with time zone`, `StartRTT`: `interval`, `ElapsedTime`: `interval`, `Outcome`: `enum.hole_punch_attempt_outcome('UNKNOWN','DIRECT_DIAL','PROTOCOL_ERROR','CANCELLED','TIMEOUT','FAILED','SUCCESS')`, `Error`: `text`, `DirectDialError`: `text`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_                       = bytes.MinRead
)

func testHolePunchAttemptsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(holePunchAttemptPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(holePunchAttemptAllColumns) == len(holePunchAttemptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHolePunchAttemptsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(holePunchAttemptAllColumns) == len(holePunchAttemptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchAttempt{}
	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holePunchAttemptDBTypes, true, holePunchAttemptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(holePunchAttemptAllColumns, holePunchAttemptPrimaryKeyColumns) {
		fields = holePunchAttemptAllColumns
	} else {
		fields = strmangle.SetComplement(
			holePunchAttemptAllColumns,
			holePunchAttemptPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, holePunchAttemptGeneratedColumns)
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

	slice := HolePunchAttemptSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHolePunchAttemptsUpsert(t *testing.T) {
	t.Parallel()

	if len(holePunchAttemptAllColumns) == len(holePunchAttemptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := HolePunchAttempt{}
	if err = randomize.Struct(seed, &o, holePunchAttemptDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert HolePunchAttempt: %s", err)
	}

	count, err := HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, holePunchAttemptDBTypes, false, holePunchAttemptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchAttempt struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert HolePunchAttempt: %s", err)
	}

	count, err = HolePunchAttempts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
