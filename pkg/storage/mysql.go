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

// NewMySQL is
func NewMySQL() *MySQL {
	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return &MySQL{
		db: db,
	}
}

// Token is
func (m *MySQL) Token(ctx context.Context, id uint64) (*model.Token, error) {
	sqlt := `SELECT id,plain_text,hash_data,trusty FROM token WHERE id<=?`
	row := m.db.QueryRowContext(ctx, sqlt, id)
	token := &model.Token{}
	if err := row.Scan(&token.Id, &token.PlainText, &token.HashData, &token.Trusty); err != nil {
		return nil, errors.WithStack(err)
	}
	return token, nil
}

// Stat is
func (m *MySQL) Stat(ctx context.Context, id uint64) (*model.Stat, error) {
	sqlt := `SELECT id,succeed,failed FROM stat WHERE id=?`
	row := m.db.QueryRowContext(ctx, sqlt, id)
	stat := &model.Stat{}
	if err := row.Scan(&stat.Id, &stat.Succeed, &stat.Failed); err != nil {
		if err == sql.ErrNoRows {
			return &model.Stat{Id: id}, nil
		}
		return nil, errors.WithStack(err)
	}
	return stat, nil
}

// MinID is
func (m *MySQL) MinID(ctx context.Context) (uint64, error) {
	sql := `SELECT min(id) FROM token`
	row := m.db.QueryRowContext(ctx, sql)
	min := uint64(0)
	if err := row.Scan(&min); err != nil {
		return 0, errors.WithStack(err)
	}
	return min, nil
}

// MaxID is
func (m *MySQL) MaxID(ctx context.Context) (uint64, error) {
	sql := `SELECT max(id) FROM token`
	row := m.db.QueryRowContext(ctx, sql)
	max := uint64(0)
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

// CreateToken is
func (m *MySQL) CreateToken(ctx context.Context, token *model.Token) error {
	sql := `INSERT IGNORE INTO token (plain_text,hash_data,trusty) VALUES (?,?,?)`
	res, err := m.db.ExecContext(ctx, sql, token.PlainText, token.HashData, token.Trusty)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	token.Id = uint64(id)
	return nil
}

// TurstToken is
func (m *MySQL) TurstToken(ctx context.Context, id uint64) error {
	sql := `UPDATE token SET trusty_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := m.db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}
