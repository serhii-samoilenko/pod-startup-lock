/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
    "net/http"
    "lock/state"
    "log"
    "net/url"
    "time"
    "strconv"
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
    timeout, ok := getRequestedTimeout(r.URL.Query())
    if !ok {
        timeout = h.defaultTimeout
    }

    if h.lock.Acquire(timeout) {
        respondOk(w, r)
    } else {
        respondLocked(w, r)
    }
}

func getRequestedTimeout(values url.Values) (time.Duration, bool) {
    timeoutStr := values.Get("timeout")
    if timeoutStr == "" {
        return 0, false
    }
    timeout, err := strconv.Atoi(timeoutStr)
    if err != nil {
        log.Printf("Invalid timeout requested: '%v'", timeoutStr)
        return 0, false
    }
    return time.Duration(timeout) * time.Second, true
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
