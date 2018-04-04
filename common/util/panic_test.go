/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
    "testing"
    "fmt"
    "github.com/stretchr/testify/require"
)

func TestPanicOnErrorNoError(t *testing.T) {
    // GIVEN
    // WHEN
    panicFunc := func() { PanicOnError(nil) }

    // THEN
    require.NotPanics(t, panicFunc)
}

func TestPanicOnErrorWithError(t *testing.T) {
    // GIVEN
    err := fmt.Errorf("err")
    expected := "err"

    // WHEN
    panicFunc := func() { PanicOnError(err) }


    // THEN
    require.PanicsWithValue(t, expected, panicFunc)
}
