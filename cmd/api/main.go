package main

import (
    "context"
    "fmt"

    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    
    "check-network/internal/api/handlers"
    "check-network/internal/api/router"
    "check-network/internal/config"
    "check-network/pkg/logger"
)

func main() {
    cfg := config.Load()
    
    appLogger := logger.New()
    
    appLogger.Info("Starting HealthCheck API", "version", cfg.Version)
    
    healthHandler := handlers.NewHealthHandler(cfg.Version)
    
    apiRouter := router.New(healthHandler, appLogger)
    
    mux := apiRouter.SetupRoutes()
    
    server := &http.Server{
        Addr:         ":" + cfg.ServerPort,
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    
    go func() {
        appLogger.Info(fmt.Sprintf("Server starting on port %s", cfg.ServerPort))
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            appLogger.Error("Could not start server", "error", err)
            os.Exit(1)
        }
    }()
    
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    appLogger.Info("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        appLogger.Error("Server forced to shutdown", "error", err)
        os.Exit(1)
    }
    
    appLogger.Info("Server exited gracefully")
}