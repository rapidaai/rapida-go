// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	lexatic_backend "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GetEndpoint(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetEndpointRequest, opts ...grpc.CallOption) (*lexatic_backend.GetEndpointResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.GetEndpoint(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetAllEndpoint(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetAllEndpointRequest, opts ...grpc.CallOption) (*lexatic_backend.GetAllEndpointResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.GetAllEndpoint(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetAllEndpointProviderModel(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetAllEndpointProviderModelRequest, opts ...grpc.CallOption) (*lexatic_backend.GetAllEndpointProviderModelResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.GetAllEndpointProviderModel(connectionConfig.WithAuth(ctx), in, opts...)
}
func UpdateEndpointVersion(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.UpdateEndpointVersionRequest, opts ...grpc.CallOption) (*lexatic_backend.UpdateEndpointVersionResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.UpdateEndpointVersion(connectionConfig.WithAuth(ctx), in, opts...)
}
func CreateEndpoint(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateEndpointRequest, opts ...grpc.CallOption) (*lexatic_backend.CreateEndpointResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateEndpoint(connectionConfig.WithAuth(ctx), in, opts...)
}
func CreateEndpointProviderModel(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateEndpointProviderModelRequest, opts ...grpc.CallOption) (*lexatic_backend.CreateEndpointProviderModelResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointProviderModel(connectionConfig.WithAuth(ctx), in, opts...)
}
func CreateEndpointCacheConfiguration(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateEndpointCacheConfigurationRequest, opts ...grpc.CallOption) (*lexatic_backend.CreateEndpointCacheConfigurationResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointCacheConfiguration(connectionConfig.WithAuth(ctx), in, opts...)
}
func CreateEndpointRetryConfiguration(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateEndpointRetryConfigurationRequest, opts ...grpc.CallOption) (*lexatic_backend.CreateEndpointRetryConfigurationResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointRetryConfiguration(connectionConfig.WithAuth(ctx), in, opts...)
}
func CreateEndpointTag(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.CreateEndpointTagRequest, opts ...grpc.CallOption) (*lexatic_backend.GetEndpointResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointTag(connectionConfig.WithAuth(ctx), in, opts...)
}
func ForkEndpoint(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.ForkEndpointRequest, opts ...grpc.CallOption) (*lexatic_backend.BaseResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.ForkEndpoint(connectionConfig.WithAuth(ctx), in, opts...)
}
func UpdateEndpointDetail(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.UpdateEndpointDetailRequest, opts ...grpc.CallOption) (*lexatic_backend.GetEndpointResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.UpdateEndpointDetail(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetAllEndpointLog(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetAllEndpointLogRequest, opts ...grpc.CallOption) (*lexatic_backend.GetAllEndpointLogResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.GetAllEndpointLog(connectionConfig.WithAuth(ctx), in, opts...)
}
func GetEndpointLog(ctx context.Context, connectionConfig connections.ConnectionConfig, in *lexatic_backend.GetEndpointLogRequest, opts ...grpc.CallOption) (*lexatic_backend.GetEndpointLogResponse, error) {

	client, err := connectionConfig.EndpointServiceClient()
	if err != nil {
		return nil, err
	}
	return client.GetEndpointLog(connectionConfig.WithAuth(ctx), in, opts...)
}
