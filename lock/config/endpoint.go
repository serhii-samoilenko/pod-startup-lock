/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package config

import (
	"fmt"
	"log"
	"regexp"
)

var endpointPattern = regexp.MustCompile(`^(\S+?)://(.*)$`)
var addressPattern = regexp.MustCompile(`^(\S+):(\d+)$`)

type Endpoint interface {
	Protocol() string
	String() string
	IsHttp() bool
}

type RawEndpoint interface {
	Endpoint
	Address() string
}

type HttpEndpoint interface {
	Endpoint
	Url() string
}

type EndpointData struct {
	protocol string
}

type RawEndpointData struct {
	EndpointData
	address string
}

type HttpEndpointData struct {
	EndpointData
	url string
}

func (e *RawEndpointData) String() string {
	return fmt.Sprintf("%s", e.address)
}

func (e *HttpEndpointData) String() string {
	return fmt.Sprintf("%s", e.url)
}

func (e *EndpointData) Protocol() string {
	return e.protocol
}

func (e *EndpointData) IsHttp() bool {
	return isHttp(e.Protocol())
}

func isHttp(protocol string) bool {
	return protocol == "http" || protocol == "https"
}

func (e *RawEndpointData) Address() string {
	return e.address
}

func (e *HttpEndpointData) Url() string {
	return e.url
}

func ParseEndpoint(str string) Endpoint {
	match := endpointPattern.FindStringSubmatch(str)
	if match == nil || len(match) != 3 {
		log.Panicf("Endpoint malformed: '%s'", str)
	}
	protocol := match[1]
	address := match[2]
	if isHttp(protocol) {
		return CreateHttp(protocol, str)
	} else {
		return CreateRaw(protocol, address)
	}
}

func CreateRaw(protocol string, address string) RawEndpoint {
	match := addressPattern.FindStringSubmatch(address)
	if match == nil || len(match) != 3 {
		log.Panicf("Address malformed: '%s'", address)
	}
	return &RawEndpointData{EndpointData{protocol}, address}
}

func CreateHttp(protocol string, url string) HttpEndpoint {
	return &HttpEndpointData{EndpointData{protocol}, url}
}
