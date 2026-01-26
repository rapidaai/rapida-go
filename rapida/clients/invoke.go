// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func Invoke(ctx context.Context, connection connections.ConnectionConfig, req *web_api.InvokeRequest, opts ...grpc.CallOption) (*web_api.InvokeResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Invoke(connection.WithAuth(ctx), req, opts...)
}

func Update(ctx context.Context, connection connections.ConnectionConfig, req *web_api.UpdateRequest, opts ...grpc.CallOption) (*web_api.UpdateResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Update(connection.WithAuth(ctx), req, opts...)
}

func Probe(ctx context.Context, connection connections.ConnectionConfig, req *web_api.ProbeRequest, opts ...grpc.CallOption) (*web_api.ProbeResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Probe(connection.WithAuth(ctx), req, opts...)
}
