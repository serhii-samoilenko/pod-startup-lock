/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRetrySuccess(t *testing.T) {
	// GIVEN
	expected := "success"
	successFunc := func() (interface{}, error) {
		return expected, nil
	}

	// WHEN
	actual := (*RetryOrPanic(1, 1, successFunc)).(string)

	// THEN
	require.Equal(t, actual, expected)
}

func TestRetryDefaultSuccess(t *testing.T) {
	// GIVEN
	expected := "success"
	successFunc := func() (interface{}, error) {
		return expected, nil
	}

	// WHEN
	actual := (*RetryOrPanicDefault(successFunc)).(string)

	// THEN
	require.Equal(t, actual, expected)
}

func TestRetryFail(t *testing.T) {
	// GIVEN
	errorFunc := func() (interface{}, error) {
		return nil, fmt.Errorf("error")
	}
	expected := "Failed after 1 attempts, last error: error"

	// WHEN
	panicFunc := func() { RetryOrPanic(1, 1, errorFunc) }

	// THEN
	require.PanicsWithValue(t, expected, panicFunc)
}

func TestMultipleRetryFail(t *testing.T) {
	// GIVEN
	errorFunc := func() (interface{}, error) {
		return nil, fmt.Errorf("error")
	}
	expected := "Failed after 2 attempts, last error: error"

	// WHEN
	panicFunc := func() { RetryOrPanic(2, 0, errorFunc) }

	// THEN
	require.PanicsWithValue(t, expected, panicFunc)
}
