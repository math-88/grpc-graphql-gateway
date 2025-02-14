package runtime

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

type MiddlewareError struct {
	Code    string
	Message string
}

func (m *MiddlewareError) Error() string {
	return m.Message
}

func NewMiddlewareError(code, message string) *MiddlewareError {
	return &MiddlewareError{
		Code:    code,
		Message: message,
	}
}

// Cors is middelware function to provide CORS headers to response headers
func Cors() MiddlewareFunc {
	return func(ctx context.Context, serveMux *ServeMux, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		w.Header().Set("Access-Control-Allow-Origin", r.URL.Host)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "1728000")
		return ctx, nil
	}
}

func WithMetadata(annotator func(context.Context, *http.Request) metadata.MD) MiddlewareFunc {
	return func(ctx context.Context, serveMux *ServeMux, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		serveMux.metadataAnnotators = append(serveMux.metadataAnnotators, annotator)
		return ctx, nil
	}
}
