/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
	"log"
	"time"
)

func RetryOrPanicDefault(call func() (interface{}, error)) *interface{} {
	return RetryOrPanic(5, 1, call)
}

func RetryOrPanic(attempts int, sleep time.Duration, call func() (interface{}, error)) *interface{} {
	for i := 0; ; i++ {
		result, err := call()
		if err == nil {
			return &result
		}
		if i >= (attempts - 1) {
			log.Panicf("Failed after %d attempts, last error: %v", attempts, err)
		}
		log.Println("Error occured, will retry:", err)
		time.Sleep(sleep * time.Second)
	}
}
