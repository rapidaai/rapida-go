// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.

package utils

import (
	"log"
	"strings"
)

type RapidaEnvironment string

const (
	PRODUCTION  RapidaEnvironment = "production"
	DEVELOPMENT RapidaEnvironment = "development"
)

// Get returns the string value of the RapidaEnvironment
func (e RapidaEnvironment) Get() string {
	return string(e)
}

// FromStr returns the corresponding RapidaEnvironment for a given string,
// or DEVELOPMENT if the string does not match any environment.
func FromEnvironmentStr(label string) RapidaEnvironment {
	switch strings.ToLower(label) {
	case "production":
		return PRODUCTION
	case "development":
		return DEVELOPMENT
	default:
		log.Printf("The environment is not supported. Only 'production' and 'development' are allowed.")
		return DEVELOPMENT
	}
}
