// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package connections

import (
	"context"
	"crypto/tls"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/configs"
	"github.com/rapidaai/rapida-go/rapida/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type ConnectionConfig interface {
	WithLocal() ConnectionConfig
	WithInsecureConnection() ConnectionConfig
	WithCustomEndpoint(endpoint map[string]string) ConnectionConfig
	// all the clients
	TalkServiceClient() (web_api.TalkServiceClient, error)
	AuthenticationServiceClient() (web_api.AuthenticationServiceClient, error)
	AssistantServiceClient() (web_api.AssistantServiceClient, error)
	AssistantDeploymentServiceClient() (web_api.AssistantDeploymentServiceClient, error)
	DeploymentClient() (web_api.DeploymentClient, error)
	EndpointServiceClient() (web_api.EndpointServiceClient, error)
	VaultServiceClient() (web_api.VaultServiceClient, error)
	// authenticaton
	WithAuth(ctx context.Context) context.Context
}

type RapidaCredential map[string]string

type connectionConfig struct {
	endpoint struct {
		assistant string
		web       string
		endpoint  string
	}
	insecure bool
	auth     map[string]string
}

func NewConnectionConfig(auth RapidaCredential, endpoint map[string]string, insecure bool) ConnectionConfig {
	cc := &connectionConfig{
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

func (cc *connectionConfig) WithCustomEndpoint(endpoint map[string]string) ConnectionConfig {
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

func (cc *connectionConfig) WithInsecureConnection() ConnectionConfig {
	cc.insecure = true
	return cc
}

func (cc *connectionConfig) GetGrpcOption() []grpc.DialOption {
	grpcOpts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
	}
	if cc.insecure {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	}
	return grpcOpts
}

// Example of one client getter method
func (cc *connectionConfig) TalkServiceClient() (web_api.TalkServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewTalkServiceClient(conn), nil
}

func (cc *connectionConfig) AuthenticationServiceClient() (web_api.AuthenticationServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.web, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewAuthenticationServiceClient(conn), nil
}

func (cc *connectionConfig) AssistantServiceClient() (web_api.AssistantServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewAssistantServiceClient(conn), nil
}

func (cc *connectionConfig) EndpointServiceClient() (web_api.EndpointServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.endpoint, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewEndpointServiceClient(conn), nil
}

func (cc *connectionConfig) AssistantDeploymentServiceClient() (web_api.AssistantDeploymentServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.assistant, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewAssistantDeploymentServiceClient(conn), nil
}

func (cc *connectionConfig) DeploymentClient() (web_api.DeploymentClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.endpoint, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewDeploymentClient(conn), nil
}

func (cc *connectionConfig) VaultServiceClient() (web_api.VaultServiceClient, error) {
	conn, err := grpc.NewClient(cc.endpoint.web, cc.GetGrpcOption()...)
	if err != nil {
		return nil, err
	}
	return web_api.NewVaultServiceClient(conn), nil
}

func (cc *connectionConfig) WithAuth(ctx context.Context) context.Context {
	md := metadata.New(cc.auth)
	return metadata.NewOutgoingContext(ctx, md)
}

// Add similar methods for other clients...

func (cc *connectionConfig) WithLocal() ConnectionConfig {
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

func WithSDK(apiKey ...string) RapidaCredential {
	if len(apiKey) < 1 {
		panic("Rapida WithSDK requires at least one API key")
	}
	credential := RapidaCredential{
		utils.HEADER_API_KEY:    apiKey[0],
		utils.HEADER_SOURCE_KEY: utils.SDK.Get(),
	}
	if len(apiKey) > 1 {
		credential[utils.HEADER_AUTH_KEY] = apiKey[1]
	}
	return credential
}

func DefaultconnectionConfig(credential RapidaCredential) ConnectionConfig {
	return NewConnectionConfig(credential, nil, false)
}
