package config

import (
    "os"
   
)

type Config struct {
    ServerPort string
    Version    string
    StartTime  string
}

func Load() *Config {
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080"
    }

    version := os.Getenv("APP_VERSION")
    if version == "" {
        version = "1.0.0"
    }

    return &Config{
        ServerPort: port,
        Version:    version,
        StartTime:  os.Getenv("START_TIME"),
    }
}