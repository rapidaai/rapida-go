// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	lexatic_backend "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func CreateProviderCredential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateProviderCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateProviderCredential(connectionConfig.WithAuth(ctx), in, opts...)
}

func GetAllOrganizationCredential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetAllOrganizationCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetAllOrganizationCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}

	return client.GetAllOrganizationCredential(connectionConfig.WithAuth(ctx), in, opts...)
}
func DeleteCredential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.DeleteCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}

	return client.DeleteCredential(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetProviderCredential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}

	return client.GetCredential(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetCredential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}

	return client.GetCredential(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetOauth2Credential(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetCredentialRequest, opts ...grpc.CallOption) (*lexatic_backend.GetCredentialResponse, error) {
	client, err := connectionConfig.VaultServiceClient()
	if err != nil {
		return nil, err
	}

	return client.GetOauth2Credential(connectionConfig.WithAuth(ctx), in, opts...)
}
