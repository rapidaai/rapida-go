// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"log"
	"strings"
)

type RapidaRegion string

const (
	AP  RapidaRegion = "ap"
	US  RapidaRegion = "us"
	EU  RapidaRegion = "eu"
	ALL RapidaRegion = "all"
)

// Get returns the string value of the RapidaRegion
func (r RapidaRegion) Get() string {
	return string(r)
}

// FromStr returns the corresponding RapidaRegion for a given string,
// or ALL if the string does not match any region.
func FromRegionStr(label string) RapidaRegion {
	switch strings.ToLower(label) {
	case "ap":
		return AP
	case "us":
		return US
	case "eu":
		return EU
	case "all":
		return ALL
	default:
		log.Printf("The region is not supported. Supported regions are 'ap', 'us', 'eu', and 'all'.")
		return ALL
	}
}
