package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/realityone/vericry/pkg/proto/api/v1"

	"github.com/stretchr/testify/assert"
)

var server = New()

func TestGetQuestion(t *testing.T) {
	reply, err := server.GetQuestion(context.Background(), &v1.QuestionRequest{})
	assert.NoError(t, err)
	fmt.Printf("Question reply: %+v\n", reply)
}
