package server

import (
	"context"

	"github.com/realityone/vericry/pkg/proto/api/v1"
	"github.com/realityone/vericry/pkg/storage"
)

// Server is
type Server struct {
	storage storage.Storage
}

var _ v1.VericryServer = &Server{}

// New is
func New() *Server {
	return &Server{
		storage: storage.NewMySQL(),
	}
}

// GetQuestion is
func (s *Server) GetQuestion(ctx context.Context, req *v1.QuestionRequest) (*v1.QuestionReply, error) {
	tokens, err := storage.RandomTokens(ctx, s.storage, 3)
	if err != nil {
		return nil, err
	}
	ptokens := make([]*v1.TokenReply, 0, len(tokens))
	for _, t := range tokens {
		pt := &v1.TokenReply{
			PlainText: t.PlainText,
		}
		ptokens = append(ptokens, pt)
	}
	return &v1.QuestionReply{Tokens: ptokens}, nil
}

// Validate is
func (s *Server) Validate(ctx context.Context, req *v1.ValidateRequest) (*v1.ValidateReply, error) {
	for plainText, hashData := range req.Answer {
		token, err := s.storage.FromPlainText(ctx, plainText)
		if err != nil {
			return &v1.ValidateReply{Valid: false}, nil
		}
		if token.HashData != hashData {
			return &v1.ValidateReply{Valid: false}, nil
		}
	}
	return &v1.ValidateReply{Valid: true}, nil
}
