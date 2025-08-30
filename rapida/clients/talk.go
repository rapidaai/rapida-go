// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func AssistantTalk(connection connections.ConnectionConfig, ctx context.Context) (web_api.TalkService_AssistantTalkClient, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.AssistantTalk(connection.WithAuth(ctx))
}

func GetAllAssistantConversation(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAllAssistantConversationRequest) (*web_api.GetAllAssistantConversationResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantConversation(connection.WithAuth(ctx), req)
}

func GetAllConversationMessage(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAllConversationMessageRequest) (*web_api.GetAllConversationMessageResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllConversationMessage(connection.WithAuth(ctx), req)
}

func CreateMessageMetric(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateMessageMetricRequest) (*web_api.CreateMessageMetricResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateMessageMetric(connection.WithAuth(ctx), req)
}

func CreateConversationMetric(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateConversationMetricRequest) (*web_api.CreateConversationMetricResponse, error) {
	c, err := connection.TalkServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateConversationMetric(connection.WithAuth(ctx), req)
}
