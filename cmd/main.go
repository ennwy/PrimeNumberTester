package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/ennwy/prime_number_tester/internal/app"
	"github.com/ennwy/prime_number_tester/internal/server"
)

func main() {
	var s app.Server = server.New()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := s.Stop(ctx); err != nil {
			log.Fatalln("server stop:", err)
		} else {
			log.Println("server has stopped successfully")
			return
		}
	}()

	log.Println("server has started")

	if err := s.Start(); err != nil {
		log.Println("server start:", err)
	}
}
