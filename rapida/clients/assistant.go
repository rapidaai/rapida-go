// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func GetAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistant(connection.WithAuth(ctx), in)
}
func GetAllAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantRequest) (*web_api.GetAllAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistant(connection.WithAuth(ctx), in)
}
func CreateAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistant(connection.WithAuth(ctx), in)
}
func DeleteAssistant(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistant(connection.WithAuth(ctx), in)
}
func GetAllAssistantProviderModel(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantProviderModelRequest) (*web_api.GetAllAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantProviderModel(connection.WithAuth(ctx), in)
}
func CreateAssistantProviderModel(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantProviderModelRequest) (*web_api.GetAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantProviderModel(connection.WithAuth(ctx), in)
}
func CreateAssistantTag(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantTagRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTag(connection.WithAuth(ctx), in)
}
func UpdateAssistantVersion(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantVersionRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantVersion(connection.WithAuth(ctx), in)
}
func UpdateAssistantDetail(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantDetailRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantDetail(connection.WithAuth(ctx), in)
}
func GetAllAssistantMessage(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantMessageRequest) (*web_api.GetAllAssistantMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantMessage(connection.WithAuth(ctx), in)
}
func GetAllMessage(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllMessageRequest) (*web_api.GetAllMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllMessage(connection.WithAuth(ctx), in)
}

func GetAssistantConversation(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantConversationRequest) (*web_api.GetAssistantConversationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantConversation(connection.WithAuth(ctx), in)
}
func GetAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantWebhookLogRequest) (*web_api.GetAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhookLog(connection.WithAuth(ctx), in)
}
func GetAllAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantWebhookLogRequest) (*web_api.GetAllAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhookLog(connection.WithAuth(ctx), in)
}
func GetAllAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantWebhookRequest) (*web_api.GetAllAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhook(connection.WithAuth(ctx), in)
}
func GetAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhook(connection.WithAuth(ctx), in)
}
func CreateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebhook(connection.WithAuth(ctx), in)
}
func UpdateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantWebhook(connection.WithAuth(ctx), in)
}
func DeleteAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantWebhook(connection.WithAuth(ctx), in)
}
func GetAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantAnalysis(connection.WithAuth(ctx), in)
}
func UpdateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantAnalysis(connection.WithAuth(ctx), in)
}
func CreateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantAnalysis(connection.WithAuth(ctx), in)
}
func DeleteAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantAnalysis(connection.WithAuth(ctx), in)
}
func GetAllAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantAnalysisRequest) (*web_api.GetAllAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantAnalysis(connection.WithAuth(ctx), in)
}
func CreateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTool(connection.WithAuth(ctx), in)
}
func GetAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantTool(connection.WithAuth(ctx), in)
}
func GetAllAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantToolRequest) (*web_api.GetAllAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTool(connection.WithAuth(ctx), in)
}
func DeleteAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantTool(connection.WithAuth(ctx), in)
}
func UpdateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantTool(connection.WithAuth(ctx), in)
}
func CreateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.CreateAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantKnowledge(connection.WithAuth(ctx), in)
}
func GetAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantKnowledge(connection.WithAuth(ctx), in)
}
func GetAllAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.GetAllAssistantKnowledgeRequest) (*web_api.GetAllAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantKnowledge(connection.WithAuth(ctx), in)
}
func DeleteAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.DeleteAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantKnowledge(connection.WithAuth(ctx), in)
}
func UpdateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *web_api.UpdateAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantKnowledge(connection.WithAuth(ctx), in)
}
