package beeceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/weanan/weanan-service/internal/pkg/logger"
	"go.elastic.co/apm/v2"
)

func (w *wrapper) GetSubdistrictByOrganizationName(ctx context.Context, name string) (resp GetOrganizationSubdistrictResponse, err error) {
	span, ctx := apm.StartSpan(ctx, "wrapper.GetSubdistrictByOrganizationName", "custom")
	defer span.End()

	path := fmt.Sprintf("/organinzation/subdistrict?name=%s", name)

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("wrapper.GetSubdistrictByOrganizationName get error: %s", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("wrapper.GetSubdistrictByOrganizationName return status %d", status)
		logger.Log.Error(ctx, err.Error())
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("wrapper.GetSubdistrictByOrganizationName unmarshal error: %s", err.Error())
		logger.Log.Error(ctx, err.Error())
		return
	}

	if !resp.Success {
		err = fmt.Errorf("wrapper.GetSubdistrictByOrganizationName not success: %s", resp.Message)
		logger.Log.Error(ctx, err.Error())
		return
	}

	return
}
