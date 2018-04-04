/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package service

import (
    . "pod-startup-lock/lock/config"
    "log"
    "time"
    "net"
    "net/http"
)

var client = new(http.Client)

type EndpointChecker struct {
    waitOnPass time.Duration
    waitOnFail time.Duration
    endpoints  []Endpoint
    isHealthy  bool
}

func NewEndpointChecker(waitOnPass time.Duration, waitOnFail time.Duration, endpoints []Endpoint) EndpointChecker {
    return EndpointChecker{waitOnPass, waitOnFail, endpoints, false}
}

func (c *EndpointChecker) HealthFunction() func() bool {
    return func() bool {
        return c.isHealthy
    }
}

func (c *EndpointChecker) Run() {
    if len(c.endpoints) == 0 {
        log.Print("No Endpoints to check")
        c.isHealthy = true
        return
    }
    for {
        if checkAll(c.endpoints) {
            log.Print("Endpoint Check passed")
            c.isHealthy = true
            time.Sleep(c.waitOnPass)
        } else {
            log.Print("Endpoint Check failed")
            c.isHealthy = false
            time.Sleep(c.waitOnFail)
        }
    }
}

func checkAll(endpoints []Endpoint) bool {
    for _, endpoint := range endpoints {
        if !check(endpoint) {
            return false
        }
    }
    return true
}

func check(endpoint Endpoint) bool {
    if endpoint.IsHttp() {
        return checkHttp(endpoint.(HttpEndpoint))
    } else {
        return checkRaw(endpoint.(RawEndpoint))
    }
}

func checkRaw(endpoint RawEndpoint) bool {
    conn, err := net.Dial(endpoint.Protocol(), endpoint.Address())
    if err != nil {
        log.Printf("'%v' endpoint connection failed: '%v'", endpoint, err)
        return false
    }
    conn.Close()
    log.Printf("'%v' endpoint OK", endpoint)
    return true
}

func checkHttp(endpoint HttpEndpoint) bool {
    resp, err := client.Get(endpoint.Url())
    if err != nil {
        log.Printf("'%v' endpoint connection failed: '%v'", endpoint, err)
        return false
    }
    if isSuccessful(resp.StatusCode) {
        log.Printf("'%v' endpoint OK (status: %v)", endpoint, resp.StatusCode)
        return true
    } else {
        log.Printf("'%v' endpoint Fail (status: %v)", endpoint, resp.StatusCode)
        return false
    }
}

func isSuccessful(code int) bool {
    return code >= 200 && code < 300
}
