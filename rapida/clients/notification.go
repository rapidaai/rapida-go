// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func UpdateNotificationSetting(ctx context.Context, connection connections.ConnectionConfig, req *web_api.UpdateNotificationSettingRequest, opts ...grpc.CallOption) (*web_api.NotificationSettingResponse, error) {
	c, err := connection.NotificationServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateNotificationSetting(connection.WithAuth(ctx), req, opts...)
}

func GetNotificationSetting(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetNotificationSettingRequest, opts ...grpc.CallOption) (*web_api.NotificationSettingResponse, error) {
	c, err := connection.NotificationServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetNotificationSettting(connection.WithAuth(ctx), req, opts...)
}
