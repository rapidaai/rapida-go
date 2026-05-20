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

func GetAllAssistantProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantProviderRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantProvider(connection.WithAuth(ctx), in, opts...)
}

func CreateAssistantProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantProviderRequest, opts ...grpc.CallOption) (*protos.GetAssistantProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantProvider(connection.WithAuth(ctx), in, opts...)
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

func GetAllAssistantTelemetry(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantTelemetryRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantTelemetryResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTelemetry(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantTelemetryProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantTelemetryProviderRequest, opts ...grpc.CallOption) (*protos.GetAssistantTelemetryProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantTelemetryProvider(connection.WithAuth(ctx), in, opts...)
}

func GetAllAssistantTelemetryProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantTelemetryProviderRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantTelemetryProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantTelemetryProvider(connection.WithAuth(ctx), in, opts...)
}

func CreateAssistantTelemetryProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantTelemetryProviderRequest, opts ...grpc.CallOption) (*protos.GetAssistantTelemetryProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantTelemetryProvider(connection.WithAuth(ctx), in, opts...)
}

func UpdateAssistantTelemetryProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.UpdateAssistantTelemetryProviderRequest, opts ...grpc.CallOption) (*protos.GetAssistantTelemetryProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateAssistantTelemetryProvider(connection.WithAuth(ctx), in, opts...)
}

func DeleteAssistantTelemetryProvider(ctx context.Context, connection connections.ConnectionConfig, in *protos.DeleteAssistantTelemetryProviderRequest, opts ...grpc.CallOption) (*protos.GetAssistantTelemetryProviderResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteAssistantTelemetryProvider(connection.WithAuth(ctx), in, opts...)
}

func CreateAssistantAuthentication(ctx context.Context, connection connections.ConnectionConfig, in *protos.CreateAssistantAuthenticationRequest, opts ...grpc.CallOption) (*protos.GetAssistantAuthenticationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantAuthentication(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantAuthentication(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantAuthenticationRequest, opts ...grpc.CallOption) (*protos.GetAssistantAuthenticationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantAuthentication(connection.WithAuth(ctx), in, opts...)
}

func DisableAssistantAuthentication(ctx context.Context, connection connections.ConnectionConfig, in *protos.DisableAssistantAuthenticationRequest, opts ...grpc.CallOption) (*protos.GetAssistantAuthenticationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantAuthentication(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantConversation(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantConversationRequest, opts ...grpc.CallOption) (*protos.GetAssistantConversationResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantConversation(connection.WithAuth(ctx), in, opts...)
}

func GetAssistantHTTPLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantHTTPLogRequest, opts ...grpc.CallOption) (*protos.GetAssistantHTTPLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantHTTPLog(connection.WithAuth(ctx), in, opts...)
}

func GetAllAssistantHTTPLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantHTTPLogRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantHTTPLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantHTTPLog(connection.WithAuth(ctx), in, opts...)
}

func RetryAssistantHTTPLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.RetryAssistantHTTPLogRequest, opts ...grpc.CallOption) (*protos.GetAssistantHTTPLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.RetryAssistantHTTPLog(connection.WithAuth(ctx), in, opts...)
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

func GetAssistantToolLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAssistantToolLogRequest, opts ...grpc.CallOption) (*protos.GetAssistantToolLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantToolLog(connection.WithAuth(ctx), in, opts...)
}

func GetAllAssistantToolLog(ctx context.Context, connection connections.ConnectionConfig, in *protos.GetAllAssistantToolLogRequest, opts ...grpc.CallOption) (*protos.GetAllAssistantToolLogResponse, error) {
	c, err := connection.AssistantServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantToolLog(connection.WithAuth(ctx), in, opts...)
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
