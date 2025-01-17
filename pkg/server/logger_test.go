package server

import (
	"net/http"
	"net/http/httptest"

	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appcontext"
)

func (s *ServerTestSuite) TestLoggerMiddleware() {
	s.Run("get a new logger with trace ID", func() {

		req := httptest.NewRequest("GET", "/systems/", nil)
		rr := httptest.NewRecorder()
		traceMiddleware := NewTraceMiddleware(s.logger)
		prodLogger, err := zap.NewProduction()
		s.NoError(err)
		loggerMiddleware := NewLoggerMiddleware(prodLogger)

		// this is the actual test, since the context is cancelled post request
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger, ok := appcontext.Logger(r.Context())

			s.True(ok)
			s.NotEqual(prodLogger, logger)
		})

		traceMiddleware(loggerMiddleware(testHandler)).ServeHTTP(rr, req)
	})

	s.Run("get the same logger with no trace ID", func() {

		req := httptest.NewRequest("GET", "/systems/", nil)
		rr := httptest.NewRecorder()
		// need a new logger, because no-op won't use options
		prodLogger, err := zap.NewProduction()
		s.NoError(err)
		loggerMiddleware := NewLoggerMiddleware(prodLogger)

		// this is the actual test, since the context is cancelled post request
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger, ok := appcontext.Logger(r.Context())

			s.True(ok)
			s.Equal(prodLogger, logger)
		})

		loggerMiddleware(testHandler).ServeHTTP(rr, req)
	})
}
