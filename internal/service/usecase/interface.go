package usecase

import (
	"context"

	"github.com/AlexandrKobalt/go-files/pkg/proto"
)

type IUseCase interface {
	SaveFile(
		ctx context.Context,
		params *proto.SaveFileRequest,
	) (result *proto.SaveFileResponse, err error)
	GetPublicURL(
		ctx context.Context,
		params *proto.GetPublicURLRequest,
	) (result *proto.GetPublicURLResponse, err error)
}
