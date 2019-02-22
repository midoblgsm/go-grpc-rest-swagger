package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) CreateObject(ctx context.Context, in *CreateObjectRequest) (*CreateObjectResponse, error) {
	log.Printf("Receive message %s", in.ObjectID)
	return &CreateObjectResponse{Status: "success"}, nil
}

func (s *Server) DeleteObject(ctx context.Context, in *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	log.Printf("Receive message %s", in.ObjectID)
	return &DeleteObjectResponse{Status: "success"}, nil
}
