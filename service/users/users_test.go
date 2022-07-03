package users_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"deall-package/proto"
	users "deall-package/service/users"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type testServer struct {
	proto.UnimplementedUsersServiceServer
}

var tcpPort = ":9999"

func DialTestServer() {

	// gRPC server section
	lis, err := net.Listen("tcp", tcpPort)

	if err != nil {
		glog.Errorf("fail to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := users.Server{}

	proto.RegisterUsersServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func TestCreate(t *testing.T) {
	go DialTestServer()
	t.Log("get list users")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(tcpPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewUsersServiceClient(conn)

	response, err := c.Create(context.Background(), &proto.UserRequest{
		FirstName: "Tendi",
		Username:  "tendif",
		Password:  "hello",
	})

	if err != nil {
		log.Fatalf("Error when calling TestGetusers: %s", err)
	}

	fmt.Println(response.Message)
}
