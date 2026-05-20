// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package clients

import (
	"context"

	web_api "github.com/rapidaai/rapida-go/protos"
	"github.com/rapidaai/rapida-go/rapida/connections"
	"google.golang.org/grpc"
)

func GetKnowledge(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetKnowledgeResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetKnowledge(connection.WithAuth(ctx), req, opts...)
}

func GetAllKnowledge(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllKnowledgeRequest, opts ...grpc.CallOption) (*web_api.GetAllKnowledgeResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllKnowledge(connection.WithAuth(ctx), req, opts...)
}

func CreateKnowledge(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateKnowledgeRequest, opts ...grpc.CallOption) (*web_api.CreateKnowledgeResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateKnowledge(connection.WithAuth(ctx), req, opts...)
}

func CreateKnowledgeTag(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateKnowledgeTagRequest, opts ...grpc.CallOption) (*web_api.GetKnowledgeResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateKnowledgeTag(connection.WithAuth(ctx), req, opts...)
}

func CreateKnowledgeDocument(ctx context.Context, connection connections.ConnectionConfig, req *web_api.CreateKnowledgeDocumentRequest, opts ...grpc.CallOption) (*web_api.CreateKnowledgeDocumentResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.CreateKnowledgeDocument(connection.WithAuth(ctx), req, opts...)
}

func GetAllKnowledgeDocument(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllKnowledgeDocumentRequest, opts ...grpc.CallOption) (*web_api.GetAllKnowledgeDocumentResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllKnowledgeDocument(connection.WithAuth(ctx), req, opts...)
}

func GetAllKnowledgeDocumentSegment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllKnowledgeDocumentSegmentRequest, opts ...grpc.CallOption) (*web_api.GetAllKnowledgeDocumentSegmentResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllKnowledgeDocumentSegment(connection.WithAuth(ctx), req, opts...)
}

func UpdateKnowledgeDetail(ctx context.Context, connection connections.ConnectionConfig, req *web_api.UpdateKnowledgeDetailRequest, opts ...grpc.CallOption) (*web_api.GetKnowledgeResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateKnowledgeDetail(connection.WithAuth(ctx), req, opts...)
}

func UpdateKnowledgeDocumentSegment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.UpdateKnowledgeDocumentSegmentRequest, opts ...grpc.CallOption) (*web_api.BaseResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.UpdateKnowledgeDocumentSegment(connection.WithAuth(ctx), req, opts...)
}

func DeleteKnowledgeDocumentSegment(ctx context.Context, connection connections.ConnectionConfig, req *web_api.DeleteKnowledgeDocumentSegmentRequest, opts ...grpc.CallOption) (*web_api.BaseResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.DeleteKnowledgeDocumentSegment(connection.WithAuth(ctx), req, opts...)
}

func GetAllKnowledgeLog(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetAllKnowledgeLogRequest, opts ...grpc.CallOption) (*web_api.GetAllKnowledgeLogResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllKnowledgeLog(connection.WithAuth(ctx), req, opts...)
}

func GetKnowledgeLog(ctx context.Context, connection connections.ConnectionConfig, req *web_api.GetKnowledgeLogRequest, opts ...grpc.CallOption) (*web_api.GetKnowledgeLogResponse, error) {
	c, err := connection.KnowledgeServiceClient()
	if err != nil {
		return nil, err
	}
	return c.GetKnowledgeLog(connection.WithAuth(ctx), req, opts...)
}
