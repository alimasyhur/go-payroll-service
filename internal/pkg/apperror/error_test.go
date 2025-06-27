package apperror_test

import (
	"errors"
	"testing"

	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

func TestApplicationError(t *testing.T) {
	err := apperror.New(404, errors.New("test error"))
	assert.NotNil(t, err)

	var appErr *apperror.ApplicationError
	assert.ErrorAs(t, err, &appErr)
	assert.Equal(t, err.Error(), "test error")
}
