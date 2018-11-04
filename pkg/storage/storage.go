package storage

import (
	"context"
	"math/rand"

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

func random(min int64, max int64) int64 {
	return int64(rand.Intn(int(max-min)) + int(min))
}

// RandomTokens is
func RandomTokens(ctx context.Context, storage Storage, count int64) ([]*model.Token, error) {
	min, err := storage.MinID(ctx)
	if err != nil {
		return nil, err
	}
	max, err := storage.MaxID(ctx)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, count)
	for i := int64(0); i < count; i++ {
		ids = append(ids, random(min, max))
	}
	tokens := make([]*model.Token, 0, count)
	for _, id := range ids {
		t, err := storage.Token(ctx, id)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, nil
}
