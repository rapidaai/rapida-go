// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GetAllPlans(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllPlansRequest, opts ...grpc.CallOption) (*web_api.GetAllPlansResponse, error) {
	c, err := connection.BillingServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllPlans(connection.WithAuth(ctx), req, opts...)
}

func GetSubscription(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetSubscriptionRequest, opts ...grpc.CallOption) (*web_api.GetSubscriptionResponse, error) {
	c, err := connection.BillingServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetSubscription(connection.WithAuth(ctx), req, opts...)
}

func UpdateSubscription(ctx context.Context, connection connections.ConnectionConfig, req *web_api.UpdateSubscriptionRequest, opts ...grpc.CallOption) (*web_api.UpdateSubscriptionResponse, error) {
	c, err := connection.BillingServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateSubscription(connection.WithAuth(ctx), req, opts...)
}
