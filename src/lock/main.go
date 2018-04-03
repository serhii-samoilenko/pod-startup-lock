/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package main

import (
    "lock/config"
    "lock/service"
    "lock/state"
)

func main() {
    conf := config.Parse()
    endpointChecker := service.NewEndpointChecker(
        conf.HealthPassTimeout,
        conf.HealthFailTimeout,
        conf.HealthEndpoints,
    )

    healthFunc := endpointChecker.HealthFunction()
    lock := state.NewLock(conf.ParallelLocks)
    handler := service.NewLockHandler(lock, conf.LockTimeout, healthFunc)

    go service.Run(conf.Host, conf.Port, handler)
    go endpointChecker.Run()

    select {} // Wait forever and let child goroutines run
}
