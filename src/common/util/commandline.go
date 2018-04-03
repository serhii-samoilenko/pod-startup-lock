/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package util

import (
    "fmt"
    "strings"
    "log"
)

type ArrayVal []string

func (v *ArrayVal) String() string {
    return fmt.Sprintf("%v", []string(*v))
}

func (v *ArrayVal) Set(value string) error {
    *v = append(*v, value)
    return nil
}

type Pair struct {
    A string
    B string
}

type PairArrayVal struct {
    sep    string
    values []Pair
}

func NewPairArrayVal(sep string) PairArrayVal {
    return PairArrayVal{sep, make([]Pair, 0)}
}

func (v *PairArrayVal) String() string {
    return fmt.Sprintf("%v", v.values)
}

func (v *PairArrayVal) Set(value string) error {
    chunks := strings.Split(value, v.sep)
    if len(chunks) != 2 || chunks[0] == "" || chunks[1] == "" {
        log.Panicf("Failed to parse value: '%s'", value)
    }
    pair := Pair{chunks[0], chunks[1]}
    v.values = append(v.values, pair)
    return nil
}

func (v *PairArrayVal) Get() []Pair {
    return v.values
}
