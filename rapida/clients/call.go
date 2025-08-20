// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func CreatePhoneCall(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreatePhoneCallRequest) (*web_api.CreatePhoneCallResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreatePhoneCall(connection.WithAuth(ctx), req)
}

func CreateBulkPhoneCall(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateBulkPhoneCallRequest) (*web_api.CreateBulkPhoneCallResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateBulkPhoneCall(connection.WithAuth(ctx), req)
}
