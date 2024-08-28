/*
 *  Copyright (c) 2024. Rapida
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 *
 *  Author: Prashant Srivastav
 *
 */

package rapida

import (
	"os"

	rapida_constants "github.com/rapidaai/rapida-sdk/rapida/constants"
)

type RapidaClientOption struct {
	RapidaApiKey       *string
	RapidaEndpointUrl  *string
	RapidaAssistantUrl *string
	RapidaRegion       *rapida_constants.RapidaRegion
	RapidaEnvironment  *rapida_constants.RapidaEnvironment
	IsSecure           bool
}

const (
	ENDPOINT_URL  = "connect.rapida.ai"
	ASSISTANT_URL = "assistant.rapida.ai:8080"
)

// NewRapidaClientOptionWithParams initializes RapidaClientOption with provided parameters.
func NewRapidaClientOptionWithParams(apiKey, endpointUrl, assistantUrl string, environment rapida_constants.RapidaEnvironment, region rapida_constants.RapidaRegion, isSecure bool) *RapidaClientOption {
	return &RapidaClientOption{
		RapidaApiKey:       &apiKey,
		RapidaEndpointUrl:  &endpointUrl,
		RapidaAssistantUrl: &assistantUrl,
		RapidaEnvironment:  &environment,
		RapidaRegion:       &region,
		IsSecure:           isSecure,
	}
}

// NewRapidaClientOption initializes RapidaClientOption with default parameters.
func NewRapidaClientOption() *RapidaClientOption {
	apiKey := os.Getenv("RAPIDA_API_KEY")
	return &RapidaClientOption{
		RapidaApiKey:       &apiKey,
		RapidaEndpointUrl:  Ptr(ENDPOINT_URL),
		RapidaAssistantUrl: Ptr(ASSISTANT_URL),
		RapidaEnvironment:  Ptr(rapida_constants.PRODUCTION),
		RapidaRegion:       Ptr(rapida_constants.ALL),
		IsSecure:           true,
	}
}

// GetRapidaApiKey returns the API key.
func (o *RapidaClientOption) GetRapidaApiKey() *string {
	return o.RapidaApiKey
}

// GetRapidaEndpointUrl returns the endpoint URL.
func (o *RapidaClientOption) GetRapidaEndpointUrl() *string {
	if o.RapidaEndpointUrl != nil {
		return o.RapidaEndpointUrl
	}
	endpointUrl := os.Getenv("RAPIDA_ENDPOINT_URL")
	if endpointUrl == "" {
		endpointUrl = ENDPOINT_URL
	}
	return &endpointUrl
}

// GetAssistantUrl returns the assistant URL.
func (o *RapidaClientOption) GetAssistantUrl() *string {
	if o.RapidaAssistantUrl != nil {
		return o.RapidaAssistantUrl
	}
	assistantUrl := os.Getenv("RAPIDA_ASSISTANT_URL")
	if assistantUrl == "" {
		assistantUrl = ASSISTANT_URL
	}
	return &assistantUrl
}

// GetRapidaRegion returns the RapidaRegion.
func (o *RapidaClientOption) GetRapidaRegion() *rapida_constants.RapidaRegion {
	if o.RapidaRegion != nil {
		return o.RapidaRegion
	}
	return Ptr(rapida_constants.ALL)
}

// GetRapidaEnvironment returns the RapidaEnvironment.
func (o *RapidaClientOption) GetRapidaEnvironment() *rapida_constants.RapidaEnvironment {
	if o.RapidaEnvironment != nil {
		return o.RapidaEnvironment
	}
	return Ptr(rapida_constants.PRODUCTION)
}

// SetRapidaEndpointUrl sets the endpoint URL.
func (o *RapidaClientOption) SetRapidaEndpointUrl(endpointUrl string) {
	o.RapidaEndpointUrl = &endpointUrl
}

// SetRapidaAssistantUrl sets the assistant URL.
func (o *RapidaClientOption) SetRapidaAssistantUrl(assistantUrl string) {
	o.RapidaAssistantUrl = &assistantUrl
}

// SetRapidaRegion sets the RapidaRegion.
func (o *RapidaClientOption) SetRapidaRegion(region rapida_constants.RapidaRegion) {
	o.RapidaRegion = &region
}

// SetRapidaEnvironment sets the RapidaEnvironment.
func (o *RapidaClientOption) SetRapidaEnvironment(environment rapida_constants.RapidaEnvironment) {
	o.RapidaEnvironment = &environment
}

// SetSecure sets the security flag.
func (o *RapidaClientOption) SetSecure(secure bool) {
	o.IsSecure = secure
}

// Utility functions to create pointers from values

func Ptr[T any](v T) *T {
	return &v
}
