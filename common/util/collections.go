/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

func ArrayContains(haystack []string, needle string) bool {
	for _, key := range haystack {
		if key == needle {
			return true
		}
	}
	return false
}

func MapContainsAllPairs(haystack map[string]string, needles []Pair) bool {
	if len(needles) == 0 {
		return false
	}
	for _, pair := range needles {
		found, ok := haystack[pair.A]
		if !ok || found != pair.B {
			return false
		}
	}
	return true
}

func MapContainsAnyPair(haystack map[string]string, needles []Pair) bool {
	for _, pair := range needles {
		found, ok := haystack[pair.A]
		if ok && found == pair.B {
			return true
		}
	}
	return false
}

func MapContainsAll(haystack map[string]string, needles map[string]string) bool {
	if len(needles) == 0 {
		return false
	}
	for key, val := range needles {
		found, ok := haystack[key]
		if !ok || found != val {
			return false
		}
	}
	return true
}
