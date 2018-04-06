/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
	"fmt"
	"log"
	"net/http"
)

func Run(host string, port int, handler http.Handler) {
	log.Print("Starting Http Service...")
	addr := fmt.Sprintf("%s:%v", host, port)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Panic("Http Service failed to start: ", err)
	}
}
