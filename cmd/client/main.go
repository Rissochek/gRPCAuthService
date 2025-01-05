package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "AuthProject/internal/proto"
)

var (
	addr = flag.String("address", "localhost:8080", "address to connect")
)

func main(){
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Registration(ctx, &pb.RegistrationRequest{Username: "rissochek", Password: "1234"})
	if err != nil{
		log.Fatalf("failed to registrate: %v", err)
	}
	log.Printf("reply from server: %v", r.GetResult())
	l, err := c.Login(ctx, &pb.LoginRequest{Username: "rissochek", Password: "1234"})
	log.Printf("reply from server: %v", l.GetResult())
	if err != nil{
		log.Fatalf("failed to login: %v", err)
	}
}