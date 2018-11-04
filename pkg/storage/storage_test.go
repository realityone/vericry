package storage

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testToken(t *testing.T, s Storage, id uint64) {
	token, err := s.Token(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	log.Printf("token: %+v", token)
}

func testStat(t *testing.T, s Storage, id uint64) {
	stat, err := s.Stat(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, stat)
	log.Printf("stat: %+v", stat)
}
