package main

import (
	"log"
	"microservice-blast-module/config/environment"
	"microservice-blast-module/internal/v1/handler"
	"strconv"
)

func main() {
	environment.LoadEnv()
	servicePort, err := strconv.Atoi(environment.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
	err = handler.StartGRPCServer(servicePort)
	if err != nil {
		log.Fatal("Failed to start gRPC server: ", err)
	}
}
