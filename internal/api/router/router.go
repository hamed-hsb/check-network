package router

import (
    "net/http"
    "check-network/internal/api/handlers"
    "check-network/pkg/logger"
)

type Router struct {
    healthHandler *handlers.HealthHandler
    logger        *logger.Logger
}

func New(healthHandler *handlers.HealthHandler, logger *logger.Logger) *Router {
    return &Router{
        healthHandler: healthHandler,
        logger:        logger,
    }
}

func (r *Router) middlewareLogging(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        r.logger.Info("Request received", "method", req.Method, "path", req.URL.Path)
        next(w, req)
    }
}

func (r *Router) SetupRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/health", r.middlewareLogging(r.healthHandler.HealthCheck))
    
    mux.HandleFunc("/health/detailed", r.middlewareLogging(r.healthHandler.HealthCheckDetailed))
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        w.Write([]byte("Welcome to Check Network API. Try /health or /health/detailed"))
    })
    
    return mux
}