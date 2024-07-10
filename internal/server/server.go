package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexandrKobalt/go-files/config"
	serviceDeliveryGRPC "github.com/AlexandrKobalt/go-files/internal/service/delivery/grpc"
	"github.com/AlexandrKobalt/go-files/pkg/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type Server struct {
	cfg          *config.Config
	deliveryGRPC *serviceDeliveryGRPC.Server
	httpServer   *fiber.App
	grpcServer   *grpc.Server
}

func New(
	cfg *config.Config,
	deliveryGRPC *serviceDeliveryGRPC.Server,
	httpServer *fiber.App,
) *Server {
	return &Server{
		cfg:          cfg,
		deliveryGRPC: deliveryGRPC,
		httpServer:   httpServer,
		grpcServer:   grpc.NewServer(),
	}
}

func (s *Server) Run() error {
	log.Println("starting server...")

	go func() {
		if err := s.httpServer.Listen(s.cfg.HTTP.Address); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", s.cfg.GRPC.Address)
		if err != nil {
			log.Fatal(err)
		}

		if err := s.grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	proto.RegisterFileServiceServer(s.grpcServer, s.deliveryGRPC)

	log.Println("server started!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("stopping server...")

	err := s.httpServer.Shutdown()
	if err != nil {
		log.Printf("error on fiber shutdown: %s", err.Error())
	}

	log.Println("server stopped!")

	return nil
}
