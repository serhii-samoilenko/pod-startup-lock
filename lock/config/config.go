/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package config

import (
	"flag"
	"github.com/lisenet/pod-startup-lock/common/util"
	"log"
	"time"
)

const defaultPort = 8888
const defaultParallelLocks = 1
const defaultLockTimeout = 10
const defaultFailTimeout = 10
const defaultPassTimeout = 60

func Parse() Config {
	host := flag.String("host", "", "Host/Ip to bind")
	port := flag.Int("port", defaultPort, "Port to bind")
	parallelLocks := flag.Int("locks", defaultParallelLocks, "Count of locks allowed to acquire in parallel")
	lockTimeout := flag.Int("timeout", defaultLockTimeout, "Default lock timeout, sec")
	failTimeout := flag.Int("failHc", defaultFailTimeout, "Pause between endpoint health checks if previous failed, sec")
	passTimeout := flag.Int("passHc", defaultPassTimeout, "Pause between endpoint health checks if previous succeeded, sec")

	var healthEndpoints util.ArrayVal
	flag.Var(&healthEndpoints, "check", "HealthCheck tcp endpoint, host:port")
	flag.Parse()

	config := Config{
		*host,
		*port,
		*parallelLocks,
		time.Duration(*lockTimeout) * time.Second,
		time.Duration(*failTimeout) * time.Second,
		time.Duration(*passTimeout) * time.Second,
		parseEndpoints(healthEndpoints),
	}
	log.Printf("Application config:\n%+v", config)
	return config
}

type Config struct {
	Host              string
	Port              int
	ParallelLocks     int
	LockTimeout       time.Duration
	HealthFailTimeout time.Duration
	HealthPassTimeout time.Duration
	HealthEndpoints   []Endpoint
}

func parseEndpoints(vals []string) []Endpoint {
	var endpoints []Endpoint
	for _, url := range vals {
		endpoints = append(endpoints, ParseEndpoint(url))
	}
	return endpoints
}
