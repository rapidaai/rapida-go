// Copyright (c) 2024 Rapida
// Licensed under the MIT License. See LICENSE file for details.
package connections

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/rapida/clients/protos"
	"github.com/rapidaai/rapida-go/rapida/configs"
	"github.com/rapidaai/rapida-go/rapida/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type RapidaCredential map[string]string

type ConnectionConfig struct {
	endpoint struct {
		assistant string
		web       string
		endpoint  string
	}
	insecure bool
	auth     map[string]string
}

func NewConnectionConfig(auth RapidaCredential, endpoint map[string]string, insecure bool) *ConnectionConfig {
	cc := &ConnectionConfig{
		auth:     auth,
		insecure: insecure,
	}
	cc.endpoint.assistant = configs.ASSISTANT_API
	cc.endpoint.web = configs.WEB_API
	cc.endpoint.endpoint = configs.ENDPOINT_API

	if endpoint != nil {
		if assistant, ok := endpoint["assistant"]; ok {
			cc.endpoint.assistant = assistant
		}
		if web, ok := endpoint["web"]; ok {
			cc.endpoint.web = web
		}
		if endpt, ok := endpoint["endpoint"]; ok {
			cc.endpoint.endpoint = endpt
		}
	}
	return cc
}

func (cc *ConnectionConfig) WithCustomEndpoint(endpoint map[string]string) *ConnectionConfig {
	if endpoint != nil {
		if assistant, ok := endpoint["assistant"]; ok {
			cc.endpoint.assistant = assistant
		}
		if web, ok := endpoint["web"]; ok {
			cc.endpoint.web = web
		}
		if endpt, ok := endpoint["endpoint"]; ok {
			cc.endpoint.endpoint = endpt
		}
	}
	return cc
}

func (cc *ConnectionConfig) WithInsecureConnection() *ConnectionConfig {
	cc.insecure = true
	return cc
}

// Example of one client getter method
func (cc *ConnectionConfig) TalkServiceClient() (web_api.TalkServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return web_api.NewTalkServiceClient(conn), nil
}

func (cc *ConnectionConfig) AuthenticationServiceClient() (web_api.AuthenticationServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.web, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return web_api.NewAuthenticationServiceClient(conn), nil
}

func (cc *ConnectionConfig) AssistantServiceClient() (web_api.AssistantServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return web_api.NewAssistantServiceClient(conn), nil
}

func (cc *ConnectionConfig) AssistantDeploymentServiceClient() (web_api.AssistantDeploymentServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return web_api.NewAssistantDeploymentServiceClient(conn), nil
}

func (cc *ConnectionConfig) DeploymentClient() (web_api.DeploymentClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return web_api.NewDeploymentClient(conn), nil
}

func (cc *ConnectionConfig) WithAuth(ctx context.Context) context.Context {
	md := metadata.New(cc.auth)
	return metadata.NewOutgoingContext(ctx, md)
}

// Add similar methods for other clients...

func (cc *ConnectionConfig) WithLocal() *ConnectionConfig {
	cc.WithCustomEndpoint(map[string]string{
		"assistant": configs.LOCAL_ASSISTANT_API,
		"web":       configs.LOCAL_WEB_API,
		"endpoint":  configs.LOCAL_ENDPOINT_API,
	})
	return cc.WithInsecureConnection()
}

func WithPersonalToken(authorization, userId, projectId string) RapidaCredential {
	return RapidaCredential{
		utils.HEADER_AUTHORIZATION: authorization,
		utils.HEADER_PROJECT_ID:    projectId,
		utils.HEADER_AUTH_KEY:      userId,
		utils.HEADER_SOURCE_KEY:    utils.SDK.Get(),
	}
}

func WithWebpluginClient(apiKey string, userId string) RapidaCredential {
	return RapidaCredential{
		utils.HEADER_API_KEY:    apiKey,
		utils.HEADER_AUTH_KEY:   userId,
		utils.HEADER_SOURCE_KEY: utils.SDK.Get(),
	}
}

func WithSDK(apiKey string, userId string) RapidaCredential {
	return RapidaCredential{
		utils.HEADER_API_KEY:    apiKey,
		utils.HEADER_AUTH_KEY:   userId,
		utils.HEADER_SOURCE_KEY: utils.SDK.Get(),
	}
}
