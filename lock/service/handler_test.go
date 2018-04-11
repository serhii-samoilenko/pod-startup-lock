/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
	"github.com/serhii-samoilenko/pod-startup-lock/lock/state"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var timeout = time.Duration(10) * time.Second

func TestAcquireIfFirst(t *testing.T) {
	// GIVEN
	permitFunction := func() bool {
		return true
	}
	lock := state.NewLock(1)
	handler := NewLockHandler(lock, timeout, permitFunction)
	req, _ := http.NewRequest("GET", "/", nil)

	// WHEN
	rr := prepareResponseRecorder(req, handler)

	// THEN
	assertResponseStatusCode(http.StatusOK, rr.Code, t)
}

func TestAcquireIfSecond(t *testing.T) {
	// GIVEN
	permitFunction := func() bool {
		return true
	}

	lock := state.NewLock(1)
	handler := NewLockHandler(lock, timeout, permitFunction)
	req, _ := http.NewRequest("GET", "/", nil)
	prepareResponseRecorder(req, handler)

	// WHEN
	rr := prepareResponseRecorder(req, handler)

	// THEN
	assertResponseStatusCode(http.StatusLocked, rr.Code, t)
}

func TestAcquireIfWrongTimeoutRequested(t *testing.T) {
	// GIVEN
	permitFunction := func() bool {
		return true
	}

	lock := state.NewLock(1)
	handler := NewLockHandler(lock, timeout, permitFunction)
	req, _ := http.NewRequest("GET", "/", nil)
	q := req.URL.Query()
	q.Add("duration", "a")
	req.URL.RawQuery = q.Encode()

	prepareResponseRecorder(req, handler)

	// WHEN
	rr := prepareResponseRecorder(req, handler)

	// THEN
	assertResponseStatusCode(http.StatusLocked, rr.Code, t)
}

func TestAcquireIfZeroTimeoutRequested(t *testing.T) {
	// GIVEN
	permitFunction := func() bool {
		return true
	}

	lock := state.NewLock(1)
	handler := NewLockHandler(lock, timeout, permitFunction)
	req, _ := http.NewRequest("GET", "/", nil)
	q := req.URL.Query()
	q.Add("duration", "0")
	req.URL.RawQuery = q.Encode()

	prepareResponseRecorder(req, handler)

	// WHEN
	rr := prepareResponseRecorder(req, handler)

	// THEN
	assertResponseStatusCode(http.StatusOK, rr.Code, t)
}

func TestAcquireIfDisabled(t *testing.T) {
	// GIVEN
	permitFunction := func() bool {
		return false
	}

	lock := state.NewLock(1)
	handler := NewLockHandler(lock, timeout, permitFunction)
	req, _ := http.NewRequest("GET", "/", nil)

	// WHEN
	rr := prepareResponseRecorder(req, handler)

	// THEN
	assertResponseStatusCode(http.StatusLocked, rr.Code, t)
}

func assertResponseStatusCode(expected int, actual int, t *testing.T) {
	if actual != expected {
		t.Errorf("handler returned wrong status code: expected %v got %v", expected, actual)
	}
}

func prepareResponseRecorder(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
