/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package config

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestNewEndpointIfHttp(t *testing.T) {
    // GIVEN
    expectedProtocol := "http"

    // WHEN
    actual := ParseEndpoint("http://localhost:1234")

    // THEN
    require.Equal(t, expectedProtocol, actual.Protocol())
    require.True(t, actual.IsHttp())
}

func TestNewEndpointIfHttps(t *testing.T) {
    // GIVEN
    expectedProtocol := "https"

    // WHEN
    actual := ParseEndpoint("https://localhost:1234")

    // THEN
    require.Equal(t, expectedProtocol, actual.Protocol())
    require.True(t, actual.IsHttp())
}

func TestNewEndpointIfTcp(t *testing.T) {
    // GIVEN
    expectedProtocol := "tcp"

    // WHEN
    actual := ParseEndpoint("tcp://localhost:1234")

    // THEN
    require.Equal(t, expectedProtocol, actual.Protocol())
    require.False(t, actual.IsHttp())
}

func TestNewEndpointIfInvalidString(t *testing.T) {
    // GIVEN
    // WHEN
    panicFunc := func() { ParseEndpoint("localhost_1234") }

    // THEN
    require.PanicsWithValue(t, "Endpoint malformed: 'localhost_1234'", panicFunc)
}

func TestNewEndpointIfInvalidPort(t *testing.T) {
    // GIVEN
    // WHEN
    panicFunc := func() { ParseEndpoint("localhost:abcd") }

    // THEN
    require.PanicsWithValue(t, "Endpoint malformed: 'localhost:abcd'", panicFunc)
}

func TestNewEndpointIfInvalidProtocol(t *testing.T) {
    // GIVEN
    // WHEN
    panicFunc := func() { ParseEndpoint("localhost:1234") }

    // THEN
    require.PanicsWithValue(t, "Endpoint malformed: 'localhost:1234'", panicFunc)
}

func TestRawEndpointAddress(t *testing.T) {
    // GIVEN
    expectedAddress := "localhost:1234"

    // WHEN
    actual := ParseEndpoint("tcp://localhost:1234").(RawEndpoint)

    // THEN
    require.Equal(t, expectedAddress, actual.Address())
}

func TestHttpEndpointUrl(t *testing.T) {
    // GIVEN
    expectedUrl := "http://localhost:1234"

    // WHEN
    actual := ParseEndpoint("http://localhost:1234").(HttpEndpoint)

    // THEN
    require.Equal(t, expectedUrl, actual.Url())
}

func TestRawEndpointAddressNoPort(t *testing.T) {
    // GIVEN
    // WHEN
    panicFunc := func() { ParseEndpoint("tcp://localhost") }

    // THEN
    require.PanicsWithValue(t, "Address malformed: 'localhost'", panicFunc)
}

func TestHttpEndpointUrlNoPort(t *testing.T) {
    // GIVEN
    expectedUrl := "http://localhost"

    // WHEN
    actual := ParseEndpoint("http://localhost").(HttpEndpoint)

    // THEN
    require.Equal(t, expectedUrl, actual.Url())
}

func TestRawEndpointString(t *testing.T) {
    // GIVEN
    expected := "localhost:1234"

    // WHEN
    actual := ParseEndpoint("tcp://localhost:1234")

    // THEN
    require.Equal(t, expected, actual.String())
}

func TestHttpEndpointString(t *testing.T) {
    // GIVEN
    expected := "http://localhost:1234"

    // WHEN
    actual := ParseEndpoint("http://localhost:1234")

    // THEN
    require.Equal(t, expected, actual.String())
}

