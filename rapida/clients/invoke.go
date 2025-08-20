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
 *  Author: Prashant <prashant@rapida.ai>
 *
 *  This module provides functions for accessing audit logs via gRPC. It includes
 *  operations for retrieving lists of audit logs and fetching specific audit log entries.
 */
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
