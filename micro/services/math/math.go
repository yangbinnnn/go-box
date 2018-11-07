package math

import (
	"context"

	core "go-box/core"
)

const port = ":50051"

type Server struct{}

func (s *Server) Add(ctx context.Context, in *AddRequest) (*AddReply, error) {
	// delegate call
	sum := core.Add(int(in.A), int(in.B))
	return &AddReply{Sum: int64(sum)}, nil
}
