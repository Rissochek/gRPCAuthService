package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"AuthProject/internal/auth"
	"AuthProject/internal/database"
	"AuthProject/internal/model"
	pb "AuthProject/internal/proto"
)

var (
	port = flag.Int("port", 8888, "server port")
	db = database.InitDataBase()
	manager = auth.JWTManager{TokenDuration: time.Minute * 10}
)

type server struct{
	pb.UnimplementedAuthServer
}

func(s *server) Registration(_ context.Context, reg_request *pb.RegistrationRequest) (*pb.RegistrationReply, error){
	user := model.User{Usermame: reg_request.Username, Password: reg_request.Password}
	err := database.AddUserToDataBase(db, &user)
	if err != nil{
		return &pb.RegistrationReply{Result: err.Error()}, err
	}
	return &pb.RegistrationReply{Result: "Success"}, nil
}

func(s *server) Login(ctx context.Context, log_request *pb.LoginRequest) (*pb.LoginReply, error){
	user := model.User{Usermame: log_request.Username, Password: log_request.Password}
	err := database.SearchUserInDB(db, &user)
	if err != nil{
		return &pb.LoginReply{Result: err.Error()}, err
	}

	token, err := manager.GenerateJWT(&user)
	if err != nil{
		return &pb.LoginReply{Result: err.Error()}, err
	}
	bearer_token := fmt.Sprintf("%s %s", "Bearer", token)
	return &pb.LoginReply{Result: "Success", Token: bearer_token}, nil
}

func main(){
	flag.Parse()
	manager.GetSecretKeyFromEnv()
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