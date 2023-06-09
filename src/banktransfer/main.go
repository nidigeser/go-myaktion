package main

import (
	"fmt"
	"net"
	"os"

	"github.com/nidigeser/go-myaktion/src/banktransfer/grpc/banktransfer"
	"github.com/nidigeser/go-myaktion/src/banktransfer/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("log level not specified, using default log level: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}

var grpcPort = 9111

func main() {
	log.Info("Starting Banktransfer service")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on grpc port %d: %v", grpcPort)
	}
	grpcServer := grpc.NewServer()
	banktransfer.RegisterBankTransferServer(grpcServer, service.NewBankTransferSerive())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port %d: %v", grpcPort, err)
	}
}
