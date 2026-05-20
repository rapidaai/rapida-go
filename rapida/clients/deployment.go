// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func CreateAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantApiDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantApiDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAllAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantApiDeployment(connection.WithAuth(ctx), req, opts...)
}

func DisableAssistantApiDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantApiDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantApiDeployment(connection.WithAuth(ctx), req, opts...)
}

func CreateAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWebpluginDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWebpluginDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAllAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWebpluginDeployment(connection.WithAuth(ctx), req, opts...)
}

func DisableAssistantWebpluginDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWebpluginDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantWebpluginDeployment(connection.WithAuth(ctx), req, opts...)
}

func CreateAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantDebuggerDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantDebuggerDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAllAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantDebuggerDeployment(connection.WithAuth(ctx), req, opts...)
}

func DisableAssistantDebuggerDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantDebuggerDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantDebuggerDeployment(connection.WithAuth(ctx), req, opts...)
}

func CreateAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantWhatsappDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantWhatsappDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAllAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantWhatsappDeployment(connection.WithAuth(ctx), req, opts...)
}

func DisableAssistantWhatsappDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantWhatsappDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantWhatsappDeployment(connection.WithAuth(ctx), req, opts...)
}

func CreateAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateAssistantPhoneDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAssistantPhoneDeployment(connection.WithAuth(ctx), req, opts...)
}

func GetAllAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAllAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllAssistantPhoneDeployment(connection.WithAuth(ctx), req, opts...)
}

func DisableAssistantPhoneDeployment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAssistantDeploymentRequest, opts ...grpc.CallOption) (*web_api.GetAssistantPhoneDeploymentResponse, error) {
	c, err := connection.AssistantDeploymentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DisableAssistantPhoneDeployment(connection.WithAuth(ctx), req, opts...)
}
