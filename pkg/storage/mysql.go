package storage

import (
	"context"
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/realityone/vericry/pkg/config"
	"github.com/realityone/vericry/pkg/proto/model"
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
func (m *MySQL) Token(ctx context.Context, id int64) (*model.Token, error) {
	sql := `SELECT id,plain_text,hash_data,trusty FROM token WHERE id=?`
	row := m.db.QueryRowContext(ctx, sql, id)
	token := &model.Token{}
	if err := row.Scan(&token.Id, &token.PlainText, &token.HashData, &token.Trusty); err != nil {
		return nil, errors.WithStack(err)
	}
	return token, nil
}

// Stat is
func (m *MySQL) Stat(ctx context.Context, id int64) (*model.Stat, error) {
	return nil, nil
}

// MinID is
func (m *MySQL) MinID(ctx context.Context) (int64, error) {
	sql := `SELECT min(id) FROM token`
	row := m.db.QueryRowContext(ctx, sql)
	min := int64(0)
	if err := row.Scan(&min); err != nil {
		return 0, errors.WithStack(err)
	}
	return min, nil
}

// MaxID is
func (m *MySQL) MaxID(ctx context.Context) (int64, error) {
	sql := `SELECT max(id) FROM token`
	row := m.db.QueryRowContext(ctx, sql)
	max := int64(0)
	if err := row.Scan(&max); err != nil {
		return 0, errors.WithStack(err)
	}
	return max, nil
}

// FromPlainText is
func (m *MySQL) FromPlainText(ctx context.Context, plain string) (*model.Token, error) {
	sql := `SELECT id,plain_text,hash_data,trusty FROM token WHERE plain_text=?`
	row := m.db.QueryRowContext(ctx, sql, plain)
	token := &model.Token{}
	if err := row.Scan(&token.Id, &token.PlainText, &token.HashData, &token.Trusty); err != nil {
		return nil, errors.WithStack(err)
	}
	return token, nil
}
