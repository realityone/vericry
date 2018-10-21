package storage

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/realityone/vericry/pkg/config"
	"github.com/realityone/vericry/pkg/proto/schema"
)

var _ Storage = &MySQL{}

// MySQL is
type MySQL struct {
	db *sql.DB
}

// New is
func New() *MySQL {
	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return &MySQL{
		db: db,
	}
}

// Token is
func (m *MySQL) Token(id int64) (*schema.Token, error) {
	return nil, nil
}

// Stat is
func (m *MySQL) Stat(id int64) (*schema.Stat, error) {
	return nil, nil
}

// MinID is
func (m *MySQL) MinID() (int64, error) {
	return 0, nil
}

// MaxID is
func (m *MySQL) MaxID() (int64, error) {
	return 0, nil
}

// FromPlainText is
func (m *MySQL) FromPlainText(plain string) (*schema.Token, error) {
	return nil, nil
}
