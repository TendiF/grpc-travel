package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"deall-package/gateway"
	"deall-package/proto"
	users "deall-package/service/users"
	"deall-package/utils"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String("endpoint", "localhost:9000", "endpoint of the gRPC service")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func serveHttp() {
	// gRPC gateway section
	API_PORT := os.Getenv("API_PORT")
	fmt.Println("API_PORT", API_PORT)
	if API_PORT == "" {
		API_PORT = "5001"
	}
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":" + API_PORT,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}

func main() {
	godotenv.Load(".env")

	GRPC_PORT := os.Getenv("GRPC_PORT")

	if GRPC_PORT == "" {
		GRPC_PORT = "5000"
	}

	go serveHttp()

	// gRPC server section
	lis, err := net.Listen("tcp", ":"+GRPC_PORT)

	if err != nil {
		glog.Errorf("fail to listen on port %s: %v", GRPC_PORT, err)
	}

	grpcServer := grpc.NewServer(
		utils.WithServerUnaryInterceptor(),
	)

	s := users.Server{}

	proto.RegisterUsersServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Failed to serve gRPC server over port %: %v", GRPC_PORT, err)
	}
}
