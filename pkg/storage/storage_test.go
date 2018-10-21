package storage

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testToken(t *testing.T, s Storage, id int64) {
	token, err := s.Token(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	log.Printf("token: %+v", token)
}
