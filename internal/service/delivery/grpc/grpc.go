package grpc

import (
	"context"

	"github.com/AlexandrKobalt/go-files/internal/service/usecase"
	"github.com/AlexandrKobalt/go-files/pkg/proto"
)

type Server struct {
	uc usecase.IUseCase
	*proto.UnimplementedFileServiceServer
}

func New(uc usecase.IUseCase) *Server {
	return &Server{uc: uc}
}

func (s *Server) SaveFile(
	ctx context.Context,
	request *proto.SaveFileRequest,
) (response *proto.SaveFileResponse, err error) {
	return s.uc.SaveFile(ctx, request)
}

func (s *Server) GetPublicURL(
	ctx context.Context,
	request *proto.GetPublicURLRequest,
) (response *proto.GetPublicURLResponse, err error) {
	return s.uc.GetPublicURL(ctx, request)
}
