package beeceptor

import (
	"context"

	"github.com/weanan/weanan-service/config"
	"github.com/weanan/weanan-service/internal/pkg/rest"
)

type BeeceptorWrapper interface {
	GetSubdistrictByOrganizationName(ctx context.Context, name string) (resp GetOrganizationSubdistrictResponse, err error)
}

type wrapper struct {
	config config.BeeceptorConfig
	client rest.RestClient
}

func NewWrapper() *wrapper {
	return &wrapper{}
}

func (w *wrapper) SetConfig(config config.BeeceptorConfig) *wrapper {
	w.config = config
	return w
}

func (w *wrapper) Setup() *wrapper {
	restOptions := rest.Options{
		Address: w.config.Host,
		Timeout: w.config.Timeout,
		SkipTLS: w.config.SkipTLS,
	}

	w.client = rest.New(restOptions)
	return w
}

func (w *wrapper) Validate() BeeceptorWrapper {
	if w.client == nil {
		panic("client is nil")
	}

	return w
}
