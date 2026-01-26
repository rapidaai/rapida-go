// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func AssistantTalk(ctx context.Context, connection connections.ConnectionConfig, opts ...grpc.CallOption) (web_api.TalkService_AssistantTalkClient, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.AssistantTalk(connection.WithAuth(ctx))
}

func GetAllAssistantConversation(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantConversationRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantConversationResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantConversation(connection.WithAuth(ctx), req)
}

func GetAllConversationMessage(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllConversationMessageRequest, opts ...grpc.CallOption) (*web_api.GetAllConversationMessageResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllConversationMessage(connection.WithAuth(ctx), req)
}

func CreateMessageMetric(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateMessageMetricRequest, opts ...grpc.CallOption) (*web_api.CreateMessageMetricResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateMessageMetric(connection.WithAuth(ctx), req)
}

func CreateConversationMetric(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateConversationMetricRequest, opts ...grpc.CallOption) (*web_api.CreateConversationMetricResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateConversationMetric(connection.WithAuth(ctx), req)
}
