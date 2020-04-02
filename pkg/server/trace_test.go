package server

import (
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"

	"github.com/cmsgov/easi-app/pkg/appcontext"
)

func (s ServerTestSuite) TestTraceMiddleware() {
	// this is the actual test, since the context is cancelled post request
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID, ok := appcontext.Trace(r.Context())

		s.True(ok)
		s.NotEqual(uuid.UUID{}, traceID)
	})

	req := httptest.NewRequest("GET", "/systems/", nil)
	rr := httptest.NewRecorder()
	middleware := NewTraceMiddleware(s.logger)

	middleware(testHandler).ServeHTTP(rr, req)
}
