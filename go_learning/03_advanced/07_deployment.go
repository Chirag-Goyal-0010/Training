package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 1. Environment Configuration
type Config struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	LogLevel string `json:"log_level"`
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// 2. Graceful Shutdown
func setupGracefulShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down server...")
		server.Close()
	}()
}

// 3. Health Check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// 4. Metrics Collection
type Metrics struct {
	RequestsTotal   int64
	RequestsFailed  int64
	ResponseTimeAvg float64
}

var metrics = &Metrics{}

func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		metrics.RequestsTotal++

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		metrics.ResponseTimeAvg = (metrics.ResponseTimeAvg*float64(metrics.RequestsTotal-1) + duration.Seconds()) / float64(metrics.RequestsTotal)
	}
}

// 5. Logging
type Logger struct {
	file *os.File
}

func NewLogger(path string) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{file: file}, nil
}

func (l *Logger) Log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
	l.file.WriteString(logEntry)
}

// 6. Docker Support
/*
# Dockerfile
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
*/

// 7. Kubernetes Support
/*
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
      - name: go-app
        image: go-app:latest
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
---
apiVersion: v1
kind: Service
metadata:
  name: go-app
spec:
  selector:
    app: go-app
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
*/

// 8. CI/CD Pipeline
/*
# .github/workflows/ci.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.16
    - name: Build
      run: go build -v ./...
    - name: Test
      run: go test -v ./...
    - name: Build Docker image
      run: docker build -t go-app:${{ github.sha }} .
*/

// 9. Environment Variables
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 10. Configuration Management
type AppConfig struct {
	Environment string
	Version     string
	Debug       bool
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Environment: getEnv("APP_ENV", "development"),
		Version:     getEnv("APP_VERSION", "1.0.0"),
		Debug:       getEnv("APP_DEBUG", "false") == "true",
	}
}

func main() {
	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Setup logger
	logger, err := NewLogger("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.file.Close()

	// Setup server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: setupRoutes(),
	}

	// Setup graceful shutdown
	setupGracefulShutdown(server)

	// Start server
	logger.Log("INFO", fmt.Sprintf("Server starting on %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())
}

func setupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthCheck)

	// Metrics endpoint
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(metrics)
	})

	// Main application endpoints
	mux.HandleFunc("/", metricsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))

	return mux
}

// 11. Monitoring and Alerting
type Alert struct {
	Level   string
	Message string
	Time    time.Time
}

func (a *Alert) Send() {
	// Implementation for sending alerts (e.g., email, Slack, etc.)
	log.Printf("ALERT: [%s] %s", a.Level, a.Message)
}

// 12. Backup and Recovery
func backupData() error {
	// Implementation for backing up data
	return nil
}

func restoreData() error {
	// Implementation for restoring data
	return nil
}

// 13. Load Balancing
type LoadBalancer struct {
	servers []string
	current int
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		current: 0,
	}
}

func (lb *LoadBalancer) NextServer() string {
	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

// 14. Service Discovery
type ServiceRegistry struct {
	services map[string][]string
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string][]string),
	}
}

func (sr *ServiceRegistry) Register(service, address string) {
	sr.services[service] = append(sr.services[service], address)
}

func (sr *ServiceRegistry) GetService(service string) []string {
	return sr.services[service]
}

// 15. Circuit Breaker
type CircuitBreaker struct {
	failures     int
	threshold    int
	resetTimeout time.Duration
	lastFailure  time.Time
}

func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold:    threshold,
		resetTimeout: resetTimeout,
	}
}

func (cb *CircuitBreaker) Execute(f func() error) error {
	if cb.failures >= cb.threshold {
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.failures = 0
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}

	err := f()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
	}
	return err
} 