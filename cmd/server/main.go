package main

import (
	"flag"
	"net"

	"github.com/realityone/vericry/pkg/config"
	"github.com/realityone/vericry/pkg/proto/api/v1"
	"github.com/realityone/vericry/pkg/server"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	config.Logging()
	glog.V(2).Info("Starting vericry server...")

	server := server.New()
	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterVericryServer(grpcServer, server)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
