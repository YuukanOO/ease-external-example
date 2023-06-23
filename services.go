package easeexternalexample

import "time"

type HealthCheckResponse struct {
	Status string    `json:"status"` // Status of services, ok if all good
	Time   time.Time `json:"time"`   // Server time (UTC)
}

// ease:api path=/api/_health
func HealthCheck() HealthCheckResponse {
	return HealthCheckResponse{
		Status: "ok",
		Time:   time.Now().UTC(),
	}
}
