package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "AuthProject/internal/proto"
	"AuthProject/internal/utils"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "server port")
	filename = flag.String("fn", "database", "File with user data")
)

type server struct{
	pb.UnimplementedAuthServer
}

func(s *server) Registration(_ context.Context, reg_request *pb.RegistrationRequest) (*pb.RegistrationReply, error){
	file := utils.OpenFile(filename)
	defer utils.CloseFile(file)
	utils.CreateUser(&reg_request.Username, &reg_request.Password, file)
	return &pb.RegistrationReply{Result: "Success"}, nil
}

func(s *server) Login(_ context.Context, log_request *pb.LoginRequest) (*pb.LoginReply, error){
	file := utils.OpenFile(filename)
	defer utils.CloseFile(file)
	err := utils.SearchInFile(&log_request.Username, &log_request.Password, file)
	if err != nil{
		return &pb.LoginReply{Result: "Fail"}, err
	}
	return &pb.LoginReply{Result: "Success"}, nil
}
func main(){
	flag.Parse()
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil{
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("Server listen on: %v", lsn.Addr())
	if err := s.Serve(lsn); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}