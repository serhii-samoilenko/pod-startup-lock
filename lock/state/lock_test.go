/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package state

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var timeout = time.Duration(10) * time.Second

func TestAcquireSingleIfFirst(t *testing.T) {
	// GIVEN
	lock := NewLock(1)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.True(t, success)
}

func TestAcquireSingleIfSecond(t *testing.T) {
	// GIVEN
	lock := NewLock(1)
	lock.Acquire(timeout)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.False(t, success)
}

func TestAcquireSingleIfReleased(t *testing.T) {
	// GIVEN
	lock := NewLock(1)
	lock.Acquire(0)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.True(t, success)
}

func TestAcquireMultipleIfFirst(t *testing.T) {
	// GIVEN
	lock := NewLock(2)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.True(t, success)
}

func TestAcquireMultipleIfSecond(t *testing.T) {
	// GIVEN
	lock := NewLock(2)
	lock.Acquire(timeout)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.True(t, success)
}

func TestAcquireMultipleIfExceed(t *testing.T) {
	// GIVEN
	lock := NewLock(2)
	lock.Acquire(timeout)
	lock.Acquire(timeout)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.False(t, success)
}

func TestAcquireMultipleIfReleased(t *testing.T) {
	// GIVEN
	lock := NewLock(2)
	lock.Acquire(0)
	lock.Acquire(timeout)

	// WHEN
	success := lock.Acquire(timeout)

	// THEN
	require.True(t, success)
}
