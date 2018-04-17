/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"io/ioutil"
)

const defaultHost = "localhost"
const defaultPort = 8888
const defaultPause = 1
const defaultTimeout = 0

const maxIdleConnections = 1
const requestTimeout = 1 * time.Second

func main() {
	host := flag.String("host", defaultHost, "Lock service host")
	port := flag.Int("port", defaultPort, "Lock service port")
	duration := flag.Int("duration", defaultTimeout, "Custom lock duration to request, sec")
	pauseSec := flag.Int("pause", defaultPause, "Pause between lock attempts, sec")
	flag.Parse()

	pause := time.Duration(*pauseSec) * time.Second
	url := fmt.Sprintf("http://%s:%v", *host, *port)
	if *duration > 0 {
		url = fmt.Sprintf("%s?duration=%v", url, *duration)
	}
	log.Printf("Will try to acquire lock at '%s' each '%v' sec", url, *pauseSec)

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: requestTimeout,
	}
	for {
		if acquireLock(client, url) {
			return
		}
		time.Sleep(pause)
	}
}

func acquireLock(client *http.Client, url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error occurred: '%v'", err)
		return false
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Lock not acquired, waiting (status: %v)", resp.StatusCode)
		return false
	}
	log.Print("Lock acquired, exiting")
	return true
}
