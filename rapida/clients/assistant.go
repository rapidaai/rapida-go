// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func GetAssistant(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistant(connection.WithAuth(ctx), in)
}
func GetAllAssistant(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantRequest) (*web_api.GetAllAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistant(connection.WithAuth(ctx), in)
}
func CreateAssistant(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistant(connection.WithAuth(ctx), in)
}
func DeleteAssistant(connection connections.ConnectionConfig, ctx context.Context, in *web_api.DeleteAssistantRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistant(connection.WithAuth(ctx), in)
}
func GetAllAssistantProviderModel(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantProviderModelRequest) (*web_api.GetAllAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantProviderModel(connection.WithAuth(ctx), in)
}
func CreateAssistantProviderModel(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantProviderModelRequest) (*web_api.GetAssistantProviderModelResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantProviderModel(connection.WithAuth(ctx), in)
}
func CreateAssistantTag(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantTagRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTag(connection.WithAuth(ctx), in)
}
func UpdateAssistantVersion(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantVersionRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantVersion(connection.WithAuth(ctx), in)
}
func UpdateAssistantDetail(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantDetailRequest) (*web_api.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantDetail(connection.WithAuth(ctx), in)
}
func GetAllAssistantMessage(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantMessageRequest) (*web_api.GetAllAssistantMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantMessage(connection.WithAuth(ctx), in)
}
func GetAllMessage(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllMessageRequest) (*web_api.GetAllMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllMessage(connection.WithAuth(ctx), in)
}

func GetAssistantConversation(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantConversationRequest) (*web_api.GetAssistantConversationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantConversation(connection.WithAuth(ctx), in)
}
func GetAssistantWebhookLog(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantWebhookLogRequest) (*web_api.GetAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhookLog(connection.WithAuth(ctx), in)
}
func GetAllAssistantWebhookLog(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantWebhookLogRequest) (*web_api.GetAllAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhookLog(connection.WithAuth(ctx), in)
}
func GetAllAssistantWebhook(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantWebhookRequest) (*web_api.GetAllAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhook(connection.WithAuth(ctx), in)
}
func GetAssistantWebhook(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhook(connection.WithAuth(ctx), in)
}
func CreateAssistantWebhook(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebhook(connection.WithAuth(ctx), in)
}
func UpdateAssistantWebhook(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantWebhook(connection.WithAuth(ctx), in)
}
func DeleteAssistantWebhook(connection connections.ConnectionConfig, ctx context.Context, in *web_api.DeleteAssistantWebhookRequest) (*web_api.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantWebhook(connection.WithAuth(ctx), in)
}
func GetAssistantAnalysis(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantAnalysis(connection.WithAuth(ctx), in)
}
func UpdateAssistantAnalysis(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantAnalysis(connection.WithAuth(ctx), in)
}
func CreateAssistantAnalysis(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantAnalysis(connection.WithAuth(ctx), in)
}
func DeleteAssistantAnalysis(connection connections.ConnectionConfig, ctx context.Context, in *web_api.DeleteAssistantAnalysisRequest) (*web_api.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantAnalysis(connection.WithAuth(ctx), in)
}
func GetAllAssistantAnalysis(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantAnalysisRequest) (*web_api.GetAllAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantAnalysis(connection.WithAuth(ctx), in)
}
func CreateAssistantTool(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTool(connection.WithAuth(ctx), in)
}
func GetAssistantTool(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantTool(connection.WithAuth(ctx), in)
}
func GetAllAssistantTool(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantToolRequest) (*web_api.GetAllAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTool(connection.WithAuth(ctx), in)
}
func DeleteAssistantTool(connection connections.ConnectionConfig, ctx context.Context, in *web_api.DeleteAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantTool(connection.WithAuth(ctx), in)
}
func UpdateAssistantTool(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantToolRequest) (*web_api.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantTool(connection.WithAuth(ctx), in)
}
func CreateAssistantKnowledge(connection connections.ConnectionConfig, ctx context.Context, in *web_api.CreateAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantKnowledge(connection.WithAuth(ctx), in)
}
func GetAssistantKnowledge(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantKnowledge(connection.WithAuth(ctx), in)
}
func GetAllAssistantKnowledge(connection connections.ConnectionConfig, ctx context.Context, in *web_api.GetAllAssistantKnowledgeRequest) (*web_api.GetAllAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantKnowledge(connection.WithAuth(ctx), in)
}
func DeleteAssistantKnowledge(connection connections.ConnectionConfig, ctx context.Context, in *web_api.DeleteAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantKnowledge(connection.WithAuth(ctx), in)
}
func UpdateAssistantKnowledge(connection connections.ConnectionConfig, ctx context.Context, in *web_api.UpdateAssistantKnowledgeRequest) (*web_api.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantKnowledge(connection.WithAuth(ctx), in)
}
