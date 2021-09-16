package main

import (
	"fmt"
	"ig.engagements/log"
	"ig.engagements/server"
	"ig.engagements/service"
)

func main() {

	logger := log.Std

	svc := service.NewIGEngagementService(logger)

	srv := server.NewServerRouter(logger, svc)

	address := "127.0.0.1:3001"

	fmt.Printf("Server running on address: %s", address)
	if err := srv.Serve(address); err != nil {
		logger.WithError(err).Error("failed to start server.")
		panic(err)
	}
}
