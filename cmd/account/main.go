package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"final/api"
	"final/internal/biz"
	"final/internal/data"
	"final/internal/db"
	"final/internal/service/account"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50052, "port number")

func main() {
	flag.Parse()

	// database
	db.InitDb()

	// gRPC
	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// business logic starts here
	repo := data.GetRepository()
	userIf := biz.NewUserImpl(repo)
	api.RegisterAccountServer(s, &account.Server{UserIf: userIf})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
