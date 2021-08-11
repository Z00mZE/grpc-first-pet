package main

import (
	"google.golang.org/grpc"
	"grpc-first-pet/pkg/adder"
	"grpc-first-pet/pkg/api"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := new(adder.GRPCServer)
	api.RegisterAdderServer(s, srv)
	listener, listenerError := net.Listen("tcp", ":8080")
	if listenerError != nil {
		log.Fatal(listenerError)
	}

	if listenError:=s.Serve(listener);listenError!=nil{
		log.Fatal(listenError)
	}
}
