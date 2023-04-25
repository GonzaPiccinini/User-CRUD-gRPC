package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/GonzaPiccinini/practice-grpc/user/userpb"
	"google.golang.org/grpc"
)

type server struct {
	userpb.UserServiceServer
}

func main() {
	// Setting logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Starting tcp connection
	fmt.Println("----- USER SERVICE -----")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	// Starting grpc server
	go func() {
		fmt.Println("Starting Server")
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// Wait for crtl + x to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until we get the signal
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Goodbye :D ...")
}
