// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func Invoke(connection connections.ConnectionConfig, ctx context.Context, req *web_api.InvokeRequest, opts ...grpc.CallOption) (*web_api.InvokeResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Invoke(connection.WithAuth(ctx), req, opts...)
}

func Update(connection connections.ConnectionConfig, ctx context.Context, req *web_api.UpdateRequest, opts ...grpc.CallOption) (*web_api.UpdateResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Update(connection.WithAuth(ctx), req, opts...)
}

func Probe(connection connections.ConnectionConfig, ctx context.Context, req *web_api.ProbeRequest, opts ...grpc.CallOption) (*web_api.ProbeResponse, error) {
	c, err := connection.DeploymentClient()
	if err != nil {
		return nil, err
	}
	return c.Probe(connection.WithAuth(ctx), req, opts...)
}
