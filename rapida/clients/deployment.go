/*
 *  Copyright (c) 2024. Rapida
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 *
 *  Author: Prashant <prashant@rapida.ai>
 *
 *  This module provides functions for accessing audit logs via gRPC. It includes
 *  operations for retrieving lists of audit logs and fetching specific audit log entries.
 */
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
