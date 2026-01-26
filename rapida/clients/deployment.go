// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func CreateAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest) (*web_api.GetAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantApiDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest) (*web_api.GetAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantApiDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest) (*web_api.GetAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebpluginDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest) (*web_api.GetAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebpluginDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest) (*web_api.GetAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantDebuggerDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest) (*web_api.GetAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantDebuggerDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest) (*web_api.GetAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWhatsappDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest) (*web_api.GetAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWhatsappDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest) (*web_api.GetAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantPhoneDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest) (*web_api.GetAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantPhoneDeployment(connection.WithAuth(ctx), req)
}
