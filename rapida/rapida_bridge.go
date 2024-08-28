package rapida

import (
	"context"
	"errors"
	"fmt"

	lexatic_backend "github.com/rapidaai/rapida-sdk/rapida/clients/protos"
	rapida_constants "github.com/rapidaai/rapida-sdk/rapida/constants"
	rapida_definitions "github.com/rapidaai/rapida-sdk/rapida/definitions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

type rapidaBridge struct {
	deploymentClient  lexatic_backend.DeploymentClient
	talkServiceClient lexatic_backend.TalkServiceClient
}

type RapidaBridge interface {
	InvokeWithContext(
		ctx context.Context,
		endpoint rapida_definitions.EndpointDefinition,
		inputs map[string]*anypb.Any,
		metadata map[string]*anypb.Any,
		options map[string]*anypb.Any,
	) (*rapida_definitions.InvokeResponseWrapper, error)
}

func NewRapidaBridge(options *RapidaClientOption) (RapidaBridge, error) {
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(rapida_constants.MaxRecvMsgSize),
			grpc.MaxCallSendMsgSize(rapida_constants.MaxSendMsgSize),
		),
	}

	endpointConnection, err := grpc.NewClient(*options.GetRapidaEndpointUrl(),
		grpcOpts...)

	if err != nil {
		return nil, err
	}

	assistantConnection, err := grpc.NewClient(*options.GetAssistantUrl(),
		grpcOpts...)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to create connection %v", err))
	}

	return &rapidaBridge{
		deploymentClient:  lexatic_backend.NewDeploymentClient(endpointConnection),
		talkServiceClient: lexatic_backend.NewTalkServiceClient(assistantConnection),
	}, nil
}

func (rb *rapidaBridge) InvokeWithContext(
	ctx context.Context,
	endpoint rapida_definitions.EndpointDefinition,
	inputs map[string]*anypb.Any,
	metadata map[string]*anypb.Any,
	options map[string]*anypb.Any,
) (*rapida_definitions.InvokeResponseWrapper, error) {

	// Building the InvokeRequest
	invokeRequest := &lexatic_backend.InvokeRequest{
		Endpoint: &lexatic_backend.EndpointDefinition{
			EndpointId: endpoint.GetEndpoint(),
			Version:    endpoint.GetEndpointVersion(),
		},
		Args:     inputs,
		Metadata: metadata,
		Options:  options,
	}
	invokeResponse, err := rb.deploymentClient.Invoke(ctx, invokeRequest)
	if err != nil {
		return nil, err
	}

	return rapida_definitions.NewInvokeResponseWrapper(invokeResponse), nil
}
