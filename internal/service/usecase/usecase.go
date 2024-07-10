package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AlexandrKobalt/go-files/config"
	"github.com/AlexandrKobalt/go-files/pkg/proto"
	"github.com/google/uuid"
)

type useCase struct {
	cfg *config.Config
}

func New(cfg *config.Config) IUseCase {
	if err := os.MkdirAll(cfg.Path, os.ModePerm); err != nil {
		log.Fatalf("could not create storage directory: %v", err)
	}

	return &useCase{cfg: cfg}
}

func (uc *useCase) SaveFile(
	ctx context.Context,
	params *proto.SaveFileRequest,
) (result *proto.SaveFileResponse, err error) {
	fileID := uuid.NewString()

	filePath := filepath.Join(uc.cfg.Path, fileID)
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Write(params.GetData())
	if err != nil {
		return nil, err
	}

	return &proto.SaveFileResponse{Uuid: fileID}, nil
}

func (uc *useCase) GetPublicURL(
	ctx context.Context,
	params *proto.GetPublicURLRequest,
) (result *proto.GetPublicURLResponse, err error) {
	filePath := filepath.Join(uc.cfg.Path, params.GetUuid())

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file with uuid %s does not exist", params.GetUuid())
	}

	return &proto.GetPublicURLResponse{
		Url: fmt.Sprintf(
			"http://%s/%s/%s",
			uc.cfg.HTTP.Address,
			uc.cfg.Path,
			params.GetUuid(),
		),
	}, nil
}
