package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

func main() {
	db := newInMemoryDB()
	svc := newService(db)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcserver := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcserver, svc)
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
