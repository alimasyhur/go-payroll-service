package validator_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/pkg/mocks"
	"github.com/weanan/weanan-service/internal/pkg/validator"
)

func TestValidate(t *testing.T) {
	type payloadStruct struct {
		Name     string `json:"name" validate:"required"`
		Location string `json:"location" validate:"required"`
	}

	cases := map[string]struct {
		ShouldError bool
		Payload     interface{}
	}{
		"ShouldErrorWhenDoValidateButPayloadIsNil": {
			ShouldError: true,
			Payload:     nil,
		},
		"ShouldErrorWhenRequiredPayloadMissing": {
			ShouldError: true,
			Payload: payloadStruct{
				Name: "mamatosai",
			},
		},
		"ShouldSuccess": {
			ShouldError: false,
			Payload: payloadStruct{
				Name:     "mamatosai",
				Location: "tangerang",
			},
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			c, _ := mocks.MockEcho("POST", "/", http.Header{}, nil)

			err := validator.Validate(c, test.Payload)

			if test.ShouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
