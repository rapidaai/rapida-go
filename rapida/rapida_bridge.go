package rapida

import (
	"context"
	"errors"

	lexatic_backend "github.com/rapidaai/rapida-go/rapida/clients/protos"
	rapida_constants "github.com/rapidaai/rapida-go/rapida/constants"
	rapida_definitions "github.com/rapidaai/rapida-go/rapida/definitions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/anypb"
)

type rapidaBridge struct {
	opts              *RapidaClientOption
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

	NewTalkerClient(ctx context.Context,
		assistant rapida_definitions.AssistantDefinition,
		metadata map[string]*anypb.Any,
		options map[string]*anypb.Any,
	) (TalkClient, error)

	SendMessage(ctx context.Context,
		assistant rapida_definitions.AssistantDefinition,
		message *lexatic_backend.Message,
		metadata map[string]*anypb.Any,
		options map[string]*anypb.Any,
	) (MessageClient, error)
}

func NewRapidaBridge(options *RapidaClientOption) (RapidaBridge, error) {
	grpcOpts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(rapida_constants.MaxRecvMsgSize),
			grpc.MaxCallSendMsgSize(rapida_constants.MaxSendMsgSize),
		),
	}

	if options.IsSecure {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	endpointConnection, err := grpc.NewClient(*options.GetRapidaEndpointUrl(),
		grpcOpts...)

	if err != nil {
		return nil, err
	}

	assistantConnection, err := grpc.NewClient(*options.GetAssistantUrl(),
		grpcOpts...)

	if err != nil {
		return nil, err
	}

	return &rapidaBridge{
		opts:              options,
		deploymentClient:  lexatic_backend.NewDeploymentClient(endpointConnection),
		talkServiceClient: lexatic_backend.NewTalkServiceClient(assistantConnection),
	}, nil
}

func (ic *rapidaBridge) WithScopeToken(c context.Context) context.Context {
	md := metadata.New(map[string]string{
		rapida_constants.HEADER_API_KEY:         *ic.opts.GetRapidaApiKey(),
		rapida_constants.HEADER_SOURCE_KEY:      lexatic_backend.Source_name[int32(lexatic_backend.Source_GO_SDK)],
		rapida_constants.HEADER_ENVIRONMENT_KEY: ic.opts.GetRapidaEnvironment().Get(),
		rapida_constants.HEADER_REGION_KEY:      ic.opts.GetRapidaRegion().Get(),
	})
	return metadata.NewOutgoingContext(c, md)
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
	invokeResponse, err := rb.deploymentClient.Invoke(rb.WithScopeToken(ctx), invokeRequest)
	if err != nil {
		return nil, err
	}
	return rapida_definitions.NewInvokeResponseWrapper(invokeResponse), nil
}

type MessageClient interface {
	Recv() (*lexatic_backend.AssistantMessagingResponse, error)
}

type messageClient struct {
	client lexatic_backend.TalkService_AssistantMessagingClient
}

func NewMessageClient(client lexatic_backend.TalkService_AssistantMessagingClient) MessageClient {
	return &messageClient{
		client: client,
	}
}

func (mc *messageClient) Recv() (*lexatic_backend.AssistantMessagingResponse, error) {
	return mc.client.Recv()
}

// CreateAssistantMessage
func (rb *rapidaBridge) SendMessage(ctx context.Context,
	assistant rapida_definitions.AssistantDefinition,
	message *lexatic_backend.Message,
	metadata map[string]*anypb.Any,
	options map[string]*anypb.Any,
) (MessageClient, error) {

	in := &lexatic_backend.AssistantMessagingRequest{
		Assistant: &lexatic_backend.AssistantDefinition{
			AssistantId: assistant.GetAssistant(),
			Version:     assistant.GetAssistantVersion(),
		},
		Metadata: metadata,
		Options:  options,
		Message:  message,
	}
	cl, err := rb.talkServiceClient.AssistantMessaging(rb.WithScopeToken(ctx), in)
	if err != nil {
		return nil, err
	}
	return NewMessageClient(cl), nil
}

type TalkClient interface {
	Send(msg *lexatic_backend.Message) error
	Recv() (*lexatic_backend.Message, error)
}

type assistantBiDirectionalClient struct {
	client                   lexatic_backend.TalkService_AssistantTalkClient
	assistantConversactionId *uint64
	assistant                *lexatic_backend.AssistantDefinition
	metadata                 map[string]*anypb.Any
	options                  map[string]*anypb.Any
}

func NewAssistantBiDirectionalClient(
	client lexatic_backend.TalkService_AssistantTalkClient,
	assistant rapida_definitions.AssistantDefinition,
	metadata map[string]*anypb.Any,
	options map[string]*anypb.Any) TalkClient {
	return &assistantBiDirectionalClient{
		client: client,
		assistant: &lexatic_backend.AssistantDefinition{
			AssistantId: assistant.GetAssistant(),
			Version:     assistant.GetAssistantVersion(),
		},
		metadata: metadata,
		options:  options,
	}

}

func (abdc *assistantBiDirectionalClient) Send(msg *lexatic_backend.Message) error {
	in := &lexatic_backend.AssistantMessagingRequest{
		Assistant:               abdc.assistant,
		Metadata:                abdc.metadata,
		Options:                 abdc.options,
		Message:                 msg,
		AssistantConversationId: abdc.assistantConversactionId,
	}
	return abdc.client.Send(in)
}

func (abdc *assistantBiDirectionalClient) Recv() (*lexatic_backend.Message, error) {
	at, err := abdc.client.Recv()
	if err != nil {
		return nil, err
	}
	if at.GetSuccess() {
		code := at.GetData()
		abdc.assistantConversactionId = &code.AssistantConversationId
		return at.GetData().GetResponse(), nil
	}
	return nil, errors.New(at.GetError().GetHumanMessage())
}

func (rb *rapidaBridge) NewTalkerClient(ctx context.Context,
	assistant rapida_definitions.AssistantDefinition,
	metadata map[string]*anypb.Any,
	options map[string]*anypb.Any,
) (TalkClient, error) {
	client, err := rb.talkServiceClient.AssistantTalk(rb.WithScopeToken(ctx))
	if err != nil {
		return nil, err
	}
	return NewAssistantBiDirectionalClient(client, assistant, metadata, options), nil
}
