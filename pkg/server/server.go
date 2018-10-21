package server

import (
	"github.com/realityone/vericry/pkg/storage"
)

// Server is
type Server struct {
	storage storage.Storage
}

// New is
func New() *Server {
	return &Server{
		storage: storage.NewMySQL(),
	}
}
