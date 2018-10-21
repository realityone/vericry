package storage

import (
	"testing"
)

var (
	mysql *MySQL
)

func init() {
	mysql = NewMySQL()
}

func TestToken(t *testing.T) {
	testToken(t, mysql, 1)
}

func TestStat(t *testing.T) {
	testStat(t, mysql, 1)
}
