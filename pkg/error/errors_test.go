package error_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"modak-rated-limited-challenge/pkg/error"
)

func TestHandlerError_Error(t *testing.T) {
	underTest := error.HandlerError{Cause: errors.New("err")}
	assert.Equal(t, "err", underTest.Error())
}

func TestHandlerNotFoundError_Error(t *testing.T) {
	underTest := error.NotFoundError{Cause: errors.New("err")}
	assert.Equal(t, "err", underTest.Error())
}
