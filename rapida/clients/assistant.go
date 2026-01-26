// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>
package clients

import (
	"context"

	"github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GetAssistant(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistant(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistant(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistant(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistant(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistant(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistant(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistant(connection.WithAuth(ctx), in, opts...)
}

func CreateAssistantTag(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantTagRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTag(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantVersion(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantVersionRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantVersion(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantDetail(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantDetailRequest, opts ...grpc.CallOption) (*protos.GetAssistantResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantDetail(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantMessage(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantMessageRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantMessage(connection.WithAuth(ctx), in, opts...)
}
func GetAllMessage(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllMessageRequest, opts ...grpc.CallOption) (*protos.GetAllMessageResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllMessage(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantConversation(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantConversationRequest, opts ...grpc.CallOption) (*protos.GetAssistantConversationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantConversation(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantWebhookLogRequest, opts ...grpc.CallOption) (*protos.GetAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhookLog(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantWebhookLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantWebhookLogRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantWebhookLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhookLog(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantWebhookRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantWebhookRequest, opts ...grpc.CallOption) (*protos.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantWebhookRequest, opts ...grpc.CallOption) (*protos.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantWebhookRequest, opts ...grpc.CallOption) (*protos.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantWebhook(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantWebhookRequest, opts ...grpc.CallOption) (*protos.GetAssistantWebhookResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantWebhook(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantAnalysisRequest, opts ...grpc.CallOption) (*protos.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantAnalysisRequest, opts ...grpc.CallOption) (*protos.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantAnalysisRequest, opts ...grpc.CallOption) (*protos.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantAnalysisRequest, opts ...grpc.CallOption) (*protos.GetAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantAnalysis(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantAnalysisRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantAnalysisResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantAnalysis(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantToolRequest, opts ...grpc.CallOption) (*protos.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantToolRequest, opts ...grpc.CallOption) (*protos.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantToolRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantToolRequest, opts ...grpc.CallOption) (*protos.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantTool(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantToolRequest, opts ...grpc.CallOption) (*protos.GetAssistantToolResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantTool(connection.WithAuth(ctx), in, opts...)
}
func CreateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantKnowledgeRequest, opts ...grpc.CallOption) (*protos.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func GetAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantKnowledgeRequest, opts ...grpc.CallOption) (*protos.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func GetAllAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantKnowledgeRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func DeleteAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantKnowledgeRequest, opts ...grpc.CallOption) (*protos.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
func UpdateAssistantKnowledge(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantKnowledgeRequest, opts ...grpc.CallOption) (*protos.GetAssistantKnowledgeResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantKnowledge(connection.WithAuth(ctx), in, opts...)
}
