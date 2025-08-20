// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
)

func CreateAssistantApiDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateAssistantApiDeploymentRequest) (*web_api.AssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantApiDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantApiDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAssistantDeploymentRequest) (*web_api.AssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantApiDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantWebpluginDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateAssistantWebpluginDeploymentRequest) (*web_api.AssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebpluginDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantWebpluginDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAssistantDeploymentRequest) (*web_api.AssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebpluginDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantDebuggerDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateAssistantDebuggerDeploymentRequest) (*web_api.AssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantDebuggerDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantDebuggerDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAssistantDeploymentRequest) (*web_api.AssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantDebuggerDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantWhatsappDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateAssistantWhatsappDeploymentRequest) (*web_api.AssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWhatsappDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantWhatsappDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAssistantDeploymentRequest) (*web_api.AssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWhatsappDeployment(connection.WithAuth(ctx), req)
}

func CreateAssistantPhoneDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.CreateAssistantPhoneDeploymentRequest) (*web_api.AssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantPhoneDeployment(connection.WithAuth(ctx), req)
}

func GetAssistantPhoneDeployment(connection connections.ConnectionConfig, ctx context.Context, req *web_api.GetAssistantDeploymentRequest) (*web_api.AssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantPhoneDeployment(connection.WithAuth(ctx), req)
}
