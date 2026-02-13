package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "check-network/internal/domain"
)

type HealthHandler struct {
    version   string
    startTime time.Time
}

func NewHealthHandler(version string) *HealthHandler {
    return &HealthHandler{
        version:   version,
        startTime: time.Now(),
    }
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    uptime := time.Since(h.startTime)
    
    healthStatus := domain.NewHealthStatus(h.version, uptime)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    if err := json.NewEncoder(w).Encode(healthStatus); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}

func (h *HealthHandler) HealthCheckDetailed(w http.ResponseWriter, r *http.Request) {
    response := map[string]interface{}{
        "status": "healthy",
        "timestamp": time.Now(),
        "version": h.version,
        "uptime": time.Since(h.startTime).String(),
        "services": map[string]string{
            "database": "connected", 
            "cache":    "connected",
        },
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}