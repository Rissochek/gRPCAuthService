package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"AuthProject/internal/database"
	"AuthProject/internal/model"
	pb "AuthProject/internal/proto"
)

var (
	port = flag.Int("port", 8888, "server port")
	db = database.InitDataBase()
)

type server struct{
	pb.UnimplementedAuthServer
}

func(s *server) Registration(_ context.Context, reg_request *pb.RegistrationRequest) (*pb.RegistrationReply, error){
	user := model.User{Usermame: reg_request.Username, Password: reg_request.Password}
	err := database.AddUserToDataBase(*db, &user)
	if err != nil{
		return &pb.RegistrationReply{Result: err.Error()}, err
	}
	return &pb.RegistrationReply{Result: "Success"}, nil
}

func(s *server) Login(_ context.Context, log_request *pb.LoginRequest) (*pb.LoginReply, error){
	user := model.User{Usermame: log_request.Username, Password: log_request.Password}
	err := database.SearchUserInDB(*db, &user)
	if err != nil{
		return &pb.LoginReply{Result: err.Error()}, err
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