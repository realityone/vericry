package storage

import (
	"context"

	"github.com/realityone/vericry/pkg/proto/model"
)

// Storage is
type Storage interface {
	Token(ctx context.Context, id int64) (*model.Token, error)
	Stat(ctx context.Context, id int64) (*model.Stat, error)

	MinID(ctx context.Context) (int64, error)
	MaxID(ctx context.Context) (int64, error)

	FromPlainText(ctx context.Context, plain string) (*model.Token, error)
}
