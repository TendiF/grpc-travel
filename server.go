package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"deall-package/gateway"
	"deall-package/proto"
	"deall-package/service/customers"
	"deall-package/service/reguler"
	users "deall-package/service/users"
	"deall-package/utils/database"
	"deall-package/utils/utils"

	customerModel "deall-package/models/customer"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func serveHttp() {
	// gRPC gateway section
	API_PORT := os.Getenv("API_PORT")
	GRPC_PORT := os.Getenv("GRPC_PORT")

	if API_PORT == "" {
		API_PORT = "5001"
	}
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":" + API_PORT,
		GRPCServer: gateway.Endpoint{
			Network: "tcp",
			Addr:    "localhost:" + GRPC_PORT,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	godotenv.Load(".env")
	database.CreateConnection(3 * time.Second)
	customerModel.CreateIndexing()
	utils.InitAdmin()
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

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		utils.WithServerUnaryInterceptor(),
	)

	proto.RegisterUsersServiceServer(grpcServer, &users.Server{})
	proto.RegisterCustomersServiceServer(grpcServer, &customers.Server{})
	proto.RegisterRegulersServiceServer(grpcServer, &reguler.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Failed to serve gRPC server over port %: %v", GRPC_PORT, err)
	}
}
