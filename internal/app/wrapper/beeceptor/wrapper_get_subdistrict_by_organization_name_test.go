package beeceptor

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/weanan/weanan-service/config"
	"github.com/weanan/weanan-service/internal/pkg/logger"
	"github.com/weanan/weanan-service/internal/pkg/rest/mocks"
)

func TestGetSubdistrictByOrganizationName(t *testing.T) {
	cfg := config.BeeceptorConfig{}
	logger.NewLogger(logger.Option{IsEnable: true})

	getOrganizationSubdistrictFalseResp := GetOrganizationSubdistrictResponse{Success: false}
	getOrganizationSubdistrictFalseRespByte, _ := json.Marshal(getOrganizationSubdistrictFalseResp)

	getOrganizationSubdistrictSuccessResp := GetOrganizationSubdistrictResponse{Success: true}
	getOrganizationSubdistrictSuccessRespByte, _ := json.Marshal(getOrganizationSubdistrictSuccessResp)

	cases := map[string]struct {
		ShouldError       bool
		GetResponseBody   []byte
		GetResponseStatus int
		GetResponseErr    error
	}{
		"ShouldErrorWhenGetResponseError": {
			ShouldError:    true,
			GetResponseErr: errors.New("http client get error"),
		},
		"ShouldErrorWhenGetResponseStatusNot200": {
			ShouldError:       true,
			GetResponseStatus: 400,
		},
		"ShouldErrorWhenGetResponseBodyIsNil": {
			ShouldError:       true,
			GetResponseStatus: 200,
			GetResponseBody:   nil,
		},
		"ShouldErrorWhenGetResponseBodyStatusFalse": {
			ShouldError:       true,
			GetResponseStatus: 200,
			GetResponseBody:   getOrganizationSubdistrictFalseRespByte,
		},
		"ShouldSuccess": {
			ShouldError:       false,
			GetResponseStatus: 200,
			GetResponseBody:   getOrganizationSubdistrictSuccessRespByte,
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			httpClient := new(mocks.RestClient)

			w := NewWrapper().SetConfig(cfg).Setup()

			// inject client with mock
			w.client = httpClient
			wrapper := w.Validate()

			httpClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(test.GetResponseBody, test.GetResponseStatus, test.GetResponseErr).Once()

			resp, err := wrapper.GetSubdistrictByOrganizationName(context.Background(), "mamatosai")

			if test.ShouldError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
