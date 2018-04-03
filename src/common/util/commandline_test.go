/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestArrayValStringWhenEmpty(t *testing.T) {
    // GIVEN
    arr := ArrayVal{}
    expected := "[]"

    // WHEN
    actual := arr.String()

    // THEN
    require.Equal(t, expected, actual)
}

func TestArrayValStringWhenNotEmpty(t *testing.T) {
    // GIVEN
    arr := ArrayVal{"val1", "val2"}
    expected := "[val1 val2]"

    // WHEN
    actual := arr.String()

    // THEN
    require.Equal(t, expected, actual)
}

func TestArrayValSet(t *testing.T) {
    // GIVEN
    arr := ArrayVal{}
    expected := ArrayVal{"val1", "val2"}

    // WHEN
    arr.Set("val1")
    arr.Set("val2")

    // THEN
    require.Equal(t, expected, arr)
}

func TestNewPairArrayVal(t *testing.T) {
    // GIVEN
    // WHEN
    arrayVal := NewPairArrayVal("-")

    // THEN
    require.Equal(t, "-", arrayVal.sep, "Wrong separator")
    require.Len(t, arrayVal.Get(), 0)
}

func TestPairArrayValString(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")
    arrayVal.Set("a:1")
    arrayVal.Set("b:2")
    expected := "[{a 1} {b 2}]"

    // WHEN
    actual := arrayVal.String()

    // THEN
    require.Equal(t, expected, actual)
}

func TestPairArrayValSetWhenEmpty(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")

    // WHEN
    panicFunc := func() { arrayVal.Set("") }

    // THEN
    require.PanicsWithValue(t, "Failed to parse value: ''", panicFunc)
}

func TestPairArrayValSetWhenNoValues(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")

    // WHEN
    panicFunc := func() { arrayVal.Set(":") }

    // THEN
    require.PanicsWithValue(t, "Failed to parse value: ':'", panicFunc)
}

func TestPairArrayValSetWhenNoKey(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")

    // WHEN
    panicFunc := func() { arrayVal.Set(":val") }

    // THEN
    require.PanicsWithValue(t, "Failed to parse value: ':val'", panicFunc)
}

func TestPairArrayValSetWhenNoValue(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")

    // WHEN
    panicFunc := func() { arrayVal.Set("key:") }

    // THEN
    require.PanicsWithValue(t, "Failed to parse value: 'key:'", panicFunc)
}

func TestPairArrayValSetSingle(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")
    expected := []Pair{{"a", "1"}}

    // WHEN
    arrayVal.Set("a:1")

    // THEN
    require.Equal(t, expected, arrayVal.Get())
}

func TestPairArrayValSetMultiple(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")
    expected := []Pair{{"a", "1"}, {"b", "2"}}

    // WHEN
    arrayVal.Set("a:1")
    arrayVal.Set("b:2")

    // THEN
    require.Equal(t, expected, arrayVal.Get())
}

func TestPairArrayValSetMultipleInvalid(t *testing.T) {
    // GIVEN
    arrayVal := NewPairArrayVal(":")

    // WHEN
    panicFunc := func() { arrayVal.Set("key::val") }

    // THEN
    require.PanicsWithValue(t, "Failed to parse value: 'key::val'", panicFunc)
}
