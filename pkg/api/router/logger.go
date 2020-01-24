package router

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

// LoggingResponseWriter will encapsulate a standard ResponseWriter with a copy of its statusCode
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// ResponseWriterWrapper is supposed to capture statusCode from ResponseWriter
func ResponseWriterWrapper(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader is a surcharge of the ResponseWriter method
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logger is a gorilla/mux middleware to add log to the API
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := ResponseWriterWrapper(w)
		inner.ServeHTTP(wrapper, r)
		logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
		_ = logger.Log(
			"remote_address", r.RemoteAddr,
			"method", r.Method,
			"request_uri", r.RequestURI,
			"protocol", r.Proto,
			"status_code", wrapper.statusCode,
			"content_length", r.ContentLength,
			"user_agent", r.Header.Get("User-Agent"))
	})
}
