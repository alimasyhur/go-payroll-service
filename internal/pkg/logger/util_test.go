package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

func TestIsSkipLog(t *testing.T) {
	cases := map[string]struct {
		ContentType   string
		ShouldSkipLog bool
	}{
		"ShouldSkipLogWhenContentTypeIsApplication/tar+gzip": {
			ContentType:   "application/tar+gzip",
			ShouldSkipLog: true,
		},
		"ShouldSkipLogWhenContentTypeIsApplication/octet-stream": {
			ContentType:   "application/tar+gzip",
			ShouldSkipLog: true,
		},
		"ShouldSkipLogWhenContentTypeIsMultipart/form-data": {
			ContentType:   "multipart/form-data",
			ShouldSkipLog: true,
		},
		"ShouldNotSkipLogWhenContentTypeIsApplication/json": {
			ContentType:   "application/json",
			ShouldSkipLog: false,
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			isSkip := logger.IsSkipLog(test.ContentType)
			assert.Equal(t, test.ShouldSkipLog, isSkip)
		})
	}
}
