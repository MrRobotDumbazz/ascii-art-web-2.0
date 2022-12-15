package main

import (
	"ascii-art-web/delivery"
	"ascii-art-web/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := new(server.Server)
	go func() {
		if err := server.Start(":8080", delivery.Handlers()); err != nil {
			log.Println(err)
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Print(err)
		return
	}
}
