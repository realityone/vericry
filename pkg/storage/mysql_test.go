package storage

import (
	"testing"
)

var (
	mysql *MySQL
)

func init() {
	mysql = New()
}

func TestToken(t *testing.T) {
	testToken(t, mysql, 1)
}
