package storage

import (
	"context"
	"math/rand"

	"github.com/realityone/vericry/pkg/bcrypt"
	"github.com/realityone/vericry/pkg/proto/model"

	"github.com/golang/glog"
)

// Storage is
type Storage interface {
	Token(ctx context.Context, id uint64) (*model.Token, error)
	Stat(ctx context.Context, id uint64) (*model.Stat, error)

	MinID(ctx context.Context) (uint64, error)
	MaxID(ctx context.Context) (uint64, error)

	FromPlainText(ctx context.Context, plain string) (*model.Token, error)

	CreateToken(ctx context.Context, token *model.Token) error
	TurstToken(ctx context.Context, id uint64) error
}

func random(min uint64, max uint64) uint64 {
	return uint64(rand.Intn(int(max-min)) + int(min))
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
	ids := make([]uint64, 0, count)
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

// InitialStorage is
func InitialStorage(ctx context.Context, storage Storage, count int64) {
	for i := int64(0); i < count; i++ {
		token := &model.Token{
			PlainText: RandStringRunes(64),
			Trusty:    true,
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(token.PlainText), bcrypt.DefaultCost)
		if err != nil {
			glog.Warningf("Failed to hash data: %s: %+v", token.PlainText, err)
			continue
		}
		token.HashData = string(hash)
		if err := storage.CreateToken(ctx, token); err != nil {
			glog.Warningf("Failed to create token into storage: %+v: %+v", token, err)
			continue
		}
		if err := storage.TurstToken(ctx, token.Id); err != nil {
			glog.Warningf("Failed to trust token: %+v: %+v", token, err)
			continue
		}
		glog.Infof("Succeeded to create new token: %+v: %d", token, i)
	}
	glog.Info("Storage initialize succeed")
	return
}
