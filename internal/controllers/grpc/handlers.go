package handlers

import (
	"example/api/proto/examplepb"
)

type Server struct {
	examplepb.UnimplementedExampleServiceServer
	service service.ServiceInterface
}

func NewControllers(s service.ServiceInterface) *Server {
	return &Server{service: s}
}
