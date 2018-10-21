package storage

import (
	"github.com/realityone/vericry/pkg/proto/schema"
)

// Storage is
type Storage interface {
	Token(id int64) (*schema.Token, error)
	Stat(id int64) (*schema.Stat, error)

	MinID() (int64, error)
	MaxID() (int64, error)

	FromPlainText(plain string) (*schema.Token, error)
}
