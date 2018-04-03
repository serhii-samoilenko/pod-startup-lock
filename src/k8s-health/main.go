/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package main

import (
    "k8s-health/config"
    "k8s-health/healthcheck"
    "k8s-health/k8s"
    "k8s-health/service"
)

func main() {
    conf := config.Parse()
    conf.Validate()

    k8sClient := k8s.NewClient(conf)
    endpointChecker := healthcheck.NewHealthChecker(conf, k8sClient)
    srv := service.NewService(conf.Host, conf.Port, endpointChecker.HealthFunction())

    go srv.Run()
    go endpointChecker.Run()

    select {} // Wait forever and let child goroutines run
}
