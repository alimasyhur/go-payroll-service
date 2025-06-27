package health_check_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/health_check"
	mockPkg "github.com/alimasyhur/go-payroll-service/internal/pkg/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
)

func TestCheck(t *testing.T) {
	t.Run("ShouldSuccess", func(t *testing.T) {
		c, rec := mockPkg.MockEcho(http.MethodGet, "/", http.Header{}, nil)

		h := health_check.NewHandler().Validate()
		err := h.Check(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		resp := response.DefaultResponse{}
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, true, resp.Success)
	})
}
