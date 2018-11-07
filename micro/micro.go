package micro

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	math "go-box/micro/services/math"
)

func RunGrpc(addr string) {
	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("grpc listen on %s\n", addr)
		// https://godoc.org/google.golang.org/grpc#Server.RegisterService
		s := grpc.NewServer()
		math.RegisterMathServer(s, &math.Server{})
		// https://godoc.org/google.golang.org/grpc/reflection
		reflection.Register(s)
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
