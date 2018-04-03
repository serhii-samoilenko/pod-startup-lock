/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package config

import (
    . "common/util"
    "log"
    "flag"
    "time"
    "os"
)

const defaultPort = 9999
const defaultFailTimeout = 10
const defaultPassTimeout = 60

func Parse() Config {
    host := flag.String("host", "", "Host/Ip to bind")
    port := flag.Int("port", defaultPort, "Port to bind")
    baseUrl := flag.String("baseUrl", "", "K8s api base url. For out-of-cluster usage only")
    namespace := flag.String("namespace", "", "K8s Namespace to check DaemonSets in. Blank for all namespaces")
    failTimeout := flag.Int("failHc", defaultFailTimeout, "Pause between DaemonSet health checks if previous failed, sec")
    passTimeout := flag.Int("passHc", defaultPassTimeout, "Pause between DaemonSet health checks if previous succeeded, sec")
    hostNetwork := flag.Bool("hostNet", false, "Host network DaemonSets only")

    nodeName, _ := os.LookupEnv("NODE_NAME")

    includeDs := NewPairArrayVal(":")
    flag.Var(&includeDs, "in", "Include DaemonSet labels, label:value")
    excludeDs := NewPairArrayVal(":")
    flag.Var(&excludeDs, "ex", "Exclude DaemonSet labels, label:value")
    flag.Parse()

    config := Config{
        *host,
        *port,
        *baseUrl,
        *namespace,
        time.Duration(*failTimeout) * time.Second,
        time.Duration(*passTimeout) * time.Second,
        nodeName,
        *hostNetwork,
        includeDs.Get(),
        excludeDs.Get(),
    }
    log.Printf("Application config:\n%+v", config)
    config.Validate()
    return config
}

type Config struct {
    Host              string
    Port              int
    K8sApiBaseUrl     string
    Namespace         string
    HealthFailTimeout time.Duration
    HealthPassTimeout time.Duration
    NodeName          string
    HostNetworkDs     bool
    IncludeDs         []Pair
    ExcludeDs         []Pair
}

func (c *Config) Validate() {
    if c.NodeName == "" {
        log.Panic("NODE_NAME not specified")
    }
    if len(c.IncludeDs) > 0 && len(c.ExcludeDs) > 0 {
        log.Panic("Cannot specify both Included and Excluded DaemonSet labels, choose one")
    }
}
