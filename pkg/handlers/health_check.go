package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

// HealthCheckHandler returns the API status
type HealthCheckHandler struct {
	Config *viper.Viper
	Logger *zap.Logger
}

type status string

const (
	statusPass status = "pass"
)

type healthCheck struct {
	Status    status `json:"status"`
	Datetime  string `json:"datetime"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

// Handle handles a web request and returns a healthcheck JSON payload
func (h HealthCheckHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusReport := healthCheck{
			Status:    statusPass,
			Version:   h.Config.GetString("APPLICATION_VERSION"),
			Datetime:  h.Config.GetString("APPLICATION_DATETIME"),
			Timestamp: h.Config.GetString("APPLICATION_TS"),
		}
		js, err := json.Marshal(statusReport)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Failed to marshal health check: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(js)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Failed to write health check to response: %v", err))
			http.Error(w, "Failed to get health check", http.StatusInternalServerError)
			return
		}
	}
}