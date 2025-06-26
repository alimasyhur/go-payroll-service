package rest

import (
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
)

const (
	// TraceparentHeader is the HTTP header for trace propagation.
	//
	// For backwards compatibility, this is currently an alias for
	// for ElasticTraceparentHeader, but the more specific constants
	// below should be preferred. In a future version this will be
	// replaced by the standard W3C header.
	TraceparentHeader = ElasticTraceparentHeader

	// ElasticTraceparentHeader is the legacy HTTP header for trace propagation,
	// maintained for backwards compatibility with older agents.
	ElasticTraceparentHeader = "Elastic-Apm-Traceparent"

	// W3CTraceparentHeader is the standard W3C Trace-Context HTTP
	// header for trace propagation.
	W3CTraceparentHeader = "Traceparent"

	// TracestateHeader is the standard W3C Trace-Context HTTP header
	// for vendor-specific trace propagation.
	TracestateHeader = "Tracestate"
)

func PopulateTraceparentHeadersFromAPMContext(ctx apm.TraceContext) (headers map[string]string) {
	headers = make(map[string]string)
	headerValue := apmhttp.FormatTraceparentHeader(ctx)

	headers[ElasticTraceparentHeader] = headerValue
	headers[W3CTraceparentHeader] = headerValue

	if tracestate := ctx.State.String(); tracestate != "" {
		headers[TracestateHeader] = tracestate
	}

	return
}
