package response_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/alimasyhur/go-payroll-service/internal/pkg/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestResponseSuccess(t *testing.T) {
	cases := map[string]struct {
		Data interface{}
	}{
		"DataShouldNilWhenParamDataIsNil": {
			Data: nil,
		},
		"ResponseDataShouldBeSameWithParamData": {
			Data: map[string]interface{}{"name": "mamatosai"},
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			c, rec := mocks.MockEcho("GET", "/", http.Header{}, nil)
			err := response.ResponseSuccess(c, test.Data)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			resp := response.DefaultResponse{}
			err = json.Unmarshal(rec.Body.Bytes(), &resp)

			assert.NoError(t, err)
			assert.Equal(t, resp.Success, true)
			assert.Equal(t, resp.Message, "Success")
			assert.Equal(t, resp.Data, test.Data)
		})
	}
}
