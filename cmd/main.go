package main

import (
	"fmt"
	"log"

	"github.com/AlexandrKobalt/go-files/config"
	"github.com/gofiber/fiber/v2"

	"github.com/AlexandrKobalt/go-files/internal/server"
	serviceDeliveryGRPC "github.com/AlexandrKobalt/go-files/internal/service/delivery/grpc"
	serviceDeliveryHTTP "github.com/AlexandrKobalt/go-files/internal/service/delivery/http"
	serviceUseCase "github.com/AlexandrKobalt/go-files/internal/service/usecase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	newSystemUseCase := serviceUseCase.New(cfg)
	newSystemDeliveryGRPC := serviceDeliveryGRPC.New(newSystemUseCase)
	newSystemDeliveryHTTP := serviceDeliveryHTTP.New(cfg)

	app := fiber.New()

	app.Get(fmt.Sprintf("/%s/:key", cfg.Path), newSystemDeliveryHTTP.ServeFile())

	srv := server.New(
		cfg,
		newSystemDeliveryGRPC,
		app,
	)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
