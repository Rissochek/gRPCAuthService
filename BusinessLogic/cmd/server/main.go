package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"AuthProject/source/auth"
	pb "AuthProject/source/proto"
)

var (
	port = flag.Int("port", 8889, "server port")
	manager = auth.JWTManager{}
)

type server struct{
	pb.UnimplementedBusinessServer
}

func(s *server) Hello(_ context.Context, hello_request *pb.HelloRequest) (*pb.HelloReply, error){
	claims, err := manager.VerifyJWT(hello_request.Authorization)
	if err != nil{
		return &pb.HelloReply{HelloName: "Failed to validate token"}, err
	}
	return &pb.HelloReply{HelloName: "Hello " + claims.Username}, nil
}

func main(){
	flag.Parse()
	manager.GetSecretKeyFromEnv()
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil{
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBusinessServer(s, &server{})
	log.Printf("Server listen on: %v", lsn.Addr())
	if err := s.Serve(lsn); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}