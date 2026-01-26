// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

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
