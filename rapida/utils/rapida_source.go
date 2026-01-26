// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"encoding/json"
	"log"
	"strings"
)

type RapidaSource string

const (
	WebPlugin RapidaSource = "web-plugin"
	Debugger  RapidaSource = "debugger"
	// api
	SDK       RapidaSource = "sdk"
	PhoneCall RapidaSource = "phone-call"
	Whatsapp  RapidaSource = "whatsapp"
)

// Get returns the string value of the RapidaRegion
func (r RapidaSource) Get() string {
	return string(r)
}

// FromStr returns the corresponding RapidaSource for a given string,
// or WebPlugin if the string does not match any source.
func FromSourceStr(label string) RapidaSource {
	switch strings.ToLower(label) {
	case "web-plugin":
		return WebPlugin
	case "debugger":
		return Debugger
	case "sdk":
		return SDK
	case "phone-call":
		return PhoneCall
	case "whatsapp":
		return Whatsapp

	default:
		log.Printf("%s The source is not supported. Supported sources are 'web-plugin', 'rapida-app', 'python-sdk', 'node-sdk', 'go-sdk', 'typescript-sdk', 'java-sdk', 'php-sdk', and 'rust-sdk'.", label)
		return Debugger
	}
}

func (c RapidaSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}
