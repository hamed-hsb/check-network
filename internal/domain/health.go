package domain

import "time"

type HealthStatus struct {
    Status    string    `json:"status"`
    Timestamp time.Time `json:"timestamp"`
    Version   string    `json:"version"`
    Uptime    string    `json:"uptime"`
}

func NewHealthStatus(version string, uptime time.Duration) HealthStatus {
    return HealthStatus{
        Status:    "healthy",
        Timestamp: time.Now(),
        Version:   version,
        Uptime:    uptime.String(),
    }
}