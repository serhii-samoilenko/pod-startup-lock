/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
	"github.com/lisenet/pod-startup-lock/lock/state"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type lockHandler struct {
	lock            state.Lock
	defaultTimeout  time.Duration
	permitAcquiring func() bool
}

func NewLockHandler(lock state.Lock, defaultTimeout time.Duration, permitOperationChecker func() bool) http.Handler {
	return &lockHandler{lock, defaultTimeout, permitOperationChecker}
}

func (h *lockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.permitAcquiring() {
		respondLocked(w, r)
		return
	}
	duration, ok := getRequestedDuration(r.URL.Query())
	if !ok {
		duration = h.defaultTimeout
	}

	if h.lock.Acquire(duration) {
		respondOk(w, r)
	} else {
		respondLocked(w, r)
	}
}

func getRequestedDuration(values url.Values) (time.Duration, bool) {
	durationStr := values.Get("duration")
	if durationStr == "" {
		return 0, false
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		log.Printf("Invalid duration requested: '%v'", durationStr)
		return 0, false
	}
	return time.Duration(duration) * time.Second, true
}

func respondOk(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	log.Printf("Responding to '%v': %v", r.RemoteAddr, status)
	w.WriteHeader(status)
	w.Write([]byte("Lock acquired"))
}

func respondLocked(w http.ResponseWriter, r *http.Request) {
	status := http.StatusLocked
	log.Printf("Responding to '%v': %v", r.RemoteAddr, status)
	w.WriteHeader(status)
	w.Write([]byte("Locked"))
}
