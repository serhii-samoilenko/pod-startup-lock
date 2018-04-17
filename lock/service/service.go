/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const readTimeout = 2 * time.Second
const writeTimeout = 2 * time.Second
const idleTimeout = 10 * time.Second

func Run(host string, port int, handler http.Handler) {
	log.Print("Starting Http Service...")
	addr := fmt.Sprintf("%s:%v", host, port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Http Service failed to start: ", err)
	}
}
