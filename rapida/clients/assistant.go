// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GetAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistant(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistant(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistant(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistant(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantProviderModel(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantProviderModelRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantProviderModel(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantProviderModel(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantProviderModelRequest, opts ...grpc.CallOption) (*web_api.GetAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantProviderModel(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantTag(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantTagRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTag(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantVersion(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantVersionRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantVersion(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantDetail(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantDetailRequest, opts ...grpc.CallOption) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantDetail(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantMessage(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantMessageRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantMessage(connection.WithAuth(ctx), in, opts...)
}
func GetAllMessage(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllMessageRequest, opts ...grpc.CallOption) (*web_api.GetAllMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllMessage(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantConversation(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantConversationRequest, opts ...grpc.CallOption) (*web_api.GetAssistantConversationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantConversation(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantWebhookLogRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhookLog(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantWebhookLogRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhookLog(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantWebhookRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantWebhookRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantWebhookRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantWebhookRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantWebhookRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantAnalysisRequest, opts ...grpc.CallOption) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantAnalysisRequest, opts ...grpc.CallOption) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantAnalysisRequest, opts ...grpc.CallOption) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantAnalysisRequest, opts ...grpc.CallOption) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantAnalysisRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantToolRequest, opts ...grpc.CallOption) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantToolRequest, opts ...grpc.CallOption) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantToolRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantToolRequest, opts ...grpc.CallOption) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantToolRequest, opts ...grpc.CallOption) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
