package test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "go-box/micro/services/math"

	"google.golang.org/grpc"
)

func TestGrpcMath(t *testing.T) {
	// https://godoc.org/google.golang.org/grpc#WithInsecure
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dail error: %v", err)
	}
	defer conn.Close()
	c := pb.NewMathClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{A: 1, B: 4})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("sum: %d", r.Sum)

}
