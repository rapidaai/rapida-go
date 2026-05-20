// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func IndexKnowledgeDocument(ctx context.Context, connection connections.ConnectionConfig, req *web_api.IndexKnowledgeDocumentRequest, opts ...grpc.CallOption) (*web_api.IndexKnowledgeDocumentResponse, error) {
	c, err := connection.DocumentServiceClient()
	if err != nil {
		return nil, err
	}
	return c.IndexKnowledgeDocument(connection.WithAuth(ctx), req, opts...)
}
