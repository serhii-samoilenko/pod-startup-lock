/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayContainsWhenEmptyArray(t *testing.T) {
	// GIVEN
	var haystack []string
	var needle string

	// WHEN
	contains := ArrayContains(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestArrayContainsWhenEmptyValue(t *testing.T) {
	// GIVEN
	haystack := []string{"a", "b", "c"}
	var needle string

	// WHEN
	contains := ArrayContains(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestArrayContainsWhenContains(t *testing.T) {
	// GIVEN
	haystack := []string{"a", "b", "c"}
	needle := "b"

	// WHEN
	contains := ArrayContains(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestArrayContainsWhenNotContains(t *testing.T) {
	// GIVEN
	haystack := []string{"a", "b", "c"}
	needle := "d"

	// WHEN
	contains := ArrayContains(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenAllEmpty(t *testing.T) {
	// GIVEN
	haystack := make(map[string]string)
	var needle []Pair

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenPairsEmpty(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	var needle []Pair

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenNotContains(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"c", "3"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenNotContainsValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "3"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenContainsSingle(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAllPairsWhenContainsMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}, {"b", "2"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAllPairsWhenContainsOneOfMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}, {"c", "3"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllPairsWhenContainsOneOfMultipleValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}, {"b", "3"}}

	// WHEN
	contains := MapContainsAllPairs(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAnyPairWhenAllEmpty(t *testing.T) {
	// GIVEN
	haystack := make(map[string]string)
	var needle []Pair

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAnyPairWhenPairsEmpty(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	var needle []Pair

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAnyPairWhenNotContains(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"c", "3"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAnyPairWhenNotContainsValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "3"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAnyPairWhenContainsSingle(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAnyPairWhenContainsMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}, {"b", "2"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAnyPairWhenContainsOneOfMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "1"}, {"c", "3"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAnyPairWhenContainsOneOfMultipleValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needle := []Pair{{"a", "3"}, {"b", "2"}}

	// WHEN
	contains := MapContainsAnyPair(haystack, needle)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAllWhenAllEmpty(t *testing.T) {
	// GIVEN
	haystack := make(map[string]string)
	needles := make(map[string]string)

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllWhenPairsEmpty(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := make(map[string]string)

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllWhenNotContains(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"c": "3"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllWhenNotContainsValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"a": "3"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllWhenContainsSingle(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"a": "1"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAllWhenContainsMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"a": "1", "b": "2"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.True(t, contains)
}

func TestMapContainsAllWhenContainsOneOfMultiple(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"a": "1", "c": "3"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}

func TestMapContainsAllWhenContainsOneOfMultipleValue(t *testing.T) {
	// GIVEN
	haystack := map[string]string{"a": "1", "b": "2"}
	needles := map[string]string{"a": "1", "b": "3"}

	// WHEN
	contains := MapContainsAll(haystack, needles)

	// THEN
	require.False(t, contains)
}
