package middleware

import (
	"errors"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

var (
	grpcGatewayTag *opentracing.Tag
)

// init Initializes constant part of request ID
func init() {
	grpcGatewayTag = &opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}
}

// TracingWrapper is a middleware which parses the incoming HTTP headers to create a new span correctly
func TracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || errors.Is(err, opentracing.ErrSpanContextNotFound) {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext,
				// and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}
