/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package state

import (
	"log"
	"sync"
	"time"
)

type Lock struct {
	maxCount int
	mutex    sync.Mutex
	locks    []time.Time
}

func NewLock(maxLockCount int) Lock {
	return Lock{maxCount: maxLockCount}
}

func (l *Lock) Acquire(duration time.Duration) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.removeExpired()
	if len(l.locks) < l.maxCount {
		l.addNew(duration)
		log.Printf("Lock acquired: %v of %v, duration: %v", len(l.locks), l.maxCount, duration)
		return true
	}
	return false
}

func (l *Lock) addNew(duration time.Duration) {
	expireTime := time.Now().Add(duration)
	l.locks = append(l.locks, expireTime)
}

func (l *Lock) removeExpired() {
	var live []time.Time
	for i := 0; i < len(l.locks); i++ {
		if !isExpired(l.locks[i]) {
			live = append(live, l.locks[i])
		}
	}
	l.locks = live
}

func isExpired(t time.Time) bool {
	return time.Now().After(t)
}
