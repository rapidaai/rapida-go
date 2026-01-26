// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
)

// GetClientSource retrieves the client source information from the given context.
// It extracts the source string from the incoming metadata in the context using the
// HEADER_SOURCE_KEY and converts it to a RapidaSource type.
//
// Parameters:
//
//	ctx (context.Context): The context from which metadata is extracted. This context
//	                        typically contains incoming metadata related to the client.
//
// Returns:
//
//	RapidaSource: The source of the client, represented as a RapidaSource constant.
//	              The function relies on the `FromSourceStr` function to convert the
//	              extracted string to the corresponding RapidaSource value.
//
// Details:
//
//	The `metadata.ExtractIncoming(ctx)` function call extracts metadata from the context,
//	which is then accessed using `Get(HEADER_SOURCE_KEY)`. This returns a string representing
//	the source of the client. The `FromSourceStr` function is used to map this string to the
//	appropriate RapidaSource constant, allowing for easier handling and validation of client
//	source information.
func GetClientSource(ctx context.Context) (RapidaSource, bool) {
	v := metadata.ExtractIncoming(ctx).Get(HEADER_SOURCE_KEY)
	if v == "" {
		return FromSourceStr(v), false
	}
	return FromSourceStr(v), true
}

// GetClientRegion retrieves the client region information from the given context.
// It extracts the region string from the incoming metadata in the context using the
// HEADER_REGION_KEY and converts it to a RapidaRegion type.
//
// Parameters:
//
//	ctx (context.Context): The context from which metadata is extracted. This context
//	                        typically contains incoming metadata related to the client.
//
// Returns:
//
//	RapidaRegion: The region of the client, represented as a RapidaRegion constant.
//	              The function relies on the `FromRegionStr` function to convert the
//	              extracted string to the corresponding RapidaRegion value.
//
// Details:
//
//	The `metadata.ExtractIncoming(ctx)` function call extracts metadata from the context,
//	which is then accessed using `Get(HEADER_REGION_KEY)`. This returns a string representing
//	the region of the client. The `FromRegionStr` function is used to map this string to the
//	appropriate RapidaRegion constant, facilitating validation and consistent handling of
//	client region information.
func GetClientRegion(ctx context.Context) (RapidaRegion, bool) {
	v := metadata.ExtractIncoming(ctx).Get(HEADER_REGION_KEY)
	if v == "" {
		// return
		return FromRegionStr(v), false
	}
	return FromRegionStr(v), true
}

// GetClientEnvironment retrieves the client environment information from the given context.
// It extracts the environment string from the incoming metadata in the context using the
// HEADER_ENVIRONMENT_KEY and converts it to a RapidaEnvironment type.
//
// Parameters:
//
//	ctx (context.Context): The context from which metadata is extracted. This context
//	                        typically contains incoming metadata related to the client.
//
// Returns:
//
//	RapidaEnvironment: The environment of the client, represented as a RapidaEnvironment
//	                   constant. The function relies on the `FromEnvironmentStr` function
//	                   to convert the extracted string to the corresponding RapidaEnvironment
//	                   value.
//
// Details:
//
//	The `metadata.ExtractIncoming(ctx)` function call extracts metadata from the context,
//	which is then accessed using `Get(HEADER_ENVIRONMENT_KEY)`. This returns a string representing
//	the environment of the client. The `FromEnvironmentStr` function is used to map this string
//	to the appropriate RapidaEnvironment constant, ensuring accurate and consistent handling of
//	client environment information.
func GetClientEnvironment(ctx context.Context) (RapidaEnvironment, bool) {
	v := metadata.ExtractIncoming(ctx).Get(HEADER_ENVIRONMENT_KEY)
	if v == "" {
		return FromEnvironmentStr(v), false
	}
	return FromEnvironmentStr(v), true
}

func GetAuthId(ctx context.Context) (*string, bool) {
	v := metadata.ExtractIncoming(ctx).Get(HEADER_AUTH_KEY)
	if v == "" {
		return nil, false
	}
	return &v, true
}
