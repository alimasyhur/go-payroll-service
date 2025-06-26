package apperror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/pkg/apperror"
)

func TestApplicationError(t *testing.T) {
	err := apperror.New(404, errors.New("test error"))
	assert.NotNil(t, err)

	var appErr *apperror.ApplicationError
	assert.ErrorAs(t, err, &appErr)
	assert.Equal(t, err.Error(), "test error")
}
