// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func CreatePhoneCall(
	ctx context.Context,
	connection connections.ConnectionConfig,
	callRequest *web_api.CreatePhoneCallRequest,
	opts ...grpc.CallOption,
) (*web_api.CreatePhoneCallResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreatePhoneCall(connection.WithAuth(ctx), callRequest, opts...)
}

func CreateBulkPhoneCall(
	ctx context.Context,
	connection connections.ConnectionConfig,
	callRequest *web_api.CreateBulkPhoneCallRequest,
	opts ...grpc.CallOption,
) (*web_api.CreateBulkPhoneCallResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateBulkPhoneCall(connection.WithAuth(ctx), callRequest, opts...)
}
