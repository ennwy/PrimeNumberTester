package server

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ennwy/prime_number_tester/internal/app"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

var _ app.Server = (*Server)(nil)

func New() *Server {
	server := &Server{
		App: fiber.New(),
	}

	server.App.Post("/primes", Primes)

	return server
}

func (s *Server) Start() error {
	addr := net.JoinHostPort(os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))
	log.Println("ADDR:", addr)
	return s.App.Listen(addr)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.App.Server().ShutdownWithContext(ctx)
}
