// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GeneralConnect(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GeneralConnectRequest, opts ...grpc.CallOption) (*web_api.GeneralConnectResponse, error) {
	c, err := connection.ConnectServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GeneralConnect(connection.WithAuth(ctx), req, opts...)
}

func GetConnectorFiles(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetConnectorFilesRequest, opts ...grpc.CallOption) (*web_api.GetConnectorFilesResponse, error) {
	c, err := connection.ConnectServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetConnectorFiles(connection.WithAuth(ctx), req, opts...)
}
